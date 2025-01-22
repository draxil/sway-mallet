package nextfree

import (
	"github.com/draxil/sway-mallet/sway"
	"github.com/spf13/cobra"
)

func NextFreeCmd(swayc sway.Sway) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nextfree",
		Short: "Commands relating to next free workspace",
	}

	cmd.AddCommand(jumpCmd(swayc))
	cmd.AddCommand(throwToCmd(swayc))

	return cmd
}
