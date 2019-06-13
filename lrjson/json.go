// Package lrjson contains functions for unmarshalling JSON responses returned by RESTful API
package lrjson

import (
	"encoding/json"
	"log"
	"time"
)

// DynamicUnmarshal takes stringified json and unmarshals it into a map with string keys and
// interface{} values. This is recommended for parsing LoginRadius API responses over individual
// functions unmarshalling JSON objects into pre-written structs.

// Golang's strict typing means the latter solution is a better practice in theory, but the former
// solution is recommended for usage with LoginRadius api end points so as to ensure the long-term
// integrity of response handling - i.e. unmarshalling into a prewritten type will throw an error in // case of unexpected field data types, and will quietly do the best it can when there is a mismatch // between incoming JSON object and the destination struct. The former results in unwanted fragility, // and the latter results in data being potentially mis-captured.

// (Note: another alternative is unmarshalling the response as it is read from the body:

// decoder := json.NewDecoder(res.Body)
// decoder.DisallowUnknownFields()
// return decoder.Decode(destinationStruct)

// DisallowUnknownFields was made available in Go 1.10 and will causes the Decoder to return an
// error when the destination is a struct and the input contains object keys which do not match any
// non-ignored, exported fields in the destination. Though this would solve the issue of data not being captured, it would increase the fragility of the API client during new endpoint updates/releases.
//)

// In addition to these reasonings, benchmarks between the two alternatives have been provided in
// benchmark-test.go in this package. The performance difference between the dynamic unmarshalling
// solution and the struct specific unmarshalling solution is about 70000 ns/op
// - insignificant enough to disregard, especially considering the amount of work required to maintain
// pre-written structs and response-specific code.

func DynamicUnmarshal(data string) (map[string]interface{}, error) {
	var unmarshalled = make(map[string]interface{})
	err := json.Unmarshal([]byte(data), &unmarshalled)
	if err != nil {
		log.Println(err)
	}
	return unmarshalled, nil
}

// A sample function for unmarshalling JSON response into pre-written struct, preserved for
// benchmarking purposes
func UnmarshalGetManageAccountProfilesByEmail(data string) (AuthProfile, error) {
	authProfile := AuthProfile{}

	error := json.Unmarshal([]byte(data), &authProfile)
	if error != nil {
		return authProfile, error
	}
	return authProfile, nil
}

