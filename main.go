package main

import (
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ribencong/go-youPipe/account"
	"github.com/ribencong/go-youPipe/service"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"time"
)

var rootCmd = &cobra.Command{
	Use: "YPManager",

	Short: "YPManager -p [password] -u [address] -s [2006-02-21] -d 14",

	Long: `""`,

	Run: mainRun,

	//Args:  cobra.MinimumNArgs(2),
}

var bootCmd = &cobra.Command{
	Use: "boot",

	Short: "YPManager boot -s [id@ip, id@ip......]",

	Long: `"YPManager boot -s [id@ip, id@ip......]"`,

	Run: bootStrapServers,

	//Args:  cobra.MinimumNArgs(2),
}

var param struct {
	password   string
	address    string
	interval   int
	startDay   string
	kingKey    string
	cipherText string
}

func main() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(bootCmd)

	rootCmd.Flags().StringVarP(&param.password, "password",
		"p", "", "Thanos's finger")

	rootCmd.Flags().StringVarP(&param.address, "address",
		"u", "", "User's address")

	rootCmd.Flags().StringVarP(&param.startDay, "startDay",
		"s", "", "License start day")

	rootCmd.Flags().IntVarP(&param.interval, "duration", "d", 0,
		"license duration in days")

	bootCmd.Flags().StringVarP(&bootServers, "server",
		"s", "", "bootstrap server list")

	bootCmd.Flags().StringVarP(&bootID, "decode",
		"d", "", "decode node id to server id@ip")
}

var bootServers = ""
var bootID = ""

func mainRun(_ *cobra.Command, _ []string) {

	thanosFinger := OpenThanosFinger(param.password)

	if !account.CheckID(param.address) {
		panic("user's address is invalid")
	}

	start := time.Now().In(time.UTC)
	if len(param.startDay) != 0 {
		s, err := time.Parse(SysTimeFormat, param.startDay)
		if err != nil {
			panic(err)
		}
		if s.Before(start) {
			panic("start time is earlier than now.")
		}
		start = s
	}

	if param.interval <= 0 {
		panic("invalid duration no in days")
	}

	l := thanosFinger.Snap(param.address, start, param.interval)
	fmt.Println(l)
	if _, err := service.ParseLicense(l); err != nil {
		panic(err)
	}
}

func bootStrapServers(_ *cobra.Command, _ []string) {

	if len(bootServers) != 0 {
		nodeIds := strings.Split(bootServers, ",")

		for _, id := range nodeIds {
			fmt.Println(base58.Encode([]byte(id)))
		}
	}

	if len(bootID) != 0 {
		fmt.Println(string(base58.Decode(bootID)))
	}

}
