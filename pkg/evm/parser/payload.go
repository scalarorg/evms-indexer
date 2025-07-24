package parser

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcutil/psbt"
	"github.com/scalarorg/bitcoin-vault/go-utils/btc"
	"github.com/scalarorg/evms-indexer/pkg/evm/abi"
)

type ContractCallWithTokenPayloadType uint8

const (
	ContractCallWithTokenPayloadType_CustodianOnly ContractCallWithTokenPayloadType = iota
	ContractCallWithTokenPayloadType_UPC
)

func (c ContractCallWithTokenPayloadType) Byte() byte {
	return byte(c)
}

func FromByte(b byte) ContractCallWithTokenPayloadType {
	return ContractCallWithTokenPayloadType(b)
}

type CustodianOnly struct {
	FeeOptions               BTCFeeOpts
	RBF                      bool
	RecipientChainIdentifier []byte
	Amount                   uint64
	TxIds                    []string
	Vouts                    []uint32
	Amounts                  []uint64
	ReqId                    [32]byte
}

type UPC struct {
	Psbt []byte
}

type ContractCallWithTokenPayload struct {
	PayloadType ContractCallWithTokenPayloadType
	*CustodianOnly
	*UPC
}

func (p *ContractCallWithTokenPayload) Parse(payload []byte) error {
	payloadType := FromByte(payload[0])
	if payloadType == ContractCallWithTokenPayloadType_CustodianOnly {
		p.CustodianOnly = &CustodianOnly{}
		return p.CustodianOnly.Parse(payload[1:])
	} else if payloadType == ContractCallWithTokenPayloadType_UPC {
		p.UPC = &UPC{}
		return p.UPC.Parse(payload[1:])
	}
	return fmt.Errorf("invalid payload type")
}

func (p *CustodianOnly) Parse(payload []byte) error {
	decoded, err := abi.ContractCallWithTokenCustodianOnly.Unpack(payload)
	if err != nil {
		return err
	}

	lockingScript := decoded[1].([]byte)
	amount := decoded[0].(uint64)
	txIds := decoded[2].([]string)
	vouts := decoded[3].([]uint32)
	amounts := decoded[4].([]uint64)
	reqId := decoded[5].([32]byte)
	p.RecipientChainIdentifier = lockingScript
	p.Amount = amount
	p.TxIds = txIds
	p.Vouts = vouts
	p.Amounts = amounts
	p.ReqId = reqId
	return nil
}

func (p *UPC) Parse(payload []byte) error {
	// decoded, err := abi.ContractCallWithTokenUPC.Unpack(payload)
	// if err != nil {
	// 	return err
	// }
	// p.Psbt = decoded[0].([]byte)
	p.Psbt = payload
	return nil
}

func (p *ContractCallWithTokenPayload) GetDestinationAddress(chainId uint64) (string, error) {
	params := btc.BtcChainsRecords().GetChainParamsByID(chainId)

	if params == nil {
		return "", fmt.Errorf("invalid destination chain: %d", chainId)
	}
	if p.CustodianOnly != nil {
		identifier := p.CustodianOnly.RecipientChainIdentifier
		address, err := btc.ScriptPubKeyToAddress(identifier, params.Name)
		if err != nil {
			return "", fmt.Errorf("failed to convert script pubkey %s to address with params name %s: %w",
				hex.EncodeToString(identifier), params.Name, err)
		}
		return address.EncodeAddress(), nil
	} else if p.UPC != nil && len(p.UPC.Psbt) > 0 {
		packet, err := psbt.NewFromRawBytes(
			bytes.NewReader(p.UPC.Psbt), false,
		)

		if err != nil {
			return "", fmt.Errorf("failed to create psbt packet: %w", err)
		}

		identifier := packet.UnsignedTx.TxOut[1].PkScript
		address, err := btc.ScriptPubKeyToAddress(identifier, params.Name)
		if err != nil {
			return "", fmt.Errorf("failed to convert script pubkey to address: %w", err)
		}
		return address.EncodeAddress(), nil
	}
	return "", fmt.Errorf("invalid payload")
}
