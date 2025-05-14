package evm

import (
	"encoding/hex"
	"fmt"
	"math/big"

	ethabi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
)

func AbiUnpack(data []byte, types ...string) ([]interface{}, error) {
	var arguments ethabi.Arguments
	for _, t := range types {
		typ, err := ethabi.NewType(t, t, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create type: %w", err)
		}
		arguments = append(arguments, ethabi.Argument{Type: typ})
	}
	args, err := arguments.Unpack(data)
	if err != nil {
		return nil, fmt.Errorf("failed to get arguments: %w", err)
	}
	return args, nil
}

func AbiUnpackIntoMap(v map[string]interface{}, data []byte, types ...byte) error {
	var arguments ethabi.Arguments
	for _, t := range types {
		arguments = append(arguments, ethabi.Argument{Type: ethabi.Type{T: t}})
	}
	err := arguments.UnpackIntoMap(v, data)
	if err != nil {
		return fmt.Errorf("failed to get arguments: %w", err)
	}
	return nil
}

func DecodeExecuteData(executeData string) (*DecodedExecuteData, error) {
	executeDataBytes, err := hex.DecodeString(executeData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode execute data: %w", err)
	}

	//First 4 bytes are the function selector
	input, err := AbiUnpack(executeDataBytes[4:], "bytes")
	if err != nil {
		log.Debug().Msgf("[EvmClient] [DecodeExecuteData] unpack executeData error: %v", err)
	}
	return DecodeInput(input[0].([]byte))
}

func DecodeInput(input []byte) (*DecodedExecuteData, error) {
	args, err := AbiUnpack(input, "bytes", "bytes")
	if err != nil {
		return nil, fmt.Errorf("failed to unpack execute input: %w", err)
	}
	// log.Debug().
	// 	Str("data", hex.EncodeToString(args[0].([]byte))).
	// 	Str("proof", hex.EncodeToString(args[1].([]byte))).
	// 	Msg("[EvmClient] [DecodeInput]")
	//Decode the data
	dataDecoded, err := AbiUnpack(args[0].([]byte), "uint256", "bytes32[]", "string[]", "bytes[]")
	if err != nil {
		return nil, fmt.Errorf("failed to unpack data: %w", err)
	}
	log.Debug().
		Uint64("chainId", dataDecoded[0].(*big.Int).Uint64()).
		Strs("commands", dataDecoded[2].([]string)).
		Int("NumberOfCommands", len(dataDecoded[2].([]string))).
		Msg("[EvmClient] [DecodeInput]")
	//Decode the proof
	proofDecoded, err := AbiUnpack(args[1].([]byte), "address[]", "uint256[]", "uint256", "bytes[]")
	if err != nil {
		return nil, fmt.Errorf("failed to unpack proof: %w", err)
	}
	chainId := dataDecoded[0].(*big.Int)
	commandIds := dataDecoded[1].([][32]byte)
	commands := dataDecoded[2].([]string)
	params := dataDecoded[3].([][]byte)
	weights := proofDecoded[1].([]*big.Int)
	weightsUint64 := make([]uint64, len(weights))
	for i, weight := range weights {
		weightsUint64[i] = weight.Uint64()
	}
	threshold := proofDecoded[2].(*big.Int)
	signaturesBytes := proofDecoded[3].([][]byte)
	signatures := make([]string, len(signaturesBytes))
	for i, signature := range signaturesBytes {
		signatures[i] = hex.EncodeToString(signature)
	}
	return &DecodedExecuteData{
		Input:      input,
		ChainId:    chainId.Uint64(),
		CommandIds: commandIds,
		Commands:   commands,
		Params:     params,
		Operators:  proofDecoded[0].([]common.Address),
		Weights:    weightsUint64,
		Threshold:  threshold.Uint64(),
		Signatures: signatures,
	}, nil
}

func DecodeApproveContractCall(input []byte) (*ApproveContractCall, error) {
	dataDecoded, err := AbiUnpack(input, "uint256", "bytes32[]", "string[]", "bytes[]")
	if err != nil {
		return nil, fmt.Errorf("failed to unpack data: %w", err)
	}
	chainId := dataDecoded[0].(*big.Int)
	commandIds := dataDecoded[1].([][32]byte)
	commands := dataDecoded[2].([]string)
	params := dataDecoded[3].([][]byte)

	return &ApproveContractCall{
		ChainId:    chainId.Uint64(),
		CommandIds: commandIds,
		Commands:   commands,
		Params:     params,
	}, nil
}
func DecodeDeployToken(input []byte) (*DeployToken, error) {
	dataDecoded, err := AbiUnpack(input, "string", "string", "uint8", "uint256", "address", "uint256")
	if err != nil {
		return nil, fmt.Errorf("failed to unpack data: %w", err)
	}
	decimals := dataDecoded[2].(uint8)
	cap := dataDecoded[3].(*big.Int)
	mintLimit := dataDecoded[5].(*big.Int)
	deployToken := DeployToken{
		Name:         dataDecoded[0].(string),
		Symbol:       dataDecoded[1].(string),
		Decimals:     decimals,
		Cap:          cap.Uint64(),
		TokenAddress: dataDecoded[4].(common.Address),
		MintLimit:    mintLimit.Uint64(),
	}
	log.Debug().
		Any("DeployToken", deployToken).
		Msg("[EvmClient] [DecodeDeployToken]")
	return &deployToken, nil
}

func DecodeStartedSwitchPhase(input []byte) (*RedeemPhase, error) {
	dataDecoded, err := AbiUnpack(input, "uint64", "uint8")
	if err != nil {
		return nil, fmt.Errorf("failed to unpack data: %w", err)
	}
	sessionSequence := dataDecoded[0].(*big.Int)
	phase := dataDecoded[1].(uint8)
	return &RedeemPhase{
		Sequence: sessionSequence.Uint64(),
		Phase:    phase,
	}, nil
}

func DecodeGroupUid(groupHex string) ([32]byte, error) {
	groupBytes, err := hex.DecodeString(groupHex)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to decode group uid: %w", err)
	}
	groupBytes32 := [32]byte{}
	copy(groupBytes32[:], groupBytes)
	return groupBytes32, nil
}
