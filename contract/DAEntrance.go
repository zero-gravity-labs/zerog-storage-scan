// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IDAEntranceCommitRootSubmission is an auto generated low-level Go binding around an user-defined struct.
type IDAEntranceCommitRootSubmission struct {
	DataRoot          [32]byte
	Epoch             *big.Int
	QuorumId          *big.Int
	ErasureCommitment BN254G1Point
	QuorumBitmap      []byte
	AggPkG2           BN254G2Point
	Signature         BN254G1Point
}

// IDASampleSampleRange is an auto generated low-level Go binding around an user-defined struct.
type IDASampleSampleRange struct {
	StartEpoch uint64
	EndEpoch   uint64
}

// IDASampleSampleTask is an auto generated low-level Go binding around an user-defined struct.
type IDASampleSampleTask struct {
	SampleHash      [32]byte
	PodasTarget     *big.Int
	RestSubmissions uint64
}

// SampleResponse is an auto generated low-level Go binding around an user-defined struct.
type SampleResponse struct {
	SampleSeed   [32]byte
	Epoch        uint64
	QuorumId     uint64
	LineIndex    uint32
	SublineIndex uint32
	Quality      *big.Int
	DataRoot     [32]byte
	BlobRoots    [3][32]byte
	Proof        [][32]byte
	Data         []byte
}

