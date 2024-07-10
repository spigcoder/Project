#pragma once
#include "Common.hpp"
#include "PageCache.hpp"
#include "ThreadCache.hpp"

static void* ConcurrentAlloc(size_t size){
    //如果要取的空间大小大于156kb，就不能使用threadCache来处理了，要用其他的
    if(size > MAX_BYTE){
        size_t page_size = size>>PAGE_SHIFT;
        PageCache::GetInstance()->_page_mul.lock();
        Span* span = PageCache::GetInstance()->NewSpan(page_size);
        PageCache::GetInstance()->_page_mul.unlock();
        void* ptr = (void*)(span->_page_id<<PAGE_SHIFT);
        return ptr;
    }else{
        if(tls_thread_cache == nullptr){
            tls_thread_cache = new ThreadCache;
        }
        void* ptr = tls_thread_cache->Allocate(size);
        return ptr;
    }
}

static void ConcurrentFree(void* ptr){
    Span* span = PageCache::GetInstance()->MapObjectToSpan(ptr);
    size_t size = span->_obj_size;
    if(size < MAX_BYTE){
        assert(tls_thread_cache);
        tls_thread_cache->DeAllocate(ptr, size);
    }else{
        //直接将内存交给PageCache
        Span* span = PageCache::GetInstance()->MapObjectToSpan(ptr);
        PageCache::GetInstance()->ReleaseSpanToPageCache(span);
    }
}
