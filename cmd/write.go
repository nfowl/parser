package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"nfowler.dev/siso-parser/pkg/siso"
)

func init() {
	rootCmd.AddCommand(writerCmd)
}

var writerCmd = &cobra.Command{
	Use:   "write",
	Short: "Write new siso enum files",
	Long:  "Create new files for the specified languages",
	Run: func(cmd *cobra.Command, args []string) {
		runParser(args)
	},
}

func runParser(args []string) {

	for _, v := range args {
		fmt.Println(v)
	}
	if SisoFile == "" {
		log.Printf("No File Specified")
		return
	}
	out, _ := siso.ReadXmlFile(SisoFile)
	fmt.Println(out)
}
