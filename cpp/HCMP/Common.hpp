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
static const size_t PAGE_SHIFT = 12;       //2^12 也就是4KB

#if __WORDSIZE == 64
    using page_id = unsigned long long;
#elif __WORDSIZE == 32
    using page_id = size_t;
#endif



static void*& NextNode(void* list){
    return *(void**)list;
}