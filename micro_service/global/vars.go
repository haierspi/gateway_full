package global

import (
    "log"

    "github.com/streadway/amqp"

    "github.com/haierspi/gateway_full/utils/config"
    "github.com/haierspi/gateway_full/utils/rpc"

    // justifying
    _ "github.com/go-sql-driver/mysql"
    _ "github.com/lib/pq"
)

var (
    MqConn    *amqp.Connection
    RpcClient *rpc.Client
    // osmMysql    *osm.Osm
    // osmPostgres *osm.Osm
)

func init() {
    log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)
    // osmPack.ShowSQL = true
    var err error

    mqURL := config.String("./config.cfg", "mq", "url")
    MqConn, err = amqp.Dial(mqURL)

    RpcClient = rpc.NewClientWithConn(MqConn, mqURL)
    if err != nil {
        panic(err)
    }

}
