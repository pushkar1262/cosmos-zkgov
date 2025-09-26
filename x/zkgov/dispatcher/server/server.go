package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/x/zkgov/types"
)

type Dispatcher struct {
	clientCtx *client.Context
}

func NewDispatcher(clientCtx *client.Context, _ interface{}) *Dispatcher {
	return &Dispatcher{clientCtx: clientCtx}
}

func (d *Dispatcher) Run(port string) error {
	http.HandleFunc("/tx", d.handleTx)
	addr := fmt.Sprintf(":%s", port)
	return http.ListenAndServe(addr, nil)
}

func (d *Dispatcher) handleTx(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var msg types.MsgVoteProposal
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Stub: acknowledge receipt; real implementation should sign and broadcast
	w.WriteHeader(http.StatusOK)
}
