/*Span是Central Cache的核心，Central也是和Thread Cache一样的桶装结构，只是他所挂的是一个一个的Span
Span中有各种各样的结构，由于有时要对空间进行合并，所以使用双向链表比单项链表更加的合理*/
#pragma once
#include <mutex>
#include "Common.hpp"
#include "MemorySize.hpp"

#if __WORDSIZE == 64
    using page_id = unsigned long long;
#elif __WORDSIZE == 32
    using page_id = size_t;
#endif

struct Span
{
    page_id _page_id = 0;       //大块内存起始页的页号
    size_t _n = 0;              //一共有多少页
    
    Span* _next = nullptr;      //双向指针的成员函数
    Span* _prev = nullptr;
    void* _free_list = nullptr; //这个span有对应大小的一个free_list
    
    size_t _obj_size = 0;       //要切割的内容的大小
    size_t _use_count = 0;      //正在使用的小块内存
};

//每一个SpanList都是一个带头双向循环链表
class SpanList{
public:
    SpanList(){
        _head = new Span;
        _head->_next = _head;
        _head->_prev = _head;
    }

    void Insert(Span* pos, Span* new_span){
        assert(new_span);
        //prev new pos
        Span* prev = pos->_prev;

        prev->_next = new_span;
        new_span->_prev = prev;
        new_span->_next = pos;
        pos->_prev = new_span;
    }

    void Erase(Span* pos){
        assert(pos);
        assert(pos == _head);
        Span* next = pos->_next;
        Span* prev = pos->_prev;

        next->_prev = prev;
        prev->_next = next;

        delete pos;
    }

private:
    Span* _head = nullptr;
public:
    std::mutex mul; 
};