// AuthProfile is a struct meant to hold the LoginRadius Profile data mapping for JSON
type AuthProfile struct {
	PasswordHash            string            `json:"PasswordHash"`
	VerificationToken       string            `json:"VerificationToken"`
	ForgotToken             string            `json:"ForgotToken"`
	IsDeleteRequestAccepted bool              `json:"IsDeleteRequestAccepted"`
	Identities              []json.RawMessage `json:"Identities"`
	UserName                string            `json:"UserName"`
	UID                     string            `json:"Uid"`
	PasswordExpirationDate  time.Time         `json:"PasswordExpirationDate"`
	Password                string            `json:"Password"`
	LastPasswordChangeDate  time.Time         `json:"LastPasswordChangeDate"`
	PreviousUids            []string          `json:"PreviousUids"`
	CustomFields            json.RawMessage   `json:"CustomFields"`
	LastPasswordChangeToken string            `json:"LastPasswordChangeToken"`
	PhoneID                 string            `json:"phoneId"`
	ExternalUserLoginID     string            `json:"ExternalUserLoginId"`
	RegistrationProvider    string            `json:"RegistrationProvider"`
	ID                      string            `json:"ID"`
	Provider                string            `json:"Provider"`
	Prefix                  string            `json:"Prefix"`
	FirstName               string            `json:"FirstName"`
	MiddleName              string            `json:"MiddleName"`
	LastName                string            `json:"LastName"`
	Suffix                  string            `json:"Suffix"`
	FullName                string            `json:"FullName"`
	NickName                string            `json:"NickName"`
	ProfileName             string            `json:"ProfileName"`
	BirthDate               string            `json:"BirthDate"`
	Gender                  string            `json:"Gender"`
	Website                 string            `json:"Website"`
	UnverifiedEmail         []struct {
		Type  string `json:"Type"`
		Value string `json:"Value"`
	} `json:"UnverifiedEmail"`
	Email []struct {
		Type  string `json:"Type"`
		Value string `json:"Value"`
	} `json:"Email"`
	Country struct {
		Code string `json:"Code"`
		Name string `json:"Name"`
	} `json:"Country"`
	ThumbnailImageURL string    `json:"ThumbnailImageUrl"`
	ImageURL          string    `json:"ImageUrl"`
	Favicon           string    `json:"Favicon"`
	ProfileURL        string    `json:"ProfileUrl"`
	HomeTown          string    `json:"HomeTown"`
	State             string    `json:"State"`
	City              string    `json:"City"`
	Industry          string    `json:"Industry"`
	About             string    `json:"About"`
	TimeZone          string    `json:"TimeZone"`
	LocalLanguage     string    `json:"LocalLanguage"`
	CoverPhoto        string    `json:"CoverPhoto"`
	TagLine           string    `json:"TagLine"`
	Language          string    `json:"Language"`
	Verified          string    `json:"Verified"`
	UpdatedTime       time.Time `json:"UpdatedTime"`
	Positions         []struct {
		Position  string    `json:"Position"`
		Summary   string    `json:"Summary"`
		StartDate time.Time `json:"StartDate"`
		EndDate   time.Time `json:"EndDate"`
		IsCurrent string    `json:"IsCurrent"`
		Location  string    `json:"Location"`
		Comapny   struct {
			Name     string `json:"Name"`
			Type     string `json:"Type"`
			Industry string `json:"Industry"`
		} `json:"Comapny"`
		Company struct {
			Name     string `json:"Name"`
			Type     string `json:"Type"`
			Industry string `json:"Industry"`
		} `json:"Company"`
	} `json:"Positions"`
	Educations []struct {
		School       string    `json:"School"`
		Year         string    `json:"year"`
		Type         string    `json:"type"`
		Notes        string    `json:"notes"`
		Activities   string    `json:"activities"`
		Degree       string    `json:"degree"`
		Fieldofstudy string    `json:"fieldofstudy"`
		StartDate    time.Time `json:"StartDate"`
		EndDate      time.Time `json:"EndDate"`
	} `json:"Educations"`
	PhoneNumbers []struct {
		PhoneType   string `json:"PhoneType"`
		PhoneNumber string `json:"PhoneNumber"`
	} `json:"PhoneNumbers"`
	IMAccounts []struct {
		AccountType string `json:"AccountType"`
		AccountName string `json:"AccountName"`
	} `json:"IMAccounts"`
	Addresses []struct {
		Type       string `json:"Type"`
		Address1   string `json:"Address1"`
		Address2   string `json:"Address2"`
		City       string `json:"City"`
		State      string `json:"State"`
		PostalCode string `json:"PostalCode"`
		Region     string `json:"Region"`
		Country    string `json:"Country"`
	} `json:"Addresses"`
	MainAddress        string    `json:"MainAddress"`
	Created            string    `json:"Created"`
	CreatedDate        time.Time `json:"CreatedDate"`
	ModifiedDate       time.Time `json:"ModifiedDate"`
	LocalCity          string    `json:"LocalCity"`
	ProfileCity        string    `json:"ProfileCity"`
	LocalCountry       string    `json:"LocalCountry"`
	ProfileCountry     string    `json:"ProfileCountry"`
	IsProtected        bool      `json:"IsProtected"`
	RelationshipStatus string    `json:"RelationshipStatus"`
	Quota              string    `json:"Quota"`
	Quote              string    `json:"Quote"`
	InterestedIn       []string  `json:"InterestedIn"`
	Interests          []struct {
		InterestedType string `json:"InterestedType"`
		InterestedName string `json:"InterestedName"`
	} `json:"Interests"`
	Religion  string `json:"Religion"`
	Political string `json:"Political"`
	Sports    []struct {
		ID   string `json:"Id"`
		Name string `json:"Name"`
	} `json:"Sports"`
	InspirationalPeople []struct {
		ID   string `json:"Id"`
		Name string `json:"Name"`
	} `json:"InspirationalPeople"`
	HTTPSImageURL      string `json:"HttpsImageUrl"`
	FollowersCount     int    `json:"FollowersCount"`
	FriendsCount       int    `json:"FriendsCount"`
	IsGeoEnabled       string `json:"IsGeoEnabled"`
	TotalStatusesCount int    `json:"TotalStatusesCount"`
	Associations       string `json:"Associations"`
	NumRecommenders    int    `json:"NumRecommenders"`
	Honors             string `json:"Honors"`
	Awards             []struct {
		ID     string `json:"Id"`
		Name   string `json:"Name"`
		Issuer string `json:"Issuer"`
	} `json:"Awards"`
	Skills []struct {
		ID   string `json:"Id"`
		Name string `json:"Name"`
	} `json:"Skills"`
	CurrentStatus []struct {
		ID          string    `json:"Id"`
		Text        string    `json:"Text"`
		Source      string    `json:"Source"`
		CreatedDate time.Time `json:"CreatedDate"`
	} `json:"CurrentStatus"`
	Certifications []struct {
		ID        string    `json:"Id"`
		Name      string    `json:"Name"`
		Authority string    `json:"Authority"`
		Number    string    `json:"Number"`
		StartDate time.Time `json:"StartDate"`
		EndDate   time.Time `json:"EndDate"`
	} `json:"Certifications"`
	Courses []struct {
		ID     string `json:"Id"`
		Name   string `json:"Name"`
		Number string `json:"Number"`
	} `json:"Courses"`
	Volunteer []struct {
		ID           string `json:"Id"`
		Role         string `json:"Role"`
		Organization string `json:"Organization"`
		Cause        string `json:"Cause"`
	} `json:"Volunteer"`
	RecommendationsReceived []struct {
		ID                 string `json:"Id"`
		RecommendationType string `json:"RecommendationType"`
		RecommendationText string `json:"RecommendationText"`
		Recommender        string `json:"Recommender"`
	} `json:"RecommendationsReceived"`
	Languages []struct {
		ID          string `json:"Id"`
		Name        string `json:"Name"`
		Proficiency string `json:"Proficiency"`
	} `json:"Languages"`
	Projects []struct {
		ID      string `json:"Id"`
		Name    string `json:"Name"`
		Summary string `json:"Summary"`
		With    []struct {
			ID   string `json:"Id"`
			Name string `json:"Name"`
		} `json:"With"`
		StartDate time.Time `json:"StartDate"`
		EndDate   time.Time `json:"EndDate"`
		IsCurrent string    `json:"isCurrent"`
	} `json:"Projects"`
	Games []struct {
		ID          string    `json:"Id"`
		Name        string    `json:"Name"`
		Category    string    `json:"Category"`
		CreatedDate time.Time `json:"CreatedDate"`
	} `json:"Games"`
	Family []struct {
		ID           string `json:"Id"`
		Name         string `json:"Name"`
		Relationship string `json:"Relationship"`
	} `json:"Family"`
	TeleVisionShow []struct {
		ID          string    `json:"Id"`
		Name        string    `json:"Name"`
		Category    string    `json:"Category"`
		CreatedDate time.Time `json:"CreatedDate"`
	} `json:"TeleVisionShow"`
	MutualFriends []struct {
		ID        string    `json:"Id"`
		Name      string    `json:"Name"`
		FirstName string    `json:"FirstName"`
		LastName  string    `json:"LastName"`
		Birthday  time.Time `json:"Birthday"`
		Hometown  string    `json:"Hometown"`
		Link      string    `json:"Link"`
		Gender    string    `json:"Gender"`
	} `json:"MutualFriends"`
	Movies []struct {
		ID          string    `json:"Id"`
		Name        string    `json:"Name"`
		Category    string    `json:"Category"`
		CreatedDate time.Time `json:"CreatedDate"`
	} `json:"Movies"`
	Books []struct {
		ID          string    `json:"Id"`
		Name        string    `json:"Name"`
		Category    string    `json:"Category"`
		CreatedDate time.Time `json:"CreatedDate"`
	} `json:"Books"`
	AgeRange struct {
		Min int `json:"Min"`
		Max int `json:"Max"`
	} `json:"AgeRange"`
	PublicRepository string `json:"PublicRepository"`
	Hireable         bool   `json:"Hireable"`
	RepositoryURL    string `json:"RepositoryUrl"`
	Age              string `json:"Age"`
	Patents          []struct {
		ID    string    `json:"Id"`
		Title string    `json:"Title"`
		Date  time.Time `json:"Date"`
	} `json:"Patents"`
	FavoriteThings []struct {
		ID   string `json:"Id"`
		Name string `json:"Name"`
		Type string `json:"Type"`
	} `json:"FavoriteThings"`
	ProfessionalHeadline     string `json:"ProfessionalHeadline"`
	ProviderAccessCredential struct {
		AccessToken string `json:"AccessToken"`
		TokenSecret string `json:"TokenSecret"`
	} `json:"ProviderAccessCredential"`
	RelatedProfileViews []struct {
		ID        string `json:"Id"`
		FirstName string `json:"FirstName"`
		LastName  string `json:"LastName"`
	} `json:"RelatedProfileViews"`
	KloutScore struct {
		KloutID string `json:"KloutId"`
		Score   string `json:"Score"`
	} `json:"KloutScore"`
	LRUserID    string `json:"LRUserID"`
	PlacesLived []struct {
		Name      string `json:"Name"`
		IsPrimary string `json:"IsPrimary"`
	} `json:"PlacesLived"`
	Publications []struct {
		ID        string `json:"Id"`
		Title     string `json:"Title"`
		Publisher string `json:"Publisher"`
		Authors   []struct {
			ID   string `json:"Id"`
			Name string `json:"Name"`
		} `json:"Authors"`
		Date    time.Time `json:"Date"`
		URL     string    `json:"Url"`
		Summary string    `json:"Summary"`
	} `json:"Publications"`
	JobBookmarks []struct {
		IsApplied      bool   `json:"IsApplied"`
		ApplyTimestamp string `json:"ApplyTimestamp"`
		IsSaved        bool   `json:"IsSaved"`
		SavedTimestamp string `json:"SavedTimestamp"`
		Job            struct {
			Active             bool   `json:"Active"`
			ID                 string `json:"Id"`
			DescriptionSnippet string `json:"DescriptionSnippet"`
			Compony            struct {
				ID   string `json:"Id"`
				Name string `json:"Name"`
			} `json:"Compony"`
			Position struct {
				Title string `json:"Title"`
			} `json:"Position"`
			PostingTimestamp string `json:"PostingTimestamp"`
		} `json:"Job"`
	} `json:"JobBookmarks"`
	Suggestions struct {
		CompaniestoFollow []struct {
			ID   string `json:"Id"`
			Name string `json:"Name"`
		} `json:"CompaniestoFollow"`
		IndustriestoFollow []struct {
			ID   string `json:"Id"`
			Name string `json:"Name"`
		} `json:"IndustriestoFollow"`
		NewssourcetoFollow []struct {
			ID   string `json:"Id"`
			Name string `json:"Name"`
		} `json:"NewssourcetoFollow"`
		PeopletoFollow []struct {
			ID   string `json:"Id"`
			Name string `json:"Name"`
		} `json:"PeopletoFollow"`
	} `json:"Suggestions"`
	Badges []struct {
		BadgeID      string `json:"BadgeId"`
		Name         string `json:"Name"`
		BadgeMessage string `json:"BadgeMessage"`
		Description  string `json:"Description"`
		ImageURL     string `json:"ImageUrl"`
	} `json:"Badges"`
	MemberURLResources []struct {
		URL     string `json:"Url"`
		URLName string `json:"UrlName"`
	} `json:"MemberUrlResources"`
	TotalPrivateRepository int    `json:"TotalPrivateRepository"`
	Currency               string `json:"Currency"`
	StarredURL             string `json:"StarredUrl"`
	GistsURL               string `json:"GistsUrl"`
	PublicGists            int    `json:"PublicGists"`
	PrivateGists           int    `json:"PrivateGists"`
	Subscription           struct {
		Name          string `json:"Name"`
		Space         string `json:"Space"`
		PrivateRepos  string `json:"PrivateRepos"`
		Collaborators string `json:"Collaborators"`
	} `json:"Subscription"`
	Company          string `json:"Company"`
	GravatarImageURL string `json:"GravatarImageUrl"`
	ProfileImageUrls struct {
		Small             string `json:"Small"`
		Square            string `json:"Square"`
		Large             string `json:"Large"`
		ImageURL          string `json:"Image Url"`
		ThumbnailImageURL string `json:"ThumbnailImage Url"`
		Profile           string `json:"Profile"`
	} `json:"ProfileImageUrls"`
	WebProfiles struct {
		Small   string `json:"Small"`
		Square  string `json:"Square"`
		Large   string `json:"Large"`
		Profile string `json:"Profile"`
	} `json:"WebProfiles"`
	IsEmailSubscribed  bool   `json:"IsEmailSubscribed"`
	IsPosted           bool   `json:"IsPosted"`
	EmailVerified      bool   `json:"EmailVerified"`
	IsActive           bool   `json:"IsActive"`
	IsDeleted          bool   `json:"IsDeleted"`
	PhoneIDVerified    bool   `json:"PhoneIdVerified"`
	NoOfLogins         int    `json:"NoOfLogins"`
	IsLoginLocked      bool   `json:"IsLoginLocked"`
	LoginLockedType    string `json:"LoginLockedType"`
	LastLoginLocation  string `json:"LastLoginLocation"`
	RegistrationSource string `json:"RegistrationSource"`
	IsCustomUID        bool   `json:"IsCustomUid"`
	IsSecurePassword   bool   `json:"IsSecurePassword"`
	PrivacyPolicy      struct {
		Version        string    `json:"Version"`
		AcceptSource   string    `json:"AcceptSource"`
		AcceptDateTime time.Time `json:"AcceptDateTime"`
	} `json:"PrivacyPolicy"`
	ExternalIDs []struct {
		SourceID string `json:"SourceId"`
		Source   string `json:"Source"`
	} `json:"ExternalIds"`
	ProfileModifiedDate time.Time `json:"ProfileModifiedDate"`
	FirstLogin          bool      `json:"FirstLogin"`
	PinsCount           int       `json:"PinsCount"`
	BoardsCount         int       `json:"BoardsCount"`
	LikesCount          int       `json:"LikesCount"`
	SignupDate          time.Time `json:"SignupDate"`
	LastLoginDate       time.Time `json:"LastLoginDate"`
	Roles               []string  `json:"Roles"`
	IdentityProviders   []string  `json:"IdentityProviders"`
}
