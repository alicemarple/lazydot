package commmand

import (
	"fmt"

	"github.com/alicemarple/lazydot/internal/constants"
	"github.com/alicemarple/lazydot/pkg/util/web"
)

func Update() {
	fmt.Println(constants.SuccessStyle.Render(fmt.Sprint("Success: Database Updated ")))
	web.GetSyncData()
}
