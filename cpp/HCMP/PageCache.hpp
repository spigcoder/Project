#pragma once
#include "Common.hpp"
#include "Span.hpp"

//同样的PageCache也应该采用单例模式
class PageCache{
public:
    static PageCache* GetInstance(){
        return &_page_ins;
    }

    Span* NewSpan(size_t k);
    Span* SlicePage(size_t i, size_t k);

public:
    std::mutex _page_mul;
private:
    PageCache(){}
    PageCache(const PageCache&) = delete;
    SpanList _page_list[MAX_PAGES];
    static PageCache _page_ins;
};