// Code generated by protofix; DO NOT EDIT.
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fix50sp2userrequest is a format of the FIX.5.0 servicepack 2 UserRequest message.
package fix50sp2userrequest

import (
	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/marshfix"
)

var (
	FIX50SP2UserRequestMarshaler   = marshfix.Marshal{Tag: "FIX50SP2", Format: FIX50SP2UserRequest}
	FIX50SP2UserRequestUnmarshaler = marshfix.Unmarshal{Tag: "FIX50SP2", Format: FIX50SP2UserRequest}
)

// FIX50SP2UserRequest is a FIX.5.0 servicepack 2 format of the UserRequest message which maps the codecs into individual fields.
var FIX50SP2UserRequest = f0.Format{
	Fields: map[int]f0.Codec{
		UserRequestID923:            f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UserRequestType924:          f0.Fld{Req, f0.ASCII, f0.IntDefault(1 /* LOG_ON_USER */, 2 /* LOG_OFF_USER */, 3 /* CHANGE_PASSWORD_FOR_USER */, 4 /* REQUEST_INDIVIDUAL_USER_STATUS */), f0.Var{1, 19}},
		Username553:                 f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		Password554:                 f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		NewPassword925:              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RawDataLength95:             f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		RawData96:                   f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		EncryptedPasswordMethod1400: f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		EncryptedPasswordLen1401:    f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncryptedPassword1402:       f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		EncryptedNewPasswordLen1403: f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncryptedNewPassword1404:    f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
	},
	Unknown0: f0.UnknownFld{f0.ASCII, f0.Unknown(), f0.Var{1, 65536}},
	Sort: []int{
		UserRequestID923,            // STRING
		UserRequestType924,          // INT
		Username553,                 // STRING
		Password554,                 // STRING
		NewPassword925,              // STRING
		RawDataLength95,             // LENGTH
		RawData96,                   // DATA
		EncryptedPasswordMethod1400, // INT
		EncryptedPasswordLen1401,    // LENGTH
		EncryptedPassword1402,       // DATA
		EncryptedNewPasswordLen1403, // LENGTH
		EncryptedNewPassword1404,    // DATA
	},
}

const Req, Opt = true, false

const (
	UserRequestID923            = 923  // STRING
	UserRequestType924          = 924  // INT
	Username553                 = 553  // STRING
	Password554                 = 554  // STRING
	NewPassword925              = 925  // STRING
	RawDataLength95             = 95   // LENGTH
	RawData96                   = 96   // DATA
	EncryptedPasswordMethod1400 = 1400 // INT
	EncryptedPasswordLen1401    = 1401 // LENGTH
	EncryptedPassword1402       = 1402 // DATA
	EncryptedNewPasswordLen1403 = 1403 // LENGTH
	EncryptedNewPassword1404    = 1404 // DATA
)
