package children

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func RunE(fn func([]int) error) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		ints := make([]int, 0, len(args))

		for _, arg := range args {
			i, err := strconv.Atoi(arg)
			if err != nil {
				return fmt.Errorf("parse argument [%s] as integer: %w", arg, err)
			}

			ints = append(ints, i)
		}

		return fn(ints)
	}
}
