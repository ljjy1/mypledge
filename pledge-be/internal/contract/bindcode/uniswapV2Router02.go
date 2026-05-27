// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindcode

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

// UniswapV2Router02MetaData contains all meta data concerning the UniswapV2Router02 contract.
var UniswapV2Router02MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60c0346100e257601f61255938819003918201601f19168301916001600160401b038311848410176100e65780849260409485528339810103126100e257610052602061004b836100fa565b92016100fa565b9060805260a05260405161244a908161010f82396080518181816101ab015281816105af0152818161081501528181610960015281816115c00152818161184d0152818161195801528181611d690152612036015260a051818181601a0152818161017c015281816108cd01528181610c2e01528181610e1a01528181610fa30152818161106d01526110e10152f35b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b03821682036100e25756fe6080604052600436101561008f575b3615610018575f80fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316330361004a57005b60405162461bcd60e51b815260206004820152601f60248201527f556e69737761705632526f757465723a20494e56414c49445f53454e444552006044820152606490fd5b5f5f3560e01c806302751cec146110c357806318cbafe5146110355780631f00ca741461102057806338ed173914610fd25780633fc8cef314610f8d5780634a25d94a14610dce5780635c11d79514610d495780637ff36ab514610c015780638803dbee14610b3f578063b6f9de95146108a1578063baa2abde14610844578063c45a0155146107ff578063d06ca61f146107d8578063e8e33700146105355763f305d7191461013f575061000e565b61015d61014b366111a1565b879497969193929596504211156113d9565b60405163e6a4390560e01b81526001600160a01b0386811660048301527f000000000000000000000000000000000000000000000000000000000000000080821660248401529195919291907f0000000000000000000000000000000000000000000000000000000000000000908116602088604481845afa978815610509578798610514575b506001600160a01b0388161561047e575b50906102088489899a9594999899611f22565b8193911580610476575b15610365575050505061022a9034979687915b611cd6565b6001600160a01b038116803b1561036157848791600460405180948193630d0e30db60e41b83525af180156103565791879186959493610333575b50916102776024928560209695611aee565b6040516335313c2160e11b81526001600160a01b0391821660048201529586938492165af190811561032757906102ef575b6102d391508334116102d7575b604051938493846040919493926060820195825260208201520152565b0390f35b6102ea6102e48534611425565b33611b73565b6102b6565b506020813d60201161031f575b81610309602093836112a0565b8101031261031b576102d390516102a9565b5f80fd5b3d91506102fc565b604051903d90823e3d90fd5b85925094610343919394956112a0565b6103525790858493925f610265565b8380fd5b6040513d87823e3d90fd5b8480fd5b90919293965061037e84610379838d611a1b565b611a2e565b933485116103f457505050811061039e57859361022a9197968791611cd6565b60405162461bcd60e51b815260206004820152602860248201527f556e69737761705632526f757465723a20494e53554646494349454e545f45546044820152671217d05353d5539560c21b6064820152608490fd5b909250610379919950610408935034611a1b565b95861061041e5761022a85933497968791611cd6565b60405162461bcd60e51b815260206004820152602a60248201527f556e69737761705632526f757465723a20494e53554646494349454e545f544f60448201526912d15397d05353d5539560b21b6064820152608490fd5b508015610212565b6040516364e329cb60e11b81526001600160a01b038a8116600483015286166024820152969750909590602090829060449082908b905af1908115610509578489949392610208928a916104da575b50989792939450506101f5565b6104fc915060203d602011610502575b6104f481836112a0565b8101906119fc565b5f6104cd565b503d6104ea565b6040513d89823e3d90fd5b61052e91985060203d602011610502576104f481836112a0565b965f6101e4565b50346107d5576101003660031901126107d557610550611175565b9061055961118b565b606435919060c4356001600160a01b0381169160443591839003610352576105854260e43510156113d9565b60405163e6a4390560e01b81526001600160a01b03878116600483015282811660248301529095907f0000000000000000000000000000000000000000000000000000000000000000908116602088604481845afa978815610509579089929188996107b4575b506001600160a01b03891615610722575b50602094888560249561061783988c9b9761063797611f22565b819291158061071a575b156106c0575061022591509c8d949c8d91611cd6565b6040516335313c2160e11b815260048101919091529485928391906001600160a01b03165af1908115610327579061068c575b6102d39150604051938493846040919493926060820195825260208201520152565b506020813d6020116106b8575b816106a6602093836112a0565b8101031261031b576102d3905161066a565b3d9150610699565b6106ce836103798389611a1b565b928284116106f657505050610225906106eb60a4358210156117d2565b9c8d949c8d91611cd6565b82965061022593506103799061070b93611a1b565b936106eb608435861015611777565b508015610621565b6040516364e329cb60e11b81526001600160a01b039384166004820152928516602484015293975092602090829060449082908a905af19081156107a9576020948994896024956106178b9a96610637968c9161078c575b509c50509550959397509150946105fd565b6107a391508b3d8d11610502576104f481836112a0565b5f61077a565b6040513d88823e3d90fd5b6107ce91995060203d602011610502576104f481836112a0565b975f6105ec565b80fd5b50346107d5576102d36107f36107ed36611349565b90611930565b60405191829182611267565b50346107d557806003193601126107d5576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b50346107d55760e03660031901126107d55761085e611175565b9061086761118b565b9160a435916001600160a01b03831683036107d557604061089560c435856084356064356044358a8961182d565b82519182526020820152f35b506108bc6108ae3661138d565b9490949291924211156113d9565b8115610b2b576108cb8161146a565b7f00000000000000000000000000000000000000000000000000000000000000009086906001600160a01b03808416916109079116821461147e565b803b15610b27578160049160405192838092630d0e30db60e41b825234905af18015610b1c57610b03575b505061093d8261146a565b908360011015610aef579061098a91610984349261095d6020870161146a565b907f0000000000000000000000000000000000000000000000000000000000000000611bdc565b90611aee565b5f198201828111610adb576001600160a01b036109b06109ab838686611446565b61146a565b1692604051956370a0823160e01b875260208760248160018060a01b038516988960048301525afa968715610ad0578897610a95575b5060209392610a1492610a066109ab93610a013684896112ee565b612030565b6001600160a01b0394611446565b16916024604051809481936370a0823160e01b835260048301525afa908115610a8a578491610a56575b50610a5392610a4c91611425565b10156114eb565b80f35b90506020813d602011610a82575b81610a71602093836112a0565b8101031261031b5751610a53610a3e565b3d9150610a64565b6040513d86823e3d90fd5b9096506020939293813d602011610ac8575b81610ab4602093836112a0565b8101031261031b57519591929160206109e6565b3d9150610aa7565b6040513d8a823e3d90fd5b634e487b7160e01b86526011600452602486fd5b634e487b7160e01b87526032600452602487fd5b81610b0d916112a0565b610b1857855f610932565b8580fd5b6040513d84823e3d90fd5b5080fd5b634e487b7160e01b85526032600452602485fd5b50346107d557610b7a610b69610b8c610b5736611218565b989491959397929690984211156113d9565b610b743686886112ee565b9061157d565b94610b84866114ca565b51111561171b565b8115610bed57610b9b8361146a565b90610ba58461146a565b908360011015610bed5750926102d39592610be0610be793610bd06107f39761095d6020870161146a565b610bd9896114ca565b5191611cd6565b36916112ee565b83611d5e565b634e487b7160e01b81526032600452602490fd5b50610c1d610c0e3661138d565b959192939490954211156113d9565b8115610bed57610c2c8361146a565b7f000000000000000000000000000000000000000000000000000000000000000091906001600160a01b0380841691610c679116821461147e565b610c7b610c753686886112ee565b34611930565b9586515f198101908111610d355790610c97610c9f92896114d7565b5110156114eb565b610ca8866114ca565b5190803b15610d31578290600460405180948193630d0e30db60e41b83525af18015610b1c57908291610d1c575b5050610ce18461146a565b908360011015610bed5750926102d39592610be0610be793610d0c6107f39761095d6020870161146a565b610d15896114ca565b5191611aee565b81610d26916112a0565b6107d557805f610cd6565b8280fd5b634e487b7160e01b84526011600452602484fd5b50346107d557610d6a610d5b36611218565b959095949193944211156113d9565b8215610dba57610d798261146a565b610d828361146a565b8460011015610da65790610da061098a939261095d6020870161146a565b90611cd6565b634e487b7160e01b88526032600452602488fd5b634e487b7160e01b86526032600452602486fd5b50346107d557610df1610de036611218565b9694939590969291924211156113d9565b5f198201828111610f7957610e54610e106109ab610e5e938686611446565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081169691610e499116871461147e565b610b743686866112ee565b95610b84876114ca565b8115610f6557610e6d8161146a565b610e768261146a565b8360011015610dba5791610ea4610ead92610e9b610eb3969561095d6020860161146a565b610bd98a6114ca565b309236916112ee565b85611d5e565b82515f198101908111610f5157610eca90846114d7565b51813b15610d31578291602483926040519485938492632e1a7d4d60e01b845260048401525af18015610b1c57908291610f3c575b505081515f19810191908211610f2857506102d392610f216107f392846114d7565b5190611b73565b634e487b7160e01b81526011600452602490fd5b81610f46916112a0565b6107d557805f610eff565b634e487b7160e01b83526011600452602483fd5b634e487b7160e01b84526032600452602484fd5b634e487b7160e01b85526011600452602485fd5b50346107d557806003193601126107d5576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b50346107d557611004610ff9610fe736611218565b979396929590979491944211156113d9565b6107ed3685876112ee565b9384515f198101908111610f515790610c97610b8c92876114d7565b50346107d5576102d36107f3610b7436611349565b50346107d557611047610de036611218565b5f198201828111610f79576110636109ab6110a7928585611446565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116959161109c9116861461147e565b6107ed3685856112ee565b9485515f198101908111610adb5790610c97610e5e92886114d7565b3461031b5761111261110a6110d7366111a1565b90969293879596927f0000000000000000000000000000000000000000000000000000000000000000953092878a61182d565b938195611aee565b6001600160a01b0316803b1561031b575f8091602460405180948193632e1a7d4d60e01b83528760048401525af1801561116a5760409461089592849261115a575b50611b73565b5f611164916112a0565b5f611154565b6040513d5f823e3d90fd5b600435906001600160a01b038216820361031b57565b602435906001600160a01b038216820361031b57565b60c090600319011261031b576004356001600160a01b038116810361031b57906024359060443590606435906084356001600160a01b038116810361031b579060a43590565b9181601f8401121561031b5782359167ffffffffffffffff831161031b576020808501948460051b01011161031b57565b60a060031982011261031b5760043591602435916044359067ffffffffffffffff821161031b5761124b916004016111e7565b90916064356001600160a01b038116810361031b579060843590565b60206040818301928281528451809452019201905f5b81811061128a5750505090565b825184526020938401939092019160010161127d565b90601f8019910116810190811067ffffffffffffffff8211176112c257604052565b634e487b7160e01b5f52604160045260245ffd5b67ffffffffffffffff81116112c25760051b60200190565b92916112f9826112d6565b9361130760405195866112a0565b602085848152019260051b810191821161031b57915b81831061132957505050565b82356001600160a01b038116810361031b5781526020928301920161131d565b90604060031983011261031b57600435916024359067ffffffffffffffff821161031b578060238301121561031b5781602461138a936004013591016112ee565b90565b90608060031983011261031b57600435916024359067ffffffffffffffff821161031b576113bd916004016111e7565b90916044356001600160a01b038116810361031b579060643590565b156113e057565b60405162461bcd60e51b815260206004820152601860248201527f556e69737761705632526f757465723a204558504952454400000000000000006044820152606490fd5b9190820391821161143257565b634e487b7160e01b5f52601160045260245ffd5b91908110156114565760051b0190565b634e487b7160e01b5f52603260045260245ffd5b356001600160a01b038116810361031b5790565b1561148557565b60405162461bcd60e51b815260206004820152601d60248201527f556e69737761705632526f757465723a20494e56414c49445f504154480000006044820152606490fd5b8051156114565760200190565b80518210156114565760209160051b010190565b156114f257565b60405162461bcd60e51b815260206004820152602b60248201527f556e69737761705632526f757465723a20494e53554646494349454e545f4f5560448201526a1514155517d05353d5539560aa1b6064820152608490fd5b90611555826112d6565b61156260405191826112a0565b8281528092611573601f19916112d6565b0190602036910137565b919061158d60028251101561147e565b611597815161154b565b9283515f198101908111611432576115af90856114d7565b5280515f19810190811161143257807f0000000000000000000000000000000000000000000000000000000000000000915b6115ea57505050565b5f198101818111611432576116276001600160a01b0361160a83876114d7565b51166001600160a01b0361161e85886114d7565b51169085611f22565b61163184896114d7565b519182156116c1578281611655921515806116b8575b61165090611fd3565b611a1b565b916103e88302928084046103e814901517156114325761167491611425565b6103e58102908082046103e514901517156114325761169291611a2e565b9060018201809211611432576116a890876114d7565b528015611432575f1901806115e1565b50831515611647565b60405162461bcd60e51b815260206004820152602c60248201527f556e697377617056324c6962726172793a20494e53554646494349454e545f4f60448201526b155514155517d05353d5539560a21b6064820152608490fd5b1561172257565b60405162461bcd60e51b815260206004820152602760248201527f556e69737761705632526f757465723a204558434553534956455f494e50555460448201526617d05353d5539560ca1b6064820152608490fd5b1561177e57565b60405162461bcd60e51b815260206004820152602660248201527f556e69737761705632526f757465723a20494e53554646494349454e545f415f604482015265105353d5539560d21b6064820152608490fd5b156117d957565b60405162461bcd60e51b815260206004820152602660248201527f556e69737761705632526f757465723a20494e53554646494349454e545f425f604482015265105353d5539560d21b6064820152608490fd5b91602460409298966118435f96994211156113d9565b6118796118718b877f0000000000000000000000000000000000000000000000000000000000000000611bdc565b938480611cd6565b835163226bf2d160e21b81526001600160a01b0391821660048201529586938492165af195861561116a575f925f976118f2575b506118b8908261228a565b506001600160a01b039182169116036118e857906118e691945b6118df8195871015611777565b10156117d2565b565b6118e691906118d2565b925095506040823d604011611928575b8161190f604093836112a0565b8101031261031b576118b86020835193015196906118ad565b3d9150611902565b91909161194160028451101561147e565b61194b835161154b565b90611955826114ca565b527f00000000000000000000000000000000000000000000000000000000000000005f5b84515f198101908111611432578110156119f6576001600160a01b0361199f82876114d7565b51169060018101808211611432576001926119ef906119e8906119d7906001600160a01b036119ce868d6114d7565b51169088611f22565b906119e2868a6114d7565b51612359565b91866114d7565b5201611979565b50509150565b9081602091031261031b57516001600160a01b038116810361031b5790565b8181029291811591840414171561143257565b8115611a38570490565b634e487b7160e01b5f52601260045260245ffd5b3d15611a86573d9067ffffffffffffffff82116112c25760405191611a7b601f8201601f1916602001846112a0565b82523d5f602084013e565b606090565b9081602091031261031b5751801515810361031b5790565b15611aaa57565b606460405162461bcd60e51b815260206004820152602060248201527f556e69737761705632526f757465723a205452414e534645525f4641494c45446044820152fd5b5f91908291826118e69560405190602082019363a9059cbb60e01b855260018060a01b03166024830152604482015260448152611b2c6064826112a0565b51925af1611b38611a4c565b81611b44575b50611aa3565b8051801592508215611b59575b50505f611b3e565b611b6c9250602080918301019101611a8b565b5f80611b51565b5f80809381935af1611b83611a4c565b5015611b8b57565b60405162461bcd60e51b8152602060048201526024808201527f556e69737761705632526f757465723a204554485f5452414e534645525f46416044820152631253115160e21b6064820152608490fd5b91611be69161228a565b604051636da62a2f60e11b81529291906020846004816001600160a01b0387165afa93841561116a575f94611ca2575b506040519060208201926001600160601b03199060601b1683526001600160601b03199060601b16603482015260288152611c526048826112a0565b5190209160405192602084019260ff60f81b84526001600160601b03199060601b1660218501526035840152605583015260558252611c926075836112a0565b905190206001600160a01b031690565b9093506020813d602011611cce575b81611cbe602093836112a0565b8101031261031b5751925f611c16565b3d9150611cb1565b5f91908291826118e6956040519060208201936323b872dd60e01b855233602484015260018060a01b03166044830152606482015260648152611b2c6084826112a0565b9260209260a09592855283850152600180861b03166040840152608060608401528051918291826080860152018484015e5f828201840152601f01601f1916010190565b602093929091905f907f0000000000000000000000000000000000000000000000000000000000000000825b84515f19810190811161143257811015611ece576001600160a01b03611db082876114d7565b51169060018101808211611432576001600160a01b03611dd082896114d7565b5116611de7611ddf828661228a565b50928a6114d7565b51916001600160a01b03168403611ec7575f91935b8851600119810190811161143257841015611ebd576002840180851161143257611e5190611e40906001600160a01b0390611e37908d6114d7565b51168489611bdc565b925b6001600160a01b039288611bdc565b1660405194611e608d876112a0565b5f865288368e880137813b1561031b575f8094611e936040519889968795869463022c0d9f60e01b865260048601611d1a565b03925af191821561116a57600192611ead575b5001611d8a565b5f611eb7916112a0565b5f611ea6565b611e518792611e42565b5f93611dfc565b5050505050509050565b51906001600160701b038216820361031b57565b9081606091031261031b57611f0081611ed8565b916040611f0f60208401611ed8565b92015163ffffffff8116810361031b5790565b906060600492611f47611f35868561228a565b50956001600160a01b03928590611bdc565b1660405193848092630240bc6b60e21b82525afa91821561116a575f905f93611f94575b506001600160701b03928316939216916001600160a01b03918216911603611f905791565b9091565b6001600160701b039350839150611fc29060603d606011611fcc575b611fba81836112a0565b810190611eec565b5093909150611f6b565b503d611fb0565b15611fda57565b60405162461bcd60e51b815260206004820152602860248201527f556e697377617056324c6962726172793a20494e53554646494349454e545f4c604482015267495155494449545960c01b6064820152608490fd5b602092917f0000000000000000000000000000000000000000000000000000000000000000915f91825b82515f19810190811161143257811015612281576001600160a01b0361208082856114d7565b51169060018101808211611432576001600160a01b03906120a190866114d7565b5116916120ae838261228a565b506001600160a01b036120c285848b611bdc565b1660405192630240bc6b60e21b8452606084600481855afa92831561116a578c5f955f9561224e575b506001600160a01b0390911682149492936024936001600160701b0391821692911686156122485791925b604051948580926370a0823160e01b82528860048301525afa92831561116a575f93612217575b5061214b8161215094611425565b612359565b9115612210575f91935b86516001198101908111611432578410156122085760028401908185116114325761219b916001600160a01b0390612192908a6114d7565b5116908a611bdc565b905b604051946121ab8c876112a0565b5f865288368d880137813b1561031b575f80946121de6040519889968795869463022c0d9f60e01b865260048601611d1a565b03925af191821561116a576001926121f8575b500161205a565b5f612202916112a0565b5f6121f1565b50849061219d565b5f9361215a565b92508c83813d8311612241575b61222e81836112a0565b8101031261031b5791519161214b61213d565b503d612224565b92612116565b6001600160701b0396506024949550612275879160603d8111611fcc57611fba81836112a0565b509097509594506120eb565b50505050509050565b9091906001600160a01b0380841690821680821461230657101561230157915b906001600160a01b038316156122bc57565b60405162461bcd60e51b815260206004820152601e60248201527f556e697377617056324c6962726172793a205a45524f5f4144445245535300006044820152606490fd5b6122aa565b60405162461bcd60e51b815260206004820152602560248201527f556e697377617056324c6962726172793a204944454e544943414c5f41444452604482015264455353455360d81b6064820152608490fd5b80156123bb578115928315806123b2575b61237390611fd3565b6103e582029182046103e503611432578161238d91611a1b565b926103e883029283046103e81417156114325781018091116114325761138a91611a2e565b5080151561236a565b60405162461bcd60e51b815260206004820152602b60248201527f556e697377617056324c6962726172793a20494e53554646494349454e545f4960448201526a1394155517d05353d5539560aa1b6064820152608490fdfea2646970667358221220c6a1813c1a70f161e24206f062a36620a500d97262369075a3bb520bb8ef7c8564736f6c634300081c0033",
}

