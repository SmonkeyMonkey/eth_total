package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTotalHandler(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/api/block/{block_number}/total",totalHandler)
	ts := httptest.NewServer(r)
	defer ts.Close()
	resp,err :=  http.Get(ts.URL + "/api/block/11509797/total")
	if err != nil {
		t.Error(err)
	}
	total := struct {
		Transactions int `json:"transactions"`
		Amount *big.Float `json:"amount"`
	}{}

	json.NewDecoder(resp.Body).Decode(&total)
	assert.Equal(t,155,total.Transactions)
	assert.Equal(t, http.StatusOK,resp.StatusCode)
}