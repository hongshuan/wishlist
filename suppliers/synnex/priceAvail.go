package synnex

type PriceAvailRequest struct {
    XMLName     xml.Name `xml:""`
}

/**
 * <?xml version="1.0" encoding="UTF-8"?>
 * <priceResponse>
 *   <customerNo>1150897</customerNo>
 *   <userName>roy@btecanada.com</userName>
 *   <PriceAvailabilityList>
 *     <synnexSKU>1031808</synnexSKU>
 *     <mfgPN>52107201</mfgPN>
 *     <mfgCode>676</mfgCode>
 *     <status>Active</status>
 *     <description>TONER-CART BLK OL400E/600E/810E/6E MP6</description>
 *     <GlobalProductStatusCode>Active</GlobalProductStatusCode>
 *     <price>27.12</price>
 *     <totalQuantity>2</totalQuantity>
 *     <AvailabilityByWarehouse>
 *       <warehouseInfo>
 *         <number>6</number>
 *         <zipcode>60446</zipcode>
 *         <city>Romeoville, IL</city>
 *         <addr>1180 Remington Blvd</addr>
 *       </warehouseInfo>
 *       <qty>0</qty>
 *     </AvailabilityByWarehouse>
 *     <AvailabilityByWarehouse>
 *       <warehouseInfo>
 *         <number>7</number>
 *         <zipcode>38654</zipcode>
 *         <city>Olive Branch, MS</city>
 *         <addr>10381 Stateline Rd</addr>
 *       </warehouseInfo>
 *       <qty>0</qty>
 *     </AvailabilityByWarehouse>
 *     <AvailabilityByWarehouse>
 *       <warehouseInfo>
 *         <number>26</number>
 *         <zipcode>B3S 1B3</zipcode>
 *         <city>Halifax, NS</city>
 *         <addr>115 Chain Lake Drive</addr>
 *       </warehouseInfo>
 *       <qty>0</qty>
 *     </AvailabilityByWarehouse>
 *     <AvailabilityByWarehouse>
 *       <warehouseInfo>
 *         <number>29</number>
 *         <zipcode>N1H 1B4</zipcode>
 *         <city>Guelph, ON</city>
 *         <addr>107 Woodlawn Road West</addr>
 *       </warehouseInfo>
 *       <qty>1</qty>
 *     </AvailabilityByWarehouse>
 *     <AvailabilityByWarehouse>
 *       <warehouseInfo>
 *         <number>31</number>
 *         <zipcode>T2C 4G6</zipcode>
 *         <city>Calgary, AB</city>
 *         <addr>5280 72 AVE SE</addr>
 *       </warehouseInfo>
 *       <qty>1</qty>
 *     </AvailabilityByWarehouse>
 *     <AvailabilityByWarehouse>
 *       <warehouseInfo>
 *         <number>60</number>
 *         <zipcode>V6V 1P6</zipcode>
 *         <city>Richmond, BC</city>
 *         <addr>3389 No. 6 Road</addr>
 *       </warehouseInfo>
 *       <qty>0</qty>
 *     </AvailabilityByWarehouse>
 *     <AvailabilityByWarehouse>
 *       <warehouseInfo>
 *         <number>98</number>
 *         <zipcode>00000</zipcode>
 *         <city>MFG Drop Shipped</city>
 *         <addr>MFG Drop Shipped</addr>
 *       </warehouseInfo>
 *       <qty>0</qty>
 *     </AvailabilityByWarehouse>
 *     <lineNumber>1</lineNumber>
 *   </PriceAvailabilityList>
 * </priceResponse>
 */

type PriceAvailResponse struct {
    XMLName     xml.Name `xml:"priceResponse"`
}
