# RakutenAPI
for Rakuten affiliate easy search.
```
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
```
