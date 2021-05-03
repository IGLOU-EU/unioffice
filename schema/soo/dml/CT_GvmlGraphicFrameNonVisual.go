// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package dml

import (
	"encoding/xml"

	"github.com/IGLOU-EU/unioffice"
)

type CT_GvmlGraphicFrameNonVisual struct {
	CNvPr             *CT_NonVisualDrawingProps
	CNvGraphicFramePr *CT_NonVisualGraphicFrameProperties
}

func NewCT_GvmlGraphicFrameNonVisual() *CT_GvmlGraphicFrameNonVisual {
	ret := &CT_GvmlGraphicFrameNonVisual{}
	ret.CNvPr = NewCT_NonVisualDrawingProps()
	ret.CNvGraphicFramePr = NewCT_NonVisualGraphicFrameProperties()
	return ret
}

func (m *CT_GvmlGraphicFrameNonVisual) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	secNvPr := xml.StartElement{Name: xml.Name{Local: "a:cNvPr"}}
	e.EncodeElement(m.CNvPr, secNvPr)
	secNvGraphicFramePr := xml.StartElement{Name: xml.Name{Local: "a:cNvGraphicFramePr"}}
	e.EncodeElement(m.CNvGraphicFramePr, secNvGraphicFramePr)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GvmlGraphicFrameNonVisual) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CNvPr = NewCT_NonVisualDrawingProps()
	m.CNvGraphicFramePr = NewCT_NonVisualGraphicFrameProperties()
lCT_GvmlGraphicFrameNonVisual:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "cNvPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "cNvPr"}:
				if err := d.DecodeElement(m.CNvPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "cNvGraphicFramePr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "cNvGraphicFramePr"}:
				if err := d.DecodeElement(m.CNvGraphicFramePr, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_GvmlGraphicFrameNonVisual %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GvmlGraphicFrameNonVisual
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GvmlGraphicFrameNonVisual and its children
func (m *CT_GvmlGraphicFrameNonVisual) Validate() error {
	return m.ValidateWithPath("CT_GvmlGraphicFrameNonVisual")
}

// ValidateWithPath validates the CT_GvmlGraphicFrameNonVisual and its children, prefixing error messages with path
func (m *CT_GvmlGraphicFrameNonVisual) ValidateWithPath(path string) error {
	if err := m.CNvPr.ValidateWithPath(path + "/CNvPr"); err != nil {
		return err
	}
	if err := m.CNvGraphicFramePr.ValidateWithPath(path + "/CNvGraphicFramePr"); err != nil {
		return err
	}
	return nil
}
