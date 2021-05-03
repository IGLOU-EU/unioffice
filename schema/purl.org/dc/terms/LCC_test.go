// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package terms_test

import (
	"encoding/xml"
	"testing"

	"github.com/IGLOU-EU/unioffice/schema/purl.org/dc/terms"
)

func TestLCCConstructor(t *testing.T) {
	v := terms.NewLCC()
	if v == nil {
		t.Errorf("terms.NewLCC must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed terms.LCC should validate: %s", err)
	}
}

func TestLCCMarshalUnmarshal(t *testing.T) {
	v := terms.NewLCC()
	buf, _ := xml.Marshal(v)
	v2 := terms.NewLCC()
	xml.Unmarshal(buf, v2)
}
