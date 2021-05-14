// Code generated by protofix; DO NOT EDIT.
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fix50sp2partydetailslistreport is a format of the FIX.5.0 servicepack 2 PartyDetailsListReport message.
package fix50sp2partydetailslistreport

import (
	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/marshfix"
)

var (
	FIX50SP2PartyDetailsListReportMarshaler   = marshfix.Marshal{Tag: "FIX50SP2", Format: FIX50SP2PartyDetailsListReport}
	FIX50SP2PartyDetailsListReportUnmarshaler = marshfix.Unmarshal{Tag: "FIX50SP2", Format: FIX50SP2PartyDetailsListReport}
)

// FIX50SP2PartyDetailsListReport is a FIX.5.0 servicepack 2 format of the PartyDetailsListReport message which maps the codecs into individual fields.
var FIX50SP2PartyDetailsListReport = f0.Format{
	Fields: map[int]f0.Codec{
		PartyDetailsListReportID1510:               f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		PartyDetailsListRequestID1505:              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		PartyDetailsRequestResult1511:              f0.Fld{Opt, f0.ASCII, f0.IntDefault(0 /* VALID_REQUEST */, 1 /* INVALID_OR_UNSUPPORTED_REQUEST */, 2 /* NO_PARTIES_OR_PARTY_DETAILS_FOUND_THAT_MATCH_SELECTION_CRITERIA */, 3 /* UNSUPPORTED_PARTYLISTRESPONSETYPE */, 4 /* NOT_AUTHORIZED_TO_RETRIEVE_PARTIES_OR_PARTY_DETAILS_DATA */, 5 /* PARTIES_OR_PARTY_DETAILS_DATA_TEMPORARILY_UNAVAILABLE */, 6 /* REQUEST_FOR_PARTIES_DATA_NOT_SUPPORTED */, 99 /* OTHER */), f0.Var{1, 19}},
		TotNoPartyList1512:                         f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		LastFragment893:                            f0.Fld{Opt, f0.ASCII, f0.BoolDefault(false /* NO */, true /* YES */), f0.Con{1}},
		Text58:                                     f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		EncodedTextLen354:                          f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		EncodedText355:                             f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		ApplID1180:                                 f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		ApplSeqNum1181:                             f0.Fld{Opt, f0.ASCII, f0.SeqNum(), f0.Var{1, 19}},
		ApplLastSeqNum1350:                         f0.Fld{Opt, f0.ASCII, f0.SeqNum(), f0.Var{1, 19}},
		ApplResendFlag1352:                         f0.Fld{Opt, f0.ASCII, f0.BoolDefault(), f0.Con{1}},
		PartyID448:                                 f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		PartyIDSource447:                           f0.Fld{Req, f0.ASCII, f0.String("1" /* KOREAN_INVESTOR_ID */, "2" /* TAIWANESE_QUALIFIED_FOREIGN_INVESTOR_ID_QFII_FID */, "3" /* TAIWANESE_TRADING_ACCT */, "4" /* MALAYSIAN_CENTRAL_DEPOSITORY */, "5" /* CHINESE_INVESTOR_ID */, "6" /* UK_NATIONAL_INSURANCE_OR_PENSION_NUMBER */, "7" /* US_SOCIAL_SECURITY_NUMBER */, "8" /* US_EMPLOYER_OR_TAX_ID_NUMBER */, "9" /* AUSTRALIAN_BUSINESS_NUMBER */, "A" /* AUSTRALIAN_TAX_FILE_NUMBER */, "B" /* BIC */, "C" /* GENERALLY_ACCEPTED_MARKET_PARTICIPANT_IDENTIFIER */, "D" /* PROPRIETARY */, "E" /* ISO_COUNTRY_CODE */, "F" /* SETTLEMENT_ENTITY_LOCATION */, "G" /* MIC */, "H" /* CSD_PARTICIPANT_MEMBER_CODE */, "I" /* DIRECTED_BROKER_THREE_CHARACTER_ACRONYM_AS_DEFINED_IN_ISITC_ETC_BEST_PRACTICE_GUIDELINES_DOCUMENT */), f0.Con{1}},
		PartyRole452:                               f0.Fld{Req, f0.ASCII, f0.IntDefault(1 /* EXECUTING_FIRM */, 10 /* SETTLEMENT_LOCATION */, 11 /* ORDER_ORIGINATION_TRADER */, 12 /* EXECUTING_TRADER */, 13 /* ORDER_ORIGINATION_FIRM */, 14 /* GIVEUP_CLEARING_FIRM */, 15 /* CORRESPONDANT_CLEARING_FIRM */, 16 /* EXECUTING_SYSTEM */, 17 /* CONTRA_FIRM */, 18 /* CONTRA_CLEARING_FIRM */, 19 /* SPONSORING_FIRM */, 2 /* BROKER_OF_CREDIT */, 20 /* UNDERLYING_CONTRA_FIRM */, 21 /* CLEARING_ORGANIZATION */, 22 /* EXCHANGE */, 24 /* CUSTOMER_ACCOUNT */, 25 /* CORRESPONDENT_CLEARING_ORGANIZATION */, 26 /* CORRESPONDENT_BROKER */, 27 /* BUYER_SELLER */, 28 /* CUSTODIAN */, 29 /* INTERMEDIARY */, 3 /* CLIENT_ID */, 30 /* AGENT */, 31 /* SUB_CUSTODIAN */, 32 /* BENEFICIARY */, 33 /* INTERESTED_PARTY */, 34 /* REGULATORY_BODY */, 35 /* LIQUIDITY_PROVIDER */, 36 /* ENTERING_TRADER */, 37 /* CONTRA_TRADER */, 38 /* POSITION_ACCOUNT */, 4 /* CLEARING_FIRM */, 5 /* INVESTOR_ID */, 6 /* INTRODUCING_FIRM */, 7 /* ENTERING_FIRM */, 8 /* LOCATE */, 9 /* FUND_MANAGER_CLIENT_ID */, 60 /* INTRODUCING_BROKER */, 41 /* CONTRA_POSITION_ACCOUNT */, 42 /* CONTRA_EXCHANGE */, 43 /* INTERNAL_CARRY_ACCOUNT */, 44 /* ORDER_ENTRY_OPERATOR_ID */, 45 /* SECONDARY_ACCOUNT_NUMBER */, 46 /* FOREIGN_FIRM */, 47 /* THIRD_PARTY_ALLOCATION_FIRM */, 48 /* CLAIMING_ACCOUNT */, 49 /* ASSET_MANAGER */, 50 /* PLEDGOR_ACCOUNT */, 51 /* PLEDGEE_ACCOUNT */, 52 /* LARGE_TRADER_REPORTABLE_ACCOUNT */, 53 /* TRADER_MNEMONIC */, 54 /* SENDER_LOCATION */, 55 /* SESSION_ID */, 56 /* ACCEPTABLE_COUNTERPARTY */, 57 /* UNACCEPTABLE_COUNTERPARTY */, 58 /* ENTERING_UNIT */, 59 /* EXECUTING_UNIT */, 39 /* CONTRA_INVESTOR_ID */, 40 /* TRANSFER_TO_FIRM */, 61 /* QUOTE_ORIGINATOR */, 62 /* REPORT_ORIGINATOR */, 63 /* SYSTEMATIC_INTERNALISER */, 64 /* MULTILATERAL_TRADING_FACILITY */, 65 /* REGULATED_MARKET */, 66 /* MARKET_MAKER */, 67 /* INVESTMENT_FIRM */, 68 /* HOST_COMPETENT_AUTHORITY */, 69 /* HOME_COMPETENT_AUTHORITY */, 70 /* COMPETENT_AUTHORITY_OF_THE_MOST_RELEVANT_MARKET_IN_TERMS_OF_LIQUIDITY */, 71 /* COMPETENT_AUTHORITY_OF_THE_TRANSACTION */, 72 /* REPORTING_INTERMEDIARY */, 73 /* EXECUTION_VENUE */, 74 /* MARKET_DATA_ENTRY_ORIGINATOR */, 75 /* LOCATION_ID */, 76 /* DESK_ID */, 77 /* MARKET_DATA_MARKET */, 78 /* ALLOCATION_ENTITY */, 79 /* PRIME_BROKER_PROVIDING_GENERAL_TRADE_SERVICES */, 80 /* STEP_OUT_FIRM */, 81 /* BROKERCLEARINGID */, 82 /* CENTRAL_REGISTRATION_DEPOSITORY */, 83 /* CLEARING_ACCOUNT */, 84 /* ACCEPTABLE_SETTLING_COUNTERPARTY */, 85 /* UNACCEPTABLE_SETTLING_COUNTERPARTY */), f0.Var{1, 19}},
		PartySubID523:                              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		PartySubIDType803:                          f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* FIRM */, 10 /* SECURITIES_ACCOUNT_NUMBER */, 11 /* REGISTRATION_NUMBER */, 12 /* REGISTERED_ADDRESS_12 */, 13 /* REGULATORY_STATUS */, 14 /* REGISTRATION_NAME */, 15 /* CASH_ACCOUNT_NUMBER */, 16 /* BIC */, 17 /* CSD_PARTICIPANT_MEMBER_CODE */, 18 /* REGISTERED_ADDRESS_18 */, 19 /* FUND_ACCOUNT_NAME */, 2 /* PERSON */, 20 /* TELEX_NUMBER */, 21 /* FAX_NUMBER */, 22 /* SECURITIES_ACCOUNT_NAME */, 23 /* CASH_ACCOUNT_NAME */, 24 /* DEPARTMENT */, 25 /* LOCATION_DESK */, 26 /* POSITION_ACCOUNT_TYPE */, 3 /* SYSTEM */, 4 /* APPLICATION */, 5 /* FULL_LEGAL_NAME_OF_FIRM */, 6 /* POSTAL_ADDRESS */, 7 /* PHONE_NUMBER */, 8 /* EMAIL_ADDRESS */, 9 /* CONTACT_NAME */, 27 /* SECURITY_LOCATE_ID */, 28 /* MARKET_MAKER */, 29 /* ELIGIBLE_COUNTERPARTY */, 30 /* PROFESSIONAL_CLIENT */, 31 /* LOCATION */, 32 /* EXECUTION_VENUE */, 33 /* CURRENCY_DELIVERY_IDENTIFIER */), f0.Var{1, 19}},
		PartyAltID1517:                             f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		PartyAltIDSource1518:                       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		PartyAltSubID1520:                          f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		PartyAltSubIDType1521:                      f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		ContextPartyID1523:                         f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		ContextPartyIDSource1524:                   f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		ContextPartyRole1525:                       f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		ContextPartySubID1527:                      f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		ContextPartySubIDType1528:                  f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		RiskLimitType1530:                          f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* GROSS_LIMIT */, 2 /* NET_LIMIT */, 3 /* EXPOSURE */, 4 /* LONG_LIMIT */, 5 /* SHORT_LIMIT */), f0.Var{1, 19}},
		RiskLimitAmount1531:                        f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		RiskLimitCurrency1532:                      f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{3}},
		RiskLimitPlatform1533:                      f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RiskInstrumentOperator1535:                 f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* INCLUDE */, 2 /* EXCLUDE */), f0.Var{1, 19}},
		RiskSymbol1536:                             f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RiskSymbolSfx1537:                          f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RiskSecurityID1538:                         f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RiskSecurityIDSource1539:                   f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RiskProduct1543:                            f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		RiskProductComplex1544:                     f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RiskSecurityGroup1545:                      f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RiskCFICode1546:                            f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RiskSecurityType1547:                       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RiskSecuritySubType1548:                    f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RiskMaturityMonthYear1549:                  f0.Fld{Opt, f0.ASCII, f0.MonthYearTime(), f0.Var{6, 8}},
		RiskMaturityTime1550:                       f0.Fld{Opt, f0.ASCII, f0.TZTime(), f0.Var{8, 12}},
		RiskRestructuringType1551:                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RiskSeniority1552:                          f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RiskPutOrCall1553:                          f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		RiskFlexibleIndicator1554:                  f0.Fld{Opt, f0.ASCII, f0.BoolDefault(), f0.Con{1}},
		RiskCouponRate1555:                         f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		RiskSecurityExchange1616:                   f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RiskSecurityDesc1556:                       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RiskEncodedSecurityDescLen1620:             f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		RiskEncodedSecurityDesc1621:                f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		RiskInstrumentSettlType1557:                f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RiskInstrumentMultiplier1558:               f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		RiskSecurityAltID1541:                      f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RiskSecurityAltIDSource1542:                f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RiskWarningLevelPercent1560:                f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		RiskWarningLevelName1561:                   f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RelatedPartyID1563:                         f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RelatedPartyIDSource1564:                   f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		RelatedPartyRole1565:                       f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		RelatedPartySubID1567:                      f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RelatedPartySubIDType1568:                  f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		RelatedPartyAltID1570:                      f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RelatedPartyAltIDSource1571:                f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		RelatedPartyAltSubID1573:                   f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RelatedPartyAltSubIDType1574:               f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		RelatedContextPartyID1576:                  f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RelatedContextPartyIDSource1577:            f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		RelatedContextPartyRole1578:                f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		RelatedContextPartySubID1580:               f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RelatedContextPartySubIDType1581:           f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		RelationshipRiskLimitType1583:              f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		RelationshipRiskLimitAmount1584:            f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		RelationshipRiskLimitCurrency1585:          f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{3}},
		RelationshipRiskLimitPlatform1586:          f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RelationshipRiskInstrumentOperator1588:     f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		RelationshipRiskSymbol1589:                 f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RelationshipRiskSymbolSfx1590:              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RelationshipRiskSecurityID1591:             f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RelationshipRiskSecurityIDSource1592:       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RelationshipRiskProduct1596:                f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		RelationshipRiskProductComplex1597:         f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RelationshipRiskSecurityGroup1598:          f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RelationshipRiskCFICode1599:                f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RelationshipRiskSecurityType1600:           f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RelationshipRiskSecuritySubType1601:        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RelationshipRiskMaturityMonthYear1602:      f0.Fld{Opt, f0.ASCII, f0.MonthYearTime(), f0.Var{6, 8}},
		RelationshipRiskMaturityTime1603:           f0.Fld{Opt, f0.ASCII, f0.TZTime(), f0.Var{8, 12}},
		RelationshipRiskRestructuringType1604:      f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RelationshipRiskSeniority1605:              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RelationshipRiskPutOrCall1606:              f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		RelationshipRiskFlexibleIndicator1607:      f0.Fld{Opt, f0.ASCII, f0.BoolDefault(), f0.Con{1}},
		RelationshipRiskCouponRate1608:             f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		RelationshipRiskSecurityExchange1609:       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RelationshipRiskSecurityDesc1610:           f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RelationshipRiskEncodedSecurityDescLen1618: f0.Fld{Opt, f0.ASCII, f0.Length(), f0.Var{1, 19}},
		RelationshipRiskEncodedSecurityDesc1619:    f0.Fld{Opt, f0.ASCII, f0.Bytes(), f0.Var{1, 65536}},
		RelationshipRiskInstrumentSettlType1611:    f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RelationshipRiskInstrumentMultiplier1612:   f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		RelationshipRiskSecurityAltID1594:          f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RelationshipRiskSecurityAltIDSource1595:    f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RelationshipRiskWarningLevelPercent1614:    f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		RelationshipRiskWarningLevelName1615:       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		PartyRelationship1515:                      f0.Fld{Opt, f0.ASCII, f0.IntDefault(0 /* IS_ALSO */, 1 /* CLEARS_FOR */, 2 /* CLEARS_THROUGH */, 3 /* TRADES_FOR */, 4 /* TRADES_THROUGH */, 5 /* SPONSORS */, 6 /* SPONSORED_THROUGH */, 7 /* PROVIDES_GUARANTEE_FOR */, 8 /* IS_GUARANTEED_BY */, 9 /* MEMBER_OF */, 10 /* HAS_MEMBERS */, 11 /* PROVIDES_MARKETPLACE_FOR */, 12 /* PARTICIPANT_OF_MARKETPLACE */, 13 /* CARRIES_POSITIONS_FOR */, 14 /* POSTS_TRADES_TO */, 15 /* ENTERS_TRADES_FOR */, 16 /* ENTERS_TRADES_THROUGH */, 17 /* PROVIDES_QUOTES_TO */, 18 /* REQUESTS_QUOTES_FROM */, 19 /* INVESTS_FOR */, 20 /* INVESTS_THROUGH */, 21 /* BROKERS_TRADES_FOR */, 22 /* BROKERS_TRADES_THROUGH */, 23 /* PROVIDES_TRADING_SERVICES_FOR */, 24 /* USES_TRADING_SERVICES_OF */, 25 /* APPROVES_OF */, 26 /* APPROVED_BY */, 27 /* PARENT_FIRM_FOR */, 28 /* SUBSIDIARY_OF */, 29 /* REGULATORY_OWNER_OF */, 30 /* OWNED_BY_30 */, 31 /* CONTROLS */, 32 /* IS_CONTROLLED_BY */, 33 /* LEGAL */, 34 /* OWNED_BY_34 */, 35 /* BENEFICIAL_OWNER_OF */, 36 /* OWNED_BY_36 */), f0.Var{1, 19}},
	},
	Unknown0: f0.UnknownFld{f0.ASCII, f0.Unknown(), f0.Var{1, 65536}},
	Sort: []int{
		PartyDetailsListReportID1510,               // STRING
		PartyDetailsListRequestID1505,              // STRING
		PartyDetailsRequestResult1511,              // INT
		TotNoPartyList1512,                         // INT
		LastFragment893,                            // BOOLEAN
		Text58,                                     // STRING
		EncodedTextLen354,                          // LENGTH
		EncodedText355,                             // DATA
		ApplID1180,                                 // STRING
		ApplSeqNum1181,                             // SEQNUM
		ApplLastSeqNum1350,                         // SEQNUM
		ApplResendFlag1352,                         // BOOLEAN
		PartyID448,                                 // STRING
		PartyIDSource447,                           // CHAR
		PartyRole452,                               // INT
		PartySubID523,                              // STRING
		PartySubIDType803,                          // INT
		PartyAltID1517,                             // STRING
		PartyAltIDSource1518,                       // CHAR
		PartyAltSubID1520,                          // STRING
		PartyAltSubIDType1521,                      // INT
		ContextPartyID1523,                         // STRING
		ContextPartyIDSource1524,                   // CHAR
		ContextPartyRole1525,                       // INT
		ContextPartySubID1527,                      // STRING
		ContextPartySubIDType1528,                  // INT
		RiskLimitType1530,                          // INT
		RiskLimitAmount1531,                        // AMT
		RiskLimitCurrency1532,                      // CURRENCY
		RiskLimitPlatform1533,                      // STRING
		RiskInstrumentOperator1535,                 // INT
		RiskSymbol1536,                             // STRING
		RiskSymbolSfx1537,                          // STRING
		RiskSecurityID1538,                         // STRING
		RiskSecurityIDSource1539,                   // STRING
		RiskProduct1543,                            // INT
		RiskProductComplex1544,                     // STRING
		RiskSecurityGroup1545,                      // STRING
		RiskCFICode1546,                            // STRING
		RiskSecurityType1547,                       // STRING
		RiskSecuritySubType1548,                    // STRING
		RiskMaturityMonthYear1549,                  // MONTHYEAR
		RiskMaturityTime1550,                       // TZTIMEONLY
		RiskRestructuringType1551,                  // STRING
		RiskSeniority1552,                          // STRING
		RiskPutOrCall1553,                          // INT
		RiskFlexibleIndicator1554,                  // BOOLEAN
		RiskCouponRate1555,                         // PERCENTAGE
		RiskSecurityExchange1616,                   // EXCHANGE
		RiskSecurityDesc1556,                       // STRING
		RiskEncodedSecurityDescLen1620,             // LENGTH
		RiskEncodedSecurityDesc1621,                // DATA
		RiskInstrumentSettlType1557,                // STRING
		RiskInstrumentMultiplier1558,               // FLOAT
		RiskSecurityAltID1541,                      // STRING
		RiskSecurityAltIDSource1542,                // STRING
		RiskWarningLevelPercent1560,                // PERCENTAGE
		RiskWarningLevelName1561,                   // STRING
		RelatedPartyID1563,                         // STRING
		RelatedPartyIDSource1564,                   // CHAR
		RelatedPartyRole1565,                       // INT
		RelatedPartySubID1567,                      // STRING
		RelatedPartySubIDType1568,                  // INT
		RelatedPartyAltID1570,                      // STRING
		RelatedPartyAltIDSource1571,                // CHAR
		RelatedPartyAltSubID1573,                   // STRING
		RelatedPartyAltSubIDType1574,               // INT
		RelatedContextPartyID1576,                  // STRING
		RelatedContextPartyIDSource1577,            // CHAR
		RelatedContextPartyRole1578,                // INT
		RelatedContextPartySubID1580,               // STRING
		RelatedContextPartySubIDType1581,           // INT
		RelationshipRiskLimitType1583,              // INT
		RelationshipRiskLimitAmount1584,            // AMT
		RelationshipRiskLimitCurrency1585,          // CURRENCY
		RelationshipRiskLimitPlatform1586,          // STRING
		RelationshipRiskInstrumentOperator1588,     // INT
		RelationshipRiskSymbol1589,                 // STRING
		RelationshipRiskSymbolSfx1590,              // STRING
		RelationshipRiskSecurityID1591,             // STRING
		RelationshipRiskSecurityIDSource1592,       // STRING
		RelationshipRiskProduct1596,                // INT
		RelationshipRiskProductComplex1597,         // STRING
		RelationshipRiskSecurityGroup1598,          // STRING
		RelationshipRiskCFICode1599,                // STRING
		RelationshipRiskSecurityType1600,           // STRING
		RelationshipRiskSecuritySubType1601,        // STRING
		RelationshipRiskMaturityMonthYear1602,      // MONTHYEAR
		RelationshipRiskMaturityTime1603,           // TZTIMEONLY
		RelationshipRiskRestructuringType1604,      // STRING
		RelationshipRiskSeniority1605,              // STRING
		RelationshipRiskPutOrCall1606,              // INT
		RelationshipRiskFlexibleIndicator1607,      // BOOLEAN
		RelationshipRiskCouponRate1608,             // PERCENTAGE
		RelationshipRiskSecurityExchange1609,       // EXCHANGE
		RelationshipRiskSecurityDesc1610,           // STRING
		RelationshipRiskEncodedSecurityDescLen1618, // LENGTH
		RelationshipRiskEncodedSecurityDesc1619,    // DATA
		RelationshipRiskInstrumentSettlType1611,    // STRING
		RelationshipRiskInstrumentMultiplier1612,   // FLOAT
		RelationshipRiskSecurityAltID1594,          // STRING
		RelationshipRiskSecurityAltIDSource1595,    // STRING
		RelationshipRiskWarningLevelPercent1614,    // PERCENTAGE
		RelationshipRiskWarningLevelName1615,       // STRING
		PartyRelationship1515,                      // INT
	},
}

