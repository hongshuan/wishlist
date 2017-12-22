package dandh

import (
    "fmt"
    "wishlist/suppliers/common"
)

const PA_PROD_URL = "https://www.dandh.ca/dhXML/xmlDispatch"
const PO_PROD_URL = "https://www.dandh.ca/dhXML/xmlDispatch"
const OS_PROD_URL = "https://www.dandh.ca/dhXML/xmlDispatch"

type Client struct {
    Config  common.DHConfig
}

func (c Client) GetPrefix() string {
    return "DH"
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
