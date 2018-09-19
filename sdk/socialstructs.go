package loginradius

import "time"

// SocialAccessToken contains info from the social access token api
type SocialAccessToken struct {
	AccessToken string    `json:"access_token"`
	ExpiresIn   time.Time `json:"expires_in"`
	IsPosted    bool
}

// SocialAlbum contains info from the social album api
type SocialAlbum []struct {
	ID            string `json:"ID"`
	OwnerID       string `json:"OwnerId"`
	OwnerName     string `json:"OwnerName"`
	Title         string `json:"Title"`
	Description   string `json:"Description"`
	Location      string `json:"Location"`
	Type          string `json:"Type"`
	CreatedDate   string `json:"CreatedDate"`
	UpdatedDate   string `json:"UpdatedDate"`
	CoverImageURL string `json:"CoverImageUrl"`
	ImageCount    string `json:"ImageCount"`
	DirectoryURL  string `json:"DirectoryUrl"`
}

// SocialAudio contains info from the social audio api
type SocialAudio []struct {
	Artist      string `json:"Artist"`
	URL         string `json:"Url"`
	Title       string `json:"Title"`
	UpdatedDate string `json:"UpdatedDate"`
	OwnerName   string `json:"OwnerName"`
	CreatedDate string `json:"CreatedDate"`
	Duration    string `json:"Duration"`
	OwnerID     string `json:"OwnerId"`
	ID          string `json:"ID"`
}

// SocialCheckin contains info from the social checkin api
type SocialCheckin []struct {
	ID          string `json:"ID"`
	CreatedDate string `json:"CreatedDate"`
	OwnerID     string `json:"OwnerId"`
	OwnerName   string `json:"OwnerName"`
	Latitude    string `json:"Latitude"`
	Longitude   string `json:"Longitude"`
	Message     string `json:"Message"`
	PlaceTitle  string `json:"PlaceTitle"`
	Address     string `json:"Address"`
	Distance    string `json:"Distance"`
	Type        string `json:"Type"`
	ImageURL    string `json:"ImageUrl"`
	City        string `json:"City"`
	Country     string `json:"Country"`
}

// SocialCompany contains info from the social company api
type SocialCompany []struct {
	ID   string `json:"ID"`
	Name string `json:"Name"`
}

// SocialContact contains info from the social contact api
type SocialContact struct {
	Data []struct {
		Name        string `json:"Name"`
		EmailID     string `json:"EmailID"`
		PhoneNumber string `json:"PhoneNumber"`
		ID          string `json:"ID"`
		ProfileURL  string `json:"ProfileUrl"`
		ImageURL    string `json:"ImageUrl"`
		Status      string `json:"Status"`
		Industry    string `json:"Industry"`
		Country     string `json:"Country"`
		Location    string `json:"Location"`
		Gender      string `json:"Gender"`
		DateOfBirth string `json:"DateOfBirth"`
	} `json:"Data"`
	NextCursor string `json:"NextCursor"`
}

