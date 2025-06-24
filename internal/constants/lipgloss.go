package constants

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	ErrorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("196")) // Red
	InfoStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("111")) // Blue
	SuccessStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("99"))  // Green
	// SuccessStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#BD93F9")) // Green

)
