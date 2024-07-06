/*包含一些都会用到的api*/
#pragma once
#include "Common.hpp"
//表示桶的最大下标

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
        _free_list = NextNode(_free_list);
        return obj;
    }

    bool empty(){
        return _free_list == nullptr;
    }

private:
    void* _free_list = nullptr;
};

