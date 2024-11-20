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

// ChunkLinearRewardMetaData contains all meta data concerning the ChunkLinearReward contract.
var ChunkLinearRewardMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"releaseSeconds_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pricingIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DistributeReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PARAMS_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pricingIndex\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"claimMineReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"donate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"beforeLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chargedSectors\",\"type\":\"uint256\"}],\"name\":\"fillReward\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstRewardableChunk\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"market_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mine_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"market\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mine\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"}],\"name\":\"payments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"releaseSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pricingIndex\",\"type\":\"uint256\"}],\"name\":\"rewardDeadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"lockedReward\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"claimableReward\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"distributedReward\",\"type\":\"uint128\"},{\"internalType\":\"uint40\",\"name\":\"startTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"lastUpdate\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serviceFeeRateBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"baseReward_\",\"type\":\"uint256\"}],\"name\":\"setBaseReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bps\",\"type\":\"uint256\"}],\"name\":\"setServiceFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"treasury_\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBaseReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"payee\",\"type\":\"address\"}],\"name\":\"withdrawPayments\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260043610620001f35760003560e01c80639010d07c116200010b578063b7a3c04c11620000a1578063e2982c21116200006c578063e2982c2114620005d9578063ed88c68e14620005fe578063f0f442601462000608578063f301af42146200062d57600080fd5b8063b7a3c04c1462000552578063c05751111462000577578063ca15c873146200058f578063d547741f14620005b457600080fd5b80639b1d309111620000e25780639b1d309114620004bf578063a217fddf14620004e4578063b15d20da14620004fb578063b3b30c1a146200052057600080fd5b80639010d07c146200045357806391d14854146200047857806399f4b251146200049d57600080fd5b806331b3eb94116200018d57806361d027b3116200015857806361d027b314620003c657806376ad03bc14620004015780637f1b5e43146200041957806380f55605146200043157600080fd5b806331b3eb94146200034057806336568abe1462000365578063485cc955146200038a57806359e9670014620003af57600080fd5b8063158ef93e11620001ce578063158ef93e146200028d5780632129593114620002b0578063248a9ca314620002e65780632f2ff15d146200031b57600080fd5b806301ffc9a714620001f85780630373a23a14620002325780630a539a191462000259575b600080fd5b3480156200020557600080fd5b506200021d6200021736600462001d6c565b620006d9565b60405190151581526020015b60405180910390f35b3480156200023f57600080fd5b50620002576200025136600462001d98565b62000707565b005b3480156200026657600080fd5b506200027e6200027836600462001d98565b62000728565b60405190815260200162000229565b3480156200029a57600080fd5b506000546200021d90600160a01b900460ff1681565b348015620002bd57600080fd5b506200027e7f000000000000000000000000000000000000000000000000000000000000000081565b348015620002f357600080fd5b506200027e6200030536600462001d98565b6000908152600160208190526040909120015490565b3480156200032857600080fd5b50620002576200033a36600462001dc8565b620007e3565b3480156200034d57600080fd5b50620002576200035f36600462001dfb565b62000812565b3480156200037257600080fd5b50620002576200038436600462001dc8565b62000876565b3480156200039757600080fd5b5062000257620003a936600462001e1b565b620008fc565b62000257620003c036600462001e4e565b62000a25565b348015620003d357600080fd5b50600954620003e8906001600160a01b031681565b6040516001600160a01b03909116815260200162000229565b3480156200040e57600080fd5b506200027e60075481565b3480156200042657600080fd5b506200027e60065481565b3480156200043e57600080fd5b50600354620003e8906001600160a01b031681565b3480156200046057600080fd5b50620003e86200047236600462001e4e565b62000e96565b3480156200048557600080fd5b506200021d6200049736600462001dc8565b62000eb0565b348015620004aa57600080fd5b50600454620003e8906001600160a01b031681565b348015620004cc57600080fd5b5062000257620004de36600462001d98565b62000edb565b348015620004f157600080fd5b506200027e600081565b3480156200050857600080fd5b506200027e6000805160206200274f83398151915281565b3480156200052d57600080fd5b506200053862000efc565b60405167ffffffffffffffff909116815260200162000229565b3480156200055f57600080fd5b50620002576200057136600462001e71565b62000fbe565b3480156200058457600080fd5b506200027e60085481565b3480156200059c57600080fd5b506200027e620005ae36600462001d98565b620011ec565b348015620005c157600080fd5b5062000257620005d336600462001dc8565b62001205565b348015620005e657600080fd5b506200027e620005f836600462001dfb565b6200122f565b62000257620012a1565b3480156200061557600080fd5b50620002576200062736600462001dfb565b620012bc565b3480156200063a57600080fd5b50620006966200064c36600462001d98565b600560205260009081526040902080546001909101546001600160801b0380831692600160801b908190048216929182169164ffffffffff918104821691600160a81b9091041685565b604080516001600160801b0396871681529486166020860152929094169183019190915264ffffffffff9081166060830152909116608082015260a00162000229565b60006001600160e01b03198216635a05180f60e01b14806200070157506200070182620012fa565b92915050565b6000805160206200274f833981519152620007228162001331565b50600755565b6000818152600560209081526040808320815160a08101835281546001600160801b038082168352600160801b918290048116958301959095526001909201549384169281019290925264ffffffffff908304811660608301819052600160a81b909304166080820152908203620007a35750600092915050565b7f0000000000000000000000000000000000000000000000000000000000000000816060015164ffffffffff16620007dc919062001ec2565b9392505050565b60008281526001602081905260409091200154620008018162001331565b6200080d838362001340565b505050565b6000546040516351cff8d960e01b81526001600160a01b038381166004830152909116906351cff8d990602401600060405180830381600087803b1580156200085a57600080fd5b505af11580156200086f573d6000803e3d6000fd5b5050505050565b6001600160a01b0381163314620008ec5760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b60648201526084015b60405180910390fd5b620008f8828262001366565b5050565b600054600160a01b900460ff1615620009645760405162461bcd60e51b8152602060048201526024808201527f5a67496e697469616c697a61626c653a20616c726561647920696e697469616c6044820152631a5e995960e21b6064820152608401620008e3565b6000805460ff60a01b1916600160a01b1781556200098a90620009843390565b62001340565b620009a56000805160206200274f8339815191523362001340565b600380546001600160a01b038085166001600160a01b0319928316179092556004805492841692909116919091179055604051620009e39062001d5e565b604051809103906000f08015801562000a00573d6000803e3d6000fd5b50600080546001600160a01b0319166001600160a01b03929092169190911790555050565b6003546001600160a01b0316336001600160a01b03161462000a8a5760405162461bcd60e51b815260206004820152601f60248201527f53656e64657220646f6573206e6f742068617665207065726d697373696f6e006044820152606401620008e3565b60006127106008543462000a9f919062001ed8565b62000aab919062001f10565b9050801562000acc5760095462000acc906001600160a01b0316826200138c565b600062000ada823462001f27565b90508260008161010062000af16104008062001ed8565b62000aff9061040062001ed8565b62000b0c90600862001ed8565b62000b18919062001f10565b62000b24908562001ed8565b62000b30919062001f10565b9050600062000b40838862001ec2565b9050600061010062000b556104008062001ed8565b62000b639061040062001ed8565b62000b7090600862001ed8565b62000b7c919062001f10565b62000b88908962001f3d565b61010062000b996104008062001ed8565b62000ba79061040062001ed8565b62000bb490600862001ed8565b62000bc0919062001f10565b62000bcc919062001f27565b90506000600161010062000be36104008062001ed8565b62000bf19061040062001ed8565b62000bfe90600862001ed8565b62000c0a919062001f10565b62000c16848c62001ec2565b62000c22919062001f10565b62000c2e919062001f27565b9050600061010062000c436104008062001ed8565b62000c519061040062001ed8565b62000c5e90600862001ed8565b62000c6a919062001f10565b62000c7760018662001f27565b62000c83919062001f3d565b62000c9090600162001ec2565b9050600061010062000ca56104008062001ed8565b62000cb39061040062001ed8565b62000cc090600862001ed8565b62000ccc919062001f10565b62000cd8838762001f27565b62000ce4919062001f10565b9050600061010062000cf96104008062001ed8565b62000d079061040062001ed8565b62000d1490600862001ed8565b62000d20919062001f10565b62000d2d83600162001ec2565b62000d39919062001ed8565b8614905081840362000d6657600084815260056020526040902062000d60908a83620014ab565b62000e88565b62000dd161010062000d7b6104008062001ed8565b62000d899061040062001ed8565b62000d9690600862001ed8565b62000da2919062001f10565b62000dae878a62001ed8565b62000dba919062001f10565b6000868152600560205260409020906001620014ab565b600062000de085600162001ec2565b90505b8281101562000e1d57600081815260056020526040902062000e0890896001620014ab565b8062000e148162001f54565b91505062000de3565b5062000e8861010062000e336104008062001ed8565b62000e419061040062001ed8565b62000e4e90600862001ed8565b62000e5a919062001f10565b62000e66858a62001ed8565b62000e72919062001f10565b60008481526005602052604090209083620014ab565b505050505050505050505050565b6000828152600260205260408120620007dc9083620015e6565b60009182526001602090815260408084206001600160a01b0393909316845291905290205460ff1690565b6000805160206200274f83398151915262000ef68162001331565b50600855565b6000806104005b62000f188167ffffffffffffffff16620015f4565b1562000f365790508062000f2e60028262001f70565b905062000f03565b8067ffffffffffffffff168267ffffffffffffffff16101562000fb8576000600262000f63848462001fa3565b62000f6f919062001fc7565b62000f7b908462001ff1565b905062000f928167ffffffffffffffff16620015f4565b1562000fad5762000fa581600162001ff1565b925062000fb1565b8091505b5062000f36565b50919050565b6004546001600160a01b0316336001600160a01b031614620010235760405162461bcd60e51b815260206004820152601f60248201527f53656e64657220646f6573206e6f742068617665207065726d697373696f6e006044820152606401620008e3565b6000838152600560209081526040808320815160a08101835281546001600160801b038082168352600160801b918290048116958301959095526001909201549384169281019290925264ffffffffff90830481166060830152600160a81b909204909116608082015290620010998262001617565b9050620010a7828262001645565b6000620010b48362001698565b60008781526005602090815260408083208751928801516001600160801b03938416600160801b9185168202178255918801516001909101805460608a015160808b0151939095166001600160a81b03199091161764ffffffffff9485169093029290921764ffffffffff60a81b1916600160a81b939091169290920291909117905590915062001147878584620016ff565b9050600081600654116200115e5760065462001160565b815b90506200116e818462001ec2565b9250806006600082825462001184919062001f27565b90915550508215620011e2576200119c878462001756565b866001600160a01b0316887f83617a1b0f847971f005bd162dde513cfe93df96e6293c3bbb5fe9c40629dd4c85604051620011d991815260200190565b60405180910390a35b5050505050505050565b60008181526002602052604081206200070190620017bd565b60008281526001602081905260409091200154620012238162001331565b6200080d838362001366565b600080546040516371d4ed8d60e11b81526001600160a01b0384811660048301529091169063e3a9db1a90602401602060405180830381865afa1580156200127b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000701919062002015565b3460066000828254620012b5919062001ec2565b9091555050565b6000805160206200274f833981519152620012d78162001331565b50600980546001600160a01b0319166001600160a01b0392909216919091179055565b60006001600160e01b03198216637965db0b60e01b14806200070157506301ffc9a760e01b6001600160e01b031983161462000701565b6200133d8133620017c8565b50565b6200134c82826200182c565b60008281526002602052604090206200080d90826200189a565b620013728282620018b1565b60008281526002602052604090206200080d90826200191b565b80471015620013de5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e63650000006044820152606401620008e3565b6000826001600160a01b03168260405160006040518083038185875af1925050503d80600081146200142d576040519150601f19603f3d011682016040523d82523d6000602084013e62001432565b606091505b50509050806200080d5760405162461bcd60e51b815260206004820152603a60248201527f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260448201527f6563697069656e74206d617920686176652072657665727465640000000000006064820152608401620008e3565b6001600160801b03821115620014f65760405162461bcd60e51b815260206004820152600f60248201526e526577617264206f766572666c6f7760881b6044820152606401620008e3565b6001830154600160801b900464ffffffffff1615620015585760405162461bcd60e51b815260206004820181905260248201527f526577617264206974656d20686173206265656e20696e697469616c697a65646044820152606401620008e3565b825482908490600090620015779084906001600160801b03166200202f565b92506101000a8154816001600160801b0302191690836001600160801b0316021790555080156200080d575050600101805469ffffffffffffffffffff60801b1916600160801b4264ffffffffff1690810264ffffffffff60a81b191691909117600160a81b91909102179055565b6000620007dc838362001932565b600080620016028362000728565b90508015801590620007dc5750421192915050565b600062000701827f00000000000000000000000000000000000000000000000000000000000000006200195f565b808260000181815162001659919062002052565b6001600160801b03169052506020820180518291906200167b9083906200202f565b6001600160801b03169052505064ffffffffff4216608090910152565b60008060028360200151620016ae919062002075565b90508083602001818151620016c4919062002052565b6001600160801b0316905250604083018051829190620016e69083906200202f565b6001600160801b03908116909152919091169392505050565b6000427f0000000000000000000000000000000000000000000000000000000000000000846060015164ffffffffff166200173b919062001ec2565b11156200174c5750600754620007dc565b5060009392505050565b60005460405163f340fa0160e01b81526001600160a01b0384811660048301529091169063f340fa019083906024016000604051808303818588803b1580156200179f57600080fd5b505af1158015620017b4573d6000803e3d6000fd5b50505050505050565b600062000701825490565b620017d4828262000eb0565b620008f857620017e48162001a34565b620017f183602062001a47565b60405160200162001804929190620020b8565b60408051601f198184030181529082905262461bcd60e51b8252620008e39160040162002131565b62001838828262000eb0565b620008f85760008281526001602081815260408084206001600160a01b0386168086529252808420805460ff19169093179092559051339285917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45050565b6000620007dc836001600160a01b03841662001c01565b620018bd828262000eb0565b15620008f85760008281526001602090815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b6000620007dc836001600160a01b03841662001c53565b60008260000182815481106200194c576200194c62002166565b9060005260206000200154905092915050565b6000826080015164ffffffffff166000036200197e5750600062000701565b6000836040015184602001516200199691906200202f565b6001600160801b0316905060008185600001516001600160801b0316620019be919062001ec2565b90506000856060015164ffffffffff1642620019db919062001f27565b9050600085620019ec838562001ed8565b620019f8919062001f10565b90508281111562001a065750815b8381101562001a1d57600094505050505062000701565b62001a29848262001f27565b979650505050505050565b6060620007016001600160a01b03831660145b6060600062001a5883600262001ed8565b62001a6590600262001ec2565b67ffffffffffffffff81111562001a805762001a806200217c565b6040519080825280601f01601f19166020018201604052801562001aab576020820181803683370190505b509050600360fc1b8160008151811062001ac95762001ac962002166565b60200101906001600160f81b031916908160001a905350600f60fb1b8160018151811062001afb5762001afb62002166565b60200101906001600160f81b031916908160001a905350600062001b2184600262001ed8565b62001b2e90600162001ec2565b90505b600181111562001bb0576f181899199a1a9b1b9c1cb0b131b232b360811b85600f166010811062001b665762001b6662002166565b1a60f81b82828151811062001b7f5762001b7f62002166565b60200101906001600160f81b031916908160001a90535060049490941c9362001ba88162002192565b905062001b31565b508315620007dc5760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152606401620008e3565b600081815260018301602052604081205462001c4a5750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915562000701565b50600062000701565b6000818152600183016020526040812054801562001d4c57600062001c7a60018362001f27565b855490915060009062001c909060019062001f27565b905081811462001cfc57600086600001828154811062001cb45762001cb462002166565b906000526020600020015490508087600001848154811062001cda5762001cda62002166565b6000918252602080832090910192909255918252600188019052604090208390555b855486908062001d105762001d10620021ac565b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505062000701565b600091505062000701565b5092915050565b61058c80620021c383390190565b60006020828403121562001d7f57600080fd5b81356001600160e01b031981168114620007dc57600080fd5b60006020828403121562001dab57600080fd5b5035919050565b6001600160a01b03811681146200133d57600080fd5b6000806040838503121562001ddc57600080fd5b82359150602083013562001df08162001db2565b809150509250929050565b60006020828403121562001e0e57600080fd5b8135620007dc8162001db2565b6000806040838503121562001e2f57600080fd5b823562001e3c8162001db2565b9150602083013562001df08162001db2565b6000806040838503121562001e6257600080fd5b50508035926020909101359150565b60008060006060848603121562001e8757600080fd5b83359250602084013562001e9b8162001db2565b929592945050506040919091013590565b634e487b7160e01b600052601160045260246000fd5b8082018082111562000701576200070162001eac565b600081600019048311821515161562001ef55762001ef562001eac565b500290565b634e487b7160e01b600052601260045260246000fd5b60008262001f225762001f2262001efa565b500490565b8181038181111562000701576200070162001eac565b60008262001f4f5762001f4f62001efa565b500690565b60006001820162001f695762001f6962001eac565b5060010190565b600067ffffffffffffffff8083168185168183048111821515161562001f9a5762001f9a62001eac565b02949350505050565b67ffffffffffffffff82811682821603908082111562001d575762001d5762001eac565b600067ffffffffffffffff8084168062001fe55762001fe562001efa565b92169190910492915050565b67ffffffffffffffff81811683821601908082111562001d575762001d5762001eac565b6000602082840312156200202857600080fd5b5051919050565b6001600160801b0381811683821601908082111562001d575762001d5762001eac565b6001600160801b0382811682821603908082111562001d575762001d5762001eac565b60006001600160801b038084168062001fe55762001fe562001efa565b60005b83811015620020af57818101518382015260200162002095565b50506000910152565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000815260008351620020f281601785016020880162002092565b7001034b99036b4b9b9b4b733903937b6329607d1b60179184019182015283516200212581602884016020880162002092565b01602801949350505050565b60208152600082518060208401526200215281604085016020870162002092565b601f01601f19169190910160400192915050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b600081620021a457620021a462001eac565b506000190190565b634e487b7160e01b600052603160045260246000fdfe608060405234801561001057600080fd5b5061001a3361001f565b61006f565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b61050e8061007e6000396000f3fe6080604052600436106100555760003560e01c806351cff8d91461005a578063715018a61461007c5780638da5cb5b14610091578063e3a9db1a146100be578063f2fde38b14610102578063f340fa0114610122575b600080fd5b34801561006657600080fd5b5061007a61007536600461048d565b610135565b005b34801561008857600080fd5b5061007a6101ac565b34801561009d57600080fd5b506000546040516001600160a01b0390911681526020015b60405180910390f35b3480156100ca57600080fd5b506100f46100d936600461048d565b6001600160a01b031660009081526001602052604090205490565b6040519081526020016100b5565b34801561010e57600080fd5b5061007a61011d36600461048d565b6101c0565b61007a61013036600461048d565b61023e565b61013d6102b0565b6001600160a01b0381166000818152600160205260408120805491905590610165908261030a565b816001600160a01b03167f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5826040516101a091815260200190565b60405180910390a25050565b6101b46102b0565b6101be6000610428565b565b6101c86102b0565b6001600160a01b0381166102325760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b61023b81610428565b50565b6102466102b0565b6001600160a01b0381166000908152600160205260408120805434928392916102709084906104b1565b90915550506040518181526001600160a01b038316907f2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4906020016101a0565b6000546001600160a01b031633146101be5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610229565b8047101561035a5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e63650000006044820152606401610229565b6000826001600160a01b03168260405160006040518083038185875af1925050503d80600081146103a7576040519150601f19603f3d011682016040523d82523d6000602084013e6103ac565b606091505b50509050806104235760405162461bcd60e51b815260206004820152603a60248201527f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260448201527f6563697069656e74206d617920686176652072657665727465640000000000006064820152608401610229565b505050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b038116811461023b57600080fd5b60006020828403121561049f57600080fd5b81356104aa81610478565b9392505050565b808201808211156104d257634e487b7160e01b600052601160045260246000fd5b9291505056fea2646970667358221220e52cf2b51f0a8d14b10d6c445edb47884f6fbd6d5fbaed7572ef106f818cf48064736f6c63430008100033b9d69e0ca90be54a40811e436234a7f7908b87ff2bec27e64f878b166da8e8e5a2646970667358221220c0f986410c492e567a1700351ec1ad79c1cca1a0b4779af3090923ab08b8d33664736f6c63430008100033",
}

