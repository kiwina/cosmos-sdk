// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/gov/types
package gov

import (
	"github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	MaxDescriptionLength         = types.MaxDescriptionLength
	MaxTitleLength               = types.MaxTitleLength
	DefaultCodespace             = types.DefaultCodespace
	CodeUnknownProposal          = types.CodeUnknownProposal
	CodeInactiveProposal         = types.CodeInactiveProposal
	CodeAlreadyActiveProposal    = types.CodeAlreadyActiveProposal
	CodeAlreadyFinishedProposal  = types.CodeAlreadyFinishedProposal
	CodeAddressNotStaked         = types.CodeAddressNotStaked
	CodeInvalidContent           = types.CodeInvalidContent
	CodeInvalidProposalType      = types.CodeInvalidProposalType
	CodeInvalidVote              = types.CodeInvalidVote
	CodeInvalidGenesis           = types.CodeInvalidGenesis
	CodeInvalidProposalStatus    = types.CodeInvalidProposalStatus
	CodeProposalHandlerNotExists = types.CodeProposalHandlerNotExists
	ModuleName                   = types.ModuleName
	StoreKey                     = types.StoreKey
	RouterKey                    = types.RouterKey
	QuerierRoute                 = types.QuerierRoute
	DefaultParamspace            = types.DefaultParamspace
	TypeMsgDeposit               = types.TypeMsgDeposit
	TypeMsgVote                  = types.TypeMsgVote
	TypeMsgSubmitProposal        = types.TypeMsgSubmitProposal
	StatusNil                    = types.StatusNil
	StatusDepositPeriod          = types.StatusDepositPeriod
	StatusVotingPeriod           = types.StatusVotingPeriod
	StatusPassed                 = types.StatusPassed
	StatusRejected               = types.StatusRejected
	StatusFailed                 = types.StatusFailed
	ProposalTypeText             = types.ProposalTypeText
	ProposalTypeSoftwareUpgrade  = types.ProposalTypeSoftwareUpgrade
	OptionEmpty                  = types.OptionEmpty
	OptionYes                    = types.OptionYes
	OptionAbstain                = types.OptionAbstain
	OptionNo                     = types.OptionNo
	OptionNoWithVeto             = types.OptionNoWithVeto
)

var (
	// functions aliases
	RegisterCodec                    = types.RegisterCodec
	RegisterProposalTypeCodec        = types.RegisterProposalTypeCodec
	ValidateAbstract                 = types.ValidateAbstract
	ErrUnknownProposal               = types.ErrUnknownProposal
	ErrInactiveProposal              = types.ErrInactiveProposal
	ErrAlreadyActiveProposal         = types.ErrAlreadyActiveProposal
	ErrAlreadyFinishedProposal       = types.ErrAlreadyFinishedProposal
	ErrAddressNotStaked              = types.ErrAddressNotStaked
	ErrInvalidProposalContent        = types.ErrInvalidProposalContent
	ErrInvalidProposalType           = types.ErrInvalidProposalType
	ErrInvalidVote                   = types.ErrInvalidVote
	ErrInvalidGenesis                = types.ErrInvalidGenesis
	ErrNoProposalHandlerExists       = types.ErrNoProposalHandlerExists
	KeyProposal                      = types.KeyProposal
	KeyDeposit                       = types.KeyDeposit
	KeyVote                          = types.KeyVote
	KeyDepositsSubspace              = types.KeyDepositsSubspace
	KeyVotesSubspace                 = types.KeyVotesSubspace
	PrefixActiveProposalQueueTime    = types.PrefixActiveProposalQueueTime
	KeyActiveProposalQueueProposal   = types.KeyActiveProposalQueueProposal
	PrefixInactiveProposalQueueTime  = types.PrefixInactiveProposalQueueTime
	KeyInactiveProposalQueueProposal = types.KeyInactiveProposalQueueProposal
	NewMsgSubmitProposal             = types.NewMsgSubmitProposal
	NewMsgDeposit                    = types.NewMsgDeposit
	NewMsgVote                       = types.NewMsgVote
	NewProposal                      = types.NewProposal
	ProposalStatusFromString         = types.ProposalStatusFromString
	ValidProposalStatus              = types.ValidProposalStatus
	NewTallyResult                   = types.NewTallyResult
	NewTallyResultFromMap            = types.NewTallyResultFromMap
	EmptyTallyResult                 = types.EmptyTallyResult
	NewTextProposal                  = types.NewTextProposal
	NewSoftwareUpgradeProposal       = types.NewSoftwareUpgradeProposal
	RegisterProposalType             = types.RegisterProposalType
	ContentFromProposalType          = types.ContentFromProposalType
	IsValidProposalType              = types.IsValidProposalType
	ProposalHandler                  = types.ProposalHandler
	VoteOptionFromString             = types.VoteOptionFromString
	ValidVoteOption                  = types.ValidVoteOption

	// variable aliases
	ModuleCdc                   = types.ModuleCdc
	KeyDelimiter                = types.KeyDelimiter
	KeyNextProposalID           = types.KeyNextProposalID
	PrefixActiveProposalQueue   = types.PrefixActiveProposalQueue
	PrefixInactiveProposalQueue = types.PrefixInactiveProposalQueue
)

type (
	Content                 = types.Content
	Handler                 = types.Handler
	Deposit                 = types.Deposit
	Deposits                = types.Deposits
	MsgSubmitProposal       = types.MsgSubmitProposal
	MsgDeposit              = types.MsgDeposit
	MsgVote                 = types.MsgVote
	Proposal                = types.Proposal
	Proposals               = types.Proposals
	TallyResult             = types.TallyResult
	TextProposal            = types.TextProposal
	SoftwareUpgradeProposal = types.SoftwareUpgradeProposal
	Vote                    = types.Vote
	Votes                   = types.Votes
	VoteOption              = types.VoteOption
)
