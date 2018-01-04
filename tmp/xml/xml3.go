package main

import (
	"encoding/xml"
	"fmt"
	//"os"
)

type Address struct {
	Street  string
	Suite   string `xml:",omitempty"`
	City    string
	State   string
	Country string
}

type Phone struct {
	Type   string `xml:"Type,attr"`
	Number string `xml:",chardata"`
}

type Customer struct {
	FirstName string   `xml:"Name>First"`
	LastName  string   `xml:"Name>Last"`
	Phone     []Phone  `xml:"Phones>Phone"`
	Address   Address  `xml:"Location"`
}

func main() {
	c := &Customer{FirstName: "Hanson", LastName:"Li"}
	c.Address = Address{ "2980 Don Mills", "", "Toronto", "Ontario", "Canada" }

	c.Phone = append(c.Phone, Phone{ Type: "Home", Number:"111-222-3333" })
	c.Phone = append(c.Phone, Phone{ Type: "Work", Number:"555-666-7777" })

	output, _ := xml.MarshalIndent(c, "", "  ")
	fmt.Println(string(output))
	//os.Stdout.Write(output)
}
