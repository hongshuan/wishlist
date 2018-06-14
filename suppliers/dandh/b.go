package main

import "fmt"
import "os"
import "encoding/xml"

type PriceAvailRequest struct {
    XMLName xml.Name `xml:"XMLFORMPOST"`
    Request string `xml:"REQUEST"`

    Login struct {
        XMLName xml.Name `xml:"LOGIN"`
        UserId string `xml:"USERID"`
        Password string `xml:"PASSWORD"`
    }

    PartNums []string `xml:"PARTNUM"`
}

type Branch struct {
	Name        string `xml:"BRANCH"`
	Qty         string `xml:"QTY"`
	InStockDate string `xml:"INSTOCKDATE"`
}

type PriceAvailItem struct {
	PartNum   string   `xml:"PARTNUM"`
	UnitPrice string   `xml:"UNITPRICE"`
	Branches  []Branch `xml:"BRANCHQTY"`
	TotalQty  string   `xml:"TOTALQTY"`
}

type PriceAvailResponse struct {
    XMLName xml.Name `xml:"XMLRESPONSE"`
    Status  string   `xml:"STATUS"`
    Items []PriceAvailItem `xml:"ITEM"`
}

func main() {
	var v PriceAvailRequest

	v.Request = "price-availability"
	v.Login.UserId = "800712XML"
	v.Login.Password = "###########"
	v.PartNums = append(v.PartNums, "US16XGCA")
	v.PartNums = append(v.PartNums, "US16XGCB")

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	if err := enc.Encode(v); err != nil {
		fmt.Printf("error: %v\n", err)
	}

	var res PriceAvailResponse

	data := `<?xml version="1.0" encoding="UTF-8"?>
<XMLRESPONSE>
    <ITEM>
        <PARTNUM>US16XGCA</PARTNUM>
        <UNITPRICE>699.06</UNITPRICE>
        <BRANCHQTY>
            <BRANCH>Toronto</BRANCH>
            <QTY>1</QTY>
            <INSTOCKDATE></INSTOCKDATE>
        </BRANCHQTY>
        <BRANCHQTY>
            <BRANCH>Vancouver</BRANCH>
            <QTY>0</QTY>
            <INSTOCKDATE>08/10/18</INSTOCKDATE>
        </BRANCHQTY>
        <TOTALQTY>1</TOTALQTY>
    </ITEM>
    <STATUS>success</STATUS>
</XMLRESPONSE>
`
	err := xml.Unmarshal([]byte(data), &res)
	if err == nil {
		fmt.Printf("%#v\n", res)
	}
}
