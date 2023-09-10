package cli

import (
	"strconv"

	"encoding/json"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"paychex/x/paychex/types"
)

var _ = strconv.Itoa(0)

func CmdSendPayroll() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-payroll [newpayroll]",
		Short: "Broadcast message send-payroll",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argNewpayroll := new(types.Payroll)
			err = json.Unmarshal([]byte(args[0]), argNewpayroll)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSendPayroll(
				clientCtx.GetFromAddress().String(),
				argNewpayroll,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
