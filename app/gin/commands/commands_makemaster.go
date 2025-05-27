package commands

import (
	"github.com/spf13/cobra"
	domainMaker "game-information/lib/domain/maker"
)

var (
	makemasterCmd = &cobra.Command{
		Use: "makemaster",
		Run: makemasterCommand,
	}
)

func makemasterCommand(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		println("no target")
	} else if args[0] == "all" {
		makemasterAllAction()
	} else if args[0] == "maker" {
		makemasterMaker()
	} else {
		println("no target")
	}
}

func makemasterAllAction() {
	println("makemasterAllAction")
}

func makemasterMaker() {
	domainMakerInstance := domainMaker.GetInstance()
	domainMakerInstance.UpdateMakerList()
}

func init() {
	RootCmd.AddCommand(makemasterCmd)
}
