package server

import (
	"context"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/smonkeymonkey/eth_total/models"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)


func StartServer() {
	r := mux.NewRouter()
	r.HandleFunc("/api/block/{block_number}/total", totalHandler).Methods(http.MethodGet)
	srv := &http.Server{
		Addr: "127.0.0.1:8080",
		Handler: r,
	}
	done := make(chan os.Signal, 1)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Warn("Server closing connection")
		}
	}()
	log.Info("Server started")

	signal.Notify(done, os.Interrupt)

	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx);err != nil{
		log.Fatal("Server shutdown failed")
	}

	log.Println("Server gracefully stopped")
	os.Exit(0)
}

func totalHandler(rw http.ResponseWriter, r *http.Request) {
	block,err := strconv.Atoi(mux.Vars(r)["block_number"])
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Invalid block number"))
		return
	}
	currentBlock := models.GetBlockByNumber(block)
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(writeAmountAndTrans(currentBlock))
}

func writeAmountAndTrans(currentBlock *models.Block) interface{}{
	resp := struct {
		Transactions int `json:"transactions"`
		Amount *big.Float `json:"amount"`
	}{}

	resp.Transactions = currentBlock.GetTotalTransactions()
	resp.Amount = currentBlock.GetTotalAmount()

	return resp
}