// ChunkLinearRewardABI is the input ABI used to generate the binding from.
// Deprecated: Use ChunkLinearRewardMetaData.ABI instead.
var ChunkLinearRewardABI = ChunkLinearRewardMetaData.ABI

// ChunkLinearRewardBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ChunkLinearRewardMetaData.Bin instead.
var ChunkLinearRewardBin = ChunkLinearRewardMetaData.Bin

// DeployChunkLinearReward deploys a new Ethereum contract, binding an instance of ChunkLinearReward to it.
func DeployChunkLinearReward(auth *bind.TransactOpts, backend bind.ContractBackend, releaseSeconds_ *big.Int) (common.Address, *types.Transaction, *ChunkLinearReward, error) {
	parsed, err := ChunkLinearRewardMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ChunkLinearRewardBin), backend, releaseSeconds_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChunkLinearReward{ChunkLinearRewardCaller: ChunkLinearRewardCaller{contract: contract}, ChunkLinearRewardTransactor: ChunkLinearRewardTransactor{contract: contract}, ChunkLinearRewardFilterer: ChunkLinearRewardFilterer{contract: contract}}, nil
}

// ChunkLinearReward is an auto generated Go binding around an Ethereum contract.
type ChunkLinearReward struct {
	ChunkLinearRewardCaller     // Read-only binding to the contract
	ChunkLinearRewardTransactor // Write-only binding to the contract
	ChunkLinearRewardFilterer   // Log filterer for contract events
}

// ChunkLinearRewardCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChunkLinearRewardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChunkLinearRewardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChunkLinearRewardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChunkLinearRewardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChunkLinearRewardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChunkLinearRewardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChunkLinearRewardSession struct {
	Contract     *ChunkLinearReward // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ChunkLinearRewardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChunkLinearRewardCallerSession struct {
	Contract *ChunkLinearRewardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ChunkLinearRewardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChunkLinearRewardTransactorSession struct {
	Contract     *ChunkLinearRewardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ChunkLinearRewardRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChunkLinearRewardRaw struct {
	Contract *ChunkLinearReward // Generic contract binding to access the raw methods on
}

// ChunkLinearRewardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChunkLinearRewardCallerRaw struct {
	Contract *ChunkLinearRewardCaller // Generic read-only contract binding to access the raw methods on
}

// ChunkLinearRewardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChunkLinearRewardTransactorRaw struct {
	Contract *ChunkLinearRewardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChunkLinearReward creates a new instance of ChunkLinearReward, bound to a specific deployed contract.
func NewChunkLinearReward(address common.Address, backend bind.ContractBackend) (*ChunkLinearReward, error) {
	contract, err := bindChunkLinearReward(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChunkLinearReward{ChunkLinearRewardCaller: ChunkLinearRewardCaller{contract: contract}, ChunkLinearRewardTransactor: ChunkLinearRewardTransactor{contract: contract}, ChunkLinearRewardFilterer: ChunkLinearRewardFilterer{contract: contract}}, nil
}

// NewChunkLinearRewardCaller creates a new read-only instance of ChunkLinearReward, bound to a specific deployed contract.
func NewChunkLinearRewardCaller(address common.Address, caller bind.ContractCaller) (*ChunkLinearRewardCaller, error) {
	contract, err := bindChunkLinearReward(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChunkLinearRewardCaller{contract: contract}, nil
}

// NewChunkLinearRewardTransactor creates a new write-only instance of ChunkLinearReward, bound to a specific deployed contract.
func NewChunkLinearRewardTransactor(address common.Address, transactor bind.ContractTransactor) (*ChunkLinearRewardTransactor, error) {
	contract, err := bindChunkLinearReward(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChunkLinearRewardTransactor{contract: contract}, nil
}

// NewChunkLinearRewardFilterer creates a new log filterer instance of ChunkLinearReward, bound to a specific deployed contract.
func NewChunkLinearRewardFilterer(address common.Address, filterer bind.ContractFilterer) (*ChunkLinearRewardFilterer, error) {
	contract, err := bindChunkLinearReward(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChunkLinearRewardFilterer{contract: contract}, nil
}

// bindChunkLinearReward binds a generic wrapper to an already deployed contract.
func bindChunkLinearReward(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ChunkLinearRewardMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChunkLinearReward *ChunkLinearRewardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChunkLinearReward.Contract.ChunkLinearRewardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChunkLinearReward *ChunkLinearRewardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.ChunkLinearRewardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChunkLinearReward *ChunkLinearRewardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.ChunkLinearRewardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChunkLinearReward *ChunkLinearRewardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChunkLinearReward.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChunkLinearReward *ChunkLinearRewardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChunkLinearReward *ChunkLinearRewardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ChunkLinearReward *ChunkLinearRewardCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ChunkLinearReward *ChunkLinearRewardSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ChunkLinearReward.Contract.DEFAULTADMINROLE(&_ChunkLinearReward.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ChunkLinearReward.Contract.DEFAULTADMINROLE(&_ChunkLinearReward.CallOpts)
}

// PARAMSADMINROLE is a free data retrieval call binding the contract method 0xb15d20da.
//
// Solidity: function PARAMS_ADMIN_ROLE() view returns(bytes32)
func (_ChunkLinearReward *ChunkLinearRewardCaller) PARAMSADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "PARAMS_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PARAMSADMINROLE is a free data retrieval call binding the contract method 0xb15d20da.
//
// Solidity: function PARAMS_ADMIN_ROLE() view returns(bytes32)
func (_ChunkLinearReward *ChunkLinearRewardSession) PARAMSADMINROLE() ([32]byte, error) {
	return _ChunkLinearReward.Contract.PARAMSADMINROLE(&_ChunkLinearReward.CallOpts)
}

// PARAMSADMINROLE is a free data retrieval call binding the contract method 0xb15d20da.
//
// Solidity: function PARAMS_ADMIN_ROLE() view returns(bytes32)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) PARAMSADMINROLE() ([32]byte, error) {
	return _ChunkLinearReward.Contract.PARAMSADMINROLE(&_ChunkLinearReward.CallOpts)
}

// BaseReward is a free data retrieval call binding the contract method 0x76ad03bc.
//
// Solidity: function baseReward() view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardCaller) BaseReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "baseReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseReward is a free data retrieval call binding the contract method 0x76ad03bc.
//
// Solidity: function baseReward() view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardSession) BaseReward() (*big.Int, error) {
	return _ChunkLinearReward.Contract.BaseReward(&_ChunkLinearReward.CallOpts)
}

// BaseReward is a free data retrieval call binding the contract method 0x76ad03bc.
//
// Solidity: function baseReward() view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) BaseReward() (*big.Int, error) {
	return _ChunkLinearReward.Contract.BaseReward(&_ChunkLinearReward.CallOpts)
}

// FirstRewardableChunk is a free data retrieval call binding the contract method 0xb3b30c1a.
//
// Solidity: function firstRewardableChunk() view returns(uint64)
func (_ChunkLinearReward *ChunkLinearRewardCaller) FirstRewardableChunk(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "firstRewardableChunk")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// FirstRewardableChunk is a free data retrieval call binding the contract method 0xb3b30c1a.
//
// Solidity: function firstRewardableChunk() view returns(uint64)
func (_ChunkLinearReward *ChunkLinearRewardSession) FirstRewardableChunk() (uint64, error) {
	return _ChunkLinearReward.Contract.FirstRewardableChunk(&_ChunkLinearReward.CallOpts)
}

// FirstRewardableChunk is a free data retrieval call binding the contract method 0xb3b30c1a.
//
// Solidity: function firstRewardableChunk() view returns(uint64)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) FirstRewardableChunk() (uint64, error) {
	return _ChunkLinearReward.Contract.FirstRewardableChunk(&_ChunkLinearReward.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ChunkLinearReward *ChunkLinearRewardCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ChunkLinearReward *ChunkLinearRewardSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ChunkLinearReward.Contract.GetRoleAdmin(&_ChunkLinearReward.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ChunkLinearReward.Contract.GetRoleAdmin(&_ChunkLinearReward.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ChunkLinearReward *ChunkLinearRewardCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ChunkLinearReward *ChunkLinearRewardSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ChunkLinearReward.Contract.GetRoleMember(&_ChunkLinearReward.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ChunkLinearReward.Contract.GetRoleMember(&_ChunkLinearReward.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ChunkLinearReward.Contract.GetRoleMemberCount(&_ChunkLinearReward.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ChunkLinearReward.Contract.GetRoleMemberCount(&_ChunkLinearReward.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ChunkLinearReward *ChunkLinearRewardCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ChunkLinearReward *ChunkLinearRewardSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ChunkLinearReward.Contract.HasRole(&_ChunkLinearReward.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ChunkLinearReward.Contract.HasRole(&_ChunkLinearReward.CallOpts, role, account)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_ChunkLinearReward *ChunkLinearRewardCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_ChunkLinearReward *ChunkLinearRewardSession) Initialized() (bool, error) {
	return _ChunkLinearReward.Contract.Initialized(&_ChunkLinearReward.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) Initialized() (bool, error) {
	return _ChunkLinearReward.Contract.Initialized(&_ChunkLinearReward.CallOpts)
}

// Market is a free data retrieval call binding the contract method 0x80f55605.
//
// Solidity: function market() view returns(address)
func (_ChunkLinearReward *ChunkLinearRewardCaller) Market(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "market")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Market is a free data retrieval call binding the contract method 0x80f55605.
//
// Solidity: function market() view returns(address)
func (_ChunkLinearReward *ChunkLinearRewardSession) Market() (common.Address, error) {
	return _ChunkLinearReward.Contract.Market(&_ChunkLinearReward.CallOpts)
}

// Market is a free data retrieval call binding the contract method 0x80f55605.
//
// Solidity: function market() view returns(address)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) Market() (common.Address, error) {
	return _ChunkLinearReward.Contract.Market(&_ChunkLinearReward.CallOpts)
}

// Mine is a free data retrieval call binding the contract method 0x99f4b251.
//
// Solidity: function mine() view returns(address)
func (_ChunkLinearReward *ChunkLinearRewardCaller) Mine(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "mine")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mine is a free data retrieval call binding the contract method 0x99f4b251.
//
// Solidity: function mine() view returns(address)
func (_ChunkLinearReward *ChunkLinearRewardSession) Mine() (common.Address, error) {
	return _ChunkLinearReward.Contract.Mine(&_ChunkLinearReward.CallOpts)
}

// Mine is a free data retrieval call binding the contract method 0x99f4b251.
//
// Solidity: function mine() view returns(address)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) Mine() (common.Address, error) {
	return _ChunkLinearReward.Contract.Mine(&_ChunkLinearReward.CallOpts)
}

// Payments is a free data retrieval call binding the contract method 0xe2982c21.
//
// Solidity: function payments(address dest) view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardCaller) Payments(opts *bind.CallOpts, dest common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "payments", dest)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Payments is a free data retrieval call binding the contract method 0xe2982c21.
//
// Solidity: function payments(address dest) view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardSession) Payments(dest common.Address) (*big.Int, error) {
	return _ChunkLinearReward.Contract.Payments(&_ChunkLinearReward.CallOpts, dest)
}

// Payments is a free data retrieval call binding the contract method 0xe2982c21.
//
// Solidity: function payments(address dest) view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) Payments(dest common.Address) (*big.Int, error) {
	return _ChunkLinearReward.Contract.Payments(&_ChunkLinearReward.CallOpts, dest)
}

// ReleaseSeconds is a free data retrieval call binding the contract method 0x21295931.
//
// Solidity: function releaseSeconds() view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardCaller) ReleaseSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "releaseSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReleaseSeconds is a free data retrieval call binding the contract method 0x21295931.
//
// Solidity: function releaseSeconds() view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardSession) ReleaseSeconds() (*big.Int, error) {
	return _ChunkLinearReward.Contract.ReleaseSeconds(&_ChunkLinearReward.CallOpts)
}

// ReleaseSeconds is a free data retrieval call binding the contract method 0x21295931.
//
// Solidity: function releaseSeconds() view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) ReleaseSeconds() (*big.Int, error) {
	return _ChunkLinearReward.Contract.ReleaseSeconds(&_ChunkLinearReward.CallOpts)
}

// RewardDeadline is a free data retrieval call binding the contract method 0x0a539a19.
//
// Solidity: function rewardDeadline(uint256 pricingIndex) view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardCaller) RewardDeadline(opts *bind.CallOpts, pricingIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "rewardDeadline", pricingIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardDeadline is a free data retrieval call binding the contract method 0x0a539a19.
//
// Solidity: function rewardDeadline(uint256 pricingIndex) view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardSession) RewardDeadline(pricingIndex *big.Int) (*big.Int, error) {
	return _ChunkLinearReward.Contract.RewardDeadline(&_ChunkLinearReward.CallOpts, pricingIndex)
}

// RewardDeadline is a free data retrieval call binding the contract method 0x0a539a19.
//
// Solidity: function rewardDeadline(uint256 pricingIndex) view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) RewardDeadline(pricingIndex *big.Int) (*big.Int, error) {
	return _ChunkLinearReward.Contract.RewardDeadline(&_ChunkLinearReward.CallOpts, pricingIndex)
}

// Rewards is a free data retrieval call binding the contract method 0xf301af42.
//
// Solidity: function rewards(uint256 ) view returns(uint128 lockedReward, uint128 claimableReward, uint128 distributedReward, uint40 startTime, uint40 lastUpdate)
func (_ChunkLinearReward *ChunkLinearRewardCaller) Rewards(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LockedReward      *big.Int
	ClaimableReward   *big.Int
	DistributedReward *big.Int
	StartTime         *big.Int
	LastUpdate        *big.Int
}, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "rewards", arg0)

	outstruct := new(struct {
		LockedReward      *big.Int
		ClaimableReward   *big.Int
		DistributedReward *big.Int
		StartTime         *big.Int
		LastUpdate        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LockedReward = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ClaimableReward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DistributedReward = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LastUpdate = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Rewards is a free data retrieval call binding the contract method 0xf301af42.
//
// Solidity: function rewards(uint256 ) view returns(uint128 lockedReward, uint128 claimableReward, uint128 distributedReward, uint40 startTime, uint40 lastUpdate)
func (_ChunkLinearReward *ChunkLinearRewardSession) Rewards(arg0 *big.Int) (struct {
	LockedReward      *big.Int
	ClaimableReward   *big.Int
	DistributedReward *big.Int
	StartTime         *big.Int
	LastUpdate        *big.Int
}, error) {
	return _ChunkLinearReward.Contract.Rewards(&_ChunkLinearReward.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0xf301af42.
//
// Solidity: function rewards(uint256 ) view returns(uint128 lockedReward, uint128 claimableReward, uint128 distributedReward, uint40 startTime, uint40 lastUpdate)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) Rewards(arg0 *big.Int) (struct {
	LockedReward      *big.Int
	ClaimableReward   *big.Int
	DistributedReward *big.Int
	StartTime         *big.Int
	LastUpdate        *big.Int
}, error) {
	return _ChunkLinearReward.Contract.Rewards(&_ChunkLinearReward.CallOpts, arg0)
}

// ServiceFeeRateBps is a free data retrieval call binding the contract method 0xc0575111.
//
// Solidity: function serviceFeeRateBps() view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardCaller) ServiceFeeRateBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "serviceFeeRateBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ServiceFeeRateBps is a free data retrieval call binding the contract method 0xc0575111.
//
// Solidity: function serviceFeeRateBps() view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardSession) ServiceFeeRateBps() (*big.Int, error) {
	return _ChunkLinearReward.Contract.ServiceFeeRateBps(&_ChunkLinearReward.CallOpts)
}

// ServiceFeeRateBps is a free data retrieval call binding the contract method 0xc0575111.
//
// Solidity: function serviceFeeRateBps() view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) ServiceFeeRateBps() (*big.Int, error) {
	return _ChunkLinearReward.Contract.ServiceFeeRateBps(&_ChunkLinearReward.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ChunkLinearReward *ChunkLinearRewardCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ChunkLinearReward *ChunkLinearRewardSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ChunkLinearReward.Contract.SupportsInterface(&_ChunkLinearReward.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ChunkLinearReward.Contract.SupportsInterface(&_ChunkLinearReward.CallOpts, interfaceId)
}

// TotalBaseReward is a free data retrieval call binding the contract method 0x7f1b5e43.
//
// Solidity: function totalBaseReward() view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardCaller) TotalBaseReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "totalBaseReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBaseReward is a free data retrieval call binding the contract method 0x7f1b5e43.
//
// Solidity: function totalBaseReward() view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardSession) TotalBaseReward() (*big.Int, error) {
	return _ChunkLinearReward.Contract.TotalBaseReward(&_ChunkLinearReward.CallOpts)
}

// TotalBaseReward is a free data retrieval call binding the contract method 0x7f1b5e43.
//
// Solidity: function totalBaseReward() view returns(uint256)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) TotalBaseReward() (*big.Int, error) {
	return _ChunkLinearReward.Contract.TotalBaseReward(&_ChunkLinearReward.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_ChunkLinearReward *ChunkLinearRewardCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChunkLinearReward.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_ChunkLinearReward *ChunkLinearRewardSession) Treasury() (common.Address, error) {
	return _ChunkLinearReward.Contract.Treasury(&_ChunkLinearReward.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_ChunkLinearReward *ChunkLinearRewardCallerSession) Treasury() (common.Address, error) {
	return _ChunkLinearReward.Contract.Treasury(&_ChunkLinearReward.CallOpts)
}

// ClaimMineReward is a paid mutator transaction binding the contract method 0xb7a3c04c.
//
// Solidity: function claimMineReward(uint256 pricingIndex, address beneficiary, bytes32 ) returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactor) ClaimMineReward(opts *bind.TransactOpts, pricingIndex *big.Int, beneficiary common.Address, arg2 [32]byte) (*types.Transaction, error) {
	return _ChunkLinearReward.contract.Transact(opts, "claimMineReward", pricingIndex, beneficiary, arg2)
}

// ClaimMineReward is a paid mutator transaction binding the contract method 0xb7a3c04c.
//
// Solidity: function claimMineReward(uint256 pricingIndex, address beneficiary, bytes32 ) returns()
func (_ChunkLinearReward *ChunkLinearRewardSession) ClaimMineReward(pricingIndex *big.Int, beneficiary common.Address, arg2 [32]byte) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.ClaimMineReward(&_ChunkLinearReward.TransactOpts, pricingIndex, beneficiary, arg2)
}

// ClaimMineReward is a paid mutator transaction binding the contract method 0xb7a3c04c.
//
// Solidity: function claimMineReward(uint256 pricingIndex, address beneficiary, bytes32 ) returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactorSession) ClaimMineReward(pricingIndex *big.Int, beneficiary common.Address, arg2 [32]byte) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.ClaimMineReward(&_ChunkLinearReward.TransactOpts, pricingIndex, beneficiary, arg2)
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactor) Donate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChunkLinearReward.contract.Transact(opts, "donate")
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_ChunkLinearReward *ChunkLinearRewardSession) Donate() (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.Donate(&_ChunkLinearReward.TransactOpts)
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactorSession) Donate() (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.Donate(&_ChunkLinearReward.TransactOpts)
}

