package common

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/version"

	furyintcmd "github.com/gridironxyz/furyint/cmd/furyint/commands"
	"github.com/gridironxyz/furyint/conv"
	"github.com/libp2p/go-libp2p"
	"github.com/spf13/cobra"
	tmcmd "github.com/tendermint/tendermint/cmd/tendermint/commands"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/p2p"
)

// add Rollapp commands
func AddRollappCommands(rootCmd *cobra.Command, defaultNodeHome string, appCreator types.AppCreator, appExport types.AppExporter, addStartFlags types.ModuleInitFlags) {
	furyintCmd := &cobra.Command{
		Use:   "furyint",
		Short: "Furyint subcommands",
	}

	furyintCmd.AddCommand(
		ShowSequencer(),
		ShowNodeIDCmd(),
		ResetAll(),
		InitFiles(),
		tmcmd.ResetStateCmd,
		server.VersionCmd(),
	)

	furyintCmd.PersistentFlags().StringP(cli.HomeFlag, "", defaultNodeHome, "directory for config and data")

	rootCmd.AddCommand(
		furyintCmd,
		server.ExportCmd(appExport, defaultNodeHome),
		version.NewVersionCommand(),
		server.NewRollbackCmd(appCreator, defaultNodeHome),
	)
}

// ShowNodeIDCmd - ported from Tendermint, dump node ID to stdout
func ShowNodeIDCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "show-node-id",
		Short: "Show this node's ID",
		RunE: func(cmd *cobra.Command, args []string) error {
			serverCtx := server.GetServerContextFromCmd(cmd)
			cfg := serverCtx.Config

			nodeKey, err := p2p.LoadNodeKey(cfg.NodeKeyFile())
			if err != nil {
				return err
			}
			signingKey, err := conv.GetNodeKey(nodeKey)
			if err != nil {
				return err
			}
			// convert nodeKey to libp2p key
			host, err := libp2p.New(libp2p.Identity(signingKey))
			if err != nil {
				return err
			}

			fmt.Println(host.ID())
			return nil
		},
	}
}

func ShowSequencer() *cobra.Command {
	showSequencer := server.ShowValidatorCmd()
	showSequencer.Use = "show-sequencer"
	showSequencer.Short = "Show the current sequencer address"

	return showSequencer
}

func ResetAll() *cobra.Command {
	resetAll := tmcmd.ResetAllCmd
	resetAll.Short = "(unsafe) Remove all the data and WAL, reset this node's sequencer to genesis state"

	return resetAll
}

func InitFiles() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize a rollapp node directory",
		RunE: func(cmd *cobra.Command, args []string) error {
			serverCtx := server.GetServerContextFromCmd(cmd)
			cfg := serverCtx.Config
			return furyintcmd.InitFilesWithConfig(cfg)
		},
	}
}
