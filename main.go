package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var config struct {
	Proxy         string
	MinTLSVersion float64
	MaxTLSVersion float64
	URL           string
}

var rootCmd = &cobra.Command{
	Use:   "https-client url",
	Short: "Perform a HTTPS get request",
	Long:  "Perform a HTTPS get request",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Proxy = ", config.Proxy)
		fmt.Println("Min TLS Version = ", config.MinTLSVersion)
		fmt.Println("Max TLS Version = ", config.MaxTLSVersion)
	},
}

func main() {

	rootCmd.Flags().StringVarP(&config.Proxy, "proxy", "p", "http://localhost:8080", "Address of proxy to use")
	rootCmd.Flags().Float64Var(&config.MinTLSVersion, "mintlsversion", 1.0, "The minimum TLS version to allow")
	rootCmd.Flags().Float64Var(&config.MaxTLSVersion, "maxtlsversion", 1.2, "The maximum TLS version to allow")

	rootCmd.Execute()
}
