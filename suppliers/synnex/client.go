package synnex

import (
    "fmt"
    "wishlist/suppliers/common"
)

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