// DAEntranceMetaData contains all meta data concerning the DAEntrance contract.
var DAEntranceMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sampleRound\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quorumId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quality\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lineIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sublineIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"DAReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quorumId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blobPrice\",\"type\":\"uint256\"}],\"name\":\"DataUpload\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quorumId\",\"type\":\"uint256\"}],\"name\":\"ErasureCommitmentVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sampleRound\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sampleHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sampleSeed\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"podasTarget\",\"type\":\"uint256\"}],\"name\":\"NewSampleRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DA_SIGNERS\",\"outputs\":[{\"internalType\":\"contractIDASigners\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_PODAS_TARGET\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PARAMS_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLICE_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLICE_NUMERATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activedReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blobPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_dataRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_quorumId\",\"type\":\"uint256\"}],\"name\":\"commitmentExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpochReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentSampleSeed\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"donate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochWindowSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextSampleHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"}],\"name\":\"payments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"podasTarget\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundSubmissions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"samplePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sampleRange\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"startEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endEpoch\",\"type\":\"uint64\"}],\"internalType\":\"structIDASample.SampleRange\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sampleRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sampleTask\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sampleHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"podasTarget\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"restSubmissions\",\"type\":\"uint64\"}],\"internalType\":\"structIDASample.SampleTask\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serviceFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serviceFeeRateBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_baseReward\",\"type\":\"uint256\"}],\"name\":\"setBaseReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blobPrice\",\"type\":\"uint256\"}],\"name\":\"setBlobPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_epochWindowSize\",\"type\":\"uint64\"}],\"name\":\"setEpochWindowSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_rewardRatio\",\"type\":\"uint64\"}],\"name\":\"setRewardRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_targetRoundSubmissions\",\"type\":\"uint64\"}],\"name\":\"setRoundSubmissions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"samplePeriod_\",\"type\":\"uint64\"}],\"name\":\"setSamplePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bps\",\"type\":\"uint256\"}],\"name\":\"setServiceFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"treasury_\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_dataRoots\",\"type\":\"bytes32[]\"}],\"name\":\"submitOriginalData\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sampleSeed\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"quorumId\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"lineIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"sublineIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"quality\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[3]\",\"name\":\"blobRoots\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structSampleResponse\",\"name\":\"rep\",\"type\":\"tuple\"}],\"name\":\"submitSamplingResponse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quorumId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"erasureCommitment\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"quorumBitmap\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"aggPkG2\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"signature\",\"type\":\"tuple\"}],\"internalType\":\"structIDAEntrance.CommitRootSubmission[]\",\"name\":\"_submissions\",\"type\":\"tuple[]\"}],\"name\":\"submitVerifiedCommitRoots\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sync\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetRoundSubmissions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetRoundSubmissionsNext\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBaseReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_dataRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_quorumId\",\"type\":\"uint256\"}],\"name\":\"verifiedErasureCommitment\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"payee\",\"type\":\"address\"}],\"name\":\"withdrawPayments\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}][{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sampleRound\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quorumId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quality\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lineIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sublineIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"DAReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quorumId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blobPrice\",\"type\":\"uint256\"}],\"name\":\"DataUpload\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quorumId\",\"type\":\"uint256\"}],\"name\":\"ErasureCommitmentVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"sampleRound\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sampleHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sampleSeed\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"podasTarget\",\"type\":\"uint256\"}],\"name\":\"NewSampleRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DA_SIGNERS\",\"outputs\":[{\"internalType\":\"contractIDASigners\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_PODAS_TARGET\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PARAMS_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLICE_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLICE_NUMERATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activedReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blobPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_dataRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_quorumId\",\"type\":\"uint256\"}],\"name\":\"commitmentExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpochReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentSampleSeed\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"donate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochWindowSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextSampleHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"}],\"name\":\"payments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"podasTarget\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundSubmissions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"samplePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sampleRange\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"startEpoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endEpoch\",\"type\":\"uint64\"}],\"internalType\":\"structIDASample.SampleRange\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sampleRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sampleTask\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sampleHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"podasTarget\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"restSubmissions\",\"type\":\"uint64\"}],\"internalType\":\"structIDASample.SampleTask\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serviceFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serviceFeeRateBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_baseReward\",\"type\":\"uint256\"}],\"name\":\"setBaseReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blobPrice\",\"type\":\"uint256\"}],\"name\":\"setBlobPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_epochWindowSize\",\"type\":\"uint64\"}],\"name\":\"setEpochWindowSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_rewardRatio\",\"type\":\"uint64\"}],\"name\":\"setRewardRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_targetRoundSubmissions\",\"type\":\"uint64\"}],\"name\":\"setRoundSubmissions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"samplePeriod_\",\"type\":\"uint64\"}],\"name\":\"setSamplePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bps\",\"type\":\"uint256\"}],\"name\":\"setServiceFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"treasury_\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_dataRoots\",\"type\":\"bytes32[]\"}],\"name\":\"submitOriginalData\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sampleSeed\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"quorumId\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"lineIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"sublineIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"quality\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[3]\",\"name\":\"blobRoots\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structSampleResponse\",\"name\":\"rep\",\"type\":\"tuple\"}],\"name\":\"submitSamplingResponse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"dataRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quorumId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"erasureCommitment\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"quorumBitmap\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"aggPkG2\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"signature\",\"type\":\"tuple\"}],\"internalType\":\"structIDAEntrance.CommitRootSubmission[]\",\"name\":\"_submissions\",\"type\":\"tuple[]\"}],\"name\":\"submitVerifiedCommitRoots\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sync\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetRoundSubmissions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetRoundSubmissionsNext\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBaseReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_dataRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_quorumId\",\"type\":\"uint256\"}],\"name\":\"verifiedErasureCommitment\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"payee\",\"type\":\"address\"}],\"name\":\"withdrawPayments\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50614b55806100206000396000f3fe608060405260043610620003795760003560e01c80638129fc1c11620001ce578063b1be17ab116200010b578063e68e035b11620000a1578063f0f442601162000078578063f0f442601462000a01578063f69027751462000a26578063ff1877481462000a4b578063fff6cae91462000a6357600080fd5b8063e68e035b14620009ba578063eafed6ce14620009d2578063ed88c68e14620009f757600080fd5b8063ca15c87311620000e2578063ca15c8731462000934578063d4ae59c91462000959578063d547741f1462000970578063e2982c21146200099557600080fd5b8063b1be17ab14620008df578063c05751111462000904578063c8d3b359146200091c57600080fd5b8063988ea94e11620001815780639da3a69b11620001585780639da3a69b146200084a5780639fae1fa4146200088b578063a217fddf14620008a3578063b15d20da14620008ba57600080fd5b8063988ea94e14620007c957806398920f57146200080d5780639b1d3091146200082557600080fd5b80638129fc1c146200071257806388521ec7146200072a5780638abdf5aa146200074f5780638bdcc71214620007675780639010d07c146200077f57806391d1485414620007a457600080fd5b80633bab2a7011620002ba5780636a53525d11620002505780637667180811620002275780637667180814620006b257806376ad03bc14620006ca5780637f1b5e4314620006e2578063807f063a14620006fa57600080fd5b80636a53525d14620006355780636efc2555146200065a5780637397eb33146200069a57600080fd5b80635626a47b11620002915780635626a47b14620005b4578063602b245a14620005cb57806361d027b314620005e2578063646033bc146200061d57600080fd5b80633bab2a70146200055f5780633d00448a14620005845780633e898337146200059c57600080fd5b806323dd60a611620003305780632f2ff15d11620003075780632f2ff15d14620004d85780632fc0534b14620004fd57806331b3eb94146200051557806336568abe146200053a57600080fd5b806323dd60a61462000466578063248a9ca3146200048b578063257a3aa314620004c057600080fd5b806301ffc9a7146200037e5780630373a23a14620003b85780631192de9a14620003df578063125770521462000404578063158ef93e146200042b578063168a062c146200044e575b600080fd5b3480156200038b57600080fd5b50620003a36200039d36600462003a58565b62000a7b565b60405190151581526020015b60405180910390f35b348015620003c557600080fd5b50620003dd620003d736600462003a84565b62000aa9565b005b348015620003ec57600080fd5b50620003dd620003fe36600462003abb565b62000ad4565b3480156200041157600080fd5b506200041c60405481565b604051908152602001620003af565b3480156200043857600080fd5b50600054620003a390600160a01b900460ff1681565b3480156200045b57600080fd5b506200041c603a5481565b3480156200047357600080fd5b50620003dd6200048536600462003a84565b62000b41565b3480156200049857600080fd5b506200041c620004aa36600462003a84565b6000908152600160208190526040909120015490565b348015620004cd57600080fd5b506200041c60435481565b348015620004e557600080fd5b50620003dd620004f736600462003aef565b62000b6c565b3480156200050a57600080fd5b506200041c603e5481565b3480156200052257600080fd5b50620003dd6200053436600462003b22565b62000b9b565b3480156200054757600080fd5b50620003dd6200055936600462003aef565b62000bff565b3480156200056c57600080fd5b50620003dd6200057e36600462003abb565b62000c81565b3480156200059157600080fd5b506200041c60475481565b348015620005a957600080fd5b506200041c60445481565b348015620005c157600080fd5b506200041c600281565b348015620005d857600080fd5b506200041c600381565b348015620005ef57600080fd5b50604a5462000604906001600160a01b031681565b6040516001600160a01b039091168152602001620003af565b3480156200062a57600080fd5b506200041c60455481565b3480156200064257600080fd5b50620003a36200065436600462003b42565b62000d10565b3480156200066757600080fd5b506200067262000d37565b6040805182516001600160401b039081168252602093840151169281019290925201620003af565b348015620006a757600080fd5b506200041c603b5481565b348015620006bf57600080fd5b506200041c60385481565b348015620006d757600080fd5b506200041c60465481565b348015620006ef57600080fd5b506200041c60415481565b3480156200070757600080fd5b506200060461100081565b3480156200071f57600080fd5b50620003dd62000dc5565b3480156200073757600080fd5b50620003dd6200074936600462003abb565b62000f83565b3480156200075c57600080fd5b506200041c60425481565b3480156200077457600080fd5b506200041c620010ff565b3480156200078c57600080fd5b50620006046200079e36600462003b6f565b62001111565b348015620007b157600080fd5b50620003a3620007c336600462003aef565b62001132565b348015620007d657600080fd5b50620007e16200115d565b604080518251815260208084015190820152918101516001600160401b031690820152606001620003af565b3480156200081a57600080fd5b506200041c60485481565b3480156200083257600080fd5b50620003dd6200084436600462003a84565b620011d5565b3480156200085757600080fd5b506200086f6200086936600462003b42565b62001200565b60408051825181526020928301519281019290925201620003af565b3480156200089857600080fd5b506200041c60395481565b348015620008b057600080fd5b506200041c600081565b348015620008c757600080fd5b506200041c60008051602062004b0083398151915281565b348015620008ec57600080fd5b50620003dd620008fe36600462003abb565b62001270565b3480156200091157600080fd5b506200041c60495481565b3480156200092957600080fd5b506200041c603c5481565b3480156200094157600080fd5b506200041c6200095336600462003a84565b620012ff565b620003dd6200096a36600462003cea565b62001318565b3480156200097d57600080fd5b50620003dd6200098f36600462003aef565b6200151a565b348015620009a257600080fd5b506200041c620009b436600462003b22565b62001544565b348015620009c757600080fd5b506200041c603f5481565b348015620009df57600080fd5b50620003dd620009f136600462003e90565b620015b6565b620003dd620018f4565b34801562000a0e57600080fd5b50620003dd62000a2036600462003b22565b62001919565b34801562000a3357600080fd5b50620003dd62000a4536600462004044565b62001961565b34801562000a5857600080fd5b506200041c603d5481565b34801562000a7057600080fd5b50620003dd62001eb4565b60006001600160e01b03198216635a05180f60e01b148062000aa3575062000aa38262001eca565b92915050565b60008051602062004b0083398151915262000ac48162001f01565b62000ace62001eb4565b50604655565b60008051602062004b0083398151915262000aef8162001f01565b62000af962001eb4565b6001600160401b038216604855603a5460000362000b3d5760485462000b2081436200418e565b62000b2d906001620041a5565b62000b399190620041bb565b603b555b5050565b60008051602062004b0083398151915262000b5c8162001f01565b62000b6662001eb4565b50604755565b6000828152600160208190526040909120015462000b8a8162001f01565b62000b96838362001f10565b505050565b6000546040516351cff8d960e01b81526001600160a01b038381166004830152909116906351cff8d990602401600060405180830381600087803b15801562000be357600080fd5b505af115801562000bf8573d6000803e3d6000fd5b5050505050565b6001600160a01b038116331462000c755760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b60648201526084015b60405180910390fd5b62000b3d828262001f36565b60008051602062004b0083398151915262000c9c8162001f01565b62000ca662001eb4565b6000826001600160401b03161162000d015760405162461bcd60e51b815260206004820152601d60248201527f52657761726420726174696f206d757374206265206e6f6e2d7a65726f000000604482015260640162000c6c565b506001600160401b0316604555565b60008062000d2085858562001200565b905062000d2d8162001f5c565b1595945050505050565b604080518082019091526000808252602082015262000d5562001eb4565b6000806000603854111562000d7857600160385462000d759190620041dd565b90505b604454811062000da257600160445462000d939190620041dd565b62000d9f9082620041dd565b91505b604080518082019091526001600160401b03928316815291166020820152919050565b600054600160a01b900460ff161562000e2d5760405162461bcd60e51b8152602060048201526024808201527f5a67496e697469616c697a61626c653a20616c726561647920696e697469616c6044820152631a5e995960e21b606482015260840162000c6c565b6000805460ff60a01b1916600160a01b17815562000e539062000e4d3390565b62001f10565b62000e6e60008051602062004b008339815191523362001f10565b6110006001600160a01b031663f4145a836040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000eaf573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000ed59190620041f3565b603855601e604881905562000eeb81436200418e565b62000ef8906001620041a5565b62000f049190620041bb565b603b5562000f1660806000196200418e565b603c55601460435561012c60445562124f806045556000604681905560475560405162000f43906200396c565b604051809103906000f08015801562000f60573d6000803e3d6000fd5b50600080546001600160a01b0319166001600160a01b0392909216919091179055565b60008051602062004b0083398151915262000f9e8162001f01565b62000fa862001eb4565b603e5462000fb8906004620041bb565b826001600160401b031611156200101e5760405162461bcd60e51b8152602060048201526024808201527f496e63726561736520726f756e64207375626d697373696f6e7320746f6f206c6044820152636172676560e01b606482015260840162000c6c565b6004603e546200102f91906200418e565b826001600160401b03161015620010955760405162461bcd60e51b8152602060048201526024808201527f446563726561736520726f756e64207375626d697373696f6e7320746f6f206c6044820152636172676560e01b606482015260840162000c6c565b6000826001600160401b031611620010f05760405162461bcd60e51b815260206004820181905260248201527f526f756e64207375626d697373696f6e732063616e6e6f74206265207a65726f604482015260640162000c6c565b506001600160401b0316604355565b6200110e60806000196200418e565b81565b60008281526002602052604081206200112b908362001f73565b9392505050565b60009182526001602090815260408084206001600160a01b0393909316845291905290205460ff1690565b60408051606081018252600080825260208201819052918101919091526200118462001eb4565b6000603e546002620011979190620041bb565b905060405180606001604052806039548152602001603c548152602001603d5483620011c49190620041dd565b6001600160401b0316905292915050565b60008051602062004b00833981519152620011f08162001f01565b620011fa62001eb4565b50604955565b604080518082018252600080825260209182018190528251808301969096528583019490945260608086019390935281518086039093018352608085018083528351938201939093208452603590529182902060c084019092528154815260019091015460a09092019190915290565b60008051602062004b008339815191526200128b8162001f01565b6200129562001eb4565b6000826001600160401b031611620012f05760405162461bcd60e51b815260206004820181905260248201527f45706f63682077696e646f772073697a652063616e6e6f74206265207a65726f604482015260640162000c6c565b506001600160401b0316604455565b600081815260026020526040812062000aa39062001f81565b6200132262001eb4565b6047548151620013339190620041bb565b3410156200137d5760405162461bcd60e51b81526020600482015260166024820152754e6f7420656e6f75676820646120626c6f622066656560501b604482015260640162000c6c565b34603f6000828254620013919190620041a5565b9091555050603854604051635ecba50360e01b815260009161100091635ecba50391620013c49160040190815260200190565b602060405180830381865afa158015620013e2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620014089190620041f3565b9050600081116200145c5760405162461bcd60e51b815260206004820152601960248201527f4441456e7472616e63653a204e6f204441205369676e65727300000000000000604482015260640162000c6c565b8060365460016200146e9190620041a5565b6200147a91906200420d565b603655815160005b8181101562001514577fb4e0ecfec4293e970525d9286428425fbdc041540ca6e58ad11bce23d16ed41c848281518110620014c157620014c162004224565b6020026020010151603854603654604754604051620014f9949392919093845260208401929092526040830152606082015260800190565b60405180910390a16200150c816200423a565b905062001482565b50505050565b60008281526001602081905260409091200154620015388162001f01565b62000b96838362001f36565b600080546040516371d4ed8d60e11b81526001600160a01b0384811660048301529091169063e3a9db1a90602401602060405180830381865afa15801562001590573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000aa39190620041f3565b805160005b8181101562000b965762001631838281518110620015dd57620015dd62004224565b602002602001015160000151848381518110620015fe57620015fe62004224565b6020026020010151602001518584815181106200161f576200161f62004224565b60200260200101516040015162000d10565b620018e15760008060006110006001600160a01b03166350b7373987868151811062001661576200166162004224565b60200260200101516020015188878151811062001682576200168262004224565b602002602001015160400151898881518110620016a357620016a362004224565b6020026020010151608001516040518463ffffffff1660e01b8152600401620016cf93929190620042aa565b608060405180830381865afa158015620016ed573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620017139190620042d4565b9250925092506200174a8387868151811062001733576200173362004224565b602002602001015162001f8c90919063ffffffff16565b62001757600382620041bb565b62001764836002620041bb565b1115620017c35760405162461bcd60e51b815260206004820152602660248201527f444152656769737472793a20696e73756666696369656e74207369676e656420604482015265736c6963657360d01b606482015260840162000c6c565b858481518110620017d857620017d862004224565b602002602001015160600151603560006200180f89888151811062001801576200180162004224565b602002602001015162002160565b81526020808201929092526040016000208251815591015160019091015585517f0f1b20d87bebd11dddaaab51f01cf2726880cb3f8073b636dbafa2aa8cacd2569087908690811062001866576200186662004224565b60200260200101516000015187868151811062001887576200188762004224565b602002602001015160200151888781518110620018a857620018a862004224565b602002602001015160400151604051620018d5939291909283526020830191909152604082015260600190565b60405180910390a15050505b620018ec816200423a565b9050620015bb565b620018fe62001eb4565b3460416000828254620019129190620041a5565b9091555050565b60008051602062004b00833981519152620019348162001f01565b6200193e62001eb4565b50604a80546001600160a01b0319166001600160a01b0392909216919091179055565b6200196b62001eb4565b8051602080830151604080850151606086015160808701518351808701979097526001600160c01b031960c095861b8116888601529290941b90911660488601526001600160e01b031960e091821b8116605087015292901b90911660548401528051808403603801815260589093019052815191012060009060008181526037602052604090205490915060ff161562001a415760405162461bcd60e51b8152602060048201526015602482015274223ab83634b1b0ba32b21039bab136b4b9b9b4b7b760591b604482015260640162000c6c565b6000818152603760205260409020805460ff19166001179055603a5462001aab5760405162461bcd60e51b815260206004820152601e60248201527f53616d706c6520726f756e6420302063616e6e6f74206265206d696e65640000604482015260640162000c6c565b603e5462001abb906002620041bb565b603d541062001b175760405162461bcd60e51b815260206004820152602160248201527f546f6f206d616e79207375626d697373696f6e7320696e206f6e6520726f756e6044820152601960fa1b606482015260840162000c6c565b60395482511462001b635760405162461bcd60e51b8152602060048201526015602482015274155b9b585d18da1959081cd85b5c1b19481cd95959605a1b604482015260640162000c6c565b603c548260a00151111562001bb15760405162461bcd60e51b8152602060048201526013602482015272145d585b1a5d1e481b9bdd081c995858da1959606a1b604482015260640162000c6c565b62001bdc8260c0015183602001516001600160401b031684604001516001600160401b031662000d10565b62001c225760405162461bcd60e51b8152602060048201526015602482015274155b9c9958dbdc9919590818dbdb5b5a5d1b595b9d605a1b604482015260640162000c6c565b60385460445483602001516001600160401b031662001c429190620041a5565b101562001c925760405162461bcd60e51b815260206004820152601a60248201527f45706f6368206861732073746f707065642073616d706c696e67000000000000604482015260640162000c6c565b60385482602001516001600160401b03161062001cf25760405162461bcd60e51b815260206004820152601b60248201527f43616e6e6f742073616d706c652063757272656e742065706f63680000000000604482015260640162000c6c565b62001cfd82620021a7565b602082015160408084015160608501519151637d37e5d360e11b81526001600160401b0393841660048201529216602483015263ffffffff1660448201526000906110009063fa6fcba690606401602060405180830381865afa15801562001d69573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062001d8f91906200432a565b90506001603d600082825462001da69190620041a5565b909155505060455460405460009162001dbf916200418e565b9050806040600082825462001dd59190620041dd565b9091555062001de5905062002556565b62001df19082620041a5565b9050801562001e065762001e06828262002591565b83602001516001600160401b0316603a54836001600160a01b03167fc3898eb7106c1cb2f727da316a76320c0035f5692950aa7f6b65d20a5efaedc587604001518860c001518960a001518a606001518b608001518960405162001ea6969594939291906001600160401b039690961686526020860194909452604085019290925263ffffffff908116606085015216608083015260a082015260c00190565b60405180910390a450505050565b62001ebe620025ef565b62001ec862002676565b565b60006001600160e01b03198216637965db0b60e01b148062000aa357506301ffc9a760e01b6001600160e01b031983161462000aa3565b62001f0d81336200273a565b50565b62001f1c82826200279e565b600082815260026020526040902062000b9690826200280c565b62001f42828262002823565b600082815260026020526040902062000b9690826200288d565b805160009015801562000aa3575050602001511590565b60006200112b8383620028a4565b600062000aa3825490565b600062001f9983620028d1565b60a084015160c085015180516020808301518751888301518651848801518951868b01516040519b9c50999a98996000997f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001996200200199909897969594939291016200436f565b6040516020818303038152906040528051906020012060001c6200202691906200420d565b9050600080620020a1620020476200203f89866200294e565b8690620029e4565b6200205162002a7e565b620020966200208e8762002087604080518082018252600080825260209182015281518083019092526001825260029082015290565b906200294e565b8a90620029e4565b886201d4c062002b40565b9150915081620021075760405162461bcd60e51b815260206004820152602a60248201527f444152656769737472793a2070616972696e6720707265636f6d70696c652063604482015269185b1b0819985a5b195960b21b606482015260840162000c6c565b80620021565760405162461bcd60e51b815260206004820181905260248201527f444152656769737472793a207369676e617475726520697320696e76616c6964604482015260640162000c6c565b5050505050505050565b805160208083015160408085015181518085019590955284820192909252606080850192909252805180850390920182526080909301909252815191012060009062000aa3565b8051620021f75760405162461bcd60e51b815260206004820152601b60248201527f53616d706c6520736565642063616e6e6f7420626520656d7074790000000000604482015260640162000c6c565b620022066003610400620043bf565b6001600160401b0316816060015163ffffffff1610620022605760405162461bcd60e51b8152602060048201526014602482015273092dcc6dee4e4cac6e840d8d2dcca40d2dcc8caf60631b604482015260640162000c6c565b60206001600160401b0316816080015163ffffffff1610620022c55760405162461bcd60e51b815260206004820152601860248201527f496e636f7272656374207375622d6c696e6520696e6465780000000000000000604482015260640162000c6c565b60006200234b826000015183602001516001600160401b031684604001516001600160401b03168560c00151866060015163ffffffff1660408051602080820197909752808201959095526060850193909352608084019190915260c01b6001600160c01b03191660a08301528051608881840301815260a89092019052805191012090565b905060006200236b82846080015163ffffffff1685610120015162002da4565b9050806200237c83600019620041dd565b1015620023bf5760405162461bcd60e51b815260206004820152601060248201526f5175616c697479206f766572666c6f7760801b604482015260640162000c6c565b60a0830151620023d08284620041a5565b14620024135760405162461bcd60e51b8152602060048201526011602482015270496e636f7272656374207175616c69747960781b604482015260640162000c6c565b60206200242361040082620043bf565b6200242f9190620043f1565b6001600160401b03168361012001515114620024865760405162461bcd60e51b8152602060048201526015602482015274092dcc6dee4e4cac6e840c8c2e8c240d8cadccee8d605b1b604482015260640162000c6c565b60006200249884610120015162002ddd565b90506000610400856060015163ffffffff16620024b69190620043f1565b905060008560e00151826001600160401b031660038110620024dc57620024dc62004224565b602002015190506000866080015163ffffffff166020610400896060015163ffffffff166200250c91906200441a565b620025189190620043bf565b62002524919062004443565b9050620025398483896101000151846200303c565b6200254d8760e001518860c0015162003155565b50505050505050565b6000604654604154116200256d5760415462002571565b6046545b90508060416000828254620025879190620041dd565b9250508190555090565b60005460405163f340fa0160e01b81526001600160a01b0384811660048301529091169063f340fa019083906024016000604051808303818588803b158015620025da57600080fd5b505af11580156200254d573d6000803e3d6000fd5b60006110006001600160a01b031663f4145a836040518163ffffffff1660e01b8152600401602060405180830381865afa15801562002632573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620026589190620041f3565b90508060385403620026675750565b603881905562001f0d620031a5565b603b544310156200268357565b603a54156200269a576200269662003216565b603c555b6001603a6000828254620026af9190620041a5565b9091555050603b54620026c4600182620041dd565b40603955604854603b8054600090620026df908490620041a5565b9091555050604354603e556000603d55603a54603954603c546040805185815260208101939093528201527fdfb5db5886e81f083727f21152a2a83457e99364e9f104e1aa10bbd6d9b4b95f9060600160405180910390a250565b62002746828262001132565b62000b3d57620027568162003307565b620027638360206200331a565b6040516020016200277692919062004466565b60408051601f198184030181529082905262461bcd60e51b825262000c6c91600401620044df565b620027aa828262001132565b62000b3d5760008281526001602081815260408084206001600160a01b0386168086529252808420805460ff19169093179092559051339285917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45050565b60006200112b836001600160a01b038416620034d3565b6200282f828262001132565b1562000b3d5760008281526001602090815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b60006200112b836001600160a01b03841662003525565b6000826000018281548110620028be57620028be62004224565b9060005260206000200154905092915050565b604080518082019091526000808252602082015281516020808401516040808601516060870151805190850151925162000aa39662002932969095949101948552602085019390935260408401919091526060830152608082015260a00190565b6040516020818303038152906040528051906020012062003630565b60408051808201909152600080825260208201526200296c6200397a565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa905080806200299c57fe5b5080620029dc5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b604482015260640162000c6c565b505092915050565b604080518082019091526000808252602082015262002a0262003998565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa9050808062002a3e57fe5b5080620029dc5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b604482015260640162000c6c565b62002a88620039b6565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b60408051808201825286815260208082018690528251808401909352868352820184905260009182919062002b74620039df565b60005b600281101562002d7557600062002b90826006620041bb565b905084826002811062002ba75762002ba762004224565b6020020151518362002bbb836000620041a5565b600c811062002bce5762002bce62004224565b602002015284826002811062002be85762002be862004224565b6020020151602001518382600162002c019190620041a5565b600c811062002c145762002c1462004224565b602002015283826002811062002c2e5762002c2e62004224565b602002015151518362002c43836002620041a5565b600c811062002c565762002c5662004224565b602002015283826002811062002c705762002c7062004224565b602002015151600160200201518362002c8b836003620041a5565b600c811062002c9e5762002c9e62004224565b602002015283826002811062002cb85762002cb862004224565b60200201516020015160006002811062002cd65762002cd662004224565b60200201518362002ce9836004620041a5565b600c811062002cfc5762002cfc62004224565b602002015283826002811062002d165762002d1662004224565b60200201516020015160016002811062002d345762002d3462004224565b60200201518362002d47836005620041a5565b600c811062002d5a5762002d5a62004224565b6020020152508062002d6c816200423a565b91505062002b77565b5062002d80620039fe565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b600083838360405160200162002dbd93929190620044f4565b60408051601f198184030181529190528051602090910120949350505050565b8051600090610100811080159062002e01575062002dfd600182620041dd565b8116155b62002e6e5760405162461bcd60e51b815260206004820152603660248201527f44617461206c656e677468206d7573742062652067726561746572207468616e60448201527510191a9b1030b7321030903837bbb2b91037b310191760511b606482015260840162000c6c565b600062002e7e610100836200418e565b90506000816001600160401b0381111562002e9d5762002e9d62003b92565b60405190808252806020026020018201604052801562002ec7578160200160208202803683370190505b50905060005b8281101562002f1c576101008181028701602001208251819084908490811062002efb5762002efb62004224565b6020908102919091010152508062002f13816200423a565b91505062002ecd565b505b60018211156200301457600062002f376002846200418e565b905060005b818110156200300b5762002fd68362002f57836002620041bb565b8151811062002f6a5762002f6a62004224565b60200260200101518483600262002f829190620041bb565b62002f8f906001620041a5565b8151811062002fa25762002fa262004224565b6020026020010151604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b83828151811062002feb5762002feb62004224565b60209081029190910101528062003002816200423a565b91505062002f3c565b50915062002f1e565b806000815181106200302a576200302a62004224565b60200260200101519350505050919050565b8360005b83518110156200310d57600084828151811062003061576200306162004224565b602002602001015190506002846200307a91906200441a565b6001600160401b0316600003620030bb57604080516020808201869052818301849052825180830384018152606090920190925280519101209250620030e6565b6040805160208082018490528183018690528251808303840181526060909201909252805191012092505b6001846001600160401b0316901c935050808062003104906200423a565b91505062003040565b5083811462000bf85760405162461bcd60e51b8152602060048201526013602482015272125b98dbdc9c9958dd08189b1bd888149bdbdd606a1b604482015260640162000c6c565b806200316183620036c8565b1462000b3d5760405162461bcd60e51b8152602060048201526012602482015271125b98dbdc9c9958dd0819185d18549bdbdd60721b604482015260640162000c6c565b6000612710604954603f54620031bc9190620041bb565b620031c891906200418e565b905080603f54620031da9190620041dd565b60406000828254620031ed9190620041a5565b90915550506000603f55801562001f0d57604a5462001f0d906001600160a01b03168262003715565b60008060006020603c54901c9050603e54603d54111562003282576008603e54603e54603d54620032489190620041dd565b620032549084620041bb565b6200326091906200418e565b6200326c91906200418e565b91506200327a8282620041dd565b9050620032ce565b6008603e54603d54603e54620032999190620041dd565b620032a59084620041bb565b620032b191906200418e565b620032bd91906200418e565b9150620032cb8282620041a5565b90505b6020620032df60806000196200418e565b901c8110620032fe57620032f760806000196200418e565b9250505090565b60201b92915050565b606062000aa36001600160a01b03831660145b606060006200332b836002620041bb565b62003338906002620041a5565b6001600160401b0381111562003352576200335262003b92565b6040519080825280601f01601f1916602001820160405280156200337d576020820181803683370190505b509050600360fc1b816000815181106200339b576200339b62004224565b60200101906001600160f81b031916908160001a905350600f60fb1b81600181518110620033cd57620033cd62004224565b60200101906001600160f81b031916908160001a9053506000620033f3846002620041bb565b62003400906001620041a5565b90505b600181111562003482576f181899199a1a9b1b9c1cb0b131b232b360811b85600f166010811062003438576200343862004224565b1a60f81b82828151811062003451576200345162004224565b60200101906001600160f81b031916908160001a90535060049490941c936200347a8162004523565b905062003403565b5083156200112b5760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e74604482015260640162000c6c565b60008181526001830160205260408120546200351c5750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915562000aa3565b50600062000aa3565b600081815260018301602052604081205480156200361e5760006200354c600183620041dd565b85549091506000906200356290600190620041dd565b9050818114620035ce57600086600001828154811062003586576200358662004224565b9060005260206000200154905080876000018481548110620035ac57620035ac62004224565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080620035e257620035e26200453d565b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505062000aa3565b600091505062000aa3565b5092915050565b6040805180820190915260008082526020820152600080806200366360008051602062004ae0833981519152866200420d565b90505b620036718162003834565b909350915060008051602062004ae08339815191528283098303620036ac576040805180820190915290815260208101919091529392505050565b60008051602062004ae083398151915260018208905062003666565b805160009062000aa3906200370c908460015b6020020151604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b836002620036db565b80471015620037675760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e6365000000604482015260640162000c6c565b6000826001600160a01b03168260405160006040518083038185875af1925050503d8060008114620037b6576040519150601f19603f3d011682016040523d82523d6000602084013e620037bb565b606091505b505090508062000b965760405162461bcd60e51b815260206004820152603a60248201527f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260448201527f6563697069656e74206d61792068617665207265766572746564000000000000606482015260840162000c6c565b6000808060008051602062004ae0833981519152600360008051602062004ae08339815191528660008051602062004ae0833981519152888909090890506000620038b0827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260008051602062004ae0833981519152620038bc565b91959194509092505050565b600080620038c9620039fe565b620038d362003a1c565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa925082806200391157fe5b5082620039615760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c757265000000000000604482015260640162000c6c565b505195945050505050565b61058c806200455483390190565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b6040518060400160405280620039cb62003a3a565b8152602001620039da62003a3a565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b60006020828403121562003a6b57600080fd5b81356001600160e01b0319811681146200112b57600080fd5b60006020828403121562003a9757600080fd5b5035919050565b80356001600160401b038116811462003ab657600080fd5b919050565b60006020828403121562003ace57600080fd5b6200112b8262003a9e565b6001600160a01b038116811462001f0d57600080fd5b6000806040838503121562003b0357600080fd5b82359150602083013562003b178162003ad9565b809150509250929050565b60006020828403121562003b3557600080fd5b81356200112b8162003ad9565b60008060006060848603121562003b5857600080fd5b505081359360208301359350604090920135919050565b6000806040838503121562003b8357600080fd5b50508035926020909101359150565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b038111828210171562003bcd5762003bcd62003b92565b60405290565b60405160e081016001600160401b038111828210171562003bcd5762003bcd62003b92565b60405161014081016001600160401b038111828210171562003bcd5762003bcd62003b92565b604051601f8201601f191681016001600160401b038111828210171562003c495762003c4962003b92565b604052919050565b60006001600160401b0382111562003c6d5762003c6d62003b92565b5060051b60200190565b600082601f83011262003c8957600080fd5b8135602062003ca262003c9c8362003c51565b62003c1e565b82815260059290921b8401810191818101908684111562003cc257600080fd5b8286015b8481101562003cdf578035835291830191830162003cc6565b509695505050505050565b60006020828403121562003cfd57600080fd5b81356001600160401b0381111562003d1457600080fd5b62003d228482850162003c77565b949350505050565b60006040828403121562003d3d57600080fd5b62003d4762003ba8565b9050813581526020820135602082015292915050565b600082601f83011262003d6f57600080fd5b81356001600160401b0381111562003d8b5762003d8b62003b92565b62003da0601f8201601f191660200162003c1e565b81815284602083860101111562003db657600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f83011262003de557600080fd5b604051604081018181106001600160401b038211171562003e0a5762003e0a62003b92565b806040525080604084018581111562003e2257600080fd5b845b8181101562003e3e57803583526020928301920162003e24565b509195945050505050565b60006080828403121562003e5c57600080fd5b62003e6662003ba8565b905062003e74838362003dd3565b815262003e85836040840162003dd3565b602082015292915050565b6000602080838503121562003ea457600080fd5b82356001600160401b038082111562003ebc57600080fd5b818501915085601f83011262003ed157600080fd5b813562003ee262003c9c8262003c51565b81815260059190911b8301840190848101908883111562003f0257600080fd5b8585015b8381101562003fd55780358581111562003f205760008081fd5b8601610180818c03601f1901121562003f395760008081fd5b62003f4362003bd3565b8882013581526040808301358a830152606080840135828401526080915062003f6f8e83860162003d2a565b9083015260c0838101358981111562003f885760008081fd5b62003f988f8d8388010162003d5d565b838501525062003fac8e60e0860162003e49565b60a084015262003fc18e610160860162003d2a565b908301525084525091860191860162003f06565b5098975050505050505050565b803563ffffffff8116811462003ab657600080fd5b600082601f8301126200400957600080fd5b604051606081018181106001600160401b03821117156200402e576200402e62003b92565b60405280606084018581111562003e2257600080fd5b6000602082840312156200405757600080fd5b81356001600160401b03808211156200406f57600080fd5b9083019061018082860312156200408557600080fd5b6200408f62003bf8565b82358152620040a16020840162003a9e565b6020820152620040b46040840162003a9e565b6040820152620040c76060840162003fe2565b6060820152620040da6080840162003fe2565b608082015260a083013560a082015260c083013560c0820152620041028660e0850162003ff7565b60e0820152610140830135828111156200411b57600080fd5b620041298782860162003c77565b61010083015250610160830135828111156200414457600080fd5b620041528782860162003d5d565b6101208301525095945050505050565b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600082620041a057620041a062004162565b500490565b8082018082111562000aa35762000aa362004178565b6000816000190483118215151615620041d857620041d862004178565b500290565b8181038181111562000aa35762000aa362004178565b6000602082840312156200420657600080fd5b5051919050565b6000826200421f576200421f62004162565b500690565b634e487b7160e01b600052603260045260246000fd5b6000600182016200424f576200424f62004178565b5060010190565b60005b838110156200427357818101518382015260200162004259565b50506000910152565b600081518084526200429681602086016020860162004256565b601f01601f19169290920160200192915050565b838152826020820152606060408201526000620042cb60608301846200427c565b95945050505050565b60008060008385036080811215620042eb57600080fd5b6040811215620042fa57600080fd5b506200430562003ba8565b8451815260208086015190820152604085015160609095015190969495509392505050565b6000602082840312156200433d57600080fd5b81516200112b8162003ad9565b8060005b6002811015620015145781518452602093840193909101906001016200434e565b8881528760208201528660408201528560608201526200439360808201866200434a565b620043a260c08201856200434a565b610100810192909252610120820152610140019695505050505050565b60006001600160401b0380831681851681830481118215151615620043e857620043e862004178565b02949350505050565b60006001600160401b03808416806200440e576200440e62004162565b92169190910492915050565b60006001600160401b038084168062004437576200443762004162565b92169190910692915050565b6001600160401b0381811683821601908082111562003629576200362962004178565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000815260008351620044a081601785016020880162004256565b7001034b99036b4b9b9b4b733903937b6329607d1b6017918401918201528351620044d381602884016020880162004256565b01602801949350505050565b6020815260006200112b60208301846200427c565b838152826020820152600082516200451481604085016020870162004256565b91909101604001949350505050565b60008162004535576200453562004178565b506000190190565b634e487b7160e01b600052603160045260246000fdfe608060405234801561001057600080fd5b5061001a3361001f565b61006f565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b61050e8061007e6000396000f3fe6080604052600436106100555760003560e01c806351cff8d91461005a578063715018a61461007c5780638da5cb5b14610091578063e3a9db1a146100be578063f2fde38b14610102578063f340fa0114610122575b600080fd5b34801561006657600080fd5b5061007a61007536600461048d565b610135565b005b34801561008857600080fd5b5061007a6101ac565b34801561009d57600080fd5b506000546040516001600160a01b0390911681526020015b60405180910390f35b3480156100ca57600080fd5b506100f46100d936600461048d565b6001600160a01b031660009081526001602052604090205490565b6040519081526020016100b5565b34801561010e57600080fd5b5061007a61011d36600461048d565b6101c0565b61007a61013036600461048d565b61023e565b61013d6102b0565b6001600160a01b0381166000818152600160205260408120805491905590610165908261030a565b816001600160a01b03167f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5826040516101a091815260200190565b60405180910390a25050565b6101b46102b0565b6101be6000610428565b565b6101c86102b0565b6001600160a01b0381166102325760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b61023b81610428565b50565b6102466102b0565b6001600160a01b0381166000908152600160205260408120805434928392916102709084906104b1565b90915550506040518181526001600160a01b038316907f2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4906020016101a0565b6000546001600160a01b031633146101be5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610229565b8047101561035a5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e63650000006044820152606401610229565b6000826001600160a01b03168260405160006040518083038185875af1925050503d80600081146103a7576040519150601f19603f3d011682016040523d82523d6000602084013e6103ac565b606091505b50509050806104235760405162461bcd60e51b815260206004820152603a60248201527f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260448201527f6563697069656e74206d617920686176652072657665727465640000000000006064820152608401610229565b505050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b038116811461023b57600080fd5b60006020828403121561049f57600080fd5b81356104aa81610478565b9392505050565b808201808211156104d257634e487b7160e01b600052601160045260246000fd5b9291505056fea26469706673582212206fb4f990ff997db88feee170cb2b7ab9937e85cb31861a96846a4d809d251fab64736f6c6343000810003330644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47b9d69e0ca90be54a40811e436234a7f7908b87ff2bec27e64f878b166da8e8e5a264697066735822122091cc965a5705bbb6f884c52bcb82f050e2c59df0e8ea1997efabecb421f042be64736f6c63430008100033",
}

