package main

import "fmt"
import "os"
import "encoding/xml"

type SkuItem struct {
	SKU        string `xml:"synnexSKU"`
	LineNumber string `xml:"lineNumber"`
}

type PriceAvailRequest struct {
    XMLName    xml.Name  `xml:"priceRequest"`
    CustomerNo string    `xml:"customerNo"`
    UserName   string    `xml:"userName"`
    Password   string    `xml:"password"`
    SkuList    []SkuItem `xml:"skuList"`
}

type AvailabilityByWarehouse struct {
	Number  string `xml:"warehouseInfo>number"`
	Zipcode string `xml:"warehouseInfo>zipcode"`
	City    string `xml:"warehouseInfo>city"`
	Addr    string `xml:"warehouseInfo>addr"`
	Qty     string `xml:"qty"`
}

type PriceAvailabilityItem struct {
	SKU string `xml:"synnexSKU"`
	MPN string `xml:"mfgPN"`
	Code string `xml:"mfgCode"`
	Status string `xml:"status"`
	Description string `xml:"description"`
	GlobalProductStatusCode string `xml:"GlobalProductStatusCode"`
	Price string `xml:"price"`
	TotalQuantity string `xml:"totalQuantity"`

	Warehouses []AvailabilityByWarehouse `xml:"AvailabilityByWarehouse"`

	LineNumber string `xml:"lineNumber"`
}

type PriceAvailResponse struct {
    XMLName    xml.Name `xml:"priceResponse"`
    CustomerNo string   `xml:"customerNo"`
    UserName   string   `xml:"userName"`

    Items []PriceAvailabilityItem `xml:"PriceAvailabilityList"`
}

func main() {
	var v PriceAvailRequest

	v.CustomerNo = "1150897"
	v.UserName = "roy@btecanada.com"
	v.Password = "###############"

	v.SkuList = append(v.SkuList, SkuItem{ SKU: "5725801", LineNumber: "1"})
	v.SkuList = append(v.SkuList, SkuItem{ SKU: "5725802", LineNumber: "2"})
	v.SkuList = append(v.SkuList, SkuItem{ SKU: "5725803", LineNumber: "3"})

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	if err := enc.Encode(v); err != nil {
		fmt.Printf("error: %v\n", err)
	}

	var res PriceAvailResponse

	data := `<?xml version="1.0" encoding="UTF-8"?>
<?xml version="1.0" encoding="UTF-8"?>
<priceResponse>
  <customerNo>1150897</customerNo>
  <userName>roy@btecanada.com</userName>
  <PriceAvailabilityList>
    <synnexSKU>5725800</synnexSKU>
    <mfgPN>051-044-NA-GE</mfgPN>
    <mfgCode>31266</mfgCode>
    <status>Active</status>
    <description>LEGENDARY WRLS HEADSET FOR PS4</description>
    <GlobalProductStatusCode>Active</GlobalProductStatusCode>
    <price>77.13</price>
    <totalQuantity>26</totalQuantity>
    <AvailabilityByWarehouse>
      <warehouseInfo>
      <number>29</number>
      <zipcode>N1H 1B4</zipcode>
      <city>Guelph, ON</city>
      <addr>107 Woodlawn Road West</addr></warehouseInfo>
      <qty>26</qty>
	</AvailabilityByWarehouse>
    <AvailabilityByWarehouse>
      <warehouseInfo>
      <number>31</number>
      <zipcode>T2C 4G6</zipcode>
      <city>Calgary, AB</city>
      <addr>5280 72 AVE SE</addr></warehouseInfo>
      <qty>0</qty>
	</AvailabilityByWarehouse>
    <AvailabilityByWarehouse>
      <warehouseInfo>
      <number>6</number>
      <zipcode>60446</zipcode>
      <city>Romeoville, IL</city>
      <addr>1180 Remington Blvd</addr></warehouseInfo>
      <qty>0</qty>
	</AvailabilityByWarehouse>
    <AvailabilityByWarehouse>
      <warehouseInfo>
      <number>60</number>
      <zipcode>V6W 1L8</zipcode>
      <city>Richmond, BC</city>
      <addr>Unit 120-18233 Blundell Road</addr></warehouseInfo>
      <qty>0</qty>
	</AvailabilityByWarehouse>
    <AvailabilityByWarehouse>
      <warehouseInfo>
      <number>7</number>
      <zipcode>38654</zipcode>
      <city>Olive Branch, MS</city>
      <addr>10381 Stateline Rd</addr></warehouseInfo>
      <qty>0</qty>
	</AvailabilityByWarehouse>
    <AvailabilityByWarehouse>
      <warehouseInfo>
      <number>98</number>
      <zipcode>00000</zipcode>
      <city>MFG Drop Shipped</city>
      <addr>MFG Drop Shipped</addr></warehouseInfo>
      <qty>0</qty>
	</AvailabilityByWarehouse>
    <lineNumber>1</lineNumber>
  </PriceAvailabilityList>
</priceResponse>

`
	err := xml.Unmarshal([]byte(data), &res)
	if err == nil {
		fmt.Printf("%#v\n", res)
	}
}
