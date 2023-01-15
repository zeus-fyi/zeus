package hestia_req_types

import signing_automation_ethereum "github.com/zeus-fyi/zeus/pkg/artemis/signing_automation/ethereum"

const (
	EthereumMainnetProtocolNetworkID  = 1
	EthereumEphemeryProtocolNetworkID = 1673748447294772000
)

type CreateValidatorServiceRequest struct {
	ValidatorServiceOrgGroupSlice
}

type ValidatorServiceOrgGroupSlice []ValidatorServiceOrgGroup

type ValidatorServiceOrgGroup struct {
	GroupName         string `json:"groupName"`
	Pubkey            string `json:"pubkey"`
	ProtocolNetworkID int    `json:"protocolNetworkID"`
	FeeRecipient      string `json:"feeRecipient"`
	Enabled           bool   `json:"enabled"`
}

type ServiceRequestWrapper struct {
	GroupName         string `json:"groupName"`
	ProtocolNetworkID int    `json:"protocolNetworkID"`
	FeeRecipient      string `json:"feeRecipient"`
	Enabled           bool   `json:"enabled"`
}

func (vsr *CreateValidatorServiceRequest) CreateValidatorServiceRequest(dp signing_automation_ethereum.ValidatorDepositSlice, srw ServiceRequestWrapper) {
	vsr.ValidatorServiceOrgGroupSlice = make([]ValidatorServiceOrgGroup, len(dp))
	for i, k := range dp {
		vsr.ValidatorServiceOrgGroupSlice[i].GroupName = srw.GroupName
		vsr.ValidatorServiceOrgGroupSlice[i].FeeRecipient = srw.FeeRecipient
		vsr.ValidatorServiceOrgGroupSlice[i].Enabled = srw.Enabled
		vsr.ValidatorServiceOrgGroupSlice[i].ProtocolNetworkID = srw.ProtocolNetworkID
		vsr.ValidatorServiceOrgGroupSlice[i].Pubkey = k.Pubkey
	}
}
