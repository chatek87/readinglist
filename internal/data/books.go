package data

import "time"

type Book struct {
	ID        int64     `json:"id"` // this json tag changes the name of the field from 'ID' to 'id' when it's marshalled to json
	CreatedAt time.Time `json:"-"`  // the '-' prevents created time from being displayed, hides it
	Title     string    `json:"title"`
	Published int       `json:"published,omitempty"`
	Pages     int       `json:"pages,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Rating    float32   `json:"rating,omitempty"`
	Version   int32     `json:"-"`
}
