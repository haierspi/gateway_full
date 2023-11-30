package services

import (
    "log"

    "github.com/haierspi/gateway_full/utils/rpc"
)

// Examples 示例微服务
type Examples struct{}

// Echo 直接返回数据
func (Examples) Echo(args rpc.BodyArgs, reply *rpc.CommonReply) error {

    // _, _, err := osmMysql.InsertBySQL(`INSERT INTO examples_echo (content) VALUES(#{Countent});`, args.Body)
    // if err != nil {
    // 	log.Println(err)
    // }

    // _, _, err = osmPostgres.InsertBySQL(`INSERT INTO examples_echo (content) VALUES(#{Countent});`, args.Body)
    // if err != nil {
    // 	log.Println(err)
    // }

    // reply.Data = args.Body

    // helloReply := map[string]string{}

    // err := global.RpcClient.Call("examples_1.0", "Examples.Hello", map[string]string{"Name": args.Body}, &helloReply)
    // if err != nil {
    //     log.Println(err)
    // }
    reply.Data = "11111"

    return nil
}

// EchoBody 直接返回数据
func (Examples) EchoBody(args rpc.BodyArgs, reply *rpc.BodyReply) error {

    reply.Body = []byte("<xml>" + args.Body + "</xml>")
    reply.ContentType = "text/xml"
    log.Println(args.Body)

    return nil
}

// Hello 添加Hello
func (Examples) Hello(args struct{ Name string }, reply *struct{ Result string }) error {

    reply.Result = "Hello, " + args.Name

    return nil
}
