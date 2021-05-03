// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package wml

import (
	"encoding/xml"
	"fmt"

	"github.com/IGLOU-EU/unioffice/schema/soo/ofc/sharedTypes"
)

type CT_Column struct {
	// Column Width
	WAttr *sharedTypes.ST_TwipsMeasure
	// Space Before Following Column
	SpaceAttr *sharedTypes.ST_TwipsMeasure
}

func NewCT_Column() *CT_Column {
	ret := &CT_Column{}
	return ret
}

func (m *CT_Column) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.WAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:w"},
			Value: fmt.Sprintf("%v", *m.WAttr)})
	}
	if m.SpaceAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:space"},
			Value: fmt.Sprintf("%v", *m.SpaceAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Column) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "w" {
			parsed, err := ParseUnionST_TwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.WAttr = &parsed
			continue
		}
		if attr.Name.Local == "space" {
			parsed, err := ParseUnionST_TwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.SpaceAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Column: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Column and its children
func (m *CT_Column) Validate() error {
	return m.ValidateWithPath("CT_Column")
}

// ValidateWithPath validates the CT_Column and its children, prefixing error messages with path
func (m *CT_Column) ValidateWithPath(path string) error {
	if m.WAttr != nil {
		if err := m.WAttr.ValidateWithPath(path + "/WAttr"); err != nil {
			return err
		}
	}
	if m.SpaceAttr != nil {
		if err := m.SpaceAttr.ValidateWithPath(path + "/SpaceAttr"); err != nil {
			return err
		}
	}
	return nil
}
