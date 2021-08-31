package rakutenapi

type TravelKeywordSearchResponse struct {
	PagingInfo			struct {
							RecordCount int `json:"recordCount"`
							PageCount   int `json:"pageCount"`
							Page  		int `json:"page"`
							First  		int `json:"first"`
							Last  		int `json:"last"`
						}				`json:"pagingInfo"`
	Hotels				[][]Hotel			`json:"hotels"`
}