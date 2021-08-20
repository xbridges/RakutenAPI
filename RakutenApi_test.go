package rakutenapi

import(
	"testing"
	"fmt"
	"encoding/json"
)

func TestRakutenApi( t *testing.T ){
	
	a := RakutenAffiliateAccount{
		ApplicationID: "YOUR APPLICATION ID",
		AffiliateID: "YOUR AFFILIATE ID",
	}
	api, _ := NewRakutenAPI(a)

	params := map[string]string{}
	params[Format.String()] = "json"
	params[Page.String()] = "1"

	resp, err := api.GetRaktenRunking(params)
	if err != nil {
		t.Fatal(err)
	}
	repr := RankingResponse{}
	json.Unmarshal(resp, &repr)

	fmt.Printf("\n\n\n%+v\n", repr)	

	params = map[string]string{}
	params[Format.String()] = "json"
	params[Page.String()] = "1"
	params[Keyword.String()] = "PS5"

	resp, err = api.GetItemSearch(params)
	if err != nil {
		t.Fatal(err)
	}
	rep := ItemSearchResponse{}
	json.Unmarshal(resp, &rep)

	fmt.Printf("\n\n\n%+v\n", rep)	
}