// FillReward is a paid mutator transaction binding the contract method 0x59e96700.
//
// Solidity: function fillReward(uint256 beforeLength, uint256 chargedSectors) payable returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactor) FillReward(opts *bind.TransactOpts, beforeLength *big.Int, chargedSectors *big.Int) (*types.Transaction, error) {
	return _ChunkLinearReward.contract.Transact(opts, "fillReward", beforeLength, chargedSectors)
}

// FillReward is a paid mutator transaction binding the contract method 0x59e96700.
//
// Solidity: function fillReward(uint256 beforeLength, uint256 chargedSectors) payable returns()
func (_ChunkLinearReward *ChunkLinearRewardSession) FillReward(beforeLength *big.Int, chargedSectors *big.Int) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.FillReward(&_ChunkLinearReward.TransactOpts, beforeLength, chargedSectors)
}

// FillReward is a paid mutator transaction binding the contract method 0x59e96700.
//
// Solidity: function fillReward(uint256 beforeLength, uint256 chargedSectors) payable returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactorSession) FillReward(beforeLength *big.Int, chargedSectors *big.Int) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.FillReward(&_ChunkLinearReward.TransactOpts, beforeLength, chargedSectors)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ChunkLinearReward.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ChunkLinearReward *ChunkLinearRewardSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.GrantRole(&_ChunkLinearReward.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.GrantRole(&_ChunkLinearReward.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address market_, address mine_) returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactor) Initialize(opts *bind.TransactOpts, market_ common.Address, mine_ common.Address) (*types.Transaction, error) {
	return _ChunkLinearReward.contract.Transact(opts, "initialize", market_, mine_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address market_, address mine_) returns()
