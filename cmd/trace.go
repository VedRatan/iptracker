/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace the IP",
	Long: `Trace any IP`,
	Run: func(cmd *cobra.Command, args []string) {
		if(len(args) > 0){
				for _, ip := range(args){
					if(is_ipv4(ip)){
						showData(ip)
					}else {
						fmt.Println("invalid ip")
					}
				}
		} else{
			fmt.Println("No IP provided to trace.")
		}
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)

}

type Ip struct{
	IP       string `json::"ip"`
	City     string `json::"city"`
	Region   string `json::"region"`
	Country  string `json::"country"`
	Loc      string `json::"loc"`
	Timezone string `json::"timezone"`
	Postal   string `json::"postal"`
}

func showData(ip string){
	url := "http://ipinfo.io/"+ip+"/geo"
	responseByte := getData(url)

	data := Ip{}

	err := json.Unmarshal(responseByte, &data)

	if(err != nil){
		log.Println("Unable to unmarshal the reponse")
	} else {

		fmt.Printf("IP: %s\nRegion: %s\nCountry: %s\nCity: %s\nLocation: %s\nTimeZone: %s\nPostalCode: %s\n",data.IP, data.Region,data.Country,data.City, data.Loc, data.Timezone, data.Postal )
	}

	
}

func getData(url string) []byte{
	response, err := http.Get(url)
	if(err!=nil){
		log.Panic("Unable to get the response")
	}

	responseByte, err := ioutil.ReadAll(response.Body)

	if(err != nil){
		log.Panic("Unable to read the response")
	}

	return responseByte
}

func is_ipv4(host string) bool {
	parts := strings.Split(host, ".")

	if len(parts) < 4 {
		return false
	}
	
	for _,x := range parts {
		if i, err := strconv.Atoi(x); err == nil {
			if i < 0 || i > 255 {
			return false
		}
		} else {
			return false
		}

	}
	return true
}
