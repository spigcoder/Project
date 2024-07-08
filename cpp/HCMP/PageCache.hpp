#pragma once
#include "Span.hpp"
#include <unordered_map>

//同样的PageCache也应该采用单例模式
class PageCache{
public:
    static PageCache* GetInstance(){
        return &_page_ins;
    }
    Span* MapObjectToSpan(void* obj);
    Span* NewSpan(size_t k);
    Span* SlicePage(size_t i, size_t k);
    void ReleaseSpanToPageCache(Span* span);

public:
    std::mutex _page_mul;
private:
    std::unordered_map<page_id, Span*> _page_span_map;
    PageCache(){}
    PageCache(const PageCache&) = delete;
    SpanList _page_list[MAX_PAGES];
    static PageCache _page_ins;
};