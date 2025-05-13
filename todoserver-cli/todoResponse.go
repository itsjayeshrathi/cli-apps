package main

import (
	"encoding/json"
	"time"
)

type todoResponse struct {
	Results List `json:"results"`
}

func (r *todoResponse) MarshalJSON() ([]byte, error) {
	resp := struct {
		Results      List  `json:"results"`
		Date         int64 `json:"date"`
		TotalResults int   `json:"total_results"`
	}{
		Results:      r.Results,
		Date:         time.Now().Unix(),
		TotalResults: len(r.Results),
	}
	return json.Marshal(resp)
}
