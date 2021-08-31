# RakutenAPI
for Rakuten affiliate easy search.
```
	a := RakutenAffiliateAccount{
		ApplicationID: "YOUR APPLICATION ID",
		Status: SearchEnabled,
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

	params = map[string]string{}
	params[Format.String()] = "json"
	params[Page.String()] = "1"
	params[Keyword.String()] = "東京"

	resp, err = api.GetTravelKeywordSearch(params)
	if err != nil {
		t.Fatal(err)
	}
	reptrv := TravelKeywordSearchResponse{}
	json.Unmarshal(resp, &reptrv)

	fmt.Printf("\n\n\n%+v\n", reptrv)	
```
