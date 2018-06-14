package main

import "fmt"
import "os"
import "encoding/xml"

type PriceAvailRequest struct {
    XMLName xml.Name `xml:"XML_PriceAvailability_Submit"`

    Header struct {
        XMLName         xml.Name `xml:"Header"`
        Username        string   `xml:"UserName"`
        Password        string   `xml:"Password"`
        ResponseVersion string   `xml:"ResponseVersion"`
    }

    Detail struct {
        XMLName xml.Name `xml:"Detail"`
        LineInfo struct {
            XMLName   xml.Name `xml:"LineInfo"`
            RefIDQual string   `xml:"RefIDQual"`
            RefID     string   `xml:"RefID"`
        }
    }
}

type PriceAvailResponseHeader struct {
	TransSetIDCode  string `xml:"TransSetIDCode"`
	TransControlID  string `xml:"TransControlID"`
	ResponseVersion string `xml:"ResponseVersion"`
}

type PriceAvailResponseWarehouse struct {
	IDCode          string `xml:"IDCode"`
	Code            string `xml:"WhseCode"`
	Qty             string `xml:"Qty"`
	TotalOnOrderQty string `xml:"TotalOnOrderQty"`
	OnOrderQty      string `xml:"OnOrderQty"`
	OnOrderETADate  string `xml:"OnOrderETADate"`
}

type PriceAvailResponseLineInfo struct {
	RefIDQual1          string   `xml:"RefIDQual1"`
	RefID1              string   `xml:"RefID1"`
	RefIDQual2          string   `xml:"RefIDQual2"`
	RefID2              string   `xml:"RefID2"`
	RefIDQual4          string   `xml:"RefIDQual4"`
	RefID4              string   `xml:"RefID4"`
	ProductDesc         string   `xml:"ProductDesc"`
	PriceIDCode1        string   `xml:"PriceIDCode1"`
	UnitPrice1          string   `xml:"UnitPrice1"`
	PriceIDCode2        string   `xml:"PriceIDCode2"`
	UnitPrice2          string   `xml:"UnitPrice2"`
	RequiredEndUserInfo string   `xml:"RequiredEndUserInfo"`
	RequiredLicenseInfo string   `xml:"RequiredLicenseInfo"`
	ProductWeight       string   `xml:"ProductWeight"`
	ItemStatus          string   `xml:"ItemStatus"`

	Warehouses []PriceAvailResponseWarehouse `xml:"WhseInfo"`
}

type PriceAvailResponseDetail struct {
	LineInfo PriceAvailResponseLineInfo `xml:"LineInfo"`
}

type PriceAvailResponseSummary struct {
	NumOfSegments string `xml:"NbrOfSegments"`
}

type PriceAvailResponse struct {
    XMLName xml.Name `xml:"XML_PriceAvailability_Response"`

    Header  PriceAvailResponseHeader  `xml:"Header"`
    Detail  PriceAvailResponseDetail  `xml:"Detail"`
    Summary PriceAvailResponseSummary `xml:"Summary"`
}

func main() {
	var v PriceAvailRequest

	v.Header.Username = "567861"
	v.Header.Password = "############"
	v.Header.ResponseVersion = "1.4"

	v.Detail.LineInfo.RefIDQual = "VP"
	v.Detail.LineInfo.RefID = "7512ZJ"

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	if err := enc.Encode(v); err != nil {
		fmt.Printf("error: %v\n", err)
	}

	var res PriceAvailResponse

	data := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE XML_PriceAvailability_Response SYSTEM "XML_PriceAvailability_Response.dtd">
<XML_PriceAvailability_Response>
  <Header>
    <TransSetIDCode>846REC</TransSetIDCode>
    <TransControlID />
    <ResponseVersion>1.4</ResponseVersion>
  </Header>
  <Detail>
    <LineInfo>
      <RefIDQual1>VP</RefIDQual1>
      <RefID1>7512ZJ</RefID1>
      <RefIDQual2>MG</RefIDQual2>
      <RefID2>PA03610-B105</RefID2>
      <RefIDQual4>UP</RefIDQual4>
      <RefID4>097564308451</RefID4>
      <ProductDesc>SCANSNAP S1100I</ProductDesc>
      <PriceIDCode1>CON</PriceIDCode1>
      <UnitPrice1>$217.54</UnitPrice1>
      <PriceIDCode2>MSR</PriceIDCode2>
      <UnitPrice2>$255.00</UnitPrice2>
      <RequiredEndUserInfo>N</RequiredEndUserInfo>
      <RequiredLicenseInfo>N</RequiredLicenseInfo>
      <ProductWeight>1.600</ProductWeight>
      <ItemStatus>ACTIVE</ItemStatus>
      <WhseInfo>
        <IDCode>MISSISSAUGA, ON</IDCode>
        <WhseCode>A1</WhseCode>
        <Qty>-1</Qty>
        <TotalOnOrderQty>15</TotalOnOrderQty>
        <OnOrderQty>15</OnOrderQty>
        <OnOrderETADate>06/27/18</OnOrderETADate>
      </WhseInfo>
      <WhseInfo>
        <IDCode>RICHMOND, BC</IDCode>
        <WhseCode>A2</WhseCode>
        <Qty>1</Qty>
      </WhseInfo>
    </LineInfo>
  </Detail>
  <Summary>
    <NbrOfSegments>5</NbrOfSegments>
  </Summary>
</XML_PriceAvailability_Response>
`
	err := xml.Unmarshal([]byte(data), &res)
	if err == nil {
		fmt.Printf("%#v\n", res)
	}
}