func (_ChunkLinearReward *ChunkLinearRewardSession) Initialize(market_ common.Address, mine_ common.Address) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.Initialize(&_ChunkLinearReward.TransactOpts, market_, mine_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address market_, address mine_) returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactorSession) Initialize(market_ common.Address, mine_ common.Address) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.Initialize(&_ChunkLinearReward.TransactOpts, market_, mine_)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ChunkLinearReward.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ChunkLinearReward *ChunkLinearRewardSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.RenounceRole(&_ChunkLinearReward.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.RenounceRole(&_ChunkLinearReward.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ChunkLinearReward.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ChunkLinearReward *ChunkLinearRewardSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.RevokeRole(&_ChunkLinearReward.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.RevokeRole(&_ChunkLinearReward.TransactOpts, role, account)
}

// SetBaseReward is a paid mutator transaction binding the contract method 0x0373a23a.
//
// Solidity: function setBaseReward(uint256 baseReward_) returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactor) SetBaseReward(opts *bind.TransactOpts, baseReward_ *big.Int) (*types.Transaction, error) {
	return _ChunkLinearReward.contract.Transact(opts, "setBaseReward", baseReward_)
}

// SetBaseReward is a paid mutator transaction binding the contract method 0x0373a23a.
//
// Solidity: function setBaseReward(uint256 baseReward_) returns()
func (_ChunkLinearReward *ChunkLinearRewardSession) SetBaseReward(baseReward_ *big.Int) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.SetBaseReward(&_ChunkLinearReward.TransactOpts, baseReward_)
}

