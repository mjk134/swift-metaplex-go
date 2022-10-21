// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package auction

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Update the authority for an auction account.
type SetAuthority struct {

	// [0] = [WRITE] auction
	// ··········· auction (pda of ['auction', program id, resource id])
	//
	// [1] = [SIGNER] authority
	// ··········· authority
	//
	// [2] = [] newAuthority
	// ··········· newAuthority
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSetAuthorityInstructionBuilder creates a new `SetAuthority` instruction builder.
func NewSetAuthorityInstructionBuilder() *SetAuthority {
	nd := &SetAuthority{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetAuctionAccount sets the "auction" account.
// auction (pda of ['auction', program id, resource id])
func (inst *SetAuthority) SetAuctionAccount(auction ag_solanago.PublicKey) *SetAuthority {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(auction).WRITE()
	return inst
}

// GetAuctionAccount gets the "auction" account.
// auction (pda of ['auction', program id, resource id])
func (inst *SetAuthority) GetAuctionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetAuthorityAccount sets the "authority" account.
// authority
func (inst *SetAuthority) SetAuthorityAccount(authority ag_solanago.PublicKey) *SetAuthority {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
// authority
func (inst *SetAuthority) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetNewAuthorityAccount sets the "newAuthority" account.
// newAuthority
func (inst *SetAuthority) SetNewAuthorityAccount(newAuthority ag_solanago.PublicKey) *SetAuthority {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(newAuthority)
	return inst
}

// GetNewAuthorityAccount gets the "newAuthority" account.
// newAuthority
func (inst *SetAuthority) GetNewAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

func (inst SetAuthority) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_SetAuthority),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SetAuthority) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetAuthority) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Auction is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.NewAuthority is not set")
		}
	}
	return nil
}

func (inst *SetAuthority) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetAuthority")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("     auction", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("   authority", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("newAuthority", inst.AccountMetaSlice[2]))
					})
				})
		})
}

func (obj SetAuthority) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *SetAuthority) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewSetAuthorityInstruction declares a new SetAuthority instruction with the provided parameters and accounts.
func NewSetAuthorityInstruction(
	// Accounts:
	auction ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	newAuthority ag_solanago.PublicKey) *SetAuthority {
	return NewSetAuthorityInstructionBuilder().
		SetAuctionAccount(auction).
		SetAuthorityAccount(authority).
		SetNewAuthorityAccount(newAuthority)
}
