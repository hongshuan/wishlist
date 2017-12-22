package asi

import (
    "fmt"
    "wishlist/suppliers/common"
)

const PA_PROD_URL = "https://www.asipartner.com/partneraccess/xml/price.asp"
const PO_PROD_URL = "https://www.asipartner.com/partneraccess/xml/order.asp"
const OS_PROD_URL = "https://www.asipartner.com/partneraccess/xml/shipping.asp"
const IN_PROD_URL = "https://www.asipartner.com/partneraccess/xml/invoice.asp"

type Client struct {
    Config  common.ASIConfig
}

func (c Client) GetPrefix() string {
    return "AS"
}

func (c Client) GetPriceAvail(sku string) {
    fmt.Println("ASI GetPriceAvail", sku)
}

func (c Client) PurchaseOrder(info string) {
    fmt.Println("ASI PurchaseOrder", info)
}

func (c Client) DropshipOrder(info string) {
    fmt.Println("ASI DropshipOrder", info)
}
