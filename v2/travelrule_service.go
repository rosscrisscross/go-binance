package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
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
	questionnaire      Questionnaire
}

type Questionnaire struct {
	IsAddressOwner int    `json:"isAddressOwner"`
	BNFType        int    `json:"bnfType,omitempty"`
	BNFName        string `json:"bnfName,omitempty"`
	Country        string `json:"country,omitempty"`
	BNFCorpName    string `json:"bnfCorpName,omitempty"`
	BNFCorpCountry string `json:"bnfCorpCountry,omitempty"`
	SendTo         int    `json:"sendTo"`
	VASP           string `json:"vasp,omitempty"`
	VASPName       string `json:"vaspName,omitempty"`
	Declaration    bool   `json:"declaration"`
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

func (s *CreateTravelRuleWithdrawService) Questionnaire(questionnaire Questionnaire) *CreateTravelRuleWithdrawService {
	s.questionnaire = questionnaire
	return s
}

// Do sends the request.
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

	r.setParam("questionnaire", string(questionnaireJSON))
	r.setParam("coin", s.coin)
	r.setParam("address", s.address)
	r.setParam("amount", s.amount)

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
	ID string `json:"id"`
}