// UniswapV2Router02ABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapV2Router02MetaData.ABI instead.
var UniswapV2Router02ABI = UniswapV2Router02MetaData.ABI

// UniswapV2Router02Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UniswapV2Router02MetaData.Bin instead.
var UniswapV2Router02Bin = UniswapV2Router02MetaData.Bin

// DeployUniswapV2Router02 deploys a new Ethereum contract, binding an instance of UniswapV2Router02 to it.
func DeployUniswapV2Router02(auth *bind.TransactOpts, backend bind.ContractBackend, _factory common.Address, _weth common.Address) (common.Address, *types.Transaction, *UniswapV2Router02, error) {
	parsed, err := UniswapV2Router02MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UniswapV2Router02Bin), backend, _factory, _weth)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UniswapV2Router02{UniswapV2Router02Caller: UniswapV2Router02Caller{contract: contract}, UniswapV2Router02Transactor: UniswapV2Router02Transactor{contract: contract}, UniswapV2Router02Filterer: UniswapV2Router02Filterer{contract: contract}}, nil
}

// UniswapV2Router02 is an auto generated Go binding around an Ethereum contract.
type UniswapV2Router02 struct {
	UniswapV2Router02Caller     // Read-only binding to the contract
	UniswapV2Router02Transactor // Write-only binding to the contract
	UniswapV2Router02Filterer   // Log filterer for contract events
}

