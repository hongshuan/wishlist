package ingram

import (
    "encoding/xml"
)

/**
 * <OrderRequest>
 *   <Version>2.0</Version>
 *   <TransactionHeader>
 *     <SenderID/>
 *     <ReceiverID/>
 *     <CountryCode>FT</CountryCode>
 *     <LoginID>**********</LoginID>
 *     <Password>**********</Password>
 *     <TransactionID/>
 *   </TransactionHeader>
 *   <OrderHeaderInformation>
 *     <BillToSuffix/>
 *     <AddressingInformation>
 *       <CustomerPO>70240827725635450</CustomerPO>
 *       <ShipToAttention/>
 *       <EndUserPO>70240827725635450</EndUserPO>
 *       <ShipTo>
 *         <Address>
 *           <ShipToAddress1>Olaoluwa Adurogbangb</ShipToAddress1>
 *           <ShipToAddress2>46 Bevan Crescent</ShipToAddress2>
 *           <ShipToAddress3/>
 *           <ShipToCity>Ajax</ShipToCity>
 *           <ShipToProvince>ON</ShipToProvince>
 *           <ShipToPostalCode>L1T 4P3</ShipToPostalCode>
 *         </Address>
 *       </ShipTo>
 *     </AddressingInformation>
 *     <ProcessingOptions>
 *       <CarrierCode>PI</CarrierCode>
 *       <AutoRelease>H</AutoRelease>
 *       <ShipmentOptions>
 *         <BackOrderFlag>Y</BackOrderFlag>
 *         <SplitShipmentFlag>N</SplitShipmentFlag>
 *         <SplitLine>N</SplitLine>
 *         <ShipFromBranches>40</ShipFromBranches>
 *         <DeliveryDate/>
 *       </ShipmentOptions>
 *     </ProcessingOptions>
 *     <DynamicMessage>
 *       <MessageLines/>
 *     </DynamicMessage>
 *   </OrderHeaderInformation>
 *   <OrderLineInformation>
 *     <ProductLine>
 *       <SKU>1116DU</SKU>
 *       <Quantity>1</Quantity>
 *       <CustomerLineNumber/>
 *     </ProductLine>
 *     <CommentLine>
 *       <CommentText/>
 *     </CommentLine>
 *   </OrderLineInformation>
 *   <ShowDetail>2</ShowDetail>
 * </OrderRequest>
 */

type PurchaseOrderRequest struct {
    XMLName     xml.Name `xml:"OrderRequest"`
}

/**
 * <OrderResponse>
 *   <Version>2.0</Version>
 *   <TransactionHeader>
 *     <SenderID/>
 *     <ReceiverID/>
 *     <ErrorStatus ErrorNumber=""/>
 *     <DocumentID>{567BBEE4-CB50-489F-8E54-04F445760130}</DocumentID>
 *     <TransactionID></TransactionID>
 *     <TimeStamp>2017-11-27T08:53:27</TimeStamp>
 *   </TransactionHeader>
 *   <OrderInfo>
 *     <OrderNumbers>
 *       <BranchOrderNumber>4047227</BranchOrderNumber>
 *       <CustomerPO>70240827725635450</CustomerPO>
 *       <ThirdPartyFreightAccount> </ThirdPartyFreightAccount>
 *       <ShipToAddress1>OLAOLUWA ADUROGBANGB</ShipToAddress1>
 *       <ShipToAddress2>46 BEVAN CRESCENT</ShipToAddress2>
 *       <ShipToCity>AJAX</ShipToCity>
 *       <ShipToProvince>ON</ShipToProvince>
 *       <ShipToPostalCode>L1T4P3</ShipToPostalCode>
 *       <AddressErrorMessage AddressErrorType="  "></AddressErrorMessage>
 *       <ContractNumber></ContractNumber>
 *       <OrderSuffix Suffix="11">
 *         <DistributionWeight>1</DistributionWeight>
 *         <SuffixErrorResponse SuffixErrorType=""/>
 *         <Carrier CarrierCode="PI">PUROLATOR GROUND</Carrier>
 *         <LineInformation>
 *           <ProductLine>
 *             <LineError/>
 *             <SKU>1116DU</SKU>
 *             <ManufacturerPartNumber>WL3-00072</ManufacturerPartNumber>
 *             <UPC>0889842196177</UPC>
 *             <UnitPrice>57.53</UnitPrice>
 *             <IngramLineNumber>001</IngramLineNumber>
 *             <CustomerLineNumber>000</CustomerLineNumber>
 *             <ShipFromBranch>40</ShipFromBranch>
 *             <OrderQuantity>1</OrderQuantity>
 *             <AllocatedQuantity>1</AllocatedQuantity>
 *             <BackOrderedQuantity>0</BackOrderedQuantity>
 *             <BackOrderETADate></BackOrderETADate>
 *             <PriceDerivedFlag/>
 *             <ForeignCurrency>0.00</ForeignCurrency>
 *             <FreightRate>0.00</FreightRate>
 *             <TransitDays>0</TransitDays>
 *             <LineBillToSuffix>000</LineBillToSuffix>
 *           </ProductLine>
 *         </LineInformation>
 *       </OrderSuffix>
 *     </OrderNumbers>
 *   </OrderInfo>
 * </OrderResponse>
 */

type PurchaseOrderResponse struct {
    XMLName     xml.Name `xml:"OrderResponse"`
}
