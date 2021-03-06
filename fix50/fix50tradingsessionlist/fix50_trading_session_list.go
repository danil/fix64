// Code generated by protofix; DO NOT EDIT.
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fix50tradingsessionlist is a format of the FIX.5.0 TradingSessionList message.
package fix50tradingsessionlist

import (
	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/marshfix"
)

var (
	FIX50TradingSessionListMarshaler   = marshfix.Marshal{Tag: "FIX50", Format: FIX50TradingSessionList}
	FIX50TradingSessionListUnmarshaler = marshfix.Unmarshal{Tag: "FIX50", Format: FIX50TradingSessionList}
)

// FIX50TradingSessionList is a FIX.5.0 format of the TradingSessionList message which maps the codecs into individual fields.
var FIX50TradingSessionList = f0.Format{
	Fields: map[int]f0.Codec{
		TradSesReqID335:           f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TradingSessionID336:       f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TradingSessionSubID625:    f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SecurityExchange207:       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TradSesMethod338:          f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* ELECTRONIC */, 2 /* OPEN_OUTCRY */, 3 /* TWO_PARTY */), f0.Var{1, 19}},
		TradSesMode339:            f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* TESTING */, 2 /* SIMULATED */, 3 /* PRODUCTION */), f0.Var{1, 19}},
		UnsolicitedIndicator325:   f0.Fld{Opt, f0.ASCII, f0.BoolDefault(false /* NO */, true /* YES */), f0.Con{1}},
		TradSesStatus340:          f0.Fld{Req, f0.ASCII, f0.IntDefault(0 /* UNKNOWN */, 1 /* HALTED */, 2 /* OPEN */, 3 /* CLOSED */, 4 /* PRE_OPEN */, 5 /* PRE_CLOSE */, 6 /* REQUEST_REJECTED */), f0.Var{1, 19}},
		TradSesStatusRejReason567: f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* UNKNOWN_OR_INVALID_TRADINGSESSIONID */, 99 /* OTHER */), f0.Var{1, 19}},
		TradSesStartTime341:       f0.Fld{Opt, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		TradSesOpenTime342:        f0.Fld{Opt, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		TradSesPreCloseTime343:    f0.Fld{Opt, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		TradSesCloseTime344:       f0.Fld{Opt, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		TradSesEndTime345:         f0.Fld{Opt, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		TotalVolumeTraded387:      f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		Text58:                    f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncodedTextLen354:         f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncodedText355:            f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
	},
	Unknown0: f0.UnknownFld{f0.ASCII, f0.Unknown(), f0.Var{1, 65536}},
	Sort: []int{
		TradSesReqID335,           // STRING
		TradingSessionID336,       // STRING
		TradingSessionSubID625,    // STRING
		SecurityExchange207,       // EXCHANGE
		TradSesMethod338,          // INT
		TradSesMode339,            // INT
		UnsolicitedIndicator325,   // BOOLEAN
		TradSesStatus340,          // INT
		TradSesStatusRejReason567, // INT
		TradSesStartTime341,       // UTCTIMESTAMP
		TradSesOpenTime342,        // UTCTIMESTAMP
		TradSesPreCloseTime343,    // UTCTIMESTAMP
		TradSesCloseTime344,       // UTCTIMESTAMP
		TradSesEndTime345,         // UTCTIMESTAMP
		TotalVolumeTraded387,      // QTY
		Text58,                    // STRING
		EncodedTextLen354,         // LENGTH
		EncodedText355,            // DATA
	},
}

const Req, Opt = true, false

const (
	TradSesReqID335           = 335 // STRING
	TradingSessionID336       = 336 // STRING
	TradingSessionSubID625    = 625 // STRING
	SecurityExchange207       = 207 // EXCHANGE
	TradSesMethod338          = 338 // INT
	TradSesMode339            = 339 // INT
	UnsolicitedIndicator325   = 325 // BOOLEAN
	TradSesStatus340          = 340 // INT
	TradSesStatusRejReason567 = 567 // INT
	TradSesStartTime341       = 341 // UTCTIMESTAMP
	TradSesOpenTime342        = 342 // UTCTIMESTAMP
	TradSesPreCloseTime343    = 343 // UTCTIMESTAMP
	TradSesCloseTime344       = 344 // UTCTIMESTAMP
	TradSesEndTime345         = 345 // UTCTIMESTAMP
	TotalVolumeTraded387      = 387 // QTY
	Text58                    = 58  // STRING
	EncodedTextLen354         = 354 // LENGTH
	EncodedText355            = 355 // DATA
)
