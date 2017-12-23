package techdata

import (
    "encoding/xml"
)

/**
 * <XML_Order_Submit>
 *   <Header>
 *     <UserName>******</UserName>
 *     <Password>************</Password>
 *     <ResponseVersion>1.6</ResponseVersion>
 *     <OrderTypeCode>BS</OrderTypeCode>
 *     <PONbr>702-6308334-2375468</PONbr>
 *     <RefNbrs>
 *       <RefIDQual>EU</RefIDQual>
 *       <RefID>702-6308334-2375468</RefID>
 *     </RefNbrs>
 *     <SalesRequirementCode/>
 *     <Name>Julia L Offrey</Name>
 *     <AddrInfo>
 *       <Addr>1988 Robertson Road RR 5</Addr>
 *       <Addr/>
 *       <Addr>902.969.9903</Addr>
 *     </AddrInfo>
 *     <CityName>Montague</CityName>
 *     <StateProvinceCode>PE</StateProvinceCode>
 *     <PostalCode>C0A 1R0</PostalCode>
 *     <ContactName>Roy Zhang</ContactName>
 *     <ContactPhoneNbr>902.969.9903</ContactPhoneNbr>
 *     <EndUserInfo>
 *       <EuiGovAgency/>
 *       <EuiGovCabinetLevel/>
 *       <EuiContractNbr/>
 *       <EuiContractType/>
 *       <EuiOrderPriority/>
 *       <EuiMarketType/>
 *       <EuiContactName>Julia L Offrey</EuiContactName>
 *       <EuiPhoneNbr>902.969.9903</EuiPhoneNbr>
 *       <EuiFaxNbr/>
 *       <EuiName>Julia L Offrey</EuiName>
 *       <EuiAddr1>1988 Robertson Road RR 5</EuiAddr1>
 *       <EuiAddr2/>
 *       <EuiAddr3/>
 *       <EuiCityName>Montague</EuiCityName>
 *       <EuiStateProvinceCode>PE</EuiStateProvinceCode>
 *       <EuiPostalCode>C0A 1R0</EuiPostalCode>
 *       <EuiCountryCode>CA</EuiCountryCode>
 *       <EuiSicCode/>
 *       <EuiOrderPromoType/>
 *       <EuiEndUserLicenseNbr/>
 *       <EuiEndUserPODate/>
 *       <EuiEndUserRef1/>
 *       <EuiEndUserRef2/>
 *       <EuiEndUserRef3/>
 *       <EuiInstallName/>
 *       <EuiDropShipType/>
 *       <EuiContactEmailAddr1/>
 *       <EuiContactEmailAddr2/>
 *     </EndUserInfo>
 *     <MyOrderTracker>
 *       <ResellerEmail>doris@btecanada.com</ResellerEmail>
 *       <ResellerEvents>OC</ResellerEvents>
 *       <ResellerEvents>OS</ResellerEvents>
 *     </MyOrderTracker>
 *   </Header>
 *   <Detail>
 *     <LineInfo>
 *       <QtyOrdered>1</QtyOrdered>
 *       <ProductIDQual>VP</ProductIDQual>
 *       <ProductID>2171ZJ</ProductID>
 *       <WhseCode>A1</WhseCode>
 *       <IDCode>UP</IDCode>
 *       <OrderMessageLine/>
 *     </LineInfo>
 *   </Detail>
 * </XML_Order_Submit>
 */

