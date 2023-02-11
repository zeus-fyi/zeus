package hestia_req_types

import (
	"fmt"
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
	GroupName         string            `json:"groupName"`
	ProtocolNetworkID int               `json:"protocolNetworkID"`
	Enabled           bool              `json:"enabled"`
	ServiceURL        string            `json:"serviceURL"`
	ServiceAuth       ServiceAuthConfig `json:"serviceAuth"`
}

// ServiceAuthConfig note: only use one auth type per service url
// only AWS lambda support for now
type ServiceAuthConfig struct {
	*AuthLamdbaAWS `json:"awsAuth"`
}

type AuthLamdbaAWS struct {
	SecretName string `json:"secretName"` // this is the name of the secret in the aws secrets manager you use for decrypting your keystores

	// these are the auth credentials you link to an IAM user that can call your aws lambda function to sign messages
	// we use these to call your lambda function to sign messages
	AccessKey    string `json:"accessKey"`
	AccessSecret string `json:"accessSecret"`
}

func (a *ServiceAuthConfig) Validate() {
	if a.AuthLamdbaAWS == nil {
		err := fmt.Errorf("you must provide an auth config, only aws lambda auth is supported for now")
		panic(err)
	}
	if a.AuthLamdbaAWS != nil {
		if len(a.AuthLamdbaAWS.SecretName) == 0 {
			err := fmt.Errorf("you must provide a secret name")
			panic(err)
		}
		if len(a.AuthLamdbaAWS.AccessSecret) == 0 {
			err := fmt.Errorf("you must provide an access secret")
			panic(err)
		}
		if len(a.AuthLamdbaAWS.AccessKey) == 0 {
			err := fmt.Errorf("you must provide an access key")
			panic(err)
		}
	}
}

func (vsr *CreateValidatorServiceRequest) CreateValidatorServiceRequest(vsg ValidatorServiceOrgGroupSlice, srw ServiceRequestWrapper) {
	srw.ServiceAuth.Validate()

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
