// Code generated by protofix; DO NOT EDIT.
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fix50ordermasscancelrequest is a format of the FIX.5.0 OrderMassCancelRequest message.
package fix50ordermasscancelrequest

import (
	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/marshfix"
)

var (
	FIX50OrderMassCancelRequestMarshaler   = marshfix.Marshal{Tag: "FIX50", Format: FIX50OrderMassCancelRequest}
	FIX50OrderMassCancelRequestUnmarshaler = marshfix.Unmarshal{Tag: "FIX50", Format: FIX50OrderMassCancelRequest}
)

// FIX50OrderMassCancelRequest is a FIX.5.0 format of the OrderMassCancelRequest message which maps the codecs into individual fields.
var FIX50OrderMassCancelRequest = f0.Format{
	Fields: map[int]f0.Codec{
		ClOrdID11:                               f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SecondaryClOrdID526:                     f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		MassCancelRequestType530:                f0.Fld{Req, f0.ASCII, f0.String("1" /* CANCEL_ORDERS_FOR_A_SECURITY */, "2" /* CANCEL_ORDERS_FOR_AN_UNDERLYING_SECURITY */, "3" /* CANCEL_ORDERS_FOR_A_PRODUCT */, "4" /* CANCEL_ORDERS_FOR_A_CFICODE */, "5" /* CANCEL_ORDERS_FOR_A_SECURITYTYPE */, "6" /* CANCEL_ORDERS_FOR_A_TRADING_SESSION */, "7" /* CANCEL_ALL_ORDERS */), f0.Con{1}},
		TradingSessionID336:                     f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TradingSessionSubID625:                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		Side54:                                  f0.Fld{Opt, f0.ASCII, f0.String("1" /* BUY */, "2" /* SELL */, "3" /* BUY_MINUS */, "4" /* SELL_PLUS */, "5" /* SELL_SHORT */, "6" /* SELL_SHORT_EXEMPT */, "7" /* UNDISCLOSED */, "8" /* CROSS */, "9" /* CROSS_SHORT */, "A" /* CROSS_SHORT_EXXMPT */, "B" /* AS_DEFINED */, "C" /* OPPOSITE */, "D" /* SUBSCRIBE */, "E" /* REDEEM */, "F" /* LEND */, "G" /* BORROW */), f0.Con{1}},
		TransactTime60:                          f0.Fld{Req, f0.ASCII, f0.UTCTimestampMillisecondTime(), f0.Var{1, 27}},
		Text58:                                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncodedTextLen354:                       f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncodedText355:                          f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		Symbol55:                                f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SymbolSfx65:                             f0.Fld{Opt, f0.ASCII, f0.String("CD" /* EUCP_WITH_LUMP_SUM_INTEREST_RATHER_THAN_DISCOUNT_PRICE */, "WI" /* WHEN_ISSUED_FOR_A_SECURITY_TO_BE_REISSUED_UNDER_AN_OLD_CUSIP_OR_ISIN */), f0.Var{1, 65536}},
		SecurityID48:                            f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SecurityIDSource22:                      f0.Fld{Opt, f0.ASCII, f0.String("1" /* CUSIP */, "2" /* SEDOL */, "3" /* QUIK */, "4" /* ISIN_NUMBER */, "5" /* RIC_CODE */, "6" /* ISO_CURRENCY_CODE */, "7" /* ISO_COUNTRY_CODE */, "8" /* EXCHANGE_SYMBOL */, "9" /* CONSOLIDATED_TAPE_ASSOCIATION */, "A" /* BLOOMBERG_SYMBOL */, "B" /* WERTPAPIER */, "C" /* DUTCH */, "D" /* VALOREN */, "E" /* SICOVAM */, "F" /* BELGIAN */, "G" /* COMMON */, "H" /* CLEARING_HOUSE */, "I" /* ISDA_FPML_PRODUCT_SPECIFICATION */, "J" /* OPTION_PRICE_REPORTING_AUTHORITY */, "L" /* LETTER_OF_CREDIT */, "K" /* ISDA_FPML_PRODUCT_URL */), f0.Var{1, 65536}},
		Product460:                              f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* AGENCY */, 10 /* MORTGAGE */, 11 /* MUNICIPAL */, 12 /* OTHER */, 13 /* FINANCING */, 2 /* COMMODITY */, 3 /* CORPORATE */, 4 /* CURRENCY */, 5 /* EQUITY */, 6 /* GOVERNMENT */, 7 /* INDEX */, 8 /* LOAN */, 9 /* MONEYMARKET */), f0.Var{1, 19}},
		CFICode461:                              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SecurityType167:                         f0.Fld{Opt, f0.ASCII, f0.String("ABS" /* ASSET_BACKED_SECURITIES */, "AMENDED" /* AMENDED_RESTATED */, "AN" /* OTHER_ANTICIPATION_NOTES */, "BA" /* BANKERS_ACCEPTANCE */, "BN" /* BANK_NOTES */, "BOX" /* BILL_OF_EXCHANGES */, "BRADY" /* BRADY_BOND */, "BRIDGE" /* BRIDGE_LOAN */, "BUYSELL" /* BUY_SELLBACK */, "CB" /* CONVERTIBLE_BOND */, "CD" /* CERTIFICATE_OF_DEPOSIT */, "CL" /* CALL_LOANS */, "CMBS" /* CORP_MORTGAGE_BACKED_SECURITIES */, "CMO" /* COLLATERALIZED_MORTGAGE_OBLIGATION */, "COFO" /* CERTIFICATE_OF_OBLIGATION */, "COFP" /* CERTIFICATE_OF_PARTICIPATION */, "CORP" /* CORPORATE_BOND */, "CP" /* COMMERCIAL_PAPER */, "CPP" /* CORPORATE_PRIVATE_PLACEMENT */, "CS" /* COMMON_STOCK */, "DEFLTED" /* DEFAULTED */, "DINP" /* DEBTOR_IN_POSSESSION */, "DN" /* DEPOSIT_NOTES */, "DUAL" /* DUAL_CURRENCY */, "EUCD" /* EURO_CERTIFICATE_OF_DEPOSIT */, "EUCORP" /* EURO_CORPORATE_BOND */, "EUCP" /* EURO_COMMERCIAL_PAPER */, "EUSOV" /* EURO_SOVEREIGNS */, "EUSUPRA" /* EURO_SUPRANATIONAL_COUPONS */, "FAC" /* FEDERAL_AGENCY_COUPON */, "FADN" /* FEDERAL_AGENCY_DISCOUNT_NOTE */, "FOR" /* FOREIGN_EXCHANGE_CONTRACT */, "FORWARD" /* FORWARD */, "FUT" /* FUTURE */, "GO" /* GENERAL_OBLIGATION_BONDS */, "IET" /* IOETTE_MORTGAGE */, "LOFC" /* LETTER_OF_CREDIT */, "LQN" /* LIQUIDITY_NOTE */, "MATURED" /* MATURED */, "MBS" /* MORTGAGE_BACKED_SECURITIES */, "MF" /* MUTUAL_FUND */, "MIO" /* MORTGAGE_INTEREST_ONLY */, "MLEG" /* MULTILEG_INSTRUMENT */, "MPO" /* MORTGAGE_PRINCIPAL_ONLY */, "MPP" /* MORTGAGE_PRIVATE_PLACEMENT */, "MPT" /* MISCELLANEOUS_PASS_THROUGH */, "MT" /* MANDATORY_TENDER */, "MTN" /* MEDIUM_TERM_NOTES */, "NONE" /* NO_SECURITY_TYPE */, "ONITE" /* OVERNIGHT */, "OPT" /* OPTION */, "PEF" /* PRIVATE_EXPORT_FUNDING */, "PFAND" /* PFANDBRIEFE */, "PN" /* PROMISSORY_NOTE */, "PS" /* PREFERRED_STOCK */, "PZFJ" /* PLAZOS_FIJOS */, "RAN" /* REVENUE_ANTICIPATION_NOTE */, "REPLACD" /* REPLACED */, "REPO" /* REPURCHASE */, "RETIRED" /* RETIRED */, "REV" /* REVENUE_BONDS */, "RVLV" /* REVOLVER_LOAN */, "RVLVTRM" /* REVOLVER_TERM_LOAN */, "SECLOAN" /* SECURITIES_LOAN */, "SECPLEDGE" /* SECURITIES_PLEDGE */, "SPCLA" /* SPECIAL_ASSESSMENT */, "SPCLO" /* SPECIAL_OBLIGATION */, "SPCLT" /* SPECIAL_TAX */, "STN" /* SHORT_TERM_LOAN_NOTE */, "STRUCT" /* STRUCTURED_NOTES */, "SUPRA" /* USD_SUPRANATIONAL_COUPONS */, "SWING" /* SWING_LINE_FACILITY */, "TAN" /* TAX_ANTICIPATION_NOTE */, "TAXA" /* TAX_ALLOCATION */, "TBA" /* TO_BE_ANNOUNCED */, "TBILL" /* US_TREASURY_BILL_TBILL */, "TBOND" /* US_TREASURY_BOND */, "TCAL" /* PRINCIPAL_STRIP_OF_A_CALLABLE_BOND_OR_NOTE */, "TD" /* TIME_DEPOSIT */, "TECP" /* TAX_EXEMPT_COMMERCIAL_PAPER */, "TERM" /* TERM_LOAN */, "TINT" /* INTEREST_STRIP_FROM_ANY_BOND_OR_NOTE */, "TIPS" /* TREASURY_INFLATION_PROTECTED_SECURITIES */, "TNOTE" /* US_TREASURY_NOTE_TNOTE */, "TPRN" /* PRINCIPAL_STRIP_FROM_A_NON_CALLABLE_BOND_OR_NOTE */, "TRAN" /* TAX_REVENUE_ANTICIPATION_NOTE */, "UST" /* US_TREASURY_NOTE_UST */, "USTB" /* US_TREASURY_BILL_USTB */, "VRDN" /* VARIABLE_RATE_DEMAND_NOTE */, "WAR" /* WARRANT */, "WITHDRN" /* WITHDRAWN */, "WLD" /* WILDCARD_ENTRY */, "XCN" /* EXTENDED_COMM_NOTE */, "XLINKD" /* INDEXED_LINKED */, "YANK" /* YANKEE_CORPORATE_BOND */, "YCD" /* YANKEE_CERTIFICATE_OF_DEPOSIT */, "OOP" /* OPTIONS_ON_PHYSICAL */, "OOF" /* OPTIONS_ON_FUTURES */, "CASH" /* CASH */), f0.Var{1, 65536}},
		SecuritySubType762:                      f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		MaturityMonthYear200:                    f0.Fld{Opt, f0.ASCII, f0.MonthYearTime(), f0.Var{6, 8}},
		MaturityDate541:                         f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		CouponPaymentDate224:                    f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		IssueDate225:                            f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		RepoCollateralSecurityType239:           f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		RepurchaseTerm226:                       f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		RepurchaseRate227:                       f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		Factor228:                               f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		CreditRating255:                         f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		InstrRegistry543:                        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		CountryOfIssue470:                       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{2}},
		StateOrProvinceOfIssue471:               f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		LocaleOfIssue472:                        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RedemptionDate240:                       f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		StrikePrice202:                          f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		StrikeCurrency947:                       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{3}},
		OptAttribute206:                         f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		ContractMultiplier231:                   f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		CouponRate223:                           f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		SecurityExchange207:                     f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		Issuer106:                               f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncodedIssuerLen348:                     f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncodedIssuer349:                        f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		SecurityDesc107:                         f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncodedSecurityDescLen350:               f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncodedSecurityDesc351:                  f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		Pool691:                                 f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		ContractSettlMonth667:                   f0.Fld{Opt, f0.ASCII, f0.MonthYearTime(), f0.Var{6, 8}},
		CPProgram875:                            f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* 3 */, 2 /* 4 */, 99 /* OTHER */), f0.Var{1, 19}},
		CPRegType876:                            f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		DatedDate873:                            f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		InterestAccrualDate874:                  f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		SecurityStatus965:                       f0.Fld{Opt, f0.ASCII, f0.String("1" /* ACTIVE */, "2" /* INACTIVE */), f0.Var{1, 65536}},
		SettleOnOpenFlag966:                     f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		InstrmtAssignmentMethod1049:             f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		StrikeMultiplier967:                     f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		StrikeValue968:                          f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		MinPriceIncrement969:                    f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		PositionLimit970:                        f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		NTPositionLimit971:                      f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		UnitOfMeasure996:                        f0.Fld{Opt, f0.ASCII, f0.String("MWh" /* MEGAWATT_HOURS */, "MMBtu" /* ONE_MILLION_BTU */, "Bbl" /* BARRELS */, "Gal" /* GALLONS */, "t" /* METRIC_TONS */, "tn" /* TONS */, "MMbbl" /* MILLION_BARRELS */, "lbs" /* POUNDS */, "oz_tr" /* TROY_OUNCES */, "USD" /* US_DOLLARS */, "Bcf" /* BILLION_CUBIC_FEET */, "Bu" /* BUSHELS */), f0.Var{1, 65536}},
		TimeUnit997:                             f0.Fld{Opt, f0.ASCII, f0.String("S" /* SECOND */, "Min" /* MINUTE */, "H" /* HOUR */, "D" /* DAY */, "Wk" /* WEEK */, "Mo" /* MONTH */, "Yr" /* YEAR */), f0.Var{1, 65536}},
		MaturityTime1079:                        f0.Fld{Opt, f0.ASCII, f0.TZTime(), f0.Var{8, 12}},
		SecurityAltID455:                        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		SecurityAltIDSource456:                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EventType865:                            f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* PUT */, 2 /* CALL */, 3 /* TENDER */, 4 /* SINKING_FUND_CALL */, 99 /* OTHER */, 5 /* ACTIVATION */, 6 /* INACTIVIATION */), f0.Var{1, 19}},
		EventDate866:                            f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		EventPx867:                              f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		EventText868:                            f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		InstrumentPartyID1019:                   f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		InstrumentPartyIDSource1050:             f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		InstrumentPartyRole1051:                 f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		InstrumentPartySubID1053:                f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		InstrumentPartySubIDType1054:            f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		UnderlyingSymbol311:                     f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingSymbolSfx312:                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingSecurityID309:                 f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingSecurityIDSource305:           f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingProduct462:                    f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		UnderlyingCFICode463:                    f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingSecurityType310:               f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingSecuritySubType763:            f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingMaturityMonthYear313:          f0.Fld{Opt, f0.ASCII, f0.MonthYearTime(), f0.Var{6, 8}},
		UnderlyingMaturityDate542:               f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		UnderlyingCouponPaymentDate241:          f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		UnderlyingIssueDate242:                  f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		UnderlyingRepoCollateralSecurityType243: f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		UnderlyingRepurchaseTerm244:             f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		UnderlyingRepurchaseRate245:             f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		UnderlyingFactor246:                     f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		UnderlyingCreditRating256:               f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingInstrRegistry595:              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingCountryOfIssue592:             f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{2}},
		UnderlyingStateOrProvinceOfIssue593:     f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingLocaleOfIssue594:              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingRedemptionDate247:             f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		UnderlyingStrikePrice316:                f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		UnderlyingStrikeCurrency941:             f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{3}},
		UnderlyingOptAttribute317:               f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		UnderlyingContractMultiplier436:         f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		UnderlyingCouponRate435:                 f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		UnderlyingSecurityExchange308:           f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingIssuer306:                     f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncodedUnderlyingIssuerLen362:           f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncodedUnderlyingIssuer363:              f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		UnderlyingSecurityDesc307:               f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncodedUnderlyingSecurityDescLen364:     f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncodedUnderlyingSecurityDesc365:        f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		UnderlyingCPProgram877:                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingCPRegType878:                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingCurrency318:                   f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{3}},
		UnderlyingQty879:                        f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		UnderlyingPx810:                         f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		UnderlyingDirtyPrice882:                 f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		UnderlyingEndPrice883:                   f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		UnderlyingStartValue884:                 f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		UnderlyingCurrentValue885:               f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		UnderlyingEndValue886:                   f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		UnderlyingAllocationPercent972:          f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		UnderlyingSettlementType975:             f0.Fld{Opt, f0.ASCII, f0.IntDefault(2 /* T_PLUS_1 */, 4 /* T_PLUS_3 */, 5 /* T_PLUS_4 */), f0.Var{1, 19}},
		UnderlyingCashAmount973:                 f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		UnderlyingCashType974:                   f0.Fld{Opt, f0.ASCII, f0.String("FIXED" /* FIXED */, "DIFF" /* DIFF */), f0.Var{1, 65536}},
		UnderlyingUnitOfMeasure998:              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingTimeUnit1000:                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingCapValue1038:                  f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		UnderlyingSettlMethod1039:               f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingAdjustedQuantity1044:          f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		UnderlyingFXRate1045:                    f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		UnderlyingFXRateCalc1046:                f0.Fld{Opt, f0.ASCII, f0.String("M" /* MULTIPLY */, "D" /* DIVIDE */), f0.Con{1}},
		UnderlyingSecurityAltID458:              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingSecurityAltIDSource459:        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingStipType888:                   f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UnderlyingStipValue889:                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UndlyInstrumentPartyID1059:              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UndlyInstrumentPartyIDSource1060:        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		UndlyInstrumentPartyRole1061:            f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		UndlyInstrumentPartySubID1063:           f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		UndlyInstrumentPartySubIDType1064:       f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		PartyID448:                              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		PartyIDSource447:                        f0.Fld{Opt, f0.ASCII, f0.String("1" /* KOREAN_INVESTOR_ID */, "2" /* TAIWANESE_QUALIFIED_FOREIGN_INVESTOR_ID_QFII_FID */, "3" /* TAIWANESE_TRADING_ACCT */, "4" /* MALAYSIAN_CENTRAL_DEPOSITORY */, "5" /* CHINESE_INVESTOR_ID */, "6" /* UK_NATIONAL_INSURANCE_OR_PENSION_NUMBER */, "7" /* US_SOCIAL_SECURITY_NUMBER */, "8" /* US_EMPLOYER_OR_TAX_ID_NUMBER */, "9" /* AUSTRALIAN_BUSINESS_NUMBER */, "A" /* AUSTRALIAN_TAX_FILE_NUMBER */, "B" /* BIC */, "C" /* GENERALLY_ACCEPTED_MARKET_PARTICIPANT_IDENTIFIER */, "D" /* PROPRIETARY */, "E" /* ISO_COUNTRY_CODE */, "F" /* SETTLEMENT_ENTITY_LOCATION */, "G" /* MIC */, "H" /* CSD_PARTICIPANT_MEMBER_CODE */, "I" /* DIRECTED_BROKER_THREE_CHARACTER_ACRONYM_AS_DEFINED_IN_ISITC_ETC_BEST_PRACTICE_GUIDELINES_DOCUMENT */), f0.Con{1}},
		PartyRole452:                            f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* EXECUTING_FIRM */, 10 /* SETTLEMENT_LOCATION */, 11 /* ORDER_ORIGINATION_TRADER */, 12 /* EXECUTING_TRADER */, 13 /* ORDER_ORIGINATION_FIRM */, 14 /* GIVEUP_CLEARING_FIRM */, 15 /* CORRESPONDANT_CLEARING_FIRM */, 16 /* EXECUTING_SYSTEM */, 17 /* CONTRA_FIRM */, 18 /* CONTRA_CLEARING_FIRM */, 19 /* SPONSORING_FIRM */, 2 /* BROKER_OF_CREDIT */, 20 /* UNDERLYING_CONTRA_FIRM */, 21 /* CLEARING_ORGANIZATION */, 22 /* EXCHANGE */, 24 /* CUSTOMER_ACCOUNT */, 25 /* CORRESPONDENT_CLEARING_ORGANIZATION */, 26 /* CORRESPONDENT_BROKER */, 27 /* BUYER_SELLER */, 28 /* CUSTODIAN */, 29 /* INTERMEDIARY */, 3 /* CLIENT_ID */, 30 /* AGENT */, 31 /* SUB_CUSTODIAN */, 32 /* BENEFICIARY */, 33 /* INTERESTED_PARTY */, 34 /* REGULATORY_BODY */, 35 /* LIQUIDITY_PROVIDER */, 36 /* ENTERING_TRADER */, 37 /* CONTRA_TRADER */, 38 /* POSITION_ACCOUNT */, 4 /* CLEARING_FIRM */, 5 /* INVESTOR_ID */, 6 /* INTRODUCING_FIRM */, 7 /* ENTERING_FIRM */, 8 /* LOCATE */, 9 /* FUND_MANAGER_CLIENT_ID */, 60 /* INTRODUCING_BROKER */, 41 /* CONTRA_POSITION_ACCOUNT */, 42 /* CONTRA_EXCHANGE */, 43 /* INTERNAL_CARRY_ACCOUNT */, 44 /* ORDER_ENTRY_OPERATOR_ID */, 45 /* SECONDARY_ACCOUNT_NUMBER */, 46 /* FORIEGN_FIRM */, 47 /* THIRD_PARTY_ALLOCATION_FIRM */, 48 /* CLAIMING_ACCOUNT */, 49 /* ASSET_MANAGER */, 50 /* PLEDGOR_ACCOUNT */, 51 /* PLEDGEE_ACCOUNT */, 52 /* LARGE_TRADER_REPORTABLE_ACCOUNT */, 53 /* TRADER_MNEMONIC */, 54 /* SENDER_LOCATION */, 55 /* SESSION_ID */, 56 /* ACCEPTABLE_COUNTERPARTY */, 57 /* UNACCEPTABLE_COUNTERPARTY */, 58 /* ENTERING_UNIT */, 59 /* EXECUTING_UNIT */, 39 /* CONTRA_INVESTOR_ID */, 40 /* TRANSFER_TO_FIRM */, 61 /* QUOTE_ORIGINATOR */, 62 /* REPORT_ORIGINATOR */, 63 /* SYSTEMATIC_INTERNALISER */, 64 /* MULTILATERAL_TRADING_FACILITY */, 65 /* REGULATED_MARKET */, 66 /* MARKET_MAKER */, 67 /* INVESTMENT_FIRM */, 68 /* HOST_COMPETENT_AUTHORITY */, 69 /* HOME_COMPETENT_AUTHORITY */, 70 /* COMPETENT_AUTHORITY_OF_THE_MOST_RELEVANT_MARKET_IN_TERMS_OF_LIQUIDITY */, 71 /* COMPETENT_AUTHORITY_OF_THE_TRANSACTION */, 72 /* REPORTING_INTERMEDIARY */, 73 /* EXECUTION_VENUE */, 74 /* MARKET_DATA_ENTRY_ORIGINATOR */, 75 /* LOCATION_ID */, 76 /* DESK_ID */, 77 /* MARKET_DATA_MARKET */, 78 /* ALLOCATION_ENTITY */), f0.Var{1, 19}},
		PartySubID523:                           f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		PartySubIDType803:                       f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* FIRM */, 10 /* SECURITIES_ACCOUNT_NUMBER */, 11 /* REGISTRATION_NUMBER */, 12 /* REGISTERED_ADDRESS_12 */, 13 /* REGULATORY_STATUS */, 14 /* REGISTRATION_NAME */, 15 /* CASH_ACCOUNT_NUMBER */, 16 /* BIC */, 17 /* CSD_PARTICIPANT_MEMBER_CODE */, 18 /* REGISTERED_ADDRESS_18 */, 19 /* FUND_ACCOUNT_NAME */, 2 /* PERSON */, 20 /* TELEX_NUMBER */, 21 /* FAX_NUMBER */, 22 /* SECURITIES_ACCOUNT_NAME */, 23 /* CASH_ACCOUNT_NAME */, 24 /* DEPARTMENT */, 25 /* LOCATION_DESK */, 26 /* POSITION_ACCOUNT_TYPE */, 3 /* SYSTEM */, 4 /* APPLICATION */, 5 /* FULL_LEGAL_NAME_OF_FIRM */, 6 /* POSTAL_ADDRESS */, 7 /* PHONE_NUMBER */, 8 /* EMAIL_ADDRESS */, 9 /* CONTACT_NAME */, 27 /* SECURITY_LOCATE_ID */, 28 /* MARKET_MAKER */, 29 /* ELIGIBLE_COUNTERPARTY */, 30 /* PROFESSIONAL_CLIENT */, 31 /* LOCATION */, 32 /* EXECUTION_VENUE */), f0.Var{1, 19}},
	},
	Unknown0: f0.UnknownFld{f0.ASCII, f0.Unknown(), f0.Var{1, 65536}},
	Sort: []int{
		ClOrdID11,                               // STRING
		SecondaryClOrdID526,                     // STRING
		MassCancelRequestType530,                // CHAR
		TradingSessionID336,                     // STRING
		TradingSessionSubID625,                  // STRING
		Side54,                                  // CHAR
		TransactTime60,                          // UTCTIMESTAMP
		Text58,                                  // STRING
		EncodedTextLen354,                       // LENGTH
		EncodedText355,                          // DATA
		Symbol55,                                // STRING
		SymbolSfx65,                             // STRING
		SecurityID48,                            // STRING
		SecurityIDSource22,                      // STRING
		Product460,                              // INT
		CFICode461,                              // STRING
		SecurityType167,                         // STRING
		SecuritySubType762,                      // STRING
		MaturityMonthYear200,                    // MONTHYEAR
		MaturityDate541,                         // LOCALMKTDATE
		CouponPaymentDate224,                    // LOCALMKTDATE
		IssueDate225,                            // LOCALMKTDATE
		RepoCollateralSecurityType239,           // INT
		RepurchaseTerm226,                       // INT
		RepurchaseRate227,                       // PERCENTAGE
		Factor228,                               // FLOAT
		CreditRating255,                         // STRING
		InstrRegistry543,                        // STRING
		CountryOfIssue470,                       // COUNTRY
		StateOrProvinceOfIssue471,               // STRING
		LocaleOfIssue472,                        // STRING
		RedemptionDate240,                       // LOCALMKTDATE
		StrikePrice202,                          // PRICE
		StrikeCurrency947,                       // CURRENCY
		OptAttribute206,                         // CHAR
		ContractMultiplier231,                   // FLOAT
		CouponRate223,                           // PERCENTAGE
		SecurityExchange207,                     // EXCHANGE
		Issuer106,                               // STRING
		EncodedIssuerLen348,                     // LENGTH
		EncodedIssuer349,                        // DATA
		SecurityDesc107,                         // STRING
		EncodedSecurityDescLen350,               // LENGTH
		EncodedSecurityDesc351,                  // DATA
		Pool691,                                 // STRING
		ContractSettlMonth667,                   // MONTHYEAR
		CPProgram875,                            // INT
		CPRegType876,                            // STRING
		DatedDate873,                            // LOCALMKTDATE
		InterestAccrualDate874,                  // LOCALMKTDATE
		SecurityStatus965,                       // STRING
		SettleOnOpenFlag966,                     // STRING
		InstrmtAssignmentMethod1049,             // CHAR
		StrikeMultiplier967,                     // FLOAT
		StrikeValue968,                          // FLOAT
		MinPriceIncrement969,                    // FLOAT
		PositionLimit970,                        // INT
		NTPositionLimit971,                      // INT
		UnitOfMeasure996,                        // STRING
		TimeUnit997,                             // STRING
		MaturityTime1079,                        // TZTIMEONLY
		SecurityAltID455,                        // STRING
		SecurityAltIDSource456,                  // STRING
		EventType865,                            // INT
		EventDate866,                            // LOCALMKTDATE
		EventPx867,                              // PRICE
		EventText868,                            // STRING
		InstrumentPartyID1019,                   // STRING
		InstrumentPartyIDSource1050,             // CHAR
		InstrumentPartyRole1051,                 // INT
		InstrumentPartySubID1053,                // STRING
		InstrumentPartySubIDType1054,            // INT
		UnderlyingSymbol311,                     // STRING
		UnderlyingSymbolSfx312,                  // STRING
		UnderlyingSecurityID309,                 // STRING
		UnderlyingSecurityIDSource305,           // STRING
		UnderlyingProduct462,                    // INT
		UnderlyingCFICode463,                    // STRING
		UnderlyingSecurityType310,               // STRING
		UnderlyingSecuritySubType763,            // STRING
		UnderlyingMaturityMonthYear313,          // MONTHYEAR
		UnderlyingMaturityDate542,               // LOCALMKTDATE
		UnderlyingCouponPaymentDate241,          // LOCALMKTDATE
		UnderlyingIssueDate242,                  // LOCALMKTDATE
		UnderlyingRepoCollateralSecurityType243, // INT
		UnderlyingRepurchaseTerm244,             // INT
		UnderlyingRepurchaseRate245,             // PERCENTAGE
		UnderlyingFactor246,                     // FLOAT
		UnderlyingCreditRating256,               // STRING
		UnderlyingInstrRegistry595,              // STRING
		UnderlyingCountryOfIssue592,             // COUNTRY
		UnderlyingStateOrProvinceOfIssue593,     // STRING
		UnderlyingLocaleOfIssue594,              // STRING
		UnderlyingRedemptionDate247,             // LOCALMKTDATE
		UnderlyingStrikePrice316,                // PRICE
		UnderlyingStrikeCurrency941,             // CURRENCY
		UnderlyingOptAttribute317,               // CHAR
		UnderlyingContractMultiplier436,         // FLOAT
		UnderlyingCouponRate435,                 // PERCENTAGE
		UnderlyingSecurityExchange308,           // EXCHANGE
		UnderlyingIssuer306,                     // STRING
		EncodedUnderlyingIssuerLen362,           // LENGTH
		EncodedUnderlyingIssuer363,              // DATA
		UnderlyingSecurityDesc307,               // STRING
		EncodedUnderlyingSecurityDescLen364,     // LENGTH
		EncodedUnderlyingSecurityDesc365,        // DATA
		UnderlyingCPProgram877,                  // STRING
		UnderlyingCPRegType878,                  // STRING
		UnderlyingCurrency318,                   // CURRENCY
		UnderlyingQty879,                        // QTY
		UnderlyingPx810,                         // PRICE
		UnderlyingDirtyPrice882,                 // PRICE
		UnderlyingEndPrice883,                   // PRICE
		UnderlyingStartValue884,                 // AMT
		UnderlyingCurrentValue885,               // AMT
		UnderlyingEndValue886,                   // AMT
		UnderlyingAllocationPercent972,          // PERCENTAGE
		UnderlyingSettlementType975,             // INT
		UnderlyingCashAmount973,                 // AMT
		UnderlyingCashType974,                   // STRING
		UnderlyingUnitOfMeasure998,              // STRING
		UnderlyingTimeUnit1000,                  // STRING
		UnderlyingCapValue1038,                  // AMT
		UnderlyingSettlMethod1039,               // STRING
		UnderlyingAdjustedQuantity1044,          // QTY
		UnderlyingFXRate1045,                    // FLOAT
		UnderlyingFXRateCalc1046,                // CHAR
		UnderlyingSecurityAltID458,              // STRING
		UnderlyingSecurityAltIDSource459,        // STRING
		UnderlyingStipType888,                   // STRING
		UnderlyingStipValue889,                  // STRING
		UndlyInstrumentPartyID1059,              // STRING
		UndlyInstrumentPartyIDSource1060,        // CHAR
		UndlyInstrumentPartyRole1061,            // INT
		UndlyInstrumentPartySubID1063,           // STRING
		UndlyInstrumentPartySubIDType1064,       // INT
		PartyID448,                              // STRING
		PartyIDSource447,                        // CHAR
		PartyRole452,                            // INT
		PartySubID523,                           // STRING
		PartySubIDType803,                       // INT
	},
}

