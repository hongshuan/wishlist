package dandh

import (
    "encoding/xml"
)

type PriceAvailRequest struct {
    XMLName     xml.Name `xml:""`
}

/**
 * <XMLRESPONSE>
 *   <ITEM>
 *      <PARTNUM>52T</PARTNUM>
 *      <UNITPRICE>6.63</UNITPRICE>
 *      <BRANCHQTY>
 *          <BRANCH>Toronto</BRANCH>
 *          <QTY>84</QTY>
 *          <INSTOCKDATE />
 *      </BRANCHQTY>
 *      <TOTALQTY>84</TOTALQTY>
 *   </ITEM>
 *   <ITEM>
 *      <PARTNUM>123ABC</PARTNUM>
 *      <MESSAGE>Invalid Item Number</MESSAGE>
 *   </ITEM>
 *   <STATUS>success</STATUS>
 * </XMLRESPONSE>
 */

type PriceAvailResponse struct {
    XMLName xml.Name `xml:"XMLRESPONSE"`

    Item struct {
         PartNum   string `xml:"PARTNUM"`
         UnitPrice string `xml:"UNITPRICE"`
         Branches struct {
             Name        string `xml:"BRANCH"`
             Qty         string `xml:"QTY"`
             InStockDate string `xml:"INSTOCKDATE"`
         } `xml:"BRANCHQTY"`
         TotalQty string `xml:"TOTALQTY"`
         Message  string `xml:"MESSAGE"`
    } `xml:"ITEM"`

    Status  string `xml:"STATUS"`
}
