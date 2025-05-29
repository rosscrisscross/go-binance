package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Create CreateTravelRuleWithdrawService submits a withdraw request along with travel rule information.
//
// See https://developers.binance.com/docs/wallet/travel-rule/withdraw
type CreateTravelRuleWithdrawService struct {
	c                  *Client
	coin               string
	withdrawOrderID    *string
	network            *string
	address            string
	addressTag         *string
	amount             string
	transactionFeeFlag *bool   // When making internal transfer, true for returning the fee to the destination account; false for returning the fee back to the departure account. Default false.
	name               *string // Description of the address. Address book cap is 200, space in name should be encoded into %20
	walletType         *int    // The wallet type for withdraw，0-spot wallet ，1-funding wallet. Default walletType is the current "selected wallet" under wallet->Fiat and Spot/Funding->Deposit
	timestamp          int64
	recvWindow         int64
	questionnaire      WithdrawalQuestionnaire
}

type WithdrawalQuestionnaire struct {
	IsAddressOwner int     `json:"isAddressOwner"`
	BNFType        *int    `json:"bnfType,omitempty"`
	BNFName        *string `json:"bnfName,omitempty"`
	Country        *string `json:"country,omitempty"`
	BNFCorpName    *string `json:"bnfCorpName,omitempty"`
	BNFCorpCountry *string `json:"bnfCorpCountry,omitempty"`
	SendTo         int     `json:"sendTo"`
	VASP           *string `json:"vasp,omitempty"`
	VASPName       *string `json:"vaspName,omitempty"`
	Declaration    bool    `json:"declaration"`
}

// Coin sets the coin parameter (MANDATORY).
func (s *CreateTravelRuleWithdrawService) Coin(v string) *CreateTravelRuleWithdrawService {
	s.coin = v
	return s
}

// WithdrawOrderID sets the withdrawOrderID parameter.
func (s *CreateTravelRuleWithdrawService) WithdrawOrderID(v string) *CreateTravelRuleWithdrawService {
	s.withdrawOrderID = &v
	return s
}

// Network sets the network parameter.
func (s *CreateTravelRuleWithdrawService) Network(v string) *CreateTravelRuleWithdrawService {
	s.network = &v
	return s
}

// Address sets the address parameter (MANDATORY).
func (s *CreateTravelRuleWithdrawService) Address(v string) *CreateTravelRuleWithdrawService {
	s.address = v
	return s
}

// AddressTag sets the addressTag parameter.
func (s *CreateTravelRuleWithdrawService) AddressTag(v string) *CreateTravelRuleWithdrawService {
	s.addressTag = &v
	return s
}

// Amount sets the amount parameter (MANDATORY).
func (s *CreateTravelRuleWithdrawService) Amount(v string) *CreateTravelRuleWithdrawService {
	s.amount = v
	return s
}

// TransactionFeeFlag sets the transactionFeeFlag parameter.
func (s *CreateTravelRuleWithdrawService) TransactionFeeFlag(v bool) *CreateTravelRuleWithdrawService {
	s.transactionFeeFlag = &v
	return s
}

// Name sets the name parameter.
func (s *CreateTravelRuleWithdrawService) Name(v string) *CreateTravelRuleWithdrawService {
	s.name = &v
	return s
}

func (s *CreateTravelRuleWithdrawService) WalletType(walletType int) *CreateTravelRuleWithdrawService {
	s.walletType = &walletType
	return s
}

func (s *CreateTravelRuleWithdrawService) Timestamp(timestamp int64) *CreateTravelRuleWithdrawService {
	s.timestamp = timestamp
	return s
}

func (s *CreateTravelRuleWithdrawService) RecvWindow(recvWindow int64) *CreateTravelRuleWithdrawService {
	s.recvWindow = recvWindow
	return s
}

func (s *CreateTravelRuleWithdrawService) Questionnaire(questionnaire WithdrawalQuestionnaire) *CreateTravelRuleWithdrawService {
	s.questionnaire = questionnaire
	return s
}