const Req, Opt = true, false

const (
	ClOrdID11                               = 11   // STRING
	SecondaryClOrdID526                     = 526  // STRING
	MassCancelRequestType530                = 530  // CHAR
	TradingSessionID336                     = 336  // STRING
	TradingSessionSubID625                  = 625  // STRING
	Side54                                  = 54   // CHAR
	TransactTime60                          = 60   // UTCTIMESTAMP
	Text58                                  = 58   // STRING
	EncodedTextLen354                       = 354  // LENGTH
	EncodedText355                          = 355  // DATA
	Symbol55                                = 55   // STRING
	SymbolSfx65                             = 65   // STRING
	SecurityID48                            = 48   // STRING
	SecurityIDSource22                      = 22   // STRING
	Product460                              = 460  // INT
	CFICode461                              = 461  // STRING
	SecurityType167                         = 167  // STRING
	SecuritySubType762                      = 762  // STRING
	MaturityMonthYear200                    = 200  // MONTHYEAR
	MaturityDate541                         = 541  // LOCALMKTDATE
	CouponPaymentDate224                    = 224  // LOCALMKTDATE
	IssueDate225                            = 225  // LOCALMKTDATE
	RepoCollateralSecurityType239           = 239  // INT
	RepurchaseTerm226                       = 226  // INT
	RepurchaseRate227                       = 227  // PERCENTAGE
	Factor228                               = 228  // FLOAT
	CreditRating255                         = 255  // STRING
	InstrRegistry543                        = 543  // STRING
	CountryOfIssue470                       = 470  // COUNTRY
	StateOrProvinceOfIssue471               = 471  // STRING
	LocaleOfIssue472                        = 472  // STRING
	RedemptionDate240                       = 240  // LOCALMKTDATE
	StrikePrice202                          = 202  // PRICE
	StrikeCurrency947                       = 947  // CURRENCY
	OptAttribute206                         = 206  // CHAR
	ContractMultiplier231                   = 231  // FLOAT
	CouponRate223                           = 223  // PERCENTAGE
	SecurityExchange207                     = 207  // EXCHANGE
	Issuer106                               = 106  // STRING
	EncodedIssuerLen348                     = 348  // LENGTH
	EncodedIssuer349                        = 349  // DATA
	SecurityDesc107                         = 107  // STRING
	EncodedSecurityDescLen350               = 350  // LENGTH
	EncodedSecurityDesc351                  = 351  // DATA
	Pool691                                 = 691  // STRING
	ContractSettlMonth667                   = 667  // MONTHYEAR
	CPProgram875                            = 875  // INT
	CPRegType876                            = 876  // STRING
	DatedDate873                            = 873  // LOCALMKTDATE
	InterestAccrualDate874                  = 874  // LOCALMKTDATE
	SecurityStatus965                       = 965  // STRING
	SettleOnOpenFlag966                     = 966  // STRING
	InstrmtAssignmentMethod1049             = 1049 // CHAR
	StrikeMultiplier967                     = 967  // FLOAT
	StrikeValue968                          = 968  // FLOAT
	MinPriceIncrement969                    = 969  // FLOAT
	PositionLimit970                        = 970  // INT
	NTPositionLimit971                      = 971  // INT
	UnitOfMeasure996                        = 996  // STRING
	TimeUnit997                             = 997  // STRING
	MaturityTime1079                        = 1079 // TZTIMEONLY
	SecurityAltID455                        = 455  // STRING
	SecurityAltIDSource456                  = 456  // STRING
	EventType865                            = 865  // INT
	EventDate866                            = 866  // LOCALMKTDATE
	EventPx867                              = 867  // PRICE
	EventText868                            = 868  // STRING
	InstrumentPartyID1019                   = 1019 // STRING
	InstrumentPartyIDSource1050             = 1050 // CHAR
	InstrumentPartyRole1051                 = 1051 // INT
	InstrumentPartySubID1053                = 1053 // STRING
	InstrumentPartySubIDType1054            = 1054 // INT
	UnderlyingSymbol311                     = 311  // STRING
	UnderlyingSymbolSfx312                  = 312  // STRING
	UnderlyingSecurityID309                 = 309  // STRING
	UnderlyingSecurityIDSource305           = 305  // STRING
	UnderlyingProduct462                    = 462  // INT
	UnderlyingCFICode463                    = 463  // STRING
	UnderlyingSecurityType310               = 310  // STRING
	UnderlyingSecuritySubType763            = 763  // STRING
	UnderlyingMaturityMonthYear313          = 313  // MONTHYEAR
	UnderlyingMaturityDate542               = 542  // LOCALMKTDATE
	UnderlyingCouponPaymentDate241          = 241  // LOCALMKTDATE
	UnderlyingIssueDate242                  = 242  // LOCALMKTDATE
	UnderlyingRepoCollateralSecurityType243 = 243  // INT
	UnderlyingRepurchaseTerm244             = 244  // INT
	UnderlyingRepurchaseRate245             = 245  // PERCENTAGE
	UnderlyingFactor246                     = 246  // FLOAT
	UnderlyingCreditRating256               = 256  // STRING
	UnderlyingInstrRegistry595              = 595  // STRING
	UnderlyingCountryOfIssue592             = 592  // COUNTRY
	UnderlyingStateOrProvinceOfIssue593     = 593  // STRING
	UnderlyingLocaleOfIssue594              = 594  // STRING
	UnderlyingRedemptionDate247             = 247  // LOCALMKTDATE
	UnderlyingStrikePrice316                = 316  // PRICE
	UnderlyingStrikeCurrency941             = 941  // CURRENCY
	UnderlyingOptAttribute317               = 317  // CHAR
	UnderlyingContractMultiplier436         = 436  // FLOAT
	UnderlyingCouponRate435                 = 435  // PERCENTAGE
	UnderlyingSecurityExchange308           = 308  // EXCHANGE
	UnderlyingIssuer306                     = 306  // STRING
	EncodedUnderlyingIssuerLen362           = 362  // LENGTH
	EncodedUnderlyingIssuer363              = 363  // DATA
	UnderlyingSecurityDesc307               = 307  // STRING
	EncodedUnderlyingSecurityDescLen364     = 364  // LENGTH
	EncodedUnderlyingSecurityDesc365        = 365  // DATA
	UnderlyingCPProgram877                  = 877  // STRING
	UnderlyingCPRegType878                  = 878  // STRING
	UnderlyingCurrency318                   = 318  // CURRENCY
	UnderlyingQty879                        = 879  // QTY
	UnderlyingPx810                         = 810  // PRICE
	UnderlyingDirtyPrice882                 = 882  // PRICE
	UnderlyingEndPrice883                   = 883  // PRICE
	UnderlyingStartValue884                 = 884  // AMT
	UnderlyingCurrentValue885               = 885  // AMT
	UnderlyingEndValue886                   = 886  // AMT
	UnderlyingAllocationPercent972          = 972  // PERCENTAGE
	UnderlyingSettlementType975             = 975  // INT
	UnderlyingCashAmount973                 = 973  // AMT
	UnderlyingCashType974                   = 974  // STRING
	UnderlyingUnitOfMeasure998              = 998  // STRING
	UnderlyingTimeUnit1000                  = 1000 // STRING
	UnderlyingCapValue1038                  = 1038 // AMT
	UnderlyingSettlMethod1039               = 1039 // STRING
	UnderlyingAdjustedQuantity1044          = 1044 // QTY
	UnderlyingFXRate1045                    = 1045 // FLOAT
	UnderlyingFXRateCalc1046                = 1046 // CHAR
	UnderlyingSecurityAltID458              = 458  // STRING
	UnderlyingSecurityAltIDSource459        = 459  // STRING
	UnderlyingStipType888                   = 888  // STRING
	UnderlyingStipValue889                  = 889  // STRING
	UndlyInstrumentPartyID1059              = 1059 // STRING
	UndlyInstrumentPartyIDSource1060        = 1060 // CHAR
	UndlyInstrumentPartyRole1061            = 1061 // INT
	UndlyInstrumentPartySubID1063           = 1063 // STRING
	UndlyInstrumentPartySubIDType1064       = 1064 // INT
	PartyID448                              = 448  // STRING
	PartyIDSource447                        = 447  // CHAR
	PartyRole452                            = 452  // INT
	PartySubID523                           = 523  // STRING
	PartySubIDType803                       = 803  // INT
)