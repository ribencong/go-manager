package main

import (
	"encoding/json"
	"github.com/btcsuite/btcutil/base58"
	"github.com/skip2/go-qrcode"
	"github.com/youpipe/go-manager/pbs"
	"github.com/youpipe/go-youPipe/account"
	"golang.org/x/crypto/ed25519"
	"time"
)

const (
	SysTimeFormat     = "2006-01-02"
	LicenseTimeFormat = "2006-01-02 15:04:05"

	//debug
	//Address    = "YPDYo3TZsLMgTHF9Vmm9arAWZHAuPTyh8XdF4MzRcqUjuT"
	//CipherText = "ffHuPBsZ7mMZ5m6XDnv66kq2fobLe1TACc4MqKjSY3ELSrvxmTwvmf6tfGsJqXFRN1fKEHZw5dnqdiHw484HiEGkcVXDXNwQhgprQr59NAVoe"

	//release
	Address    = "YP5rttHPzRsAe2RmF52sLzbBk4jpoPwJLtABaMv6qn7kVm"
	CipherText = "347FrZuRaDL7dKGeG1fWzZuf2irc3qtXjxpSn762uNxHi8wBjTDongteyLvNDykbnTcXKokvhnvV3kMmnMP1RSYjRUwaGLAGVpkdfkx6CQWKiq"
)

type ThanosFinger ed25519.PrivateKey

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

	return ThanosFinger(acc.Key.PriKey)
}

func (tf ThanosFinger) Snap(id string, startDay time.Time, duration int) string {

	endTime := startDay.Add(time.Hour * 24 * time.Duration(duration))

	ldata := &pbs.LicenseData{
		StartTime: startDay.Format(LicenseTimeFormat),
		EndTime:   endTime.Format(LicenseTimeFormat),
		UserAddr:  id,
	}
	data, err := json.Marshal(ldata)
	if err != nil {
		panic(err)
	}

	sig := ed25519.Sign(ed25519.PrivateKey(tf), data)

	l := &pbs.License{
		Data: ldata,
		Sig:  sig,
	}

	data, err = json.Marshal(l)

	err = qrcode.WriteFile(string(data), qrcode.Medium, 256, ldata.UserAddr+".png")
	if err != nil {
		panic(err)
	}
	return string(data)
}