// SocialEvent contains info from the social event api
type SocialEvent []struct {
	ID          string `json:"ID"`
	OwnerID     string `json:"OwnerId"`
	OwnerName   string `json:"OwnerName"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	RsvpStatus  string `json:"RsvpStatus"`
	Location    string `json:"Location"`
	StartTime   string `json:"StartTime"`
	UpdatedDate string `json:"UpdatedDate"`
	EndTime     string `json:"EndTime"`
	Privacy     string `json:"Privacy"`
}

// SocialFollowing contains info from the social following api
type SocialFollowing []struct {
	Name        string `json:"Name"`
	EmailID     string `json:"EmailID"`
	PhoneNumber string `json:"PhoneNumber"`
	ID          string `json:"ID"`
	ProfileURL  string `json:"ProfileUrl"`
	ImageURL    string `json:"ImageUrl"`
	Status      string `json:"Status"`
	Industry    string `json:"Industry"`
	Country     string `json:"Country"`
	Location    string `json:"Location"`
	Gender      string `json:"Gender"`
	DateOfBirth string `json:"DateOfBirth"`
}

// SocialGroup contains info from the social group api
type SocialGroup []struct {
	ID          string `json:"ID"`
	Name        string `json:"Name"`
	Email       string `json:"Email"`
	Description string `json:"Description"`
	Type        string `json:"Type"`
	Country     string `json:"Country"`
	PostalCode  string `json:"PostalCode"`
	Logo        string `json:"Logo"`
	Image       string `json:"Image"`
	MemberCount int    `json:"MemberCount"`
}

// SocialLike contains info from the social like api
type SocialLike []struct {
	ID          string `json:"ID"`
	Name        string `json:"Name"`
	Category    string `json:"Category"`
	CreatedDate string `json:"CreatedDate"`
	Website     string `json:"Website"`
	Description string `json:"Description"`
}

// SocialMention contains info from the social mention api
type SocialMention []struct {
	ID       string `json:"Id"`
	Text     string `json:"Text"`
	DateTime string `json:"DateTime"`
	Likes    int    `json:"Likes"`
	Place    string `json:"Place"`
	Source   string `json:"Source"`
	ImageURL string `json:"ImageUrl"`
	LinkURL  string `json:"LinkUrl"`
	Name     string `json:"Name"`
}

// SocialMessageAPI contains info from the social message api
type SocialMessageAPI struct {
	IsPosted bool   `json:"isPosted"`
	Data     string `json:"Data"`
}

// SocialPage contains info from the social page api
type SocialPage struct {
	ID                string    `json:"ID"`
	Name              string    `json:"Name"`
	URL               string    `json:"Url"`
	Category          string    `json:"Category"`
	Likes             string    `json:"Likes"`
	Phone             string    `json:"Phone"`
	Image             string    `json:"Image"`
	Website           string    `json:"Website"`
	About             string    `json:"About"`
	Description       string    `json:"Description"`
	Awards            string    `json:"Awards"`
	CheckinCount      string    `json:"CheckinCount"`
	Founded           string    `json:"Founded"`
	Mission           string    `json:"Mission"`
	Products          string    `json:"Products"`
	ReleaseDate       time.Time `json:"ReleaseDate"`
	TalkingAboutCount string    `json:"TalkingAboutCount"`
	Published         bool      `json:"Published"`
	UserName          string    `json:"UserName"`
	Locations         string    `json:"Locations"`
	CategoryList      []struct {
		ID   string `json:"Id"`
		Name string `json:"Name"`
	} `json:"CategoryList"`
	CoverImage struct {
		ID      string `json:"Id"`
		Source  string `json:"Source"`
		OffsetY string `json:"OffsetY"`
		OffsetX string `json:"OffsetX"`
	} `json:"CoverImage"`
	EmployeeCountRange string `json:"EmployeeCountRange"`
	Industries         string `json:"Industries"`
	Specialties        string `json:"Specialties"`
	Status             string `json:"Status"`
	StockExchange      string `json:"StockExchange"`
}

// SocialPhoto contains info from the social photo api
type SocialPhoto []struct {
	ID          string `json:"ID"`
	AlbumID     string `json:"AlbumId"`
	OwnerID     string `json:"OwnerId"`
	OwnerName   string `json:"OwnerName"`
	Name        string `json:"Name"`
	DirectURL   string `json:"DirectUrl"`
	ImageURL    string `json:"ImageUrl"`
	Location    string `json:"Location"`
	Link        string `json:"Link"`
	Description string `json:"Description"`
	Height      string `json:"Height"`
	Width       string `json:"Width"`
	CreatedDate string `json:"CreatedDate"`
	UpdatedDate string `json:"UpdatedDate"`
	Images      []struct {
		Dimensions string `json:"Dimensions"`
		Image      string `json:"Image"`
	} `json:"Images"`
}

// SocialPost contains info from the social post api
type SocialPost []struct {
	ID         string `json:"ID"`
	Name       string `json:"Name"`
	Title      string `json:"Title"`
	StartTime  string `json:"StartTime"`
	UpdateTime string `json:"UpdateTime"`
	Message    string `json:"Message"`
	Place      string `json:"Place"`
	Picture    string `json:"Picture"`
	Likes      int    `json:"Likes"`
	Share      int    `json:"Share"`
	Type       string `json:"Type"`
}

// SocialStatus contains info from the social status api
type SocialStatus []struct {
	IsPosted bool   `json:"isPosted"`
	Data     string `json:"data"`
	ID       string `json:"Id"`
	Text     string `json:"Text"`
	DateTime string `json:"DateTime"`
	Likes    int    `json:"Likes"`
	Place    string `json:"Place"`
	Source   string `json:"Source"`
	ImageURL string `json:"ImageUrl"`
	LinkURL  string `json:"LinkUrl"`
	Name     string `json:"Name"`
}

// SocialVideo contains info from the social video api
type SocialVideo struct {
	Data []struct {
		ID          string `json:"ID"`
		Description string `json:"Description"`
		Name        string `json:"Name"`
		Image       string `json:"Image"`
		Source      string `json:"Source"`
		CreatedDate string `json:"CreatedDate"`
		OwnerID     string `json:"OwnerId"`
		OwnerName   string `json:"OwnerName"`
		EmbedHTML   string `json:"EmbedHtml"`
		UpdatedDate string `json:"UpdatedDate"`
		Duration    string `json:"Duration"`
		DirectLink  string `json:"DirectLink"`
	} `json:"Data"`
	NextCursor string `json:"NextCursor"`
}
