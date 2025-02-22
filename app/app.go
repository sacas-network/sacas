package app

import (
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
)

type App struct {
	*baseapp.BaseApp
	cdc *codec.Codec
}

func NewApp(name string, codec *codec.Codec) *App {
	return &App{
		BaseApp: baseapp.NewBaseApp(name, codec),
		cdc:     codec,
	}
}

