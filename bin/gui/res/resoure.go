package res

import (
	"encoding/json"
	"io/ioutil"
	"path"
	"path/filepath"

	"fyne.io/fyne"
	"github.com/lengzhao/wallet/bin/gui/conf"
)

var localI18n map[string]string

var i18n = map[string]string{
	"GOVM":             "GOVM",
	"login.msg":        "Please enter password",
	"login.empty_pwd":  "not enter password",
	"transfer.peer":    "Peer",
	"transfer.desc":    "After the transfer, wait until the block index of the transaction is not 0 to be considered a success. Normally less than 2 minutes.",
	"move.from_chain":  "From Chain",
	"move.to_chain":    "To Chain",
	"move.desc":        "Move coin to other chain,need wait 10min after BlockID not equal 0.",
	"setting.unit":     "Coin Unit:",
	"setting.language": "Language:",
	"setting.server":   "API Server",
	"opCode.0":         "Transfer",
	"opCode.1":         "Move coins to other chain",
	"opCode.2":         "New Chain",
	"opCode.3":         "New App",
	"opCode.4":         "Run App",
	"opCode.5":         "Update App Life",
	"opCode.6":         "Register Miner",
	"setting.desc":     "Restart after setting",
	"vote_desc":        "After voting, collect the income yourself (every month through 0 vote)",
	"wgovm_desc":       "govm swap to eth.wGovm\n1. run this.Contract to lock coins.\n2. sign by admin(1 day)\n3. relayMint by yourself(http://govm.net:9090/wgovm.html)",
	"new_app_desc":     "create app",
}

var readFile func(fn string) ([]byte, error)

func assetRead(fn string) []byte {
	bytes, err := ioutil.ReadFile(path.Join("assets", fn))
	if err != nil {
		return nil
	}
	return bytes
}

// LoadLanguage load language
func LoadLanguage() error {
	lang := conf.Get().Langure
	data := assetRead(lang + ".json")
	if len(data) == 0 {
		return nil
	}
	err := json.Unmarshal(data, &localI18n)
	if err != nil {
		return err
	}
	return nil
}

// GetResource returns a resource by name
func GetResource(fn string) fyne.Resource {
	bytes := assetRead(fn)
	if len(bytes) == 0 {
		return nil
	}
	name := filepath.Base(fn)
	return fyne.NewStaticResource(name, bytes)
}

// GetResourceW get resource
func GetResourceW(fn string, fallback fyne.Resource) fyne.Resource {
	out := GetResource(fn)
	if out == nil {
		return fallback
	}
	return out
}

// GetLocalString get local string
func GetLocalString(id string) string {
	out := localI18n[id]
	if out != "" {
		return out
	}
	out = i18n[id]
	if out != "" {
		return out
	}
	return id
}

// GetBaseOfUnit get base of unit
func GetBaseOfUnit(in string) uint64 {
	switch in {
	case "tc":
		return 1000000000000
	case "t9":
		return 1000000000
	case "gvm":
		return 1000000000
	case "govm":
		return 1000000000
	case "t6":
		return 1000000
	case "t3":
		return 1000
	case "t0":
		return 1
	}
	return 1000000000
}
