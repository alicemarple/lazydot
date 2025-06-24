package commmand

import (
	"fmt"

	"github.com/alicemarple/lazydot/internal/constants"
	"github.com/alicemarple/lazydot/pkg/model"
	"github.com/alicemarple/lazydot/pkg/util/file"
	"github.com/alicemarple/lazydot/pkg/util/tables"
)

// Query function
func Query(packageName string) {
	// Read data from the YAML
	dt := file.ReadYml[model.MetaData](constants.LocalFile)
	if packageName == "all" {
		tables.PrintTable(dt)
		return
	}

	for _, v := range dt {
		if v.Name == packageName {
			tables.PrintTable([]model.MetaData{v})
			return
		}
	}

	fmt.Printf("no package found with name: %s\n", packageName)

	// tables.PrintTable(dt)
}
