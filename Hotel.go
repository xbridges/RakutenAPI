package rakutenapi

type Hotel struct {
	HotelBasicInfo	struct {
		HotelNo 			int			`json:"hotelNo"`
		HotelName 			string		`json:"hotelName"`
		HotelInformationUrl string		`json:"hotelInformationUrl"`
		PlanListUrl 		string		`json:"planListUrl"`
		DpPlanListUrl 		string		`json:"dpPlanListUrl"`
		ReviewUrl 			string		`json:"reviewUrl"`
		HotelKanaName 		string		`json:"hotelKanaName"`
		HotelSpecial 		string		`json:"hotelSpecial"`
		HotelMinCharge 		int			`json:"hotelMinCharge"`
		Latitude 			float64		`json:"latitude"`
		Longitude 			float64		`json:"longitude"`
		PostalCode 			string		`json:"postalCode"`
		Address1 			string		`json:"address1"`
		Address2 			string		`json:"address2"`
		TelephoneNo 		string		`json:"telephoneNo"`
		FaxNo 				string		`json:"faxNo"`
		Access 				string		`json:"access"`
		ParkingInformation 	string		`json:"parkingInformation"`
		NearestStation 		string		`json:"nearestStation"`
		HotelImageUrl 		string		`json:"hotelImageUrl"`
		HotelThumbnailUrl 	string		`json:"hotelThumbnailUrl"`
		RoomImageUrl 		string		`json:"roomImageUrl"`
		RoomThumbnailUrl 	string		`json:"roomThumbnailUrl"`
		HotelMapImageUrl 	string		`json:"hotelMapImageUrl"`
		ReviewCount 		int			`json:"reviewCount"`
		ReviewAverage  		float64		`json:"reviewAverage"`
		UserReview 			string		`json:"userReview"`
	}	`json:"hotelBasicInfo"`
	HotelRatingInfo	struct {
		ServiceAverage 		float64		`json:"serviceAverage"`
		LocationAverage 	float64		`json:"locationAverage"`
		RoomAverage 		float64		`json:"roomAverage"`
		EquipmentAverage 	float64		`json:"equipmentAverage"`
		BathAverage 		float64		`json:"bathAverage"`
		MealAverage 		float64		`json:"mealAverage"`
	}	`json:"hotelRatingInfo"`
}