package loginradius

import (
	"encoding/json"
)

// AuthProfile is a struct meant to hold the LoginRadius Profile data mapping for JSON
type AuthProfile struct {
	PasswordHash            string            `json:"PasswordHash"`
	VerificationToken       string            `json:"VerificationToken"`
	ForgotToken             string            `json:"ForgotToken"`
	IsDeleteRequestAccepted bool              `json:"IsDeleteRequestAccepted"`
	Identities              []json.RawMessage `json:"Identities"`
	UserName                string            `json:"UserName"`
	UID                     string            `json:"Uid"`
	PasswordExpirationDate  TimeAlt           `json:"PasswordExpirationDate"`
	Password                string            `json:"Password"`
	LastPasswordChangeDate  TimeAlt           `json:"LastPasswordChangeDate"`
	PreviousUids            []string          `json:"PreviousUids"`
	CustomFields            json.RawMessage `json:"CustomFields"`
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
	ThumbnailImageURL string  `json:"ThumbnailImageUrl"`
	ImageURL          string  `json:"ImageUrl"`
	Favicon           string  `json:"Favicon"`
	ProfileURL        string  `json:"ProfileUrl"`
	HomeTown          string  `json:"HomeTown"`
	State             string  `json:"State"`
	City              string  `json:"City"`
	Industry          string  `json:"Industry"`
	About             string  `json:"About"`
	TimeZone          string  `json:"TimeZone"`
	LocalLanguage     string  `json:"LocalLanguage"`
	CoverPhoto        string  `json:"CoverPhoto"`
	TagLine           string  `json:"TagLine"`
	Language          string  `json:"Language"`
	Verified          string  `json:"Verified"`
	UpdatedTime       TimeAlt `json:"UpdatedTime"`
	Positions         []struct {
		Position  string  `json:"Position"`
		Summary   string  `json:"Summary"`
		StartDate TimeAlt `json:"StartDate"`
		EndDate   TimeAlt `json:"EndDate"`
		IsCurrent string  `json:"IsCurrent"`
		Location  string  `json:"Location"`
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
		School       string  `json:"School"`
		Year         string  `json:"year"`
		Type         string  `json:"type"`
		Notes        string  `json:"notes"`
		Activities   string  `json:"activities"`
		Degree       string  `json:"degree"`
		Fieldofstudy string  `json:"fieldofstudy"`
		StartDate    TimeAlt `json:"StartDate"`
		EndDate      TimeAlt `json:"EndDate"`
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
	MainAddress        string   `json:"MainAddress"`
	Created            string   `json:"Created"`
	CreatedDate        TimeAlt  `json:"CreatedDate"`
	ModifiedDate       TimeAlt  `json:"ModifiedDate"`
	LocalCity          string   `json:"LocalCity"`
	ProfileCity        string   `json:"ProfileCity"`
	LocalCountry       string   `json:"LocalCountry"`
	ProfileCountry     string   `json:"ProfileCountry"`
	IsProtected        bool     `json:"IsProtected"`
	RelationshipStatus string   `json:"RelationshipStatus"`
	Quota              string   `json:"Quota"`
	Quote              string   `json:"Quote"`
	InterestedIn       []string `json:"InterestedIn"`
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
		ID          string  `json:"Id"`
		Text        string  `json:"Text"`
		Source      string  `json:"Source"`
		CreatedDate TimeAlt `json:"CreatedDate"`
	} `json:"CurrentStatus"`
	Certifications []struct {
		ID        string  `json:"Id"`
		Name      string  `json:"Name"`
		Authority string  `json:"Authority"`
		Number    string  `json:"Number"`
		StartDate TimeAlt `json:"StartDate"`
		EndDate   TimeAlt `json:"EndDate"`
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
		StartDate TimeAlt `json:"StartDate"`
		EndDate   TimeAlt `json:"EndDate"`
		IsCurrent string  `json:"isCurrent"`
	} `json:"Projects"`
	Games []struct {
		ID          string  `json:"Id"`
		Name        string  `json:"Name"`
		Category    string  `json:"Category"`
		CreatedDate TimeAlt `json:"CreatedDate"`
	} `json:"Games"`
	Family []struct {
		ID           string `json:"Id"`
		Name         string `json:"Name"`
		Relationship string `json:"Relationship"`
	} `json:"Family"`
	TeleVisionShow []struct {
		ID          string  `json:"Id"`
		Name        string  `json:"Name"`
		Category    string  `json:"Category"`
		CreatedDate TimeAlt `json:"CreatedDate"`
	} `json:"TeleVisionShow"`
	MutualFriends []struct {
		ID        string  `json:"Id"`
		Name      string  `json:"Name"`
		FirstName string  `json:"FirstName"`
		LastName  string  `json:"LastName"`
		Birthday  TimeAlt `json:"Birthday"`
		Hometown  string  `json:"Hometown"`
		Link      string  `json:"Link"`
		Gender    string  `json:"Gender"`
	} `json:"MutualFriends"`
	Movies []struct {
		ID          string  `json:"Id"`
		Name        string  `json:"Name"`
		Category    string  `json:"Category"`
		CreatedDate TimeAlt `json:"CreatedDate"`
	} `json:"Movies"`
	Books []struct {
		ID          string  `json:"Id"`
		Name        string  `json:"Name"`
		Category    string  `json:"Category"`
		CreatedDate TimeAlt `json:"CreatedDate"`
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
		ID    string  `json:"Id"`
		Title string  `json:"Title"`
		Date  TimeAlt `json:"Date"`
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
		Date    TimeAlt `json:"Date"`
		URL     string  `json:"Url"`
		Summary string  `json:"Summary"`
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
		Version        string  `json:"Version"`
		AcceptSource   string  `json:"AcceptSource"`
		AcceptDateTime TimeAlt `json:"AcceptDateTime"`
	} `json:"PrivacyPolicy"`
	ExternalIDs []struct {
		SourceID string `json:"SourceId"`
		Source   string `json:"Source"`
	} `json:"ExternalIds"`
	ProfileModifiedDate TimeAlt  `json:"ProfileModifiedDate"`
	FirstLogin          bool     `json:"FirstLogin"`
	PinsCount           int      `json:"PinsCount"`
	BoardsCount         int      `json:"BoardsCount"`
	LikesCount          int      `json:"LikesCount"`
	SignupDate          TimeAlt  `json:"SignupDate"`
	LastLoginDate       TimeAlt  `json:"LastLoginDate"`
	Roles               []string `json:"Roles"`
	IdentityProviders   []string `json:"IdentityProviders"`
}

