// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package powerpoint

import "github.com/IGLOU-EU/unioffice"

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	unioffice.RegisterConstructor("urn:schemas-microsoft-com:office:powerpoint", "CT_Empty", NewCT_Empty)
	unioffice.RegisterConstructor("urn:schemas-microsoft-com:office:powerpoint", "CT_Rel", NewCT_Rel)
	unioffice.RegisterConstructor("urn:schemas-microsoft-com:office:powerpoint", "iscomment", NewIscomment)
	unioffice.RegisterConstructor("urn:schemas-microsoft-com:office:powerpoint", "textdata", NewTextdata)
}
