package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

type DHConfig struct {
    Username    string `json:"username"`
    Password    string `json:"password"`
    DropShipPW  string `json:"dropshippw"`
    PartShip    string `json:"partship"`
    BackOrder   string `json:"backorder"`
    ShipCarrier string `json:"shipcarrier"`
    ShipService string `json:"shipservice"`
    OnlyBranch  string `json:"onlybranch"`
    Branches    string `json:"branches"`
}

type SYNConfig struct {
    CustomerNo string `json:"customerNo"`
    AccountNo  string `json:"accountNo"`
    Username   string `json:"username"`
    Password   string `json:"password"`
    ShipMethod string `json:"shipmethod"`
}

type INGConfig struct {
    LoginId       string `json:"loginId"`
    Password      string `json:"password"`
    AutoRelease   string `json:"autoRelease"`
    CarrierCode   string `json:"carrierCode"`
    BackOrder     string `json:"backOrder"`
    SplitShipment string `json:"splitShipment"`
    SplitLine     string `json:"splitLine"`
}

type TDConfig struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type ASConfig struct {
    CID  string `json:"CID"`
    Cert string `json:"CERT"`
}

type Config struct {
    DH        DHConfig   `json:"DH"`
    SYNNEX    SYNConfig  `json:"SYNNEX"`
    INGRAM    INGConfig  `json:"INGRAM"`
    TECHDDATA TDConfig   `json:"TECHDATA"`
    ASI       ASConfig   `json:"ASI"`
}

func LoadConfig() Config {
    text, err := ioutil.ReadFile("./config.json")
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    var c Config
    json.Unmarshal(text, &c)
    return c
}
