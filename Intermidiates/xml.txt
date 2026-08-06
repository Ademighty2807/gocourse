package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age,omitempty"`
	City    string   `xml:"city,omitempty"`
	Address Address  `xml:"address,omitempty"`
	Email   string   `xml:"-"`
	// Email   string `xml:"email"`
}

type Address struct {
	City  string `xml:"city"`
	State string `xml:"state"`
}

func main() {
	person := Person{
		Name:  "John Doe",
		Age:   30,
		City:  "New York",
		Email: "john.doe@example.com",
		Address: Address{
			City:  "New York",
			State: "NY",
		},
	}

	// person := Person{
	// 	Name: "John Doe",
	// 	// Age:   30,
	// 	// City:  "New York",
	// 	Email: "john.doe@example.com",
	// 	Address: Address{
	// 		City:  "New York",
	// 		State: "NY",
	// 	},
	// }
	fmt.Println(person)

	xmlData, err := xml.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling XML:", err)
		return
	}
	fmt.Println(string(xmlData))

	xmlData1, err := xml.MarshalIndent(person, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling XML:", err)
		return
	}
	fmt.Println(string(xmlData1))

	xmlRaw := `<person><name>Jane Doe</name><age>25</age><address><city>Los Angeles</city><state>CA</state></address></person>`
	var personxml Person
	err = xml.Unmarshal([]byte(xmlRaw), &personxml)
	if err != nil {
		fmt.Println("Error unmarshalling XML:", err)
		return
	}
	fmt.Println(personxml)
	fmt.Println("Local String:", personxml.XMLName.Local)
	fmt.Println("Namespace:", personxml.XMLName.Space)

	xmlDataAttr, err := xml.MarshalIndent(Book{
		Title:      "The Go Programming Language",
		Author:     "Alan A. A. Donovan & Brian W. Kernighan",
		ISBN:       "978-0134190550",
		Pseudo:     "This is a pseudo element",
		PseudoAttr: "This is a pseudo attribute",
	}, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling XML:", err)
		return
	}
	fmt.Println(string(xmlDataAttr))
}

type Book struct {
	XMLName    xml.Name `xml:"book"`
	Title      string   `xml:"title,attr"`
	Author     string   `xml:"author,attr"`
	ISBN       string   `xml:"isbn,attr"`
	Pseudo     string   `xml:"pseudo"`
	PseudoAttr string   `xml:"pseudoattr,attr"`
}

// <book isbn="978-3-16-148410-0" color="blue">
