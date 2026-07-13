package models

type ReviewRequest struct {
	Language string `json:"language"`

	Code string `json:"code"`
}

type Issue struct {
	Severity string `json:"severity"`
	Line     int    `json:"line"`
	Message  string `json:"message"`
}

type ReviewResponse struct {
	Score       int      `json:"score"`
	Grade       string   `json:"grade"`
	Summary     string   `json:"summary"`
	Issues      []Issue  `json:"issues"`
	Suggestions []string `json:"suggestions"`
}
