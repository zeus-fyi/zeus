package hestia_req_types

import (
	signing_automation_ethereum "github.com/zeus-fyi/zeus/pkg/artemis/signing_automation/ethereum"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
)

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
	ServiceURL        string `json:"serviceURL"`
}

type ServiceRequestWrapper struct {
	GroupName         string `json:"groupName"`
	ProtocolNetworkID int    `json:"protocolNetworkID"`
	FeeRecipient      string `json:"feeRecipient"`
	Enabled           bool   `json:"enabled"`
	ServiceURL        string `json:"serviceURL"`
}

func (vsr *CreateValidatorServiceRequest) CreateValidatorServiceRequest(dp signing_automation_ethereum.ValidatorDepositSlice, srw ServiceRequestWrapper) {
	vsr.ValidatorServiceOrgGroupSlice = make([]ValidatorServiceOrgGroup, len(dp))
	for i, k := range dp {
		vsr.ValidatorServiceOrgGroupSlice[i].GroupName = srw.GroupName
		vsr.ValidatorServiceOrgGroupSlice[i].FeeRecipient = srw.FeeRecipient
		vsr.ValidatorServiceOrgGroupSlice[i].Enabled = srw.Enabled
		vsr.ValidatorServiceOrgGroupSlice[i].ProtocolNetworkID = srw.ProtocolNetworkID
		vsr.ValidatorServiceOrgGroupSlice[i].Pubkey = strings_filter.AddHexPrefix(k.Pubkey)
		vsr.ValidatorServiceOrgGroupSlice[i].ProtocolNetworkID = srw.ProtocolNetworkID
		if strings_filter.ValidateHttpsURL(srw.ServiceURL) {
			vsr.ValidatorServiceOrgGroupSlice[i].ServiceURL = srw.ServiceURL
		} else {
			panic("you must provide a valid https service link")
		}
	}
}
