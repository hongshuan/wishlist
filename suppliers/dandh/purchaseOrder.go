package dandh

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
    XMLName     xml.Name `xml:""`
}

/**
 * <XMLRESPONSE>
 *   <ORDERNUM>3257770</ORDERNUM>
 *   <STATUS>success</STATUS>
 * </XMLRESPONSE>
 */

type PurchaseOrderResponse struct {
    XMLName     xml.Name `xml:""`
}
