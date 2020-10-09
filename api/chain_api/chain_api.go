package chain_api

import (
	"context"
	"net/http"
	"perch/internal/p2p/blockchain"
	"perch/web/metric"
	"perch/web/model"
)

func CreateBlockChain(writer http.ResponseWriter, request *http.Request) {
	metric.ProcessMetricFunc(writer, request, nil, func(ctx context.Context, bean interface{}, respone *model.ResultReponse) error {

		newBlock, err := blockchain.GenerateBlock(blockchain.BlockChain[len(blockchain.BlockChain)-1], 5)
		if err != nil {
			//respone.ErrMsg= err
			return err
		}
		if blockchain.IsBlockValied(newBlock, blockchain.BlockChain[len(blockchain.BlockChain)-1]) {
			newBlockChain := append(blockchain.BlockChain, newBlock)
			blockchain.ReplaceChain(newBlockChain)
		}
		respone.Spec = newBlock
		respone.Kind = "blockchain"
		respone.Code = http.StatusOK

		return nil
	})

}

func GetBlockChain(writer http.ResponseWriter, request *http.Request) {
	metric.ProcessMetricFunc(writer, request, nil, func(ctx context.Context, bean interface{}, respone *model.ResultReponse) error {
		respone.Spec = blockchain.BlockChain
		respone.Total = len(blockchain.BlockChain)
		respone.Kind = "blockchain"
		respone.ErrMsg = nil
		respone.Code = http.StatusOK
		return nil
	})
}
