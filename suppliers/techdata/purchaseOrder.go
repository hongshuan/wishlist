package techdata

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
}
