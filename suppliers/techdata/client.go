package techdata

import (
    "fmt"
    "wishlist/suppliers/common"
)

const PURPOSE_ORDER_DETAIL    = "01"
const PURPOSE_SHIPMENT_DETAIL = "03"
const PURPOSE_INVOICE_DETAIL  = "02"

const PA_TEST_URL = "http://tdxml.cstenet.com/xmlservlet"
const PA_PROD_URL = "https://tdxml.techdata.com/xmlservlet"

const PO_TEST_URL = "http://tdxml.cstenet.com/xmlservlet"
const PO_PROD_URL = "https://tdxml.techdata.com/xmlservlet"

const OS_PROD_URL = "https://tdxml.techdata.com/xmlservlet"

type Client struct {
    Config  common.TDConfig
}

func (c Client) GetPrefix() string {
    return "TD"
}

func (c Client) GetPriceAvail(sku string) {
    fmt.Println("TD GetPriceAvail", sku)
}

func (c Client) PurchaseOrder(info string) {
    fmt.Println("TD PurchaseOrder", info)
}

func (c Client) DropshipOrder(info string) {
    fmt.Println("TD DropshipOrder", info)
}
