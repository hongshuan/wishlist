package asi

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
    XMLName     xml.Name `xml:""`
}

/**
 * <ASIOrderReply>
 *   <order>
 *     <orderid>11622239</orderid>
 *   </order>
 * </ASIOrderReply>
 */
type PurchaseOrderResponse struct {
    XMLName     xml.Name `xml:""`
}
