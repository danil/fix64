// Code generated by protofix; DO NOT EDIT.
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fix43ioi is a format of the FIX.4.3 IOI message.
package fix43ioi

import (
	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/marshfix"
)

var (
	FIX43IOIMarshaler   = marshfix.Marshal{Tag: "FIX43", Format: FIX43IOI}
	FIX43IOIUnmarshaler = marshfix.Unmarshal{Tag: "FIX43", Format: FIX43IOI}
)

// FIX43IOI is a FIX.4.3 format of the IOI message which maps the codecs into individual fields.
var FIX43IOI = f0.Format{
	Fields: map[int]f0.Codec{
		BeginString8:                  f0.Fld{Req, f0.ASCII, f0.StringDefault("FIX.4.3"), f0.Con{7}},
		MsgType35:                     f0.Fld{Req, f0.ASCII, f0.String("m" /* LIST_STRIKE_PRICE */, "l" /* BID_RESPONSE */, "k" /* BID_REQUEST */, "j" /* BUSINESS_MESSAGE_REJECT */, "i" /* MASS_QUOTE */, "h" /* TRADING_SESSION_STATUS */, "q" /* ORDER_MASS_CANCEL_REQUEST */, "AE" /* TRADE_CAPTURE_REPORT */, "o" /* REGISTRATION_INSTRUCTIONS */, "z" /* DERIVATIVE_SECURITY_LIST_REQUEST */, "AI" /* QUOTE_STATUS_REPORT */, "AH" /* RFQ_REQUEST */, "AG" /* QUOTE_REQUEST_REJECT */, "AF" /* ORDER_MASS_STATUS_REQUEST */, "n" /* XML_MESSAGE */, "AD" /* TRADE_CAPTURE_REPORT_REQUEST */, "g" /* TRADING_SESSION_STATUS_REQUEST */, "AC" /* MULTILEG_ORDER_CANCEL_REPLACE */, "AA" /* DERIVATIVE_SECURITY_LIST */, "r" /* ORDER_MASS_CANCEL_REPORT */, "y" /* SECURITY_LIST */, "x" /* SECURITY_LIST_REQUEST */, "w" /* SECURITY_TYPES */, "v" /* SECURITY_TYPE_REQUEST */, "u" /* CROSS_ORDER_CANCEL_REQUEST */, "t" /* CROSS_ORDER_CANCEL_REPLACE_REQUEST */, "s" /* NEW_ORDER_s */, "AB" /* NEW_ORDER_AB */, "9" /* ORDER_CANCEL_REJECT */, "7" /* ADVERTISEMENT */, "M" /* LIST_STATUS_REQUEST */, "E" /* ORDER_LIST */, "D" /* ORDER_SINGLE */, "C" /* EMAIL */, "B" /* NEWS */, "A" /* LOGON */, "N" /* LIST_STATUS */, "f" /* SECURITY_STATUS */, "P" /* ALLOCATION_ACK */, "8" /* EXECUTION_REPORT */, "0" /* HEARTBEAT */, "1" /* TEST_REQUEST */, "2" /* RESEND_REQUEST */, "3" /* REJECT */, "4" /* SEQUENCE_RESET */, "5" /* LOGOUT */, "6" /* INDICATION_OF_INTEREST */, "G" /* ORDER_CANCEL_REPLACE_REQUEST */, "K" /* LIST_CANCEL_REQUEST */, "e" /* SECURITY_STATUS_REQUEST */, "d" /* SECURITY_DEFINITION */, "c" /* SECURITY_DEFINITION_REQUEST */, "b" /* MASS_QUOTE_ACKNOWLEDGEMENT */, "a" /* QUOTE_STATUS_REQUEST */, "Z" /* QUOTE_CANCEL */, "Y" /* MARKET_DATA_REQUEST_REJECT */, "F" /* ORDER_CANCEL_REQUEST */, "J" /* ALLOCATION */, "p" /* REGISTRATION_INSTRUCTIONS_RESPONSE */, "L" /* LIST_EXECUTE */, "X" /* MARKET_DATA_INCREMENTAL_REFRESH */, "W" /* MARKET_DATA_SNAPSHOT_FULL_REFRESH */, "V" /* MARKET_DATA_REQUEST */, "T" /* SETTLEMENT_INSTRUCTIONS */, "S" /* QUOTE */, "R" /* QUOTE_REQUEST */, "Q" /* DONT_KNOW_TRADE */, "H" /* ORDER_STATUS_REQUEST */), f0.Var{1, 2}},
		SenderCompID49:                f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TargetCompID56:                f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		OnBehalfOfCompID115:           f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		DeliverToCompID128:            f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SecureDataLen90:               f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		SecureData91:                  f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		MsgSeqNum34:                   f0.Fld{Req, f0.ASCII, f0.SeqNum(), f0.Var{1, 19}},
		SenderSubID50:                 f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SenderLocationID142:           f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TargetSubID57:                 f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TargetLocationID143:           f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		OnBehalfOfSubID116:            f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		OnBehalfOfLocationID144:       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		DeliverToSubID129:             f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		DeliverToLocationID145:        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		PossDupFlag43:                 f0.Fld{Opt, f0.ASCII, f0.BoolDefault(false /* NO */, true /* YES */), f0.Con{1}},
		PossResend97:                  f0.Fld{Opt, f0.ASCII, f0.BoolDefault(false /* NO */, true /* YES */), f0.Con{1}},
		SendingTime52:                 f0.Fld{Req, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		OrigSendingTime122:            f0.Fld{Opt, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		XmlDataLen212:                 f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		XmlData213:                    f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		MessageEncoding347:            f0.Fld{Opt, f0.ASCII, f0.String("UTF-8" /* UTF_8 */, "ISO-2022-JP" /* ISO_2022_JP */, "EUC-JP" /* EUC_JP */, "SHIFT_JIS" /* SHIFT_JIS */), f0.Var{1, 65536}},
		LastMsgSeqNumProcessed369:     f0.Fld{Opt, f0.ASCII, f0.SeqNum(), f0.Var{1, 19}},
		OnBehalfOfSendingTime370:      f0.Fld{Opt, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		HopCompID628:                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		HopSendingTime629:             f0.Fld{Opt, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		HopRefID630:                   f0.Fld{Opt, f0.ASCII, f0.SeqNum(), f0.Var{1, 19}},
		IOIid23:                       f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		IOITransType28:                f0.Fld{Req, f0.ASCII, f0.String("C" /* CANCEL */, "N" /* NEW */, "R" /* REPLACE */), f0.Con{1}},
		IOIRefID26:                    f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		Side54:                        f0.Fld{Req, f0.ASCII, f0.String("6" /* SELL_SHORT_EXEMPT */, "B" /* AS_DEFINED */, "C" /* OPPOSITE */, "8" /* CROSS */, "9" /* CROSS_SHORT */, "1" /* BUY */, "2" /* SELL */, "3" /* BUY_MINUS */, "4" /* SELL_PLUS */, "A" /* CROSS_SHORT_EXEMPT */, "5" /* SELL_SHORT */, "7" /* UNDISCLOSED */), f0.Con{1}},
		QuantityType465:               f0.Fld{Opt, f0.ASCII, f0.IntDefault(6 /* CONTRACTS */, 7 /* OTHER */, 5 /* CURRENCY */, 4 /* ORIGINALFACE */, 3 /* CURRENTFACE */, 2 /* BONDS */, 1 /* SHARES */, 8 /* PAR */), f0.Var{1, 19}},
		IOIQty27:                      f0.Fld{Req, f0.ASCII, f0.String("L" /* LARGE */, "M" /* MEDIUM */, "S" /* SMALL */), f0.Var{1, 65536}},
		PriceType423:                  f0.Fld{Opt, f0.ASCII, f0.IntDefault(3 /* FIXED_AMOUNT */, 1 /* PERCENTAGE */, 4 /* DISCOUNT */, 6 /* BASIS_POINTS_RELATIVE_TO_BENCHMARK */, 7 /* TED_PRICE */, 8 /* TED_YIELD */, 5 /* PREMIUM */, 2 /* PER_SHARE */), f0.Var{1, 19}},
		Price44:                       f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		Currency15:                    f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{3}},
		ValidUntilTime62:              f0.Fld{Opt, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		IOIQltyInd25:                  f0.Fld{Opt, f0.ASCII, f0.String("M" /* MEDIUM */, "H" /* HIGH */, "L" /* LOW */), f0.Con{1}},
		IOINaturalFlag130:             f0.Fld{Opt, f0.ASCII, f0.BoolDefault(true /* YES */, false /* NO */), f0.Con{1}},
		Text58:                        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncodedTextLen354:             f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncodedText355:                f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		TransactTime60:                f0.Fld{Opt, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		URLLink149:                    f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		Benchmark219:                  f0.Fld{Opt, f0.ASCII, f0.String("5" /* OLD_10 */, "1" /* CURVE */, "2" /* 5_YR */, "4" /* 10_YR */, "6" /* 30_YR */, "7" /* OLD_30 */, "8" /* 3_MO_LIBOR */, "9" /* 6_MO_LIBOR */, "3" /* OLD_5 */), f0.Con{1}},
		Symbol55:                      f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SymbolSfx65:                   f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SecurityID48:                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SecurityIDSource22:            f0.Fld{Opt, f0.ASCII, f0.String("E" /* SICOVAM */, "2" /* SEDOL */, "1" /* CUSIP */, "3" /* QUIK */, "F" /* BELGIAN */, "D" /* VALOREN */, "C" /* DUTCH */, "B" /* WERTPAPIER */, "A" /* BLOOMBERG_SYMBOL */, "9" /* CONSOLIDATED_TAPE_ASSOCIATION */, "8" /* EXCHANGE_SYMBOL */, "7" /* ISO_COUNTRY_CODE */, "6" /* ISO_CURRENCY_CODE */, "5" /* RIC_CODE */, "4" /* ISIN_NUMBER */, "G" /* COMMON */), f0.Var{1, 65536}},
		Product460:                    f0.Fld{Opt, f0.ASCII, f0.IntDefault(8 /* LOAN */, 12 /* OTHER */, 11 /* MUNICIPAL */, 1 /* AGENCY */, 3 /* CORPORATE */, 4 /* CURRENCY */, 2 /* COMMODITY */, 6 /* GOVERNMENT */, 10 /* MORTGAGE */, 7 /* INDEX */, 9 /* MONEYMARKET */, 5 /* EQUITY */), f0.Var{1, 19}},
		CFICode461:                    f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SecurityType167:               f0.Fld{Opt, f0.ASCII, f0.String("CP" /* COMMERCIAL_PAPER */, "VRDN" /* VARIABLE_RATE_DEMAND_NOTE */, "PZFJ" /* PLAZOS_FIJOS */, "PN" /* PROMISSORY_NOTE */, "ONITE" /* OVERNIGHT */, "MTN" /* MEDIUM_TERM_NOTES */, "TECP" /* TAX_EXEMPT_COMMERCIAL_PAPER */, "AMENDED" /* AMENDED_RESTATED */, "BRIDGE" /* BRIDGE_LOAN */, "LOFC" /* LETTER_OF_CREDIT */, "SWING" /* SWING_LINE_FACILITY */, "DINP" /* DEBTOR_IN_POSSESSION */, "DEFLTED" /* DEFAULTED */, "WITHDRN" /* WITHDRAWN */, "LQN" /* LIQUIDITY_NOTE */, "MATURED" /* MATURED */, "DN" /* DEPOSIT_NOTES */, "RETIRED" /* RETIRED */, "BA" /* BANKERS_ACCEPTANCE */, "BN" /* BANK_NOTES */, "BOX" /* BILL_OF_EXCHANGES */, "CD" /* CERTIFICATE_OF_DEPOSIT */, "CL" /* CALL_LOANS */, "REPLACD" /* REPLACED */, "MT" /* MANDATORY_TENDER */, "RVLVTRM" /* REVOLVER_TERM_LOAN */, "MPP" /* MORTGAGE_PRIVATE_PLACEMENT */, "STN" /* SHORT_TERM_LOAN_NOTE */, "MPT" /* MISCELLANEOUS_PASS_THROUGH */, "TBA" /* TO_BE_ANNOUNCED */, "AN" /* OTHER_ANTICIPATION_NOTES_BAN_GAN_ETC */, "MIO" /* MORTGAGE_INTEREST_ONLY */, "COFP" /* CERTIFICATE_OF_PARTICIPATION */, "MBS" /* MORTGAGE_BACKED_SECURITIES */, "REV" /* REVENUE_BONDS */, "SPCLA" /* SPECIAL_ASSESSMENT */, "SPCLO" /* SPECIAL_OBLIGATION */, "SPCLT" /* SPECIAL_TAX */, "TAN" /* TAX_ANTICIPATION_NOTE */, "TAXA" /* TAX_ALLOCATION */, "COFO" /* CERTIFICATE_OF_OBLIGATION */, "TD" /* TIME_DEPOSIT */, "GO" /* GENERAL_OBLIGATION_BONDS */, "?" /* WILDCARD_ENTRY */, "WAR" /* WARRANT */, "MF" /* MUTUAL_FUND */, "MLEG" /* MULTI_LEG_INSTRUMENT */, "TRAN" /* TAX_REVENUE_ANTICIPATION_NOTE */, "MPO" /* MORTGAGE_PRINCIPAL_ONLY */, "RP" /* REPURCHASE_AGREEMENT */, "NONE" /* NO_SECURITY_TYPE */, "XCN" /* EXTENDED_COMM_NOTE */, "POOL" /* AGENCY_POOLS */, "ABS" /* ASSET_BACKED_SECURITIES */, "CMBS" /* CORP_MORTGAGE_BACKED_SECURITIES */, "CMO" /* COLLATERALIZED_MORTGAGE_OBLIGATION */, "IET" /* IOETTE_MORTGAGE */, "RVRP" /* REVERSE_REPURCHASE_AGREEMENT */, "FOR" /* FOREIGN_EXCHANGE_CONTRACT */, "RAN" /* REVENUE_ANTICIPATION_NOTE */, "RVLV" /* REVOLVER_LOAN */, "FAC" /* FEDERAL_AGENCY_COUPON */, "FADN" /* FEDERAL_AGENCY_DISCOUNT_NOTE */, "PEF" /* PRIVATE_EXPORT_FUNDING */, "CORP" /* CORPORATE_BOND */, "CPP" /* CORPORATE_PRIVATE_PLACEMENT */, "CB" /* CONVERTIBLE_BOND */, "DUAL" /* DUAL_CURRENCY */, "XLINKD" /* INDEXED_LINKED */, "YANK" /* YANKEE_CORPORATE_BOND */, "CS" /* COMMON_STOCK */, "PS" /* PREFERRED_STOCK */, "BRADY" /* BRADY_BOND */, "TBOND" /* US_TREASURY_BOND */, "TINT" /* INTEREST_STRIP_FROM_ANY_BOND_OR_NOTE */, "TIPS" /* TREASURY_INFLATION_PROTECTED_SECURITIES */, "TCAL" /* PRINCIPAL_STRIP_OF_A_CALLABLE_BOND_OR_NOTE */, "TPRN" /* PRINCIPAL_STRIP_FROM_A_NON_CALLABLE_BOND_OR_NOTE */, "UST" /* US_TREASURY_NOTE_BOND */, "USTB" /* US_TREASURY_BILL */, "TERM" /* TERM_LOAN */, "STRUCT" /* STRUCTURED_NOTES */), f0.Var{1, 65536}},
		MaturityMonthYear200:          f0.Fld{Opt, f0.ASCII, f0.MonthYearTime(), f0.Var{6, 8}},
		MaturityDate541:               f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		CouponPaymentDate224:          f0.Fld{Opt, f0.ASCII, f0.UTCDateOnlyTime(), f0.Con{8}},
		IssueDate225:                  f0.Fld{Opt, f0.ASCII, f0.UTCDateOnlyTime(), f0.Con{8}},
		RepoCollateralSecurityType239: f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		RepurchaseTerm226:             f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		RepurchaseRate227:             f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		Factor228:                     f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		CreditRating255:               f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		InstrRegistry543:              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		CountryOfIssue470:             f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{2}},
		StateOrProvinceOfIssue471:     f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		LocaleOfIssue472:              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RedemptionDate240:             f0.Fld{Opt, f0.ASCII, f0.UTCDateOnlyTime(), f0.Con{8}},
		StrikePrice202:                f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		OptAttribute206:               f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		ContractMultiplier231:         f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		CouponRate223:                 f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		SecurityExchange207:           f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		Issuer106:                     f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncodedIssuerLen348:           f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncodedIssuer349:              f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		SecurityDesc107:               f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncodedSecurityDescLen350:     f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncodedSecurityDesc351:        f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		SecurityAltID455:              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SecurityAltIDSource456:        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		Spread218:                     f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		BenchmarkCurveCurrency220:     f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{3}},
		BenchmarkCurveName221:         f0.Fld{Opt, f0.ASCII, f0.String("SWAP" /* SWAP */, "LIBID" /* LIBID */, "OTHER" /* OTHER */, "Treasury" /* TREASURY */, "Euribor" /* EURIBOR */, "Pfandbriefe" /* PFANDBRIEFE */, "FutureSWAP" /* FUTURESWAP */, "MuniAAA" /* MUNIAAA */, "LIBOR" /* LIBOR */), f0.Var{1, 65536}},
		BenchmarkCurvePoint222:        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SignatureLength93:             f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		Signature89:                   f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
	},
	BodyLength9: f0.BodyLengthFld{f0.ASCII, f0.BodyLength(), f0.Var{1, 18}},
	CheckSum10:  f0.CheckSumFld{f0.ASCII, f0.CheckSum(), f0.Con{3}},
	Unknown0:    f0.UnknownFld{f0.ASCII, f0.Unknown(), f0.Var{1, 65536}},
	Sort: []int{
		BeginString8,                  // STRING
		BodyLength9,                   // LENGTH
		MsgType35,                     // STRING
		SenderCompID49,                // STRING
		TargetCompID56,                // STRING
		OnBehalfOfCompID115,           // STRING
		DeliverToCompID128,            // STRING
		SecureDataLen90,               // LENGTH
		SecureData91,                  // DATA
		MsgSeqNum34,                   // SEQNUM
		SenderSubID50,                 // STRING
		SenderLocationID142,           // STRING
		TargetSubID57,                 // STRING
		TargetLocationID143,           // STRING
		OnBehalfOfSubID116,            // STRING
		OnBehalfOfLocationID144,       // STRING
		DeliverToSubID129,             // STRING
		DeliverToLocationID145,        // STRING
		PossDupFlag43,                 // BOOLEAN
		PossResend97,                  // BOOLEAN
		SendingTime52,                 // UTCTIMESTAMP
		OrigSendingTime122,            // UTCTIMESTAMP
		XmlDataLen212,                 // LENGTH
		XmlData213,                    // DATA
		MessageEncoding347,            // STRING
		LastMsgSeqNumProcessed369,     // SEQNUM
		OnBehalfOfSendingTime370,      // UTCTIMESTAMP
		HopCompID628,                  // STRING
		HopSendingTime629,             // UTCTIMESTAMP
		HopRefID630,                   // SEQNUM
		IOIid23,                       // STRING
		IOITransType28,                // CHAR
		IOIRefID26,                    // STRING
		Side54,                        // CHAR
		QuantityType465,               // INT
		IOIQty27,                      // STRING
		PriceType423,                  // INT
		Price44,                       // PRICE
		Currency15,                    // CURRENCY
		ValidUntilTime62,              // UTCTIMESTAMP
		IOIQltyInd25,                  // CHAR
		IOINaturalFlag130,             // BOOLEAN
		Text58,                        // STRING
		EncodedTextLen354,             // LENGTH
		EncodedText355,                // DATA
		TransactTime60,                // UTCTIMESTAMP
		URLLink149,                    // STRING
		Benchmark219,                  // CHAR
		Symbol55,                      // STRING
		SymbolSfx65,                   // STRING
		SecurityID48,                  // STRING
		SecurityIDSource22,            // STRING
		Product460,                    // INT
		CFICode461,                    // STRING
		SecurityType167,               // STRING
		MaturityMonthYear200,          // MONTHYEAR
		MaturityDate541,               // LOCALMKTDATE
		CouponPaymentDate224,          // UTCDATE
		IssueDate225,                  // UTCDATE
		RepoCollateralSecurityType239, // INT
		RepurchaseTerm226,             // INT
		RepurchaseRate227,             // PERCENTAGE
		Factor228,                     // FLOAT
		CreditRating255,               // STRING
		InstrRegistry543,              // STRING
		CountryOfIssue470,             // COUNTRY
		StateOrProvinceOfIssue471,     // STRING
		LocaleOfIssue472,              // STRING
		RedemptionDate240,             // UTCDATE
		StrikePrice202,                // PRICE
		OptAttribute206,               // CHAR
		ContractMultiplier231,         // FLOAT
		CouponRate223,                 // PERCENTAGE
		SecurityExchange207,           // EXCHANGE
		Issuer106,                     // STRING
		EncodedIssuerLen348,           // LENGTH
		EncodedIssuer349,              // DATA
		SecurityDesc107,               // STRING
		EncodedSecurityDescLen350,     // LENGTH
		EncodedSecurityDesc351,        // DATA
		SecurityAltID455,              // STRING
		SecurityAltIDSource456,        // STRING
		Spread218,                     // PRICEOFFSET
		BenchmarkCurveCurrency220,     // CURRENCY
		BenchmarkCurveName221,         // STRING
		BenchmarkCurvePoint222,        // STRING
		SignatureLength93,             // LENGTH
		Signature89,                   // DATA
		CheckSum10,                    // STRING
	},
}

const Req, Opt = true, false

const (
	BeginString8                  = 8   // STRING
	BodyLength9                   = 9   // LENGTH
	MsgType35                     = 35  // STRING
	SenderCompID49                = 49  // STRING
	TargetCompID56                = 56  // STRING
	OnBehalfOfCompID115           = 115 // STRING
	DeliverToCompID128            = 128 // STRING
	SecureDataLen90               = 90  // LENGTH
	SecureData91                  = 91  // DATA
	MsgSeqNum34                   = 34  // SEQNUM
	SenderSubID50                 = 50  // STRING
	SenderLocationID142           = 142 // STRING
	TargetSubID57                 = 57  // STRING
	TargetLocationID143           = 143 // STRING
	OnBehalfOfSubID116            = 116 // STRING
	OnBehalfOfLocationID144       = 144 // STRING
	DeliverToSubID129             = 129 // STRING
	DeliverToLocationID145        = 145 // STRING
	PossDupFlag43                 = 43  // BOOLEAN
	PossResend97                  = 97  // BOOLEAN
	SendingTime52                 = 52  // UTCTIMESTAMP
	OrigSendingTime122            = 122 // UTCTIMESTAMP
	XmlDataLen212                 = 212 // LENGTH
	XmlData213                    = 213 // DATA
	MessageEncoding347            = 347 // STRING
	LastMsgSeqNumProcessed369     = 369 // SEQNUM
	OnBehalfOfSendingTime370      = 370 // UTCTIMESTAMP
	HopCompID628                  = 628 // STRING
	HopSendingTime629             = 629 // UTCTIMESTAMP
	HopRefID630                   = 630 // SEQNUM
	IOIid23                       = 23  // STRING
	IOITransType28                = 28  // CHAR
	IOIRefID26                    = 26  // STRING
	Side54                        = 54  // CHAR
	QuantityType465               = 465 // INT
	IOIQty27                      = 27  // STRING
	PriceType423                  = 423 // INT
	Price44                       = 44  // PRICE
	Currency15                    = 15  // CURRENCY
	ValidUntilTime62              = 62  // UTCTIMESTAMP
	IOIQltyInd25                  = 25  // CHAR
	IOINaturalFlag130             = 130 // BOOLEAN
	Text58                        = 58  // STRING
	EncodedTextLen354             = 354 // LENGTH
	EncodedText355                = 355 // DATA
	TransactTime60                = 60  // UTCTIMESTAMP
	URLLink149                    = 149 // STRING
	Benchmark219                  = 219 // CHAR
	Symbol55                      = 55  // STRING
	SymbolSfx65                   = 65  // STRING
	SecurityID48                  = 48  // STRING
	SecurityIDSource22            = 22  // STRING
	Product460                    = 460 // INT
	CFICode461                    = 461 // STRING
	SecurityType167               = 167 // STRING
	MaturityMonthYear200          = 200 // MONTHYEAR
	MaturityDate541               = 541 // LOCALMKTDATE
	CouponPaymentDate224          = 224 // UTCDATE
	IssueDate225                  = 225 // UTCDATE
	RepoCollateralSecurityType239 = 239 // INT
	RepurchaseTerm226             = 226 // INT
	RepurchaseRate227             = 227 // PERCENTAGE
	Factor228                     = 228 // FLOAT
	CreditRating255               = 255 // STRING
	InstrRegistry543              = 543 // STRING
	CountryOfIssue470             = 470 // COUNTRY
	StateOrProvinceOfIssue471     = 471 // STRING
	LocaleOfIssue472              = 472 // STRING
	RedemptionDate240             = 240 // UTCDATE
	StrikePrice202                = 202 // PRICE
	OptAttribute206               = 206 // CHAR
	ContractMultiplier231         = 231 // FLOAT
	CouponRate223                 = 223 // PERCENTAGE
	SecurityExchange207           = 207 // EXCHANGE
	Issuer106                     = 106 // STRING
	EncodedIssuerLen348           = 348 // LENGTH
	EncodedIssuer349              = 349 // DATA
	SecurityDesc107               = 107 // STRING
	EncodedSecurityDescLen350     = 350 // LENGTH
	EncodedSecurityDesc351        = 351 // DATA
	SecurityAltID455              = 455 // STRING
	SecurityAltIDSource456        = 456 // STRING
	Spread218                     = 218 // PRICEOFFSET
	BenchmarkCurveCurrency220     = 220 // CURRENCY
	BenchmarkCurveName221         = 221 // STRING
	BenchmarkCurvePoint222        = 222 // STRING
	SignatureLength93             = 93  // LENGTH
	Signature89                   = 89  // DATA
	CheckSum10                    = 10  // STRING
)
