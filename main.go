package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/youpipe/go-youPipe/account"
	"os"
	"time"
)

var rootCmd = &cobra.Command{
	Use: "YPManager",

	Short: "YPManager -p [password] -a [address] -s [2006-02-21] -d 14",

	Long: `""`,

	Run: mainRun,

	//Args:  cobra.MinimumNArgs(2),
}
var param struct {
	password      string
	address       string
	interval      int
	startDay      string
	bootStrapIP   string
	bootStrapAddr string
	kingKey       string
	cipherText    string
}

func main() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var sysInitCmd = &cobra.Command{
	Use:   "init",
	Short: " YPManager init -p [password] -b [bootIP] -a [bootAddress].",
	Long:  `TODO::.`,
	Run:   initSys,
}

func init() {

	rootCmd.AddCommand(sysInitCmd)

	rootCmd.Flags().StringVarP(&param.password, "password",
		"p", "", "Thanos's finger")

	rootCmd.Flags().StringVarP(&param.address, "address",
		"u", "", "User's address")

	rootCmd.Flags().StringVarP(&param.startDay, "startDay",
		"s", "", "License start day")

	rootCmd.Flags().IntVarP(&param.interval, "duration", "d", 0,
		"license duration in days")

	sysInitCmd.Flags().StringVarP(&param.password, "password",
		"p", "", "Thanos's finger")

	sysInitCmd.Flags().StringVarP(&param.bootStrapIP, "bootIP",
		"i", "", "Boot strap server's IP")

	sysInitCmd.Flags().StringVarP(&param.bootStrapAddr, "bootID",
		"a", "", "Boot strap server's YouPipe account address")
}

func mainRun(_ *cobra.Command, _ []string) {

	thanosFinger := OpenThanosFinger(param.password)

	if !account.CheckID(param.address) {
		panic("user's address is invalid")
	}
	aid := account.ID(param.address)

	start := time.Now()
	if len(param.startDay) != 0 {
		s, err := time.Parse(SysTimeFormat, param.startDay)
		if err != nil {
			panic(err)
		}
		start = s
	}

	if param.interval <= 0 {
		panic("invalid duration no in days")
	}

	thanosFinger.Snap(aid, start, param.interval)
}

func initSys(_ *cobra.Command, _ []string) {
	thanosFinger := OpenThanosFinger(param.password)
	thanosFinger.CreateConfig(param.bootStrapIP, param.bootStrapAddr)
}
