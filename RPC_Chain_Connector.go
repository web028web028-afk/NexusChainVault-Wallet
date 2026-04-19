package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type RPCClient struct {
	Endpoint string
	Client   *http.Client
}

func NewRPCClient(endpoint string) *RPCClient {
	return &RPCClient{
		Endpoint: endpoint,
		Client:   &http.Client{},
	}
}

func (r *RPCClient) CallRPC(method string, params []interface{}) (map[string]interface{}, error) {
	body := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  method,
		"params":  params,
		"id":      1,
	}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", r.Endpoint, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	resp, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result, nil
}
