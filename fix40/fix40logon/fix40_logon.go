// Code generated by protofix; DO NOT EDIT.
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fix40logon is a format of the FIX.4.0 Logon message.
package fix40logon

import (
	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/marshfix"
)

var (
	FIX40LogonMarshaler   = marshfix.Marshal{Tag: "FIX40", Format: FIX40Logon}
	FIX40LogonUnmarshaler = marshfix.Unmarshal{Tag: "FIX40", Format: FIX40Logon}
)

// FIX40Logon is a FIX.4.0 format of the Logon message which maps the codecs into individual fields.
var FIX40Logon = f0.Format{
	Fields: map[int]f0.Codec{
		BeginString8:        f0.Fld{Req, f0.ASCII, f0.StringDefault("FIX.4.0"), f0.Con{7}},
		MsgType35:           f0.Fld{Req, f0.ASCII, f0.String("0" /* HEARTBEAT */, "1" /* TEST_REQUEST */, "2" /* RESEND_REQUEST */, "3" /* REJECT */, "4" /* SEQUENCE_RESET */, "5" /* LOGOUT */, "6" /* INDICATION_OF_INTEREST */, "7" /* ADVERTISEMENT */, "8" /* EXECUTION_REPORT */, "9" /* ORDER_CANCEL_REJECT */, "A" /* LOGON */, "B" /* NEWS */, "C" /* EMAIL */, "D" /* ORDER_D */, "E" /* ORDER_E */, "F" /* ORDER_CANCEL_REQUEST */, "G" /* ORDER_CANCEL_REPLACE_REQUEST */, "H" /* ORDER_STATUS_REQUEST */, "J" /* ALLOCATION */, "K" /* LIST_CANCEL_REQUEST */, "L" /* LIST_EXECUTE */, "M" /* LIST_STATUS_REQUEST */, "N" /* LIST_STATUS */, "P" /* ALLOCATION_ACK */, "Q" /* DONT_KNOW_TRADE */, "R" /* QUOTE_REQUEST */, "S" /* QUOTE */), f0.Var{1, 2}},
		SenderCompID49:      f0.Fld{Req, f0.ASCII, f0.String(), f0.Con{1}},
		TargetCompID56:      f0.Fld{Req, f0.ASCII, f0.String(), f0.Con{1}},
		OnBehalfOfCompID115: f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		DeliverToCompID128:  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		SecureDataLen90:     f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		SecureData91:        f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		MsgSeqNum34:         f0.Fld{Req, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		SenderSubID50:       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		TargetSubID57:       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		OnBehalfOfSubID116:  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		DeliverToSubID129:   f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		PossDupFlag43:       f0.Fld{Opt, f0.ASCII, f0.String("N" /* NO */, "Y" /* YES */), f0.Con{1}},
		PossResend97:        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		SendingTime52:       f0.Fld{Req, f0.ASCII, f0.TZTime(), f0.Var{8, 12}},
		OrigSendingTime122:  f0.Fld{Opt, f0.ASCII, f0.TZTime(), f0.Var{8, 12}},
		EncryptMethod98:     f0.Fld{Req, f0.ASCII, f0.IntDefault(0 /* NONE */, 1 /* PKCS */, 2 /* DES */, 3 /* PKCS_DES */, 4 /* PGP_DES */, 5 /* PGP_DES_MD5 */, 6 /* PEM_DES_MD5 */), f0.Con{1}},
		HeartBtInt108:       f0.Fld{Req, f0.ASCII, f0.SecondsDuration(), f0.Var{1, 18}},
		RawDataLength95:     f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		RawData96:           f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		SignatureLength93:   f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		Signature89:         f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
	},
	BodyLength9: f0.BodyLengthFld{f0.ASCII, f0.BodyLength(), f0.Var{1, 18}},
	CheckSum10:  f0.CheckSumFld{f0.ASCII, f0.CheckSum(), f0.Con{3}},
	Unknown0:    f0.UnknownFld{f0.ASCII, f0.Unknown(), f0.Var{1, 65536}},
	Sort: []int{
		BeginString8,        // CHAR
		BodyLength9,         // INT
		MsgType35,           // CHAR
		SenderCompID49,      // CHAR
		TargetCompID56,      // CHAR
		OnBehalfOfCompID115, // CHAR
		DeliverToCompID128,  // CHAR
		SecureDataLen90,     // LENGTH
		SecureData91,        // DATA
		MsgSeqNum34,         // INT
		SenderSubID50,       // CHAR
		TargetSubID57,       // CHAR
		OnBehalfOfSubID116,  // CHAR
		DeliverToSubID129,   // CHAR
		PossDupFlag43,       // CHAR
		PossResend97,        // CHAR
		SendingTime52,       // TIME
		OrigSendingTime122,  // TIME
		EncryptMethod98,     // INT
		HeartBtInt108,       // INT
		RawDataLength95,     // LENGTH
		RawData96,           // DATA
		SignatureLength93,   // LENGTH
		Signature89,         // DATA
		CheckSum10,          // CHAR
	},
}

const Req, Opt = true, false

const (
	BeginString8        = 8   // CHAR
	BodyLength9         = 9   // INT
	MsgType35           = 35  // CHAR
	SenderCompID49      = 49  // CHAR
	TargetCompID56      = 56  // CHAR
	OnBehalfOfCompID115 = 115 // CHAR
	DeliverToCompID128  = 128 // CHAR
	SecureDataLen90     = 90  // LENGTH
	SecureData91        = 91  // DATA
	MsgSeqNum34         = 34  // INT
	SenderSubID50       = 50  // CHAR
	TargetSubID57       = 57  // CHAR
	OnBehalfOfSubID116  = 116 // CHAR
	DeliverToSubID129   = 129 // CHAR
	PossDupFlag43       = 43  // CHAR
	PossResend97        = 97  // CHAR
	SendingTime52       = 52  // TIME
	OrigSendingTime122  = 122 // TIME
	EncryptMethod98     = 98  // INT
	HeartBtInt108       = 108 // INT
	RawDataLength95     = 95  // LENGTH
	RawData96           = 96  // DATA
	SignatureLength93   = 93  // LENGTH
	Signature89         = 89  // DATA
	CheckSum10          = 10  // CHAR
)