// Identities is the struct that contains identities for social profiles
type Identities struct {
	ID          string `json:"ID"`
	Provider    string `json:"Provider"`
	Prefix      string `json:"Prefix"`
	FirstName   string `json:"FirstName"`
	MiddleName  string `json:"MiddleName"`
	LastName    string `json:"LastName"`
	Suffix      string `json:"Suffix"`
	FullName    string `json:"FullName"`
	NickName    string `json:"NickName"`
	ProfileName string `json:"ProfileName"`
	BirthDate   string `json:"BirthDate"`
	Gender      string `json:"Gender"`
	Website     string `json:"Website"`
	Email       []struct {
		Type  string `json:"Type"`
		Value string `json:"Value"`
	} `json:"Email"`
	Country struct {
		Code string `json:"Code"`
		Name string `json:"Name"`
	} `json:"Country"`
	ThumbnailImageURL string  `json:"ThumbnailImageUrl"`
	ImageURL          string  `json:"ImageUrl"`
	Favicon           string  `json:"Favicon"`
	ProfileURL        string  `json:"ProfileUrl"`
	HomeTown          string  `json:"HomeTown"`
	State             string  `json:"State"`
	City              string  `json:"City"`
	Industry          string  `json:"Industry"`
	About             string  `json:"About"`
	TimeZone          string  `json:"TimeZone"`
	LocalLanguage     string  `json:"LocalLanguage"`
	CoverPhoto        string  `json:"CoverPhoto"`
	TagLine           string  `json:"TagLine"`
	Language          string  `json:"Language"`
	Verified          string  `json:"Verified"`
	UpdatedTime       TimeAlt `json:"UpdatedTime"`
	Positions         []struct {
		Position  string  `json:"Position"`
		Summary   string  `json:"Summary"`
		StartDate TimeAlt `json:"StartDate"`
		EndDate   TimeAlt `json:"EndDate"`
		IsCurrent string  `json:"IsCurrent"`
		Location  string  `json:"Location"`
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
		School       string  `json:"School"`
		Year         string  `json:"year"`
		Type         string  `json:"type"`
		Notes        string  `json:"notes"`
		Activities   string  `json:"activities"`
		Degree       string  `json:"degree"`
		Fieldofstudy string  `json:"fieldofstudy"`
		StartDate    TimeAlt `json:"StartDate"`
		EndDate      TimeAlt `json:"EndDate"`
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
	MainAddress        string   `json:"MainAddress"`
	Created            bool     `json:"Created"`
	CreatedDate        TimeAlt  `json:"CreatedDate"`
	ModifiedDate       TimeAlt  `json:"ModifiedDate"`
	LocalCity          string   `json:"LocalCity"`
	ProfileCity        string   `json:"ProfileCity"`
	LocalCountry       string   `json:"LocalCountry"`
	ProfileCountry     string   `json:"ProfileCountry"`
	IsProtected        bool     `json:"IsProtected"`
	RelationshipStatus string   `json:"RelationshipStatus"`
	Quota              string   `json:"Quota"`
	Quote              string   `json:"Quote"`
	InterestedIn       []string `json:"InterestedIn"`
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
	IsGeoEnabled       bool   `json:"IsGeoEnabled"`
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
		ID          string  `json:"Id"`
		Text        string  `json:"Text"`
		Source      string  `json:"Source"`
		CreatedDate TimeAlt `json:"CreatedDate"`
	} `json:"CurrentStatus"`
	Certifications []struct {
		ID        string  `json:"Id"`
		Name      string  `json:"Name"`
		Authority string  `json:"Authority"`
		Number    string  `json:"Number"`
		StartDate TimeAlt `json:"StartDate"`
		EndDate   TimeAlt `json:"EndDate"`
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
		StartDate TimeAlt `json:"StartDate"`
		EndDate   TimeAlt `json:"EndDate"`
		IsCurrent string  `json:"isCurrent"`
	} `json:"Projects"`
	Games []struct {
		ID          string  `json:"Id"`
		Name        string  `json:"Name"`
		Category    string  `json:"Category"`
		CreatedDate TimeAlt `json:"CreatedDate"`
	} `json:"Games"`
	Family []struct {
		ID           string `json:"Id"`
		Name         string `json:"Name"`
		Relationship string `json:"Relationship"`
	} `json:"Family"`
	TeleVisionShow []struct {
		ID          string  `json:"Id"`
		Name        string  `json:"Name"`
		Category    string  `json:"Category"`
		CreatedDate TimeAlt `json:"CreatedDate"`
	} `json:"TeleVisionShow"`
	MutualFriends []struct {
		ID        string  `json:"Id"`
		Name      string  `json:"Name"`
		FirstName string  `json:"FirstName"`
		LastName  string  `json:"LastName"`
		Birthday  TimeAlt `json:"Birthday"`
		Hometown  string  `json:"Hometown"`
		Link      string  `json:"Link"`
		Gender    string  `json:"Gender"`
	} `json:"MutualFriends"`
	Movies []struct {
		ID          string  `json:"Id"`
		Name        string  `json:"Name"`
		Category    string  `json:"Category"`
		CreatedDate TimeAlt `json:"CreatedDate"`
	} `json:"Movies"`
	Books []struct {
		ID          string  `json:"Id"`
		Name        string  `json:"Name"`
		Category    string  `json:"Category"`
		CreatedDate TimeAlt `json:"CreatedDate"`
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
		ID    string  `json:"Id"`
		Title string  `json:"Title"`
		Date  TimeAlt `json:"Date"`
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
		Date    TimeAlt `json:"Date"`
		URL     string  `json:"Url"`
		Summary string  `json:"Summary"`
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
	IsEmailSubscribed  bool                       `json:"IsEmailSubscribed"`
	IsPosted           bool                       `json:"IsPosted"`
	Data               map[string]json.RawMessage `json:"Data"`
	EmailVerified      bool                       `json:"EmailVerified"`
	IsActive           bool                       `json:"IsActive"`
	IsDeleted          bool                       `json:"IsDeleted"`
	PhoneIDVerified    bool                       `json:"PhoneIdVerified"`
	NoOfLogins         int                        `json:"NoOfLogins"`
	IsLoginLocked      bool                       `json:"IsLoginLocked"`
	LoginLockedType    string                     `json:"LoginLockedType"`
	LastLoginLocation  string                     `json:"LastLoginLocation"`
	RegistrationSource string                     `json:"RegistrationSource"`
	IsCustomUID        bool                       `json:"IsCustomUid"`
	IsSecurePassword   bool                       `json:"IsSecurePassword"`
	PrivacyPolicy      struct {
		Version        string  `json:"Version"`
		AcceptSource   string  `json:"AcceptSource"`
		AcceptDateTime TimeAlt `json:"AcceptDateTime"`
	} `json:"PrivacyPolicy"`
	ExternalIDs []struct {
		SourceID string `json:"SourceId"`
		Source   string `json:"Source"`
	} `json:"ExternalIds"`
	ProfileModifiedDate TimeAlt  `json:"ProfileModifiedDate"`
	FirstLogin          bool     `json:"FirstLogin"`
	PinsCount           int      `json:"PinsCount"`
	BoardsCount         int      `json:"BoardsCount"`
	LikesCount          int      `json:"LikesCount"`
	SignupDate          TimeAlt  `json:"SignupDate"`
	LastLoginDate       TimeAlt  `json:"LastLoginDate"`
	Roles               []string `json:"Roles"`
	Medium              string
}

// AuthEmail is a struct for the Authorization Email Endpoint response
type AuthEmail struct {
	IsPosted  bool `json:"IsPosted"`
	IsExist   bool `json:"IsExist"`
	IsDeleted bool `json:"IsDeleted"`
	Data      struct {
		Email       string      `json:"Email"`
		Profile     AuthProfile `json:"Profile"`
		AccessToken string      `json:"access_token"`
		ExpiresIn   TimeAlt     `json:"expires_in"`
	} `json:"Data"`
}

// AuthRegister is a struct for the Authorization Register Endpoint response
type AuthRegister struct {
	IsPosted bool `json:"IsPosted"`
	Data     json.RawMessage
}

// AuthLogin is a struct for the Authorization Login Endpoint response
type AuthLogin struct {
	Profile     AuthProfile `json:"Profile"`
	AccessToken string      `json:"access_token"`
	ExpiresIn   TimeAlt     `json:"expires_in"`
}

// AuthBool is a struct for the Authorization Endpoints that return a boolean reply
type AuthBool struct {
	IsExist                 bool `json:"IsExist"`
	IsPosted                bool `json:"IsPosted"`
	IsDeleted               bool `json:"IsDeleted"`
	IsDeleteRequestAccepted bool `json:"IsDeleteRequestAccepted"`
}

// AuthAccessToken is a struct for the Authorization AccessToken Endpoint response
type AuthAccessToken struct {
	AccessToken string  `json:"access_token"`
	ExpiresIn   TimeAlt `json:"expires_in"`
	IsPosted    bool    `json:"IsPosted"`
}

// AuthSecurityQuestion is a struct for the Authorization SecurityQuestion Endpoint response
type AuthSecurityQuestion []struct {
	QuestionID string `json:"QuestionId"`
	Question   string `json:"Question"`
}

// AuthUpdate is a struct for the Auth Update APIs
type AuthUpdate struct {
	IsPosted bool        `json:"IsPosted"`
	Data     AuthProfile `json:"Data"`
}
