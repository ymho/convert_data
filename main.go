package main

import (
	"./src/ngsiv2/covid19"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type InputData struct {
	LastUpdate string `json:"lastUpdate"`
	Patients   struct {
		Date string `json:"date"`
		Data []struct {
			No                string `json:"No"`
			PublishedDateTime string `json:"発表日"`
			AgeGender         string `json:"年代性別"`
			Residence         string `json:"国籍"`
			Livingin          string `json:"住居地"`
			Condition         string `json:"接触状況"`
			Remarks           string `json:"備考"`
			Date              string `json:"date"`
			W                 string `json:"w"`
			ShortDate         string `json:"short_date"`
		} `json:"data"`
	} `json:"patients"`
}

type Request struct {
	ActionType string           `json:"actionType"`
	Entities   covid19.Patients `json:"entities"`
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
	var patients []covid19.Patient
	for index, data := range inputData.Patients.Data {
		no, err := strconv.Atoi(data.No)
		if err != nil {
			log.Fatal(err)
		}
		age := ""
		if strings.Index(data.AgeGender, "代") != -1 {
			age = data.AgeGender[0:strings.Index(data.AgeGender, "代")]
		}
		if strings.Index(data.AgeGender, "未満") != -1 {
			age = data.AgeGender[0:strings.Index(data.AgeGender, "未満")]
		}
		gender := ""
		if (strings.Index(data.AgeGender, "男")) != -1 {
			gender = "男"
		}
		if (strings.Index(data.AgeGender, "女")) != -1 {
			gender = "女"
		}
		patient := covid19.Patient{
			ID:             "urn:ngsi-ld:covid19:Patients:aichi:" + strconv.Itoa(index+1),
			Type:           "Patients",
			Age:            covid19.Age{Type: "Text", Value: age},
			CityName:       covid19.CityName{Type: "Text", Value: data.Livingin},
			Condition:      covid19.Condition{Type: "Text", Value: ""},
			DayOfWeek:      covid19.DayOfWeek{Type: "Text", Value: ""},
			Details:        covid19.Details{Type: "Text", Value: data.Condition},
			Discharged:     covid19.Discharged{Type: "Text", Value: ""},
			Gender:         covid19.Gender{Type: "Text", Value: gender},
			No:             covid19.No{Type: "Number", Value: no},
			OnsetDate:      covid19.OnsetDate{Type: "Text", Value: ""},
			PrefectureCode: covid19.PrefectureCode{Type: "Text", Value: "230006"},
			PrefectureName: covid19.PrefectureName{Type: "Text", Value: "愛知県"},
			PublishedDate:  covid19.PublishedDate{Type: "Text", Value: data.Date},
			Remarks:        covid19.Remarks{Type: "Text", Value: ""},
			Residence:      covid19.Residence{Type: "Text", Value: data.Residence},
			Symptoms:       covid19.Symptoms{Type: "Text", Value: ""},
			TravelRecord:   covid19.TravelRecord{Type: "Text", Value: ""},
		}
		patients = append(patients, patient)
	}
	sample_json, _ := json.MarshalIndent(patients, "", "  ")
	fmt.Printf("%s\n", string(sample_json))

	//request := Request{
	//	ActionType: "append",
	//	Entities:   patients,
	//}

	//rq, _ := json.MarshalIndent(request,"","  ")

}
