package withGETH

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"log"

	"github.com/ethereum/go-ethereum"
	_ "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type IService interface {
	SubNewHeads()
	Monitoring()
	UnSubNewHeads()
	StopMonitor()
}

type service struct {
	client      *ethclient.Client
	sub         ethereum.Subscription
	headers     chan *types.Header
	url         string
	monitorAddr string
	quitChan    chan int
	repo        IRepository
}

func NewService(url string, monitorAddr string, repo IRepository) IService {
	return &service{url: url, monitorAddr: monitorAddr, repo: repo}
}

func (mon *service) SubNewHeads() {
	client, err := ethclient.Dial(mon.url)
	if err != nil {
		log.Panicf(err.Error())
	}

	mon.client = client

	mon.headers = make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), mon.headers)
	if err != nil {
		log.Panicf(err.Error())
	}
	mon.sub = sub
	mon.quitChan = make(chan int)
}

func (mon *service) Monitoring() {
	fmt.Printf("start monitoring address: %s\n", mon.monitorAddr)
	for {
		select {
		case err := <-mon.sub.Err():
			log.Panic(err.Error())
		case header := <-mon.headers:
			block, err := mon.client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Println("Error: ", err)
				continue
			}
			log.Println("Get Block No", block.Number())
			mon.filterNSaveTransactions(block.Transactions())
		case <-mon.quitChan:
			fmt.Println("cancel monitoring")
			return
		}
	}
}

func (mon *service) UnSubNewHeads() {
	mon.StopMonitor()
	mon.sub.Unsubscribe()
	mon.client.Close()
	close(mon.headers)
	close(mon.quitChan)
}

func (mon *service) filterNSaveTransactions(transactions types.Transactions) {
	log.Println("Transaction len: ", len(transactions))
	for _, tx := range transactions {

		//log.Println("Transaction index:", i+1)
		fromAddress, err := getFrom(tx)

		if err != nil {
			log.Println("Get From Error", err)
			continue
		}

		to := fmt.Sprintf("%v", tx.To())
		if to == mon.monitorAddr || fromAddress == mon.monitorAddr {
			log.Println("Found TX: ", tx.Hash())
			txJson, err := json.Marshal(tx)
			if err != nil {
				log.Printf("Error marshaling JSON: %v\n", err)
				continue
			}
			var data map[string]interface{}
			if err := json.Unmarshal(txJson, &data); err != nil {
				log.Printf("Error unmarshaling JSON: %v\n", err)
				continue
			}
			data["monitor_address"] = mon.monitorAddr
			data["from"] = fromAddress
			if err := mon.repo.InsertTx(data); err != nil {
				log.Printf("Cannot save Transaction: %v\n", err)
				continue
			}
			log.Println(`Save transaction into DB success`)

		}
		//log.Println("Transaction index AGAIN:", i+1)

	}
	log.Println("Finish filter")
}

func (mon *service) StopMonitor() {
	mon.quitChan <- 0
}

func getFrom(tx *types.Transaction) (string, error) {
	from, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", from), nil
}
