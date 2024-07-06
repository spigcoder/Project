#pragma once
#include <assert.h>
#include <iostream>

static const size_t MAX_LIST_NUM = 208;
static const size_t MAX_BYTE = 256*1024;

static void*& NextNode(void* list){
    return *(void**)list;
}