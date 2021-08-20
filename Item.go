package rakutenapi

// FormatVersion 2専用オブジェクト
type Item struct {
	AffiliateUrl		string		`json:"affiliateUrl"`
	AffiliateRate		float64		`json:"affiliateRate"`
	AsurakuArea			string		`json:"asurakuArea"`
	AsurakuFlag			int			`json:"asurakuFlag"`
	AsurakuClosingTime	string		`json:"asurakuClosingTime"`
	Availability		int			`json:"availability"`
	CreditCardFlag		int			`json:"creditCardFlag"`
	Catchcopy			string		`json:"catchcopy"`
	Carrier				int			`json:"carrier"`
	GenreId				int			`json:"genreId"`
	EndTime				string		`json:"endTime"`
	MediumImageUrls		[]string	`json:"mediumImageUrls"`
	ItemCaption			string 		`json:"itemCaption"`
	ImageFlag			int			`json:"imageFlag"`
	ItemCode			string		`json:"itemCode"`
	ItemName			string		`json:"itemName"`
	ItemPrice			int			`json:"itemPrice"`
	ItemUrl				string		`json:"itemUrl"`
	PointRate			int			`json:"pointRate"`
	PointRateEndTime	string		`json:"pointRateEndTime"`
	PointRateStartTime	string		`json:"pointRateStartTime"`
	PostageFlag			int			`json:"postageFlag"`
	Rank				int			`json:"rank"`
	ReviewAverage		float64		`json:"reviewAverage"`
	ReviewCount			int			`json:"reviewCount"`
	SmallImageUrls		[]string	`json:"imageUrl"`
	ShipOverseasArea	string		`json:"shipOverseasArea"`
	ShipOverseasFlag	int			`json:"shipOverseasFlag"`
	ShopAffiliateUrl	string		`json:"shopAffiliateUrl"`
	ShopCode			string		`json:"shopCode"`
	ShopName			string		`json:"shopName"`
	ShopOfTheYearFlag	int			`json:"shopOfTheYearFlag"`
	ShopUrl				string		`json:"shopUrl"`
	StartTime			string		`json:"startTime"`
	TaxFlag				int			`json:"taxFlagmediumImageUrls"`
	TagIds				[]int		`json:"tagIds"`
}