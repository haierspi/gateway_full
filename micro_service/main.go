package main

import (
    "log"

    "github.com/haierspi/gateway_full/micro_service/global"
    "github.com/haierspi/gateway_full/micro_service/services"
    "github.com/haierspi/gateway_full/utils/rpc"

    // justifying
    _ "github.com/go-sql-driver/mysql"
    _ "github.com/lib/pq"
)

// StartServer 开启服务
func main() {
    server := rpc.NewServer()
    server.Register(&services.Examples{})

    log.Println("examples_1.0 started...")
    server.ServeConn(global.MqConn, "examples_1.0")
}
