package main

import (
    "fmt"

    "wishlist/suppliers/asi"
    "wishlist/suppliers/dandh"
    "wishlist/suppliers/synnex"
    "wishlist/suppliers/ingram"
    "wishlist/suppliers/techdata"
)

func main() {
    fmt.Println("Hello, wishlist\n")

    c1 := &asi.Client{}
    c1.GetPriceAvail("ASI-111")
    c1.PurchaseOrder("ASI-111")
    c1.DropshipOrder("ASI-111")
    fmt.Println()

    c2 := &dandh.Client{}
    c2.GetPriceAvail("DH-222")
    c2.PurchaseOrder("DH-222")
    c2.DropshipOrder("DH-222")
    fmt.Println()

    c3 := &synnex.Client{}
    c3.GetPriceAvail("SYN-333")
    c3.PurchaseOrder("SYN-333")
    c3.DropshipOrder("SYN-333")
    fmt.Println()

    c4 := &ingram.Client{}
    c4.GetPriceAvail("ING-444")
    c4.PurchaseOrder("ING-444")
    c4.DropshipOrder("ING-444")
    fmt.Println()

    c5 := &techdata.Client{}
    c5.GetPriceAvail("TD-555")
    c5.PurchaseOrder("TD-555")
    c5.DropshipOrder("TD-555")
    fmt.Println()
}
