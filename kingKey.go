package main

import (
	"encoding/json"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ribencong/go-youPipe/account"
	"github.com/ribencong/go-youPipe/service"
	"golang.org/x/crypto/ed25519"
	"time"
)

const (
	SysTimeFormat = "2006-01-02"
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

	content := &service.LicenseData{
		StartDate: service.JsonTime(startDay),
		EndDate:   service.JsonTime(endTime),
		UserAddr:  id,
	}
	data, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	sig := ed25519.Sign(ed25519.PrivateKey(tf), data)
	l := &service.License{
		Signature:   sig,
		LicenseData: content,
	}

	data, err = json.Marshal(l)
	return string(data)
}
