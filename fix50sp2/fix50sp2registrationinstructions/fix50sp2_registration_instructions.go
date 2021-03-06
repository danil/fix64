// Code generated by protofix; DO NOT EDIT.
// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fix50sp2registrationinstructions is a format of the FIX.5.0 servicepack 2 RegistrationInstructions message.
package fix50sp2registrationinstructions

import (
	f0 "github.com/protofix/protofix/codecfix"
	"github.com/protofix/protofix/marshfix"
)

var (
	FIX50SP2RegistrationInstructionsMarshaler   = marshfix.Marshal{Tag: "FIX50SP2", Format: FIX50SP2RegistrationInstructions}
	FIX50SP2RegistrationInstructionsUnmarshaler = marshfix.Unmarshal{Tag: "FIX50SP2", Format: FIX50SP2RegistrationInstructions}
)

// FIX50SP2RegistrationInstructions is a FIX.5.0 servicepack 2 format of the RegistrationInstructions message which maps the codecs into individual fields.
var FIX50SP2RegistrationInstructions = f0.Format{
	Fields: map[int]f0.Codec{
		RegistID513:                   f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RegistTransType514:            f0.Fld{Req, f0.ASCII, f0.String("0" /* NEW */, "1" /* REPLACE */, "2" /* CANCEL */), f0.Con{1}},
		RegistRefID508:                f0.Fld{Req, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		ClOrdID11:                     f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		Account1:                      f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		AcctIDSource660:               f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* BIC */, 2 /* SID_CODE */, 3 /* TFM */, 4 /* OMGEO */, 5 /* DTCC_CODE */, 99 /* OTHER */), f0.Var{1, 19}},
		RegistAcctType493:             f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		TaxAdvantageType495:           f0.Fld{Opt, f0.ASCII, f0.IntDefault(0 /* NONE_NOT_APPLICABLE */, 1 /* MAXI_ISA */, 10 /* EMPLOYEE_10 */, 11 /* EMPLOYER_11 */, 12 /* EMPLOYER_12 */, 13 /* NON_FUND_PROTOTYPE_IRA */, 14 /* NON_FUND_QUALIFIED_PLAN */, 15 /* DEFINED_CONTRIBUTION_PLAN */, 16 /* INDIVIDUAL_RETIREMENT_ACCOUNT_16 */, 17 /* INDIVIDUAL_RETIREMENT_ACCOUNT_17 */, 18 /* KEOGH */, 19 /* PROFIT_SHARING_PLAN */, 2 /* TESSA */, 20 /* 401 */, 21 /* SELF_DIRECTED_IRA */, 22 /* 403 */, 23 /* 457 */, 24 /* ROTH_IRA_24 */, 25 /* ROTH_IRA_25 */, 26 /* ROTH_CONVERSION_IRA_26 */, 27 /* ROTH_CONVERSION_IRA_27 */, 28 /* EDUCATION_IRA_28 */, 29 /* EDUCATION_IRA_29 */, 3 /* MINI_CASH_ISA */, 4 /* MINI_STOCKS_AND_SHARES_ISA */, 5 /* MINI_INSURANCE_ISA */, 6 /* CURRENT_YEAR_PAYMENT */, 7 /* PRIOR_YEAR_PAYMENT */, 8 /* ASSET_TRANSFER */, 9 /* EMPLOYEE_9 */, 999 /* OTHER */), f0.Var{1, 19}},
		OwnershipType517:              f0.Fld{Opt, f0.ASCII, f0.String("2" /* JOINT_TRUSTEES */, "J" /* JOINT_INVESTORS */, "T" /* TENANTS_IN_COMMON */), f0.Con{1}},
		PartyID448:                    f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		PartyIDSource447:              f0.Fld{Opt, f0.ASCII, f0.String("1" /* KOREAN_INVESTOR_ID */, "2" /* TAIWANESE_QUALIFIED_FOREIGN_INVESTOR_ID_QFII_FID */, "3" /* TAIWANESE_TRADING_ACCT */, "4" /* MALAYSIAN_CENTRAL_DEPOSITORY */, "5" /* CHINESE_INVESTOR_ID */, "6" /* UK_NATIONAL_INSURANCE_OR_PENSION_NUMBER */, "7" /* US_SOCIAL_SECURITY_NUMBER */, "8" /* US_EMPLOYER_OR_TAX_ID_NUMBER */, "9" /* AUSTRALIAN_BUSINESS_NUMBER */, "A" /* AUSTRALIAN_TAX_FILE_NUMBER */, "B" /* BIC */, "C" /* GENERALLY_ACCEPTED_MARKET_PARTICIPANT_IDENTIFIER */, "D" /* PROPRIETARY */, "E" /* ISO_COUNTRY_CODE */, "F" /* SETTLEMENT_ENTITY_LOCATION */, "G" /* MIC */, "H" /* CSD_PARTICIPANT_MEMBER_CODE */, "I" /* DIRECTED_BROKER_THREE_CHARACTER_ACRONYM_AS_DEFINED_IN_ISITC_ETC_BEST_PRACTICE_GUIDELINES_DOCUMENT */), f0.Con{1}},
		PartyRole452:                  f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* EXECUTING_FIRM */, 10 /* SETTLEMENT_LOCATION */, 11 /* ORDER_ORIGINATION_TRADER */, 12 /* EXECUTING_TRADER */, 13 /* ORDER_ORIGINATION_FIRM */, 14 /* GIVEUP_CLEARING_FIRM */, 15 /* CORRESPONDANT_CLEARING_FIRM */, 16 /* EXECUTING_SYSTEM */, 17 /* CONTRA_FIRM */, 18 /* CONTRA_CLEARING_FIRM */, 19 /* SPONSORING_FIRM */, 2 /* BROKER_OF_CREDIT */, 20 /* UNDERLYING_CONTRA_FIRM */, 21 /* CLEARING_ORGANIZATION */, 22 /* EXCHANGE */, 24 /* CUSTOMER_ACCOUNT */, 25 /* CORRESPONDENT_CLEARING_ORGANIZATION */, 26 /* CORRESPONDENT_BROKER */, 27 /* BUYER_SELLER */, 28 /* CUSTODIAN */, 29 /* INTERMEDIARY */, 3 /* CLIENT_ID */, 30 /* AGENT */, 31 /* SUB_CUSTODIAN */, 32 /* BENEFICIARY */, 33 /* INTERESTED_PARTY */, 34 /* REGULATORY_BODY */, 35 /* LIQUIDITY_PROVIDER */, 36 /* ENTERING_TRADER */, 37 /* CONTRA_TRADER */, 38 /* POSITION_ACCOUNT */, 4 /* CLEARING_FIRM */, 5 /* INVESTOR_ID */, 6 /* INTRODUCING_FIRM */, 7 /* ENTERING_FIRM */, 8 /* LOCATE */, 9 /* FUND_MANAGER_CLIENT_ID */, 60 /* INTRODUCING_BROKER */, 41 /* CONTRA_POSITION_ACCOUNT */, 42 /* CONTRA_EXCHANGE */, 43 /* INTERNAL_CARRY_ACCOUNT */, 44 /* ORDER_ENTRY_OPERATOR_ID */, 45 /* SECONDARY_ACCOUNT_NUMBER */, 46 /* FOREIGN_FIRM */, 47 /* THIRD_PARTY_ALLOCATION_FIRM */, 48 /* CLAIMING_ACCOUNT */, 49 /* ASSET_MANAGER */, 50 /* PLEDGOR_ACCOUNT */, 51 /* PLEDGEE_ACCOUNT */, 52 /* LARGE_TRADER_REPORTABLE_ACCOUNT */, 53 /* TRADER_MNEMONIC */, 54 /* SENDER_LOCATION */, 55 /* SESSION_ID */, 56 /* ACCEPTABLE_COUNTERPARTY */, 57 /* UNACCEPTABLE_COUNTERPARTY */, 58 /* ENTERING_UNIT */, 59 /* EXECUTING_UNIT */, 39 /* CONTRA_INVESTOR_ID */, 40 /* TRANSFER_TO_FIRM */, 61 /* QUOTE_ORIGINATOR */, 62 /* REPORT_ORIGINATOR */, 63 /* SYSTEMATIC_INTERNALISER */, 64 /* MULTILATERAL_TRADING_FACILITY */, 65 /* REGULATED_MARKET */, 66 /* MARKET_MAKER */, 67 /* INVESTMENT_FIRM */, 68 /* HOST_COMPETENT_AUTHORITY */, 69 /* HOME_COMPETENT_AUTHORITY */, 70 /* COMPETENT_AUTHORITY_OF_THE_MOST_RELEVANT_MARKET_IN_TERMS_OF_LIQUIDITY */, 71 /* COMPETENT_AUTHORITY_OF_THE_TRANSACTION */, 72 /* REPORTING_INTERMEDIARY */, 73 /* EXECUTION_VENUE */, 74 /* MARKET_DATA_ENTRY_ORIGINATOR */, 75 /* LOCATION_ID */, 76 /* DESK_ID */, 77 /* MARKET_DATA_MARKET */, 78 /* ALLOCATION_ENTITY */, 79 /* PRIME_BROKER_PROVIDING_GENERAL_TRADE_SERVICES */, 80 /* STEP_OUT_FIRM */, 81 /* BROKERCLEARINGID */, 82 /* CENTRAL_REGISTRATION_DEPOSITORY */, 83 /* CLEARING_ACCOUNT */, 84 /* ACCEPTABLE_SETTLING_COUNTERPARTY */, 85 /* UNACCEPTABLE_SETTLING_COUNTERPARTY */), f0.Var{1, 19}},
		PartySubID523:                 f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		PartySubIDType803:             f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* FIRM */, 10 /* SECURITIES_ACCOUNT_NUMBER */, 11 /* REGISTRATION_NUMBER */, 12 /* REGISTERED_ADDRESS_12 */, 13 /* REGULATORY_STATUS */, 14 /* REGISTRATION_NAME */, 15 /* CASH_ACCOUNT_NUMBER */, 16 /* BIC */, 17 /* CSD_PARTICIPANT_MEMBER_CODE */, 18 /* REGISTERED_ADDRESS_18 */, 19 /* FUND_ACCOUNT_NAME */, 2 /* PERSON */, 20 /* TELEX_NUMBER */, 21 /* FAX_NUMBER */, 22 /* SECURITIES_ACCOUNT_NAME */, 23 /* CASH_ACCOUNT_NAME */, 24 /* DEPARTMENT */, 25 /* LOCATION_DESK */, 26 /* POSITION_ACCOUNT_TYPE */, 3 /* SYSTEM */, 4 /* APPLICATION */, 5 /* FULL_LEGAL_NAME_OF_FIRM */, 6 /* POSTAL_ADDRESS */, 7 /* PHONE_NUMBER */, 8 /* EMAIL_ADDRESS */, 9 /* CONTACT_NAME */, 27 /* SECURITY_LOCATE_ID */, 28 /* MARKET_MAKER */, 29 /* ELIGIBLE_COUNTERPARTY */, 30 /* PROFESSIONAL_CLIENT */, 31 /* LOCATION */, 32 /* EXECUTION_VENUE */, 33 /* CURRENCY_DELIVERY_IDENTIFIER */), f0.Var{1, 19}},
		RegistDtls509:                 f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		RegistEmail511:                f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		MailingDtls474:                f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		MailingInst482:                f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		OwnerType522:                  f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* INDIVIDUAL_INVESTOR */, 10 /* NETWORKING_SUB_ACCOUNT */, 11 /* NON_PROFIT_ORGANIZATION */, 12 /* CORPORATE_BODY */, 13 /* NOMINEE */, 2 /* PUBLIC_COMPANY */, 3 /* PRIVATE_COMPANY */, 4 /* INDIVIDUAL_TRUSTEE */, 5 /* COMPANY_TRUSTEE */, 6 /* PENSION_PLAN */, 7 /* CUSTODIAN_UNDER_GIFTS_TO_MINORS_ACT */, 8 /* TRUSTS */, 9 /* FIDUCIARIES */), f0.Var{1, 19}},
		DateOfBirth486:                f0.Fld{Opt, f0.ASCII, f0.LocalMktDateTime(), f0.Con{8}},
		InvestorCountryOfResidence475: f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{2}},
		NestedPartyID524:              f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		NestedPartyIDSource525:        f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{1}},
		NestedPartyRole538:            f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		NestedPartySubID545:           f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		NestedPartySubIDType805:       f0.Fld{Opt, f0.ASCII, f0.IntDefault(), f0.Var{1, 19}},
		DistribPaymentMethod477:       f0.Fld{Opt, f0.ASCII, f0.IntDefault(1 /* CREST */, 10 /* BPAY */, 11 /* HIGH_VALUE_CLEARING_SYSTEM_HVACS */, 12 /* REINVEST_IN_FUND */, 2 /* NSCC */, 3 /* EUROCLEAR */, 4 /* CLEARSTREAM */, 5 /* CHEQUE */, 6 /* TELEGRAPHIC_TRANSFER */, 7 /* FED_WIRE */, 8 /* DIRECT_CREDIT */, 9 /* ACH_CREDIT */), f0.Var{1, 19}},
		DistribPercentage512:          f0.Fld{Opt, f0.ASCII, f0.Float(), f0.Var{1, 19}},
		CashDistribCurr478:            f0.Fld{Opt, f0.ASCII, f0.String(), f0.Con{3}},
		CashDistribAgentName498:       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		CashDistribAgentCode499:       f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		CashDistribAgentAcctNumber500: f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		CashDistribPayRef501:          f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
		CashDistribAgentAcctName502:   f0.Fld{Opt, f0.ASCII, f0.String(), f0.Var{1, 65536}},
	},
	Unknown0: f0.UnknownFld{f0.ASCII, f0.Unknown(), f0.Var{1, 65536}},
	Sort: []int{
		RegistID513,                   // STRING
		RegistTransType514,            // CHAR
		RegistRefID508,                // STRING
		ClOrdID11,                     // STRING
		Account1,                      // STRING
		AcctIDSource660,               // INT
		RegistAcctType493,             // STRING
		TaxAdvantageType495,           // INT
		OwnershipType517,              // CHAR
		PartyID448,                    // STRING
		PartyIDSource447,              // CHAR
		PartyRole452,                  // INT
		PartySubID523,                 // STRING
		PartySubIDType803,             // INT
		RegistDtls509,                 // STRING
		RegistEmail511,                // STRING
		MailingDtls474,                // STRING
		MailingInst482,                // STRING
		OwnerType522,                  // INT
		DateOfBirth486,                // LOCALMKTDATE
		InvestorCountryOfResidence475, // COUNTRY
		NestedPartyID524,              // STRING
		NestedPartyIDSource525,        // CHAR
		NestedPartyRole538,            // INT
		NestedPartySubID545,           // STRING
		NestedPartySubIDType805,       // INT
		DistribPaymentMethod477,       // INT
		DistribPercentage512,          // PERCENTAGE
		CashDistribCurr478,            // CURRENCY
		CashDistribAgentName498,       // STRING
		CashDistribAgentCode499,       // STRING
		CashDistribAgentAcctNumber500, // STRING
		CashDistribPayRef501,          // STRING
		CashDistribAgentAcctName502,   // STRING
	},
}

