// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package vml_test

import (
	"encoding/xml"
	"testing"

	"github.com/IGLOU-EU/unioffice/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcTopConstructor(t *testing.T) {
	v := vml.NewOfcTop()
	if v == nil {
		t.Errorf("vml.NewOfcTop must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcTop should validate: %s", err)
	}
}

func TestOfcTopMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcTop()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcTop()
	xml.Unmarshal(buf, v2)
}
