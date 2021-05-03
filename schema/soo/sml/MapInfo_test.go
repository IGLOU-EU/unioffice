// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/IGLOU-EU/unioffice/schema/soo/sml"
)

func TestMapInfoConstructor(t *testing.T) {
	v := sml.NewMapInfo()
	if v == nil {
		t.Errorf("sml.NewMapInfo must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.MapInfo should validate: %s", err)
	}
}

func TestMapInfoMarshalUnmarshal(t *testing.T) {
	v := sml.NewMapInfo()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewMapInfo()
	xml.Unmarshal(buf, v2)
}
