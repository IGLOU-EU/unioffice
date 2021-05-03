// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/IGLOU-EU/unioffice"
)

type CT_Items struct {
	// Field Count
	CountAttr *uint32
	// PivotTable Field Item
	Item []*CT_Item
}

func NewCT_Items() *CT_Items {
	ret := &CT_Items{}
	return ret
}

func (m *CT_Items) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	seitem := xml.StartElement{Name: xml.Name{Local: "ma:item"}}
	for _, c := range m.Item {
		e.EncodeElement(c, seitem)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Items) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
			continue
		}
	}
lCT_Items:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "item"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "item"}:
				tmp := NewCT_Item()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Item = append(m.Item, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_Items %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Items
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Items and its children
func (m *CT_Items) Validate() error {
	return m.ValidateWithPath("CT_Items")
}

// ValidateWithPath validates the CT_Items and its children, prefixing error messages with path
func (m *CT_Items) ValidateWithPath(path string) error {
	for i, v := range m.Item {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Item[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
