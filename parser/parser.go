package parser

type DateRange struct {
	Begin int64 `xml:"begin"`
	End int64 `xml"end"`
}

type ReportMetaData struct {
	OrgName string `xml:"org_name"`
	Email string `xml:"email"`
	ExtraContactInfo string `xml:"extra_contact_info"`
	ReportId string `xml:"report_id"`
	DateRange DateRange `xml:"date_range"`
}

type PolicyPublished struct {
	Domain string `xml:"domain"`
	Adkim string `xml:"adkim"` // DKIM alignment mode: r=relaxed s=strict
	Aspf string `xml:"aspf"` // SPF alignment mode: same values as above
	P string `xml:"p"` // policy: none, quarantine, reject
	Sp string `xml:"sp"` // policy for subdomains: same as above, fallback to p if absent
	Pct uint8  `xml:"pct"` // % of mail policy applies to
	Np string `xml:"np"` // policy for non-existent subdomains: same values as above	
}


type Feedback struct {
	Version string `xml:"version"`
	ReportMetaData ReportMetaData `xml:"report_metadata"`
	PolicyPublished PolicyPublished `xml:"policy_published"`
}
