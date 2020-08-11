package main

import (
	"os"

	"github.com/spf13/cobra"
)

func maindda() {

	cmd := &cobra.Command{
		Use:          "fii",
		Short:        "Hello Sacaners",
		SilenceUsage: true,
	}

	cmd.AddCommand(printTimeCmd())

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}

}
