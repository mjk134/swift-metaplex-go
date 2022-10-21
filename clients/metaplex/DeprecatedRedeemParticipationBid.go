// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package metaplex

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Note: This requires that auction manager be in a Running state.
//
// If an auction is complete, you can redeem your bid for an Open Edition token if it is eligible. If you are the first to do this,
// The auction manager will switch from Running state to Disbursing state. If you are the last, this may change
// the auction manager state to Finished provided that no authorities remain to be delegated for Master Edition tokens.
//
// NOTE: Please note that it is totally possible to redeem a bid 2x - once for a prize you won and once at this end point for a open edition
// that comes as a 'token of appreciation' for bidding. They are not mutually exclusive unless explicitly set to be that way.
//
// NOTE: If you are redeeming a newly minted Open Edition, you must actually supply a destination account containing a token from a brand new
// mint. We do not provide the token to you. Our job with this action is to christen this mint + token combo as an official Open Edition.
type DeprecatedRedeemParticipationBid struct {

	// [0] = [WRITE] auctionManager
	// ··········· Auction manager
	//
	// [1] = [WRITE] safetyDepositTokenStorage
	// ··········· Safety deposit token storage account
	//
	// [2] = [WRITE] destinationAccount
	// ··········· Destination account for limited edition authority token. Must be same mint as master edition Printing mint.
	//
	// [3] = [WRITE] bidRedemptionKey
	// ··········· Bid redemption key -
	// ··········· Just a PDA with seed ['metaplex', auction_key, bidder_metadata_key] that we will allocate to mark that you redeemed your bid
	//
	// [4] = [] safetyDepositBox
	// ··········· Safety deposit box account
	//
	// [5] = [] vaultAccount
	// ··········· Vault account
	//
	// [6] = [] safetyDepositConfig
	// ··········· Safety deposit config pda of ['metaplex', program id, auction manager, safety deposit]
	// ··········· This account will only get used in the event this is an AuctionManagerV2
	//
	// [7] = [] auction
	// ··········· Auction
	//
	// [8] = [] bidderMetadata
	// ··········· Your BidderMetadata account
	//
	// [9] = [WRITE, SIGNER] bidder
	// ··········· Your Bidder account - Only needs to be signer if payer does not own
	//
	// [10] = [SIGNER] payer
	// ··········· Payer
	//
	// [11] = [] tokenProgram
	// ··········· Token program
	//
	// [12] = [] tokenVaultProgram
	// ··········· Token Vault program
	//
	// [13] = [] tokenMetadataProgram
	// ··········· Token metadata program
	//
	// [14] = [] store
	// ··········· Store
	//
	// [15] = [] system
	// ··········· System
	//
	// [16] = [] rentSysvar
	// ··········· Rent sysvar
	//
	// [17] = [SIGNER] transferAuthority
	// ··········· Transfer authority to move the payment in the auction's token_mint coin from the bidder account for the participation_fixed_price
	// ··········· on the auction manager to the auction manager account itself.
	//
	// [18] = [WRITE] acceptPayment
	// ··········· The accept payment account for the auction manager
	//
	// [19] = [WRITE] tokenAccount
	// ··········· The token account you will potentially pay for the open edition bid with if necessary
	//
	// [20] = [WRITE] participationNFTPrintingHoldingAccount
	// ··········· Participation NFT printing holding account (present on participation_state)
	//
	// [21] = [] auctionExtendedPDA
	// ··········· Auction extended (pda relative to auction of ['auction', program id, vault key, 'extended'])
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewDeprecatedRedeemParticipationBidInstructionBuilder creates a new `DeprecatedRedeemParticipationBid` instruction builder.
func NewDeprecatedRedeemParticipationBidInstructionBuilder() *DeprecatedRedeemParticipationBid {
	nd := &DeprecatedRedeemParticipationBid{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 22),
	}
	return nd
}

// SetAuctionManagerAccount sets the "auctionManager" account.
// Auction manager
func (inst *DeprecatedRedeemParticipationBid) SetAuctionManagerAccount(auctionManager ag_solanago.PublicKey) *DeprecatedRedeemParticipationBid {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(auctionManager).WRITE()
	return inst
}

