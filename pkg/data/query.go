package data

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Types for all Query.
const (
	TypeListen = iota
	TypeUnlisten
	TypeSet
	TypeUpdate
	TypeRemove
)

// Query defines the constraints for retrieving data.
type Query struct {
	// Type defines the query type.
	Type int

	// Ref defines the reference path.
	Ref string
	// RequestID defines the id of client request.
	RequestID int64

	// Data defines the update payload if type is Set or Update.
	Data interface{}

	// QueryID defines the id of query if type is Listen.
	QueryID int64
	// StartAt defines the value of start, should be with OrderBy.
	StartAt interface{}
	// StartKey defines the key of start.
	StartKey string
	// EndAt defines the value of end, should be with OrderBy.
	EndAt interface{}
	// EndKey defines the key of end.
	EndKey string
	// OrderBy defines the field to be ordered, can be ".key", ".value" or child field name.
	OrderBy string
	// Limit defines the query limit.
	Limit int
	// LimitOrder defines the query limit asc "l" or desc "r".
	LimitOrder string
	// Shallow defines the flag to returns shallowed keys.
	Shallow bool
}

// r defines the internal data schema of Query.
type r struct {
	T string `json:"t"`
	D *struct {
		// A indicates the request type.
		A string `json:"a"`
		// R indicates the request ID.
		R int64 `json:"r"`
		// B indicates the body content for the request.
		B *struct {
			// P indicates the Firebase reference path.
			P string `json:"p"`
			// D indicates the data payload to be written.
			D interface{} `json:"d"`
			// T indicates the query ID.
			T int64 `json:"t"`
			// Q indicates the query to retrieve data.
			Q *struct {
				// SP indicates "start at" query.
				SP interface{} `json:"sp"`
				// SN indicates "start key" query
				SN string `json:"sn"`
				// EP indicates "end at" query.
				EP interface{} `json:"ep"`
				// EN indicates "end key" query.
				EN string `json:"en"`
				// I indicates "order by" query
				I string `json:"i"`
				// L indicates "limit".
				L int `json:"l"`
				// VF indicates order, can be either "l" or "r".
				VF string `json:"vf"`
			} `json:"q"`
		} `json:"b"`
	} `json:"d"`
}

// UnmarshalJSON overrides the json.Unmarshal for Query struct.
func (q *Query) UnmarshalJSON(bytes []byte) error {
	// unmarshal the bytes into internal model.
	var r *r
	if err := json.Unmarshal(bytes, &r); err != nil {
		return err
	}

	// validate the internal model.
	if r == nil {
		return errors.New("missing request")
	} else if r.T != "d" {
		return fmt.Errorf("invalid r.t: %s", r.T)
	} else if r.D == nil {
		return errors.New("missing r.d")
	} else if r.D.B == nil {
		return errors.New("missing r.d.b")
	}

	// check the query type.
	switch r.D.A {
	case "l", "q":
		q.Type = TypeListen
	case "n":
		q.Type = TypeUnlisten
	case "m":
		q.Type = TypeUpdate
	case "p":
		q.Type = TypeSet
	default:
		return fmt.Errorf("invalid type: %s", r.D.A)
	}

	// convert the internal model to Query.
	q.RequestID = r.D.R
	q.Ref = r.D.B.P
	q.Data = r.D.B.D
	q.QueryID = r.D.B.T

	// convert query parameters.
	if r.D.B.Q != nil {
		q.StartAt = r.D.B.Q.SP
		q.StartKey = r.D.B.Q.SN
		q.EndAt = r.D.B.Q.EP
		q.EndKey = r.D.B.Q.EN
		q.OrderBy = r.D.B.Q.I
		q.Limit = r.D.B.Q.L
		q.LimitOrder = r.D.B.Q.VF
	}

	return nil
}
