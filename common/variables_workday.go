package common

import (
	"fmt"
	"os"
)

// WorkDay companies list
var (
	ASML                    = "ASML"
	CrowdStrike             = "CRWD"
	Intel                   = "INTL"
	Nvidia                  = "NVDA"
	Qualys                  = "QLYS"
	SalesForce              = "SLSF"
	Walmart                 = "WLMT"
	Target                  = "TRGT"
	Samsung                 = "SMSN"
	Disney                  = "DSNY"
	Sony                    = "SONY"
	Twitter                 = "TWTR"
	CapitalOne              = "CONE"
	Boeing                  = "BOEG"
	Bose                    = "BOSE"
	Snapchat                = "SNPT"
	CVS                     = "CVSS"
	CCCIntelligentSolutions = "CCCI"
	NorthropGrumman         = "NTGN"
	Phinia                  = "PHNA"
	Nissan                  = "NISN"
	HP                      = "HPHP"
	Barclays                = "BARC"
	Blueorigin              = "BLON"
	Medtronic               = "MDTC"
	Tancent                 = "TNCT"
	Mantech                 = "MTCH"
	SPGlobal                = "SPGL"
	Adobe                   = "ADBE"
	Comcast                 = "CMCT"
	Redhat                  = "REDH"
	Symbotic                = "SYMC"
) //////////////// Edit here

var WorkdayCompanies = make(map[string]bool)

func checkAndInitWorkdayCompaniesList() {
	values := []string{
		ASML,
		CrowdStrike,
		Intel,
		Nvidia,
		Qualys,
		SalesForce,
		Walmart,
		Target,
		Samsung,
		Disney,
		Sony,
		Twitter,
		CapitalOne,
		Boeing,
		Bose,
		Snapchat,
		CVS,
		CCCIntelligentSolutions,
		NorthropGrumman,
		Phinia,
		Nissan,
		HP,
		Barclays,
		Blueorigin,
		Medtronic,
		Tancent,
		Mantech,
		SPGlobal,
		Adobe,
		Comcast,
		Redhat,
		Symbotic,
	} /////////////////// Edit here

	for _, value := range values {
		if WorkdayCompanies[value] || AllCompanies[value] {
			fmt.Printf("Duplicate company code found: %s\n", value)
			os.Exit(1)
		} else {
			WorkdayCompanies[value] = true
			AllCompanies[value] = true
		}
	}
}
