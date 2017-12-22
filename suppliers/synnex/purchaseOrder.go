package synnex

/**
 * <SynnexB2B>
 *   <Credential>
 *     <UserID>*****************</UserID>
 *     <Password>***************</Password>
 *   </Credential>
 *   <OrderRequest>
 *     <CustomerNumber>1150897</CustomerNumber>
 *     <PONumber>702-1079188-0724236</PONumber>
 *     <DropShipFlag>Y</DropShipFlag>
 *     <Shipment>
 *       <ShipFromWarehouse>29</ShipFromWarehouse>
 *       <ShipTo>
 *         <AddressName1>Prashant vasudeva</AddressName1>
 *         <AddressLine1>1206-552 Wellington Street West</AddressLine1>
 *         <AddressLine2/>
 *         <City>Toronto</City>
 *         <State>ON</State>
 *         <ZipCode>M5V 2V5</ZipCode>
 *         <Country>CA</Country>
 *       </ShipTo>
 *       <ShipToContact>
 *         <ContactName>Prashant vasudeva</ContactName>
 *         <PhoneNumber>4168327984</PhoneNumber>
 *         <EmailAddress>ghpgpc54rfrkbfr@marketplace.amazon.ca</EmailAddress>
 *       </ShipToContact>
 *       <ShipMethod>
 *         <Code>UPG</Code>
 *       </ShipMethod>
 *     </Shipment>
 *     <Items>
 *       <Item lineNumber="1">
 *         <SKU>5707434</SKU>
 *         <UnitPrice>418.80</UnitPrice>
 *         <OrderQuantity>1</OrderQuantity>
 *       </Item>
 *     </Items>
 *   </OrderRequest>
 * </SynnexB2B>
 */

type PurchaseOrderRequest struct {
    XMLName     xml.Name `xml:""`
}

/**
 * <SynnexB2B>
 *   <OrderResponse>
 *     <CustomerNumber>1150897</CustomerNumber>
 *     <PONumber>702-1079188-0724236</PONumber>
 *     <Code>accepted</Code>
 *     <ResponseDateTime>2017-11-27T11:48:54</ResponseDateTime>
 *     <ResponseElapsedTime>1.399s</ResponseElapsedTime>
 *     <Items>
 *       <Item lineNumber="1">
 *         <SKU>5707434</SKU>
 *         <OrderQuantity>1</OrderQuantity>
 *         <Code>accepted</Code>
 *         <OrderNumber>43061107</OrderNumber>
 *         <OrderType>SO</OrderType>
 *         <ShipFromWarehouse/>
 *         <SynnexInternalReference>SALESORDER,A-CIS---SOLINE,A-CIS</SynnexInternalReference>
 *       </Item>
 *     </Items>
 *   </OrderResponse>
 * </SynnexB2B>
 */

type PurchaseOrderResponse struct {
    XMLName     xml.Name `xml:""`
}