// DAEntranceABI is the input ABI used to generate the binding from.
// Deprecated: Use DAEntranceMetaData.ABI instead.
var DAEntranceABI = DAEntranceMetaData.ABI

// DAEntranceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DAEntranceMetaData.Bin instead.
var DAEntranceBin = DAEntranceMetaData.Bin

// DeployDAEntrance deploys a new Ethereum contract, binding an instance of DAEntrance to it.
func DeployDAEntrance(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DAEntrance, error) {
	parsed, err := DAEntranceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DAEntranceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DAEntrance{DAEntranceCaller: DAEntranceCaller{contract: contract}, DAEntranceTransactor: DAEntranceTransactor{contract: contract}, DAEntranceFilterer: DAEntranceFilterer{contract: contract}}, nil
}

// DAEntrance is an auto generated Go binding around an Ethereum contract.
type DAEntrance struct {
	DAEntranceCaller     // Read-only binding to the contract
	DAEntranceTransactor // Write-only binding to the contract
	DAEntranceFilterer   // Log filterer for contract events
}

// DAEntranceCaller is an auto generated read-only Go binding around an Ethereum contract.
type DAEntranceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DAEntranceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DAEntranceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DAEntranceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DAEntranceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DAEntranceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DAEntranceSession struct {
	Contract     *DAEntrance       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DAEntranceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DAEntranceCallerSession struct {
	Contract *DAEntranceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DAEntranceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DAEntranceTransactorSession struct {
	Contract     *DAEntranceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DAEntranceRaw is an auto generated low-level Go binding around an Ethereum contract.
type DAEntranceRaw struct {
	Contract *DAEntrance // Generic contract binding to access the raw methods on
}

// DAEntranceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DAEntranceCallerRaw struct {
	Contract *DAEntranceCaller // Generic read-only contract binding to access the raw methods on
}

// DAEntranceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DAEntranceTransactorRaw struct {
	Contract *DAEntranceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDAEntrance creates a new instance of DAEntrance, bound to a specific deployed contract.
func NewDAEntrance(address common.Address, backend bind.ContractBackend) (*DAEntrance, error) {
	contract, err := bindDAEntrance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DAEntrance{DAEntranceCaller: DAEntranceCaller{contract: contract}, DAEntranceTransactor: DAEntranceTransactor{contract: contract}, DAEntranceFilterer: DAEntranceFilterer{contract: contract}}, nil
}

// NewDAEntranceCaller creates a new read-only instance of DAEntrance, bound to a specific deployed contract.
func NewDAEntranceCaller(address common.Address, caller bind.ContractCaller) (*DAEntranceCaller, error) {
	contract, err := bindDAEntrance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DAEntranceCaller{contract: contract}, nil
}

// NewDAEntranceTransactor creates a new write-only instance of DAEntrance, bound to a specific deployed contract.
func NewDAEntranceTransactor(address common.Address, transactor bind.ContractTransactor) (*DAEntranceTransactor, error) {
	contract, err := bindDAEntrance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DAEntranceTransactor{contract: contract}, nil
}

// NewDAEntranceFilterer creates a new log filterer instance of DAEntrance, bound to a specific deployed contract.
func NewDAEntranceFilterer(address common.Address, filterer bind.ContractFilterer) (*DAEntranceFilterer, error) {
	contract, err := bindDAEntrance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DAEntranceFilterer{contract: contract}, nil
}

// bindDAEntrance binds a generic wrapper to an already deployed contract.
func bindDAEntrance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DAEntranceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DAEntrance *DAEntranceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DAEntrance.Contract.DAEntranceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DAEntrance *DAEntranceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DAEntrance.Contract.DAEntranceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DAEntrance *DAEntranceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DAEntrance.Contract.DAEntranceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DAEntrance *DAEntranceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DAEntrance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DAEntrance *DAEntranceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DAEntrance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DAEntrance *DAEntranceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DAEntrance.Contract.contract.Transact(opts, method, params...)
}

// DASIGNERS is a free data retrieval call binding the contract method 0x807f063a.
//
// Solidity: function DA_SIGNERS() view returns(address)
func (_DAEntrance *DAEntranceCaller) DASIGNERS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DAEntrance.contract.Call(opts, &out, "DA_SIGNERS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DASIGNERS is a free data retrieval call binding the contract method 0x807f063a.
//
// Solidity: function DA_SIGNERS() view returns(address)
func (_DAEntrance *DAEntranceSession) DASIGNERS() (common.Address, error) {
	return _DAEntrance.Contract.DASIGNERS(&_DAEntrance.CallOpts)
}

// DASIGNERS is a free data retrieval call binding the contract method 0x807f063a.
//
// Solidity: function DA_SIGNERS() view returns(address)
func (_DAEntrance *DAEntranceCallerSession) DASIGNERS() (common.Address, error) {
	return _DAEntrance.Contract.DASIGNERS(&_DAEntrance.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DAEntrance *DAEntranceCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DAEntrance.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DAEntrance *DAEntranceSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _DAEntrance.Contract.DEFAULTADMINROLE(&_DAEntrance.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DAEntrance *DAEntranceCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _DAEntrance.Contract.DEFAULTADMINROLE(&_DAEntrance.CallOpts)
}

// MAXPODASTARGET is a free data retrieval call binding the contract method 0x8bdcc712.
//
// Solidity: function MAX_PODAS_TARGET() view returns(uint256)
func (_DAEntrance *DAEntranceCaller) MAXPODASTARGET(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DAEntrance.contract.Call(opts, &out, "MAX_PODAS_TARGET")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPODASTARGET is a free data retrieval call binding the contract method 0x8bdcc712.
//
// Solidity: function MAX_PODAS_TARGET() view returns(uint256)
func (_DAEntrance *DAEntranceSession) MAXPODASTARGET() (*big.Int, error) {
	return _DAEntrance.Contract.MAXPODASTARGET(&_DAEntrance.CallOpts)
}

// MAXPODASTARGET is a free data retrieval call binding the contract method 0x8bdcc712.
//
// Solidity: function MAX_PODAS_TARGET() view returns(uint256)
func (_DAEntrance *DAEntranceCallerSession) MAXPODASTARGET() (*big.Int, error) {
	return _DAEntrance.Contract.MAXPODASTARGET(&_DAEntrance.CallOpts)
}

// PARAMSADMINROLE is a free data retrieval call binding the contract method 0xb15d20da.
//
// Solidity: function PARAMS_ADMIN_ROLE() view returns(bytes32)
func (_DAEntrance *DAEntranceCaller) PARAMSADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DAEntrance.contract.Call(opts, &out, "PARAMS_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PARAMSADMINROLE is a free data retrieval call binding the contract method 0xb15d20da.
//
// Solidity: function PARAMS_ADMIN_ROLE() view returns(bytes32)
func (_DAEntrance *DAEntranceSession) PARAMSADMINROLE() ([32]byte, error) {
	return _DAEntrance.Contract.PARAMSADMINROLE(&_DAEntrance.CallOpts)
}

// PARAMSADMINROLE is a free data retrieval call binding the contract method 0xb15d20da.
//
// Solidity: function PARAMS_ADMIN_ROLE() view returns(bytes32)
func (_DAEntrance *DAEntranceCallerSession) PARAMSADMINROLE() ([32]byte, error) {
	return _DAEntrance.Contract.PARAMSADMINROLE(&_DAEntrance.CallOpts)
}

// SLICEDENOMINATOR is a free data retrieval call binding the contract method 0x602b245a.
//
// Solidity: function SLICE_DENOMINATOR() view returns(uint256)
func (_DAEntrance *DAEntranceCaller) SLICEDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DAEntrance.contract.Call(opts, &out, "SLICE_DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SLICEDENOMINATOR is a free data retrieval call binding the contract method 0x602b245a.
//
// Solidity: function SLICE_DENOMINATOR() view returns(uint256)
func (_DAEntrance *DAEntranceSession) SLICEDENOMINATOR() (*big.Int, error) {
	return _DAEntrance.Contract.SLICEDENOMINATOR(&_DAEntrance.CallOpts)
}

// SLICEDENOMINATOR is a free data retrieval call binding the contract method 0x602b245a.
//
// Solidity: function SLICE_DENOMINATOR() view returns(uint256)
func (_DAEntrance *DAEntranceCallerSession) SLICEDENOMINATOR() (*big.Int, error) {
	return _DAEntrance.Contract.SLICEDENOMINATOR(&_DAEntrance.CallOpts)
}

// SLICENUMERATOR is a free data retrieval call binding the contract method 0x5626a47b.
//
// Solidity: function SLICE_NUMERATOR() view returns(uint256)
func (_DAEntrance *DAEntranceCaller) SLICENUMERATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DAEntrance.contract.Call(opts, &out, "SLICE_NUMERATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SLICENUMERATOR is a free data retrieval call binding the contract method 0x5626a47b.
//
// Solidity: function SLICE_NUMERATOR() view returns(uint256)
func (_DAEntrance *DAEntranceSession) SLICENUMERATOR() (*big.Int, error) {
	return _DAEntrance.Contract.SLICENUMERATOR(&_DAEntrance.CallOpts)
}

// SLICENUMERATOR is a free data retrieval call binding the contract method 0x5626a47b.
//
// Solidity: function SLICE_NUMERATOR() view returns(uint256)
func (_DAEntrance *DAEntranceCallerSession) SLICENUMERATOR() (*big.Int, error) {
	return _DAEntrance.Contract.SLICENUMERATOR(&_DAEntrance.CallOpts)
}

// ActivedReward is a free data retrieval call binding the contract method 0x12577052.
//
// Solidity: function activedReward() view returns(uint256)
func (_DAEntrance *DAEntranceCaller) ActivedReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DAEntrance.contract.Call(opts, &out, "activedReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActivedReward is a free data retrieval call binding the contract method 0x12577052.
//
// Solidity: function activedReward() view returns(uint256)
func (_DAEntrance *DAEntranceSession) ActivedReward() (*big.Int, error) {
	return _DAEntrance.Contract.ActivedReward(&_DAEntrance.CallOpts)
}

// ActivedReward is a free data retrieval call binding the contract method 0x12577052.
//
// Solidity: function activedReward() view returns(uint256)
func (_DAEntrance *DAEntranceCallerSession) ActivedReward() (*big.Int, error) {
	return _DAEntrance.Contract.ActivedReward(&_DAEntrance.CallOpts)
}

// BaseReward is a free data retrieval call binding the contract method 0x76ad03bc.
//
// Solidity: function baseReward() view returns(uint256)
func (_DAEntrance *DAEntranceCaller) BaseReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DAEntrance.contract.Call(opts, &out, "baseReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseReward is a free data retrieval call binding the contract method 0x76ad03bc.
//
// Solidity: function baseReward() view returns(uint256)
func (_DAEntrance *DAEntranceSession) BaseReward() (*big.Int, error) {
	return _DAEntrance.Contract.BaseReward(&_DAEntrance.CallOpts)
}

// BaseReward is a free data retrieval call binding the contract method 0x76ad03bc.
//
// Solidity: function baseReward() view returns(uint256)
func (_DAEntrance *DAEntranceCallerSession) BaseReward() (*big.Int, error) {
	return _DAEntrance.Contract.BaseReward(&_DAEntrance.CallOpts)
}

// BlobPrice is a free data retrieval call binding the contract method 0x3d00448a.
//
// Solidity: function blobPrice() view returns(uint256)
func (_DAEntrance *DAEntranceCaller) BlobPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DAEntrance.contract.Call(opts, &out, "blobPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlobPrice is a free data retrieval call binding the contract method 0x3d00448a.
//
// Solidity: function blobPrice() view returns(uint256)
func (_DAEntrance *DAEntranceSession) BlobPrice() (*big.Int, error) {
	return _DAEntrance.Contract.BlobPrice(&_DAEntrance.CallOpts)
}

// BlobPrice is a free data retrieval call binding the contract method 0x3d00448a.
//
// Solidity: function blobPrice() view returns(uint256)
func (_DAEntrance *DAEntranceCallerSession) BlobPrice() (*big.Int, error) {
	return _DAEntrance.Contract.BlobPrice(&_DAEntrance.CallOpts)
}

// CommitmentExists is a free data retrieval call binding the contract method 0x6a53525d.
//
// Solidity: function commitmentExists(bytes32 _dataRoot, uint256 _epoch, uint256 _quorumId) view returns(bool)
func (_DAEntrance *DAEntranceCaller) CommitmentExists(opts *bind.CallOpts, _dataRoot [32]byte, _epoch *big.Int, _quorumId *big.Int) (bool, error) {
	var out []interface{}
	err := _DAEntrance.contract.Call(opts, &out, "commitmentExists", _dataRoot, _epoch, _quorumId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CommitmentExists is a free data retrieval call binding the contract method 0x6a53525d.
//
// Solidity: function commitmentExists(bytes32 _dataRoot, uint256 _epoch, uint256 _quorumId) view returns(bool)
func (_DAEntrance *DAEntranceSession) CommitmentExists(_dataRoot [32]byte, _epoch *big.Int, _quorumId *big.Int) (bool, error) {
	return _DAEntrance.Contract.CommitmentExists(&_DAEntrance.CallOpts, _dataRoot, _epoch, _quorumId)
}

// CommitmentExists is a free data retrieval call binding the contract method 0x6a53525d.
//
// Solidity: function commitmentExists(bytes32 _dataRoot, uint256 _epoch, uint256 _quorumId) view returns(bool)
func (_DAEntrance *DAEntranceCallerSession) CommitmentExists(_dataRoot [32]byte, _epoch *big.Int, _quorumId *big.Int) (bool, error) {
	return _DAEntrance.Contract.CommitmentExists(&_DAEntrance.CallOpts, _dataRoot, _epoch, _quorumId)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_DAEntrance *DAEntranceCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DAEntrance.contract.Call(opts, &out, "currentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_DAEntrance *DAEntranceSession) CurrentEpoch() (*big.Int, error) {
	return _DAEntrance.Contract.CurrentEpoch(&_DAEntrance.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_DAEntrance *DAEntranceCallerSession) CurrentEpoch() (*big.Int, error) {
	return _DAEntrance.Contract.CurrentEpoch(&_DAEntrance.CallOpts)
}

// CurrentEpochReward is a free data retrieval call binding the contract method 0xe68e035b.
//
// Solidity: function currentEpochReward() view returns(uint256)
func (_DAEntrance *DAEntranceCaller) CurrentEpochReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DAEntrance.contract.Call(opts, &out, "currentEpochReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpochReward is a free data retrieval call binding the contract method 0xe68e035b.
//
// Solidity: function currentEpochReward() view returns(uint256)
func (_DAEntrance *DAEntranceSession) CurrentEpochReward() (*big.Int, error) {
	return _DAEntrance.Contract.CurrentEpochReward(&_DAEntrance.CallOpts)
}

// CurrentEpochReward is a free data retrieval call binding the contract method 0xe68e035b.
//
// Solidity: function currentEpochReward() view returns(uint256)
func (_DAEntrance *DAEntranceCallerSession) CurrentEpochReward() (*big.Int, error) {
	return _DAEntrance.Contract.CurrentEpochReward(&_DAEntrance.CallOpts)
}

// CurrentSampleSeed is a free data retrieval call binding the contract method 0x9fae1fa4.
//
// Solidity: function currentSampleSeed() view returns(bytes32)
func (_DAEntrance *DAEntranceCaller) CurrentSampleSeed(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DAEntrance.contract.Call(opts, &out, "currentSampleSeed")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentSampleSeed is a free data retrieval call binding the contract method 0x9fae1fa4.
//
// Solidity: function currentSampleSeed() view returns(bytes32)
func (_DAEntrance *DAEntranceSession) CurrentSampleSeed() ([32]byte, error) {
	return _DAEntrance.Contract.CurrentSampleSeed(&_DAEntrance.CallOpts)
}

// CurrentSampleSeed is a free data retrieval call binding the contract method 0x9fae1fa4.
//
// Solidity: function currentSampleSeed() view returns(bytes32)
func (_DAEntrance *DAEntranceCallerSession) CurrentSampleSeed() ([32]byte, error) {
	return _DAEntrance.Contract.CurrentSampleSeed(&_DAEntrance.CallOpts)
}

// EpochWindowSize is a free data retrieval call binding the contract method 0x3e898337.
//
// Solidity: function epochWindowSize() view returns(uint256)
func (_DAEntrance *DAEntranceCaller) EpochWindowSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DAEntrance.contract.Call(opts, &out, "epochWindowSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochWindowSize is a free data retrieval call binding the contract method 0x3e898337.
//
// Solidity: function epochWindowSize() view returns(uint256)
func (_DAEntrance *DAEntranceSession) EpochWindowSize() (*big.Int, error) {
	return _DAEntrance.Contract.EpochWindowSize(&_DAEntrance.CallOpts)
}

// EpochWindowSize is a free data retrieval call binding the contract method 0x3e898337.
//
// Solidity: function epochWindowSize() view returns(uint256)
func (_DAEntrance *DAEntranceCallerSession) EpochWindowSize() (*big.Int, error) {
	return _DAEntrance.Contract.EpochWindowSize(&_DAEntrance.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DAEntrance *DAEntranceCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _DAEntrance.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DAEntrance *DAEntranceSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _DAEntrance.Contract.GetRoleAdmin(&_DAEntrance.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DAEntrance *DAEntranceCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _DAEntrance.Contract.GetRoleAdmin(&_DAEntrance.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_DAEntrance *DAEntranceCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DAEntrance.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_DAEntrance *DAEntranceSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _DAEntrance.Contract.GetRoleMember(&_DAEntrance.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_DAEntrance *DAEntranceCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _DAEntrance.Contract.GetRoleMember(&_DAEntrance.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_DAEntrance *DAEntranceCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _DAEntrance.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_DAEntrance *DAEntranceSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _DAEntrance.Contract.GetRoleMemberCount(&_DAEntrance.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_DAEntrance *DAEntranceCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _DAEntrance.Contract.GetRoleMemberCount(&_DAEntrance.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DAEntrance *DAEntranceCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _DAEntrance.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DAEntrance *DAEntranceSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _DAEntrance.Contract.HasRole(&_DAEntrance.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DAEntrance *DAEntranceCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _DAEntrance.Contract.HasRole(&_DAEntrance.CallOpts, role, account)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_DAEntrance *DAEntranceCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DAEntrance.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_DAEntrance *DAEntranceSession) Initialized() (bool, error) {
	return _DAEntrance.Contract.Initialized(&_DAEntrance.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_DAEntrance *DAEntranceCallerSession) Initialized() (bool, error) {
	return _DAEntrance.Contract.Initialized(&_DAEntrance.CallOpts)
}

// NextSampleHeight is a free data retrieval call binding the contract method 0x7397eb33.
//
// Solidity: function nextSampleHeight() view returns(uint256)
func (_DAEntrance *DAEntranceCaller) NextSampleHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DAEntrance.contract.Call(opts, &out, "nextSampleHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextSampleHeight is a free data retrieval call binding the contract method 0x7397eb33.
//
// Solidity: function nextSampleHeight() view returns(uint256)
func (_DAEntrance *DAEntranceSession) NextSampleHeight() (*big.Int, error) {
	return _DAEntrance.Contract.NextSampleHeight(&_DAEntrance.CallOpts)
}

// NextSampleHeight is a free data retrieval call binding the contract method 0x7397eb33.
//
// Solidity: function nextSampleHeight() view returns(uint256)
func (_DAEntrance *DAEntranceCallerSession) NextSampleHeight() (*big.Int, error) {
	return _DAEntrance.Contract.NextSampleHeight(&_DAEntrance.CallOpts)
}

// Payments is a free data retrieval call binding the contract method 0xe2982c21.
//
// Solidity: function payments(address dest) view returns(uint256)
func (_DAEntrance *DAEntranceCaller) Payments(opts *bind.CallOpts, dest common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DAEntrance.contract.Call(opts, &out, "payments", dest)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Payments is a free data retrieval call binding the contract method 0xe2982c21.
//
// Solidity: function payments(address dest) view returns(uint256)
func (_DAEntrance *DAEntranceSession) Payments(dest common.Address) (*big.Int, error) {
	return _DAEntrance.Contract.Payments(&_DAEntrance.CallOpts, dest)
}

// Payments is a free data retrieval call binding the contract method 0xe2982c21.
//
// Solidity: function payments(address dest) view returns(uint256)
func (_DAEntrance *DAEntranceCallerSession) Payments(dest common.Address) (*big.Int, error) {
	return _DAEntrance.Contract.Payments(&_DAEntrance.CallOpts, dest)
}

// PodasTarget is a free data retrieval call binding the contract method 0xc8d3b359.
//
// Solidity: function podasTarget() view returns(uint256)
func (_DAEntrance *DAEntranceCaller) PodasTarget(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DAEntrance.contract.Call(opts, &out, "podasTarget")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PodasTarget is a free data retrieval call binding the contract method 0xc8d3b359.
//
// Solidity: function podasTarget() view returns(uint256)
func (_DAEntrance *DAEntranceSession) PodasTarget() (*big.Int, error) {
	return _DAEntrance.Contract.PodasTarget(&_DAEntrance.CallOpts)
}

// PodasTarget is a free data retrieval call binding the contract method 0xc8d3b359.
//
// Solidity: function podasTarget() view returns(uint256)
func (_DAEntrance *DAEntranceCallerSession) PodasTarget() (*big.Int, error) {
	return _DAEntrance.Contract.PodasTarget(&_DAEntrance.CallOpts)
}

// RewardRatio is a free data retrieval call binding the contract method 0x646033bc.
//
// Solidity: function rewardRatio() view returns(uint256)
func (_DAEntrance *DAEntranceCaller) RewardRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DAEntrance.contract.Call(opts, &out, "rewardRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRatio is a free data retrieval call binding the contract method 0x646033bc.
//
// Solidity: function rewardRatio() view returns(uint256)
func (_DAEntrance *DAEntranceSession) RewardRatio() (*big.Int, error) {
	return _DAEntrance.Contract.RewardRatio(&_DAEntrance.CallOpts)
}

// RewardRatio is a free data retrieval call binding the contract method 0x646033bc.
//
// Solidity: function rewardRatio() view returns(uint256)
func (_DAEntrance *DAEntranceCallerSession) RewardRatio() (*big.Int, error) {
	return _DAEntrance.Contract.RewardRatio(&_DAEntrance.CallOpts)
}

// RoundSubmissions is a free data retrieval call binding the contract method 0xff187748.
//
// Solidity: function roundSubmissions() view returns(uint256)
func (_DAEntrance *DAEntranceCaller) RoundSubmissions(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DAEntrance.contract.Call(opts, &out, "roundSubmissions")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoundSubmissions is a free data retrieval call binding the contract method 0xff187748.
//
// Solidity: function roundSubmissions() view returns(uint256)
func (_DAEntrance *DAEntranceSession) RoundSubmissions() (*big.Int, error) {
	return _DAEntrance.Contract.RoundSubmissions(&_DAEntrance.CallOpts)
}

// RoundSubmissions is a free data retrieval call binding the contract method 0xff187748.
//
// Solidity: function roundSubmissions() view returns(uint256)
func (_DAEntrance *DAEntranceCallerSession) RoundSubmissions() (*big.Int, error) {
	return _DAEntrance.Contract.RoundSubmissions(&_DAEntrance.CallOpts)
}

// SamplePeriod is a free data retrieval call binding the contract method 0x98920f57.
//
// Solidity: function samplePeriod() view returns(uint256)
func (_DAEntrance *DAEntranceCaller) SamplePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DAEntrance.contract.Call(opts, &out, "samplePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SamplePeriod is a free data retrieval call binding the contract method 0x98920f57.
//
// Solidity: function samplePeriod() view returns(uint256)
func (_DAEntrance *DAEntranceSession) SamplePeriod() (*big.Int, error) {
	return _DAEntrance.Contract.SamplePeriod(&_DAEntrance.CallOpts)
}

// SamplePeriod is a free data retrieval call binding the contract method 0x98920f57.
//
// Solidity: function samplePeriod() view returns(uint256)
func (_DAEntrance *DAEntranceCallerSession) SamplePeriod() (*big.Int, error) {
	return _DAEntrance.Contract.SamplePeriod(&_DAEntrance.CallOpts)
}

// SampleRound is a free data retrieval call binding the contract method 0x168a062c.
//
// Solidity: function sampleRound() view returns(uint256)
func (_DAEntrance *DAEntranceCaller) SampleRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DAEntrance.contract.Call(opts, &out, "sampleRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SampleRound is a free data retrieval call binding the contract method 0x168a062c.
//
// Solidity: function sampleRound() view returns(uint256)
func (_DAEntrance *DAEntranceSession) SampleRound() (*big.Int, error) {
	return _DAEntrance.Contract.SampleRound(&_DAEntrance.CallOpts)
}

// SampleRound is a free data retrieval call binding the contract method 0x168a062c.
//
// Solidity: function sampleRound() view returns(uint256)
func (_DAEntrance *DAEntranceCallerSession) SampleRound() (*big.Int, error) {
	return _DAEntrance.Contract.SampleRound(&_DAEntrance.CallOpts)
}

// ServiceFee is a free data retrieval call binding the contract method 0x8abdf5aa.
//
// Solidity: function serviceFee() view returns(uint256)
func (_DAEntrance *DAEntranceCaller) ServiceFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DAEntrance.contract.Call(opts, &out, "serviceFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ServiceFee is a free data retrieval call binding the contract method 0x8abdf5aa.
//
// Solidity: function serviceFee() view returns(uint256)
func (_DAEntrance *DAEntranceSession) ServiceFee() (*big.Int, error) {
	return _DAEntrance.Contract.ServiceFee(&_DAEntrance.CallOpts)
}

// ServiceFee is a free data retrieval call binding the contract method 0x8abdf5aa.
//
// Solidity: function serviceFee() view returns(uint256)
func (_DAEntrance *DAEntranceCallerSession) ServiceFee() (*big.Int, error) {
	return _DAEntrance.Contract.ServiceFee(&_DAEntrance.CallOpts)
}

// ServiceFeeRateBps is a free data retrieval call binding the contract method 0xc0575111.
//
// Solidity: function serviceFeeRateBps() view returns(uint256)
func (_DAEntrance *DAEntranceCaller) ServiceFeeRateBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DAEntrance.contract.Call(opts, &out, "serviceFeeRateBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ServiceFeeRateBps is a free data retrieval call binding the contract method 0xc0575111.
//
// Solidity: function serviceFeeRateBps() view returns(uint256)
func (_DAEntrance *DAEntranceSession) ServiceFeeRateBps() (*big.Int, error) {
	return _DAEntrance.Contract.ServiceFeeRateBps(&_DAEntrance.CallOpts)
}

// ServiceFeeRateBps is a free data retrieval call binding the contract method 0xc0575111.
//
// Solidity: function serviceFeeRateBps() view returns(uint256)
func (_DAEntrance *DAEntranceCallerSession) ServiceFeeRateBps() (*big.Int, error) {
	return _DAEntrance.Contract.ServiceFeeRateBps(&_DAEntrance.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DAEntrance *DAEntranceCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _DAEntrance.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DAEntrance *DAEntranceSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DAEntrance.Contract.SupportsInterface(&_DAEntrance.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DAEntrance *DAEntranceCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DAEntrance.Contract.SupportsInterface(&_DAEntrance.CallOpts, interfaceId)
}

// TargetRoundSubmissions is a free data retrieval call binding the contract method 0x2fc0534b.
//
// Solidity: function targetRoundSubmissions() view returns(uint256)
func (_DAEntrance *DAEntranceCaller) TargetRoundSubmissions(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DAEntrance.contract.Call(opts, &out, "targetRoundSubmissions")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TargetRoundSubmissions is a free data retrieval call binding the contract method 0x2fc0534b.
//
// Solidity: function targetRoundSubmissions() view returns(uint256)
func (_DAEntrance *DAEntranceSession) TargetRoundSubmissions() (*big.Int, error) {
	return _DAEntrance.Contract.TargetRoundSubmissions(&_DAEntrance.CallOpts)
}

// TargetRoundSubmissions is a free data retrieval call binding the contract method 0x2fc0534b.
//
// Solidity: function targetRoundSubmissions() view returns(uint256)
func (_DAEntrance *DAEntranceCallerSession) TargetRoundSubmissions() (*big.Int, error) {
	return _DAEntrance.Contract.TargetRoundSubmissions(&_DAEntrance.CallOpts)
}

// TargetRoundSubmissionsNext is a free data retrieval call binding the contract method 0x257a3aa3.
//
// Solidity: function targetRoundSubmissionsNext() view returns(uint256)
func (_DAEntrance *DAEntranceCaller) TargetRoundSubmissionsNext(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DAEntrance.contract.Call(opts, &out, "targetRoundSubmissionsNext")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TargetRoundSubmissionsNext is a free data retrieval call binding the contract method 0x257a3aa3.
//
// Solidity: function targetRoundSubmissionsNext() view returns(uint256)
func (_DAEntrance *DAEntranceSession) TargetRoundSubmissionsNext() (*big.Int, error) {
	return _DAEntrance.Contract.TargetRoundSubmissionsNext(&_DAEntrance.CallOpts)
}

// TargetRoundSubmissionsNext is a free data retrieval call binding the contract method 0x257a3aa3.
//
// Solidity: function targetRoundSubmissionsNext() view returns(uint256)
func (_DAEntrance *DAEntranceCallerSession) TargetRoundSubmissionsNext() (*big.Int, error) {
	return _DAEntrance.Contract.TargetRoundSubmissionsNext(&_DAEntrance.CallOpts)
}

// TotalBaseReward is a free data retrieval call binding the contract method 0x7f1b5e43.
//
// Solidity: function totalBaseReward() view returns(uint256)
func (_DAEntrance *DAEntranceCaller) TotalBaseReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DAEntrance.contract.Call(opts, &out, "totalBaseReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBaseReward is a free data retrieval call binding the contract method 0x7f1b5e43.
//
// Solidity: function totalBaseReward() view returns(uint256)
func (_DAEntrance *DAEntranceSession) TotalBaseReward() (*big.Int, error) {
	return _DAEntrance.Contract.TotalBaseReward(&_DAEntrance.CallOpts)
}

// TotalBaseReward is a free data retrieval call binding the contract method 0x7f1b5e43.
//
// Solidity: function totalBaseReward() view returns(uint256)
func (_DAEntrance *DAEntranceCallerSession) TotalBaseReward() (*big.Int, error) {
	return _DAEntrance.Contract.TotalBaseReward(&_DAEntrance.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_DAEntrance *DAEntranceCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DAEntrance.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_DAEntrance *DAEntranceSession) Treasury() (common.Address, error) {
	return _DAEntrance.Contract.Treasury(&_DAEntrance.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_DAEntrance *DAEntranceCallerSession) Treasury() (common.Address, error) {
	return _DAEntrance.Contract.Treasury(&_DAEntrance.CallOpts)
}

// VerifiedErasureCommitment is a free data retrieval call binding the contract method 0x9da3a69b.
//
// Solidity: function verifiedErasureCommitment(bytes32 _dataRoot, uint256 _epoch, uint256 _quorumId) view returns((uint256,uint256))
func (_DAEntrance *DAEntranceCaller) VerifiedErasureCommitment(opts *bind.CallOpts, _dataRoot [32]byte, _epoch *big.Int, _quorumId *big.Int) (BN254G1Point, error) {
	var out []interface{}
	err := _DAEntrance.contract.Call(opts, &out, "verifiedErasureCommitment", _dataRoot, _epoch, _quorumId)

	if err != nil {
		return *new(BN254G1Point), err
	}

	out0 := *abi.ConvertType(out[0], new(BN254G1Point)).(*BN254G1Point)

	return out0, err

}

// VerifiedErasureCommitment is a free data retrieval call binding the contract method 0x9da3a69b.
//
// Solidity: function verifiedErasureCommitment(bytes32 _dataRoot, uint256 _epoch, uint256 _quorumId) view returns((uint256,uint256))
func (_DAEntrance *DAEntranceSession) VerifiedErasureCommitment(_dataRoot [32]byte, _epoch *big.Int, _quorumId *big.Int) (BN254G1Point, error) {
	return _DAEntrance.Contract.VerifiedErasureCommitment(&_DAEntrance.CallOpts, _dataRoot, _epoch, _quorumId)
}

// VerifiedErasureCommitment is a free data retrieval call binding the contract method 0x9da3a69b.
//
// Solidity: function verifiedErasureCommitment(bytes32 _dataRoot, uint256 _epoch, uint256 _quorumId) view returns((uint256,uint256))
func (_DAEntrance *DAEntranceCallerSession) VerifiedErasureCommitment(_dataRoot [32]byte, _epoch *big.Int, _quorumId *big.Int) (BN254G1Point, error) {
	return _DAEntrance.Contract.VerifiedErasureCommitment(&_DAEntrance.CallOpts, _dataRoot, _epoch, _quorumId)
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_DAEntrance *DAEntranceTransactor) Donate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DAEntrance.contract.Transact(opts, "donate")
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_DAEntrance *DAEntranceSession) Donate() (*types.Transaction, error) {
	return _DAEntrance.Contract.Donate(&_DAEntrance.TransactOpts)
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_DAEntrance *DAEntranceTransactorSession) Donate() (*types.Transaction, error) {
	return _DAEntrance.Contract.Donate(&_DAEntrance.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DAEntrance *DAEntranceTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DAEntrance.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DAEntrance *DAEntranceSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DAEntrance.Contract.GrantRole(&_DAEntrance.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DAEntrance *DAEntranceTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DAEntrance.Contract.GrantRole(&_DAEntrance.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_DAEntrance *DAEntranceTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DAEntrance.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_DAEntrance *DAEntranceSession) Initialize() (*types.Transaction, error) {
	return _DAEntrance.Contract.Initialize(&_DAEntrance.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_DAEntrance *DAEntranceTransactorSession) Initialize() (*types.Transaction, error) {
	return _DAEntrance.Contract.Initialize(&_DAEntrance.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_DAEntrance *DAEntranceTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DAEntrance.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_DAEntrance *DAEntranceSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DAEntrance.Contract.RenounceRole(&_DAEntrance.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_DAEntrance *DAEntranceTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DAEntrance.Contract.RenounceRole(&_DAEntrance.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DAEntrance *DAEntranceTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DAEntrance.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DAEntrance *DAEntranceSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DAEntrance.Contract.RevokeRole(&_DAEntrance.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DAEntrance *DAEntranceTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DAEntrance.Contract.RevokeRole(&_DAEntrance.TransactOpts, role, account)
}

// SampleRange is a paid mutator transaction binding the contract method 0x6efc2555.
//
// Solidity: function sampleRange() returns((uint64,uint64))
func (_DAEntrance *DAEntranceTransactor) SampleRange(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DAEntrance.contract.Transact(opts, "sampleRange")
}

// SampleRange is a paid mutator transaction binding the contract method 0x6efc2555.
//
// Solidity: function sampleRange() returns((uint64,uint64))
func (_DAEntrance *DAEntranceSession) SampleRange() (*types.Transaction, error) {
	return _DAEntrance.Contract.SampleRange(&_DAEntrance.TransactOpts)
}

// SampleRange is a paid mutator transaction binding the contract method 0x6efc2555.
//
// Solidity: function sampleRange() returns((uint64,uint64))
func (_DAEntrance *DAEntranceTransactorSession) SampleRange() (*types.Transaction, error) {
	return _DAEntrance.Contract.SampleRange(&_DAEntrance.TransactOpts)
}

// SampleTask is a paid mutator transaction binding the contract method 0x988ea94e.
//
// Solidity: function sampleTask() returns((bytes32,uint256,uint64))
func (_DAEntrance *DAEntranceTransactor) SampleTask(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DAEntrance.contract.Transact(opts, "sampleTask")
}

// SampleTask is a paid mutator transaction binding the contract method 0x988ea94e.
//
// Solidity: function sampleTask() returns((bytes32,uint256,uint64))
func (_DAEntrance *DAEntranceSession) SampleTask() (*types.Transaction, error) {
	return _DAEntrance.Contract.SampleTask(&_DAEntrance.TransactOpts)
}

// SampleTask is a paid mutator transaction binding the contract method 0x988ea94e.
//
// Solidity: function sampleTask() returns((bytes32,uint256,uint64))
func (_DAEntrance *DAEntranceTransactorSession) SampleTask() (*types.Transaction, error) {
	return _DAEntrance.Contract.SampleTask(&_DAEntrance.TransactOpts)
}

// SetBaseReward is a paid mutator transaction binding the contract method 0x0373a23a.
//
// Solidity: function setBaseReward(uint256 _baseReward) returns()
func (_DAEntrance *DAEntranceTransactor) SetBaseReward(opts *bind.TransactOpts, _baseReward *big.Int) (*types.Transaction, error) {
	return _DAEntrance.contract.Transact(opts, "setBaseReward", _baseReward)
}

// SetBaseReward is a paid mutator transaction binding the contract method 0x0373a23a.
//
// Solidity: function setBaseReward(uint256 _baseReward) returns()
func (_DAEntrance *DAEntranceSession) SetBaseReward(_baseReward *big.Int) (*types.Transaction, error) {
	return _DAEntrance.Contract.SetBaseReward(&_DAEntrance.TransactOpts, _baseReward)
}

// SetBaseReward is a paid mutator transaction binding the contract method 0x0373a23a.
//
// Solidity: function setBaseReward(uint256 _baseReward) returns()
func (_DAEntrance *DAEntranceTransactorSession) SetBaseReward(_baseReward *big.Int) (*types.Transaction, error) {
	return _DAEntrance.Contract.SetBaseReward(&_DAEntrance.TransactOpts, _baseReward)
}

// SetBlobPrice is a paid mutator transaction binding the contract method 0x23dd60a6.
//
// Solidity: function setBlobPrice(uint256 _blobPrice) returns()
func (_DAEntrance *DAEntranceTransactor) SetBlobPrice(opts *bind.TransactOpts, _blobPrice *big.Int) (*types.Transaction, error) {
	return _DAEntrance.contract.Transact(opts, "setBlobPrice", _blobPrice)
}

// SetBlobPrice is a paid mutator transaction binding the contract method 0x23dd60a6.
//
// Solidity: function setBlobPrice(uint256 _blobPrice) returns()
func (_DAEntrance *DAEntranceSession) SetBlobPrice(_blobPrice *big.Int) (*types.Transaction, error) {
	return _DAEntrance.Contract.SetBlobPrice(&_DAEntrance.TransactOpts, _blobPrice)
}

// SetBlobPrice is a paid mutator transaction binding the contract method 0x23dd60a6.
//
// Solidity: function setBlobPrice(uint256 _blobPrice) returns()
func (_DAEntrance *DAEntranceTransactorSession) SetBlobPrice(_blobPrice *big.Int) (*types.Transaction, error) {
	return _DAEntrance.Contract.SetBlobPrice(&_DAEntrance.TransactOpts, _blobPrice)
}

// SetEpochWindowSize is a paid mutator transaction binding the contract method 0xb1be17ab.
//
// Solidity: function setEpochWindowSize(uint64 _epochWindowSize) returns()
func (_DAEntrance *DAEntranceTransactor) SetEpochWindowSize(opts *bind.TransactOpts, _epochWindowSize uint64) (*types.Transaction, error) {
	return _DAEntrance.contract.Transact(opts, "setEpochWindowSize", _epochWindowSize)
}

// SetEpochWindowSize is a paid mutator transaction binding the contract method 0xb1be17ab.
//
// Solidity: function setEpochWindowSize(uint64 _epochWindowSize) returns()
func (_DAEntrance *DAEntranceSession) SetEpochWindowSize(_epochWindowSize uint64) (*types.Transaction, error) {
	return _DAEntrance.Contract.SetEpochWindowSize(&_DAEntrance.TransactOpts, _epochWindowSize)
}

// SetEpochWindowSize is a paid mutator transaction binding the contract method 0xb1be17ab.
//
// Solidity: function setEpochWindowSize(uint64 _epochWindowSize) returns()
func (_DAEntrance *DAEntranceTransactorSession) SetEpochWindowSize(_epochWindowSize uint64) (*types.Transaction, error) {
	return _DAEntrance.Contract.SetEpochWindowSize(&_DAEntrance.TransactOpts, _epochWindowSize)
}

// SetRewardRatio is a paid mutator transaction binding the contract method 0x3bab2a70.
//
// Solidity: function setRewardRatio(uint64 _rewardRatio) returns()
func (_DAEntrance *DAEntranceTransactor) SetRewardRatio(opts *bind.TransactOpts, _rewardRatio uint64) (*types.Transaction, error) {
	return _DAEntrance.contract.Transact(opts, "setRewardRatio", _rewardRatio)
}

// SetRewardRatio is a paid mutator transaction binding the contract method 0x3bab2a70.
//
// Solidity: function setRewardRatio(uint64 _rewardRatio) returns()
func (_DAEntrance *DAEntranceSession) SetRewardRatio(_rewardRatio uint64) (*types.Transaction, error) {
	return _DAEntrance.Contract.SetRewardRatio(&_DAEntrance.TransactOpts, _rewardRatio)
}

// SetRewardRatio is a paid mutator transaction binding the contract method 0x3bab2a70.
//
// Solidity: function setRewardRatio(uint64 _rewardRatio) returns()
func (_DAEntrance *DAEntranceTransactorSession) SetRewardRatio(_rewardRatio uint64) (*types.Transaction, error) {
	return _DAEntrance.Contract.SetRewardRatio(&_DAEntrance.TransactOpts, _rewardRatio)
}

// SetRoundSubmissions is a paid mutator transaction binding the contract method 0x88521ec7.
//
// Solidity: function setRoundSubmissions(uint64 _targetRoundSubmissions) returns()
func (_DAEntrance *DAEntranceTransactor) SetRoundSubmissions(opts *bind.TransactOpts, _targetRoundSubmissions uint64) (*types.Transaction, error) {
	return _DAEntrance.contract.Transact(opts, "setRoundSubmissions", _targetRoundSubmissions)
}

// SetRoundSubmissions is a paid mutator transaction binding the contract method 0x88521ec7.
//
// Solidity: function setRoundSubmissions(uint64 _targetRoundSubmissions) returns()
func (_DAEntrance *DAEntranceSession) SetRoundSubmissions(_targetRoundSubmissions uint64) (*types.Transaction, error) {
	return _DAEntrance.Contract.SetRoundSubmissions(&_DAEntrance.TransactOpts, _targetRoundSubmissions)
}

// SetRoundSubmissions is a paid mutator transaction binding the contract method 0x88521ec7.
//
// Solidity: function setRoundSubmissions(uint64 _targetRoundSubmissions) returns()
func (_DAEntrance *DAEntranceTransactorSession) SetRoundSubmissions(_targetRoundSubmissions uint64) (*types.Transaction, error) {
	return _DAEntrance.Contract.SetRoundSubmissions(&_DAEntrance.TransactOpts, _targetRoundSubmissions)
}

// SetSamplePeriod is a paid mutator transaction binding the contract method 0x1192de9a.
//
// Solidity: function setSamplePeriod(uint64 samplePeriod_) returns()
func (_DAEntrance *DAEntranceTransactor) SetSamplePeriod(opts *bind.TransactOpts, samplePeriod_ uint64) (*types.Transaction, error) {
	return _DAEntrance.contract.Transact(opts, "setSamplePeriod", samplePeriod_)
}

// SetSamplePeriod is a paid mutator transaction binding the contract method 0x1192de9a.
//
// Solidity: function setSamplePeriod(uint64 samplePeriod_) returns()
func (_DAEntrance *DAEntranceSession) SetSamplePeriod(samplePeriod_ uint64) (*types.Transaction, error) {
	return _DAEntrance.Contract.SetSamplePeriod(&_DAEntrance.TransactOpts, samplePeriod_)
}

// SetSamplePeriod is a paid mutator transaction binding the contract method 0x1192de9a.
//
// Solidity: function setSamplePeriod(uint64 samplePeriod_) returns()
func (_DAEntrance *DAEntranceTransactorSession) SetSamplePeriod(samplePeriod_ uint64) (*types.Transaction, error) {
	return _DAEntrance.Contract.SetSamplePeriod(&_DAEntrance.TransactOpts, samplePeriod_)
}

// SetServiceFeeRate is a paid mutator transaction binding the contract method 0x9b1d3091.
//
// Solidity: function setServiceFeeRate(uint256 bps) returns()
func (_DAEntrance *DAEntranceTransactor) SetServiceFeeRate(opts *bind.TransactOpts, bps *big.Int) (*types.Transaction, error) {
	return _DAEntrance.contract.Transact(opts, "setServiceFeeRate", bps)
}

// SetServiceFeeRate is a paid mutator transaction binding the contract method 0x9b1d3091.
//
// Solidity: function setServiceFeeRate(uint256 bps) returns()
func (_DAEntrance *DAEntranceSession) SetServiceFeeRate(bps *big.Int) (*types.Transaction, error) {
	return _DAEntrance.Contract.SetServiceFeeRate(&_DAEntrance.TransactOpts, bps)
}

// SetServiceFeeRate is a paid mutator transaction binding the contract method 0x9b1d3091.
//
// Solidity: function setServiceFeeRate(uint256 bps) returns()
func (_DAEntrance *DAEntranceTransactorSession) SetServiceFeeRate(bps *big.Int) (*types.Transaction, error) {
	return _DAEntrance.Contract.SetServiceFeeRate(&_DAEntrance.TransactOpts, bps)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address treasury_) returns()
func (_DAEntrance *DAEntranceTransactor) SetTreasury(opts *bind.TransactOpts, treasury_ common.Address) (*types.Transaction, error) {
	return _DAEntrance.contract.Transact(opts, "setTreasury", treasury_)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address treasury_) returns()
func (_DAEntrance *DAEntranceSession) SetTreasury(treasury_ common.Address) (*types.Transaction, error) {
	return _DAEntrance.Contract.SetTreasury(&_DAEntrance.TransactOpts, treasury_)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address treasury_) returns()
func (_DAEntrance *DAEntranceTransactorSession) SetTreasury(treasury_ common.Address) (*types.Transaction, error) {
	return _DAEntrance.Contract.SetTreasury(&_DAEntrance.TransactOpts, treasury_)
}

// SubmitOriginalData is a paid mutator transaction binding the contract method 0xd4ae59c9.
//
// Solidity: function submitOriginalData(bytes32[] _dataRoots) payable returns()
func (_DAEntrance *DAEntranceTransactor) SubmitOriginalData(opts *bind.TransactOpts, _dataRoots [][32]byte) (*types.Transaction, error) {
	return _DAEntrance.contract.Transact(opts, "submitOriginalData", _dataRoots)
}

// SubmitOriginalData is a paid mutator transaction binding the contract method 0xd4ae59c9.
//
// Solidity: function submitOriginalData(bytes32[] _dataRoots) payable returns()
func (_DAEntrance *DAEntranceSession) SubmitOriginalData(_dataRoots [][32]byte) (*types.Transaction, error) {
	return _DAEntrance.Contract.SubmitOriginalData(&_DAEntrance.TransactOpts, _dataRoots)
}

// SubmitOriginalData is a paid mutator transaction binding the contract method 0xd4ae59c9.
//
// Solidity: function submitOriginalData(bytes32[] _dataRoots) payable returns()
func (_DAEntrance *DAEntranceTransactorSession) SubmitOriginalData(_dataRoots [][32]byte) (*types.Transaction, error) {
	return _DAEntrance.Contract.SubmitOriginalData(&_DAEntrance.TransactOpts, _dataRoots)
}

// SubmitSamplingResponse is a paid mutator transaction binding the contract method 0xf6902775.
//
// Solidity: function submitSamplingResponse((bytes32,uint64,uint64,uint32,uint32,uint256,bytes32,bytes32[3],bytes32[],bytes) rep) returns()
func (_DAEntrance *DAEntranceTransactor) SubmitSamplingResponse(opts *bind.TransactOpts, rep SampleResponse) (*types.Transaction, error) {
	return _DAEntrance.contract.Transact(opts, "submitSamplingResponse", rep)
}

// SubmitSamplingResponse is a paid mutator transaction binding the contract method 0xf6902775.
//
// Solidity: function submitSamplingResponse((bytes32,uint64,uint64,uint32,uint32,uint256,bytes32,bytes32[3],bytes32[],bytes) rep) returns()
func (_DAEntrance *DAEntranceSession) SubmitSamplingResponse(rep SampleResponse) (*types.Transaction, error) {
	return _DAEntrance.Contract.SubmitSamplingResponse(&_DAEntrance.TransactOpts, rep)
}

// SubmitSamplingResponse is a paid mutator transaction binding the contract method 0xf6902775.
//
// Solidity: function submitSamplingResponse((bytes32,uint64,uint64,uint32,uint32,uint256,bytes32,bytes32[3],bytes32[],bytes) rep) returns()
func (_DAEntrance *DAEntranceTransactorSession) SubmitSamplingResponse(rep SampleResponse) (*types.Transaction, error) {
	return _DAEntrance.Contract.SubmitSamplingResponse(&_DAEntrance.TransactOpts, rep)
}

// SubmitVerifiedCommitRoots is a paid mutator transaction binding the contract method 0xeafed6ce.
//
// Solidity: function submitVerifiedCommitRoots((bytes32,uint256,uint256,(uint256,uint256),bytes,(uint256[2],uint256[2]),(uint256,uint256))[] _submissions) returns()
func (_DAEntrance *DAEntranceTransactor) SubmitVerifiedCommitRoots(opts *bind.TransactOpts, _submissions []IDAEntranceCommitRootSubmission) (*types.Transaction, error) {
	return _DAEntrance.contract.Transact(opts, "submitVerifiedCommitRoots", _submissions)
}

// SubmitVerifiedCommitRoots is a paid mutator transaction binding the contract method 0xeafed6ce.
//
// Solidity: function submitVerifiedCommitRoots((bytes32,uint256,uint256,(uint256,uint256),bytes,(uint256[2],uint256[2]),(uint256,uint256))[] _submissions) returns()
func (_DAEntrance *DAEntranceSession) SubmitVerifiedCommitRoots(_submissions []IDAEntranceCommitRootSubmission) (*types.Transaction, error) {
	return _DAEntrance.Contract.SubmitVerifiedCommitRoots(&_DAEntrance.TransactOpts, _submissions)
}

// SubmitVerifiedCommitRoots is a paid mutator transaction binding the contract method 0xeafed6ce.
//
// Solidity: function submitVerifiedCommitRoots((bytes32,uint256,uint256,(uint256,uint256),bytes,(uint256[2],uint256[2]),(uint256,uint256))[] _submissions) returns()
func (_DAEntrance *DAEntranceTransactorSession) SubmitVerifiedCommitRoots(_submissions []IDAEntranceCommitRootSubmission) (*types.Transaction, error) {
	return _DAEntrance.Contract.SubmitVerifiedCommitRoots(&_DAEntrance.TransactOpts, _submissions)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_DAEntrance *DAEntranceTransactor) Sync(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DAEntrance.contract.Transact(opts, "sync")
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_DAEntrance *DAEntranceSession) Sync() (*types.Transaction, error) {
	return _DAEntrance.Contract.Sync(&_DAEntrance.TransactOpts)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_DAEntrance *DAEntranceTransactorSession) Sync() (*types.Transaction, error) {
	return _DAEntrance.Contract.Sync(&_DAEntrance.TransactOpts)
}

// WithdrawPayments is a paid mutator transaction binding the contract method 0x31b3eb94.
//
// Solidity: function withdrawPayments(address payee) returns()
func (_DAEntrance *DAEntranceTransactor) WithdrawPayments(opts *bind.TransactOpts, payee common.Address) (*types.Transaction, error) {
	return _DAEntrance.contract.Transact(opts, "withdrawPayments", payee)
}

// WithdrawPayments is a paid mutator transaction binding the contract method 0x31b3eb94.
//
// Solidity: function withdrawPayments(address payee) returns()
func (_DAEntrance *DAEntranceSession) WithdrawPayments(payee common.Address) (*types.Transaction, error) {
	return _DAEntrance.Contract.WithdrawPayments(&_DAEntrance.TransactOpts, payee)
}

// WithdrawPayments is a paid mutator transaction binding the contract method 0x31b3eb94.
//
// Solidity: function withdrawPayments(address payee) returns()
func (_DAEntrance *DAEntranceTransactorSession) WithdrawPayments(payee common.Address) (*types.Transaction, error) {
	return _DAEntrance.Contract.WithdrawPayments(&_DAEntrance.TransactOpts, payee)
}

// DAEntranceDARewardIterator is returned from FilterDAReward and is used to iterate over the raw logs and unpacked data for DAReward events raised by the DAEntrance contract.
type DAEntranceDARewardIterator struct {
	Event *DAEntranceDAReward // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DAEntranceDARewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAEntranceDAReward)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DAEntranceDAReward)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DAEntranceDARewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAEntranceDARewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAEntranceDAReward represents a DAReward event raised by the DAEntrance contract.
type DAEntranceDAReward struct {
	Beneficiary  common.Address
	SampleRound  *big.Int
	Epoch        *big.Int
	QuorumId     *big.Int
	DataRoot     [32]byte
	Quality      *big.Int
	LineIndex    *big.Int
	SublineIndex *big.Int
	Reward       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDAReward is a free log retrieval operation binding the contract event 0xc3898eb7106c1cb2f727da316a76320c0035f5692950aa7f6b65d20a5efaedc5.
//
// Solidity: event DAReward(address indexed beneficiary, uint256 indexed sampleRound, uint256 indexed epoch, uint256 quorumId, bytes32 dataRoot, uint256 quality, uint256 lineIndex, uint256 sublineIndex, uint256 reward)
func (_DAEntrance *DAEntranceFilterer) FilterDAReward(opts *bind.FilterOpts, beneficiary []common.Address, sampleRound []*big.Int, epoch []*big.Int) (*DAEntranceDARewardIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var sampleRoundRule []interface{}
	for _, sampleRoundItem := range sampleRound {
		sampleRoundRule = append(sampleRoundRule, sampleRoundItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _DAEntrance.contract.FilterLogs(opts, "DAReward", beneficiaryRule, sampleRoundRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &DAEntranceDARewardIterator{contract: _DAEntrance.contract, event: "DAReward", logs: logs, sub: sub}, nil
}

// WatchDAReward is a free log subscription operation binding the contract event 0xc3898eb7106c1cb2f727da316a76320c0035f5692950aa7f6b65d20a5efaedc5.
//
// Solidity: event DAReward(address indexed beneficiary, uint256 indexed sampleRound, uint256 indexed epoch, uint256 quorumId, bytes32 dataRoot, uint256 quality, uint256 lineIndex, uint256 sublineIndex, uint256 reward)
func (_DAEntrance *DAEntranceFilterer) WatchDAReward(opts *bind.WatchOpts, sink chan<- *DAEntranceDAReward, beneficiary []common.Address, sampleRound []*big.Int, epoch []*big.Int) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var sampleRoundRule []interface{}
	for _, sampleRoundItem := range sampleRound {
		sampleRoundRule = append(sampleRoundRule, sampleRoundItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _DAEntrance.contract.WatchLogs(opts, "DAReward", beneficiaryRule, sampleRoundRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAEntranceDAReward)
				if err := _DAEntrance.contract.UnpackLog(event, "DAReward", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDAReward is a log parse operation binding the contract event 0xc3898eb7106c1cb2f727da316a76320c0035f5692950aa7f6b65d20a5efaedc5.
//
// Solidity: event DAReward(address indexed beneficiary, uint256 indexed sampleRound, uint256 indexed epoch, uint256 quorumId, bytes32 dataRoot, uint256 quality, uint256 lineIndex, uint256 sublineIndex, uint256 reward)
func (_DAEntrance *DAEntranceFilterer) ParseDAReward(log types.Log) (*DAEntranceDAReward, error) {
	event := new(DAEntranceDAReward)
	if err := _DAEntrance.contract.UnpackLog(event, "DAReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAEntranceDataUploadIterator is returned from FilterDataUpload and is used to iterate over the raw logs and unpacked data for DataUpload events raised by the DAEntrance contract.
type DAEntranceDataUploadIterator struct {
	Event *DAEntranceDataUpload // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DAEntranceDataUploadIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAEntranceDataUpload)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DAEntranceDataUpload)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DAEntranceDataUploadIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAEntranceDataUploadIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAEntranceDataUpload represents a DataUpload event raised by the DAEntrance contract.
type DAEntranceDataUpload struct {
	DataRoot  [32]byte
	Epoch     *big.Int
	QuorumId  *big.Int
	BlobPrice *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDataUpload is a free log retrieval operation binding the contract event 0xb4e0ecfec4293e970525d9286428425fbdc041540ca6e58ad11bce23d16ed41c.
//
// Solidity: event DataUpload(bytes32 dataRoot, uint256 epoch, uint256 quorumId, uint256 blobPrice)
func (_DAEntrance *DAEntranceFilterer) FilterDataUpload(opts *bind.FilterOpts) (*DAEntranceDataUploadIterator, error) {

	logs, sub, err := _DAEntrance.contract.FilterLogs(opts, "DataUpload")
	if err != nil {
		return nil, err
	}
	return &DAEntranceDataUploadIterator{contract: _DAEntrance.contract, event: "DataUpload", logs: logs, sub: sub}, nil
}

// WatchDataUpload is a free log subscription operation binding the contract event 0xb4e0ecfec4293e970525d9286428425fbdc041540ca6e58ad11bce23d16ed41c.
//
// Solidity: event DataUpload(bytes32 dataRoot, uint256 epoch, uint256 quorumId, uint256 blobPrice)
func (_DAEntrance *DAEntranceFilterer) WatchDataUpload(opts *bind.WatchOpts, sink chan<- *DAEntranceDataUpload) (event.Subscription, error) {

	logs, sub, err := _DAEntrance.contract.WatchLogs(opts, "DataUpload")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAEntranceDataUpload)
				if err := _DAEntrance.contract.UnpackLog(event, "DataUpload", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDataUpload is a log parse operation binding the contract event 0xb4e0ecfec4293e970525d9286428425fbdc041540ca6e58ad11bce23d16ed41c.
//
// Solidity: event DataUpload(bytes32 dataRoot, uint256 epoch, uint256 quorumId, uint256 blobPrice)
func (_DAEntrance *DAEntranceFilterer) ParseDataUpload(log types.Log) (*DAEntranceDataUpload, error) {
	event := new(DAEntranceDataUpload)
	if err := _DAEntrance.contract.UnpackLog(event, "DataUpload", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAEntranceErasureCommitmentVerifiedIterator is returned from FilterErasureCommitmentVerified and is used to iterate over the raw logs and unpacked data for ErasureCommitmentVerified events raised by the DAEntrance contract.
type DAEntranceErasureCommitmentVerifiedIterator struct {
	Event *DAEntranceErasureCommitmentVerified // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DAEntranceErasureCommitmentVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAEntranceErasureCommitmentVerified)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DAEntranceErasureCommitmentVerified)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DAEntranceErasureCommitmentVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAEntranceErasureCommitmentVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAEntranceErasureCommitmentVerified represents a ErasureCommitmentVerified event raised by the DAEntrance contract.
type DAEntranceErasureCommitmentVerified struct {
	DataRoot [32]byte
	Epoch    *big.Int
	QuorumId *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterErasureCommitmentVerified is a free log retrieval operation binding the contract event 0x0f1b20d87bebd11dddaaab51f01cf2726880cb3f8073b636dbafa2aa8cacd256.
//
// Solidity: event ErasureCommitmentVerified(bytes32 dataRoot, uint256 epoch, uint256 quorumId)
func (_DAEntrance *DAEntranceFilterer) FilterErasureCommitmentVerified(opts *bind.FilterOpts) (*DAEntranceErasureCommitmentVerifiedIterator, error) {

	logs, sub, err := _DAEntrance.contract.FilterLogs(opts, "ErasureCommitmentVerified")
	if err != nil {
		return nil, err
	}
	return &DAEntranceErasureCommitmentVerifiedIterator{contract: _DAEntrance.contract, event: "ErasureCommitmentVerified", logs: logs, sub: sub}, nil
}

// WatchErasureCommitmentVerified is a free log subscription operation binding the contract event 0x0f1b20d87bebd11dddaaab51f01cf2726880cb3f8073b636dbafa2aa8cacd256.
//
// Solidity: event ErasureCommitmentVerified(bytes32 dataRoot, uint256 epoch, uint256 quorumId)
func (_DAEntrance *DAEntranceFilterer) WatchErasureCommitmentVerified(opts *bind.WatchOpts, sink chan<- *DAEntranceErasureCommitmentVerified) (event.Subscription, error) {

	logs, sub, err := _DAEntrance.contract.WatchLogs(opts, "ErasureCommitmentVerified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAEntranceErasureCommitmentVerified)
				if err := _DAEntrance.contract.UnpackLog(event, "ErasureCommitmentVerified", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseErasureCommitmentVerified is a log parse operation binding the contract event 0x0f1b20d87bebd11dddaaab51f01cf2726880cb3f8073b636dbafa2aa8cacd256.
//
// Solidity: event ErasureCommitmentVerified(bytes32 dataRoot, uint256 epoch, uint256 quorumId)
func (_DAEntrance *DAEntranceFilterer) ParseErasureCommitmentVerified(log types.Log) (*DAEntranceErasureCommitmentVerified, error) {
	event := new(DAEntranceErasureCommitmentVerified)
	if err := _DAEntrance.contract.UnpackLog(event, "ErasureCommitmentVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAEntranceNewSampleRoundIterator is returned from FilterNewSampleRound and is used to iterate over the raw logs and unpacked data for NewSampleRound events raised by the DAEntrance contract.
type DAEntranceNewSampleRoundIterator struct {
	Event *DAEntranceNewSampleRound // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DAEntranceNewSampleRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAEntranceNewSampleRound)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DAEntranceNewSampleRound)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DAEntranceNewSampleRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAEntranceNewSampleRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAEntranceNewSampleRound represents a NewSampleRound event raised by the DAEntrance contract.
type DAEntranceNewSampleRound struct {
	SampleRound  *big.Int
	SampleHeight *big.Int
	SampleSeed   [32]byte
	PodasTarget  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewSampleRound is a free log retrieval operation binding the contract event 0xdfb5db5886e81f083727f21152a2a83457e99364e9f104e1aa10bbd6d9b4b95f.
//
// Solidity: event NewSampleRound(uint256 indexed sampleRound, uint256 sampleHeight, bytes32 sampleSeed, uint256 podasTarget)
func (_DAEntrance *DAEntranceFilterer) FilterNewSampleRound(opts *bind.FilterOpts, sampleRound []*big.Int) (*DAEntranceNewSampleRoundIterator, error) {

	var sampleRoundRule []interface{}
	for _, sampleRoundItem := range sampleRound {
		sampleRoundRule = append(sampleRoundRule, sampleRoundItem)
	}

	logs, sub, err := _DAEntrance.contract.FilterLogs(opts, "NewSampleRound", sampleRoundRule)
	if err != nil {
		return nil, err
	}
	return &DAEntranceNewSampleRoundIterator{contract: _DAEntrance.contract, event: "NewSampleRound", logs: logs, sub: sub}, nil
}

// WatchNewSampleRound is a free log subscription operation binding the contract event 0xdfb5db5886e81f083727f21152a2a83457e99364e9f104e1aa10bbd6d9b4b95f.
//
// Solidity: event NewSampleRound(uint256 indexed sampleRound, uint256 sampleHeight, bytes32 sampleSeed, uint256 podasTarget)
func (_DAEntrance *DAEntranceFilterer) WatchNewSampleRound(opts *bind.WatchOpts, sink chan<- *DAEntranceNewSampleRound, sampleRound []*big.Int) (event.Subscription, error) {

	var sampleRoundRule []interface{}
	for _, sampleRoundItem := range sampleRound {
		sampleRoundRule = append(sampleRoundRule, sampleRoundItem)
	}

	logs, sub, err := _DAEntrance.contract.WatchLogs(opts, "NewSampleRound", sampleRoundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAEntranceNewSampleRound)
				if err := _DAEntrance.contract.UnpackLog(event, "NewSampleRound", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewSampleRound is a log parse operation binding the contract event 0xdfb5db5886e81f083727f21152a2a83457e99364e9f104e1aa10bbd6d9b4b95f.
//
// Solidity: event NewSampleRound(uint256 indexed sampleRound, uint256 sampleHeight, bytes32 sampleSeed, uint256 podasTarget)
func (_DAEntrance *DAEntranceFilterer) ParseNewSampleRound(log types.Log) (*DAEntranceNewSampleRound, error) {
	event := new(DAEntranceNewSampleRound)
	if err := _DAEntrance.contract.UnpackLog(event, "NewSampleRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAEntranceRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the DAEntrance contract.
type DAEntranceRoleAdminChangedIterator struct {
	Event *DAEntranceRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DAEntranceRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAEntranceRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DAEntranceRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DAEntranceRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAEntranceRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAEntranceRoleAdminChanged represents a RoleAdminChanged event raised by the DAEntrance contract.
type DAEntranceRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_DAEntrance *DAEntranceFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*DAEntranceRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _DAEntrance.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &DAEntranceRoleAdminChangedIterator{contract: _DAEntrance.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_DAEntrance *DAEntranceFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *DAEntranceRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _DAEntrance.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAEntranceRoleAdminChanged)
				if err := _DAEntrance.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_DAEntrance *DAEntranceFilterer) ParseRoleAdminChanged(log types.Log) (*DAEntranceRoleAdminChanged, error) {
	event := new(DAEntranceRoleAdminChanged)
	if err := _DAEntrance.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAEntranceRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the DAEntrance contract.
type DAEntranceRoleGrantedIterator struct {
	Event *DAEntranceRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DAEntranceRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAEntranceRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DAEntranceRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DAEntranceRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAEntranceRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAEntranceRoleGranted represents a RoleGranted event raised by the DAEntrance contract.
type DAEntranceRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_DAEntrance *DAEntranceFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DAEntranceRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DAEntrance.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DAEntranceRoleGrantedIterator{contract: _DAEntrance.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_DAEntrance *DAEntranceFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *DAEntranceRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DAEntrance.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAEntranceRoleGranted)
				if err := _DAEntrance.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_DAEntrance *DAEntranceFilterer) ParseRoleGranted(log types.Log) (*DAEntranceRoleGranted, error) {
	event := new(DAEntranceRoleGranted)
	if err := _DAEntrance.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAEntranceRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the DAEntrance contract.
type DAEntranceRoleRevokedIterator struct {
	Event *DAEntranceRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DAEntranceRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAEntranceRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DAEntranceRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DAEntranceRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAEntranceRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAEntranceRoleRevoked represents a RoleRevoked event raised by the DAEntrance contract.
type DAEntranceRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_DAEntrance *DAEntranceFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DAEntranceRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DAEntrance.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DAEntranceRoleRevokedIterator{contract: _DAEntrance.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_DAEntrance *DAEntranceFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *DAEntranceRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DAEntrance.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAEntranceRoleRevoked)
				if err := _DAEntrance.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_DAEntrance *DAEntranceFilterer) ParseRoleRevoked(log types.Log) (*DAEntranceRoleRevoked, error) {
	event := new(DAEntranceRoleRevoked)
	if err := _DAEntrance.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
