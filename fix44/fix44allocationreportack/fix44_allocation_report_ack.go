// Code generated by protofix; DO NOT EDIT.
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fix44allocationreportack is a format of the FIX.4.4 AllocationReportAck message.
package fix44allocationreportack

import (
	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/marshfix"
)

var (
	FIX44AllocationReportAckMarshaler   = marshfix.Marshal{Tag: "FIX44", Format: FIX44AllocationReportAck}
	FIX44AllocationReportAckUnmarshaler = marshfix.Unmarshal{Tag: "FIX44", Format: FIX44AllocationReportAck}
)

// FIX44AllocationReportAck is a FIX.4.4 format of the AllocationReportAck message which maps the codecs into individual fields.
var FIX44AllocationReportAck = f0.Format{
	Fields: map[int]f0.Codec{
		BeginString8:              f0.Fld{Req, f0.ASCII, f0.StringDefault("FIX.4.4"), f0.Con{7}},
		MsgType35:                 f0.Fld{Req, f0.ASCII, f0.String("0" /* HEARTBEAT */, "1" /* TESTREQUEST */, "2" /* RESENDREQUEST */, "3" /* REJECT */, "4" /* SEQUENCERESET */, "5" /* LOGOUT */, "6" /* IOI */, "7" /* ADVERTISEMENT */, "8" /* EXECUTIONREPORT */, "9" /* ORDERCANCELREJECT */, "a" /* QUOTESTATUSREQUEST */, "A" /* LOGON */, "AA" /* DERIVATIVESECURITYLIST */, "AB" /* NEWORDERMULTILEG */, "AC" /* MULTILEGORDERCANCELREPLACE */, "AD" /* TRADECAPTUREREPORTREQUEST */, "AE" /* TRADECAPTUREREPORT */, "AF" /* ORDERMASSSTATUSREQUEST */, "AG" /* QUOTEREQUESTREJECT */, "AH" /* RFQREQUEST */, "AI" /* QUOTESTATUSREPORT */, "AJ" /* QUOTERESPONSE */, "AK" /* CONFIRMATION */, "AL" /* POSITIONMAINTENANCEREQUEST */, "AM" /* POSITIONMAINTENANCEREPORT */, "AN" /* REQUESTFORPOSITIONS */, "AO" /* REQUESTFORPOSITIONSACK */, "AP" /* POSITIONREPORT */, "AQ" /* TRADECAPTUREREPORTREQUESTACK */, "AR" /* TRADECAPTUREREPORTACK */, "AS" /* ALLOCATIONREPORT */, "AT" /* ALLOCATIONREPORTACK */, "AU" /* CONFIRMATIONACK */, "AV" /* SETTLEMENTINSTRUCTIONREQUEST */, "AW" /* ASSIGNMENTREPORT */, "AX" /* COLLATERALREQUEST */, "AY" /* COLLATERALASSIGNMENT */, "AZ" /* COLLATERALRESPONSE */, "B" /* NEWS */, "b" /* MASSQUOTEACKNOWLEDGEMENT */, "BA" /* COLLATERALREPORT */, "BB" /* COLLATERALINQUIRY */, "BC" /* NETWORKCOUNTERPARTYSYSTEMSTATUSREQUEST */, "BD" /* NETWORKCOUNTERPARTYSYSTEMSTATUSRESPONSE */, "BE" /* USERREQUEST */, "BF" /* USERRESPONSE */, "BG" /* COLLATERALINQUIRYACK */, "BH" /* CONFIRMATIONREQUEST */, "C" /* EMAIL */, "c" /* SECURITYDEFINITIONREQUEST */, "d" /* SECURITYDEFINITION */, "D" /* NEWORDERSINGLE */, "e" /* SECURITYSTATUSREQUEST */, "E" /* NEWORDERLIST */, "F" /* ORDERCANCELREQUEST */, "f" /* SECURITYSTATUS */, "G" /* ORDERCANCELREPLACEREQUEST */, "g" /* TRADINGSESSIONSTATUSREQUEST */, "H" /* ORDERSTATUSREQUEST */, "h" /* TRADINGSESSIONSTATUS */, "i" /* MASSQUOTE */, "j" /* BUSINESSMESSAGEREJECT */, "J" /* ALLOCATIONINSTRUCTION */, "k" /* BIDREQUEST */, "K" /* LISTCANCELREQUEST */, "l" /* BIDRESPONSE */, "L" /* LISTEXECUTE */, "m" /* LISTSTRIKEPRICE */, "M" /* LISTSTATUSREQUEST */, "n" /* XMLNONFIX */, "N" /* LISTSTATUS */, "o" /* REGISTRATIONINSTRUCTIONS */, "p" /* REGISTRATIONINSTRUCTIONSRESPONSE */, "P" /* ALLOCATIONINSTRUCTIONACK */, "q" /* ORDERMASSCANCELREQUEST */, "Q" /* DONTKNOWTRADEDK */, "R" /* QUOTEREQUEST */, "r" /* ORDERMASSCANCELREPORT */, "S" /* QUOTE */, "s" /* NEWORDERCROSS */, "T" /* SETTLEMENTINSTRUCTIONS */, "t" /* CROSSORDERCANCELREPLACEREQUEST */, "u" /* CROSSORDERCANCELREQUEST */, "V" /* MARKETDATAREQUEST */, "v" /* SECURITYTYPEREQUEST */, "w" /* SECURITYTYPES */, "W" /* MARKETDATASNAPSHOTFULLREFRESH */, "x" /* SECURITYLISTREQUEST */, "X" /* MARKETDATAINCREMENTALREFRESH */, "Y" /* MARKETDATAREQUESTREJECT */, "y" /* SECURITYLIST */, "Z" /* QUOTECANCEL */, "z" /* DERIVATIVESECURITYLISTREQUEST */), f0.Var{1, 2}},
		SenderCompID49:            f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TargetCompID56:            f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		OnBehalfOfCompID115:       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		DeliverToCompID128:        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SecureDataLen90:           f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		SecureData91:              f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		MsgSeqNum34:               f0.Fld{Req, f0.ASCII, f0.SeqNum(), f0.Var{1, 19}},
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
		MessageEncoding347:        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		LastMsgSeqNumProcessed369: f0.Fld{Opt, f0.ASCII, f0.SeqNum(), f0.Var{1, 19}},
		HopCompID628:              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		HopSendingTime629:         f0.Fld{Opt, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		HopRefID630:               f0.Fld{Opt, f0.ASCII, f0.SeqNum(), f0.Var{1, 19}},
		AllocReportID755:          f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		AllocID70:                 f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SecondaryAllocID793:       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TradeDate75:               f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		TransactTime60:            f0.Fld{Req, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		AllocStatus87:             f0.Fld{Req, f0.ASCII, f0.IntDefault(0 /* ACCEPTEDSUCCESSFULLYPROCESSED */, 1 /* BLOCKLEVELREJECT */, 2 /* ACCOUNTLEVELREJECT */, 3 /* RECEIVEDRECEIVEDNOTYETPROCESSED */, 4 /* INCOMPLETE */, 5 /* REJECTEDBYINTERMEDIARY */), f0.Var{1, 19}},
		AllocRejCode88:            f0.Fld{Opt, f0.ASCII, f0.IntDefault(0 /* UNKNOWNACCT */, 1 /* INCORRECTQTY */, 10 /* UNKNOWNORSTALEEXECID */, 11 /* MISMATCHEDDATA */, 12 /* UNKNOWNCLORDID */, 13 /* WAREHOUSEREQUESTREJECTED */, 2 /* INCORRECTAVGPRC */, 3 /* INCORRECTBRKMNC */, 4 /* COMMDIFF */, 5 /* UNKNOWNORDID */, 6 /* UNKNOWNLISTID */, 7 /* OTHER */, 8 /* INCORRECTALLOCATEDQUANTITY */, 9 /* CALCULATIONDIFFERENCE */), f0.Var{1, 19}},
		AllocReportType794:        f0.Fld{Opt, f0.ASCII, f0.IntDefault(3 /* SELLSIDECALCULATEDUSINGPRELIMINARY */, 4 /* SELLSIDECALCULATEDWITHOUTPRELIMINARY */, 5 /* WAREHOUSERECAP */, 8 /* REQUESTTOINTERMEDIARY */), f0.Var{1, 19}},
		AllocIntermedReqType808:   f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* PENDINGACCEPT */, 2 /* PENDINGRELEASE */, 3 /* PENDINGREVERSAL */, 4 /* ACCEPT */, 5 /* BLOCKLEVELREJECT */, 6 /* ACCOUNTLEVELREJECT */), f0.Var{1, 19}},
		MatchStatus573:            f0.Fld{Opt, f0.ASCII, f0.String("0" /* COMPMATAFF */, "1" /* UNCOMPUNMATUNAFF */, "2" /* ADVALERT */), f0.Con{1}},
		Product460:                f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* AGENCY */, 10 /* MORTGAGE */, 11 /* MUNICIPAL */, 12 /* OTHER */, 13 /* FINANCING */, 2 /* COMMODITY */, 3 /* CORPORATE */, 4 /* CURRENCY */, 5 /* EQUITY */, 6 /* GOVERNMENT */, 7 /* INDEX */, 8 /* LOAN */, 9 /* MONEYMARKET */), f0.Var{1, 19}},
		SecurityType167:           f0.Fld{Opt, f0.ASCII, f0.String("ABS" /* ASSETBACKEDSECURITIES */, "AMENDED" /* AMENDEDRESTATED */, "AN" /* OTHERANTICIPATIONNOTESBANGANETC */, "BA" /* BANKERSACCEPTANCE */, "BN" /* BANKNOTES */, "BOX" /* BILLOFEXCHANGES */, "BRADY" /* BRADYBOND */, "BRIDGE" /* BRIDGELOAN */, "BUYSELL" /* BUYSELLBACK */, "CB" /* CONVERTIBLEBOND */, "CD" /* CERTIFICATEOFDEPOSIT */, "CL" /* CALLLOANS */, "CMBS" /* CORPMORTGAGEBACKEDSECURITIES */, "CMO" /* COLLATERALIZEDMORTGAGEOBLIGATION */, "COFO" /* CERTIFICATEOFOBLIGATION */, "COFP" /* CERTIFICATEOFPARTICIPATION */, "CORP" /* CORPORATEBOND */, "CP" /* COMMERCIALPAPER */, "CPP" /* CORPORATEPRIVATEPLACEMENT */, "CS" /* COMMONSTOCK */, "DEFLTED" /* DEFAULTED */, "DINP" /* DEBTORINPOSSESSION */, "DN" /* DEPOSITNOTES */, "DUAL" /* DUALCURRENCY */, "EUCD" /* EUROCERTIFICATEOFDEPOSIT */, "EUCORP" /* EUROCORPORATEBOND */, "EUCP" /* EUROCOMMERCIALPAPER */, "EUSOV" /* EUROSOVEREIGNS */, "EUSUPRA" /* EUROSUPRANATIONALCOUPONS */, "FAC" /* FEDERALAGENCYCOUPON */, "FADN" /* FEDERALAGENCYDISCOUNTNOTE */, "FOR" /* FOREIGNEXCHANGECONTRACT */, "FORWARD" /* FORWARD */, "FUT" /* FUTURE */, "GO" /* GENERALOBLIGATIONBONDS */, "IET" /* IOETTEMORTGAGE */, "LOFC" /* LETTEROFCREDIT */, "LQN" /* LIQUIDITYNOTE */, "MATURED" /* MATURED */, "MBS" /* MORTGAGEBACKEDSECURITIES */, "MF" /* MUTUALFUND */, "MIO" /* MORTGAGEINTERESTONLY */, "MLEG" /* MULTILEGINSTRUMENT */, "MPO" /* MORTGAGEPRINCIPALONLY */, "MPP" /* MORTGAGEPRIVATEPLACEMENT */, "MPT" /* MISCELLANEOUSPASSTHROUGH */, "MT" /* MANDATORYTENDER */, "MTN" /* MEDIUMTERMNOTES */, "NONE" /* NOSECURITYTYPE */, "ONITE" /* OVERNIGHT */, "OPT" /* OPTION */, "PEF" /* PRIVATEEXPORTFUNDING */, "PFAND" /* PFANDBRIEFE */, "PN" /* PROMISSORYNOTE */, "PS" /* PREFERREDSTOCK */, "PZFJ" /* PLAZOSFIJOS */, "RAN" /* REVENUEANTICIPATIONNOTE */, "REPLACD" /* REPLACED */, "REPO" /* REPURCHASE */, "RETIRED" /* RETIRED */, "REV" /* REVENUEBONDS */, "RVLV" /* REVOLVERLOAN */, "RVLVTRM" /* REVOLVERTERMLOAN */, "SECLOAN" /* SECURITIESLOAN */, "SECPLEDGE" /* SECURITIESPLEDGE */, "SPCLA" /* SPECIALASSESSMENT */, "SPCLO" /* SPECIALOBLIGATION */, "SPCLT" /* SPECIALTAX */, "STN" /* SHORTTERMLOANNOTE */, "STRUCT" /* STRUCTUREDNOTES */, "SUPRA" /* USDSUPRANATIONALCOUPONS */, "SWING" /* SWINGLINEFACILITY */, "TAN" /* TAXANTICIPATIONNOTE */, "TAXA" /* TAXALLOCATION */, "TBA" /* TOBEANNOUNCED */, "TBILL" /* USTREASURYBILL */, "TBOND" /* USTREASURYBOND */, "TCAL" /* PRINCIPALSTRIPOFACALLABLEBONDORNOTE */, "TD" /* TIMEDEPOSIT */, "TECP" /* TAXEXEMPTCOMMERCIALPAPER */, "TERM" /* TERMLOAN */, "TINT" /* INTERESTSTRIPFROMANYBONDORNOTE */, "TIPS" /* TREASURYINFLATIONPROTECTEDSECURITIES */, "TNOTE" /* USTREASURYNOTE */, "TPRN" /* PRINCIPALSTRIPFROMANONCALLABLEBONDORNOTE */, "TRAN" /* TAXREVENUEANTICIPATIONNOTE */, "UST" /* USTREASURYNOTEDEPRECATEDVALUEUSETNOTE */, "USTB" /* USTREASURYBILLDEPRECATEDVALUEUSETBILL */, "VRDN" /* VARIABLERATEDEMANDNOTE */, "WAR" /* WARRANT */, "WITHDRN" /* WITHDRAWN */, "WLD" /* WILDCARDENTRY */, "XCN" /* EXTENDEDCOMMNOTE */, "XLINKD" /* INDEXEDLINKED */, "YANK" /* YANKEECORPORATEBOND */, "YCD" /* YANKEECERTIFICATEOFDEPOSIT */), f0.Var{1, 65536}},
		Text58:                    f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncodedTextLen354:         f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncodedText355:            f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		PartyID448:                f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		PartyIDSource447:          f0.Fld{Opt, f0.ASCII, f0.String("1" /* KOREANINVESTORID */, "2" /* TAIWANESEQUALIFIED */, "3" /* TAIWANESETRADINGACCT */, "4" /* MCDNUMBER */, "5" /* CHINESEBSHARE */, "6" /* UKNATIONALINSPENNUMBER */, "7" /* USSOCIALSECURITY */, "8" /* USEMPLOYERIDNUMBER */, "9" /* AUSTRALIANBUSINESSNUMBER */, "A" /* AUSTRALIANTAXFILENUMBER */, "B" /* BIC */, "C" /* ACCPTMARKETPART */, "D" /* PROPCODE */, "E" /* ISOCODE */, "F" /* SETTLENTLOC */, "G" /* MIC */, "H" /* CSDPARTCODE */, "I" /* DIRECTEDDEFINEDISITC */), f0.Con{1}},
		PartyRole452:              f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* EXECUTINGFIRM */, 10 /* SETTLEMENTLOCATION */, 11 /* INITIATINGTRADER */, 12 /* EXECUTINGTRADER */, 13 /* ORDERORIGINATOR */, 14 /* GIVEUPCLEARINGFIRM */, 15 /* CORRESPONDANTCLEARINGFIRM */, 16 /* EXECUTINGSYSTEM */, 17 /* CONTRAFIRM */, 18 /* CONTRACLEARINGFIRM */, 19 /* SPONSORINGFIRM */, 2 /* BROKEROFCREDIT */, 20 /* UNDRCONTRAFIRM */, 21 /* CLEARINGORGANIZATION */, 22 /* EXCHANGE */, 24 /* CUSTOMERACCOUNT */, 25 /* CORRESPONDENTCLEARINGORGANIZATION */, 26 /* CORRESPONDENTBROKER */, 27 /* BUYERSELLERRECEIVERDELIVERER */, 28 /* CUSTODIAN */, 29 /* INTERMEDIARY */, 3 /* CLIENTID */, 30 /* AGENT */, 31 /* SUBCUSTODIAN */, 32 /* BENEFICIARY */, 33 /* INTERESTEDPARTY */, 34 /* REGULATORYBODY */, 35 /* LIQUIDITYPROVIDER */, 36 /* ENTERINGTRADER */, 37 /* CONTRATRADER */, 38 /* POSITIONACCOUNT */, 39 /* ALLOCENTITY */, 4 /* CLEARINGFIRM */, 5 /* INVESTORID */, 6 /* INTRODUCINGFIRM */, 7 /* ENTERINGFIRM */, 8 /* LOCATELENDINGFIRM */, 9 /* FUNDMANAGER */), f0.Var{1, 19}},
		PartySubID523:             f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		PartySubIDType803:         f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* FIRM */, 10 /* SECURITIESACCOUNTNUMBER */, 11 /* REGISTRATIONNUMBER */, 12 /* REGISTEREDADDRESS_12 */, 13 /* REGULATORYSTATUS */, 14 /* REGISTRATIONNAME */, 15 /* CASHACCOUNT */, 16 /* BIC */, 17 /* CSDPARTICIPANTMEMBERCODE */, 18 /* REGISTEREDADDRESS_18 */, 19 /* FUNDACCOUNTNAME */, 2 /* PERSON */, 20 /* TELEXNUMBER */, 21 /* FAXNUMBER */, 22 /* SECURITIESACCOUNTNAME */, 23 /* CASHACCOUNTNAME */, 24 /* DEPARTMENT */, 25 /* LOCATIONDESK */, 26 /* POSITIONACCOUNTTYPE */, 3 /* SYSTEM */, 4 /* APPLICATION */, 4000 /* RESERVEDANDAVAILABLEFORBILATERALLYAGREEDUPONUSERDEFINEDVALUES */, 5 /* FULLLEGALNAMEOFFIRM */, 6 /* POSTALADDRESS */, 7 /* PHONENUMBER */, 8 /* EMAILADDRESS */, 9 /* CONTACTNAME */), f0.Var{1, 19}},
		SignatureLength93:         f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		Signature89:               f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
	},
	BodyLength9: f0.BodyLengthFld{f0.ASCII, f0.BodyLength(), f0.Var{1, 18}},
	CheckSum10:  f0.CheckSumFld{f0.ASCII, f0.CheckSum(), f0.Con{3}},
	Unknown0:    f0.UnknownFld{f0.ASCII, f0.Unknown(), f0.Var{1, 65536}},
	Sort: []int{
		BeginString8,              // STRING
		BodyLength9,               // LENGTH
		MsgType35,                 // STRING
		SenderCompID49,            // STRING
		TargetCompID56,            // STRING
		OnBehalfOfCompID115,       // STRING
		DeliverToCompID128,        // STRING
		SecureDataLen90,           // LENGTH
		SecureData91,              // DATA
		MsgSeqNum34,               // SEQNUM
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
		LastMsgSeqNumProcessed369, // SEQNUM
		HopCompID628,              // STRING
		HopSendingTime629,         // UTCTIMESTAMP
		HopRefID630,               // SEQNUM
		AllocReportID755,          // STRING
		AllocID70,                 // STRING
		SecondaryAllocID793,       // STRING
		TradeDate75,               // LOCALMKTDATE
		TransactTime60,            // UTCTIMESTAMP
		AllocStatus87,             // INT
		AllocRejCode88,            // INT
		AllocReportType794,        // INT
		AllocIntermedReqType808,   // INT
		MatchStatus573,            // CHAR
		Product460,                // INT
		SecurityType167,           // STRING
		Text58,                    // STRING
		EncodedTextLen354,         // LENGTH
		EncodedText355,            // DATA
		PartyID448,                // STRING
		PartyIDSource447,          // CHAR
		PartyRole452,              // INT
		PartySubID523,             // STRING
		PartySubIDType803,         // INT
		SignatureLength93,         // LENGTH
		Signature89,               // DATA
		CheckSum10,                // STRING
	},
}

