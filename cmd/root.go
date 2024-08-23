/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var traceCmd = &cobra.Command{
	Use:   "iptracker",
	Short: "IPtracker CLI app",
	Long:  `IP tracker CLI app`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Initializing Tracing")
	},
}

var tracingCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace the IP",
	Long:  "Trace the IP",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				showData(ip)
			}
		} else {
			fmt.Println("Please provide IP to trace.")
		}
	},
}

func Execute() {
	err := traceCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	traceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	traceCmd.AddCommand(tracingCmd)
}

type Ip struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	Anycast  bool   `json:"anycast"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
	Readme   string `json:"readme"`
}

func showData(ip string) {
	url := "https://ipinfo.io/" + ip + "/geo"
	responseByte := getData(url)

	var data Ip
	err := json.Unmarshal(responseByte, &data)
	if err != nil {
		log.Println("Unable to read Data")
	}

	fmt.Println("DATA FOUND")
	fmt.Printf("IP :%s\nCity  :%s\nRegion : %s\nCountry : %s\nLocation : %s\nTimeZone : %s\nPostal : %s",
		data.IP,
		data.City,
		data.Region,
		data.Country,
		data.Loc,
		data.Timezone,
		data.Postal)

}

func getData(url string) []byte {

	response, err := http.Get(url)
	if err != nil {
		log.Println("Unable to get the response")
	}

	responseByte, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Unable to send the response")
	}

	return responseByte

}
