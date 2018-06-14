package main

import "fmt"
import "os"
import "encoding/xml"

type PriceAvailRequest struct {
    XMLName xml.Name `xml:"PNARequest"`
    Version string `xml:"Version"`

    Header struct {
        XMLName xml.Name `xml:"TransactionHeader"`
        SenderID string `xml:"SenderID"`
        ReceiverID string `xml:"ReceiverID"`
        CountryCode string `xml:"CountryCode"`
        LoginID string `xml:"LoginID"`
        Password string `xml:"Password"`
        TransactionID string `xml:"TransactionID"`
    }

    Item struct {
        XMLName xml.Name `xml:"PNAInformation"`
		SKU     string   `xml:"SKU,attr"`
		Qty     string   `xml:"Quantity,attr"`
    }

    ShowDetail string `xml:"ShowDetail"`
}

type ErrorStatus struct {
	ErrorNumber string `xml:"ErrorNumber,attr"`
}

type TransactionHeader struct {
	SenderID      string      `xml:"SenderID"`
	ReceiverID    string      `xml:"ReceiverID"`
	ErrorStatus   ErrorStatus `xml:"ErrorStatus"`
	DocumentID    string      `xml:"DocumentID"`
	TransactionID string      `xml:"TransactionID"`
	TimeStamp     string      `xml:"TimeStamp"`
}

type Branch struct {
	ID       string `xml:"ID,attr"`
	Name     string `xml:"Name,attr"`
	Qty      string `xml:"Availability"`
	OnOrder  string `xml:"OnOrder"`
	ETADate  string `xml:"ETADate"`
}

type PriceAvailabilityItem struct {
	SKU string `xml:"SKU,attr"`
	Price string `xml:"Price"`
	SpecialPriceFlag string `xml:"SpecialPriceFlag"`
	ManufacturerPartNumber string `xml:"ManufacturerPartNumber"`
	ManufacturerPartNumberOccurs string `xml:"ManufacturerPartNumberOccurs"`
	VendorNumber string `xml:"VendorNumber"`
	Description string `xml:"Description"`
	ReserveInventoryFlag string `xml:"ReserveInventoryFlag"`
	AvailableRebQty string `xml:"AvailableRebQty"`

	Branches []Branch `xml:"Branch"`

	UPC string `xml:"UPC"`
	CustomerPartNumber string `xml:"CustomerPartNumber"`
}

type PriceAvailResponse struct {
    XMLName xml.Name `xml:"PNAResponse"`
    Version string `xml:"Version"`
    Header  TransactionHeader `xml:"TransactionHeader"`
    Item    PriceAvailabilityItem `xml:"PriceAndAvailability"`
}

func main() {
	var v PriceAvailRequest

	v.Version = "2.0"
	v.ShowDetail = "2"

	v.Header.SenderID = "ME"
	v.Header.ReceiverID = "YOU"
	v.Header.CountryCode = "FT"
	v.Header.LoginID = "TrEv8fEbes"
	v.Header.Password = "##########"
	v.Header.TransactionID = "1"

	v.Item.SKU = "78616Z"

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	if err := enc.Encode(v); err != nil {
		fmt.Printf("error: %v\n", err)
	}

	var res PriceAvailResponse

	data := `<?xml version="1.0" encoding="UTF-8"?>
<PNAResponse>
    <Version>2.0</Version>
    <TransactionHeader>
        <SenderID>YOU</SenderID>
        <ReceiverID>ME</ReceiverID>
        <ErrorStatus ErrorNumber=""></ErrorStatus>
        <DocumentID>{48F8D37C-21C1-4BA7-BD2D-3521E8DC3710}</DocumentID>
        <TransactionID>1</TransactionID>
        <TimeStamp>2018-06-14T09:46:32</TimeStamp>
    </TransactionHeader>
    <PriceAndAvailability SKU="78616Z" Quantity="">
        <Price>214.15</Price>
        <SpecialPriceFlag></SpecialPriceFlag>
        <ManufacturerPartNumber>PA03610-B105</ManufacturerPartNumber>
        <ManufacturerPartNumberOccurs></ManufacturerPartNumberOccurs>
        <VendorNumber>1937</VendorNumber>
        <Description>SCANSNAP S1100I CLR 600</Description>
        <ReserveInventoryFlag>Y</ReserveInventoryFlag>
        <AvailableRebQty>0</AvailableRebQty>
        <Branch Name="Vancouver" ID="10">
            <Availability>5</Availability>
            <OnOrder>0</OnOrder>
            <ETADate></ETADate>
        </Branch>
        <Branch Name="Toronto" ID="40">
            <Availability>0</Availability>
            <OnOrder>0</OnOrder>
            <ETADate></ETADate>
        </Branch>
        <UPC>0097564308451</UPC>
        <CustomerPartNumber></CustomerPartNumber>
    </PriceAndAvailability>
</PNAResponse>
`
	err := xml.Unmarshal([]byte(data), &res)
	if err == nil {
		fmt.Printf("%#v\n", res)
	} else {
		fmt.Println(err)
	}
}
