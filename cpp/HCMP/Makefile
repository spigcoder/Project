CC = g++
CFLAGS = -std=c++11 -O0

.PHONY:all
all: HCMP

HCMP: test.cc CentralCache.o  PageCache.o ThreadCache.o
	$(CC) $(CFLAGS) -o HCMP -g test.cc ThreadCache.o CentralCache.o  PageCache.o


CentralCache.o: CentralCache.cc CentralCache.hpp
	$(CC) $(CFLAGS) -c CentralCache.cc -g

ThreadCache.o: ThreadCache.cc ThreadCache.hpp
	$(CC) $(CFLAGS) -c ThreadCache.cc -g

PageCache.o:PageCache.cc PageCache.hpp
	$(CC) $(CFLAGS) -c PageCache.cc -g

.PHONY:clean
clean:
	rm -f *.o HCMP

