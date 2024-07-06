//目的是建立一个定长内存池，然后自行来进行分配和释放
#pragma once
#include <iostream>
using std::cout;
using std::endl;

const size_t DEFAULTSIZE = 1024*128;

template <class T>
class ObjectPool{
public:
    T* New(){
        T* obj = nullptr;
        size_t T_size = sizeof(T);
        //首先再free_list中寻找空间
        if(_free_list){
            void* next = *(void**)_free_list;
            obj = (T*)_free_list;
            _free_list = next;
        }
        else{
            //这时才是使用memory来进行操作
            if(_remain_size < T_size){
                //证明是第一次
                _remain_size = DEFAULTSIZE;
                _memory = (char*)malloc(_remain_size);
            }
            //这里要从空间中进行抽取
            size_t obj_size = T_size < sizeof(void*) ? sizeof(void*) : T_size;
            obj = (T*)_memory;
            _memory += obj_size;
            _remain_size -= obj_size;
        }

        //定位new调用这个指针的构造函数
        new(obj)T;
        return obj;
    } 

    void Delete(T* obj){
        //调用析构函数
        obj->~T();
        //我们对free_list的管理方法就是将前面一个指针的大小存放下一块空间的地址
        //头插
        //这里*(void**)是取出一个地址大小的空间，用来存放下一块空间的地址，因为32和64位环境下地址的大小不同
        *(void**)obj = _free_list;
        _free_list = obj;
    }

private:
    char* _memory = nullptr;          //开辟的整块大的空间
    size_t _remain_size = 0;    //剩余的空间的大小
    void* _free_list = nullptr;       //释放的空间用链表链接起来进行管理
};