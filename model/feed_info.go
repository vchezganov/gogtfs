//Done
package model

import "time"

type FeedInfo struct {
	FeedPublisherName string    `csv:"feed_publisher_name"`
	FeedPublisherURL  string    `csv:"feed_publisher_url"`
	FeedLang          string    `csv:"feed_lang"`
	DefaultLang       string    `csv:"default_lang"`
	FeedStartDate     time.Time `csv:"feed_start_date,ParseFeedStartDate"`
	FeedEndDate       time.Time `csv:"feed_end_date,ParseFeedEndDate"`
	FeedVersion       string    `csv:"feed_version"`
	FeedContactEmail  string    `csv:"feed_contact_email"`
	FeedContactURL    string    `csv:"feed_contact_url"`
}

func (f FeedInfo) TableFileName() string {
	return "feed_info.txt"
}

func (f *FeedInfo) ParseFeedStartDate(value string) (err error) {
	f.FeedStartDate, err = parseDate(value)
	return
}

func (f *FeedInfo) ParseFeedEndDate(value string) (err error) {
	f.FeedEndDate, err = parseDate(value)
	return
}
