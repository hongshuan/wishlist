package common

const STATUS_OK    = "OK"
const STATUS_ERROR = "ERR"

type SupplierClient interface {
    GetPrefix() string
    GetPriceAvail(sku string)
    PurchaseOrder(info string)
    DropshipOrder(info string)
}

type PriceAvailBranch struct {
    Name     string
    Code     string
    Qty      string
}

type PriceAvailItem struct {
    Sku      string
    Status   string
    Price    string
    Branches []PriceAvailBranch
}

type PriceAvailResult struct {
    Status       string
    ErrorMessage string
    Items        []PriceAvailItem
}

type PurchaseOrderResult struct {
    Status       string
    OrderNo      string
    ErrorMessage string
}
