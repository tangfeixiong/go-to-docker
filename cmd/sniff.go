package cmd

import (
	"github.com/spf13/cobra"
)

func commandSnoop() *cobra.Command {
	command := &cobra.Command{
		Use: "sniff",
	}
	command.AddCommand(commandNetwork())
	return command
}

func commandNetwork() *cobra.Command {
	command := &cobra.Command{
		Use: "vnets [config file]",
	}
	return command
}
