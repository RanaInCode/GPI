package insights

type Insight struct {
	Title      string   `json:"title"`
	Summary    string   `json:"summary"`
	Evidence   []string `json: "evidence"`
    Confidence float64  `json:"confidence"`
	Category   string   `json:"category"`

}