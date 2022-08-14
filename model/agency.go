//Done
package model

type Agency struct {
	AgencyID       string `csv:"agency_id"`
	AgencyName     string `csv:"agency_name"`
	AgencyURL      string `csv:"agency_url"`
	AgencyTimezone string `csv:"agency_timezone"`
	AgencyLang     string `csv:"agency_lang"`
	AgencyPhone    string `csv:"agency_phone"`
	AgencyFareURL  string `csv:"agency_fare_url"`
	AgencyEmail    string `csv:"agency_email"`
}

func (a Agency) TableFileName() string {
	return "agency.txt"
}
