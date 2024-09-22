#pragma once
#include <mutex>
#include <memory>
#include <thread>
#include <string>
#include <stdio.h>
#include <cstring>
#include <unistd.h>
#include <stdlib.h>
#include <iostream>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>

enum{
    SocketCreateFail = 1,
    SocketBindFail,
    SocketListenFail,
    SocketAcceptFaail,
};


class Util{
public:
    //目的是处理不同系统的HTTP请求
    static int ReadLine(int accept_socket, std::string& line){
        //这个字符是随意给的
        char ch = 'x';
        while(ch != '\n'){
            ssize_t s = recv(accept_socket, &ch, 1, 0);
            if(s > 0){
                if (s == '\r'){
                    //只窥探，不进行修改
                    recv(accept_socket, &ch, 1, MSG_PEEK);
                    if(ch == '\n')
                        recv(accept_socket, &ch, 1, 0);
                    else    
                        ch = '\n';
                }
                //普通字符
                line += ch;
            }else if(0 == s){
                return 0;
            }else {
                return -1;
            }
        }
        return line.size();
    }
};