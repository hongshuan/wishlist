package common

type DHConfig struct {
    Username      string `json:"username"`
    Password      string `json:"password"`
    DropShipPW    string `json:"dropshippw"`
    PartShip      string `json:"partship"`
    BackOrder     string `json:"backorder"`
    ShipCarrier   string `json:"shipcarrier"`
    ShipService   string `json:"shipservice"`
    OnlyBranch    string `json:"onlybranch"`
    Branches      string `json:"branches"`
}

type SYNConfig struct {
    CustomerNo    string `json:"customerNo"`
    AccountNo     string `json:"accountNo"`
    Username      string `json:"username"`
    Password      string `json:"password"`
    ShipMethod    string `json:"shipmethod"`
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
    Username      string `json:"username"`
    Password      string `json:"password"`
}

type ASIConfig struct {
    CID           string `json:"CID"`
    Cert          string `json:"CERT"`
}
