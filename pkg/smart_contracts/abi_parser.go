package smart_contracts

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type SmartContractABI struct {
	abi.ABI
	Functions map[string]Field
}

// NewSmartContractABI just parses for functions for now
func NewSmartContractABI(jsonData []byte) (SmartContractABI, error) {
	sca := SmartContractABI{
		ABI:       abi.ABI{},
		Functions: make(map[string]Field),
	}

	var fields []Field
	err := json.Unmarshal(jsonData, &fields)
	if err != nil {
		return sca, err
	}
	for _, field := range fields {
		if field.Type == "function" {
			if field.StateMutability == "payable" {
				field.Payable = true
			}
			sca.Functions[field.Name] = field
		}
	}
	return sca, err
}
