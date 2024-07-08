#include "PageCache.hpp"

PageCache PageCache::_page_ins;

//将大的Page切分为小的Page
Span* PageCache::SlicePage(size_t i, size_t k){
    Span* n_span = _page_list[i].PopFront();
    Span* k_span = new Span; 

    k_span->_n = k;
    cout << "n_span ptr add is: " << (void*)(n_span->_page_id<<PAGE_SHIFT) <<endl; 
    k_span->_page_id = n_span->_page_id;
    //这里要建立映射关系，对于要返回的span，每一页都要建立对应的span和page_id的关系
    for(int i = 0; i < k; ++i){
        _page_span_map[k_span->_page_id+i] = k_span;
    }

    n_span->_n -= k;
    n_span->_page_id += k;
    //对于n_span只用让头和唯页指向这个span即可
    _page_span_map[n_span->_page_id] = n_span;
    _page_span_map[n_span->_page_id+n_span->_n-1] = n_span;

    _page_list[n_span->_n].PushFront(n_span);
    return k_span;
}

//获取一个新的span还给central cache, page cache的桶结构和central cache
//的结构不一样，他是以页数来进行悬挂的，而且如果当前桶没有，他会向更大的桶进行申请，然后把更大的桶进行切分
//将需要的大小返回，将切剩下的
Span* PageCache::NewSpan(size_t k){
    if(!_page_list[k].Empty()){
        //这里证明当前Page_List中有Page,将第一个page返回jike
        return _page_list[k].PopFront();
    } 
    //走到这里证明当前桶中是没有这么大的元素的，那么我们可以找更大的元素进行切分操作
    for(int i = k+1; i < MAX_PAGES; ++i){
        if(!_page_list[i].Empty()){
            cout << i << endl;
            Span* span = SlicePage(i, k);
            cout << "ptr add is: " << (void*)(span->_page_id<<PAGE_SHIFT) <<endl; 
            return span;
        }
    }

    //走到这里证明整个page cache都没有空间，那么就要向系统申请一个大小为128页的空间
    Span* big_span = new Span;
    void* ptr = SysAlloc(MAX_PAGES-1);
    cout << "ptr add is: " << ptr << endl;
    big_span->_page_id = (page_id)ptr>>PAGE_SHIFT;
    big_span->_n = MAX_PAGES-1;

    _page_list[big_span->_n].PushFront(big_span);

    Span* span = SlicePage(MAX_PAGES-1, k);
    return span;

}

/*根据映射表，返回该指针对应的页对应的span*/
Span* PageCache::MapObjectToSpan(void* obj){
    assert(obj);
    page_id id = (page_id)obj>>PAGE_SHIFT;
    auto it = _page_span_map.find(id);
    if(it == _page_span_map.end()){
        //没有找到
        assert(false);
        return nullptr;
    }else{
        //找到了
        return it->second;
    }
}

/*现在这个span已经换回来了，但是有一个问题需要解决，就是如果把span直接挂在span_list的化
就会出现外碎片问题,所以我们要把相邻的span进行合并：怎么合并-> 使用page_id和每个span的大小来判断是否相邻
以及这个span是否正在被使用*/
void PageCache::ReleaseSpanToPageCache(Span* span){
    //向前合并
    while(true){
        page_id prev_page = span->_page_id-1;
        auto it = _page_span_map.find(prev_page);
        //这里证明没有前一个，可以退出
        if(it == _page_span_map.end()) {
            break;
        }
        Span* prev_span = it->second;
        //这里证明前一个正在使用，退出
        if(prev_span->_is_use == true){
            break;
        }
        //这里证明两个页相加大于最大的页，也不进行操作
        if(prev_span->_n + span->_n > MAX_PAGES-1){
            break; 
        }
        //走到这里证明既有前一个span，而且满足合并的条件
        _page_list[prev_span->_n].Erase(prev_span);
        span->_n += prev_span->_n;
        span->_page_id = prev_span->_page_id;
        delete prev_span;
    }

    //向后合并
    while(true){
        page_id next_page = span->_page_id+span->_n;
        auto it = _page_span_map.find(next_page);
        //这里证明没有前一个，可以退出
        if(it == _page_span_map.end()) {
            break;
        }
        Span* next_span = it->second;
        //这里证明前一个正在使用，退出
        if(next_span->_is_use == true){
            break;
        }
        //这里证明两个页相加大于最大的页，也不进行操作
        if(next_span->_n + span->_n > MAX_PAGES-1){
            break; 
        }
        //走到这里证明既有前一个span，而且满足合并的条件
        _page_list[next_span->_n].Erase(next_span);
        span->_n += next_span->_n;
        delete next_span;
    }
    //要把合并后的span挂起来
    _page_list[span->_n].PushFront(span);
    span->_is_use = false;
    //把这个span也映射到表里面，方便后面的span进行合并
    _page_span_map[span->_page_id] = span;
    _page_span_map[span->_page_id+span->_n-1] = span;
}














