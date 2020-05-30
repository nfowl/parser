package cmd

import (
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
	out, _ := siso.ReadXmlFile(SisoFile)
	for _, v := range args {
		writer, _ := siso.AvailableWriters.GetWriter(v)
		for _, enum := range out.Enums {
			writer.WriteEnum(&enum)
		}
		for _, cet := range out.Entities {
			writer.WriteCet(&cet)
		}
	}
	if SisoFile == "" {
		log.Printf("No File Specified")
		return
	}
}
