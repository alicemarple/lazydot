package commmand

import (
	"fmt"

	"github.com/alicemarple/lazydot/internal/constants"
	"github.com/alicemarple/lazydot/pkg/model"
	"github.com/alicemarple/lazydot/pkg/util/file"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func Search(packageName string) {
	dt := file.ReadYml[model.MetaData](constants.SyncFile)

	if packageName == "all" {
		printSearchTable(dt)
		return
	}

	for _, v := range dt {
		if v.Name == packageName {
			printSearchTable([]model.MetaData{v})
			return
		}
	}

	fmt.Printf("no package found with name: %s\n", packageName)
}

func printSearchTable[T model.PrintData](dt []T) {
	// Styles
	var (
		purple    = lipgloss.Color("99")
		gray      = lipgloss.Color("245")
		lightGray = lipgloss.Color("241")

		headerStyle  = lipgloss.NewStyle().Foreground(purple).Bold(true).Align(lipgloss.Center)
		cellStyle    = lipgloss.NewStyle().Padding(0, 1).Width(23)
		oddRowStyle  = cellStyle.Foreground(gray)
		evenRowStyle = cellStyle.Foreground(lightGray)
	)

	// Build table
	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(purple)).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return headerStyle
			case row%2 == 0:
				return evenRowStyle
			default:
				return oddRowStyle
			}
		}).
		Headers("Name", "Version", "Destination")

	// Add rows
	for _, v := range dt {
		t.Row(v.GetPName(), v.GetPVersion(), v.GetPConfDir())
	}

	// Print the table
	fmt.Println(t)
}
