package piiredactionprocessor

type Entity struct {
	Entity string  `json:"entity"`
	Score  float64 `json:"score"`
	Index  int     `json:"index"`
	Word   string  `json:"word"`
	Start  int     `json:"start"`
	End    int     `json:"end"`
}
