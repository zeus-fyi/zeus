package serverless_aws_automation

import (
	"context"
	"fmt"
	"github.com/zeus-fyi/zeus/builds"
	signing_automation_ethereum "github.com/zeus-fyi/zeus/pkg/artemis/signing_automation/ethereum"
	hestia_client "github.com/zeus-fyi/zeus/pkg/hestia/client"
	hestia_req_types "github.com/zeus-fyi/zeus/pkg/hestia/client/req_types"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
)

func CreateHestiaValidatorsServiceRequest(ctx context.Context, keystoresPath filepaths.Path, sr hestia_req_types.ServiceRequestWrapper, bearerToken, feeRecipientAddr string) {
	fmt.Println("INFO: Creating Hestia Validators Service Request...")
	hc := hestia_client.NewDefaultHestiaClient(bearerToken)
	builds.ChangeToBuildsDir()
	filter := &strings_filter.FilterOpts{StartsWith: "deposit_data"}
	keystoresPath.FilterFiles = filter
	dpSlice, err := signing_automation_ethereum.ParseValidatorDepositSliceJSON(ctx, keystoresPath)
	if err != nil {
		panic(err)
	}
	pubkeys := hestia_req_types.ValidatorServiceOrgGroupSlice{}
	for _, validatorDepositInfo := range dpSlice {
		pubkeys = append(pubkeys, hestia_req_types.ValidatorServiceOrgGroup{
			Pubkey:       validatorDepositInfo.Pubkey,
			FeeRecipient: feeRecipientAddr,
		})
	}
	hs := hestia_req_types.CreateValidatorServiceRequest{}
	hs.CreateValidatorServiceRequest(pubkeys, sr)
	resp, err := hc.ValidatorsServiceRequest(ctx, hs)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
