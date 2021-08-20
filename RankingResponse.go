package rakutenapi

type RankingResponse struct {
	Title				string			`json:"title"`
	LastBuildDate		string			`json:"lastBuildDate"`
	Items				[]Item			`json:"Items"`
}
