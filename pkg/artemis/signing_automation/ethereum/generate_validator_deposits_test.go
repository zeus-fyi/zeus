package signing_automation_ethereum

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/pkg/errors"
	"github.com/wealdtech/go-ed25519hd"
	e2types "github.com/wealdtech/go-eth2-types/v2"
	util "github.com/wealdtech/go-eth2-util"
	keystorev4 "github.com/wealdtech/go-eth2-wallet-encryptor-keystorev4"
	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	"github.com/zeus-fyi/zeus/test/configs"
)

var depositDataPath = filepaths.Path{
	PackageName: "",
	DirIn:       "./mocks/validator_keys",
	DirOut:      "../mocks/validator_keys",
	FnOut:       "",
	Env:         "",
	FilterFiles: strings_filter.FilterOpts{},
}

/*
EIP-2334 defines derivation path indices for withdrawal and validator keys.
For a given index i the keys will be at the following paths:

withdrawal key: m/12381/3600/i/0
validator key: m/12381/3600/i/0/0
*/

// TestEphemeralDepositsFromMnemonicInEth2KeystoreFormat is useful for generating deposits for testing using
// the config file to provide your mnemonic

func (t *Web3SignerClientTestSuite) TestEphemeralDepositsFromMnemonicInEth2KeystoreFormat() {
	configs.ForceDirToConfigLocation()
	seed, serr := ed25519hd.SeedFromMnemonic(t.TestMnemonic, t.TestHDWalletPassword)
	t.Require().Nil(serr)
	ks := keystorev4.New()

	numKeys := 3
	for i := 0; i < numKeys; i++ {
		path := fmt.Sprintf("m/12381/3600/%d/0/0", i)

		sk, err := util.PrivateKeyFromSeedAndPath(seed, path)
		t.Require().Nil(err)

		enc, err := ks.Encrypt(sk.Marshal(), t.TestHDWalletPassword)
		t.Require().Nil(err)
		enc["path"] = path
		enc["pubkey"] = bls_signer.ConvertBytesToString(sk.PublicKey().Marshal())
		enc["version"] = "4"

		b, err := json.Marshal(enc)
		t.Require().Nil(err)

		slashSplit := strings.Split(path, "/")
		underScoreStr := strings.Join(slashSplit, "_")

		depositDataPath.FnOut = fmt.Sprintf("keystore-ephemeral-%s.json", underScoreStr)
		err = depositDataPath.WriteToFileOutPath(b)
		t.Require().Nil(err)
	}
}

func (t *Web3SignerClientTestSuite) TestEphemeralDepositGenerator() {
	s := bls_signer.NewEthBLSAccount()
	wd, err := ValidateAndReturnEcdsaPubkeyBytes(t.TestAccount1.PublicKey())
	t.Require().Nil(err)
	dd, err := GenerateEphemeralDepositData(s, wd)
	t.Require().Nil(err)
	t.Assert().NotEmpty(dd)
	depositDataPath.FnOut = fmt.Sprintf("deposit_data-ephemeral-%d.json", time.Now().Unix())
	dd.PrintJSON(depositDataPath)
}

type DepositInfo struct {
	PublicKey             []byte
	WithdrawalCredentials []byte
	Signature             []byte
	DepositDataRoot       []byte
	DepositMessageRoot    []byte
	ForkVersion           []byte
	Amount                uint64
}

/*
di := DepositInfo{
	PublicKey:             s.PublicKey().Marshal(),
	WithdrawalCredentials: wd,
	Signature:             blsSig.Marshal(),
	DepositDataRoot:       nil,
	DepositMessageRoot:    nil,
	ForkVersion:           nil,
	Amount:                0,
}
*/

