package main

import (
	"encoding/json"
	"github.com/btcsuite/btcutil/base58"
	"github.com/youpipe/go-youPipe/account"
	"golang.org/x/crypto/ed25519"
	"io/ioutil"
	"net"
	"time"
)

const (
	ConfFile      = ".finger"
	SysTimeFormat = "2006-01-02"
	Address       = "YP5rttHPzRsAe2RmF52sLzbBk4jpoPwJLtABaMv6qn7kVm"
	CipherText    = "347FrZuRaDL7dKGeG1fWzZuf2irc3qtXjxpSn762uNxHi8wBjTDongteyLvNDykbnTcXKokvhnvV3kMmnMP1RSYjRUwaGLAGVpkdfkx6CQWKiq"
)

type SysConf struct {
	bootStrapIP   string
	bootStrapAddr string
	kingKey       string
	cipherTxt     string
}

type ThanosFinger ed25519.PrivateKey

func OpenThanosFinger(password string) ThanosFinger {

	if len(param.password) == 0 {
		panic("please input king's account password")
	}

	acc := &account.Account{
		Address: Address,
		Key: &account.Key{
			LockedKey: base58.Decode(CipherText),
		},
	}

	if ok := acc.UnlockAcc(param.password); ok {
		panic("You're not Thanos")
	}

	return ThanosFinger(acc.Key.PriKey)
}

func (tf ThanosFinger) Snap(id account.ID, startDay time.Time, duration int) {

}

func (tf ThanosFinger) CreateConfig(ip, addr string) {

	if !account.CheckID(addr) {
		panic("boot strap server's YouPipe node address is invalid")
	}

	if ipAddr := net.ParseIP(ip); ipAddr == nil {
		panic("boot strap server's ip is invalid")
	}

	conf := SysConf{
		bootStrapAddr: addr,
		bootStrapIP:   ip,
		kingKey:       Address,
		cipherTxt:     CipherText,
	}

	data, err := json.Marshal(conf)
	if err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile(ConfFile, data, 0644); err != nil {
		panic(err)
	}
}