// GetAuctionManagerAccount gets the "auctionManager" account.
// Auction manager
func (inst *DeprecatedRedeemParticipationBid) GetAuctionManagerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetSafetyDepositTokenStorageAccount sets the "safetyDepositTokenStorage" account.
// Safety deposit token storage account
func (inst *DeprecatedRedeemParticipationBid) SetSafetyDepositTokenStorageAccount(safetyDepositTokenStorage ag_solanago.PublicKey) *DeprecatedRedeemParticipationBid {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(safetyDepositTokenStorage).WRITE()
	return inst
}

// GetSafetyDepositTokenStorageAccount gets the "safetyDepositTokenStorage" account.
// Safety deposit token storage account
func (inst *DeprecatedRedeemParticipationBid) GetSafetyDepositTokenStorageAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetDestinationAccount sets the "destinationAccount" account.
// Destination account for limited edition authority token. Must be same mint as master edition Printing mint.
func (inst *DeprecatedRedeemParticipationBid) SetDestinationAccount(destinationAccount ag_solanago.PublicKey) *DeprecatedRedeemParticipationBid {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(destinationAccount).WRITE()
	return inst
}

// GetDestinationAccount gets the "destinationAccount" account.
// Destination account for limited edition authority token. Must be same mint as master edition Printing mint.
func (inst *DeprecatedRedeemParticipationBid) GetDestinationAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetBidRedemptionKeyAccount sets the "bidRedemptionKey" account.
// Bid redemption key -
// Just a PDA with seed ['metaplex', auction_key, bidder_metadata_key] that we will allocate to mark that you redeemed your bid
func (inst *DeprecatedRedeemParticipationBid) SetBidRedemptionKeyAccount(bidRedemptionKey ag_solanago.PublicKey) *DeprecatedRedeemParticipationBid {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(bidRedemptionKey).WRITE()
	return inst
}

// GetBidRedemptionKeyAccount gets the "bidRedemptionKey" account.
// Bid redemption key -
// Just a PDA with seed ['metaplex', auction_key, bidder_metadata_key] that we will allocate to mark that you redeemed your bid
func (inst *DeprecatedRedeemParticipationBid) GetBidRedemptionKeyAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSafetyDepositBoxAccount sets the "safetyDepositBox" account.
// Safety deposit box account
func (inst *DeprecatedRedeemParticipationBid) SetSafetyDepositBoxAccount(safetyDepositBox ag_solanago.PublicKey) *DeprecatedRedeemParticipationBid {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(safetyDepositBox)
	return inst
}

// GetSafetyDepositBoxAccount gets the "safetyDepositBox" account.
// Safety deposit box account
func (inst *DeprecatedRedeemParticipationBid) GetSafetyDepositBoxAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetVaultAccount sets the "vaultAccount" account.
// Vault account
func (inst *DeprecatedRedeemParticipationBid) SetVaultAccount(vaultAccount ag_solanago.PublicKey) *DeprecatedRedeemParticipationBid {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(vaultAccount)
	return inst
}

// GetVaultAccount gets the "vaultAccount" account.
// Vault account
func (inst *DeprecatedRedeemParticipationBid) GetVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetSafetyDepositConfigAccount sets the "safetyDepositConfig" account.
// Safety deposit config pda of ['metaplex', program id, auction manager, safety deposit]
// This account will only get used in the event this is an AuctionManagerV2
func (inst *DeprecatedRedeemParticipationBid) SetSafetyDepositConfigAccount(safetyDepositConfig ag_solanago.PublicKey) *DeprecatedRedeemParticipationBid {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(safetyDepositConfig)
	return inst
}

// GetSafetyDepositConfigAccount gets the "safetyDepositConfig" account.
// Safety deposit config pda of ['metaplex', program id, auction manager, safety deposit]
// This account will only get used in the event this is an AuctionManagerV2
func (inst *DeprecatedRedeemParticipationBid) GetSafetyDepositConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetAuctionAccount sets the "auction" account.
// Auction
func (inst *DeprecatedRedeemParticipationBid) SetAuctionAccount(auction ag_solanago.PublicKey) *DeprecatedRedeemParticipationBid {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(auction)
	return inst
}

// GetAuctionAccount gets the "auction" account.
// Auction
func (inst *DeprecatedRedeemParticipationBid) GetAuctionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetBidderMetadataAccount sets the "bidderMetadata" account.
// Your BidderMetadata account
func (inst *DeprecatedRedeemParticipationBid) SetBidderMetadataAccount(bidderMetadata ag_solanago.PublicKey) *DeprecatedRedeemParticipationBid {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(bidderMetadata)
	return inst
}

