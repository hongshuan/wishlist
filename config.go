package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"

    "wishlist/suppliers/common"
)

type Config struct {
    DH   common.DHConfig   `json:"DH"`
    SYN  common.SYNConfig  `json:"SYNNEX"`
    ING  common.INGConfig  `json:"INGRAM"`
    TD   common.TDConfig   `json:"TECHDATA"`
    ASI  common.ASIConfig  `json:"ASI"`
}

func LoadConfig() Config {
    raw, err := ioutil.ReadFile("./config.json")
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    var cfg Config
    json.Unmarshal(raw, &cfg)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    //fmt.Printf("%+v\n", cfg)
    //fmt.Println(cfg.DH.ShipCarrier)

    return cfg
}
