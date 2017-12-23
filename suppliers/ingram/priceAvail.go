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
    XMLName xml.Name `xml:"PNARequest"`
    Version string `xml:"Version"`

    Header struct {
        SenderID      string `xml:"SenderID"`
        ReceiverID    string `xml:"ReceiverID"`
        CountryCode   string `xml:"CountryCode"`
        LoginID       string `xml:"LoginID"`
        Password      string `xml:"Password"`
        TransactionID string `xml:"TransactionID"`
    } `xml:"TransactionHeader"`

    Item struct {
     // CPartNum string `xml:"CustomerPartNumber"`
     // MPartNum string `xml:"ManufacturerPartNumber"`
     // UPC      string `xml:"UPC"`
        Sku      string `xml:"SKU,attr"`
        Qty      string `xml:"Quantity,attr"`
    } `xml:"PNAInformation"`

    ShowDetail string `xml:"ShowDetail"`
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
    XMLName xml.Name `xml:"PNAResponse"`
    Version string `xml:"Version"`

    Header struct {
        SenderID   string `xml:"SenderID"`
        ReceiverID string `xml:"ReceiverID"`

        ErrorStatus struct {
            ErrorNumber string `xml:"ErrorNumber,attr"`
        } `xml:"ErrorStatus"`

        DocumentID    string `xml:"DocumentID"`
        TransactionID string `xml:"TransactionID"`
        TimeStamp     string `xml:"TimeStamp"`
    } `xml:"TransactionHeader"`

    Item struct {
        Sku                          string `xml:"SKU,attr"`
        Qty                          string `xml:"Quantity,attr"`
        Price                        string `xml:"Price"`
        SpecialPriceFlag             string `xml:"SpecialPriceFlag"`
        ManufacturerPartNumber       string `xml:"ManufacturerPartNumber"`
        ManufacturerPartNumberOccurs string `xml:"ManufacturerPartNumberOccurs"`
        VendorNumber                 string `xml:"VendorNumber"`
        Description                  string `xml:"Description"`
        ReserveInventoryFlag         string `xml:"ReserveInventoryFlag"`
        AvailableRebQty              string `xml:"AvailableRebQty"`

        Branches []struct {
            ID      string `xml:"ID,attr"`
            Name    string `xml:"Name,attr"`
            Qty     string `xml:"Availability"`
            OnOrder string `xml:"OnOrder"`
            ETADate string `xml:"ETADate"`
        } `xml:"Branch"`

        UPC string `xml:"UPC"`
        CustomerPartNumber string `xml:"CustomerPartNumber"`
    } `xml:"PriceAndAvailability"`
}
