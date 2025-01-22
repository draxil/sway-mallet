package nextfree

import (
	"github.com/draxil/sway-mallet/sway"
	"github.com/spf13/cobra"
)

func jumpCmd(swayc sway.Sway) *cobra.Command {
	return &cobra.Command{
		Use:   "jump",
		Short: "Jump to the next free workspace",
		RunE: func(cmd *cobra.Command, args []string) error {
			ws, err := swayc.NextFreeWS(cmd.Context())
			if err != nil {
				return err
			}

			err = swayc.Run(cmd.Context(), "workspace %d", ws)
			if err != nil {
				return err
			}

			return nil
		},
	}
}
