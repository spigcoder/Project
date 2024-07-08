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
        _sz++;
    }

    void InsertRange(void* start, void* end, size_t sz){
        _sz += sz;
        *(void**)end = _free_list;
        _free_list = start;
    }

    void* pop(){
        assert(_free_list);
        _sz--;
        void* obj = _free_list;
        _free_list = NextNode(_free_list);
        return obj;
    }

    //从链表中删除n个元素并且分别将start 和 end储存在这两个指针当中
    void PopRange(void*& start, void*& end, size_t n){
        assert(n <= _sz);
        start = _free_list;
        for(int i = 0; i < n-1; ++i){
            _free_list = NextNode(_free_list);
        }
         end = _free_list;
         _free_list = NextNode(_free_list);
         NextNode(end) = nullptr;
         _sz -= n;
    }

    size_t& Size(){ return _sz; }
    size_t& MaxSize(){ return max_size; }
    bool empty(){ return _free_list == nullptr; }

private:
    void* _free_list = nullptr;
    size_t _sz = 0;
    size_t max_size = 1;
};

