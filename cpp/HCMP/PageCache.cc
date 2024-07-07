#include "PageCache.hpp"

PageCache PageCache::_page_ins;

//将大的Page切分为小的Page
Span* PageCache::SlicePage(size_t i, size_t k){
    Span* n_span = _page_list[i].PopFront();
    Span* k_span = new Span; 

    k_span->_n = k;
    k_span->_page_id = n_span->_page_id;

    n_span->_n -= k;
    n_span->_page_id += k;

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
            return SlicePage(i, k);
        }
    }

    //走到这里证明整个page cache都没有空间，那么就要向系统申请一个大小为128页的空间
    Span* big_span = new Span;
    void* ptr = SysAlloc(MAX_PAGES-1);
    big_span->_page_id = (page_id)ptr>>PAGE_SHIFT;
    big_span->_n = MAX_PAGES-1;

    _page_list[big_span->_n].PushFront(big_span);

    return SlicePage(MAX_PAGES-1, k);

}