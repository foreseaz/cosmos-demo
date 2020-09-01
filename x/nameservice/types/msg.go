package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgSetName defines a SetName message
type MsgSetName struct {
	Name  string         `json:"name"`
	Value string         `json:"value"`
	Owner sdk.AccAddress `json:"owner"`
}

// NewMsgSetName is a constructor function for MsgSetName
func NewMsgSetName(name string, value string, owner sdk.AccAddress) MsgSetName {
	return MsgSetName{
		Name:  name,
		Value: value,
		Owner: owner,
	}
}

// Route should return the name of the module
func (msg MsgSetName) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSetName) Type() string { return "set_name" }

// ValidateBasic runs stateless checks on the message
func (msg MsgSetName) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Owner.String())
	}
	if len(msg.Name) == 0 || len(msg.Value) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Name and/or Value cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSetName) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgSetName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// const RouterKey = ModuleName // this was defined in your key.go file

// // TODO: Describe your actions, these will implment the interface of `sdk.Msg`
// // verify interface at compile time
// var _ sdk.Msg = &Msg<Action>{}

// // Msg<Action> - struct for unjailing jailed validator
// type Msg<Action> struct {
// 	ValidatorAddr sdk.ValAddress `json:"address" yaml:"address"` // address of the validator operator
// }

// // NewMsg<Action> creates a new Msg<Action> instance
// func NewMsg<Action>(validatorAddr sdk.ValAddress) Msg<Action> {
// 	return Msg<Action>{
// 		ValidatorAddr: validatorAddr,
// 	}
// }

// const <action>Const = "<action>"
