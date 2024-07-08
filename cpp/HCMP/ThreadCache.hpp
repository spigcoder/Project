#pragma once
#include "FreeList.hpp"
#include "MemorySize.hpp"
#include "CentralCache.hpp"
#include "Common.hpp"

class ThreadCache{
public:
    void* Allocate(size_t size);
    void  DeAllocate(void* ptr, size_t size);
    void* FetchFromCentralCache(size_t index, size_t mem_num);
    void ListTooLong(FreeList& list, size_t size);
private:
    FreeList _free_list[MAX_LIST_NUM];
};

static __thread ThreadCache* tls_thread_cache = nullptr;