// UniswapV2Router02Caller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapV2Router02Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV2Router02Transactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapV2Router02Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV2Router02Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapV2Router02Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV2Router02Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapV2Router02Session struct {
	Contract     *UniswapV2Router02 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// UniswapV2Router02CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapV2Router02CallerSession struct {
	Contract *UniswapV2Router02Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// UniswapV2Router02TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapV2Router02TransactorSession struct {
	Contract     *UniswapV2Router02Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// UniswapV2Router02Raw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapV2Router02Raw struct {
	Contract *UniswapV2Router02 // Generic contract binding to access the raw methods on
}

// UniswapV2Router02CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapV2Router02CallerRaw struct {
	Contract *UniswapV2Router02Caller // Generic read-only contract binding to access the raw methods on
}

// UniswapV2Router02TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapV2Router02TransactorRaw struct {
	Contract *UniswapV2Router02Transactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapV2Router02 creates a new instance of UniswapV2Router02, bound to a specific deployed contract.
func NewUniswapV2Router02(address common.Address, backend bind.ContractBackend) (*UniswapV2Router02, error) {
	contract, err := bindUniswapV2Router02(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapV2Router02{UniswapV2Router02Caller: UniswapV2Router02Caller{contract: contract}, UniswapV2Router02Transactor: UniswapV2Router02Transactor{contract: contract}, UniswapV2Router02Filterer: UniswapV2Router02Filterer{contract: contract}}, nil
}

// NewUniswapV2Router02Caller creates a new read-only instance of UniswapV2Router02, bound to a specific deployed contract.
func NewUniswapV2Router02Caller(address common.Address, caller bind.ContractCaller) (*UniswapV2Router02Caller, error) {
	contract, err := bindUniswapV2Router02(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV2Router02Caller{contract: contract}, nil
}

// NewUniswapV2Router02Transactor creates a new write-only instance of UniswapV2Router02, bound to a specific deployed contract.
func NewUniswapV2Router02Transactor(address common.Address, transactor bind.ContractTransactor) (*UniswapV2Router02Transactor, error) {
	contract, err := bindUniswapV2Router02(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV2Router02Transactor{contract: contract}, nil
}

// NewUniswapV2Router02Filterer creates a new log filterer instance of UniswapV2Router02, bound to a specific deployed contract.
func NewUniswapV2Router02Filterer(address common.Address, filterer bind.ContractFilterer) (*UniswapV2Router02Filterer, error) {
	contract, err := bindUniswapV2Router02(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapV2Router02Filterer{contract: contract}, nil
}

// bindUniswapV2Router02 binds a generic wrapper to an already deployed contract.
func bindUniswapV2Router02(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UniswapV2Router02MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV2Router02 *UniswapV2Router02Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV2Router02.Contract.UniswapV2Router02Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV2Router02 *UniswapV2Router02Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV2Router02.Contract.UniswapV2Router02Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV2Router02 *UniswapV2Router02Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV2Router02.Contract.UniswapV2Router02Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV2Router02 *UniswapV2Router02CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV2Router02.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV2Router02 *UniswapV2Router02TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV2Router02.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV2Router02 *UniswapV2Router02TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV2Router02.Contract.contract.Transact(opts, method, params...)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_UniswapV2Router02 *UniswapV2Router02Caller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV2Router02.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_UniswapV2Router02 *UniswapV2Router02Session) Factory() (common.Address, error) {
	return _UniswapV2Router02.Contract.Factory(&_UniswapV2Router02.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_UniswapV2Router02 *UniswapV2Router02CallerSession) Factory() (common.Address, error) {
	return _UniswapV2Router02.Contract.Factory(&_UniswapV2Router02.CallOpts)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_UniswapV2Router02 *UniswapV2Router02Caller) GetAmountsIn(opts *bind.CallOpts, amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _UniswapV2Router02.contract.Call(opts, &out, "getAmountsIn", amountOut, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_UniswapV2Router02 *UniswapV2Router02Session) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _UniswapV2Router02.Contract.GetAmountsIn(&_UniswapV2Router02.CallOpts, amountOut, path)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_UniswapV2Router02 *UniswapV2Router02CallerSession) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _UniswapV2Router02.Contract.GetAmountsIn(&_UniswapV2Router02.CallOpts, amountOut, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_UniswapV2Router02 *UniswapV2Router02Caller) GetAmountsOut(opts *bind.CallOpts, amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _UniswapV2Router02.contract.Call(opts, &out, "getAmountsOut", amountIn, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_UniswapV2Router02 *UniswapV2Router02Session) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _UniswapV2Router02.Contract.GetAmountsOut(&_UniswapV2Router02.CallOpts, amountIn, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_UniswapV2Router02 *UniswapV2Router02CallerSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _UniswapV2Router02.Contract.GetAmountsOut(&_UniswapV2Router02.CallOpts, amountIn, path)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_UniswapV2Router02 *UniswapV2Router02Caller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV2Router02.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_UniswapV2Router02 *UniswapV2Router02Session) Weth() (common.Address, error) {
	return _UniswapV2Router02.Contract.Weth(&_UniswapV2Router02.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_UniswapV2Router02 *UniswapV2Router02CallerSession) Weth() (common.Address, error) {
	return _UniswapV2Router02.Contract.Weth(&_UniswapV2Router02.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_UniswapV2Router02 *UniswapV2Router02Transactor) AddLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Router02.contract.Transact(opts, "addLiquidity", tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_UniswapV2Router02 *UniswapV2Router02Session) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Router02.Contract.AddLiquidity(&_UniswapV2Router02.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_UniswapV2Router02 *UniswapV2Router02TransactorSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Router02.Contract.AddLiquidity(&_UniswapV2Router02.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_UniswapV2Router02 *UniswapV2Router02Transactor) AddLiquidityETH(opts *bind.TransactOpts, token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Router02.contract.Transact(opts, "addLiquidityETH", token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_UniswapV2Router02 *UniswapV2Router02Session) AddLiquidityETH(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Router02.Contract.AddLiquidityETH(&_UniswapV2Router02.TransactOpts, token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_UniswapV2Router02 *UniswapV2Router02TransactorSession) AddLiquidityETH(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Router02.Contract.AddLiquidityETH(&_UniswapV2Router02.TransactOpts, token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_UniswapV2Router02 *UniswapV2Router02Transactor) RemoveLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Router02.contract.Transact(opts, "removeLiquidity", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_UniswapV2Router02 *UniswapV2Router02Session) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Router02.Contract.RemoveLiquidity(&_UniswapV2Router02.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_UniswapV2Router02 *UniswapV2Router02TransactorSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Router02.Contract.RemoveLiquidity(&_UniswapV2Router02.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_UniswapV2Router02 *UniswapV2Router02Transactor) RemoveLiquidityETH(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Router02.contract.Transact(opts, "removeLiquidityETH", token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_UniswapV2Router02 *UniswapV2Router02Session) RemoveLiquidityETH(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Router02.Contract.RemoveLiquidityETH(&_UniswapV2Router02.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_UniswapV2Router02 *UniswapV2Router02TransactorSession) RemoveLiquidityETH(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Router02.Contract.RemoveLiquidityETH(&_UniswapV2Router02.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_UniswapV2Router02 *UniswapV2Router02Transactor) SwapExactETHForTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Router02.contract.Transact(opts, "swapExactETHForTokens", amountOutMin, path, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_UniswapV2Router02 *UniswapV2Router02Session) SwapExactETHForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Router02.Contract.SwapExactETHForTokens(&_UniswapV2Router02.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_UniswapV2Router02 *UniswapV2Router02TransactorSession) SwapExactETHForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Router02.Contract.SwapExactETHForTokens(&_UniswapV2Router02.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb6f9de95.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_UniswapV2Router02 *UniswapV2Router02Transactor) SwapExactETHForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Router02.contract.Transact(opts, "swapExactETHForTokensSupportingFeeOnTransferTokens", amountOutMin, path, to, deadline)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb6f9de95.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_UniswapV2Router02 *UniswapV2Router02Session) SwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Router02.Contract.SwapExactETHForTokensSupportingFeeOnTransferTokens(&_UniswapV2Router02.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb6f9de95.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_UniswapV2Router02 *UniswapV2Router02TransactorSession) SwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Router02.Contract.SwapExactETHForTokensSupportingFeeOnTransferTokens(&_UniswapV2Router02.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_UniswapV2Router02 *UniswapV2Router02Transactor) SwapExactTokensForETH(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Router02.contract.Transact(opts, "swapExactTokensForETH", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_UniswapV2Router02 *UniswapV2Router02Session) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Router02.Contract.SwapExactTokensForETH(&_UniswapV2Router02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_UniswapV2Router02 *UniswapV2Router02TransactorSession) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Router02.Contract.SwapExactTokensForETH(&_UniswapV2Router02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_UniswapV2Router02 *UniswapV2Router02Transactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Router02.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_UniswapV2Router02 *UniswapV2Router02Session) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Router02.Contract.SwapExactTokensForTokens(&_UniswapV2Router02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_UniswapV2Router02 *UniswapV2Router02TransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Router02.Contract.SwapExactTokensForTokens(&_UniswapV2Router02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_UniswapV2Router02 *UniswapV2Router02Transactor) SwapExactTokensForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Router02.contract.Transact(opts, "swapExactTokensForTokensSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_UniswapV2Router02 *UniswapV2Router02Session) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Router02.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_UniswapV2Router02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_UniswapV2Router02 *UniswapV2Router02TransactorSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Router02.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_UniswapV2Router02.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_UniswapV2Router02 *UniswapV2Router02Transactor) SwapTokensForExactETH(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Router02.contract.Transact(opts, "swapTokensForExactETH", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_UniswapV2Router02 *UniswapV2Router02Session) SwapTokensForExactETH(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Router02.Contract.SwapTokensForExactETH(&_UniswapV2Router02.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_UniswapV2Router02 *UniswapV2Router02TransactorSession) SwapTokensForExactETH(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Router02.Contract.SwapTokensForExactETH(&_UniswapV2Router02.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_UniswapV2Router02 *UniswapV2Router02Transactor) SwapTokensForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Router02.contract.Transact(opts, "swapTokensForExactTokens", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_UniswapV2Router02 *UniswapV2Router02Session) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Router02.Contract.SwapTokensForExactTokens(&_UniswapV2Router02.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_UniswapV2Router02 *UniswapV2Router02TransactorSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _UniswapV2Router02.Contract.SwapTokensForExactTokens(&_UniswapV2Router02.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniswapV2Router02 *UniswapV2Router02Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV2Router02.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniswapV2Router02 *UniswapV2Router02Session) Receive() (*types.Transaction, error) {
	return _UniswapV2Router02.Contract.Receive(&_UniswapV2Router02.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniswapV2Router02 *UniswapV2Router02TransactorSession) Receive() (*types.Transaction, error) {
	return _UniswapV2Router02.Contract.Receive(&_UniswapV2Router02.TransactOpts)
}
