#pragma once
#include <mutex>
#include <thread>
#include <assert.h>
#include <iostream>
#include <algorithm>
#include <unistd.h>
#include <sys/mman.h>

using std::cout;
using std::endl;

static const size_t MAX_LIST_NUM = 208;    //thread cache 和 central cache的桶数
static const size_t MAX_BYTE = 256*1024;   //允许申请的最大字节数
static const size_t MAX_PAGES = 129;       //最大的页数
static const size_t PAGE_SHIFT = 13;       //2^13

#if __WORDSIZE == 64
    using page_id = unsigned long long;
#elif __WORDSIZE == 32
    using page_id = size_t;
#endif

inline static void* SysAlloc(size_t k_page){
    #ifdef _WIN32
    	void* ptr = VirtualAlloc(0, k_page << PAGE_SHIFT, MEM_COMMIT | MEM_RESERVE, PAGE_READWRITE);
    #else
    	// linux下brk mmap等
        void *ptr = mmap(NULL, k_page<<PAGE_SHIFT, PROT_READ | PROT_WRITE, MAP_PRIVATE | MAP_ANONYMOUS, -1, 0);
    #endif

    	if (ptr == nullptr)
    		throw std::bad_alloc();

    	return ptr;
}



static void*& NextNode(void* list){
    return *(void**)list;
}