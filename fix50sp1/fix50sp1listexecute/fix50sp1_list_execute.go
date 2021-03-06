// Code generated by protofix; DO NOT EDIT.
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fix50sp1listexecute is a format of the FIX.5.0 servicepack 1 ListExecute message.
package fix50sp1listexecute

import (
	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/marshfix"
)

var (
	FIX50SP1ListExecuteMarshaler   = marshfix.Marshal{Tag: "FIX50SP1", Format: FIX50SP1ListExecute}
	FIX50SP1ListExecuteUnmarshaler = marshfix.Unmarshal{Tag: "FIX50SP1", Format: FIX50SP1ListExecute}
)

// FIX50SP1ListExecute is a FIX.5.0 servicepack 1 format of the ListExecute message which maps the codecs into individual fields.
var FIX50SP1ListExecute = f0.Format{
	Fields: map[int]f0.Codec{
		ListID66:          f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		ClientBidID391:    f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		BidID390:          f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TransactTime60:    f0.Fld{Req, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		Text58:            f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncodedTextLen354: f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncodedText355:    f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
	},
	Unknown0: f0.UnknownFld{f0.ASCII, f0.Unknown(), f0.Var{1, 65536}},
	Sort: []int{
		ListID66,          // STRING
		ClientBidID391,    // STRING
		BidID390,          // STRING
		TransactTime60,    // UTCTIMESTAMP
		Text58,            // STRING
		EncodedTextLen354, // LENGTH
		EncodedText355,    // DATA
	},
}

const Req, Opt = true, false

const (
	ListID66          = 66  // STRING
	ClientBidID391    = 391 // STRING
	BidID390          = 390 // STRING
	TransactTime60    = 60  // UTCTIMESTAMP
	Text58            = 58  // STRING
	EncodedTextLen354 = 354 // LENGTH
	EncodedText355    = 355 // DATA
)