// SetBaseReward is a paid mutator transaction binding the contract method 0x0373a23a.
//
// Solidity: function setBaseReward(uint256 baseReward_) returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactorSession) SetBaseReward(baseReward_ *big.Int) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.SetBaseReward(&_ChunkLinearReward.TransactOpts, baseReward_)
}

// SetServiceFeeRate is a paid mutator transaction binding the contract method 0x9b1d3091.
//
// Solidity: function setServiceFeeRate(uint256 bps) returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactor) SetServiceFeeRate(opts *bind.TransactOpts, bps *big.Int) (*types.Transaction, error) {
	return _ChunkLinearReward.contract.Transact(opts, "setServiceFeeRate", bps)
}

// SetServiceFeeRate is a paid mutator transaction binding the contract method 0x9b1d3091.
//
// Solidity: function setServiceFeeRate(uint256 bps) returns()
func (_ChunkLinearReward *ChunkLinearRewardSession) SetServiceFeeRate(bps *big.Int) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.SetServiceFeeRate(&_ChunkLinearReward.TransactOpts, bps)
}

// SetServiceFeeRate is a paid mutator transaction binding the contract method 0x9b1d3091.
//
// Solidity: function setServiceFeeRate(uint256 bps) returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactorSession) SetServiceFeeRate(bps *big.Int) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.SetServiceFeeRate(&_ChunkLinearReward.TransactOpts, bps)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address treasury_) returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactor) SetTreasury(opts *bind.TransactOpts, treasury_ common.Address) (*types.Transaction, error) {
	return _ChunkLinearReward.contract.Transact(opts, "setTreasury", treasury_)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address treasury_) returns()
