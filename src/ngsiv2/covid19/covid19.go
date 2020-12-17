package covid19

type Age struct {
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Metadata struct{} `json:"metadata"`
}

type CityName struct {
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Metadata struct{} `json:"metadata"`
}

type Condition struct {
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Metadata struct{} `json:"metadata"`
}

type DayOfWeek struct {
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Metadata struct{} `json:"metadata"`
}

type Details struct {
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Metadata struct{} `json:"metadata"`
}

type Discharged struct {
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Metadata struct{} `json:"metadata"`
}

type Gender struct {
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Metadata struct{} `json:"metadata"`
}

type No struct {
	Type     string   `json:"type"`
	Value    int      `json:"value"`
	Metadata struct{} `json:"metadata"`
}

type OnsetDate struct {
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Metadata struct{} `json:"metadata"`
}

type PrefectureCode struct {
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Metadata struct{} `json:"metadata"`
}

type PrefectureName struct {
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Metadata struct{} `json:"metadata"`
}

type PublishedDate struct {
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Metadata struct{} `json:"metadata"`
}

type Remarks struct {
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Metadata struct{} `json:"metadata"`
}

type Residence struct {
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Metadata struct{} `json:"metadata"`
}

type Symptoms struct {
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Metadata struct{} `json:"metadata"`
}

type TravelRecord struct {
	Type     string   `json:"type"`
	Value    string   `json:"value"`
	Metadata struct{} `json:"metadata"`
}

type Patient struct {
	ID             string         `json:"id"`
	Type           string         `json:"type"`
	Age            Age            `json:"age"`
	CityName       CityName       `json:"cityName"`
	Condition      Condition      `json:"condition"`
	DayOfWeek      DayOfWeek      `json:"dayOfWeek"`
	Details        Details        `json:"details"`
	Discharged     Discharged     `json:"discharged"`
	Gender         Gender         `json:"gender"`
	No             No             `json:"no"`
	OnsetDate      OnsetDate      `json:"onsetDate"`
	PrefectureCode PrefectureCode `json:"prefectureCode"`
	PrefectureName PrefectureName `json:"prefectureName"`
	PublishedDate  PublishedDate  `json:"publishedDate"`
	Remarks        Remarks        `json:"remarks"`
	Residence      Residence      `json:"residence"`
	Symptoms       Symptoms       `json:"symptoms"`
	TravelRecord   TravelRecord   `json:"travelRecord"`
}

type Patients []Patient
