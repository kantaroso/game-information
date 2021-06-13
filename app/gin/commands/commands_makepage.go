package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"local.packages/game-information/controllers"
)

var (
	makepageCmd = &cobra.Command{
		Use: "makepage",
		Run: makepageCommand,
	}
)

func makepageCommand(cmd *cobra.Command, args []string) {
	makepageAllAction()
}

func makepageAllAction() {
	fmt.Println("path:/common/pv")
	controllers.MakePvJson()

	fmt.Println("path:/maker/list")
	controllers.MakeMakerListJson()

	fmt.Println("path:/maker/detail/*")
	controllers.MakeMakerInfoJson()

	fmt.Println("path:/maker/videos/*")
	controllers.MakeMakerVideosJson()

	fmt.Println("success!!")
}

func init() {
	RootCmd.AddCommand(makepageCmd)
}
