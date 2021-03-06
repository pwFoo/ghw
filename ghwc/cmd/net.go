//
// Use and distribution licensed under the Apache license version 2.
//
// See the COPYING file in the root project directory for full text.
//

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// netCmd represents the install command
var netCmd = &cobra.Command{
	Use:   "net",
	Short: "Show network information for the host system",
	RunE:  showNetwork,
}

// showNetwork show network information for the host system.
func showNetwork(cmd *cobra.Command, args []string) error {
	net := info.Network
	fmt.Printf("%v\n", net)

	for _, nic := range net.NICs {
		fmt.Printf(" %v\n", nic)

		enabledCaps := make([]int, 0)
		for x, cap := range nic.Capabilities {
			if cap.IsEnabled {
				enabledCaps = append(enabledCaps, x)
			}
		}
		if len(enabledCaps) > 0 {
			fmt.Printf("  enabled capabilities:\n")
			for _, x := range enabledCaps {
				fmt.Printf("   - %s\n", nic.Capabilities[x].Name)
			}
		}
	}
	return nil
}

func init() {
	rootCmd.AddCommand(netCmd)
}
