package flags

import (
	"fmt"

	"github.com/alicemarple/lazydot/internal/commmand"
	"github.com/alicemarple/lazydot/internal/constants"

	// "github.com/alicemarple/lazydot/pkg/util/setup"
	"github.com/spf13/cobra"
)

func FLagSetup(cmd *cobra.Command, packageName string) {
	isSync := cmd.Flags().Changed("sync")
	isRemove := cmd.Flags().Changed("remove")
	isQuery := cmd.Flags().Changed("query")
	isUpdate := cmd.Flags().Changed("update")
	isSearch := cmd.Flags().Changed("search")

	if isSync {
		// setup.Setup(constants.ConfigFile)
		commmand.Sync(constants.SyncFile, packageName)
	} else if isRemove {
		commmand.Remove(packageName)
	} else if isQuery {
		commmand.Query()
	} else if isUpdate {
		commmand.Update()
	} else if isSearch {
		commmand.Search(packageName)
	} else {
		fmt.Println("Give the proper flag")
	}
}
