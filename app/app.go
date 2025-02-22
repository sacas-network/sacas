package app

import (
	"encoding/json"
	"io"
	"os"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/tendermint/tendermint/libs/log"
	"github.com/tendermint/tendermint/libs/cli"
)

var (
	_ types.App = (*SacasApp)(nil)
)

type SacasApp struct {
	*baseapp.BaseApp
	cdc         *codec.LegacyAmino
	appCodec    codec.Codec
	invCheckPeriod uint
}

func NewSacasApp(logger log.Logger, db dbm.DB, traceStore io.Writer, loadLatest bool, home string, invCheckPeriod uint, encodingConfig EncodingConfig, options ...func(*baseapp.BaseApp)) *SacasApp {
	bApp := baseapp.NewBaseApp("sacas", logger, db, encodingConfig.TxConfig.TxDecoder(), options...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion("0.1.0")
	bApp.SetInterfaceRegistry(encodingConfig.InterfaceRegistry)

	app := &SacasApp{
		BaseApp:        bApp,
		invCheckPeriod: invCheckPeriod,
	}

	return app
}

func (app *SacasApp) Name() string { return app.BaseApp.Name() }

