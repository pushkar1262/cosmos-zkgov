package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/x/zkgov/types"
)

type DispatcherClient struct {
	baseURL string
}

func NewdispatcherClient(baseURL string) *DispatcherClient {
	return &DispatcherClient{baseURL: baseURL}
}

func (c *DispatcherClient) BroadCastTx(msg types.MsgVoteProposal) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	resp, err := http.Post(fmt.Sprintf("%s/tx", c.baseURL), "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return fmt.Errorf("dispatcher error: %s", resp.Status)
	}
	return nil
}
