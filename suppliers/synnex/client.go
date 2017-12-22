package synnex

import (
    "fmt"
    "wishlist/suppliers/common"
)

const PA_TEST_URL = "https://testec.synnex.ca/SynnexXML/PriceAvailability"
const PA_PROD_URL = "https://ec.synnex.ca/SynnexXML/PriceAvailability"

const PO_TEST_URL = "https://testec.synnex.ca/SynnexXML/PO"
const PO_PROD_URL = "https://ec.synnex.ca/SynnexXML/PO"

type Client struct {
    Config  common.SYNConfig
}

func (c Client) GetPrefix() string {
    return "SYN"
}

func (c Client) GetPriceAvail(sku string) {
    fmt.Println("SYN GetPriceAvail", sku)
}

func (c Client) PurchaseOrder(info string) {
    fmt.Println("SYN PurchaseOrder", info)
}

func (c Client) DropshipOrder(info string) {
    fmt.Println("SYN DropshipOrder", info)
}
