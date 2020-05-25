package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&SisoFile, "sisofile", "s", "", "The Siso File to use")
}

var SisoFile string

var rootCmd = &cobra.Command{
	Use:   "siso-parser",
	Short: "SISO File Parser",
	Long:  `A Go Based Parser for the SISO Enums file`,
}

//Execute is the main entry point for the sim
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
