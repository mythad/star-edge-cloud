package models

import "encoding/xml"

// Rule - 规则
type Rule struct {
	Name      string `xml:"name,attr"`
	Yql       string `xml:"Yql"`
	PostAddr  string `xml:"PostAddr"`
	ChildRule []Rule `xml:"Rule"`
}

// Rules 规则集合
type Rules struct {
	XMLName   xml.Name `xml:"Rules"`
	Version   string   `xml:"Version"`
	RuleItems []Rule   `xml:"Rule"`
	State     string   `xml:"State"`
	// Groups  []string `xml:"Group>Value"`
}
