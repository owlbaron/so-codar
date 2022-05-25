/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/marcobarao/parser/internal/io"
	"github.com/marcobarao/parser/internal/lexer"
	"github.com/marcobarao/parser/internal/parser"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run is the commad used to run a program",
	Long:  "run is the commad used to run a program.",
	Run: func(cmd *cobra.Command, args []string) {
		var filePath string

		if len(args) > 0 && args[0] != "" {
			filePath = args[0]
		} else {
			panic("no file path provided")
		}

		fileContent, err := io.ReadFileContent(filePath)

		if err != nil {
			panic(err)
		}

		lexer := lexer.NewLexer(fileContent)
		parser := parser.NewParser(lexer)

		err = parser.Program()

    if err != nil {
      panic(err)
    }
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
