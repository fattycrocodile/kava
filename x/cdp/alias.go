package cdp

import (
	"github.com/kava-labs/kava/x/cdp/keeper"
	"github.com/kava-labs/kava/x/cdp/types"
)

// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/kava-labs/kava/x/cdp/keeper
// ALIASGEN: github.com/kava-labs/kava/x/cdp/types

const (
	BaseDigitFactor                 = keeper.BaseDigitFactor
	EventTypeCreateCdp              = types.EventTypeCreateCdp
	EventTypeCdpDeposit             = types.EventTypeCdpDeposit
	EventTypeCdpDraw                = types.EventTypeCdpDraw
	EventTypeCdpRepay               = types.EventTypeCdpRepay
	EventTypeCdpClose               = types.EventTypeCdpClose
	EventTypeCdpWithdrawal          = types.EventTypeCdpWithdrawal
	EventTypeCdpLiquidation         = types.EventTypeCdpLiquidation
	EventTypeBeginBlockerFatal      = types.EventTypeBeginBlockerFatal
	AttributeKeyCdpID               = types.AttributeKeyCdpID
	AttributeKeyDeposit             = types.AttributeKeyDeposit
	AttributeValueCategory          = types.AttributeValueCategory
	AttributeKeyError               = types.AttributeKeyError
	ModuleName                      = types.ModuleName
	StoreKey                        = types.StoreKey
	RouterKey                       = types.RouterKey
	QuerierRoute                    = types.QuerierRoute
	DefaultParamspace               = types.DefaultParamspace
	LiquidatorMacc                  = types.LiquidatorMacc
	SavingsRateMacc                 = types.SavingsRateMacc
	QueryGetCdp                     = types.QueryGetCdp
	QueryGetCdpDeposits             = types.QueryGetCdpDeposits
	QueryGetCdps                    = types.QueryGetCdps
	QueryGetCdpsByCollateralization = types.QueryGetCdpsByCollateralization
	QueryGetParams                  = types.QueryGetParams
	RestOwner                       = types.RestOwner
	RestCollateralDenom             = types.RestCollateralDenom
	RestRatio                       = types.RestRatio
)