// verifyDeposit is mostly borrowed from https://github.com/wealdtech/ethdo
func verifyDeposit(deposit *DepositInfo, withdrawalCredentials []byte, validatorPubKeys map[[48]byte]bool, depositVerifyForkVersion string) (bool, error) {
	if withdrawalCredentials == nil {
		fmt.Println("Withdrawal public key or address not supplied; withdrawal credentials NOT checked")
		return false, nil
	}
	if depositVerifyForkVersion == "" {
		fmt.Println("Fork version not supplied")
		return false, nil
	}
	if !bytes.Equal(deposit.WithdrawalCredentials, withdrawalCredentials) {
		fmt.Println("Withdrawal credentials incorrect")
		return false, nil
	}
	fmt.Println("Withdrawal credentials verified")
	var key [48]byte
	copy(key[:], deposit.PublicKey)
	if _, exists := validatorPubKeys[key]; !exists {
		fmt.Println("Validator public key incorrect")
		return false, nil
	}
	var pubKey phase0.BLSPubKey
	copy(pubKey[:], deposit.PublicKey)
	var signature phase0.BLSSignature
	copy(signature[:], deposit.Signature)
	depositData := &phase0.DepositData{
		PublicKey:             pubKey,
		WithdrawalCredentials: deposit.WithdrawalCredentials,
		Amount:                phase0.Gwei(deposit.Amount),
		Signature:             signature,
	}
	depositDataRoot, err := depositData.HashTreeRoot()
	if err != nil {
		return false, errors.Wrap(err, "failed to generate deposit data root")
	}
	if bytes.Equal(deposit.DepositDataRoot, depositDataRoot[:]) {
		fmt.Println("Deposit data root verified")
	} else {
		fmt.Println("Deposit data root incorrect")
		return false, nil
	}
	if len(deposit.ForkVersion) == 0 {
		if depositVerifyForkVersion != "" {
			fmt.Println("Data format does not contain fork version for verification; NOT verified")
			return false, nil
		}
	}
	forkVersion, err := hex.DecodeString(strings.TrimPrefix(depositVerifyForkVersion, "0x"))
	if err != nil {
		return false, errors.Wrap(err, "failed to decode fork version")
	}
	if bytes.Equal(deposit.ForkVersion, forkVersion) {
		fmt.Println("Fork version verified")
	} else {
		fmt.Println("Fork version incorrect")
		return false, nil
	}
	if len(deposit.DepositMessageRoot) != 32 {
		fmt.Println("Deposit message root not supplied; not checked")
		return false, nil
	}
	// We can also verify the deposit message signature.
	depositMessage := &phase0.DepositMessage{
		PublicKey:             pubKey,
		WithdrawalCredentials: withdrawalCredentials,
		Amount:                phase0.Gwei(deposit.Amount),
	}
	depositMessageRoot, err := depositMessage.HashTreeRoot()
	if err != nil {
		return false, errors.Wrap(err, "failed to generate deposit message root")
	}
	if bytes.Equal(deposit.DepositMessageRoot, depositMessageRoot[:]) {
		fmt.Println("Deposit message root verified")
	} else {
		fmt.Println("Deposit message root incorrect")
		return false, nil
	}
	domainBytes, err := e2types.ComputeDomain(e2types.DomainDeposit, forkVersion, e2types.ZeroGenesisValidatorsRoot)
	if err != nil {
		return false, errors.Wrap(err, "failed to compute domain")
	}
	var domain phase0.Domain
	copy(domain[:], domainBytes)
	container := &phase0.SigningData{
		ObjectRoot: depositMessageRoot,
		Domain:     domain,
	}
	containerRoot, err := container.HashTreeRoot()
	if err != nil {
		return false, errors.New("failed to generate root for container")
	}
	validatorPubKey, err := e2types.BLSPublicKeyFromBytes(pubKey[:])
	if err != nil {
		return false, errors.Wrap(err, "failed to generate validator public key")
	}
	blsSig, err := e2types.BLSSignatureFromBytes(signature[:])
	if err != nil {
		return false, errors.New("failed to verify BLS signature")
	}
	signatureVerified := blsSig.Verify(containerRoot[:], validatorPubKey)
	if signatureVerified {
		fmt.Println("Deposit message signature verified")
	} else {
		fmt.Println("Deposit message signature NOT verified")
		return false, nil
	}
	return true, nil
}
