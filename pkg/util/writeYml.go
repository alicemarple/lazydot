package util

import (
	"fmt"
	"os"

	"github.com/alicemarple/lazydot/pkg/model"
	"gopkg.in/yaml.v3"
)

// WriteYml function
func WriteYml(filename string, data []model.MetaData) error {
	// Read existing data
	existing := ReadYml(filename)

	for _, newItem := range data {
		found := false
		for i, existingItem := range existing {
			if existingItem.Name == newItem.Name {
				// Found: Replace the entry
				existing[i] = newItem
				found = true
				break
			}
		}
		if !found {
			// Not found: Append the new item
			existing = append(existing, newItem)
		}
	}

	// Marshal the combined result
	out, err := yaml.Marshal(existing)
	if err != nil {
		return fmt.Errorf("failed to marshal YAML: %w", err)
	}

	// Write the whole thing (overwrite!)
	err = os.WriteFile(filename, out, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
