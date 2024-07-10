#include "ConcurrentAlloc.hpp"
#include <vector>
#include <atomic>
// void Alloc1()
// {
// 	for (size_t i = 0; i < 5; ++i)
// 	{
// 		void* ptr = ConcurrentAlloc(6);
// 	}
// }

// void Alloc2()
// {
// 	for (size_t i = 0; i < 5; ++i)
// 	{
// 		void* ptr = ConcurrentAlloc(7);
// 	}
// }

// void TLSTest()
// {
// 	std::thread t1(Alloc1);
// 	t1.join();

// }
// void TestConcurrentAlloc1()
// {
// 	void* p1 = ConcurrentAlloc(6);
// 	void* p2 = ConcurrentAlloc(8);
// 	void* p3 = ConcurrentAlloc(1);
// 	void* p4 = ConcurrentAlloc(7);
// 	void* p5 = ConcurrentAlloc(8);
// 	void* p6 = ConcurrentAlloc(8);
// 	void* p7 = ConcurrentAlloc(8);
// 	void* p8 = ConcurrentAlloc(8);


// 	cout << p1 << endl;
// 	cout << p2 << endl;
// 	cout << p3 << endl;
// 	cout << p4 << endl;
// 	cout << p5 << endl;

// 	ConcurrentFree(p1);
// 	ConcurrentFree(p2);
// 	ConcurrentFree(p3);
// 	ConcurrentFree(p4);
// 	ConcurrentFree(p5);
// 	ConcurrentFree(p6);
// 	ConcurrentFree(p7);
// 	ConcurrentFree(p8);
// }

// void TestConcurrentAlloc2()
// {
// 	for (size_t i = 0; i < 1024; ++i)
// 	{
// 		void* p1 = ConcurrentAlloc(6);
// 		cout << p1 << endl;
// 	}

// 	void* p2 = ConcurrentAlloc(8);
// 	cout << p2 << endl;
// }


// void MultiThreadAlloc1()
// {
// 	std::vector<void*> v;
// 	for (size_t i = 0; i < 7; ++i)
// 	{
// 		void* ptr = ConcurrentAlloc(6);
// 		v.push_back(ptr);
// 	}

// 	for (auto e : v)
// 	{
// 		ConcurrentFree(e);
// 	}
// }

// void MultiThreadAlloc2()
// {
// 	std::vector<void*> v;
// 	for (size_t i = 0; i < 7; ++i)
// 	{
// 		void* ptr = ConcurrentAlloc(16);
// 		v.push_back(ptr);
// 	}

// 	for (auto e : v)
// 	{
// 		ConcurrentFree(e);
// 	}
// }

// void TestMultiThread()
// {
// 	std::thread t1(MultiThreadAlloc1);
// 	std::thread t2(MultiThreadAlloc2);

// 	t1.join();
// 	t2.join();
// }

// int main()
// {
// 	//TestObjectPool();
// 	//TLSTest();

// 	TestConcurrentAlloc1();
// 	//TestAddressShift();

// 	// TestMultiThread();

// 	return 0;
// }

void BenchmarkMalloc(size_t ntimes, size_t nworks, size_t rounds)
{
	std::vector<std::thread> vthread(nworks);
	std::atomic<size_t> malloc_costtime(0) ;
	std::atomic<size_t> free_costtime(0);

	for (size_t k = 0; k < nworks; ++k)
	{
		vthread[k] = std::thread([&, k]() {
			std::vector<void*> v;
			v.reserve(ntimes);

			for (size_t j = 0; j < rounds; ++j)
			{
				size_t begin1 = clock();
				for (size_t i = 0; i < ntimes; i++)
				{
					v.push_back(malloc(16));
					// v.push_back(malloc((16 + i) % 8192 + 1));
				}
				size_t end1 = clock();

				size_t begin2 = clock();
				for (size_t i = 0; i < ntimes; i++)
				{
					free(v[i]);
				}
				size_t end2 = clock();
				v.clear();

				malloc_costtime += (end1 - begin1);
				free_costtime += (end2 - begin2);
			}
		});
	}

	for (auto& t : vthread)
	{
		t.join();
	}

	printf("%lu个线程并发执行%u轮次，每轮次malloc %u次: 花费：%u ms\n",
		nworks, rounds, ntimes, malloc_costtime.load());

	printf("%lu个线程并发执行%u轮次，每轮次free %u次: 花费：%u ms\n",
		nworks, rounds, ntimes, free_costtime.load());

	printf("%lu个线程并发malloc&free %u次，总计花费：%u ms\n",
		nworks, nworks*rounds*ntimes, malloc_costtime.load() + free_costtime.load());
}


// 单轮次申请释放次数 线程数 轮次
void BenchmarkConcurrentMalloc(size_t ntimes, size_t nworks, size_t rounds)
{
	std::vector<std::thread> vthread(nworks);
	std::atomic<size_t> malloc_costtime(0);
	std::atomic<size_t> free_costtime(0);

	for (size_t k = 0; k < nworks; ++k)
	{
		vthread[k] = std::thread([&]() {
			std::vector<void*> v;
			v.reserve(ntimes);

			for (size_t j = 0; j < rounds; ++j)
			{
				size_t begin1 = clock();
				for (size_t i = 0; i < ntimes; i++)
				{
					v.push_back(ConcurrentAlloc(16));
					// v.push_back(ConcurrentAlloc((16 + i) % 8192 + 1));
				}
				size_t end1 = clock();

				size_t begin2 = clock();
				for (size_t i = 0; i < ntimes; i++)
				{
					ConcurrentFree(v[i]);
				}
				size_t end2 = clock();
				v.clear();

				malloc_costtime += (end1 - begin1);
				free_costtime += (end2 - begin2);
			}
		});
	}

	for (auto& t : vthread)
	{
		t.join();
	}

	printf("%lu个线程并发执行%u轮次，每轮次concurrent alloc %u次: 花费：%u ms\n",
		nworks, rounds, ntimes, malloc_costtime.load());

	printf("%lu个线程并发执行%u轮次，每轮次concurrent dealloc %u次: 花费：%u ms\n",
		nworks, rounds, ntimes, free_costtime.load());

	printf("%lu个线程并发concurrent alloc&dealloc %u次，总计花费：%u ms\n",
		nworks, nworks*rounds*ntimes, malloc_costtime.load() + free_costtime.load());
}

int main()
{
	size_t n = 1000;
	cout << "==========================================================" << endl;
	BenchmarkConcurrentMalloc(n, 4, 10);
	cout << endl << endl;

	BenchmarkMalloc(n, 4, 10);
	cout << "==========================================================" << endl;

	return 0;
}