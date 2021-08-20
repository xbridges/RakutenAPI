package rakutenapi

type ItemSearchResponse struct {
	Count				int				`json:"count"`
	Page				int				`json:"page"`
	First				int				`json:"first"`
	Last				int				`json:"last"`
	Hits				int				`json:"hits"`
	Carrier				int				`json:"carrier"`
	PageCount			int				`json:"pageCount"`
	Items				[]Item			`json:"Items"`
	GenreInformation	[]interface{}	`json:"GenreInformation"`
	TagInformation		[]interface{}	`json:"TagInformation"`
}