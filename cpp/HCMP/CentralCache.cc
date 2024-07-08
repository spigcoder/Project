#include "CentralCache.hpp"

//static成员变量要在内外进行定义，为了避免包含头文件是重复定义所以又新建一个.cc文件
CentralCache CentralCache::_cen_cache_ins;

Span* CentralCache::GetOpenSpan(SpanList& span_list, size_t mem_num){
    //从给出的pan_list中寻找有没有空余的span_list，如果有就进行返回
    //否则就像page cache中申请
    Span* it = span_list.Begin();
    while(it != span_list.End()){
        if(it->_free_list != nullptr){
            return it;
        }else{
            it = it->_next;
        }
    }
    //走到这里证明现在已有的span_list中已经没有span了，那么我们就要向page_cache中申请内存

    //这里可以经这个桶的桶锁解开，这样如果有线程想要进行内存的释放就可以进行
    span_list.mul.unlock();

    //获取页时也要进行加锁操作
    PageCache::GetInstance()->_page_mul.lock();
    //从page cache 中得到一个span
    Span* span = PageCache::GetInstance()->NewSpan(PageMoveSize(mem_num));
    span->_is_use = true;
    PageCache::GetInstance()->_page_mul.unlock();

    //将则个span的地址拆分为一个一个大小为size_的free_list
    char* start = (char*)((span->_page_id)<<PAGE_SHIFT);
    cout << "start add is: " << (void*)start << endl;
    size_t size = (span->_n)<<PAGE_SHIFT;
    span->_free_list = start;
    void* tail = start;
    char* end = start + size; 
    start+= mem_num;
    //开始进行切分操作
    int i = 0;
    while(tail != end && start != end){
        i++;
        NextNode(tail) = start;
        start += mem_num;
        tail = NextNode(tail);
    }
    //将span插入到队列当中，这个过程是需要加锁的
    span_list.mul.lock();
    span_list.PushFront(span);
    return span;
}

//从已有的Span中获取batch_num个大小为size的对象，将他们以链表的形式放在start和end上面，
//因为最终的数量可能不够，所以最后要返回实际获得到的数量
size_t CentralCache::FetchRangeObj(void*& start, void*& end, size_t batch_num, size_t mem_num){
   //这里的size实际上传递的是mem_num但是都是可以通过GetIndex来获得下标的
    size_t index = GetIndex(mem_num);

    //加上桶锁保证可以进行并发访问
    _span_list[index].mul.lock();
    Span* span = GetOpenSpan(_span_list[index], mem_num);
    assert(span); 
    //这里我们已经得到了一个span，现在可以中span的free_list中获取我们想要的数据了
    start = span->_free_list;
    end = span->_free_list;
    int i = 0, actual_num = 1;
    while(i < batch_num-1 && NextNode(end)){
         i++, actual_num++;
         end = NextNode(end);
    }
    //处理操作
    span->_free_list = NextNode(end);
    NextNode(end) = nullptr;
    span->_use_count += actual_num;
    cout << span->_use_count << endl;
    _span_list[index].mul.unlock();

    return actual_num;
}

/*我们现在有一个链表以及上面的各个节点了，现在我们要把这个链表中的节点转移到spanlist
的各个span的free_list上面去, 那么有一个问题，就是我们怎么知道哪个块属于哪一个span呢
可以使用unorder_map来解决这个问题，在page cache中建立一个unordered_map来进行page_id和
span的一个映射，只要我们知道了一个块的起始地址，就知道它是哪一个page进而知道它是哪一个
span就可以进行操作了*/
void CentralCache::ReleaseListToSpans(void* start, size_t size){
    size_t index = GetIndex(size);
    _span_list[index].mul.lock();
    while(start){
        //因为central cache是多线程共同使用的，所以对他进行操作时要上锁
        //先将start的下一个位置保存起来
        void* next = NextNode(start);
        Span* span = PageCache::GetInstance()->MapObjectToSpan(start);
        //找到span之后将自己头插到span的free_list节点当中
        NextNode(start) = span->_free_list;
        span->_free_list = start;
        span->_use_count--;
        //当_usecount减为0时将它还给page cache
        if(span->_use_count == 0){
           //这里证明这个span划分出去的所有的小的单元块都已经换回来了，那么他的_free_list
           //是不是乱序已经不重要了，因为全都在，所以相当于拥有一块大的内存空间
           span->_free_list = nullptr;
           _span_list[index].Erase(span); 

           //这时可以将桶锁给解开了，因为span已经不能被外界访问到了，
           _span_list[index].mul.unlock();
           //这里将span还给Page Cache
           PageCache::GetInstance()->_page_mul.lock();
           PageCache::GetInstance()->ReleaseSpanToPageCache(span);
           PageCache::GetInstance()->_page_mul.unlock();
           _span_list[index].mul.lock();
        }
        start = next; 
    }
    _span_list[index].mul.unlock();
}















