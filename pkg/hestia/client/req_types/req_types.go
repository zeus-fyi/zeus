package hestia_req_types

import (
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
)

const (
	EthereumMainnetProtocolNetworkID  = 1
	EthereumEphemeryProtocolNetworkID = 1673748447294772000
)

type CreateValidatorServiceRequest struct {
	ServiceRequestWrapper
	ValidatorServiceOrgGroupSlice
}

type ValidatorServiceOrgGroupSlice []ValidatorServiceOrgGroup

type ValidatorServiceOrgGroup struct {
	Pubkey       string `json:"pubkey"`
	FeeRecipient string `json:"feeRecipient"`
}

type ServiceRequestWrapper struct {
	GroupName         string `json:"groupName"`
	ProtocolNetworkID int    `json:"protocolNetworkID"`
	Enabled           bool   `json:"enabled"`
	ServiceURL        string `json:"serviceURL"`
}

func (vsr *CreateValidatorServiceRequest) CreateValidatorServiceRequest(vsg ValidatorServiceOrgGroupSlice, srw ServiceRequestWrapper) {

	if !strings_filter.ValidateHttpsURL(srw.ServiceURL) {
		panic("you must provide a valid https service link")
	}

	if len(vsr.GroupName) == 0 {
		panic("you must provide a group name")
	}

	if srw.ProtocolNetworkID != EthereumMainnetProtocolNetworkID && srw.ProtocolNetworkID != EthereumEphemeryProtocolNetworkID {
		panic("you must provide a supported protocol network identifier")
	}

	vsr.ServiceRequestWrapper = srw
	vsr.ValidatorServiceOrgGroupSlice = vsg
	for i, _ := range vsg {
		vsr.ValidatorServiceOrgGroupSlice[i].Pubkey = strings_filter.AddHexPrefix(vsr.ValidatorServiceOrgGroupSlice[i].Pubkey)
		if len(vsr.ValidatorServiceOrgGroupSlice[i].Pubkey) != 98 {
			panic("you must provide a valid 0x prefixed bls pubkey value")
		}
		vsr.ValidatorServiceOrgGroupSlice[i].FeeRecipient = strings_filter.AddHexPrefix(vsr.ValidatorServiceOrgGroupSlice[i].FeeRecipient)
	}
}
