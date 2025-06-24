package commmand

import (
	"fmt"
	"os"

	"github.com/alicemarple/lazydot/internal/constants"
	"github.com/alicemarple/lazydot/pkg/model"
	"github.com/alicemarple/lazydot/pkg/util/file"
	"github.com/alicemarple/lazydot/pkg/util/tables"
	"gopkg.in/yaml.v2"
)

func Remove(packageName string) {
	var tempmd []model.MetaData

	md := file.ReadYml[model.MetaData](constants.LocalFile)
	for _, v := range md {
		if v.Name != packageName {
			tempmd = append(tempmd, v)
		} else {
			tables.PrintTable([]model.MetaData{v})
		}
	}

	// 3️⃣ Marshal the new list
	data, err := yaml.Marshal(tempmd)
	if err != nil {
		fmt.Printf("error marshalling data: %v\n", err)
		return
	}

	// 4️⃣ Write to the same file
	err = os.WriteFile(constants.LocalFile, data, 0644)
	if err != nil {
		fmt.Printf("error writing file: %v\n", err)
		return
	}

	removeScript := "/mnt/e/projects/golang/lazydot/scripts/remove.sh"
	Runscript(packageName, removeScript)
	fmt.Println(constants.SuccessStyle.Render(fmt.Sprint("Done! : Package removed.")))
}
