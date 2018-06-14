package main

import "fmt"
import "encoding/xml"

type PriceAvailRequest struct {
}

type Branch struct {
	Code string `xml:"Code,attr"`
	Name string `xml:"Name,attr"`
}

type PriceAvailResponseItem struct {
	SKU         string `xml:"ItemId"`
	Description string `xml:"Description"`
	Vendor      string `xml:"Vendor"`
	Category    string `xml:"Category"`
	SubCategory string `xml:"SubCategory"`
	UPC         string `xml:"UPC"`
	Price       string `xml:"Price"`
	Rebate      string `xml:"Rebate"`
	Term        string `xml:"Term"`
	Weight      string `xml:"Weight"`
	Status      string `xml:"Status"`

	Branches  []Branch `xml:"Qty>Branch"`
}

type PriceAvailResponse struct {
    XMLName   xml.Name `xml:"ASIInventory"`
	Timestamp string   `xml:"request,attr"`
	Excute    string   `xml:"excute,attr"`

	Item PriceAvailResponseItem `xml:"Inventory"`
}

func main() {
	var res PriceAvailResponse

	data := `<?xml version="1.0" encoding="UTF-8"?>
<ASIInventory request="06/14/2018 09:45" excute="812">
  <Inventory SKU="205784">
    <ItemId>910-004558</ItemId>
    <Description><![CDATA[Logitech Mouse 910-004558 G900 Wireless USB2.0 Full-speed Optical RGB Gaming Mouse Retail]]></Description>
    <Vendor>Logitech</Vendor>
    <Category>MC</Category>
    <SubCategory>99</SubCategory>
    <UPC>097855118424</UPC>
    <Price>216.08</Price>
    <Rebate><![CDATA[none]]></Rebate>
    <Term>none</Term>
    <Weight>3.8</Weight>
    <Status>ACTIVE</Status>
    <Qty>
      <Branch Code="2216" Name="Toronto">0</Branch>
      <Branch Code="3116" Name="Vancouver">0</Branch>
    </Qty>
  </Inventory>
</ASIInventory>

`
	err := xml.Unmarshal([]byte(data), &res)
	if err == nil {
		fmt.Printf("%#v\n", res)
	}
}
