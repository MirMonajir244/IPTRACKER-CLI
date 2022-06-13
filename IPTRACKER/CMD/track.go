package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
)

// trackCmd represents the track command
var trackCmd = &cobra.Command{
	Use:   "track",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			fmt.Println("Please Enter the valid IP address")
		} else {
			for _, ip := range args {
				showData(ip)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(trackCmd)
}

func getData(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Unable to get the response")
	}
	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Unable to read the response")
	}
	return respByte
}

type Ip struct {
	IP       string `json:"ip"`
	CITY     string `json:"city"`
	REGION   string `json:"region"`
	COUNTRY  string `json:"country"`
	LOCATION string `json:"location"`
	POSTAL   string `json:"postal"`
	TIMEZONE string `json:"timezone"`
}

func showData(ip string) {
	url := "https://ipinfo.io/" + ip + "/geo"
	respData := getData(url)
	data := Ip{}
	err := json.Unmarshal(respData, &data)
	if err != nil {
		log.Println("Unable to read the response")
	}
	fmt.Println("IP DATA FOUND: ")
	fmt.Printf("IP:=%s\n COUNTRY=%s\n CITY=%s\n REGION=%s\n TIMEZONE=%s\n POSTAL=%s\n LOCATION=%s", data.IP, data.COUNTRY, data.REGION, data.TIMEZONE, data.POSTAL, data.LOCATION)
}
