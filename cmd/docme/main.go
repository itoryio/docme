package main

import (
	"fmt"
	"os"

	"github.com/itoryio/docme/agent"
	"github.com/spf13/cobra"
)

//RootCmd comment
var RootCmd = &cobra.Command{
	Use: "docme [command]",

	Short: "docme",
	Long:  `Docme docker volume plugin`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

//Log from logrus

func main() {

	RootCmd.AddCommand(agent.AgentCmd)

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

}
