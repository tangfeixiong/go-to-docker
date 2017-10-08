package cmd

import (
	"github.com/spf13/cobra"
)

func RootCommandFor(name string) *cobra.Command {
	// in, out, errout := os.Stdin, os.Stdout, os.Stderr

	root := &cobra.Command{
		Use:   name,
		Short: "Tracing application for OpenTracing",
		Long: `
        tracing
        
        This is a server of OpenTracing system, and serving to export
        span into destination, It is co-operated with jaeger.
        `,
	}
	root.AddCommand(createJaegerAgentCommand())
	root.AddCommand(createJaegerCollectorCommand())
	// root.AddCommand(createClientCommand())

	return root
}

func createClientCommand() *cobra.Command {
	command := &cobra.Command{
		Use: "client",
	}
	return command
}
