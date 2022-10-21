// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_vault

import ag_binary "github.com/gagliardetto/binary"

type InitVaultArgs struct {
	AllowFurtherShareCreation bool
}

func (obj InitVaultArgs) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `AllowFurtherShareCreation` param:
	err = encoder.Encode(obj.AllowFurtherShareCreation)
	if err != nil {
		return err
	}
	return nil
}

func (obj *InitVaultArgs) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `AllowFurtherShareCreation`:
	err = decoder.Decode(&obj.AllowFurtherShareCreation)
	if err != nil {
		return err
	}
	return nil
}

type AmountArgs struct {
	Amount uint64
}

func (obj AmountArgs) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (obj *AmountArgs) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

type NumberOfShareArgs struct {
	NumberOfShares uint64
}

func (obj NumberOfShareArgs) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `NumberOfShares` param:
	err = encoder.Encode(obj.NumberOfShares)
	if err != nil {
		return err
	}
	return nil
}

func (obj *NumberOfShareArgs) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `NumberOfShares`:
	err = decoder.Decode(&obj.NumberOfShares)
	if err != nil {
		return err
	}
	return nil
}

type MintEditionProxyArgs struct {
	Edition uint64
}

func (obj MintEditionProxyArgs) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Edition` param:
	err = encoder.Encode(obj.Edition)
	if err != nil {
		return err
	}
	return nil
}

func (obj *MintEditionProxyArgs) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Edition`:
	err = decoder.Decode(&obj.Edition)
	if err != nil {
		return err
	}
	return nil
}