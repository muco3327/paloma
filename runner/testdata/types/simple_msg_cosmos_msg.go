package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var amino = codec.NewLegacyAmino()
var ModuleCdc = codec.NewAminoCodec(amino)

func (msg SimpleMessage) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{sender}
}

func (msg SimpleMessage) ValidateBasic() error {
	return nil
}

func (msg SimpleMessage) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}