package main

import "fmt"
import "os"
import "encoding/xml"

type Credential struct {
	UserID string `xml:"UserID"`
	Password string `xml:"Password"`
}

type ShipMethod struct {
	Code string `xml:"Code"`
}

type ShipToContact struct {
	Name         string `xml:"ContactName"`
	PhoneNumber  string `xml:"PhoneNumber"`
	EmailAddress string `xml:"EmailAddress"`
}

type ShipTo struct {
	AddressName  string `xml:"AddressName1"`
	AddressLine1 string `xml:"AddressLine1"`
	AddressLine2 string `xml:"AddressLine2"`
	City         string `xml:"City"`
	State        string `xml:"State"`
	ZipCode      string `xml:"ZipCode"`
	Country      string `xml:"Country"`
}

type Shipment struct {
	FromWarehouse string        `xml:"ShipFromWarehouse"`
	ShipTo        ShipTo        `xml:"ShipTo"`
	Contact       ShipToContact `xml:"ShipToContact"`
	ShipMethod    ShipMethod    `xml:"ShipMethod"`
}

type Item struct {
	LineNumber    string    `xml:"lineNumber,attr"`
	SKU           string    `xml:"SKU"`
	UnitPrice     string    `xml:"UnitPrice"`
	Quantity      string    `xml:"OrderQuantity"`
}

type Detail struct {
	CustomerNumber string   `xml:"CustomerNumber"`
	PONumber       string   `xml:"PONumber"`
	DropShipFlag   string   `xml:"DropShipFlag"`
	Shipment       Shipment `xml:"Shipment"`
	Comment        string   `xml:"Comment"`
	Items          []Item   `xml:"Items>Item"`
}

type OrderRequest struct {
    XMLName    xml.Name   `xml:"SynnexB2B"`
    Credential Credential `xml:"Credential"`
	Detail     Detail     `xml:"OrderRequest"`
}

type OrderResponse struct {
    XMLName    xml.Name   `xml:"SynnexB2B"`
    Status     string     `xml:"OrderResponse>Items>Item>Code"`
	OrderNum   string     `xml:"OrderResponse>Items>Item>OrderNumber"`
}

func main() {
	var v OrderRequest

	v.Credential.UserID = "roy@btecanada.com"
	v.Credential.Password = "###############"

	v.Detail.CustomerNumber = "1150897"
	v.Detail.PONumber       = "702-7479202-2375455"
	v.Detail.DropShipFlag   = "Y"

	v.Detail.Shipment.FromWarehouse = "29"

	v.Detail.Shipment.ShipTo.AddressName  = "George Vatcher"
	v.Detail.Shipment.ShipTo.AddressLine1 = "3075 rue Paul-David Apt 439"
	v.Detail.Shipment.ShipTo.AddressLine2 = ""
	v.Detail.Shipment.ShipTo.City         = "Montreal"
	v.Detail.Shipment.ShipTo.State        = "QC"
	v.Detail.Shipment.ShipTo.ZipCode      = "H1N 0A8"
	v.Detail.Shipment.ShipTo.Country      = "CA"

	v.Detail.Shipment.Contact.Name         = "George Vatcher"
	v.Detail.Shipment.Contact.PhoneNumber  = "514-725-4476"
	v.Detail.Shipment.Contact.EmailAddress = "84s3v7g2w57y2z8@marketplace.amazon.ca"

	v.Detail.Shipment.ShipMethod.Code = "PUX"

	v.Detail.Comment = "match ing's $7.38"
	v.Detail.Items = append(v.Detail.Items, Item{
		LineNumber: "1",
		SKU: "4627972",
		UnitPrice: "7.71",
		Quantity: "1",
	})

	v.Detail.Items = append(v.Detail.Items, Item{
		LineNumber: "2",
		SKU: "4627973",
		UnitPrice: "8.81",
		Quantity: "2",
	})

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	if err := enc.Encode(v); err != nil {
		fmt.Printf("error: %v\n", err)
	}

	var res1 OrderResponse

	data := `<?xml version="1.0" encoding="UTF-8"?>
<SynnexB2B>
  <OrderResponse>
    <CustomerNumber>1150897</CustomerNumber>
    <PONumber>702-7479202-2375455</PONumber>
    <Code>accepted</Code>
    <ResponseDateTime>2018-06-13T12:11:33</ResponseDateTime>
    <ResponseElapsedTime>1.627s</ResponseElapsedTime>
    <Items>
      <Item lineNumber="1">
        <SKU>4627972</SKU>
        <OrderQuantity>1</OrderQuantity>
        <Code>accepted</Code>
        <OrderNumber>23821269</OrderNumber>
        <OrderType>99</OrderType>
        <ShipFromWarehouse/>
        <SynnexInternalReference>OKCOMMENTS,A-SALESPND---CONVERTOK,A-SALESPND</SynnexInternalReference>
      </Item>
    </Items>
  </OrderResponse>
</SynnexB2B>
`
	err := xml.Unmarshal([]byte(data), &res1)
	if err == nil {
		fmt.Printf("OrderNum: %q\n",  res1.OrderNum)
		fmt.Printf("Error:   %q\n",   res1.Status)
	}
}
