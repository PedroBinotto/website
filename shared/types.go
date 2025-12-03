package shared

import "time"

type Tag struct {
	Url         string `json:"url"`
	DisplayName string `json:"display_name"`
}

type BlogPayload struct {
	ID        int32
	Title     string
	Synopsis  string
	Url       string
	Body      string
	CreatedAt time.Time
	Tags      []Tag
}
