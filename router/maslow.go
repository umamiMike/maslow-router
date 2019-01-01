package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

type policy struct {
	Name string
}

var rootCmd = &cobra.Command{
	Use:   "maslow",
	Short: "a web traffic shaper",
	Long:  "Think of all the wonderful things you will be able to do with your time",
}

var parseLeases = &cobra.Command{
	Use:   "parse-leases",
	Short: "Parse the system dnsmasq.leases file and upload system data to firebase",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("must supply the path to the dnsmasq.leases file")
			return
		}
		readAndParseLeases(args[0])
	},
}

// Write script that scans dnsmasq.log output and builds a dictionary of names and IP addresses
//currently log to stdout
//used by iptables writing command
var parseDNS = &cobra.Command{
	Use:   "parse-dns",
	Short: "Parse the system dnsmasq.log file and write to stdout",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("must supply the path to the dnsmasq.log file")
			return
		}
		readAndParseDNS(args[0])
	},
}

func init() {
	rootCmd.AddCommand(parseLeases)
	rootCmd.AddCommand(parseDNS)
	rootCmd.Execute()

}
func main() {
}