// GetBidderMetadataAccount gets the "bidderMetadata" account.
// Your BidderMetadata account
func (inst *DeprecatedRedeemParticipationBid) GetBidderMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetBidderAccount sets the "bidder" account.
// Your Bidder account - Only needs to be signer if payer does not own
func (inst *DeprecatedRedeemParticipationBid) SetBidderAccount(bidder ag_solanago.PublicKey) *DeprecatedRedeemParticipationBid {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(bidder).WRITE().SIGNER()
	return inst
}

// GetBidderAccount gets the "bidder" account (optional).
// Your Bidder account - Only needs to be signer if payer does not own
func (inst *DeprecatedRedeemParticipationBid) GetBidderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetPayerAccount sets the "payer" account.
// Payer
func (inst *DeprecatedRedeemParticipationBid) SetPayerAccount(payer ag_solanago.PublicKey) *DeprecatedRedeemParticipationBid {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(payer).SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
// Payer
func (inst *DeprecatedRedeemParticipationBid) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
// Token program
func (inst *DeprecatedRedeemParticipationBid) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *DeprecatedRedeemParticipationBid {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
// Token program
func (inst *DeprecatedRedeemParticipationBid) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetTokenVaultProgramAccount sets the "tokenVaultProgram" account.
// Token Vault program
func (inst *DeprecatedRedeemParticipationBid) SetTokenVaultProgramAccount(tokenVaultProgram ag_solanago.PublicKey) *DeprecatedRedeemParticipationBid {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(tokenVaultProgram)
	return inst
}

// GetTokenVaultProgramAccount gets the "tokenVaultProgram" account.
// Token Vault program
func (inst *DeprecatedRedeemParticipationBid) GetTokenVaultProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetTokenMetadataProgramAccount sets the "tokenMetadataProgram" account.
// Token metadata program
func (inst *DeprecatedRedeemParticipationBid) SetTokenMetadataProgramAccount(tokenMetadataProgram ag_solanago.PublicKey) *DeprecatedRedeemParticipationBid {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(tokenMetadataProgram)
	return inst
}

// GetTokenMetadataProgramAccount gets the "tokenMetadataProgram" account.
// Token metadata program
func (inst *DeprecatedRedeemParticipationBid) GetTokenMetadataProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetStoreAccount sets the "store" account.
// Store
func (inst *DeprecatedRedeemParticipationBid) SetStoreAccount(store ag_solanago.PublicKey) *DeprecatedRedeemParticipationBid {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(store)
	return inst
}

// GetStoreAccount gets the "store" account.
// Store
func (inst *DeprecatedRedeemParticipationBid) GetStoreAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

// SetSystemAccount sets the "system" account.
// System
func (inst *DeprecatedRedeemParticipationBid) SetSystemAccount(system ag_solanago.PublicKey) *DeprecatedRedeemParticipationBid {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(system)
	return inst
}

// GetSystemAccount gets the "system" account.
// System
func (inst *DeprecatedRedeemParticipationBid) GetSystemAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(15)
}

// SetRentSysvarAccount sets the "rentSysvar" account.
// Rent sysvar
func (inst *DeprecatedRedeemParticipationBid) SetRentSysvarAccount(rentSysvar ag_solanago.PublicKey) *DeprecatedRedeemParticipationBid {
	inst.AccountMetaSlice[16] = ag_solanago.Meta(rentSysvar)
	return inst
}

// GetRentSysvarAccount gets the "rentSysvar" account.
// Rent sysvar
func (inst *DeprecatedRedeemParticipationBid) GetRentSysvarAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(16)
}

// SetTransferAuthorityAccount sets the "transferAuthority" account.
// Transfer authority to move the payment in the auction's token_mint coin from the bidder account for the participation_fixed_price
// on the auction manager to the auction manager account itself.
func (inst *DeprecatedRedeemParticipationBid) SetTransferAuthorityAccount(transferAuthority ag_solanago.PublicKey) *DeprecatedRedeemParticipationBid {
	inst.AccountMetaSlice[17] = ag_solanago.Meta(transferAuthority).SIGNER()
	return inst
}

