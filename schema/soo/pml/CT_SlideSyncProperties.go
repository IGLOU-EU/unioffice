// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package pml

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/IGLOU-EU/unioffice"
)

type CT_SlideSyncProperties struct {
	// Server's Slide File ID
	ServerSldIdAttr string
	// Server's Slide File's modification date/time
	ServerSldModifiedTimeAttr time.Time
	// Client Slide Insertion date/time
	ClientInsertedTimeAttr time.Time
	ExtLst                 *CT_ExtensionList
}

func NewCT_SlideSyncProperties() *CT_SlideSyncProperties {
	ret := &CT_SlideSyncProperties{}
	return ret
}

func (m *CT_SlideSyncProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "serverSldId"},
		Value: fmt.Sprintf("%v", m.ServerSldIdAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "serverSldModifiedTime"},
		Value: fmt.Sprintf("%v", m.ServerSldModifiedTimeAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "clientInsertedTime"},
		Value: fmt.Sprintf("%v", m.ClientInsertedTimeAttr)})
	e.EncodeToken(start)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SlideSyncProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "serverSldId" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ServerSldIdAttr = parsed
			continue
		}
		if attr.Name.Local == "serverSldModifiedTime" {
			parsed, err := ParseStdlibTime(attr.Value)
			if err != nil {
				return err
			}
			m.ServerSldModifiedTimeAttr = parsed
			continue
		}
		if attr.Name.Local == "clientInsertedTime" {
			parsed, err := ParseStdlibTime(attr.Value)
			if err != nil {
				return err
			}
			m.ClientInsertedTimeAttr = parsed
			continue
		}
	}
lCT_SlideSyncProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_SlideSyncProperties %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SlideSyncProperties
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SlideSyncProperties and its children
func (m *CT_SlideSyncProperties) Validate() error {
	return m.ValidateWithPath("CT_SlideSyncProperties")
}

// ValidateWithPath validates the CT_SlideSyncProperties and its children, prefixing error messages with path
func (m *CT_SlideSyncProperties) ValidateWithPath(path string) error {
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
