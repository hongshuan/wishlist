package asi

import (
    "encoding/xml"
)

/**
 * <ASIOrderRequest custid="*****" custpo="701-8531915-6574603" cert="***************">
 *   <header>
 *     <shipto>
 *       <name>Jim Parker</name>
 *       <address1>205 Hammerstone Cr.</address1>
 *       <address2/>
 *       <address3>905-709-2093</address3>
 *       <city>Thornhill</city>
 *       <state>ON</state>
 *       <zip>L4J 8A9</zip>
 *       <country>CA</country>
 *     </shipto>
 *     <shipinfo>
 *       <via>PGD</via>
 *       <instructions/>
 *     </shipinfo>
 *   </header>
 *   <detail>
 *     <line>
 *       <sku>204816</sku>
 *       <qty>2</qty>
 *       <price>829.65</price>
 *       <branch>2216</branch>
 *     </line>
 *   </detail>
 * </ASIOrderRequest>
 */

type PurchaseOrderRequest struct {
    XMLName xml.Name `xml:"ASIOrderRequest"`
    CustId  string   `xml:"custid,attr"`
    CustPo  string   `xml:"custpo,attr"`
    Cert    string   `xml:"cert,attr"`

    Header struct {
        ShipTo struct {
            Name     string `xml:"name"`
            Address1 string `xml:"address1"`
            Address2 string `xml:"address2"`
            Address3 string `xml:"address3"`
            City     string `xml:"city"`
            State    string `xml:"state"`
            Zip      string `xml:"zip"`
            Country  string `xml:"country"`
        } `xml:"shipto"`

        ShipInfo struct {
            Via          string `xml:"via"`
            Instructions string `xml:"instructions"`
        } `xml:"shipinfo"`
    } `xml:"header"`

    Detail struct {
        Sku     string  `xml:"line>sku"`
        Qty     string  `xml:"line>qty"`
        Price   string  `xml:"line>price"`
        Branch  string  `xml:"line>branch"`
    } `xml:"detail"`
}

/**
 * <ASIOrderReply>
 *   <order>
 *     <orderid>11622239</orderid>
 *   </order>
 * </ASIOrderReply>
 */
type PurchaseOrderResponse struct {
    XMLName     xml.Name `xml:"ASIOrderReply"`
    OrderId     string   `xml:"order>orderid"`
}
