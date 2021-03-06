// Code generated by protofix; DO NOT EDIT.
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fix42executionreport is a format of the FIX.4.2 ExecutionReport message.
package fix42executionreport

import (
	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/marshfix"
)

var (
	FIX42ExecutionReportMarshaler   = marshfix.Marshal{Tag: "FIX42", Format: FIX42ExecutionReport}
	FIX42ExecutionReportUnmarshaler = marshfix.Unmarshal{Tag: "FIX42", Format: FIX42ExecutionReport}
)

// FIX42ExecutionReport is a FIX.4.2 format of the ExecutionReport message which maps the codecs into individual fields.
var FIX42ExecutionReport = f0.Format{
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
		OrderID37:                 f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SecondaryOrderID198:       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		ClOrdID11:                 f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		OrigClOrdID41:             f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		ClientID109:               f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		ExecBroker76:              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		ListID66:                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		ExecID17:                  f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		ExecTransType20:           f0.Fld{Req, f0.ASCII, f0.String("0" /* NEW */, "1" /* CANCEL */, "2" /* CORRECT */, "3" /* STATUS */), f0.Con{1}},
		ExecRefID19:               f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		ExecType150:               f0.Fld{Req, f0.ASCII, f0.String("0" /* NEW */, "1" /* PARTIAL_FILL */, "2" /* FILL */, "3" /* DONE_FOR_DAY */, "4" /* CANCELED */, "5" /* REPLACE */, "6" /* PENDING_CANCEL */, "7" /* STOPPED */, "8" /* REJECTED */, "9" /* SUSPENDED */, "A" /* PENDING_NEW */, "B" /* CALCULATED */, "C" /* EXPIRED */, "D" /* RESTATED */, "E" /* PENDING_REPLACE */), f0.Con{1}},
		OrdStatus39:               f0.Fld{Req, f0.ASCII, f0.String("0" /* NEW */, "1" /* PARTIALLY_FILLED */, "2" /* FILLED */, "3" /* DONE_FOR_DAY */, "4" /* CANCELED */, "5" /* REPLACED */, "6" /* PENDING_CANCEL */, "7" /* STOPPED */, "8" /* REJECTED */, "9" /* SUSPENDED */, "A" /* PENDING_NEW */, "B" /* CALCULATED */, "C" /* EXPIRED */, "D" /* ACCEPTED_FOR_BIDDING */, "E" /* PENDING_REPLACE */), f0.Con{1}},
		OrdRejReason103:           f0.Fld{Opt, f0.ASCII, f0.IntDefault(0 /* BROKER_OPTION */, 1 /* UNKNOWN_SYMBOL */, 2 /* EXCHANGE_CLOSED */, 3 /* ORDER_EXCEEDS_LIMIT */, 4 /* TOO_LATE_TO_ENTER */, 5 /* UNKNOWN_ORDER */, 6 /* DUPLICATE_ORDER */, 7 /* DUPLICATE_OF_A_VERBALLY_COMMUNICATED_ORDER */, 8 /* STALE_ORDER */), f0.Var{1, 19}},
		ExecRestatementReason378:  f0.Fld{Opt, f0.ASCII, f0.IntDefault(0 /* GT_CORPORATE_ACTION */, 1 /* GT_RENEWAL */, 2 /* VERBAL_CHANGE */, 3 /* REPRICING_OF_ORDER */, 4 /* BROKER_OPTION */, 5 /* PARTIAL_DECLINE_OF_ORDERQTY */), f0.Var{1, 19}},
		Account1:                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SettlmntTyp63:             f0.Fld{Opt, f0.ASCII, f0.String("0" /* REGULAR */, "1" /* CASH */, "2" /* NEXT_DAY */, "3" /* T_PLUS_2 */, "4" /* T_PLUS_3 */, "5" /* T_PLUS_4 */, "6" /* FUTURE */, "7" /* WHEN_ISSUED */, "8" /* SELLERS_OPTION */, "9" /* T_PLUS_5 */), f0.Con{1}},
		FutSettDate64:             f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
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
		OrderQty38:                f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		CashOrderQty152:           f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		OrdType40:                 f0.Fld{Opt, f0.ASCII, f0.String("1" /* MARKET */, "2" /* LIMIT */, "3" /* STOP */, "4" /* STOP_LIMIT */, "5" /* MARKET_ON_CLOSE */, "6" /* WITH_OR_WITHOUT */, "7" /* LIMIT_OR_BETTER */, "8" /* LIMIT_WITH_OR_WITHOUT */, "9" /* ON_BASIS */, "A" /* ON_CLOSE */, "B" /* LIMIT_ON_CLOSE */, "C" /* FOREX_C */, "D" /* PREVIOUSLY_QUOTED */, "E" /* PREVIOUSLY_INDICATED */, "F" /* FOREX_F */, "G" /* FOREX_G */, "H" /* FOREX_H */, "I" /* FUNARI */, "P" /* PEGGED */), f0.Con{1}},
		Price44:                   f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		StopPx99:                  f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		PegDifference211:          f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		DiscretionInst388:         f0.Fld{Opt, f0.ASCII, f0.String("0" /* RELATED_TO_DISPLAYED_PRICE */, "1" /* RELATED_TO_MARKET_PRICE */, "2" /* RELATED_TO_PRIMARY_PRICE */, "3" /* RELATED_TO_LOCAL_PRIMARY_PRICE */, "4" /* RELATED_TO_MIDPOINT_PRICE */, "5" /* RELATED_TO_LAST_TRADE_PRICE */), f0.Con{1}},
		DiscretionOffset389:       f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		Currency15:                f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{3}},
		ComplianceID376:           f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SolicitedFlag377:          f0.Fld{Opt, f0.ASCII, f0.BoolDefault(false /* NO */, true /* YES */), f0.Con{1}},
		TimeInForce59:             f0.Fld{Opt, f0.ASCII, f0.String("0" /* DAY */, "1" /* GOOD_TILL_CANCEL */, "2" /* AT_THE_OPENING */, "3" /* IMMEDIATE_OR_CANCEL */, "4" /* FILL_OR_KILL */, "5" /* GOOD_TILL_CROSSING */, "6" /* GOOD_TILL_DATE */), f0.Con{1}},
		EffectiveTime168:          f0.Fld{Opt, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		ExpireDate432:             f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		ExpireTime126:             f0.Fld{Opt, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		ExecInst18:                f0.Fld{Opt, f0.ASCII, f0.String("0" /* STAY_ON_OFFERSIDE */, "1" /* NOT_HELD */, "2" /* WORK */, "3" /* GO_ALONG */, "4" /* OVER_THE_DAY */, "5" /* HELD */, "6" /* PARTICIPATE_DONT_INITIATE */, "7" /* STRICT_SCALE */, "8" /* TRY_TO_SCALE */, "9" /* STAY_ON_BIDSIDE */, "A" /* NO_CROSS */, "B" /* OK_TO_CROSS */, "C" /* CALL_FIRST */, "D" /* PERCENT_OF_VOLUME */, "E" /* DO_NOT_INCREASE */, "F" /* DO_NOT_REDUCE */, "G" /* ALL_OR_NONE */, "I" /* INSTITUTIONS_ONLY */, "L" /* LAST_PEG */, "M" /* MID_PRICE_PEG */, "N" /* NON_NEGOTIABLE */, "O" /* OPENING_PEG */, "P" /* MARKET_PEG */, "R" /* PRIMARY_PEG */, "S" /* SUSPEND */, "T" /* FIXED_PEG_TO_LOCAL_BEST_BID_OR_OFFER_AT_TIME_OF_ORDER */, "U" /* CUSTOMER_DISPLAY_INSTRUCTION */, "V" /* NETTING */, "W" /* PEG_TO_VWAP */), f0.Var{1, 65536}},
		Rule80A47:                 f0.Fld{Opt, f0.ASCII, f0.String("A" /* AGENCY_SINGLE_ORDER */, "B" /* SHORT_EXEMPT_TRANSACTION_B */, "C" /* PROGRAM_ORDER_NON_INDEX_ARB_FOR_MEMBER_FIRM_ORG */, "D" /* PROGRAM_ORDER_INDEX_ARB_FOR_MEMBER_FIRM_ORG */, "E" /* REGISTERED_EQUITY_MARKET_MAKER_TRADES */, "F" /* SHORT_EXEMPT_TRANSACTION_F */, "H" /* SHORT_EXEMPT_TRANSACTION_H */, "I" /* INDIVIDUAL_INVESTOR_SINGLE_ORDER */, "J" /* PROGRAM_ORDER_INDEX_ARB_FOR_INDIVIDUAL_CUSTOMER */, "K" /* PROGRAM_ORDER_NON_INDEX_ARB_FOR_INDIVIDUAL_CUSTOMER */, "L" /* SHORT_EXEMPT_TRANSACTION_FOR_MEMBER_COMPETING_MARKET_MAKER_AFFILIATED_WITH_THE_FIRM_CLEARING_THE_TRADE */, "M" /* PROGRAM_ORDER_INDEX_ARB_FOR_OTHER_MEMBER */, "N" /* PROGRAM_ORDER_NON_INDEX_ARB_FOR_OTHER_MEMBER */, "O" /* COMPETING_DEALER_TRADES_O */, "P" /* PRINCIPAL */, "R" /* COMPETING_DEALER_TRADES_R */, "S" /* SPECIALIST_TRADES */, "T" /* COMPETING_DEALER_TRADES_T */, "U" /* PROGRAM_ORDER_INDEX_ARB_FOR_OTHER_AGENCY */, "W" /* ALL_OTHER_ORDERS_AS_AGENT_FOR_OTHER_MEMBER */, "X" /* SHORT_EXEMPT_TRANSACTION_FOR_MEMBER_COMPETING_MARKET_MAKER_NOT_AFFILIATED_WITH_THE_FIRM_CLEARING_THE_TRADE */, "Y" /* PROGRAM_ORDER_NON_INDEX_ARB_FOR_OTHER_AGENCY */, "Z" /* SHORT_EXEMPT_TRANSACTION_FOR_NON_MEMBER_COMPETING_MARKET_MAKER */), f0.Con{1}},
		LastShares32:              f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		LastPx31:                  f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		LastSpotRate194:           f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		LastForwardPoints195:      f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		LastMkt30:                 f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TradingSessionID336:       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		LastCapacity29:            f0.Fld{Opt, f0.ASCII, f0.String("1" /* AGENT */, "2" /* CROSS_AS_AGENT */, "3" /* CROSS_AS_PRINCIPAL */, "4" /* PRINCIPAL */), f0.Con{1}},
		LeavesQty151:              f0.Fld{Req, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		CumQty14:                  f0.Fld{Req, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		AvgPx6:                    f0.Fld{Req, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		DayOrderQty424:            f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		DayCumQty425:              f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		DayAvgPx426:               f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		GTBookingInst427:          f0.Fld{Opt, f0.ASCII, f0.IntDefault(0 /* BOOK_OUT_ALL_TRADES_ON_DAY_OF_EXECUTION */, 1 /* ACCUMULATE_EXECUTIONS_UNTIL_ORDER_IS_FILLED_OR_EXPIRES */, 2 /* ACCUMULATE_UNTIL_VERBALLY_NOTIFIED_OTHERWISE */), f0.Var{1, 19}},
		TradeDate75:               f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		TransactTime60:            f0.Fld{Opt, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		ReportToExch113:           f0.Fld{Opt, f0.ASCII, f0.BoolDefault(false /* NO */, true /* YES */), f0.Con{1}},
		Commission12:              f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		CommType13:                f0.Fld{Opt, f0.ASCII, f0.String("1" /* PER_SHARE */, "2" /* PERCENTAGE */, "3" /* ABSOLUTE */), f0.Con{1}},
		GrossTradeAmt381:          f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		SettlCurrAmt119:           f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		SettlCurrency120:          f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{3}},
		SettlCurrFxRate155:        f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		SettlCurrFxRateCalc156:    f0.Fld{Opt, f0.ASCII, f0.String("M" /* MULTIPLY */, "D" /* DIVIDE */), f0.Con{1}},
		HandlInst21:               f0.Fld{Opt, f0.ASCII, f0.String("1" /* AUTOMATED_EXECUTION_ORDER_PRIVATE_NO_BROKER_INTERVENTION */, "2" /* AUTOMATED_EXECUTION_ORDER_PUBLIC_BROKER_INTERVENTION_OK */, "3" /* MANUAL_ORDER_BEST_EXECUTION */), f0.Con{1}},
		MinQty110:                 f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		MaxFloor111:               f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		OpenClose77:               f0.Fld{Opt, f0.ASCII, f0.String("C" /* CLOSE */, "O" /* OPEN */), f0.Con{1}},
		MaxShow210:                f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		Text58:                    f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncodedTextLen354:         f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncodedText355:            f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		FutSettDate2193:           f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		OrderQty2192:              f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		ClearingFirm439:           f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		ClearingAccount440:        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		MultiLegReportingType442:  f0.Fld{Opt, f0.ASCII, f0.String("1" /* SINGLE_SECURITY */, "2" /* INDIVIDUAL_LEG_OF_A_MULTI_LEG_SECURITY */, "3" /* MULTI_LEG_SECURITY */), f0.Con{1}},
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
		OrderID37,                 // STRING
		SecondaryOrderID198,       // STRING
		ClOrdID11,                 // STRING
		OrigClOrdID41,             // STRING
		ClientID109,               // STRING
		ExecBroker76,              // STRING
		ListID66,                  // STRING
		ExecID17,                  // STRING
		ExecTransType20,           // CHAR
		ExecRefID19,               // STRING
		ExecType150,               // CHAR
		OrdStatus39,               // CHAR
		OrdRejReason103,           // INT
		ExecRestatementReason378,  // INT
		Account1,                  // STRING
		SettlmntTyp63,             // CHAR
		FutSettDate64,             // LOCALMKTDATE
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
		OrderQty38,                // QTY
		CashOrderQty152,           // QTY
		OrdType40,                 // CHAR
		Price44,                   // PRICE
		StopPx99,                  // PRICE
		PegDifference211,          // PRICEOFFSET
		DiscretionInst388,         // CHAR
		DiscretionOffset389,       // PRICEOFFSET
		Currency15,                // CURRENCY
		ComplianceID376,           // STRING
		SolicitedFlag377,          // BOOLEAN
		TimeInForce59,             // CHAR
		EffectiveTime168,          // UTCTIMESTAMP
		ExpireDate432,             // LOCALMKTDATE
		ExpireTime126,             // UTCTIMESTAMP
		ExecInst18,                // MULTIPLEVALUESTRING
		Rule80A47,                 // CHAR
		LastShares32,              // QTY
		LastPx31,                  // PRICE
		LastSpotRate194,           // PRICE
		LastForwardPoints195,      // PRICEOFFSET
		LastMkt30,                 // EXCHANGE
		TradingSessionID336,       // STRING
		LastCapacity29,            // CHAR
		LeavesQty151,              // QTY
		CumQty14,                  // QTY
		AvgPx6,                    // PRICE
		DayOrderQty424,            // QTY
		DayCumQty425,              // QTY
		DayAvgPx426,               // PRICE
		GTBookingInst427,          // INT
		TradeDate75,               // LOCALMKTDATE
		TransactTime60,            // UTCTIMESTAMP
		ReportToExch113,           // BOOLEAN
		Commission12,              // AMT
		CommType13,                // CHAR
		GrossTradeAmt381,          // AMT
		SettlCurrAmt119,           // AMT
		SettlCurrency120,          // CURRENCY
		SettlCurrFxRate155,        // FLOAT
		SettlCurrFxRateCalc156,    // CHAR
		HandlInst21,               // CHAR
		MinQty110,                 // QTY
		MaxFloor111,               // QTY
		OpenClose77,               // CHAR
		MaxShow210,                // QTY
		Text58,                    // STRING
		EncodedTextLen354,         // LENGTH
		EncodedText355,            // DATA
		FutSettDate2193,           // LOCALMKTDATE
		OrderQty2192,              // QTY
		ClearingFirm439,           // STRING
		ClearingAccount440,        // STRING
		MultiLegReportingType442,  // CHAR
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
	OrderID37                 = 37  // STRING
	SecondaryOrderID198       = 198 // STRING
	ClOrdID11                 = 11  // STRING
	OrigClOrdID41             = 41  // STRING
	ClientID109               = 109 // STRING
	ExecBroker76              = 76  // STRING
	ListID66                  = 66  // STRING
	ExecID17                  = 17  // STRING
	ExecTransType20           = 20  // CHAR
	ExecRefID19               = 19  // STRING
	ExecType150               = 150 // CHAR
	OrdStatus39               = 39  // CHAR
	OrdRejReason103           = 103 // INT
	ExecRestatementReason378  = 378 // INT
	Account1                  = 1   // STRING
	SettlmntTyp63             = 63  // CHAR
	FutSettDate64             = 64  // LOCALMKTDATE
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
	OrderQty38                = 38  // QTY
	CashOrderQty152           = 152 // QTY
	OrdType40                 = 40  // CHAR
	Price44                   = 44  // PRICE
	StopPx99                  = 99  // PRICE
	PegDifference211          = 211 // PRICEOFFSET
	DiscretionInst388         = 388 // CHAR
	DiscretionOffset389       = 389 // PRICEOFFSET
	Currency15                = 15  // CURRENCY
	ComplianceID376           = 376 // STRING
	SolicitedFlag377          = 377 // BOOLEAN
	TimeInForce59             = 59  // CHAR
	EffectiveTime168          = 168 // UTCTIMESTAMP
	ExpireDate432             = 432 // LOCALMKTDATE
	ExpireTime126             = 126 // UTCTIMESTAMP
	ExecInst18                = 18  // MULTIPLEVALUESTRING
	Rule80A47                 = 47  // CHAR
	LastShares32              = 32  // QTY
	LastPx31                  = 31  // PRICE
	LastSpotRate194           = 194 // PRICE
	LastForwardPoints195      = 195 // PRICEOFFSET
	LastMkt30                 = 30  // EXCHANGE
	TradingSessionID336       = 336 // STRING
	LastCapacity29            = 29  // CHAR
	LeavesQty151              = 151 // QTY
	CumQty14                  = 14  // QTY
	AvgPx6                    = 6   // PRICE
	DayOrderQty424            = 424 // QTY
	DayCumQty425              = 425 // QTY
	DayAvgPx426               = 426 // PRICE
	GTBookingInst427          = 427 // INT
	TradeDate75               = 75  // LOCALMKTDATE
	TransactTime60            = 60  // UTCTIMESTAMP
	ReportToExch113           = 113 // BOOLEAN
	Commission12              = 12  // AMT
	CommType13                = 13  // CHAR
	GrossTradeAmt381          = 381 // AMT
	SettlCurrAmt119           = 119 // AMT
	SettlCurrency120          = 120 // CURRENCY
	SettlCurrFxRate155        = 155 // FLOAT
	SettlCurrFxRateCalc156    = 156 // CHAR
	HandlInst21               = 21  // CHAR
	MinQty110                 = 110 // QTY
	MaxFloor111               = 111 // QTY
	OpenClose77               = 77  // CHAR
	MaxShow210                = 210 // QTY
	Text58                    = 58  // STRING
	EncodedTextLen354         = 354 // LENGTH
	EncodedText355            = 355 // DATA
	FutSettDate2193           = 193 // LOCALMKTDATE
	OrderQty2192              = 192 // QTY
	ClearingFirm439           = 439 // STRING
	ClearingAccount440        = 440 // STRING
	MultiLegReportingType442  = 442 // CHAR
	SignatureLength93         = 93  // LENGTH
	Signature89               = 89  // DATA
	CheckSum10                = 10  // STRING
)
