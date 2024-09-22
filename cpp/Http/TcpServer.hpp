#pragma once
#include "common.hpp"


constexpr int BACKLOG = 10;

class TcpServer{
public:
    void Socket(){
        _listen_socket = socket(AF_INET, SOCK_STREAM, 0);
        if (_listen_socket < 0){
            perror("socket:");
            exit(SocketCreateFail);
        }
        int opt = 1;
        //设置，让其在time_wait状态也可以重现建立其链接
        setsockopt(_listen_socket, SOL_SOCKET, SO_REUSEADDR, &opt, sizeof(opt));
    }

    void Listen(){
        if(listen(_listen_socket, BACKLOG) == -1){
            perror("listen fail");
            exit(SocketListenFail);
        }
    }

    void Bind(){
        struct sockaddr_in local;
        memset(&local, 0, sizeof(local));
        local.sin_family = AF_INET;
        local.sin_port = htons(_port);
        local.sin_addr.s_addr = INADDR_ANY;

        if(bind(_listen_socket, (sockaddr*)&local, sizeof(local)) == -1){
            perror("socket bind fail");
            exit(SocketBindFail);
        }
    }

    void SockInit(){
        Socket();
        Bind();
        Listen();
    }

    static std::shared_ptr<TcpServer> GetInstance(int port){
        static pthread_mutex_t lock = PTHREAD_MUTEX_INITIALIZER;
        //使用锁来对单例模式进行保护
        mtx.lock();
        if(nullptr == _instance){
            _instance = std::shared_ptr<TcpServer>(new TcpServer(port));
            _instance->SockInit();
        }
        mtx.unlock();
        return _instance;
    }

    int GetSocket(){
        return _listen_socket;
    }

private:
    TcpServer(int port)
    :_port(port),
    _listen_socket(-1)
    {}

    TcpServer(const TcpServer& server) = delete;
    TcpServer& operator=(const TcpServer& server) = delete;

private:
    int _port;
    int _listen_socket;
    inline static std::mutex mtx;
    inline static std::shared_ptr<TcpServer> _instance = nullptr;
};