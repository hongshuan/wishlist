package main

import "fmt"
import "os"
import "encoding/xml"

type ShipTo struct {
	Name     string `xml:"name"`
	Address1 string `xml:"address1"`
	Address2 string `xml:"address2"`
	Address3 string `xml:"address3"`
	City     string `xml:"city"`
	State    string `xml:"state"`
	Zip      string `xml:"zip"`
	Country  string `xml:"country"`
}

type ShipInfo struct {
	Via          string `xml:"via"`
	Instructions string `xml:"instructions"`
}

type Header struct {
	ShipTo   ShipTo   `xml:"shipto"`
	ShipInfo ShipInfo `xml:"shipinfo"`
}

type Line struct {
	Sku    string `xml:"sku"`
	Qty    string `xml:"qty"`
	Price  string `xml:"price"`
	Branch string `xml:"branch"`
}

type Detail struct {
	Line Line `xml:"line"`
}

type OrderRequest struct {
    XMLName xml.Name `xml:"ASIOrderRequest"`

	CustId string `xml:"custid,attr"`
	CustPo string `xml:"custpo,attr"`
	Cert   string `xml:"cert,attr"`

	Header Header `xml:"header"`
	Detail Detail `xml:"detail"`
}

type Error struct {
	Code    string `xml:"code"`
	Message string `xml:"message"`
}

type OrderReply struct {
    XMLName xml.Name `xml:"ASIOrderReply"`
    OrderNum string  `xml:"order>orderid"`
	Error    Error   `xml:"error"`
}

func main() {
	var v OrderRequest

	v.CustId = "75692"
	v.CustPo = "702-8551087-2974644"
	v.Cert   = "###############"

	v.Header.ShipTo.Name     = "Devin Giftopoulos"
	v.Header.ShipTo.Address1 = "2955 Regional Rd 12"
	v.Header.ShipTo.Address2 = ""
	v.Header.ShipTo.Address3 = "2899216585"
	v.Header.ShipTo.City     = "Grassie"
	v.Header.ShipTo.State    = "ON"
	v.Header.ShipTo.Zip      = "L0R 1M0"
	v.Header.ShipTo.Country  = "CA"

	v.Header.ShipInfo.Via = "PGD"
	v.Header.ShipInfo.Instructions = ""

	v.Detail.Line.Sku    = "209770"
	v.Detail.Line.Qty    = "1"
	v.Detail.Line.Price  = "216.61"
	v.Detail.Line.Branch = "2216"

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	if err := enc.Encode(v); err != nil {
		fmt.Printf("error: %v\n", err)
	}

	var res1 OrderReply

	data := `<?xml version="1.0" encoding="UTF-8"?>
<?xml version="1.0" encoding="utf-8"?>
<ASIOrderReply>
  <order>
    <orderid>11868789</orderid>
  </order>
</ASIOrderReply>
`
	err := xml.Unmarshal([]byte(data), &res1)
	if err == nil {
		fmt.Printf("OrderNum: %q\n",  res1.OrderNum)
		fmt.Printf("Error:   %q\n",   res1.Error)
	}

	var res2 OrderReply

	data = `<?xml version="1.0" encoding="UTF-8"?>
<ASIOrderReply>
  <error>
    <code>502</code>
    <message>Your certificate is invalid</message>
  </error>
</ASIOrderReply>
`
	err = xml.Unmarshal([]byte(data), &res2)
	if err == nil {
		fmt.Printf("OrderNum: %q\n",  res2.OrderNum)
		fmt.Printf("Error:   %q\n",   res2.Error)
	}
}
