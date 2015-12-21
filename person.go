package gofc

type ContactInfo struct {
	FamilyName  string   `json:"familyName,omitempty"`
	GivenName   string   `json:"givenName,omitempty"`
	FullName    string   `json:"fullName,omitempty"`
	MiddleNames []string `json:"middleNames,omitempty"`
	Websites    []struct {
		URL string `json:"url"`
	} `json:"websites,omitempty"`
	Chats []struct {
		Handle string `json:"handle,omitempty"`
		Client string `json:"client,omitempty"`
	}
}

type Demographics struct {
	LocationGeneral string `json:"locationGeneral,omitempty"`
	LocationDeduced struct {
		NormalizedLocation string `json:"normalizedLocation,omitempty"`
		DeducedLocation    string `json:"deducedLocation,omitempty"`
		City               struct {
			Deduced bool   `json:"deduced"`
			Name    string `json:"name"`
		} `json:"city,omitempty"`
		State struct {
			Deduced bool   `json:"deduced"`
			Name    string `json:"name"`
			Code    string `json:"code"`
		} `json:"state,omitempty"`
		Country struct {
			Deduced bool   `json:"deduced"`
			Name    string `json:"name"`
			Code    string `json:"code"`
		} `json:"country,omitempty"`
		Continent struct {
			Deduced bool   `json:"deduced"`
			Name    string `json:"name"`
		} `json:"continent,omitempty"`
		County struct {
			Deduced bool   `json:"deduced"`
			Name    string `json:"name"`
			Code    string `json:"code"`
		} `json:"county,omitempty"`
		likelihood float64 `json:"likelihood"`
	} `json:"locationDeduced,omitempty"`
	Age      string `json:"age,omitempty"`
	Gender   string `json:"gender,omitempty"`
	AgeRange string `json:"ageRange,omitempty"`
}

type Photo struct {
	TypeID    string `json:"typeId"`
	TypeName  string `json:"typeName"`
	URL       string `json:"url"`
	IsPrimary bool   `json:"isPrimary"`
}

type SocialProfile struct {
	TypeID    string `json:"typeId"`
	TypeName  string `json:"typeName"`
	ID        string `json:"id"`
	Username  string `json:"username"`
	URL       string `json:"url"`
	Bio       string `json:"bio,omitempty"`
	RSS       string `json:"rss,omitempty"`
	Following int    `json:"following,omitempty"`
	Followers int    `json:"followers,omitempty"`
}

type DigitalFootprint struct {
	Topics []struct {
		Value    string `json:"value"`
		Provider string `json:"provider"`
	} `json:"topics,omitempty"`
	Scores []struct {
		Provider string  `json:"provider"`
		Type     string  `json:"type"`
		Value    float64 `json:"value"`
	} `json:"scores,omitempty"`
}

type Organization struct {
	Title     string `json:"title"`
	Name      string `json:"name"`
	StartDate string `json:"startDate,omitempty"`
	EndDate   string `json:"endDate,omitempty"`
	IsPrimary bool   `json:"isPrimary"`
	Current   bool   `json:"current"`
}

type PersonResponse struct {
	Status           int              `json:"status"`
	RequestID        string           `json:"requestId"`
	Likelihood       float64          `json:"Likelihood"`
	ContactInfo      ContactInfo      `json:"contactInfo,omitempty"`
	Demographics     Demographics     `json:"demographics,omitempty"`
	Photos           []Photo          `json:"photos,omitempty"`
	SocialProfiles   []SocialProfile  `json:"socialProfiles,omitempty"`
	DigitalFootprint DigitalFootprint `json:"digitalFootprint,omitempty"`
	Organizations    []Organization   `json:"organizations,omitempty"`
}
