package main

import (
	"github.com/BinArchitecture/GoRocketmqSender/rocketmq"
	"flag"
	"runtime"
	"git.apache.org/thrift.git/lib/go/thrift"
	"fmt"
	"os"
	"github.com/BinArchitecture/GoRocketmqSender/rmq"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()
	producer, err := rocketmq.NewDefaultProducer("test1Group", "10.6.30.109:9876","prodInstance")
	if err != nil {
		panic(err)
	}
	prod,er:=rocketmq.NewRoutingProducer(producer,1000)
	if er != nil {
		panic(er)
	}
	//transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	transportFactory := thrift.NewTTransportFactory()
	protocolFactory := thrift.NewTBinaryProtocolFactory(true, true)
	serverTransport, err := thrift.NewTServerSocket("10.6.30.109:7912")
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}
	handler := &rmq.RmqThriftProdServiceImpl{
		prod,
	}
	handler.Start()
	var processor =rmq.NewRmqThriftProdServiceProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("thrift server start...")
	server.Serve()
}