const Req, Opt = true, false

const (
	BeginString8              = 8   // STRING
	BodyLength9               = 9   // LENGTH
	MsgType35                 = 35  // STRING
	SenderCompID49            = 49  // STRING
	TargetCompID56            = 56  // STRING
	OnBehalfOfCompID115       = 115 // STRING
	DeliverToCompID128        = 128 // STRING
	SecureDataLen90           = 90  // LENGTH
	SecureData91              = 91  // DATA
	MsgSeqNum34               = 34  // SEQNUM
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
	LastMsgSeqNumProcessed369 = 369 // SEQNUM
	HopCompID628              = 628 // STRING
	HopSendingTime629         = 629 // UTCTIMESTAMP
	HopRefID630               = 630 // SEQNUM
	AllocReportID755          = 755 // STRING
	AllocID70                 = 70  // STRING
	SecondaryAllocID793       = 793 // STRING
	TradeDate75               = 75  // LOCALMKTDATE
	TransactTime60            = 60  // UTCTIMESTAMP
	AllocStatus87             = 87  // INT
	AllocRejCode88            = 88  // INT
	AllocReportType794        = 794 // INT
	AllocIntermedReqType808   = 808 // INT
	MatchStatus573            = 573 // CHAR
	Product460                = 460 // INT
	SecurityType167           = 167 // STRING
	Text58                    = 58  // STRING
	EncodedTextLen354         = 354 // LENGTH
	EncodedText355            = 355 // DATA
	PartyID448                = 448 // STRING
	PartyIDSource447          = 447 // CHAR
	PartyRole452              = 452 // INT
	PartySubID523             = 523 // STRING
	PartySubIDType803         = 803 // INT
	SignatureLength93         = 93  // LENGTH
	Signature89               = 89  // DATA
	CheckSum10                = 10  // STRING
)
