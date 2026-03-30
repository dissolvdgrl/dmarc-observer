// Package that parses an XML DMARC report and returns its structure

package parser

import (
	"encoding/xml"
	"io"
)

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

type PolicyEvaluated struct {
	Disposition string `xml:"disposition"`
	Dkim string `xml:"dkim"`
	Spf string `xml:"spf"`
}

type Row struct {
	SourceIp string `xml:"source_ip"`
	Count uint8 `xml:"count"`
	PolicyEvaluated PolicyEvaluated `xml:"policy_evaluated"`
}

type Identifiers struct {
	HeaderFrom string `xml:"header_fromt"`
}

type Dkim struct {
	Domain string `xml:"domain"`
	Result string `xml:"pass"`
	Selector string `xml:"selector"`
}

type Spf struct {
	Domain string `xml:"domain"`
	Result string `xml:"result"`
}

type AuthResults struct {
	Dkim Dkim `xml:"dkim"`
	Spf Spf `xml:"spf"`
}

type Record struct {
	Row Row `xml:"row"`
	Identifiers Identifiers `xml:"identifiers"`
}

type Feedback struct {
	Version string `xml:"version"`
	ReportMetaData ReportMetaData `xml:"report_metadata"`
	PolicyPublished PolicyPublished `xml:"policy_published"`
	Record Record `xml:"record"`
}

func ParseReport(r io.Reader) (*Feedback, error) {
	var report Feedback
	decoder := xml.NewDecoder(r)
	error := decoder.Decode(&report)

	if error != nil {
		return nil, error
	}

	return &report, nil
}
