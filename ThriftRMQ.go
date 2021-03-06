package main

import (
	"github.com/BinArchitecture/GoRocketmqSender/rocketmq"
	"flag"
	"runtime"
	"git.apache.org/thrift.git/lib/go/thrift"
	"os"
	"github.com/BinArchitecture/GoRocketmqSender/rmq"
	"github.com/golang/glog"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()
	prod,er:=rocketmq.NewRoutingPoolProducer(50000,500000,"test1Group", "10.6.30.141:9876","prodInstance")
	if er != nil {
		panic(er)
	}
	transportFactory := thrift.NewTTransportFactory()
	protocolFactory := thrift.NewTBinaryProtocolFactory(true, true)
	serverTransport, err := thrift.NewTServerSocket("10.6.30.141:7912")
	if err != nil {
		glog.Errorf("Error%v!\n", err)
		os.Exit(1)
	}
	handler := &rmq.RmqThriftProdServiceImpl{
		prod,
	}
	handler.Start()
	var processor =rmq.NewRmqThriftProdServiceProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	glog.Info("thrift server start...")
	server.Serve()
}
