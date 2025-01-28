package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	CalendarId    string             `bson:"calendar_id"`
	Date          string             `bson:"date"`
	Country       string             `bson:"country"`
	Category      string             `bson:"category"`
	Event         string             `bson:"event"`
	Reference     string             `bson:"reference"`
	ReferenceDate string             `bson:"reference_date"`
	Source        string             `bson:"source"`
	SourceURL     string             `bson:"source_url"`
	Actual        string             `bson:"actual"`
	Previous      string             `bson:"previous"`
	Forecast      string             `bson:"forecast"`
	TEForecast    string             `bson:"te_forecast"`
	URL           string             `bson:"url"`
	DateSpan      string             `bson:"date_span"`
	Importance    int                `bson:"importance"`
	LastUpdate    string             `bson:"last_update"`
	Revised       string             `bson:"revised"`
	Currency      string             `bson:"currency"`
	Unit          string             `bson:"unit"`
	Ticker        string             `bson:"ticker"`
	Symbol        string             `bson:"symbol"`
}

type UserEvent struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Email      string             `bson:"email"`
	CalendarId string             `bson:"calendar_id"`
	Importance int                `bson:"importance"`
	Event      string             `bson:"event"`
	Date       string             `bson:"date"`
}