const Req, Opt = true, false

const (
	RegistID513                   = 513 // STRING
	RegistTransType514            = 514 // CHAR
	RegistRefID508                = 508 // STRING
	ClOrdID11                     = 11  // STRING
	Account1                      = 1   // STRING
	AcctIDSource660               = 660 // INT
	RegistAcctType493             = 493 // STRING
	TaxAdvantageType495           = 495 // INT
	OwnershipType517              = 517 // CHAR
	PartyID448                    = 448 // STRING
	PartyIDSource447              = 447 // CHAR
	PartyRole452                  = 452 // INT
	PartySubID523                 = 523 // STRING
	PartySubIDType803             = 803 // INT
	RegistDtls509                 = 509 // STRING
	RegistEmail511                = 511 // STRING
	MailingDtls474                = 474 // STRING
	MailingInst482                = 482 // STRING
	OwnerType522                  = 522 // INT
	DateOfBirth486                = 486 // LOCALMKTDATE
	InvestorCountryOfResidence475 = 475 // COUNTRY
	NestedPartyID524              = 524 // STRING
	NestedPartyIDSource525        = 525 // CHAR
	NestedPartyRole538            = 538 // INT
	NestedPartySubID545           = 545 // STRING
	NestedPartySubIDType805       = 805 // INT
	DistribPaymentMethod477       = 477 // INT
	DistribPercentage512          = 512 // PERCENTAGE
	CashDistribCurr478            = 478 // CURRENCY
	CashDistribAgentName498       = 498 // STRING
	CashDistribAgentCode499       = 499 // STRING
	CashDistribAgentAcctNumber500 = 500 // STRING
	CashDistribPayRef501          = 501 // STRING
	CashDistribAgentAcctName502   = 502 // STRING
)
