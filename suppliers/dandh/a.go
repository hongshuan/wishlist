package main

import "fmt"
import "os"
import "encoding/xml"

type Login struct {
	UserId   string `xml:"USERID"`
	Password string `xml:"PASSWORD"`
}

type OrderHeader struct {
	OnlyBranch       string `xml:"ONLYBRANCH"`
	Branches         string `xml:"BRANCHES"`
	PartShipAllow    string `xml:"PARTSHIPALLOW"`
	BackOrderAllow   string `xml:"BACKORDERALLOW"`
	DropShipPW       string `xml:"DROPSHIPPW"`
	ShipToName       string `xml:"SHIPTONAME"`
	ShipToAttn       string `xml:"SHIPTOATTN"`
	ShipToAddress    string `xml:"SHIPTOADDRESS"`
	ShipToAddress2   string `xml:"SHIPTOADDRESS2"`
	ShipToCity       string `xml:"SHIPTOCITY"`
	ShipToProvince   string `xml:"SHIPTOPROVINCE"`
	ShipToPostalCode string `xml:"SHIPTOPOSTALCODE"`
	ShipCarrier      string `xml:"SHIPCARRIER"`
	ShipService      string `xml:"SHIPSERVICE"`
	Ponum            string `xml:"PONUM"`
	Remarks          string `xml:"REMARKS"`
}

type OrderItem struct {
	PartNum string   `xml:"PARTNUM"`
	Qty     string   `xml:"QTY"`
	Branch  string   `xml:"BRANCH"`
}

type OrderEntryRequest struct {
    XMLName xml.Name    `xml:"XMLFORMPOST"`
    Request string      `xml:"REQUEST"`
	Login   Login       `xml:"LOGIN"`
	Header  OrderHeader `xml:"ORDERHEADER"`
	Items   []OrderItem `xml:"ORDERITEMS>ITEM"`
}

type OrderEntryResponse struct {
	XMLName  xml.Name `xml:"XMLRESPONSE"`
	OrderNum string   `xml:"ORDERNUM"`
	Status   string   `xml:"STATUS"`
}

func main() {
	var v OrderEntryRequest

	v.Request = "orderEntry"
	v.Login.UserId = "800712XML"
	v.Login.Password = "###########"

	v.Header.OnlyBranch       = "Toronto"
	v.Header.Branches         = ""
	v.Header.PartShipAllow    = "N"
	v.Header.BackOrderAllow   = "N"
	v.Header.DropShipPW       = ""
	v.Header.ShipToName       = "Roy Zhang"
	v.Header.ShipToAttn       = "(Tel:905-480-0618)"
	v.Header.ShipToAddress    = "Unit 5, 270 Esna Park Dr."
	v.Header.ShipToAddress2   = ""
	v.Header.ShipToCity       = "Markham"
	v.Header.ShipToProvince   = "ON"
	v.Header.ShipToPostalCode = "L3R 1H3"
	v.Header.ShipCarrier      = "Purolator"
	v.Header.ShipService      = "Ground"
	v.Header.Ponum            = "W-1528899007-DH"
	v.Header.Remarks          = "Please HOLD"

	v.Items = append(v.Items, OrderItem{ "HYPERM2X16CARDC", "1", "Vancouver"})

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	if err := enc.Encode(v); err != nil {
		fmt.Printf("error: %v\n", err)
	}

	var res OrderEntryResponse

	data := `<?xml version="1.0" encoding="UTF-8"?>
<XMLRESPONSE>
  <ORDERNUM>3584591</ORDERNUM>
  <STATUS>success</STATUS>
</XMLRESPONSE>`

	err := xml.Unmarshal([]byte(data), &res)
	if err == nil {
		fmt.Printf("XMLName:  %#v\n", res.XMLName)
		fmt.Printf("OrderNum: %q\n",  res.OrderNum)
		fmt.Printf("Status:   %q\n",  res.Status)
	}
}
