#pragma once
#include "Span.hpp"

//thread 可以由多个线程同时访问，但是Central最好是搞成一个，大家都可以访问, 这样就可以使用单例模式
//单例模式就是说整个程序中只有一个实例，构造函数和拷贝构造都私有化，通过一个GetInstance来进行成员函数的调用
class CentralCache{
    static CentralCache* GetInstance(){return &_cen_cache_ins;}
    size_t FetchRangeObj(void*& start, void*& end, size_t batchNum, size_t size);
    Span* GetOpenSpan(SpanList& list, size_t size);

private:
    SpanList _span_list[MAX_LIST_NUM];
private:
    static CentralCache _cen_cache_ins;
    CentralCache(){}
    CentralCache(const CentralCache&) = delete;
};