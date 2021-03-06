// Code generated by protofix; DO NOT EDIT.
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fix42ordercancelrequest is a format of the FIX.4.2 OrderCancelRequest message.
package fix42ordercancelrequest

import (
	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/marshfix"
)

var (
	FIX42OrderCancelRequestMarshaler   = marshfix.Marshal{Tag: "FIX42", Format: FIX42OrderCancelRequest}
	FIX42OrderCancelRequestUnmarshaler = marshfix.Unmarshal{Tag: "FIX42", Format: FIX42OrderCancelRequest}
)

// FIX42OrderCancelRequest is a FIX.4.2 format of the OrderCancelRequest message which maps the codecs into individual fields.
var FIX42OrderCancelRequest = f0.Format{
	Fields: map[int]f0.Codec{
		BeginString8:              f0.Fld{Req, f0.ASCII, f0.StringDefault("FIX.4.2"), f0.Con{7}},
		MsgType35:                 f0.Fld{Req, f0.ASCII, f0.String("0" /* HEARTBEAT */, "1" /* TEST_REQUEST */, "2" /* RESEND_REQUEST */, "3" /* REJECT */, "4" /* SEQUENCE_RESET */, "5" /* LOGOUT */, "6" /* INDICATION_OF_INTEREST */, "7" /* ADVERTISEMENT */, "8" /* EXECUTION_REPORT */, "9" /* ORDER_CANCEL_REJECT */, "a" /* QUOTE_STATUS_REQUEST */, "A" /* LOGON */, "B" /* NEWS */, "b" /* QUOTE_ACKNOWLEDGEMENT */, "C" /* EMAIL */, "c" /* SECURITY_DEFINITION_REQUEST */, "D" /* ORDER_SINGLE */, "d" /* SECURITY_DEFINITION */, "E" /* ORDER_LIST */, "e" /* SECURITY_STATUS_REQUEST */, "f" /* SECURITY_STATUS */, "F" /* ORDER_CANCEL_REQUEST */, "G" /* ORDER_CANCEL_REPLACE_REQUEST */, "g" /* TRADING_SESSION_STATUS_REQUEST */, "H" /* ORDER_STATUS_REQUEST */, "h" /* TRADING_SESSION_STATUS */, "i" /* MASS_QUOTE */, "j" /* BUSINESS_MESSAGE_REJECT */, "J" /* ALLOCATION */, "K" /* LIST_CANCEL_REQUEST */, "k" /* BID_REQUEST */, "l" /* BID_RESPONSE */, "L" /* LIST_EXECUTE */, "m" /* LIST_STRIKE_PRICE */, "M" /* LIST_STATUS_REQUEST */, "N" /* LIST_STATUS */, "P" /* ALLOCATION_ACK */, "Q" /* DONT_KNOW_TRADE */, "R" /* QUOTE_REQUEST */, "S" /* QUOTE */, "T" /* SETTLEMENT_INSTRUCTIONS */, "V" /* MARKET_DATA_REQUEST */, "W" /* MARKET_DATA_SNAPSHOT_FULL_REFRESH */, "X" /* MARKET_DATA_INCREMENTAL_REFRESH */, "Y" /* MARKET_DATA_REQUEST_REJECT */, "Z" /* QUOTE_CANCEL */), f0.Var{1, 2}},
		SenderCompID49:            f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TargetCompID56:            f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		OnBehalfOfCompID115:       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		DeliverToCompID128:        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SecureDataLen90:           f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		SecureData91:              f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		MsgSeqNum34:               f0.Fld{Req, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		SenderSubID50:             f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SenderLocationID142:       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TargetSubID57:             f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TargetLocationID143:       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		OnBehalfOfSubID116:        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		OnBehalfOfLocationID144:   f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		DeliverToSubID129:         f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		DeliverToLocationID145:    f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		PossDupFlag43:             f0.Fld{Opt, f0.ASCII, f0.BoolDefault(false /* NO */, true /* YES */), f0.Con{1}},
		PossResend97:              f0.Fld{Opt, f0.ASCII, f0.BoolDefault(false /* NO */, true /* YES */), f0.Con{1}},
		SendingTime52:             f0.Fld{Req, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		OrigSendingTime122:        f0.Fld{Opt, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		XmlDataLen212:             f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		XmlData213:                f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		MessageEncoding347:        f0.Fld{Opt, f0.ASCII, f0.String("EUC-JP" /* EUC_JP */, "ISO-2022-JP" /* ISO_2022_JP */, "SHIFT_JIS" /* SHIFT_JIS */, "UTF-8" /* UTF_8 */), f0.Var{1, 65536}},
		LastMsgSeqNumProcessed369: f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		OnBehalfOfSendingTime370:  f0.Fld{Opt, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		OrigClOrdID41:             f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		OrderID37:                 f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		ClOrdID11:                 f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		ListID66:                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		Account1:                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		ClientID109:               f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		ExecBroker76:              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		Symbol55:                  f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SymbolSfx65:               f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SecurityID48:              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		IDSource22:                f0.Fld{Opt, f0.ASCII, f0.String("1" /* CUSIP */, "2" /* SEDOL */, "3" /* QUIK */, "4" /* ISIN_NUMBER */, "5" /* RIC_CODE */, "6" /* ISO_CURRENCY_CODE */, "7" /* ISO_COUNTRY_CODE */, "8" /* EXCHANGE_SYMBOL */, "9" /* CONSOLIDATED_TAPE_ASSOCIATION */), f0.Var{1, 65536}},
		SecurityType167:           f0.Fld{Opt, f0.ASCII, f0.String("?" /* WILDCARD_ENTRY */, "BA" /* BANKERS_ACCEPTANCE */, "CB" /* CONVERTIBLE_BOND */, "CD" /* CERTIFICATE_OF_DEPOSIT */, "CMO" /* COLLATERALIZE_MORTGAGE_OBLIGATION */, "CORP" /* CORPORATE_BOND */, "CP" /* COMMERCIAL_PAPER */, "CPP" /* CORPORATE_PRIVATE_PLACEMENT */, "CS" /* COMMON_STOCK */, "FHA" /* FEDERAL_HOUSING_AUTHORITY */, "FHL" /* FEDERAL_HOME_LOAN */, "FN" /* FEDERAL_NATIONAL_MORTGAGE_ASSOCIATION */, "FOR" /* FOREIGN_EXCHANGE_CONTRACT */, "FUT" /* FUTURE */, "GN" /* GOVERNMENT_NATIONAL_MORTGAGE_ASSOCIATION */, "GOVT" /* TREASURIES_PLUS_AGENCY_DEBENTURE */, "IET" /* MORTGAGE_IOETTE */, "MF" /* MUTUAL_FUND */, "MIO" /* MORTGAGE_INTEREST_ONLY */, "MPO" /* MORTGAGE_PRINCIPAL_ONLY */, "MPP" /* MORTGAGE_PRIVATE_PLACEMENT */, "MPT" /* MISCELLANEOUS_PASS_THRU */, "MUNI" /* MUNICIPAL_BOND */, "NONE" /* NO_ISITC_SECURITY_TYPE */, "OPT" /* OPTION */, "PS" /* PREFERRED_STOCK */, "RP" /* REPURCHASE_AGREEMENT */, "RVRP" /* REVERSE_REPURCHASE_AGREEMENT */, "SL" /* STUDENT_LOAN_MARKETING_ASSOCIATION */, "TD" /* TIME_DEPOSIT */, "USTB" /* US_TREASURY_BILL */, "WAR" /* WARRANT */, "ZOO" /* CATS_TIGERS_LIONS */), f0.Var{1, 65536}},
		MaturityMonthYear200:      f0.Fld{Opt, f0.ASCII, f0.MonthYearTime(), f0.Var{6, 8}},
		MaturityDay205:            f0.Fld{Opt, f0.ASCII, f0.UTCDateOnlyTime(), f0.Con{8}},
		PutOrCall201:              f0.Fld{Opt, f0.ASCII, f0.IntDefault(0 /* PUT */, 1 /* CALL */), f0.Var{1, 19}},
		StrikePrice202:            f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		OptAttribute206:           f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		ContractMultiplier231:     f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		CouponRate223:             f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		SecurityExchange207:       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		Issuer106:                 f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncodedIssuerLen348:       f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncodedIssuer349:          f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		SecurityDesc107:           f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncodedSecurityDescLen350: f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncodedSecurityDesc351:    f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		Side54:                    f0.Fld{Req, f0.ASCII, f0.String("1" /* BUY */, "2" /* SELL */, "3" /* BUY_MINUS */, "4" /* SELL_PLUS */, "5" /* SELL_SHORT */, "6" /* SELL_SHORT_EXEMPT */, "7" /* UNDISCLOSED */, "8" /* CROSS */, "9" /* CROSS_SHORT */), f0.Con{1}},
		TransactTime60:            f0.Fld{Req, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		OrderQty38:                f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		CashOrderQty152:           f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		ComplianceID376:           f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SolicitedFlag377:          f0.Fld{Opt, f0.ASCII, f0.BoolDefault(false /* NO */, true /* YES */), f0.Con{1}},
		Text58:                    f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncodedTextLen354:         f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncodedText355:            f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		SignatureLength93:         f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		Signature89:               f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
	},
	BodyLength9: f0.BodyLengthFld{f0.ASCII, f0.BodyLength(), f0.Var{1, 18}},
	CheckSum10:  f0.CheckSumFld{f0.ASCII, f0.CheckSum(), f0.Con{3}},
	Unknown0:    f0.UnknownFld{f0.ASCII, f0.Unknown(), f0.Var{1, 65536}},
	Sort: []int{
		BeginString8,              // STRING
		BodyLength9,               // INT
		MsgType35,                 // STRING
		SenderCompID49,            // STRING
		TargetCompID56,            // STRING
		OnBehalfOfCompID115,       // STRING
		DeliverToCompID128,        // STRING
		SecureDataLen90,           // LENGTH
		SecureData91,              // DATA
		MsgSeqNum34,               // INT
		SenderSubID50,             // STRING
		SenderLocationID142,       // STRING
		TargetSubID57,             // STRING
		TargetLocationID143,       // STRING
		OnBehalfOfSubID116,        // STRING
		OnBehalfOfLocationID144,   // STRING
		DeliverToSubID129,         // STRING
		DeliverToLocationID145,    // STRING
		PossDupFlag43,             // BOOLEAN
		PossResend97,              // BOOLEAN
		SendingTime52,             // UTCTIMESTAMP
		OrigSendingTime122,        // UTCTIMESTAMP
		XmlDataLen212,             // LENGTH
		XmlData213,                // DATA
		MessageEncoding347,        // STRING
		LastMsgSeqNumProcessed369, // INT
		OnBehalfOfSendingTime370,  // UTCTIMESTAMP
		OrigClOrdID41,             // STRING
		OrderID37,                 // STRING
		ClOrdID11,                 // STRING
		ListID66,                  // STRING
		Account1,                  // STRING
		ClientID109,               // STRING
		ExecBroker76,              // STRING
		Symbol55,                  // STRING
		SymbolSfx65,               // STRING
		SecurityID48,              // STRING
		IDSource22,                // STRING
		SecurityType167,           // STRING
		MaturityMonthYear200,      // MONTHYEAR
		MaturityDay205,            // DAYOFMONTH
		PutOrCall201,              // INT
		StrikePrice202,            // PRICE
		OptAttribute206,           // CHAR
		ContractMultiplier231,     // FLOAT
		CouponRate223,             // FLOAT
		SecurityExchange207,       // EXCHANGE
		Issuer106,                 // STRING
		EncodedIssuerLen348,       // LENGTH
		EncodedIssuer349,          // DATA
		SecurityDesc107,           // STRING
		EncodedSecurityDescLen350, // LENGTH
		EncodedSecurityDesc351,    // DATA
		Side54,                    // CHAR
		TransactTime60,            // UTCTIMESTAMP
		OrderQty38,                // QTY
		CashOrderQty152,           // QTY
		ComplianceID376,           // STRING
		SolicitedFlag377,          // BOOLEAN
		Text58,                    // STRING
		EncodedTextLen354,         // LENGTH
		EncodedText355,            // DATA
		SignatureLength93,         // LENGTH
		Signature89,               // DATA
		CheckSum10,                // STRING
	},
}

const Req, Opt = true, false

const (
	BeginString8              = 8   // STRING
	BodyLength9               = 9   // INT
	MsgType35                 = 35  // STRING
	SenderCompID49            = 49  // STRING
	TargetCompID56            = 56  // STRING
	OnBehalfOfCompID115       = 115 // STRING
	DeliverToCompID128        = 128 // STRING
	SecureDataLen90           = 90  // LENGTH
	SecureData91              = 91  // DATA
	MsgSeqNum34               = 34  // INT
	SenderSubID50             = 50  // STRING
	SenderLocationID142       = 142 // STRING
	TargetSubID57             = 57  // STRING
	TargetLocationID143       = 143 // STRING
	OnBehalfOfSubID116        = 116 // STRING
	OnBehalfOfLocationID144   = 144 // STRING
	DeliverToSubID129         = 129 // STRING
	DeliverToLocationID145    = 145 // STRING
	PossDupFlag43             = 43  // BOOLEAN
	PossResend97              = 97  // BOOLEAN
	SendingTime52             = 52  // UTCTIMESTAMP
	OrigSendingTime122        = 122 // UTCTIMESTAMP
	XmlDataLen212             = 212 // LENGTH
	XmlData213                = 213 // DATA
	MessageEncoding347        = 347 // STRING
	LastMsgSeqNumProcessed369 = 369 // INT
	OnBehalfOfSendingTime370  = 370 // UTCTIMESTAMP
	OrigClOrdID41             = 41  // STRING
	OrderID37                 = 37  // STRING
	ClOrdID11                 = 11  // STRING
	ListID66                  = 66  // STRING
	Account1                  = 1   // STRING
	ClientID109               = 109 // STRING
	ExecBroker76              = 76  // STRING
	Symbol55                  = 55  // STRING
	SymbolSfx65               = 65  // STRING
	SecurityID48              = 48  // STRING
	IDSource22                = 22  // STRING
	SecurityType167           = 167 // STRING
	MaturityMonthYear200      = 200 // MONTHYEAR
	MaturityDay205            = 205 // DAYOFMONTH
	PutOrCall201              = 201 // INT
	StrikePrice202            = 202 // PRICE
	OptAttribute206           = 206 // CHAR
	ContractMultiplier231     = 231 // FLOAT
	CouponRate223             = 223 // FLOAT
	SecurityExchange207       = 207 // EXCHANGE
	Issuer106                 = 106 // STRING
	EncodedIssuerLen348       = 348 // LENGTH
	EncodedIssuer349          = 349 // DATA
	SecurityDesc107           = 107 // STRING
	EncodedSecurityDescLen350 = 350 // LENGTH
	EncodedSecurityDesc351    = 351 // DATA
	Side54                    = 54  // CHAR
	TransactTime60            = 60  // UTCTIMESTAMP
	OrderQty38                = 38  // QTY
	CashOrderQty152           = 152 // QTY
	ComplianceID376           = 376 // STRING
	SolicitedFlag377          = 377 // BOOLEAN
	Text58                    = 58  // STRING
	EncodedTextLen354         = 354 // LENGTH
	EncodedText355            = 355 // DATA
	SignatureLength93         = 93  // LENGTH
	Signature89               = 89  // DATA
	CheckSum10                = 10  // STRING
)
