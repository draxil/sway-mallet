package main

import (
	"context"
	"os"

	"github.com/draxil/sway-mallet/cmd/nextfree"
	"github.com/draxil/sway-mallet/sway"
	"github.com/draxil/sway-mallet/util"
	gosway "github.com/joshuarubin/go-sway"
	"github.com/spf13/cobra"
)

func main() {
	cmd := cobra.Command{
		Use:   "sway-mallet",
		Short: "A tool to get sway to do things!",
	}

	cl, err := gosway.New(context.Background())
	cobra.CheckErr(err)

	swayc := sway.New(cl)

	cmd.AddCommand(nextfree.NextFreeCmd(swayc))

	err = cmd.Execute()
	if err != nil {
		util.Warn("%v", err)
		os.Exit(1)
	}
}
