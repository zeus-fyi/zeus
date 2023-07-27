package hestia_req_types

import (
	"fmt"

	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
)

const (
	EthereumMainnetProtocolNetworkID  = 1
	EthereumGoerliProtocolNetworkID   = 5
	EthereumEphemeryProtocolNetworkID = 1673748447294772000
	Goerli                            = "goerli"
	Mainnet                           = "mainnet"
	Ephemery                          = "ephemery"
)

func ProtocolNetworkStringToID(network string) int {
	switch network {
	case Mainnet:
		return EthereumMainnetProtocolNetworkID
	case Goerli:
		return EthereumGoerliProtocolNetworkID
	case Ephemery:
		return EthereumEphemeryProtocolNetworkID
	default:
		return 0
	}
}

func ProtocolNetworkIDToString(id int) string {
	switch id {
	case EthereumMainnetProtocolNetworkID:
		return Mainnet
	case EthereumGoerliProtocolNetworkID:
		return Goerli
	case EthereumEphemeryProtocolNetworkID:
		return Ephemery
	default:
		return "unknown"
	}
}

type CreateValidatorServiceRequest struct {
	ServiceRequestWrapper         `json:"serviceRequestWrapper"`
	ValidatorServiceOrgGroupSlice `json:"validatorServiceOrgGroupSlice"`
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
	MevEnabled        bool              `json:"mevEnabled"`
	ServiceAuth       ServiceAuthConfig `json:"serviceAuth"`
}

// ServiceAuthConfig note: only use one auth type per service url
// only AWS lambda support for now
type ServiceAuthConfig struct {
	*AuthLamdbaAWS `json:"awsAuth"`
}

type AuthLamdbaAWS struct {
	ServiceURL string `json:"serviceURL"`

	SecretName string `json:"secretName"` // this is the name of the secret in the aws secrets manager you use for decrypting your keystores
	// these are the auth credentials you link to an IAM user that can call your aws lambda function to sign messages
	// we use these to call your lambda function to sign messages
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"accessSecret"`
}

func (a *ServiceAuthConfig) Validate() error {
	if a.AuthLamdbaAWS == nil {
		err := fmt.Errorf("you must provide an auth config, only aws lambda auth is supported for now")
		return err
	}
	if a.AuthLamdbaAWS != nil {
		if !strings_filter.ValidateHttpsURL(a.AuthLamdbaAWS.ServiceURL) {
			err := fmt.Errorf("you must provide a valid service url")
			return err
		}
		if len(a.AuthLamdbaAWS.SecretName) == 0 {
			err := fmt.Errorf("you must provide a secret name")
			return err
		}
		if len(a.AuthLamdbaAWS.SecretKey) == 0 {
			err := fmt.Errorf("you must provide an access secret")
			return err
		}
		if len(a.AuthLamdbaAWS.AccessKey) == 0 {
			err := fmt.Errorf("you must provide an access key")
			return err
		}
	}
	return nil
}

func (vsr *CreateValidatorServiceRequest) CreateValidatorServiceRequest(vsg ValidatorServiceOrgGroupSlice, srw ServiceRequestWrapper) {
	err := srw.ServiceAuth.Validate()
	if err != nil {
		panic(err)
	}

	if len(srw.GroupName) == 0 {
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

type CreateOrgRoutesRequest struct {
	Routes []string `json:"routes"`
}
