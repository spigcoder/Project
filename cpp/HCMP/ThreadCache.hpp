#pragma once
#include "FreeList.hpp"
#include "MemorySize.hpp"

class ThreadCache{
public:
    void* Allocate(size_t size);
    void  DeAllocate(void* ptr, size_t size);
    void* FetchFromCentralCache(size_t index, size_t mem_num);
private:
    FreeList _free_list[MAX_LIST_NUM];
};

//tls(thread local storage)每个线程都会有一份独立的，这种变量是不用加锁访问的
static __thread ThreadCache* tls_thread_cache = nullptr;

void* ThreadCache::Allocate(size_t size){
    assert(size<MAX_BYTE);
    //首先获取要分贝得到多少内存
    size_t mem_num = GetMemNum(size);
    size_t index = GetIndex(size);
    if(!_free_list[index].empty()){
        return _free_list[index].pop();
    }else{
        //从central cache中获取内存
        FetchFromCentralCache(index, mem_num);
    }
}

void ThreadCache::DeAllocate(void* ptr, size_t size){
    assert(ptr);
    assert(size<MAX_BYTE);

    size_t index = GetIndex(size);
    _free_list[index].push(ptr);
}