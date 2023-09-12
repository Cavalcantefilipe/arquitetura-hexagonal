/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/filipe/exagonal/adapters/cli"
	"github.com/spf13/cobra"
)

var action string
var productId string
var productName string
var productPrice float64

// cliCmd represents the cli command
var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := cli.Run(&productService, action, productId, productName, productPrice)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)

	cliCmd.Flags().StringVarP(&action, "action", "a", "get", "Action to perform on the product")
	cliCmd.Flags().StringVarP(&productId, "productId", "i", "", "Id of the product")
	cliCmd.Flags().StringVarP(&productName, "productName", "n", "", "Name of the product")
	cliCmd.Flags().Float64VarP(&productPrice, "productPrice", "p", 0.0, "Price of the product")

}
