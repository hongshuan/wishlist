package ingram

import (
    "fmt"
)

type Client struct {
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
