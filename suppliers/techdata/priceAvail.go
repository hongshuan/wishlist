package techdata

/**
 * <XML_PriceAvailability_Submit>
 *   <Header>
 *     <UserName>******</UserName>  
 *     <Password>************</Password>
 *     <ResponseVersion>1.4</ResponseVersion>
 *   </Header>
 *   <Detail>
 *     <LineInfo>
 *       <RefIDQual>VP</RefIDQual>
 *       <RefID>296460</RefID>
 *     </LineInfo>
 *     <LineInfo>
 *       <RefIDQual>VP</RefIDQual>
 *       <RefID>180054</RefID>
 *     </LineInfo>
 *   </Detail>
 * </XML_PriceAvailability_Submit>
 */

type PriceAvailRequest struct {
    XMLName     xml.Name `xml:"XML_PriceAvailability_Submit"`
}

/**
 * <!DOCTYPE XML_PriceAvailability_Response SYSTEM "XML_PriceAvailability_Response.dtd">
 * <XML_PriceAvailability_Response>
 *   <Header>
 *     <TransSetIDCode>846REC</TransSetIDCode>
 *     <TransControlID />
 *     <ResponseVersion>1.4</ResponseVersion>
 *   </Header>
 *   <Detail>
 *     <LineInfo>
 *       <RefIDQual1>VP</RefIDQual1>
 *       <RefID1>296460</RefID1>
 *       <ErrorInfo>
 *         <RefIDQual3>1Q</RefIDQual3>
 *         <RefID3>0</RefID3>
 *         <ErrorDesc>ITEM NOT FOUND</ErrorDesc>
 *       </ErrorInfo>
 *     </LineInfo>
 *     <LineInfo>
 *       <RefIDQual1>VP</RefIDQual1>
 *       <RefID1>180054</RefID1>
 *       <ErrorInfo>
 *         <RefIDQual3>1Q</RefIDQual3>
 *         <RefID3>0</RefID3>
 *         <ErrorDesc>ITEM NOT FOUND</ErrorDesc>
 *       </ErrorInfo>
 *     </LineInfo>
 *   </Detail>
 *   <Summary>
 *     <NbrOfSegments>5</NbrOfSegments>
 *   </Summary>
 * </XML_PriceAvailability_Response>
 */

type PriceAvailResponse struct {
    XMLName     xml.Name `xml:"XML_PriceAvailability_Response"`
}
