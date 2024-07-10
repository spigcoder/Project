#pragma once
//来计算对应size分配的空间的大小还有哈希桶的下标
#include "Common.hpp"

inline static void* SysAlloc(size_t k_page){
    #ifdef _WIN32
    	void* ptr = VirtualAlloc(0, k_page << PAGE_SHIFT, MEM_COMMIT | MEM_RESERVE, PAGE_READWRITE);
    #else
    	// linux下brk mmap等
        void *ptr = mmap(NULL, k_page<<PAGE_SHIFT, PROT_READ | PROT_WRITE, MAP_PRIVATE | MAP_ANONYMOUS, -1, 0);
    #endif

    	if (ptr == (void*)-1)
            perror("can't open a address");
    	return ptr;
}

inline static void SystemFree(void* ptr, size_t size) {
#ifdef _WIN32 
    VirtualFree(ptr, 0, MEM_RELEASE); #else
#else
    munmap(ptr, size);
// sbrk unmmap等 
#endif
}

//区间的划分
// 整体控制在最多10%左右的内碎⽚浪费 
// [1,128]                  8byte对⻬        freelist[0,16) 
// [128+1,1024]             16byte对⻬       freelist[16,72)
// [1024+1,8*1024]          128byte对⻬      freelist[72,128) 
// [8*1024+1,64*1024]       1024byte对⻬     freelist[128,184)     
// [64*1024+1,256*1024]     8*1024byte对⻬   freelist[184,208) 

//其实这里的思路就是找到第一个大于等于size的大小，每两个空间之间的距离是align_num
// static inline size_t _GetMemNum(size_t size, size_t align_num){
//     if(size % align_num == 0){
//         return size;
//     }else{
//         //找出最接近size的可开辟空间
//         return (size/align_num+1)*align_num;
//     }
// }

//还有一种使用位运算的写法也可以处理这个问题,这两种解决方法最后得到的结果是一样的
static inline size_t _GetMemNum(size_t size, size_t align_num){
    return ((size+(align_num-1))&(~(align_num-1)));
}

static inline size_t GetMemNum(size_t size){
    if(size <= 128){
        return _GetMemNum(size, 8);
    }else if(size <= 1024){
        return _GetMemNum(size, 16);
    }else if(size <= 8*1024){
        return _GetMemNum(size, 128);
    }else if(size <= 64*1024){
        return _GetMemNum(size, 1024);
    }else if(size <= 256*1024){
        return _GetMemNum(size, 8*1024);
    }else{
        std::cout << "size is big than MAX_BYTE" << std::endl;
        return -1;
    }
}

static inline size_t _GetIndex(size_t size, size_t align_num){
    //就是算出当前这个大小，在align_num位间隔的在哪个区间中
    if(size % align_num == 0){
        return size/align_num-1;
    }else{
        return size/align_num;
    }
}

//同样的可以通过位运算的方法来进行计算操作,这里的align_shift是2的多少次方
// static inline size_t _GetIndex(size_t size, size_t align_shift){
//     return ((size+((1<<align_shift)-1))>>align_shift)-1;
// }

//得到free_list的下标，就是上面对应的free_list后面的数字
static inline size_t GetIndex(size_t size){
   static size_t space[2] = {16, 56}; 
   if(size <= 128){
        return _GetIndex(size, 8);
   }else if(size <= 1024){
        return _GetIndex(size-128, 16)+space[0];
   }else if(size <= 8*1024){
        return _GetIndex(size-1024, 128)+space[0]+space[1];
   }else if(size <= 64*1024){
        return _GetIndex(size-8*1024, 1024)+space[0]+space[1]*2;
   }else if(size <= 256*1024){
        return _GetIndex(size-64*1024, 8*1024)+space[0]+space[1]*3;
   }
   return -1;
}

//这里计算的是应该从central cache中获取多少个mem_num对象
static size_t MemMoveSize(size_t mem_num){
    int n = MAX_BYTE / mem_num;
    if(n < 2) return 2;
    if(n > 512) return 512;
    return n;
}

//这里是计算的是应该从page cache中获取几页对象
static size_t PageMoveSize(size_t mem_num){
    size_t n = MemMoveSize(mem_num);
    size_t size = n*mem_num;
    //得到这个大小的内存是多少页
    size >>= PAGE_SHIFT;
    if(size < 1) size = 1;
    return size;
}




