package main

import "fmt"
import "os"
import "encoding/xml"

type TransactionHeader struct {
	SenderID      string `xml:"SenderID"`
	ReceiverID    string `xml:"ReceiverID"`
	CountryCode   string `xml:"CountryCode"`
	LoginID       string `xml:"LoginID"`
	Password      string `xml:"Password"`
	TransactionID string `xml:"TransactionID"`
}

type Address struct {
	Address1   string `xml:"ShipToAddress1"`
	Address2   string `xml:"ShipToAddress2"`
	Address3   string `xml:"ShipToAddress3"`
	City       string `xml:"ShipToCity"`
	Province   string `xml:"ShipToProvince"`
	PostalCode string `xml:"ShipToPostalCode"`
}

type ShipTo struct {
	Address Address  `xml:"Address"`
}

type AddressingInformation struct {
	CustomerPO      string `xml:"CustomerPO"`
	ShipToAttention string `xml:"ShipToAttention"`
	EndUserPO       string `xml:"EndUserPO"`
	ShipTo          ShipTo `xml:"ShipTo"`
}

type ShipmentOptions struct {
	BackOrderFlag     string `xml:"BackOrderFlag"`
	SplitShipmentFlag string `xml:"SplitShipmentFlag"`
	SplitLine         string `xml:"SplitLine"`
	ShipFromBranches  string `xml:"ShipFromBranches"`
	DeliveryDate      string `xml:"DeliveryDate"`
}

type ProcessingOptions struct {
	CarrierCode string `xml:"CarrierCode"`
	AutoRelease string `xml:"AutoRelease"`

	ShipmentOptions ShipmentOptions `xml:"ShipmentOptions"`
}

type DynamicMessage struct {
	MessageLines string `xml:"MessageLines"`
}

type OrderHeaderInformation struct {
	BillToSuffix string `xml:"BillToSuffix"`

	Addressing         AddressingInformation `xml:"AddressingInformation"`
	ProcessingOptions  ProcessingOptions     `xml:"ProcessingOptions"`
	DynamicMessage     DynamicMessage        `xml:"DynamicMessage"`
}

type ProductLine struct {
	SKU        string `xml:"SKU"`
	Quantity   string `xml:"Quantity"`
	LineNumber string `xml:"CustomerLineNumber"`
}

type CommentLine struct {
	Text string `xml:"CommentText"`
}

type OrderLineInformation struct {
	Product ProductLine `xml:"ProductLine"`
	Comment CommentLine `xml:"CommentLine"`
}

type OrderRequest struct {
    XMLName xml.Name `xml:"OrderRequest"`
    Version string   `xml:"Version"`

	Transaction TransactionHeader      `xml:"TransactionHeader"`
	OrderHeader OrderHeaderInformation `xml:"OrderHeaderInformation"`
	OrderLine   OrderLineInformation   `xml:"OrderLineInformation"`

    ShowDetail string `xml:"ShowDetail"`
}

type ErrorStatus struct {
    ErrorNumber string `xml:"ErrorNumber,attr"`
}

type OrderResponse struct {
    XMLName  xml.Name    `xml:"OrderResponse"`
    Status   ErrorStatus `xml:"TransactionHeader>ErrorStatus"`
	OrderNum string      `xml:"OrderInfo>OrderNumbers>BranchOrderNumber"`
}

func main() {

	var v OrderRequest

	v.Version = "2"
	v.Transaction.SenderID = ""
	v.Transaction.ReceiverID = ""
	v.Transaction.CountryCode = "FT"
	v.Transaction.LoginID = "TrEv8fEbes"
	v.Transaction.Password = "##########"
	v.Transaction.TransactionID = ""

	v.OrderHeader.BillToSuffix = ""

	v.OrderHeader.Addressing.CustomerPO = "70266665747063427"
	v.OrderHeader.Addressing.ShipToAttention = ""
	v.OrderHeader.Addressing.EndUserPO = "70266665747063427"
	v.OrderHeader.Addressing.ShipTo.Address.Address1 = "Elizabeth Batten"
	v.OrderHeader.Addressing.ShipTo.Address.Address2 = "General Delivery"
	v.OrderHeader.Addressing.ShipTo.Address.Address3 = ""
	v.OrderHeader.Addressing.ShipTo.Address.City = "Brigus Junction"
	v.OrderHeader.Addressing.ShipTo.Address.Province = "NL"
	v.OrderHeader.Addressing.ShipTo.Address.PostalCode = "A0B 1G0"

	v.OrderHeader.ProcessingOptions.CarrierCode = "PI"
	v.OrderHeader.ProcessingOptions.AutoRelease = "H"

	v.OrderHeader.ProcessingOptions.ShipmentOptions.BackOrderFlag     = "Y"
	v.OrderHeader.ProcessingOptions.ShipmentOptions.SplitShipmentFlag = "N"
	v.OrderHeader.ProcessingOptions.ShipmentOptions.SplitLine         = "N"
	v.OrderHeader.ProcessingOptions.ShipmentOptions.ShipFromBranches  = "10"
	v.OrderHeader.ProcessingOptions.ShipmentOptions.DeliveryDate      = ""

	v.OrderHeader.DynamicMessage.MessageLines = ""

	v.OrderLine.Product.SKU = "8410EB"
	v.OrderLine.Product.Quantity = "1"
	v.OrderLine.Product.LineNumber = ""

	v.OrderLine.Comment.Text = ""

	v.ShowDetail = "2"

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	if err := enc.Encode(v); err != nil {
		fmt.Printf("error: %v\n", err)
	}

	var res OrderResponse

	//TODO: encoding ??
	//data := `<?xml version="1.0" encoding="ISO-8859-1"?>
	data := `<?xml version="1.0"?>
<OrderResponse>
  <TransactionHeader>
    <ErrorStatus ErrorNumber=""/>
  </TransactionHeader>
  <OrderInfo>
    <OrderNumbers>
      <BranchOrderNumber>4011852</BranchOrderNumber>
      <CustomerPO>70266665747063427</CustomerPO>
    </OrderNumbers>
  </OrderInfo>
</OrderResponse>

`
	err := xml.Unmarshal([]byte(data), &res)
	if err == nil {
		fmt.Printf("OrderNum: %q\n",  res.OrderNum)
		fmt.Printf("Status:   %q\n",  res.Status.ErrorNumber)
	} else {
		fmt.Println(err)
	}
}
