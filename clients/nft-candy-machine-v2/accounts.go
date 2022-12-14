// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package nft_candy_machine_v2

import (
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type CandyMachine struct {
	Authority     ag_solanago.PublicKey
	Wallet        ag_solanago.PublicKey
	TokenMint     *ag_solanago.PublicKey `bin:"optional"`
	ItemsRedeemed uint64
	Data          CandyMachineData
}

type MagicEden struct {
	Authority     ag_solanago.PublicKey
	WalletAuthority        ag_solanago.PublicKey
	Config ag_solanago.PublicKey
	ItemsRedeemedNormal uint64
	ItemsRedeemedRaffle uint64
	RaffleTicketsPurchased uint64
	Uuid string
	ItemsAvailable uint64
	RaffleSeed uint64
	Bump uint8
	Notary ag_solanago.PublicKey
	OrderInfo ag_solanago.PublicKey
}

var CandyMachineDiscriminator = [8]byte{51, 173, 177, 113, 25, 241, 109, 189}

func (obj CandyMachine) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(CandyMachineDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Authority` param:
	err = encoder.Encode(obj.Authority)
	if err != nil {
		return err
	}
	// Serialize `Wallet` param:
	err = encoder.Encode(obj.Wallet)
	if err != nil {
		return err
	}
	// Serialize `TokenMint` param (optional):
	{
		if obj.TokenMint == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.TokenMint)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `ItemsRedeemed` param:
	err = encoder.Encode(obj.ItemsRedeemed)
	if err != nil {
		return err
	}
	// Serialize `Data` param:
	err = encoder.Encode(obj.Data)
	if err != nil {
		return err
	}
	return nil
}

func (obj *CandyMachine) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(CandyMachineDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[51 173 177 113 25 241 109 189]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Authority`:
	err = decoder.Decode(&obj.Authority)
	if err != nil {
		return err
	}
	// Deserialize `Wallet`:
	err = decoder.Decode(&obj.Wallet)
	if err != nil {
		return err
	}
	// Deserialize `TokenMint` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.TokenMint)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `ItemsRedeemed`:
	err = decoder.Decode(&obj.ItemsRedeemed)
	if err != nil {
		return err
	}
	// Deserialize `Data`:
	err = decoder.Decode(&obj.Data)
	if err != nil {
		return err
	}
	return nil
}
