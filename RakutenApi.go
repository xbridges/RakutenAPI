package rakutenapi

import(
	"fmt"
	"errors"
	"net/http"
	"net/url"
	"io/ioutil" 
)

const (
	ItemRankingURL   = "https://app.rakuten.co.jp/services/api/IchibaItem/Ranking/20170628"
	ItemSearchURL    = "https://app.rakuten.co.jp/services/api/IchibaItem/Search/20170706"
	TravelKeyWordURL = "https://app.rakuten.co.jp/services/api/Travel/KeywordHotelSearch/20170426"
)

const (
	SearchDisabled int = iota
	SearchEnabled
)

type RakutenAPI struct{
	RakutenAffiliateAccount
}

func NewRakutenAPI(account RakutenAffiliateAccount) (*RakutenAPI, error) {
	if len(account.ApplicationID) == 0 {
		return nil, errors.New("Need your application id.")
	} 
	return &RakutenAPI{RakutenAffiliateAccount: account}, nil
}

func accessRakutenAPI(uri string) ([]byte, error){

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rakutenResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return rakutenResp, nil
}

func buildQuery(api *RakutenAPI, params map[string]string) string{

	// URL = https://app.rakuten.co.jp/services/api/IchibaItem/Ranking/20170628 
	// ?[parameter1]=[value1]&[parameter2]=[value2]…
	// A value must be encoded to URL by UTF8.
	// ApplicationID parameter required.
	query := fmt.Sprintf("%s=%s", ApplicationID.String(), url.QueryEscape(api.ApplicationID))
	if len(api.AffiliateID) != 0 {
		query = fmt.Sprintf("%s&%s=%s", query, AffiliateID.String(), url.QueryEscape(api.AffiliateID))
	}
	if _, ok := params[FormatVersion.String()]; !ok {
		query = fmt.Sprintf("%s&%s=%s", query, FormatVersion.String(), url.QueryEscape("2"))	
	}
	for p, v := range params {
		if p == ApplicationID.String() || p == AffiliateID.String() {
			continue
		}
		query = fmt.Sprintf("%s&%s=%s", query, p, url.QueryEscape(v))
	}
	return query
}

func (api *RakutenAPI) GetRaktenRunking(params map[string]string) ([]byte, error) {

	if api.Status != SearchEnabled {
		return nil, errors.New("Search disabled")
	}
	query := buildQuery(api, params)
	uri := fmt.Sprintf("%s?%s", ItemRankingURL, query)
	// fmt.Printf("%s\n", uri)

	return accessRakutenAPI(uri)
}

func (api *RakutenAPI) GetItemSearch(params map[string]string) ([]byte, error) {

	if api.Status != SearchEnabled {
		return nil, errors.New("Search disabled")
	}
	query := buildQuery(api, params)
	uri := fmt.Sprintf("%s?%s", ItemSearchURL, query)
	// fmt.Printf("%s\n", uri)

	return accessRakutenAPI(uri)
}

func (api *RakutenAPI) GetTravelKeywordSearch(params map[string]string) ([]byte, error) {

	if api.Status != SearchEnabled {
		return nil, errors.New("TravelSearch disabled")
	}
	query := buildQuery(api, params)
	uri := fmt.Sprintf("%s?%s", TravelKeyWordURL, query)
	fmt.Printf("%s\n", uri)

	return accessRakutenAPI(uri)
}
