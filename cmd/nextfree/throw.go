package nextfree

import (
	"github.com/draxil/sway-mallet/sway"
	"github.com/draxil/sway-mallet/util"
	"github.com/spf13/cobra"
)

func throwToCmd(swayc sway.Sway) *cobra.Command {
	follow := false

	cmd := &cobra.Command{
		Use:   "throw",
		Short: "throw focused window to the next free workspace",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			focused, err := swayc.FocusedWin(ctx)
			if err != nil {
				return err
			}

			if focused == "" {
				util.Warn("no focused window")
				return nil
			}

			nextFree, err := swayc.NextFreeWS(ctx)
			if err != nil {
				return err
			}

			err = swayc.Run(ctx, "move window to workspace number %d",
				nextFree)
			if err != nil {
				return err
			}

			if !follow {
				return nil
			}

			err = swayc.Run(ctx, "workspace %d", nextFree)
			if err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().BoolVarP(&follow, "follow", "f", false, "follow the window to the desktop")

	return cmd
}
