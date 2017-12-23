package dandh

import (
    "encoding/xml"
)

/**
 * <?xml version="1.0"?>
 * <XMLFORMPOST>
 *   <REQUEST>orderEntry</REQUEST>
 *   <LOGIN>
 *     <USERID>*********</USERID>
 *     <PASSWORD>***********</PASSWORD>
 *   </LOGIN>
 *   <ORDERHEADER>
 *     <ONLYBRANCH>Toronto</ONLYBRANCH>
 *     <BRANCHES/>
 *     <PARTSHIPALLOW>N</PARTSHIPALLOW>
 *     <BACKORDERALLOW>N</BACKORDERALLOW>
 *     <DROPSHIPPW/>
 *     <SHIPTONAME>Jennifer Danjal</SHIPTONAME>
 *     <SHIPTOATTN>(Tel:506-386-2206)</SHIPTOATTN>
 *     <SHIPTOADDRESS>548 Steiger way</SHIPTOADDRESS>
 *     <SHIPTOADDRESS2/>
 *     <SHIPTOCITY>Saskatoon</SHIPTOCITY>
 *     <SHIPTOPROVINCE>SK</SHIPTOPROVINCE>
 *     <SHIPTOPOSTALCODE>S7N 4k2</SHIPTOPOSTALCODE>
 *     <SHIPCARRIER>Purolator</SHIPCARRIER>
 *     <SHIPSERVICE>Ground</SHIPSERVICE>
 *     <PONUM>195501401-A</PONUM>
 *     <REMARKS>Drop ship</REMARKS>
 *   </ORDERHEADER>
 *   <ORDERITEMS>
 *     <ITEM>
 *       <PARTNUM>XBS90102PNK</PARTNUM>
 *       <QTY>1</QTY>
 *       <BRANCH>Toronto</BRANCH>
 *     </ITEM>
 *   </ORDERITEMS>
 * </XMLFORMPOST>
 */

type PurchaseOrderRequest struct {
    XMLName  xml.Name `xml:"XMLFORMPOST"`
    Request  string   `xml:"REQUEST"`

    UserId   string   `xml:"LOGIN>USERID"`
    Password string   `xml:"LOGIN>PASSWORD"`

    Header struct {
        OnlyBranch       string `xml:"ONLYBRANCH"`
        Branches         string `xml:"BRANCHES"`
        PartShipAllow    string `xml:"PARTSHIPALLOW"`
        BackOrderAllow   string `xml:"BACKORDERALLOW"`
        DropshipPW       string `xml:"DROPSHIPPW"`
        ShipToName       string `xml:"SHIPTONAME"`
        ShipToAttn       string `xml:"SHIPTOATTN"`
        ShipToAddress    string `xml:"SHIPTOADDRESS"`
        ShipToAddress2   string `xml:"SHIPTOADDRESS2"`
        ShipToCity       string `xml:"SHIPTOCITY"`
        ShipToProvince   string `xml:"SHIPTOPROVINCE"`
        ShipToPostalCode string `xml:"SHIPTOPOSTALCODE"`
        ShipCarrier      string `xml:"SHIPCARRIER"`
        ShipService      string `xml:"SHIPSERVICE"`
        PoNum            string `xml:"PONUM"`
        Remarks          string `xml:"REMARKS"`
    } `xml:"ORDERHEADER"`

    Items struct {
        PartNum string `xml:"PARTNUM"`
        Qty     string `xml:"QTY"`
        Branch  string `xml:"BRANCH"`
    } `xml:"ORDERITEMS>ITEM"`
}

/**
 * <XMLRESPONSE>
 *   <ORDERNUM>3257770</ORDERNUM>
 *   <STATUS>success</STATUS>
 * </XMLRESPONSE>
 */

type PurchaseOrderResponse struct {
    XMLName     xml.Name `xml:"XMLRESPONSE"`
    OrderNum    string   `xml:"ORDERNUM"`
    Status      string   `xml:"STATUS"`
}
