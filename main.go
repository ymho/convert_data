package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Age struct {
	Type     string `json:"type"`
	Value    string `json:"value"`
	Metadata struct {} `json:"metadata"`
}

type CityName struct {
	Type     string `json:"type"`
	Value    string `json:"value"`
	Metadata struct {} `json:"metadata"`
}

type Condition struct {
	Type     string `json:"type"`
	Value    string `json:"value"`
	Metadata struct {} `json:"metadata"`
}

type DayOfWeek struct {
	Type     string `json:"type"`
	Value    string `json:"value"`
	Metadata struct {} `json:"metadata"`
}

type Details struct {
	Type     string `json:"type"`
	Value    string `json:"value"`
	Metadata struct {} `json:"metadata"`
}

type Discharged struct {
	Type     string `json:"type"`
	Value    string `json:"value"`
	Metadata struct {} `json:"metadata"`
}

type Gender struct {
	Type     string `json:"type"`
	Value    string `json:"value"`
	Metadata struct {} `json:"metadata"`
}

type No struct {
	Type     string `json:"type"`
	Value    int    `json:"value"`
	Metadata struct {} `json:"metadata"`
}

type OnsetDate struct {
	Type     string `json:"type"`
	Value    string `json:"value"`
	Metadata struct {} `json:"metadata"`
}

type PrefectureCode struct {
	Type     string `json:"type"`
	Value    string `json:"value"`
	Metadata struct {} `json:"metadata"`
}

type PrefectureName struct {
	Type     string `json:"type"`
	Value    string `json:"value"`
	Metadata struct {} `json:"metadata"`
}

type PublishedDate struct {
	Type     string `json:"type"`
	Value    string `json:"value"`
	Metadata struct {} `json:"metadata"`
}

type Remarks struct {
	Type     string `json:"type"`
	Value    string `json:"value"`
	Metadata struct {} `json:"metadata"`
}

type Residence struct {
	Type     string `json:"type"`
	Value    string `json:"value"`
	Metadata struct {} `json:"metadata"`
}

type Symptoms struct {
	Type     string `json:"type"`
	Value    string `json:"value"`
	Metadata struct {} `json:"metadata"`
}

type TravelRecord struct {
	Type     string `json:"type"`
	Value    string `json:"value"`
	Metadata struct {} `json:"metadata"`
}

type Patient struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Age  Age `json:"age"`
	CityName CityName `json:"cityName"`
	Condition Condition `json:"condition"`
	DayOfWeek DayOfWeek `json:"dayOfWeek"`
	Details Details `json:"details"`
	Discharged Discharged `json:"discharged"`
	Gender Gender `json:"gender"`
	No No `json:"no"`
	OnsetDate OnsetDate `json:"onsetDate"`
	PrefectureCode PrefectureCode `json:"prefectureCode"`
	PrefectureName PrefectureName`json:"prefectureName"`
	PublishedDate PublishedDate`json:"publishedDate"`
	Remarks Remarks `json:"remarks"`
	Residence Residence `json:"residence"`
	Symptoms Symptoms`json:"symptoms"`
	TravelRecord TravelRecord `json:"travelRecord"`
}

type Patients []Patient

type InputData struct {
	LastUpdate string `json:"lastUpdate"`
	Patients   struct {
		Date string `json:"date"`
		Data []struct {
			No            string `json:"No"`
			PublishedDateTime string `json:"発表日"`
			AgeGender 	  string `json:"年代性別"`
			Residence 	  string `json:"国籍"`
			Livingin	  string `json:"住居地"`
			Condition	  string `json:"接触状況"`
			Remarks	      string `json:"備考"`
			Date          string `json:"date"`
			W             string `json:"w"`
			ShortDate     string `json:"short_date"`
		} `json:"data"`
	} `json:"patients"`
}

func main() {
	bytes, err := ioutil.ReadFile("./source/data.json")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Print(string(bytes))
	var inputData InputData
	if err := json.Unmarshal(bytes, &inputData); err != nil {
		log.Fatal(err)
	}
	var patients []Patient
	for index, data := range inputData.Patients.Data {
		no, err := strconv.Atoi(data.No)
		if err != nil {
			log.Fatal(err)
		}
		age := ""
		if strings.Index(data.AgeGender,"代") != -1{
			age = data.AgeGender[0:strings.Index(data.AgeGender,"代")]
		}
		if strings.Index(data.AgeGender,"未満") != -1{
			age = data.AgeGender[0:strings.Index(data.AgeGender,"未満")]
		}
		gender := ""
		if (strings.Index(data.AgeGender, "男")) != -1{
			gender = "男"
		}
		if (strings.Index(data.AgeGender, "女")) != -1{
			gender = "女"
		}
		patient := Patient{
			ID:             "urn:ngsi-ld:covid19:Patients:aichi:" + strconv.Itoa(index+1),
			Type:           "Patients",
			Age:            Age {Type: "Text", Value: age},
			CityName:       CityName {Type: "Text", Value: data.Livingin},
			Condition:      Condition{Type: "Text", Value: ""},
			DayOfWeek:      DayOfWeek{Type: "Text", Value: ""},
			Details:        Details{Type: "Text", Value: data.Condition},
			Discharged:     Discharged{Type: "Text", Value: ""},
			Gender:         Gender{Type: "Text", Value: gender},
			No:             No{Type: "Number", Value: no},
			OnsetDate:      OnsetDate{Type: "Text", Value: ""},
			PrefectureCode: PrefectureCode{Type: "Text", Value: "230006"},
			PrefectureName: PrefectureName{Type: "Text", Value: "愛知県"},
			PublishedDate:  PublishedDate{Type: "Text", Value: data.Date},
			Remarks:        Remarks{Type: "Text", Value: ""},
			Residence:      Residence{Type: "Text", Value: data.Residence},
			Symptoms:       Symptoms{Type: "Text", Value: ""},
			TravelRecord:   TravelRecord{Type: "Text", Value: ""},
		}
		patients = append(patients, patient)
	}
	sample_json, _ := json.MarshalIndent(patients,"","  ")
	fmt.Printf("[+] %s\n", string(sample_json))
}