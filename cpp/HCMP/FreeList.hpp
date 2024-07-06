/*包含一些都会用到的api*/
#pragma once
#include <assert.h>
#include <iostream>
//表示桶的最大下标

static void*& NextList(void* list){
    return *(void**)list;
}

class FreeList{
public:
    void push(void* obj){
        assert(obj);
        //头插
        *(void**)obj = _free_list;
        _free_list = obj;
    }

    void* pop(){
        assert(_free_list);

        void* obj = _free_list;
        _free_list = NextList(_free_list);
        return obj;
    }

    bool empty(){
        return _free_list == nullptr;
    }

private:
    void* _free_list = nullptr;
};

