#include <iostream>
#include "HttpServer.hpp"

int main(int argc, char* argv[]){
    if(argc != 2){
        std::cout << argv[0] << " need two argument" << std::endl;
        exit(-1);
    }
    int port = atoi(argv[1]);
    HttpServer http_server(port);
    http_server.Loop();
}

