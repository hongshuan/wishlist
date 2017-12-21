package dandh

import (
    "fmt"
)

type Client struct {
}

func (c Client) GetPriceAvail(sku string) {
    fmt.Println("DH GetPriceAvail", sku)
}

func (c Client) PurchaseOrder(info string) {
    fmt.Println("DH PurchaseOrder", info)
}

func (c Client) DropshipOrder(info string) {
    fmt.Println("DH DropshipOrder", info)
}