// GetTransferAuthorityAccount gets the "transferAuthority" account.
// Transfer authority to move the payment in the auction's token_mint coin from the bidder account for the participation_fixed_price
// on the auction manager to the auction manager account itself.
func (inst *DeprecatedRedeemParticipationBid) GetTransferAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(17)
}

// SetAcceptPaymentAccount sets the "acceptPayment" account.
// The accept payment account for the auction manager
func (inst *DeprecatedRedeemParticipationBid) SetAcceptPaymentAccount(acceptPayment ag_solanago.PublicKey) *DeprecatedRedeemParticipationBid {
	inst.AccountMetaSlice[18] = ag_solanago.Meta(acceptPayment).WRITE()
	return inst
}

// GetAcceptPaymentAccount gets the "acceptPayment" account.
// The accept payment account for the auction manager
func (inst *DeprecatedRedeemParticipationBid) GetAcceptPaymentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(18)
}

// SetTokenAccount sets the "tokenAccount" account.
// The token account you will potentially pay for the open edition bid with if necessary
func (inst *DeprecatedRedeemParticipationBid) SetTokenAccount(tokenAccount ag_solanago.PublicKey) *DeprecatedRedeemParticipationBid {
	inst.AccountMetaSlice[19] = ag_solanago.Meta(tokenAccount).WRITE()
	return inst
}

// GetTokenAccount gets the "tokenAccount" account.
// The token account you will potentially pay for the open edition bid with if necessary
func (inst *DeprecatedRedeemParticipationBid) GetTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(19)
}

// SetParticipationNFTPrintingHoldingAccount sets the "participationNFTPrintingHoldingAccount" account.
// Participation NFT printing holding account (present on participation_state)
func (inst *DeprecatedRedeemParticipationBid) SetParticipationNFTPrintingHoldingAccount(participationNFTPrintingHoldingAccount ag_solanago.PublicKey) *DeprecatedRedeemParticipationBid {
	inst.AccountMetaSlice[20] = ag_solanago.Meta(participationNFTPrintingHoldingAccount).WRITE()
	return inst
}

// GetParticipationNFTPrintingHoldingAccount gets the "participationNFTPrintingHoldingAccount" account.
// Participation NFT printing holding account (present on participation_state)
func (inst *DeprecatedRedeemParticipationBid) GetParticipationNFTPrintingHoldingAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(20)
}

// SetAuctionExtendedPDAAccount sets the "auctionExtendedPDA" account.
// Auction extended (pda relative to auction of ['auction', program id, vault key, 'extended'])
func (inst *DeprecatedRedeemParticipationBid) SetAuctionExtendedPDAAccount(auctionExtendedPDA ag_solanago.PublicKey) *DeprecatedRedeemParticipationBid {
	inst.AccountMetaSlice[21] = ag_solanago.Meta(auctionExtendedPDA)
	return inst
}

// GetAuctionExtendedPDAAccount gets the "auctionExtendedPDA" account.
// Auction extended (pda relative to auction of ['auction', program id, vault key, 'extended'])
func (inst *DeprecatedRedeemParticipationBid) GetAuctionExtendedPDAAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(21)
}

func (inst DeprecatedRedeemParticipationBid) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_DeprecatedRedeemParticipationBid),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst DeprecatedRedeemParticipationBid) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *DeprecatedRedeemParticipationBid) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.AuctionManager is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.SafetyDepositTokenStorage is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.DestinationAccount is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.BidRedemptionKey is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.SafetyDepositBox is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.VaultAccount is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.SafetyDepositConfig is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.Auction is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.BidderMetadata is not set")
		}

		// [9] = Bidder is optional

		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.TokenVaultProgram is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.TokenMetadataProgram is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.Store is not set")
		}
		if inst.AccountMetaSlice[15] == nil {
			return errors.New("accounts.System is not set")
		}
		if inst.AccountMetaSlice[16] == nil {
			return errors.New("accounts.RentSysvar is not set")
		}
		if inst.AccountMetaSlice[17] == nil {
			return errors.New("accounts.TransferAuthority is not set")
		}
		if inst.AccountMetaSlice[18] == nil {
			return errors.New("accounts.AcceptPayment is not set")
		}
		if inst.AccountMetaSlice[19] == nil {
			return errors.New("accounts.TokenAccount is not set")
		}
		if inst.AccountMetaSlice[20] == nil {
			return errors.New("accounts.ParticipationNFTPrintingHoldingAccount is not set")
		}
		if inst.AccountMetaSlice[21] == nil {
			return errors.New("accounts.AuctionExtendedPDA is not set")
		}
	}
	return nil
}

