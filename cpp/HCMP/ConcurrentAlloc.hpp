#pragma once
#include "ThreadCache.hpp"
#include "Common.hpp"

static void* ConcurrentAlloc(size_t size){
    if(tls_thread_cache == nullptr){
        tls_thread_cache = new ThreadCache;
    }

    return tls_thread_cache->Allocate(size);
}

static void ConcurrentFree(void* ptr, size_t size){
    assert(tls_thread_cache);

    tls_thread_cache->DeAllocate(ptr, size);
}