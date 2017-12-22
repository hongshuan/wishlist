package ingram

import (
    "encoding/xml"
)

/**
 * <PNARequest>
 *   <Version>2.0</Version>
 *   <TransactionHeader>
 *     <SenderID>MD</SenderID>
 *     <ReceiverID>YOU</ReceiverID>
 *     <CountryCode>CA</CountryCode>
 *     <LoginID>**********</LoginID>
 *     <Password>**********</Password>
 *     <TransactionID>1</TransactionID>
 *   </TransactionHeader>
 *   <PNAInformation CustomerPartNumber="E0123456" Quantity="1"/>
 *   <PNAInformation ManufacturerPartNumber="RG0322" Quantity="1"/>
 *   <PNAInformation UPC="0760884154205" Quantity="1"/>
 *   <PNAInformation SKU="123512" Quantity="" />
 *   <ShowDetail>2</ShowDetail>
 * </PNARequest>
 */

type PriceAvailRequest struct {
    XMLName     xml.Name `xml:"PNARequest"`
}

/**
 * <PNAResponse>
 *   <Version>2.0</Version>
 *   <TransactionHeader>
 *     <SenderID>YOU</SenderID>
 *     <ReceiverID>MD</ReceiverID>
 *     <ErrorStatus ErrorNumber=""></ErrorStatus>
 *     <DocumentID>{454B8D7B-84C8-4B9C-9F72-4E179022A492}</DocumentID>
 *     <TransactionID>1</TransactionID>
 *     <TimeStamp>2014-06-11T16:55:04</TimeStamp>
 *   </TransactionHeader>
 *   <PriceAndAvailability SKU="T42794" Quantity="">
 *     <Price>18.65</Price>
 *     <SpecialPriceFlag></SpecialPriceFlag>
 *     <ManufacturerPartNumber>AVP-6</ManufacturerPartNumber>
 *     <ManufacturerPartNumberOccurs></ManufacturerPartNumberOccurs>
 *     <VendorNumber>Q410</VendorNumber>
 *     <Description>MIDLAND DESKTOP CHARGER</Description>
 *     <ReserveInventoryFlag>Y</ReserveInventoryFlag>
 *     <AvailableRebQty>0</AvailableRebQty>
 *     <Branch Name="Vancouver" ID="10">
 *       <Availability>0</Availability>
 *       <OnOrder>0</OnOrder>
 *       <ETADate></ETADate>
 *     </Branch>
 *     <Branch Name="Toronto" ID="40">
 *       <Availability>1</Availability>
 *       <OnOrder>0</OnOrder>
 *       <ETADate></ETADate>
 *     </Branch>
 *     <UPC>0046014298750</UPC>
 *     <CustomerPartNumber></CustomerPartNumber>
 *   </PriceAndAvailability>
 * </PNAResponse>
 */

type PriceAvailResponse struct {
    XMLName     xml.Name `xml:"PNAResponse"`
}
