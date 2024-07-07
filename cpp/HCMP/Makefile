CC = g++
CFLAGS = -std=c++11

.PHONY:all
all: HCMP

HCMP: test.o CentralCache.o  PageCache.o
	$(CC) $(CFLAGS) -o HCMP test.o CentralCache.o  PageCache.o

test.o: test.cc
	$(CC) $(CFLAGS) -c test.cc

CentralCache.o: CentralCache.cc CentralCache.hpp
	$(CC) $(CFLAGS) -c CentralCache.cc



PageCache.o:PageCache.cc PageCache.hpp
	$(CC) $(CFLAGS) -c PageCache.cc

.PHONY:clean
clean:
	rm -f *.o HCMP

