// Package that parses an XML DMARC report and returns its structure

package parser

import (
	"encoding/xml"
	"fmt"
	"io"
	"strings"
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


var separator string = "-------------------------------------------------------------\n"

func ParseReport(r io.Reader) (*Feedback, error) {
	var report Feedback
	decoder := xml.NewDecoder(r)
	err := decoder.Decode(&report)

	if err != nil {
		return nil, err
	}

	return &report, nil
}

func PrintReportSummary (reports []Feedback) string {
	var sb strings.Builder
	header := "Reports Summary - Press the number on the left to view details \n"

	sb.WriteString(header)
	sb.WriteString(separator)

	for index, report := range reports {
		sb.WriteString(fmt.Sprintf(
			"%d \t %s - %s - %s\n",
			index,
			report.ReportMetaData.OrgName, 
			report.PolicyPublished.Domain,
			report.PolicyPublished.P,
		))
	}

	return sb.String()
}

func PrintFullReport(report *Feedback) (string) {
	var sb strings.Builder
	var header = "Name \t \t \t Value \n"
	
	sb.WriteString(header)
	sb.WriteString(separator)
	sb.WriteString(fmt.Sprintf("Org Name \t \t %s\n", report.ReportMetaData.OrgName))
	sb.WriteString(fmt.Sprintf("Domain \t \t \t %s\n", report.PolicyPublished.Domain))
	sb.WriteString(fmt.Sprintf("P \t \t \t %s\n", report.PolicyPublished.P))
	sb.WriteString(fmt.Sprintf("DKIM \t \t \t %s\n", report.Record.Row.PolicyEvaluated.Dkim))
	sb.WriteString(fmt.Sprintf("SPF \t \t \t %s\n", report.Record.Row.PolicyEvaluated.Spf))
	sb.WriteString(fmt.Sprintf("COUNT \t \t \t %d\n", report.Record.Row.Count))
	sb.WriteString(fmt.Sprintf("SOURCE IP \t \t %s\n", report.Record.Row.SourceIp))

	return sb.String()
}
