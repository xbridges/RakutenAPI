package rakutenapi

type RakutenSearchParameter uint
const (
	// For general use
	ApplicationID RakutenSearchParameter = iota	/* Required */
	AffiliateID				/* Your affiliate id*/
	Format 					/* Responsed format "json" or "xml", default "json" */
	CallBack				/* For jsonp callback function name */
	Elements				/* Elements responsed target columns, default "ALL" */
	FormatVersion			/* "1" or "2", recommended "2", responsed be a simple. */
	Carrier
	Page
	GenreID					/* use Ranking and Search. */
	
	// For Raunking API use
	Age
	Sex
	Period

	// For ItemSearch API use
	Keyword
	ShopCode
	ItemCode
	TagId
	Hits
	Sort
	MinPrice
	MaxPrice
	Availability
	Field
	ImageFlag
	OrFlag
	NGKeyword
	PurchaseType
	ShipOverseasFlag
	ShipOverseasArea
	AsurakuFlag
	AsurakuArea
	PointRateFlag
	PointRate
	PostageFlag
	CreditCardFlag
	GiftFlag
	HasReviewFlag
	MaxAffiliateRate
	MinAffiliateRate
	HasMovieFlag
	PamphletFlag
	AppointDeliveryDateFlag
	GenreInformationFlag
	TagInformationFlag
)


var (
    enumParams = [...]string{
    	// general
    	"applicationId", "affiliateId", "format", 
    	"callback", "elements", "formatVersion", 
    	"carrier", "page", "genreId", 
    	// ranking
    	"age", "sex", "period",
    	// search
    	"keyword", "shopCode", "itemCode", "genreId", 
    	"tagId", "hits", "sort", "minPrice", "maxPrice", 
    	"availability", "field", "imageFlag", "orFlag", 
    	"NGKeyword", "purchaseType", "shipOverseasFlag", 
    	"shipOverseasArea", "asurakuFlag", "asurakuArea", 
    	"pointRateFlag", "pointRate", "postageFlag", 
    	"creditCardFlag", "giftFlag", "hasReviewFlag", 
    	"maxAffiliateRate", "minAffiliateRate", 
    	"hasMovieFlag", "pamphletFlag", 
    	"appointDeliveryDateFlag", "genreInformationFlag", 
    	"tagInformationFlag",
    }
)

func (v RakutenSearchParameter) String() string {
    return enumParams[v]
}
