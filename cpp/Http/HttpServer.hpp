#pragma once
#include "TcpServer.hpp"

constexpr int PORT = 8080;

static void HandleRequest(int accept_socket){
    //处理请求，首先我们要获得请求
    std::string str;
    std::cout << "------------begin-------------" << std::endl;
    while(Util::ReadLine(accept_socket, str)){
        std::cout << str << std::endl;
        str.clear();
    }
    std::cout << "------------end---------------" << std::endl;
}

class HttpServer{
public:
    HttpServer(int port = PORT)
    :_port(port)
    {
        tcp_server = TcpServer::GetInstance(port);
    }

    //开始进行监听
    void Loop(){
        int listen_socket = tcp_server->GetSocket();
        while(true){
            sockaddr_in client_server;
            socklen_t client_len = (socklen_t)sizeof(client_server);
            int accept_socket = accept(listen_socket, (sockaddr*)&client_server, &client_len);
            if(accept_socket < 0){
                perror("accept fail");
                exit(SocketAcceptFaail);
            }
            std::thread t1(HandleRequest, accept_socket);
            t1.detach();
        }
    }
private:
    int _port;
    std::shared_ptr<TcpServer> tcp_server;
};