func (_ChunkLinearReward *ChunkLinearRewardSession) SetTreasury(treasury_ common.Address) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.SetTreasury(&_ChunkLinearReward.TransactOpts, treasury_)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address treasury_) returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactorSession) SetTreasury(treasury_ common.Address) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.SetTreasury(&_ChunkLinearReward.TransactOpts, treasury_)
}

// WithdrawPayments is a paid mutator transaction binding the contract method 0x31b3eb94.
//
// Solidity: function withdrawPayments(address payee) returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactor) WithdrawPayments(opts *bind.TransactOpts, payee common.Address) (*types.Transaction, error) {
	return _ChunkLinearReward.contract.Transact(opts, "withdrawPayments", payee)
}

// WithdrawPayments is a paid mutator transaction binding the contract method 0x31b3eb94.
//
// Solidity: function withdrawPayments(address payee) returns()
func (_ChunkLinearReward *ChunkLinearRewardSession) WithdrawPayments(payee common.Address) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.WithdrawPayments(&_ChunkLinearReward.TransactOpts, payee)
}

// WithdrawPayments is a paid mutator transaction binding the contract method 0x31b3eb94.
//
// Solidity: function withdrawPayments(address payee) returns()
func (_ChunkLinearReward *ChunkLinearRewardTransactorSession) WithdrawPayments(payee common.Address) (*types.Transaction, error) {
	return _ChunkLinearReward.Contract.WithdrawPayments(&_ChunkLinearReward.TransactOpts, payee)
}

