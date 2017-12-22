package ingram

import (
    "fmt"
    "wishlist/suppliers/common"
)

const PA_PROD_URL = "https://newport.ingrammicro.com/mustang"
const PO_PROD_URL = "https://newport.ingrammicro.com/mustang"
const OS_PROD_URL = "https://newport.ingrammicro.com/mustang"

type Client struct {
    Config  common.INGConfig
}

func (c Client) GetPrefix() string {
    return "ING"
}

func (c Client) GetPriceAvail(sku string) {
    fmt.Println("ING GetPriceAvail", sku)
}

func (c Client) PurchaseOrder(info string) {
    fmt.Println("ING PurchaseOrder", info)
}

func (c Client) DropshipOrder(info string) {
    fmt.Println("ING DropshipOrder", info)
}
