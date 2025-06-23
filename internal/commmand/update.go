package commmand

import (
	"fmt"

	"github.com/alicemarple/lazydot/pkg/util/web"
)

func Update() {
	fmt.Println("update the database")
	web.GetSyncData()
}