var (
	// functions aliases
	NewKeeper                   = keeper.NewKeeper
	NewQuerier                  = keeper.NewQuerier
	NewCDP                      = types.NewCDP
	NewAugmentedCDP             = types.NewAugmentedCDP
	RegisterCodec               = types.RegisterCodec
	NewDeposit                  = types.NewDeposit
	NewGenesisState             = types.NewGenesisState
	DefaultGenesisState         = types.DefaultGenesisState
	GetCdpIDBytes               = types.GetCdpIDBytes
	GetCdpIDFromBytes           = types.GetCdpIDFromBytes
	CdpKey                      = types.CdpKey
	SplitCdpKey                 = types.SplitCdpKey
	DenomIterKey                = types.DenomIterKey
	SplitDenomIterKey           = types.SplitDenomIterKey
	DepositKey                  = types.DepositKey
	SplitDepositKey             = types.SplitDepositKey
	DepositIterKey              = types.DepositIterKey
	SplitDepositIterKey         = types.SplitDepositIterKey
	CollateralRatioBytes        = types.CollateralRatioBytes
	CollateralRatioKey          = types.CollateralRatioKey
	SplitCollateralRatioKey     = types.SplitCollateralRatioKey
	CollateralRatioIterKey      = types.CollateralRatioIterKey
	SplitCollateralRatioIterKey = types.SplitCollateralRatioIterKey
	NewMsgCreateCDP             = types.NewMsgCreateCDP
	NewMsgDeposit               = types.NewMsgDeposit
	NewMsgWithdraw              = types.NewMsgWithdraw
	NewMsgDrawDebt              = types.NewMsgDrawDebt
	NewMsgRepayDebt             = types.NewMsgRepayDebt
	NewParams                   = types.NewParams
	DefaultParams               = types.DefaultParams
	ParamKeyTable               = types.ParamKeyTable
	NewQueryCdpsParams          = types.NewQueryCdpsParams
	NewQueryCdpParams           = types.NewQueryCdpParams
	NewQueryCdpDeposits         = types.NewQueryCdpDeposits
	NewQueryCdpsByRatioParams   = types.NewQueryCdpsByRatioParams
	ValidSortableDec            = types.ValidSortableDec
	SortableDecBytes            = types.SortableDecBytes
	ParseDecBytes               = types.ParseDecBytes
	RelativePow                 = types.RelativePow

	// variable aliases
	ModuleCdc                           = types.ModuleCdc
	ErrCdpAlreadyExists                 = types.ErrCdpAlreadyExists
	ErrInvalidCollateralLength          = types.ErrInvalidCollateralLength
	ErrCollateralNotSupported           = types.ErrCollateralNotSupported
	ErrDebtNotSupported                 = types.ErrDebtNotSupported
	ErrExceedsDebtLimit                 = types.ErrExceedsDebtLimit
	ErrInvalidCollateralRatio           = types.ErrInvalidCollateralRatio
	ErrCdpNotFound                      = types.ErrCdpNotFound
	ErrDepositNotFound                  = types.ErrDepositNotFound
	ErrInvalidDeposit                   = types.ErrInvalidDeposit
	ErrInvalidPayment                   = types.ErrInvalidPayment
	ErrDepositNotAvailable              = types.ErrDepositNotAvailable
	ErrInvalidWithdrawAmount            = types.ErrInvalidWithdrawAmount
	ErrCdpNotAvailable                  = types.ErrCdpNotAvailable
	ErrBelowDebtFloor                   = types.ErrBelowDebtFloor
	ErrLoadingAugmentedCDP              = types.ErrLoadingAugmentedCDP
	ErrInvalidDebtRequest               = types.ErrInvalidDebtRequest
	ErrDenomPrefixNotFound              = types.ErrDenomPrefixNotFound
	CdpIDKeyPrefix                      = types.CdpIDKeyPrefix
	CdpKeyPrefix                        = types.CdpKeyPrefix
	CollateralRatioIndexPrefix          = types.CollateralRatioIndexPrefix
	CdpIDKey                            = types.CdpIDKey
	DebtDenomKey                        = types.DebtDenomKey
	GovDenomKey                         = types.GovDenomKey
	DepositKeyPrefix                    = types.DepositKeyPrefix
	PrincipalKeyPrefix                  = types.PrincipalKeyPrefix
	PreviousDistributionTimeKey         = types.PreviousDistributionTimeKey
	KeyGlobalDebtLimit                  = types.KeyGlobalDebtLimit
	KeyCollateralParams                 = types.KeyCollateralParams
	KeyDebtParam                        = types.KeyDebtParam
	KeyDistributionFrequency            = types.KeyDistributionFrequency
	KeyCircuitBreaker                   = types.KeyCircuitBreaker
	KeyDebtThreshold                    = types.KeyDebtThreshold
	KeySurplusThreshold                 = types.KeySurplusThreshold
	DefaultGlobalDebt                   = types.DefaultGlobalDebt
	DefaultCircuitBreaker               = types.DefaultCircuitBreaker
	DefaultCollateralParams             = types.DefaultCollateralParams
	DefaultDebtParam                    = types.DefaultDebtParam
	DefaultCdpStartingID                = types.DefaultCdpStartingID
	DefaultDebtDenom                    = types.DefaultDebtDenom
	DefaultGovDenom                     = types.DefaultGovDenom
	DefaultStableDenom                  = types.DefaultStableDenom
	DefaultSurplusThreshold             = types.DefaultSurplusThreshold
	DefaultDebtThreshold                = types.DefaultDebtThreshold
	DefaultPreviousDistributionTime     = types.DefaultPreviousDistributionTime
	DefaultSavingsDistributionFrequency = types.DefaultSavingsDistributionFrequency
	MaxSortableDec                      = types.MaxSortableDec
)

type (
	Keeper                 = keeper.Keeper
	CDP                    = types.CDP
	CDPs                   = types.CDPs
	AugmentedCDP           = types.AugmentedCDP
	AugmentedCDPs          = types.AugmentedCDPs
	Deposit                = types.Deposit
	Deposits               = types.Deposits
	GenesisState           = types.GenesisState
	MsgCreateCDP           = types.MsgCreateCDP
	MsgDeposit             = types.MsgDeposit
	MsgWithdraw            = types.MsgWithdraw
	MsgDrawDebt            = types.MsgDrawDebt
	MsgRepayDebt           = types.MsgRepayDebt
	Params                 = types.Params
	CollateralParam        = types.CollateralParam
	CollateralParams       = types.CollateralParams
	DebtParam              = types.DebtParam
	DebtParams             = types.DebtParams
	QueryCdpsParams        = types.QueryCdpsParams
	QueryCdpParams         = types.QueryCdpParams
	QueryCdpDeposits       = types.QueryCdpDeposits
	QueryCdpsByRatioParams = types.QueryCdpsByRatioParams
)
