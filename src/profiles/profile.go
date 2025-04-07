package profiles

type Profile struct {
	ID       string  `json:"id"`
	Password string  `json:"password"`
	Fullname string  `json:"fullname"`
	Wallet   float64 `json:"wallet"`
	Basket 	 Clothe
}

type Clothe struct {
	Hat int
	Shoes int
	Pant int
	Shirt int
}

var Profiles []Profile

func AddProfile(new Profile) {
	Profiles = append(Profiles, new)
}