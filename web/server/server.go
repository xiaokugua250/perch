package main

import (
	_ "expvar"
	"net/http"
	"perch/api/chain_api"
	"perch/internal/p2p/blockchain"
	"perch/web/service"
)

// 路由信息

func main() {
	service.WebService{

		Name: "BlockP2P",

		Router: []service.WebRouter{
			{RouterPath: "/blockchain/", RouterHandlerFunc: chain_api.GetBlockChain, RouterMethod: http.MethodGet},
			{RouterPath: "/", RouterHandlerFunc: chain_api.CreateBlockChain, RouterMethod: http.MethodPost},
		},
		InitFunc: []func() error{
			blockchain.GenInitialChainBlock,
		},
	}.WebServiceStart()

}