func (inst *DeprecatedRedeemParticipationBid) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("DeprecatedRedeemParticipationBid")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=22]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                 auctionManager", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("      safetyDepositTokenStorage", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                    destination", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("               bidRedemptionKey", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("               safetyDepositBox", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("                          vault", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("            safetyDepositConfig", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("                        auction", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("                 bidderMetadata", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("                         bidder", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("                          payer", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("                   tokenProgram", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("              tokenVaultProgram", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("           tokenMetadataProgram", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("                          store", inst.AccountMetaSlice.Get(14)))
						accountsBranch.Child(ag_format.Meta("                         system", inst.AccountMetaSlice.Get(15)))
						accountsBranch.Child(ag_format.Meta("                     rentSysvar", inst.AccountMetaSlice.Get(16)))
						accountsBranch.Child(ag_format.Meta("              transferAuthority", inst.AccountMetaSlice.Get(17)))
						accountsBranch.Child(ag_format.Meta("                  acceptPayment", inst.AccountMetaSlice.Get(18)))
						accountsBranch.Child(ag_format.Meta("                          token", inst.AccountMetaSlice.Get(19)))
						accountsBranch.Child(ag_format.Meta("participationNFTPrintingHolding", inst.AccountMetaSlice.Get(20)))
						accountsBranch.Child(ag_format.Meta("             auctionExtendedPDA", inst.AccountMetaSlice.Get(21)))
					})
				})
		})
}

func (obj DeprecatedRedeemParticipationBid) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *DeprecatedRedeemParticipationBid) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewDeprecatedRedeemParticipationBidInstruction declares a new DeprecatedRedeemParticipationBid instruction with the provided parameters and accounts.
func NewDeprecatedRedeemParticipationBidInstruction(
	// Accounts:
	auctionManager ag_solanago.PublicKey,
	safetyDepositTokenStorage ag_solanago.PublicKey,
	destinationAccount ag_solanago.PublicKey,
	bidRedemptionKey ag_solanago.PublicKey,
	safetyDepositBox ag_solanago.PublicKey,
	vaultAccount ag_solanago.PublicKey,
	safetyDepositConfig ag_solanago.PublicKey,
	auction ag_solanago.PublicKey,
	bidderMetadata ag_solanago.PublicKey,
	bidder ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	tokenVaultProgram ag_solanago.PublicKey,
	tokenMetadataProgram ag_solanago.PublicKey,
	store ag_solanago.PublicKey,
	system ag_solanago.PublicKey,
	rentSysvar ag_solanago.PublicKey,
	transferAuthority ag_solanago.PublicKey,
	acceptPayment ag_solanago.PublicKey,
	tokenAccount ag_solanago.PublicKey,
	participationNFTPrintingHoldingAccount ag_solanago.PublicKey,
	auctionExtendedPDA ag_solanago.PublicKey) *DeprecatedRedeemParticipationBid {
	return NewDeprecatedRedeemParticipationBidInstructionBuilder().
		SetAuctionManagerAccount(auctionManager).
		SetSafetyDepositTokenStorageAccount(safetyDepositTokenStorage).
		SetDestinationAccount(destinationAccount).
		SetBidRedemptionKeyAccount(bidRedemptionKey).
		SetSafetyDepositBoxAccount(safetyDepositBox).
		SetVaultAccount(vaultAccount).
		SetSafetyDepositConfigAccount(safetyDepositConfig).
		SetAuctionAccount(auction).
		SetBidderMetadataAccount(bidderMetadata).
		SetBidderAccount(bidder).
		SetPayerAccount(payer).
		SetTokenProgramAccount(tokenProgram).
		SetTokenVaultProgramAccount(tokenVaultProgram).
		SetTokenMetadataProgramAccount(tokenMetadataProgram).
		SetStoreAccount(store).
		SetSystemAccount(system).
		SetRentSysvarAccount(rentSysvar).
		SetTransferAuthorityAccount(transferAuthority).
		SetAcceptPaymentAccount(acceptPayment).
		SetTokenAccount(tokenAccount).
		SetParticipationNFTPrintingHoldingAccount(participationNFTPrintingHoldingAccount).
		SetAuctionExtendedPDAAccount(auctionExtendedPDA)
}
