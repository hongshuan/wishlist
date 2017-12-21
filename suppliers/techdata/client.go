package techdata

import (
    "fmt"
)

type Client struct {
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
