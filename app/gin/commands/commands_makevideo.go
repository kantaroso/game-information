package commands

import (
	"errors"

	"github.com/spf13/cobra"
	domainMaker "game-information/lib/domain/maker"
)

var (
	makevideoCmd = &cobra.Command{
		Use: "makevideo",
		Run: makevideoCommand,
	}
)

func makevideoCommand(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		makevideoAllAction()
	} else {
		makevideoAction(args[0])
	}
}

func makevideoAllAction() {

	domainMakerInstance := domainMaker.GetInstance()
	makers := domainMakerInstance.GetMakerList()
	if len(*makers) == 0 {
		Exit(errors.New("no targets.."), 1)
	}
	println("target...")
	for _, item := range *makers {
		println("code:" + item.Code)
		if !domainMakerInstance.UpdateVideoList(item.ID) {
			Exit(errors.New("Error update video list code:"+item.Code), 1)
		}
	}
	println("success!!")
}

func makevideoAction(path string) {
	domainMakerInstance := domainMaker.GetInstance()
	maker := domainMakerInstance.GetMaker(path)
	if maker.ID == 0 {
		Exit(errors.New("no targets.."), 1)
	}
	println("code:" + maker.Code)
	if !domainMakerInstance.UpdateVideoList(maker.ID) {
		Exit(errors.New("Error update video list ID:"+maker.Code), 1)
	}
	println("success!!")
}

func init() {
	RootCmd.AddCommand(makevideoCmd)
}