type PurchaseOrderRequest struct {
    XMLName     xml.Name `xml:"XML_Order_Submit"`
    Header struct {
        UserName        string `xml:"UserName"`
        Password        string `xml:"Password"`
        ResponseVersion string `xml:"ResponseVersion"`
        OrderTypeCode   string `xml:"OrderTypeCode"`
        PONbr           string `xml:"PONbr"`

        RefNbrs struct {
            RefIDQual string `xml:"RefIDQual"`
            RefID     string `xml:"RefID"`
        } `xml:"RefNbrs"`

        SalesRequirementCode string `xml:"SalesRequirementCode"`
        Name                 string `xml:"Name"`
        AddrInfo           []string `xml:"AddrInfo>Addr"`
        CityName             string `xml:"CityName"`
        StateProvinceCode    string `xml:"StateProvinceCode"`
        PostalCode           string `xml:"PostalCode"`
        ContactName          string `xml:"ContactName"`
        ContactPhoneNbr      string `xml:"ContactPhoneNbr"`

        EndUserInfo struct {
            GovAgency         string `xml:"EuiGovAgency"`
            GovCabinetLevel   string `xml:"EuiGovCabinetLevel"`
            ContractNbr       string `xml:"EuiContractNbr"`
            ContractType      string `xml:"EuiContractType"`
            OrderPriority     string `xml:"EuiOrderPriority"`
            MarketType        string `xml:"EuiMarketType"`
            ContactName       string `xml:"EuiContactName"`
            PhoneNbr          string `xml:"EuiPhoneNbr"`
            FaxNbr            string `xml:"EuiFaxNbr"`
            Name              string `xml:"EuiName"`
            Addr1             string `xml:"EuiAddr1"`
            Addr2             string `xml:"EuiAddr2"`
            Addr3             string `xml:"EuiAddr3"`
            CityName          string `xml:"EuiCityName"`
            StateProvinceCode string `xml:"EuiStateProvinceCode"`
            PostalCode        string `xml:"EuiPostalCode"`
            CountryCode       string `xml:"EuiCountryCode"`
            SicCode           string `xml:"EuiSicCode"`
            OrderPromoType    string `xml:"EuiOrderPromoType"`
            EndUserLicenseNbr string `xml:"EuiEndUserLicenseNbr"`
            EndUserPODate     string `xml:"EuiEndUserPODate"`
            EndUserRef1       string `xml:"EuiEndUserRef1"`
            EndUserRef2       string `xml:"EuiEndUserRef2"`
            EndUserRef3       string `xml:"EuiEndUserRef3"`
            InstallName       string `xml:"EuiInstallName"`
            DropShipType      string `xml:"EuiDropShipType"`
            ContactEmailAddr1 string `xml:"EuiContactEmailAddr1"`
            ContactEmailAddr2 string `xml:"EuiContactEmailAddr2"`
        } `xml:"EndUserInfo"`

        MyOrderTracker struct {
            ResellerEmail  string `xml:"ResellerEmail"`
            ResellerEvents string `xml:"ResellerEvents"`
            ResellerEvents string `xml:"ResellerEvents"`
        } `xml:"MyOrderTracker"`

    } `xml:"Header"`

    Detail struct {
        LineInfo struct {
            QtyOrdered       string `xml:"QtyOrdered"`
            ProductIDQual    string `xml:"ProductIDQual"`
            ProductID        string `xml:"ProductID"`
            WhseCode         string `xml:"WhseCode"`
            IDCode           string `xml:"IDCode"`
            OrderMessageLine string `xml:"OrderMessageLine"`
        } `xml:"LineInfo"`
    } `xml:"Detail"`
}

/**
 * <!DOCTYPE XML_Order_Response SYSTEM "XML_Order_Response.dtd">
 * <XML_Order_Response>
 *   <Header>
 *     <UserName>567861</UserName>
 *     <TransSetIDCode>855</TransSetIDCode>
 *     <TransControlID/>
 *     <ResponseVersion>1.6</ResponseVersion>
 *     <PurposeCode/>
 *     <PONbr>702-6308334-2375468</PONbr>
 *     <RefID>3266709</RefID>
 *   </Header>
 *   <Summary/>
 * </XML_Order_Response>
 */

type PurchaseOrderResponse struct {
    XMLName     xml.Name `xml:"XML_Order_Response"`

    Header struct {
        Username        string `xml:"UserName"`
        TransSetIDCode  string `xml:"TransSetIDCode "`
        TransControlID  string `xml:"TransControlID "`
        ResponseVersion string `xml:"ResponseVersion "`
        PurposeCode     string `xml:"PurposeCode "`
        PoNum           string `xml:"PONbr"`
        RefID           string `xml:"RefID"`
    } `xml:"Header"`

    Summary string `xml:"Summary"`
}
