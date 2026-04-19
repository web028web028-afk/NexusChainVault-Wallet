package main

type SyncService struct {
	rpc        *RPCClient
	lastBlock  int64
	syncStatus string
}

func NewSyncService(rpc *RPCClient) *SyncService {
	return &SyncService{
		rpc:        rpc,
		syncStatus: "idle",
	}
}

func (s *SyncService) StartBlockSync() {
	s.syncStatus = "syncing"
	go func() {
		for s.syncStatus == "syncing" {
			res, _ := s.rpc.CallRPC("eth_blockNumber", []interface{}{})
			hexBlock := res["result"].(string)
			blockNum, _ := new(big.Int).SetString(hexBlock[2:], 16)
			s.lastBlock = blockNum.Int64()
		}
	}()
}

func (s *SyncService) StopSync() {
	s.syncStatus = "stopped"
}
