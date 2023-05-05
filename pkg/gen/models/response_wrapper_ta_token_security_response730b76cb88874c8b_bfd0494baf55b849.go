// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849 ResponseWrapper«TaTokenSecurityResponse»-730b76cb-8887-4c8b-bfd0-494baf55b849
//
// swagger:model ResponseWrapper«TaTokenSecurityResponse»-730b76cb-8887-4c8b-bfd0-494baf55b849
type ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849 struct {

	// Code 1：Success
	Code int32 `json:"code,omitempty"`

	// Response message
	Message string `json:"message,omitempty"`

	// Response result
	Result map[string]ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849ResultAnon `json:"result,omitempty"`
}

// Validate validates this response wrapper ta token security response 730b76cb 8887 4c8b bfd0 494baf55b849
func (m *ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(m.Result) { // not required
		return nil
	}

	for k := range m.Result {

		if swag.IsZero(m.Result[k]) { // not required
			continue
		}
		if val, ok := m.Result[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849) UnmarshalBinary(b []byte) error {
	var res ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849ResultAnon contarct address
//
// swagger:model ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849ResultAnon
type ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849ResultAnon struct {

	// It describes whether the contract has the function to modify the maximum amount of transactions or the maximum token position.
	// "1" means true;
	// "0" means false;
	// No return means unknown.(Notice:(1) When "is_open_source": "0", there will be no return.
	// (2) Sometimes, when "is_proxy": "1", there will be no return.
	// (3)When the anti whale value is set to a very small value, all tradinge would fail.)
	AntiWhaleModifiable string `json:"anti_whale_modifiable,omitempty"`

	// It describes the tax when buying the token.
	// Example: "buy_tax": 0.1%.
	// No return means unknown.(Notice:(1) When "is_in_dex": "0", there will be no return.
	// (2) Buy tax will cause the actual value received when buying a token to be less than expected, and too much buy tax may lead to heavy losses.
	// (3) When "buy_tax": "1", it means buy tax is 100% or cannot buy.
	// (4) Sometimes token's anti-bot mechanism would affect our sandbox system, leading to "cannoy_buy": "1",  causing the display of "buy_tax": "1".
	// (5)Some of the token is deisgned not for sale, leading to "cannot_buy":1, causing the display of "buy_tax": "1".)
	BuyTax string `json:"buy_tax,omitempty"`

	// It describes whether this contract has the function to take back ownership.
	// "1" means true;
	// "0" means false;
	// No return means unknown.(Notice:(1) When "is_open_source": "0", there will be no return.
	// (2) Sometimes, when "is_proxy": "1", there will be no return.
	// (3) Ownership is mostly used to adjust the parameters and status of the contract, such as minting, modification of slippage, suspension of trading, setting blacklsit, etc.
	// When the contract does not have an owner (or if the owner is a black hole address) and the owner cannot be retrieved, these functions will most likely be disabled.)
	CanTakeBackOwnership string `json:"can_take_back_ownership,omitempty"`

	// It deiscribes whether the Token can be bought.
	// "1" means true;
	// "0" means false;
	// No return means unknown.(Notice:(1) Generally, "cannot_buy": "1" would be found in Reward Tokens. Such Tokens are issued as rewards for some on-chain applications and cannot be bought directly by users.
	// (2) Sometimes token's anti-bot mechanism would affect our sandbox system, causing the display of "buy_tax": "1".
	// (3) When “cannot_buy”: "1", our sandbox system might be bloked, causing the display of "buy_tax": "1" and "sell_tax": "1")
	CannotBuy string `json:"cannot_buy,omitempty"`

	// It describes whether the contract has the function restricting token holder selling all the token.
	// "1" means true;
	// "0" means false;
	// No return means unknown.(Notice:(1) When "is_in_dex": "0", there will be no return.
	// (2) This feature means that you will not be able to sell all your tokens in a single sale. Sometimes you need to leave a certain percentage of the token, e.g. 10%, sometimes you need to leave a fixed number of token, such as 10 token.
	// (3) When "buy_tax": "1", there will be no return.)
	CannotSellAll string `json:"cannot_sell_all,omitempty"`

	// It describes this contract's owner address.
	// Example: "creator_address": "0x744aF9cBb7606BB040f6FBf1c0a0B0dcBA6385E5";
	CreatorAddress string `json:"creator_address,omitempty"`

	// It describes the balance of the contract owner.
	// Example:"owner_balance": 100000000.
	CreatorBalance string `json:"creator_balance,omitempty"`

	// It describes the percentage of tokens held by the contract owner. Example:"owner_balance": 0.1.(Notice:1 means 100% here.)
	CreatorPercent string `json:"creator_percent,omitempty"`

	// It describes Dex information of where the token that can be traded.(Notice:When "is_in_dex": "0", there will be empty array. )
	Dex []*ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849ResultAnonDexItems0 `json:"dex"`

	// It describes whether the contract would call functions of other contracts when primary methods are executed.
	// "1" means true;
	// "0" means false;
	// No return means unknown.(Notice:(1) When "is_open_source": "0", there will be no return.
	// (2) External call would cause the implementation of this contract to be highly dependent on other external contracts, which may be a potential risk.)
	ExternalCall string `json:"external_call,omitempty"`

	// It describes whether the contract has hidden owners. For contract with a hidden owner, developer can still manipulate the contract even if the ownership has been abandoned.
	// “1” means true;
	// "0" means false;
	// No return means unknown.(Notice:(1) When "is_open_source": "0", there will be no return.
	// (2) Sometimes, when "is_proxy": "1", there will be no return.
	// (3) Hidden owner is often used by developers to hide ownership and is often accompanied by malicious functionality. When the hidden owner exists, it is assumed that ownership has not been abandoned.)
	HiddenOwner string `json:"hidden_owner,omitempty"`

	// It describes the number of token holders.
	// Example:"holder_count": "4342"
	HolderCount string `json:"holder_count,omitempty"`

	// Top10 holders info
	Holders []*ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849ResultAnonHoldersItems0 `json:"holders"`

	// It describes whether the token is an airdrop scam.
	// "1" means true;
	// "0" means false;
	// None means no result (Because We did not find conclusive information on whether token is an airdrop scam).(Notice:Only "is_airdrop_scam": "1" means it is an airdrop scam.)
	IsAirdropScam string `json:"is_airdrop_scam,omitempty"`

	// It describes whether the contract has the function to limit the maximum amount of transactions or the maximum token position that for single address.
	// "1" means true;
	// "0" means false;
	// No return means unknown.(Notice:(1) When "is_open_source": "0", there will be no return.
	// (2) Sometimes, when "is_proxy": "1", there will be no return. )
	IsAntiWhale string `json:"is_anti_whale,omitempty"`

	// It describes whether the blacklist function is not included in the contract. If there is a blacklist, some addresses may not be able to trade normally.
	// "1" means true;
	// "0" means false;
	// No return means unknown.(Notice:(1) When "is_open_source": "0",  there will be no return.
	// (2) Sometimes, when "is_proxy": "1", there will be no return.
	// (3) The contract owner may add any address into the blacklist, and the token holder in blacklist will not be able to trade. Abuse of the blacklist function will lead to great risks.
	// (4) For contracts without an owner (or the owner is a black hole address), the blacklist will not be able to get updated. However, the existing blacklist is still in effect.)
	IsBlacklisted string `json:"is_blacklisted,omitempty"`

	// It describes whether the token is a honeypot. "HoneyPot" means that the token maybe cannot be sold because of the token contract's function, Or the token contains malicious code.
	// "1" means true;
	// "0" means false;
	// No return means unknown.(Notice:(1) When "is_open_source": "0", there will be no return.
	// (2) Sometimes, when "is_proxy": "1", there will be no return.
	// (3) Hight risk, definitely scam.)
	IsHoneypot string `json:"is_honeypot,omitempty"`

	// It describes whether the token can be traded on the main Dex.
	// "1" means true;
	// "0" means false(Notice:It only counts when the token has a marketing pair with mainstream coins/tokens.)
	IsInDex string `json:"is_in_dex,omitempty"`

	// It describes whether this contract has the function to mint tokens.
	// "1" means true;
	// "0" means false;
	// No return means unknown.(Notice:(1) When "is_open_source": "0", there will be no return.
	// (2) Sometimes, when "is_proxy": "1", there will be no return.
	// (3) Mint function will directly trigger a massive sell-off, causing the coin price to plummet. It is extremely risky.
	// (4) This function generally relies on ownership. When the contract does not have an owner (or if the owner is a black hole address) and the owner cannot be retrieved, this function will most likely be disabled.)
	IsMintable string `json:"is_mintable,omitempty"`

	// It describes whether this contract is open source.
	// "1" means true;
	// "0" means false.(Notice:Un-open-sourced contracts may hide various unknown mechanisms and are extremely risky. When the contract is not open source, we will not be able to detect other risk items.)
	IsOpenSource string `json:"is_open_source,omitempty"`

	// It describes whether this contract has a proxy contract.
	// "1" means true;
	// "0" means false;
	// No return means unknown.(Notice:(1) When "is_open_source": "0", there will be no return.
	// (2) Most Proxy contracts are accompanied by modifiable implementation contracts, and implementation contracts may contain significant potential risk. When the contract is a Proxy, we will stop detecting other risk items.)
	IsProxy string `json:"is_proxy,omitempty"`

	// It describes whether the token is true or fake.
	// "1" means true token;
	// "0" means fake token;
	// None means no result (Because we did not find decisive information about the truth or falsity)(Notice:Only "is_true_token": "0" means it is a fake token.)
	IsTrueToken string `json:"is_true_token,omitempty"`

	// It describes whether the whitelist function is not included in the contract. If there is a whitelist, some addresses may not be able to trade normally.
	// "1" means true;
	// "0" means false;
	// No return means unknown.(Notice:(1) When "is_open_source": "0", there will be no return.
	// (2) Sometimes, when "is_proxy": "1", there will be no return.
	// (3) Whitelisting is mostly used to allow specific addresses to make early transactions, tax-free, and not affected by transaction suspension.
	// (4) For contracts without an owner (or the owner is a black hole address), the whitelist will not be able to get updated. However, the existing whitelist is still in effect.)
	IsWhitelisted string `json:"is_whitelisted,omitempty"`

	// It describes the number of LP token holders.
	// Example:"lp_holder_count": "4342".
	// No return means no LP.(Notice:When "is_in_dex": "0", there will be no return.)
	LpHolderCount string `json:"lp_holder_count,omitempty"`

	// Top10 LP token holders info(Notice:When "is_in_dex": "0", there will be no return. )
	LpHolders []*ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849ResultAnonLpHoldersItems0 `json:"lp_holders"`

	// It describes the supply number of the LP token.
	// Example:"lp_total_supply": "100000000".
	// No return means no LP.(Notice:(1) When "is_in_dex": "0", there will be no return.
	// (2) It is LP token number, NOT token number)
	LpTotalSupply string `json:"lp_total_supply,omitempty"`

	// It describes whether the contract has other things investors need to know.
	// Example:
	// "note”: “Contract owner is a multisign contract.”(Notice:(1) If we haven't found any other thing which is valuable yet, there will be no return.
	// (2) Type: string.)
	Note string `json:"note,omitempty"`

	// It describes whether the contract has other potential risks.
	// Example:
	// “other_potential_risks”: “Owner can set different transaction taxes for each user, which can trigger serious losses.”(Notice:(1) If we haven't found any other potential risk yet, there will be no return.
	// (2) Type: string.)
	OtherPotentialRisks string `json:"other_potential_risks,omitempty"`

	// It describes this contract's owner address.
	// Example: "owner_address": "0x744aF9cBb7606BB040f6FBf1c0a0B0dcBA6385E5";
	// No return means unknown; Return empty means there is no ownership or can't find ownership.(Notice:(1) When "is_open_source": "0", there will be no return.
	// (2) Sometimes, when "is_proxy": "1", there will be no return.
	// (3) Ownership is mostly used to adjust the parameters and status of the contract, such as minting, modification of slippage, suspension of trading, setting blacklist, etc.
	// When the contract does not have an owner (or if the owner is a black hole address) and the owner cannot be retrieved, these functions will most likely be disabled.)
	OwnerAddress string `json:"owner_address,omitempty"`

	// It describes the balance of the contract owner.
	// Example: "owner_balance": "100000000".
	// No return or return empty means there is no ownership or can't find ownership.(Notice:When "owner_address" returns empty, or no return, there will be no return.)
	OwnerBalance string `json:"owner_balance,omitempty"`

	// It describes whether the contract owner has the authority to change the balance of any token holder.
	// "1" means true;
	// "0" means false;
	// No return means unknown.(Notice:(1) When "is_open_source": "0", there will be no return.
	// (2) Sometimes, when "is_proxy": "1", there will be no return.
	// (3) Token with this feature means that the owner can modify anyone's balance, resulting in an asset straight to zero or a massive minting and sell-off.
	// (4) This function generally relies on ownership. When the contract does not have an owner (or if the owner is a black hole address) and the owner cannot be retrieved, this function will most likely be disabled.)
	OwnerChangeBalance string `json:"owner_change_balance,omitempty"`

	// It describes the percentage of tokens held by the contract owner.
	// Example:"owner_balance": "0.1".
	// No return or return empty means there is no ownership or can't find ownership.(Notice:(1) 1 means 100% here.
	// (2) When "owner_address" returns empty, or no return, there will be no return.)
	OwnerPercent string `json:"owner_percent,omitempty"`

	// It describes whether the owner can set a different tax rate for every assigned address.
	// "1" means ture;
	// "0" means false;
	// No return means unknown.(Notice:(1) When "is_open_source": "0",  there will be no return.
	// (2) Sometimes, when "is_proxy": "1", there will be no return.
	// (3) The contract owner may set a very outrageous tax rate for assigned address to block it from trading. Abuse of this funtcion will lead to great risks.
	// (4) For contracts without an owner (or the owner is a black hole address), this function would not able to be used. However, the existing tax rate would be still in effect.)
	PersonalSlippageModifiable string `json:"personal_slippage_modifiable,omitempty"`

	// It describes whether this contract can self destruct.
	// "1" means true;
	// "0" means false;
	// No return means unknown.(Notice:(1) When "is_open_source": "0", there will be no return.
	// (2) When the self-destruct function is triggered, this contract will be destroyed, all functions will be unavailable, and all related assets will be erased.)
	Selfdestruct string `json:"selfdestruct,omitempty"`

	// It describes the tax when selling the token.
	// Example: "sell_tax": 0.1%.
	// No return means unknown.(Notice:(1) When "is_in_dex": "0", there will be no return.
	// (2) Sell tax will cause the actual value received when selling a token to be less than expected, and too much buy tax may lead to large losses.
	// (3) When "sell_tax": "1", it means sell-tax is 100% or this token cannot be sold.
	// (4) Sometimes token's  trading-cool-down mechanism would affect our sandbox system. When "trading_cooldown": "1", "sell_tax" may return "1".)
	SellTax string `json:"sell_tax,omitempty"`

	// It describes whether the trading tax can be modifiable by token contract.
	// "1" means true;
	// "0" means false;
	// No return means unknown.(Notice:(1) When When "is_open_source": "0", there will be no return.
	// (2) Sometimes, when "is_proxy": "1", there will be no return.
	// (3) Token with modifiable tax means that the contract owner can modify the buy tax or sell tax of the token. This may cause some losses, especially since some contracts have unlimited modifiable tax rates, which would make the token untradeable.
	// (4) This function generally relies on ownership. When the contract does not have an owner (or if the owner is a black hole address) and the owner cannot be retrieved, this function will most likely be disabled.)
	SlippageModifiable string `json:"slippage_modifiable,omitempty"`

	// Token Name
	TokenName string `json:"token_name,omitempty"`

	// Token Symbol
	TokenSymbol string `json:"token_symbol,omitempty"`

	// It describes the supply number of the token.
	// Example:"total_supply": 100000000
	TotalSupply string `json:"total_supply,omitempty"`

	// It describes whether the contract has trading-cool-down mechanism which can limits the minimum time between two transactions.
	// "1" means true;
	// "0" means false;
	// No return means unknown.(Notice:(1) When "is_open_source": "0", there will be no return.
	// (2) Sometimes, when "is_proxy": "1", there will be no return. )
	TradingCooldown string `json:"trading_cooldown,omitempty"`

	// It describes whether trading can be pausable by token contract.
	// "1" means true;
	// "0" means false;
	// No return means unknown.(Notice:(1) When "is_open_source": "0", there will be no return.
	// (2) Sometimes, when "is_proxy": "1", there will be no return.
	// (3) This feature means that the contract owner will be able to suspend trading at any time, after that anyone will not be able to sell, except those who have special authority.
	// (4) This function generally relies on ownership. When the contract does not have an owner (or if the owner is a black hole address) and the owner cannot be retrieved, this function will most likely be disabled.)
	TransferPausable string `json:"transfer_pausable,omitempty"`

	// It describes whether the token is a famous and trustworthy one. "1" means true; No return no result (Because We did not find conclusive information on whether token is a airdrop scam).(Notice:(1) Only "trust_list": "1" means it is a famous and trustworthy token.
	// (2) No return doesn't mean it is risky.)
	TrustList string `json:"trust_list,omitempty"`
}

// Validate validates this response wrapper ta token security response730b76cb88874c8b bfd0494baf55b849 result anon
func (m *ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849ResultAnon) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHolders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLpHolders(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849ResultAnon) validateDex(formats strfmt.Registry) error {

	if swag.IsZero(m.Dex) { // not required
		return nil
	}

	for i := 0; i < len(m.Dex); i++ {
		if swag.IsZero(m.Dex[i]) { // not required
			continue
		}

		if m.Dex[i] != nil {
			if err := m.Dex[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dex" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849ResultAnon) validateHolders(formats strfmt.Registry) error {

	if swag.IsZero(m.Holders) { // not required
		return nil
	}

	for i := 0; i < len(m.Holders); i++ {
		if swag.IsZero(m.Holders[i]) { // not required
			continue
		}

		if m.Holders[i] != nil {
			if err := m.Holders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("holders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849ResultAnon) validateLpHolders(formats strfmt.Registry) error {

	if swag.IsZero(m.LpHolders) { // not required
		return nil
	}

	for i := 0; i < len(m.LpHolders); i++ {
		if swag.IsZero(m.LpHolders[i]) { // not required
			continue
		}

		if m.LpHolders[i] != nil {
			if err := m.LpHolders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("lp_holders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849ResultAnon) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849ResultAnon) UnmarshalBinary(b []byte) error {
	var res ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849ResultAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849ResultAnonDexItems0 response wrapper ta token security response730b76cb88874c8b bfd0494baf55b849 result anon dex items0
//
// swagger:model ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849ResultAnonDexItems0
type ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849ResultAnonDexItems0 struct {

	// Liquidity is converted to USDT denomination.
	Liquidity string `json:"liquidity,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	//  It only counts when the token has a marketing pair with mainstream
	Pair string `json:"pair,omitempty"`
}

// Validate validates this response wrapper ta token security response730b76cb88874c8b bfd0494baf55b849 result anon dex items0
func (m *ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849ResultAnonDexItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849ResultAnonDexItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849ResultAnonDexItems0) UnmarshalBinary(b []byte) error {
	var res ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849ResultAnonDexItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849ResultAnonHoldersItems0 response wrapper ta token security response730b76cb88874c8b bfd0494baf55b849 result anon holders items0
//
// swagger:model ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849ResultAnonHoldersItems0
type ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849ResultAnonHoldersItems0 struct {

	// It describes the holder address;
	Address string `json:"address,omitempty"`

	// It describes the balance of the holder.
	Balance string `json:"balance,omitempty"`

	// It describes whether the holder is a contract "1" means true; "0" means false.
	IsContract int32 `json:"is_contract,omitempty"`

	// It describes whether the tokens owned by the holder are locked "1" means true; "0" means false;
	// (3) “tag” describes the address's public tag. Example:Burn (Notice:About "locked": We only support the token lock addresses or black hole addresses that we have included. )
	IsLocked int32 `json:"is_locked,omitempty"`

	// It is an array, decribes lock position info of this holder, only shows when "locked": 1. This Array may contain multiple objects for multiple locking info. In every objetc, "amount" describes the number of token locked, "end_time" describes when the token will be unlocked, "opt_time" describes when the token was locked.(Notice:When "locked":0, or lock address is a black hole address,  "locked_detail" will be no return.)
	LockedDetail []string `json:"locked_detail"`

	// It  describes the percentage of tokens held by this holder (Notice:About "percent": 1 means 100% here.)
	Percent string `json:"percent,omitempty"`

	// It describes the address's public tag. Example:Burn Address/Deployer;
	Tag string `json:"tag,omitempty"`
}

// Validate validates this response wrapper ta token security response730b76cb88874c8b bfd0494baf55b849 result anon holders items0
func (m *ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849ResultAnonHoldersItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849ResultAnonHoldersItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849ResultAnonHoldersItems0) UnmarshalBinary(b []byte) error {
	var res ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849ResultAnonHoldersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849ResultAnonLpHoldersItems0 response wrapper ta token security response730b76cb88874c8b bfd0494baf55b849 result anon lp holders items0
//
// swagger:model ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849ResultAnonLpHoldersItems0
type ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849ResultAnonLpHoldersItems0 struct {

	// It describes the holder address;
	Address string `json:"address,omitempty"`

	// It describes the balance of the holder.
	Balance string `json:"balance,omitempty"`

	// It describes whether the holder is a contract "1" means true; "0" means false.
	IsContract int32 `json:"is_contract,omitempty"`

	// It describes whether the tokens owned by the holder are locked "1" means true; "0" means false;
	// (3) “tag” describes the address's public tag. Example:Burn (Notice:About "locked": We only support the token lock addresses or black hole addresses that we have included. )
	IsLocked int32 `json:"is_locked,omitempty"`

	// It is an array, decribes lock position info of this holder, only shows when "locked": 1. This Array may contain multiple objects for multiple locking info. In every objetc, "amount" describes the number of token locked, "end_time" describes when the token will be unlocked, "opt_time" describes when the token was locked.(Notice:When "locked":0, or lock address is a black hole address,  "locked_detail" will be no return.)
	LockedDetail []string `json:"locked_detail"`

	// It  describes the percentage of tokens held by this holder (Notice:About "percent": 1 means 100% here.)
	Percent string `json:"percent,omitempty"`

	// It describes the address's public tag. Example:Burn Address/Deployer;
	Tag string `json:"tag,omitempty"`
}

// Validate validates this response wrapper ta token security response730b76cb88874c8b bfd0494baf55b849 result anon lp holders items0
func (m *ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849ResultAnonLpHoldersItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849ResultAnonLpHoldersItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849ResultAnonLpHoldersItems0) UnmarshalBinary(b []byte) error {
	var res ResponseWrapperTaTokenSecurityResponse730b76cb88874c8bBfd0494baf55b849ResultAnonLpHoldersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
