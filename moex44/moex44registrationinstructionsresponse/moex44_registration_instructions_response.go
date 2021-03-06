// Code generated by protofix; DO NOT EDIT.
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package moex44registrationinstructionsresponse is a format of the MOEX.4.4 (FIX.4.4) RegistrationInstructionsResponse message.
package moex44registrationinstructionsresponse

import (
	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/marshfix"
)

var (
	MOEX44RegistrationInstructionsResponseMarshaler   = marshfix.Marshal{Tag: "MOEX44", Format: MOEX44RegistrationInstructionsResponse}
	MOEX44RegistrationInstructionsResponseUnmarshaler = marshfix.Unmarshal{Tag: "MOEX44", Format: MOEX44RegistrationInstructionsResponse}
)

// MOEX44RegistrationInstructionsResponse is a MOEX.4.4 (FIX.4.4) format of the RegistrationInstructionsResponse message which maps the codecs into individual fields.
var MOEX44RegistrationInstructionsResponse = f0.Format{
	Fields: map[int]f0.Codec{
		BeginString8:           f0.Fld{Req, f0.ASCII, f0.StringDefault("FIX.4.4"), f0.Con{7}},
		MsgType35:              f0.Fld{Req, f0.ASCII, f0.String("0" /* HEARTBEAT */, "1" /* TESTREQUEST */, "2" /* RESENDREQUEST */, "3" /* REJECT */, "4" /* SEQUENCERESET */, "5" /* LOGOUT */, "6" /* IOI */, "7" /* ADVERTISEMENT */, "8" /* EXECUTIONREPORT */, "9" /* ORDERCANCELREJECT */, "a" /* QUOTESTATUSREQUEST */, "A" /* LOGON */, "AA" /* DERIVATIVESECURITYLIST */, "AB" /* NEWORDERMULTILEG */, "AC" /* MULTILEGORDERCANCELREPLACE */, "AD" /* TRADECAPTUREREPORTREQUEST */, "AE" /* TRADECAPTUREREPORT */, "AF" /* ORDERMASSSTATUSREQUEST */, "AG" /* QUOTEREQUESTREJECT */, "AH" /* RFQREQUEST */, "AI" /* QUOTESTATUSREPORT */, "AJ" /* QUOTERESPONSE */, "AK" /* CONFIRMATION */, "AL" /* POSITIONMAINTENANCEREQUEST */, "AM" /* POSITIONMAINTENANCEREPORT */, "AN" /* REQUESTFORPOSITIONS */, "AO" /* REQUESTFORPOSITIONSACK */, "AP" /* POSITIONREPORT */, "AQ" /* TRADECAPTUREREPORTREQUESTACK */, "AR" /* TRADECAPTUREREPORTACK */, "AS" /* ALLOCATIONREPORT */, "AT" /* ALLOCATIONREPORTACK */, "AU" /* CONFIRMATIONACK */, "AV" /* SETTLEMENTINSTRUCTIONREQUEST */, "AW" /* ASSIGNMENTREPORT */, "AX" /* COLLATERALREQUEST */, "AY" /* COLLATERALASSIGNMENT */, "AZ" /* COLLATERALRESPONSE */, "B" /* NEWS */, "b" /* MASSQUOTEACKNOWLEDGEMENT */, "BA" /* COLLATERALREPORT */, "BB" /* COLLATERALINQUIRY */, "BC" /* NETWORKCOUNTERPARTYSYSTEMSTATUSREQUEST */, "BD" /* NETWORKCOUNTERPARTYSYSTEMSTATUSRESPONSE */, "BE" /* USERREQUEST */, "BF" /* USERRESPONSE */, "BG" /* COLLATERALINQUIRYACK */, "BH" /* CONFIRMATIONREQUEST */, "C" /* EMAIL */, "c" /* SECURITYDEFINITIONREQUEST */, "d" /* SECURITYDEFINITION */, "D" /* NEWORDERSINGLE */, "e" /* SECURITYSTATUSREQUEST */, "E" /* NEWORDERLIST */, "F" /* ORDERCANCELREQUEST */, "f" /* SECURITYSTATUS */, "G" /* ORDERCANCELREPLACEREQUEST */, "g" /* TRADINGSESSIONSTATUSREQUEST */, "H" /* ORDERSTATUSREQUEST */, "h" /* TRADINGSESSIONSTATUS */, "i" /* MASSQUOTE */, "j" /* BUSINESSMESSAGEREJECT */, "J" /* ALLOCATIONINSTRUCTION */, "k" /* BIDREQUEST */, "K" /* LISTCANCELREQUEST */, "l" /* BIDRESPONSE */, "L" /* LISTEXECUTE */, "m" /* LISTSTRIKEPRICE */, "M" /* LISTSTATUSREQUEST */, "n" /* XMLNONFIX */, "N" /* LISTSTATUS */, "o" /* REGISTRATIONINSTRUCTIONS */, "p" /* REGISTRATIONINSTRUCTIONSRESPONSE */, "P" /* ALLOCATIONINSTRUCTIONACK */, "q" /* ORDERMASSCANCELREQUEST */, "Q" /* DONTKNOWTRADEDK */, "R" /* QUOTEREQUEST */, "r" /* ORDERMASSCANCELREPORT */, "S" /* QUOTE */, "s" /* NEWORDERCROSS */, "T" /* SETTLEMENTINSTRUCTIONS */, "t" /* CROSSORDERCANCELREPLACEREQUEST */, "u" /* CROSSORDERCANCELREQUEST */, "V" /* MARKETDATAREQUEST */, "v" /* SECURITYTYPEREQUEST */, "w" /* SECURITYTYPES */, "W" /* MARKETDATASNAPSHOTFULLREFRESH */, "x" /* SECURITYLISTREQUEST */, "X" /* MARKETDATAINCREMENTALREFRESH */, "Y" /* MARKETDATAREQUESTREJECT */, "y" /* SECURITYLIST */, "Z" /* QUOTECANCEL */, "z" /* DERIVATIVESECURITYLISTREQUEST */), f0.Var{1, 2}},
		SenderCompID49:         f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TargetCompID56:         f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		MsgSeqNum34:            f0.Fld{Opt, f0.ASCII, f0.SeqNum(), f0.Var{1, 19}},
		PossDupFlag43:          f0.Fld{Opt, f0.ASCII, f0.BoolDefault(false /* NO */, true /* YES */), f0.Con{1}},
		PossResend97:           f0.Fld{Opt, f0.ASCII, f0.BoolDefault(false /* NO */, true /* YES */), f0.Con{1}},
		SendingTime52:          f0.Fld{Req, f0.ASCII, f0.UTCTimestampNanosecondTime(), f0.Var{1, 27}},
		OrigSendingTime122:     f0.Fld{Opt, f0.ASCII, f0.UTCTimestampNanosecondTime(), f0.Var{1, 27}},
		RegistID513:            f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RegistTransType514:     f0.Fld{Req, f0.ASCII, f0.String("0" /* NEW */, "1" /* REPLACE */, "2" /* CANCEL */), f0.Con{1}},
		RegistRefID508:         f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		ClOrdID11:              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		Account1:               f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		AcctIDSource660:        f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* BIC */, 2 /* SIDCODE */, 3 /* TFMGSPTA */, 4 /* OMGEOALERTID */, 5 /* DTCCCODE */, 99 /* OTHER */), f0.Var{1, 19}},
		RegistStatus506:        f0.Fld{Req, f0.ASCII, f0.String("A" /* ACCEPTED */, "H" /* HELD */, "N" /* REMINDER_IE_REGISTRATION_INSTRUCTIONS_ARE_STILL_OUTSTANDING */, "R" /* REJECTED */), f0.Con{1}},
		RegistRejReasonCode507: f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* INVALIDACCOUNTTYPE */, 10 /* INVALIDINVESTORIDSOURCE */, 11 /* INVALIDDATEOFBIRTH */, 12 /* INVALIDINVESTORCOUNTRYOFRESIDENCE */, 13 /* INVALIDNODISTRIBINSTNS */, 14 /* INVALIDDISTRIBPERCENTAGE */, 15 /* INVALIDDISTRIBPAYMENTMETHOD */, 16 /* INVALIDCASHDISTRIBAGENTACCTNAME */, 17 /* INVALIDCASHDISTRIBAGENTCODE */, 18 /* INVALIDCASHDISTRIBAGENTACCTNUM */, 2 /* INVALIDTAXEXEMPTTYPE */, 3 /* INVALIDOWNERSHIPTYPE */, 4 /* INVALIDNOREGDETLS */, 5 /* INVALIDREGSEQNO */, 6 /* INVALIDREGDTLS */, 7 /* INVALIDMAILINGDTLS */, 8 /* INVALIDMAILINGINST */, 9 /* INVALIDINVESTORID */, 99 /* OTHER */), f0.Var{1, 19}},
		RegistRejReasonText496: f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		PartyID448:             f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		PartyIDSource447:       f0.Fld{Opt, f0.ASCII, f0.String("1" /* KOREANINVESTORID */, "2" /* TAIWANESEQUALIFIED */, "3" /* TAIWANESETRADINGACCT */, "4" /* MCDNUMBER */, "5" /* CHINESEBSHARE */, "6" /* UKNATIONALINSPENNUMBER */, "7" /* USSOCIALSECURITY */, "8" /* USEMPLOYERIDNUMBER */, "9" /* AUSTRALIANBUSINESSNUMBER */, "A" /* AUSTRALIANTAXFILENUMBER */, "B" /* BIC */, "C" /* ACCPTMARKETPART */, "D" /* PROPCODE */, "E" /* ISOCODE */, "F" /* SETTLENTLOC */, "G" /* MIC */, "H" /* CSDPARTCODE */, "I" /* DIRECTEDDEFINEDISITC */), f0.Con{1}},
		PartyRole452:           f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* EXECUTINGFIRM */, 10 /* SETTLEMENTLOCATION */, 11 /* INITIATINGTRADER */, 12 /* EXECUTINGTRADER */, 13 /* ORDERORIGINATOR */, 14 /* GIVEUPCLEARINGFIRM */, 15 /* CORRESPONDANTCLEARINGFIRM */, 16 /* EXECUTINGSYSTEM */, 17 /* CONTRAFIRM */, 18 /* CONTRACLEARINGFIRM */, 19 /* SPONSORINGFIRM */, 2 /* BROKEROFCREDIT */, 20 /* UNDRCONTRAFIRM */, 21 /* CLEARINGORGANIZATION */, 22 /* EXCHANGE */, 24 /* CUSTOMERACCOUNT */, 25 /* CORRESPONDENTCLEARINGORGANIZATION */, 26 /* CORRESPONDENTBROKER */, 27 /* BUYERSELLERRECEIVERDELIVERER */, 28 /* CUSTODIAN */, 29 /* INTERMEDIARY */, 3 /* CLIENTID */, 30 /* AGENT */, 31 /* SUBCUSTODIAN */, 32 /* BENEFICIARY */, 33 /* INTERESTEDPARTY */, 34 /* REGULATORYBODY */, 35 /* LIQUIDITYPROVIDER */, 36 /* ENTERINGTRADER */, 37 /* CONTRATRADER */, 38 /* POSITIONACCOUNT */, 39 /* ALLOCENTITY */, 4 /* CLEARINGFIRM */, 5 /* INVESTORID */, 6 /* INTRODUCINGFIRM */, 7 /* ENTERINGFIRM */, 8 /* LOCATELENDINGFIRM */, 9 /* FUNDMANAGER */), f0.Var{1, 19}},
		PartySubID523:          f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		PartySubIDType803:      f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* FIRM */, 10 /* SECURITIESACCOUNTNUMBER */, 11 /* REGISTRATIONNUMBER */, 12 /* REGISTEREDADDRESS_12 */, 13 /* REGULATORYSTATUS */, 14 /* REGISTRATIONNAME */, 15 /* CASHACCOUNT */, 16 /* BIC */, 17 /* CSDPARTICIPANTMEMBERCODE */, 18 /* REGISTEREDADDRESS_18 */, 19 /* FUNDACCOUNTNAME */, 2 /* PERSON */, 20 /* TELEXNUMBER */, 21 /* FAXNUMBER */, 22 /* SECURITIESACCOUNTNAME */, 23 /* CASHACCOUNTNAME */, 24 /* DEPARTMENT */, 25 /* LOCATIONDESK */, 26 /* POSITIONACCOUNTTYPE */, 3 /* SYSTEM */, 4 /* APPLICATION */, 4000 /* RESERVEDANDAVAILABLEFORBILATERALLYAGREEDUPONUSERDEFINEDVALUES */, 5 /* FULLLEGALNAMEOFFIRM */, 6 /* POSTALADDRESS */, 7 /* PHONENUMBER */, 8 /* EMAILADDRESS */, 9 /* CONTACTNAME */), f0.Var{1, 19}},
	},
	BodyLength9: f0.BodyLengthFld{f0.ASCII, f0.BodyLength(), f0.Var{1, 18}},
	CheckSum10:  f0.CheckSumFld{f0.ASCII, f0.CheckSum(), f0.Con{3}},
	Unknown0:    f0.UnknownFld{f0.ASCII, f0.Unknown(), f0.Var{1, 65536}},
	Sort: []int{
		BeginString8,           // STRING
		BodyLength9,            // LENGTH
		MsgType35,              // STRING
		SenderCompID49,         // STRING
		TargetCompID56,         // STRING
		MsgSeqNum34,            // SEQNUM
		PossDupFlag43,          // BOOLEAN
		PossResend97,           // BOOLEAN
		SendingTime52,          // UTCTIMESTAMP
		OrigSendingTime122,     // UTCTIMESTAMP
		RegistID513,            // STRING
		RegistTransType514,     // CHAR
		RegistRefID508,         // STRING
		ClOrdID11,              // STRING
		Account1,               // STRING
		AcctIDSource660,        // INT
		RegistStatus506,        // CHAR
		RegistRejReasonCode507, // INT
		RegistRejReasonText496, // STRING
		PartyID448,             // STRING
		PartyIDSource447,       // CHAR
		PartyRole452,           // INT
		PartySubID523,          // STRING
		PartySubIDType803,      // INT
		CheckSum10,             // STRING
	},
}

const Req, Opt = true, false

const (
	BeginString8           = 8   // STRING
	BodyLength9            = 9   // LENGTH
	MsgType35              = 35  // STRING
	SenderCompID49         = 49  // STRING
	TargetCompID56         = 56  // STRING
	MsgSeqNum34            = 34  // SEQNUM
	PossDupFlag43          = 43  // BOOLEAN
	PossResend97           = 97  // BOOLEAN
	SendingTime52          = 52  // UTCTIMESTAMP
	OrigSendingTime122     = 122 // UTCTIMESTAMP
	RegistID513            = 513 // STRING
	RegistTransType514     = 514 // CHAR
	RegistRefID508         = 508 // STRING
	ClOrdID11              = 11  // STRING
	Account1               = 1   // STRING
	AcctIDSource660        = 660 // INT
	RegistStatus506        = 506 // CHAR
	RegistRejReasonCode507 = 507 // INT
	RegistRejReasonText496 = 496 // STRING
	PartyID448             = 448 // STRING
	PartyIDSource447       = 447 // CHAR
	PartyRole452           = 452 // INT
	PartySubID523          = 523 // STRING
	PartySubIDType803      = 803 // INT
	CheckSum10             = 10  // STRING
)
