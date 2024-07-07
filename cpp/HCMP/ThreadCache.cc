#include "ThreadCache.hpp"

void*  ThreadCache::FetchFromCentralCache(size_t index, size_t mem_num){
    void *start = nullptr, *end = nullptr;
    size_t& max_size = _free_list[index].MaxSize();
    size_t batch_num = std::min(max_size, MemMoveSize(mem_num));
    //如果和max_size的到的数据量相同，就扩大max_size
    if(batch_num == max_size) max_size++;
    size_t actual_num = CentralCache::GetInstance()->FetchRangeObj(start, end, batch_num, mem_num); 
    assert(actual_num > 0);
    if(actual_num == 1){
        //证明最终只得到了一个对象，那么将这个对象返回即可
        assert(start == end);
    }else{
        //证明不知有一个对象，要把多余的对象连接到链表当中去
        _free_list[index].InsertRange(NextNode(start), end);
    }
    return start;
}

void* ThreadCache::Allocate(size_t size){
    assert(size<MAX_BYTE);
    //首先获取要分贝得到多少内存
    size_t mem_num = GetMemNum(size);
    size_t index = GetIndex(size);
    if(!_free_list[index].empty()){
        return _free_list[index].pop();
    }else{
        //从central cache中获取内存
        return FetchFromCentralCache(index, mem_num);
    }
}

void ThreadCache::DeAllocate(void* ptr, size_t size){
    assert(ptr);
    assert(size<MAX_BYTE);

    size_t index = GetIndex(size);
    _free_list[index].push(ptr);
}