const Req, Opt = true, false

const (
	PartyDetailsListReportID1510               = 1510 // STRING
	PartyDetailsListRequestID1505              = 1505 // STRING
	PartyDetailsRequestResult1511              = 1511 // INT
	TotNoPartyList1512                         = 1512 // INT
	LastFragment893                            = 893  // BOOLEAN
	Text58                                     = 58   // STRING
	EncodedTextLen354                          = 354  // LENGTH
	EncodedText355                             = 355  // DATA
	ApplID1180                                 = 1180 // STRING
	ApplSeqNum1181                             = 1181 // SEQNUM
	ApplLastSeqNum1350                         = 1350 // SEQNUM
	ApplResendFlag1352                         = 1352 // BOOLEAN
	PartyID448                                 = 448  // STRING
	PartyIDSource447                           = 447  // CHAR
	PartyRole452                               = 452  // INT
	PartySubID523                              = 523  // STRING
	PartySubIDType803                          = 803  // INT
	PartyAltID1517                             = 1517 // STRING
	PartyAltIDSource1518                       = 1518 // CHAR
	PartyAltSubID1520                          = 1520 // STRING
	PartyAltSubIDType1521                      = 1521 // INT
	ContextPartyID1523                         = 1523 // STRING
	ContextPartyIDSource1524                   = 1524 // CHAR
	ContextPartyRole1525                       = 1525 // INT
	ContextPartySubID1527                      = 1527 // STRING
	ContextPartySubIDType1528                  = 1528 // INT
	RiskLimitType1530                          = 1530 // INT
	RiskLimitAmount1531                        = 1531 // AMT
	RiskLimitCurrency1532                      = 1532 // CURRENCY
	RiskLimitPlatform1533                      = 1533 // STRING
	RiskInstrumentOperator1535                 = 1535 // INT
	RiskSymbol1536                             = 1536 // STRING
	RiskSymbolSfx1537                          = 1537 // STRING
	RiskSecurityID1538                         = 1538 // STRING
	RiskSecurityIDSource1539                   = 1539 // STRING
	RiskProduct1543                            = 1543 // INT
	RiskProductComplex1544                     = 1544 // STRING
	RiskSecurityGroup1545                      = 1545 // STRING
	RiskCFICode1546                            = 1546 // STRING
	RiskSecurityType1547                       = 1547 // STRING
	RiskSecuritySubType1548                    = 1548 // STRING
	RiskMaturityMonthYear1549                  = 1549 // MONTHYEAR
	RiskMaturityTime1550                       = 1550 // TZTIMEONLY
	RiskRestructuringType1551                  = 1551 // STRING
	RiskSeniority1552                          = 1552 // STRING
	RiskPutOrCall1553                          = 1553 // INT
	RiskFlexibleIndicator1554                  = 1554 // BOOLEAN
	RiskCouponRate1555                         = 1555 // PERCENTAGE
	RiskSecurityExchange1616                   = 1616 // EXCHANGE
	RiskSecurityDesc1556                       = 1556 // STRING
	RiskEncodedSecurityDescLen1620             = 1620 // LENGTH
	RiskEncodedSecurityDesc1621                = 1621 // DATA
	RiskInstrumentSettlType1557                = 1557 // STRING
	RiskInstrumentMultiplier1558               = 1558 // FLOAT
	RiskSecurityAltID1541                      = 1541 // STRING
	RiskSecurityAltIDSource1542                = 1542 // STRING
	RiskWarningLevelPercent1560                = 1560 // PERCENTAGE
	RiskWarningLevelName1561                   = 1561 // STRING
	RelatedPartyID1563                         = 1563 // STRING
	RelatedPartyIDSource1564                   = 1564 // CHAR
	RelatedPartyRole1565                       = 1565 // INT
	RelatedPartySubID1567                      = 1567 // STRING
	RelatedPartySubIDType1568                  = 1568 // INT
	RelatedPartyAltID1570                      = 1570 // STRING
	RelatedPartyAltIDSource1571                = 1571 // CHAR
	RelatedPartyAltSubID1573                   = 1573 // STRING
	RelatedPartyAltSubIDType1574               = 1574 // INT
	RelatedContextPartyID1576                  = 1576 // STRING
	RelatedContextPartyIDSource1577            = 1577 // CHAR
	RelatedContextPartyRole1578                = 1578 // INT
	RelatedContextPartySubID1580               = 1580 // STRING
	RelatedContextPartySubIDType1581           = 1581 // INT
	RelationshipRiskLimitType1583              = 1583 // INT
	RelationshipRiskLimitAmount1584            = 1584 // AMT
	RelationshipRiskLimitCurrency1585          = 1585 // CURRENCY
	RelationshipRiskLimitPlatform1586          = 1586 // STRING
	RelationshipRiskInstrumentOperator1588     = 1588 // INT
	RelationshipRiskSymbol1589                 = 1589 // STRING
	RelationshipRiskSymbolSfx1590              = 1590 // STRING
	RelationshipRiskSecurityID1591             = 1591 // STRING
	RelationshipRiskSecurityIDSource1592       = 1592 // STRING
	RelationshipRiskProduct1596                = 1596 // INT
	RelationshipRiskProductComplex1597         = 1597 // STRING
	RelationshipRiskSecurityGroup1598          = 1598 // STRING
	RelationshipRiskCFICode1599                = 1599 // STRING
	RelationshipRiskSecurityType1600           = 1600 // STRING
	RelationshipRiskSecuritySubType1601        = 1601 // STRING
	RelationshipRiskMaturityMonthYear1602      = 1602 // MONTHYEAR
	RelationshipRiskMaturityTime1603           = 1603 // TZTIMEONLY
	RelationshipRiskRestructuringType1604      = 1604 // STRING
	RelationshipRiskSeniority1605              = 1605 // STRING
	RelationshipRiskPutOrCall1606              = 1606 // INT
	RelationshipRiskFlexibleIndicator1607      = 1607 // BOOLEAN
	RelationshipRiskCouponRate1608             = 1608 // PERCENTAGE
	RelationshipRiskSecurityExchange1609       = 1609 // EXCHANGE
	RelationshipRiskSecurityDesc1610           = 1610 // STRING
	RelationshipRiskEncodedSecurityDescLen1618 = 1618 // LENGTH
	RelationshipRiskEncodedSecurityDesc1619    = 1619 // DATA
	RelationshipRiskInstrumentSettlType1611    = 1611 // STRING
	RelationshipRiskInstrumentMultiplier1612   = 1612 // FLOAT
	RelationshipRiskSecurityAltID1594          = 1594 // STRING
	RelationshipRiskSecurityAltIDSource1595    = 1595 // STRING
	RelationshipRiskWarningLevelPercent1614    = 1614 // PERCENTAGE
	RelationshipRiskWarningLevelName1615       = 1615 // STRING
	PartyRelationship1515                      = 1515 // INT
)