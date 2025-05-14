package db

import "github.com/scalarorg/data-models/chains"

type QueryOptions struct {
	IncludeCallContract          *bool
	IncludeCallContractWithToken *bool
}

type ContractCallExecuteResult struct {
	EventId string
	Status  chains.ContractCallStatus
}
