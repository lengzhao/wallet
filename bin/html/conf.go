package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"syscall"
	"time"

	"github.com/lengzhao/wallet/trans"
	"golang.org/x/crypto/ssh/terminal"
)

// TConfig config of app
type TConfig struct {
	Password    string `json:"password"`
	WalletFile  string `json:"wallet_file"`
	APIServer   string `json:"api_server"`
	StaticFiles string `json:"static_files"`
	HTTPPort    int    `json:"http_port"`
}

// Version Version
const Version = "v0.5.0"

var (
	confFile = "conf.json"
	conf     TConfig
	wallet   trans.Wallet
)

func getWorkDir() string {
	fullexecpath, err := os.Executable()
	if err != nil {
		return "."
	}

	dir, _ := filepath.Split(fullexecpath)
	return dir
}

func init() {
	confFile = path.Join(getWorkDir(), "conf.json")
	err := loadConfig()
	if err != nil {
		log.Println("fail to load config,", err)
		time.Sleep(time.Second * 3)
		os.Exit(2)
	}
}

func loadConfig() error {
	writeConf := false
	data, err := ioutil.ReadFile(confFile)
	if err == nil {
		err = json.Unmarshal(data, &conf)
		if err != nil {
			log.Println("fail to Unmarshal configure,conf.json")
			return err
		}
	} else {
		log.Println("not found config file,", confFile)
		writeConf = true
	}

	if conf.WalletFile == "" {
		conf.WalletFile = path.Join(getWorkDir(), "wallet.key")
	}

	if conf.Password == "" {
		fmt.Println("please enter the password of wallet:(default is govm_pwd@2019)")
		password, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return err
		}
		conf.Password = string(password)
		if conf.Password == "" {
			return fmt.Errorf("need password")
		}
	}
	if conf.APIServer == "" {
		conf.APIServer = "http://govm.top:9090"
	}
	if conf.StaticFiles == "" {
		conf.StaticFiles = path.Join(getWorkDir(), "static")
	}
	if _, err := os.Stat(conf.WalletFile); os.IsNotExist(err) {
		wallet.New(conf.Password)
		out := wallet.String()
		ioutil.WriteFile(conf.WalletFile, []byte(out), 666)
	} else {
		data, err = ioutil.ReadFile(conf.WalletFile)
		if err != nil {
			log.Println("fail to read wallet file:", conf.WalletFile, err)
			return err
		}
		err = wallet.Load(conf.Password, string(data))
		if err != nil {
			log.Println("fail to load wallet.", err)
			return err
		}
	}
	if writeConf {
		conf.Password = ""
		data, _ = json.MarshalIndent(conf, "", "  ")
		ioutil.WriteFile(confFile, data, 666)
	}
	return nil
}
