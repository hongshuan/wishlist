package common

type PriceAvailResult struct {
}

type SupplierClient interface {
    GetPriceAvail(sku string)
    PurchaseOrder(info string)
    DropshipOrder(info string)
}
