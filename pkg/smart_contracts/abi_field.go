package smart_contracts

import "github.com/ethereum/go-ethereum/accounts/abi"

type Field struct {
	Type    string
	Name    string
	Inputs  []abi.Argument
	Outputs []abi.Argument

	// Status indicator which can be: "pure", "view",
	// "nonpayable" or "payable".
	StateMutability string

	// Deprecated Status indicators, but removed in v0.6.0.
	Constant bool // True if function is either pure or view
	Payable  bool // True if function is payable

	// Event relevant indicator represents the event is
	// declared as anonymous.
	Anonymous bool
}
