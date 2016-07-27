package main

import (
	"log"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

// RootCmd is root command of remote-docker-exec
var rootCmd = &cobra.Command{
	Use:     "remote-docker-exec [ContainerId]",
	Short:   "remote-docker-exec is a WebSocket with SSL shell, it can connect to docker container, and act as a remote docker exec.",
	Long:    "remote-docker-exec is a WebSocket with SSL shell, it can connect to docker container, and act as a remote docker exec. For linker internal use only.",
	Example: "./remote-docker-exec ffcede5a47cb",
	Run: func(cmd *cobra.Command, args []string) {
		if runtime.GOOS != "linux" {
			log.Fatalln("Only linux is supported now.")
			return
		}

		// example
		// remote-docker-exec c3598346f0d7
		if cmd.Flags().NArg() == 1 {
			containerId := cmd.Flags().Args()[0]
			remoteDockerExec(containerId)
		}

		// others
		cmd.SetArgs([]string{"--help"})
		if err := cmd.Execute(); err != nil {
			log.Fatalf("command arguments error: %v", err)
			os.Exit(-1)
		}
		return
	},
}
