package commands

import (
	"github.com/spf13/cobra"
)

// CoreOS creates the coreos command.
func CoreOS() *Command {
	cmd := &Command{
		Command: &cobra.Command{
			Use:     "coreos",
			Aliases: []string{"core"},
			Short:   "coreos commands",
			Long:    "coreos is used to access coreos commands",
		},
		DocCategories: []string{"coreos"},
		IsIndex:       true,
	}

	return cmd
}
