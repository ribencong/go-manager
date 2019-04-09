package main

import (
	"encoding/json"
	"github.com/btcsuite/btcutil/base58"
	"github.com/youpipe/go-youPipe/account"
	"golang.org/x/crypto/ed25519"
	"io/ioutil"
	"time"
)

const (
	ConfFile          = "finger.json"
	SysTimeFormat     = "2006-01-02"
	LicenseTimeFormat = "2006-01-02 15:04:05"

	//debug
	Address    = "YPDYo3TZsLMgTHF9Vmm9arAWZHAuPTyh8XdF4MzRcqUjuT"
	CipherText = "ffHuPBsZ7mMZ5m6XDnv66kq2fobLe1TACc4MqKjSY3ELSrvxmTwvmf6tfGsJqXFRN1fKEHZw5dnqdiHw484HiEGkcVXDXNwQhgprQr59NAVoe"

	//release
	//Address       = "YP5rttHPzRsAe2RmF52sLzbBk4jpoPwJLtABaMv6qn7kVm"
	//CipherText    = "347FrZuRaDL7dKGeG1fWzZuf2irc3qtXjxpSn762uNxHi8wBjTDongteyLvNDykbnTcXKokvhnvV3kMmnMP1RSYjRUwaGLAGVpkdfkx6CQWKiq"
)

type SysConf struct {
	KingKey   string `json:"KingKey"`
	CipherTxt string `json:"CipherTxt"`
}

type License struct {
	Signature string `json:"signature"`
	StartTime string `json:"start"`
	EndTime   string `json:"end"`
	Address   string `json:"user"`
}

type ThanosFinger ed25519.PrivateKey

var SysConfig = &SysConf{}

func OpenThanosFinger(password string) ThanosFinger {

	if len(password) == 0 {
		panic("please input king's account password")
	}

	acc := &account.Account{
		Address: Address,
		Key: &account.Key{
			LockedKey: base58.Decode(CipherText),
		},
	}

	if ok := acc.UnlockAcc(password); !ok {
		panic("You're not Thanos")
	}

	data, err := ioutil.ReadFile(ConfFile)
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(data, SysConfig); err != nil {
		panic(err)
	}

	return ThanosFinger(acc.Key.PriKey)
}

func (tf ThanosFinger) Snap(id string, startDay time.Time, duration int) *License {

	endTime := startDay.Add(time.Hour * 24 * time.Duration(duration))
	l := &License{
		StartTime: startDay.Format(LicenseTimeFormat),
		EndTime:   endTime.Format(LicenseTimeFormat),
		Address:   id,
	}
	data, err := json.Marshal(l)
	if err != nil {
		panic(err)
	}

	sig := ed25519.Sign(ed25519.PrivateKey(tf), data)
	l.Signature = base58.Encode(sig)

	return l
}
