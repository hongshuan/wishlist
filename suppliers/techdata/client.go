package techdata

import (
    "fmt"
    "wishlist/suppliers/common"
)

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
