package models

import "time"

type Events struct {
	UserId    string
	TimeStamp time.Time
	Event     string
}

type Ingest struct {
	Path string `json:"path"`
}

type Filters struct {
	UserId   string
	DateFrom string
	DateTo   string
	Event    string
}

type CountResponse struct {
	Count string `json:"count"`
}

type CountDistinctUsersResponse struct {
	CountDistinctUsers string `json:"count_distinct_users"`
}

type ExistsResponse struct {
	Exists string `json:"exists"`
}
