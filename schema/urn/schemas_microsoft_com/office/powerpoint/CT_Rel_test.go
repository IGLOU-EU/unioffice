// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package powerpoint_test

import (
	"encoding/xml"
	"testing"

	"github.com/IGLOU-EU/unioffice/schema/urn/schemas_microsoft_com/office/powerpoint"
)

func TestCT_RelConstructor(t *testing.T) {
	v := powerpoint.NewCT_Rel()
	if v == nil {
		t.Errorf("powerpoint.NewCT_Rel must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed powerpoint.CT_Rel should validate: %s", err)
	}
}

func TestCT_RelMarshalUnmarshal(t *testing.T) {
	v := powerpoint.NewCT_Rel()
	buf, _ := xml.Marshal(v)
	v2 := powerpoint.NewCT_Rel()
	xml.Unmarshal(buf, v2)
}
