#include "CenctralCache.hpp"
//static成员变量要在内外进行定义，为了避免包含头文件是重复定义所以又新建一个.cc文件
CentralCache CentralCache::_cen_cache_ins;

Span* GetOpenSpan(SpanList& list, size_t size){
    //...
    return nullptr;
}

//从已有的Span中获取batch_num个大小为size的对象，将他们以链表的形式放在start和end上面，
//因为最终的数量可能不够，所以最后要返回实际获得到的数量
size_t CentralCache::FetchRangeObj(void*& start, void*& end, size_t batch_num, size_t size){
   size_t index = GetIndex(size);

    //加上桶锁保证可以进行并发访问
    _span_list[index].mul.lock();
   Span* span = GetOpenSpan(_span_list[index], size);
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
    _span_list[index].mul.unlock();
    return actual_num;
}