// ChunkLinearRewardDistributeRewardIterator is returned from FilterDistributeReward and is used to iterate over the raw logs and unpacked data for DistributeReward events raised by the ChunkLinearReward contract.
type ChunkLinearRewardDistributeRewardIterator struct {
	Event *ChunkLinearRewardDistributeReward // Event containing the contract specifics and raw log

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
func (it *ChunkLinearRewardDistributeRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChunkLinearRewardDistributeReward)
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
		it.Event = new(ChunkLinearRewardDistributeReward)
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
func (it *ChunkLinearRewardDistributeRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChunkLinearRewardDistributeRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChunkLinearRewardDistributeReward represents a DistributeReward event raised by the ChunkLinearReward contract.
type ChunkLinearRewardDistributeReward struct {
	PricingIndex *big.Int
	Beneficiary  common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDistributeReward is a free log retrieval operation binding the contract event 0x83617a1b0f847971f005bd162dde513cfe93df96e6293c3bbb5fe9c40629dd4c.
//
// Solidity: event DistributeReward(uint256 indexed pricingIndex, address indexed beneficiary, uint256 amount)
func (_ChunkLinearReward *ChunkLinearRewardFilterer) FilterDistributeReward(opts *bind.FilterOpts, pricingIndex []*big.Int, beneficiary []common.Address) (*ChunkLinearRewardDistributeRewardIterator, error) {

	var pricingIndexRule []interface{}
	for _, pricingIndexItem := range pricingIndex {
		pricingIndexRule = append(pricingIndexRule, pricingIndexItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _ChunkLinearReward.contract.FilterLogs(opts, "DistributeReward", pricingIndexRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &ChunkLinearRewardDistributeRewardIterator{contract: _ChunkLinearReward.contract, event: "DistributeReward", logs: logs, sub: sub}, nil
}

// WatchDistributeReward is a free log subscription operation binding the contract event 0x83617a1b0f847971f005bd162dde513cfe93df96e6293c3bbb5fe9c40629dd4c.
//
// Solidity: event DistributeReward(uint256 indexed pricingIndex, address indexed beneficiary, uint256 amount)
func (_ChunkLinearReward *ChunkLinearRewardFilterer) WatchDistributeReward(opts *bind.WatchOpts, sink chan<- *ChunkLinearRewardDistributeReward, pricingIndex []*big.Int, beneficiary []common.Address) (event.Subscription, error) {

	var pricingIndexRule []interface{}
	for _, pricingIndexItem := range pricingIndex {
		pricingIndexRule = append(pricingIndexRule, pricingIndexItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _ChunkLinearReward.contract.WatchLogs(opts, "DistributeReward", pricingIndexRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChunkLinearRewardDistributeReward)
				if err := _ChunkLinearReward.contract.UnpackLog(event, "DistributeReward", log); err != nil {
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

// ParseDistributeReward is a log parse operation binding the contract event 0x83617a1b0f847971f005bd162dde513cfe93df96e6293c3bbb5fe9c40629dd4c.
//
// Solidity: event DistributeReward(uint256 indexed pricingIndex, address indexed beneficiary, uint256 amount)
func (_ChunkLinearReward *ChunkLinearRewardFilterer) ParseDistributeReward(log types.Log) (*ChunkLinearRewardDistributeReward, error) {
	event := new(ChunkLinearRewardDistributeReward)
	if err := _ChunkLinearReward.contract.UnpackLog(event, "DistributeReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChunkLinearRewardRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ChunkLinearReward contract.
type ChunkLinearRewardRoleAdminChangedIterator struct {
	Event *ChunkLinearRewardRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ChunkLinearRewardRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChunkLinearRewardRoleAdminChanged)
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
		it.Event = new(ChunkLinearRewardRoleAdminChanged)
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
func (it *ChunkLinearRewardRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChunkLinearRewardRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChunkLinearRewardRoleAdminChanged represents a RoleAdminChanged event raised by the ChunkLinearReward contract.
type ChunkLinearRewardRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ChunkLinearReward *ChunkLinearRewardFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ChunkLinearRewardRoleAdminChangedIterator, error) {

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

	logs, sub, err := _ChunkLinearReward.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ChunkLinearRewardRoleAdminChangedIterator{contract: _ChunkLinearReward.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ChunkLinearReward *ChunkLinearRewardFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ChunkLinearRewardRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ChunkLinearReward.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChunkLinearRewardRoleAdminChanged)
				if err := _ChunkLinearReward.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_ChunkLinearReward *ChunkLinearRewardFilterer) ParseRoleAdminChanged(log types.Log) (*ChunkLinearRewardRoleAdminChanged, error) {
	event := new(ChunkLinearRewardRoleAdminChanged)
	if err := _ChunkLinearReward.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChunkLinearRewardRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ChunkLinearReward contract.
type ChunkLinearRewardRoleGrantedIterator struct {
	Event *ChunkLinearRewardRoleGranted // Event containing the contract specifics and raw log

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
func (it *ChunkLinearRewardRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChunkLinearRewardRoleGranted)
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
		it.Event = new(ChunkLinearRewardRoleGranted)
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
func (it *ChunkLinearRewardRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChunkLinearRewardRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChunkLinearRewardRoleGranted represents a RoleGranted event raised by the ChunkLinearReward contract.
type ChunkLinearRewardRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ChunkLinearReward *ChunkLinearRewardFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ChunkLinearRewardRoleGrantedIterator, error) {

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

	logs, sub, err := _ChunkLinearReward.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ChunkLinearRewardRoleGrantedIterator{contract: _ChunkLinearReward.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ChunkLinearReward *ChunkLinearRewardFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ChunkLinearRewardRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ChunkLinearReward.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChunkLinearRewardRoleGranted)
				if err := _ChunkLinearReward.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_ChunkLinearReward *ChunkLinearRewardFilterer) ParseRoleGranted(log types.Log) (*ChunkLinearRewardRoleGranted, error) {
	event := new(ChunkLinearRewardRoleGranted)
	if err := _ChunkLinearReward.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChunkLinearRewardRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ChunkLinearReward contract.
type ChunkLinearRewardRoleRevokedIterator struct {
	Event *ChunkLinearRewardRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ChunkLinearRewardRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChunkLinearRewardRoleRevoked)
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
		it.Event = new(ChunkLinearRewardRoleRevoked)
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
func (it *ChunkLinearRewardRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChunkLinearRewardRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChunkLinearRewardRoleRevoked represents a RoleRevoked event raised by the ChunkLinearReward contract.
type ChunkLinearRewardRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ChunkLinearReward *ChunkLinearRewardFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ChunkLinearRewardRoleRevokedIterator, error) {

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

	logs, sub, err := _ChunkLinearReward.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ChunkLinearRewardRoleRevokedIterator{contract: _ChunkLinearReward.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ChunkLinearReward *ChunkLinearRewardFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ChunkLinearRewardRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ChunkLinearReward.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChunkLinearRewardRoleRevoked)
				if err := _ChunkLinearReward.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_ChunkLinearReward *ChunkLinearRewardFilterer) ParseRoleRevoked(log types.Log) (*ChunkLinearRewardRoleRevoked, error) {
	event := new(ChunkLinearRewardRoleRevoked)
	if err := _ChunkLinearReward.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
