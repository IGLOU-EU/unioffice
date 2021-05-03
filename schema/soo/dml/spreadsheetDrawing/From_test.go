// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package spreadsheetDrawing_test

import (
	"encoding/xml"
	"testing"

	"github.com/IGLOU-EU/unioffice/schema/soo/dml/spreadsheetDrawing"
)

func TestFromConstructor(t *testing.T) {
	v := spreadsheetDrawing.NewFrom()
	if v == nil {
		t.Errorf("spreadsheetDrawing.NewFrom must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetDrawing.From should validate: %s", err)
	}
}

func TestFromMarshalUnmarshal(t *testing.T) {
	v := spreadsheetDrawing.NewFrom()
	buf, _ := xml.Marshal(v)
	v2 := spreadsheetDrawing.NewFrom()
	xml.Unmarshal(buf, v2)
}
