package asi

import (
    "fmt"
)

type Client struct {
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
