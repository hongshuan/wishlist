package asi

type PriceAvailRequest struct {
    XMLName     xml.Name `xml:""`
}

/**
 * <ASIInventory request="04/20/2014 18:45" excute="423">
 *   <Inventory SKU="102040">
 *     <ItemId>920-002478</ItemId>
 *     <Description><![CDATA[Logitech Keyboard 920-002478 Desktop K120 USB Black Retail]]></Description>
 *     <Vendor>Logitech</Vendor>
 *     <Category>KB</Category>
 *     <SubCategory>98</SubCategory>
 *     <UPC>097855065537</UPC>
 *     <Price>10.84</Price>
 *     <Rebate><![CDATA[none]]></Rebate>
 *     <Term>none</Term>
 *     <Weight>1.54</Weight>
 *     <Status>ACTIVE</Status>
 *     <Qty>
 *       <Branch Code="1016" Name="Fremont">678</Branch>
 *       <Branch Code="1028" Name="Los Angeles">915</Branch>
 *       <Branch Code="1116" Name="Atlanta">18</Branch>
 *       <Branch Code="1216" Name="Chicago">0</Branch>
 *       <Branch Code="1316" Name="Dallas">40</Branch>
 *       <Branch Code="1516" Name="Kansas">0</Branch>
 *       <Branch Code="1716" Name="New Jersey">27</Branch>
 *       <Branch Code="2416" Name="Miami">0</Branch>
 *       <Branch Code="2616" Name="Portland">0</Branch>
 *       <Branch Code="1016" Name="DC">0</Branch>
 *     </Qty>
 *   </Inventory>
 * </ASIInventory>
 */

/**
 * <ASIInventory>
 *   <error>
 *     <code>101</code>
 *     <message>No inventory found</message>
 *   </error>
 * </ASIInventory>
 */

type PriceAvailResponse struct {
    XMLName     xml.Name `xml:""`
}
