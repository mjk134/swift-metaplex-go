// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package metaplex

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Given an existing store, add or update an existing whitelisted creator for the store. This creates
// a PDA with seed ['metaplex', store key, creator key] if it does not already exist to store attributes there.
type SetWhitelistedCreator struct {
	Args *SetWhitelistedCreatorArgs

	// [0] = [WRITE] whitelistedCreatorPDAKey
	// ··········· The whitelisted creator pda key, seed of ['metaplex', store key, creator key]
	//
	// [1] = [SIGNER] adminWallet
	// ··········· The admin wallet
	//
	// [2] = [SIGNER] payer
	// ··········· Payer
	//
	// [3] = [] creatorKey
	// ··········· The creator key
	//
	// [4] = [] storeKey
	// ··········· The store key, seed of ['metaplex', admin wallet]
	//
	// [5] = [] system
	// ··········· System
	//
	// [6] = [] rentSysvar
	// ··········· Rent sysvar
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSetWhitelistedCreatorInstructionBuilder creates a new `SetWhitelistedCreator` instruction builder.
func NewSetWhitelistedCreatorInstructionBuilder() *SetWhitelistedCreator {
	nd := &SetWhitelistedCreator{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 7),
	}
	return nd
}

// SetArgs sets the "args" parameter.
func (inst *SetWhitelistedCreator) SetArgs(args SetWhitelistedCreatorArgs) *SetWhitelistedCreator {
	inst.Args = &args
	return inst
}

// SetWhitelistedCreatorPDAKeyAccount sets the "whitelistedCreatorPDAKey" account.
// The whitelisted creator pda key, seed of ['metaplex', store key, creator key]
func (inst *SetWhitelistedCreator) SetWhitelistedCreatorPDAKeyAccount(whitelistedCreatorPDAKey ag_solanago.PublicKey) *SetWhitelistedCreator {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(whitelistedCreatorPDAKey).WRITE()
	return inst
}

// GetWhitelistedCreatorPDAKeyAccount gets the "whitelistedCreatorPDAKey" account.
// The whitelisted creator pda key, seed of ['metaplex', store key, creator key]
func (inst *SetWhitelistedCreator) GetWhitelistedCreatorPDAKeyAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAdminWalletAccount sets the "adminWallet" account.
// The admin wallet
func (inst *SetWhitelistedCreator) SetAdminWalletAccount(adminWallet ag_solanago.PublicKey) *SetWhitelistedCreator {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(adminWallet).SIGNER()
	return inst
}

// GetAdminWalletAccount gets the "adminWallet" account.
// The admin wallet
func (inst *SetWhitelistedCreator) GetAdminWalletAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPayerAccount sets the "payer" account.
// Payer
func (inst *SetWhitelistedCreator) SetPayerAccount(payer ag_solanago.PublicKey) *SetWhitelistedCreator {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(payer).SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
// Payer
func (inst *SetWhitelistedCreator) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetCreatorKeyAccount sets the "creatorKey" account.
// The creator key
func (inst *SetWhitelistedCreator) SetCreatorKeyAccount(creatorKey ag_solanago.PublicKey) *SetWhitelistedCreator {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(creatorKey)
	return inst
}

// GetCreatorKeyAccount gets the "creatorKey" account.
// The creator key
func (inst *SetWhitelistedCreator) GetCreatorKeyAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetStoreKeyAccount sets the "storeKey" account.
// The store key, seed of ['metaplex', admin wallet]
func (inst *SetWhitelistedCreator) SetStoreKeyAccount(storeKey ag_solanago.PublicKey) *SetWhitelistedCreator {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(storeKey)
	return inst
}

// GetStoreKeyAccount gets the "storeKey" account.
// The store key, seed of ['metaplex', admin wallet]
func (inst *SetWhitelistedCreator) GetStoreKeyAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetSystemAccount sets the "system" account.
// System
func (inst *SetWhitelistedCreator) SetSystemAccount(system ag_solanago.PublicKey) *SetWhitelistedCreator {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(system)
	return inst
}

// GetSystemAccount gets the "system" account.
// System
func (inst *SetWhitelistedCreator) GetSystemAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetRentSysvarAccount sets the "rentSysvar" account.
// Rent sysvar
func (inst *SetWhitelistedCreator) SetRentSysvarAccount(rentSysvar ag_solanago.PublicKey) *SetWhitelistedCreator {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(rentSysvar)
	return inst
}

// GetRentSysvarAccount gets the "rentSysvar" account.
// Rent sysvar
func (inst *SetWhitelistedCreator) GetRentSysvarAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

func (inst SetWhitelistedCreator) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_SetWhitelistedCreator),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SetWhitelistedCreator) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetWhitelistedCreator) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Args == nil {
			return errors.New("Args parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.WhitelistedCreatorPDAKey is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.AdminWallet is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.CreatorKey is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.StoreKey is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.System is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.RentSysvar is not set")
		}
	}
	return nil
}

func (inst *SetWhitelistedCreator) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetWhitelistedCreator")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Args", *inst.Args))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=7]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("whitelistedCreatorPDAKey", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("             adminWallet", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                   payer", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("              creatorKey", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("                storeKey", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("                  system", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("              rentSysvar", inst.AccountMetaSlice.Get(6)))
					})
				})
		})
}

func (obj SetWhitelistedCreator) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Args` param:
	err = encoder.Encode(obj.Args)
	if err != nil {
		return err
	}
	return nil
}
func (obj *SetWhitelistedCreator) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Args`:
	err = decoder.Decode(&obj.Args)
	if err != nil {
		return err
	}
	return nil
}

// NewSetWhitelistedCreatorInstruction declares a new SetWhitelistedCreator instruction with the provided parameters and accounts.
func NewSetWhitelistedCreatorInstruction(
	// Parameters:
	args SetWhitelistedCreatorArgs,
	// Accounts:
	whitelistedCreatorPDAKey ag_solanago.PublicKey,
	adminWallet ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	creatorKey ag_solanago.PublicKey,
	storeKey ag_solanago.PublicKey,
	system ag_solanago.PublicKey,
	rentSysvar ag_solanago.PublicKey) *SetWhitelistedCreator {
	return NewSetWhitelistedCreatorInstructionBuilder().
		SetArgs(args).
		SetWhitelistedCreatorPDAKeyAccount(whitelistedCreatorPDAKey).
		SetAdminWalletAccount(adminWallet).
		SetPayerAccount(payer).
		SetCreatorKeyAccount(creatorKey).
		SetStoreKeyAccount(storeKey).
		SetSystemAccount(system).
		SetRentSysvarAccount(rentSysvar)
}