// Do sends the travel rule withdraw request.
func (s *CreateTravelRuleWithdrawService) Do(ctx context.Context, opts ...RequestOption) (*CreateTravelRuleWithdrawResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/localentity/withdraw/apply",
		secType:  secTypeSigned,
	}

	questionnaireJSON, err := json.Marshal(s.questionnaire)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal questionnaire: %v", err)
	}

	r.setParam("questionnaire", url.QueryEscape(string(questionnaireJSON)))
	r.setParam("coin", s.coin)
	r.setParam("address", s.address)
	r.setParam("amount", s.amount)
	r.setParam("timestamp", s.timestamp)
	r.setParam("recvWindow", s.recvWindow)

	if v := s.withdrawOrderID; v != nil {
		r.setParam("withdrawOrderId", *v)
	}
	if v := s.network; v != nil {
		r.setParam("network", *v)
	}
	if v := s.addressTag; v != nil {
		r.setParam("addressTag", *v)
	}
	if v := s.transactionFeeFlag; v != nil {
		r.setParam("transactionFeeFlag", *v)
	}
	if v := s.name; v != nil {
		r.setParam("name", *v)
	}
	if s.walletType != nil {
		r.setParam("walletType", *s.walletType)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := &CreateTravelRuleWithdrawResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

// CreateTravelRuleWithdrawResponse represents a response from CreateTravelRuleWithdrawService.
type CreateTravelRuleWithdrawResponse struct {
	ID       int    `json:"trId"`
	Accepted bool   `json:"accepted"`
	Info     string `json:"info"`
}

// ListTravelRuleDepositsService fetches deposits that require travel rule information
//
// See https://developers.binance.com/docs/wallet/travel-rule/deposit-history
type ListTravelRuleDepositsService struct {
	c                    *Client
	trID                 *string
	txID                 *string
	tranID               *string
	network              *string
	coin                 *string
	travelRuleStatus     *int
	pendingQuestionnaire *bool
	startTime            *int64
	endTime              *int64
	offset               *int
	limit                *int
	timestamp            int64
}

// TRID sets the TR ID parameter
func (s *ListTravelRuleDepositsService) TRID(v string) *ListTravelRuleDepositsService {
	s.trID = &v
	return s
}

// TXID sets the TX ID parameter
func (s *ListTravelRuleDepositsService) TXID(v string) *ListTravelRuleDepositsService {
	s.txID = &v
	return s
}

// TranID sets the transaction ID parameter
func (s *ListTravelRuleDepositsService) TranID(v string) *ListTravelRuleDepositsService {
	s.tranID = &v
	return s
}

// Network sets the network parameter
func (s *ListTravelRuleDepositsService) Network(v string) *ListTravelRuleDepositsService {
	s.network = &v
	return s
}

// Coin sets the coin parameter
func (s *ListTravelRuleDepositsService) Coin(v string) *ListTravelRuleDepositsService {
	s.coin = &v
	return s
}

// TravelRuleStatus sets the travel rule status parameter
func (s *ListTravelRuleDepositsService) TravelRuleStatus(v int) *ListTravelRuleDepositsService {
	s.travelRuleStatus = &v
	return s
}

// PendingQuestionnaire sets the pending questionnaire parameter
func (s *ListTravelRuleDepositsService) PendingQuestionnaire(v bool) *ListTravelRuleDepositsService {
	s.pendingQuestionnaire = &v
	return s
}

// StartTime sets the start time parameter
func (s *ListTravelRuleDepositsService) StartTime(v int64) *ListTravelRuleDepositsService {
	s.startTime = &v
	return s
}

// EndTime sets the end time parameter
func (s *ListTravelRuleDepositsService) EndTime(v int64) *ListTravelRuleDepositsService {
	s.endTime = &v
	return s
}

// Offset sets the offset parameter
func (s *ListTravelRuleDepositsService) Offset(v int) *ListTravelRuleDepositsService {
	s.offset = &v
	return s
}

// Limit sets the limit parameter
func (s *ListTravelRuleDepositsService) Limit(v int) *ListTravelRuleDepositsService {
	s.limit = &v
	return s
}

// Timestamp sets the timestamp parameter
func (s *ListTravelRuleDepositsService) Timestamp(v int64) *ListTravelRuleDepositsService {
	s.timestamp = v
	return s
}

func (s *ListTravelRuleDepositsService) Do(ctx context.Context, opts ...RequestOption) ([]*TravelRuleDeposit, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/localentity/deposit/history",
		secType:  secTypeSigned,
	}

	r.setParam("timestamp", s.timestamp)

	if v := s.trID; v != nil {
		r.setParam("trId", *v)
	}
	if v := s.txID; v != nil {
		r.setParam("txId", *v)
	}
	if v := s.tranID; v != nil {
		r.setParam("tranId", *v)
	}
	if v := s.network; v != nil {
		r.setParam("network", *v)
	}
	if v := s.coin; v != nil {
		r.setParam("coin", *v)
	}
	if v := s.travelRuleStatus; v != nil {
		r.setParam("travelRuleStatus", *v)
	}
	if v := s.pendingQuestionnaire; v != nil {
		r.setParam("pendingQuestionnaire", *v)
	}
	if v := s.startTime; v != nil {
		r.setParam("startTime", *v)
	}
	if v := s.endTime; v != nil {
		r.setParam("endTime", *v)
	}
	if v := s.offset; v != nil {
		r.setParam("offset", *v)
	}
	if v := s.limit; v != nil {
		r.setParam("limit", *v)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := make([]*TravelRuleDeposit, 0)
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res, nil
}

type TravelRuleDeposit struct {
	TrID                 int64                  `json:"trId"`
	TranID               int64                  `json:"tranId"`
	Amount               string                 `json:"amount"`
	Coin                 string                 `json:"coin"`
	Network              string                 `json:"network"`
	DepositStatus        int                    `json:"depositStatus"`
	TravelRuleStatus     int                    `json:"travelRuleStatus"`
	Address              string                 `json:"address"`
	AddressTag           string                 `json:"addressTag"`
	TxID                 string                 `json:"txId"`
	InsertTime           int64                  `json:"insertTime"`
	TransferType         int                    `json:"transferType"`
	ConfirmTimes         string                 `json:"confirmTimes"`
	UnlockConfirm        int                    `json:"unlockConfirm"`
	WalletType           int                    `json:"walletType"`
	RequireQuestionnaire bool                   `json:"requireQuestionnaire"`
	Questionnaire        map[string]interface{} `json:"questionnaire"`
}

// ProvideTravelRuleDepositInfoService submits travel rule information for a deposit
//
// See https://developers.binance.com/docs/wallet/travel-rule/deposit-provide-info
type ProvideTravelRuleDepositInfoService struct {
	c             *Client
	tranID        int64
	questionnaire DepositQuestionnaire
	timestamp     int64
}

type DepositQuestionnaire struct {
	DepositOriginator int     `json:"depositOriginator"`
	OrgType           *int    `json:"orgType"`
	OrgName           *string `json:"orgName"`
	Country           *string `json:"country"`
	CorpName          *string `json:"corpName"`
	CorpCountry       *string `json:"corpCountry"`
	ReceiveFrom       int     `json:"receiveFrom"`
	VASP              *string `json:"vasp"`
	VASPName          *string `json:"vaspName"`
	Declaration       bool    `json:"declaration"`
}

// TranID sets the TR ID parameter
func (s *ProvideTravelRuleDepositInfoService) TranID(v int64) *ProvideTravelRuleDepositInfoService {
	s.tranID = v
	return s
}

// Questionnaire sets the questionnaire parameter
func (s *ProvideTravelRuleDepositInfoService) Questionnaire(v DepositQuestionnaire) *ProvideTravelRuleDepositInfoService {
	s.questionnaire = v
	return s
}

// Timestamp sets the questionnaire parameter
func (s *ProvideTravelRuleDepositInfoService) Timestamp(v int64) *ProvideTravelRuleDepositInfoService {
	s.timestamp = v
	return s
}

// Do provides travel rule info to a deposit.
func (s *ProvideTravelRuleDepositInfoService) Do(ctx context.Context, opts ...RequestOption) (*ProvideTravelRuleDepositInfoResponse, error) {
	r := &request{
		method:   http.MethodPut,
		endpoint: "/sapi/v1/localentity/deposit/provide-info",
		secType:  secTypeSigned,
	}

	questionnaireJSON, err := json.Marshal(s.questionnaire)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal questionnaire: %v", err)
	}

	r.setParam("tranId", s.tranID)
	r.setParam("questionnaire", url.QueryEscape(string(questionnaireJSON)))
	r.setParam("timestamp", s.timestamp)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := &ProvideTravelRuleDepositInfoResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

type ProvideTravelRuleDepositInfoResponse struct {
	ID       int    `json:"trId"`
	Accepted bool   `json:"accepted"`
	Info     string `json:"info"`
}
