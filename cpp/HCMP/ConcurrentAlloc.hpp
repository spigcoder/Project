#pragma once
#include "ThreadCache.hpp"
#include "Common.hpp"

static void* ConcurrentAlloc(size_t size){
    if(tls_thread_cache == nullptr){
        tls_thread_cache = new ThreadCache;
    }
    void* ptr = tls_thread_cache->Allocate(size);
    return ptr;
}

static void ConcurrentFree(void* ptr, size_t size){
    assert(tls_thread_cache);

    tls_thread_cache->DeAllocate(ptr, size);
}
