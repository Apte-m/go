package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
Данная задача ориентирована на реальную работу с данными в формате JSON. Для решения мы будем
использовать справочник ОКВЭД (Общероссийский классификатор видов экономической деятельности),
который был размещен на web-портале data.gov.ru.
Необходимая вам информация о структуре данных содержится в файле structure-20190514T0000.json,
а сами данные, которые вам потребуется декодировать, содержатся в файле data-20190514T0100.json.
Файлы размещены в https://github.com/semyon-dev/stepik-go/tree/master/work_with_json.
Для того, чтобы показать, что вы действительно смогли декодировать документ вам необходимо в
качестве ответа записать сумму полей global_id всех элементов, закодированных в файле.
*/

type Student struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Address    string
	Phone      string
	Rating     []int
}

type Group struct {
	ID       int
	Number   string
	Year     int
	Students []Student
}

type Result struct {
	Average float64
}

func exa() {

	type Item struct {
		SystemObjectId string `json:"system_object_id"`
		Kod            string `json:"Kod"`
		IsDeleted      int    `json:"is_deleted"`
		SignatureDate  string `json:"signature_date"`
		Nomdescr       string `json:"Nomdescr"`
		GlobalId       int    `json:"global_id"`
		Idx            string `json:"Idx"`
		Razdel         string `json:"Razdel"`
		Name           string `json:"Name"`
	}

	const URL = "https://raw.githubusercontent.com/semyon-dev/stepik-go/master/work_with_json/data-20190514T0100.json"

	rs, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	defer rs.Body.Close()

	dec := json.NewDecoder(rs.Body)

	var items []Item
	err = dec.Decode(&items)
	if err != nil {
		panic(err)
	}

	sum := 0
	for _, x := range items {
		sum += x.GlobalId
	}
	fmt.Println(sum)
	// 763804981288
}

func exa2() {
	var group Group
	data := []byte(`
		{
			"ID":134,
			"Number":"ИЛМ-1274",
			"Year":2,
			"Students":[
				{
					"LastName":"Вещий",
					"FirstName":"Лифон",
					"MiddleName":"Вениаминович",
					"Birthday":"4апреля1970года",
					"Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
					"Phone":"+7(948)709-47-24",
					"Rating":[1,2,3]
				},
				{
					"LastName":"Вещий",
					"FirstName":"Лифон",
					"MiddleName":"Вениаминович",
					"Birthday":"4апреля1970года",
					"Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
					"Phone":"+7(948)709-47-24",
					"Rating":[4,2,4]
				}
			]
		}
	`)
	//fmt.Println(string(data))

	if err := json.Unmarshal(data, &group); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(data, &group); err != nil {
		panic(err)
	}

	var numberRating int
	for _, student := range group.Students {
		numberRating += len(student.Rating)
	}

	result := Result{
		Average: float64(numberRating) / float64(len(group.Students)),
	}
	dataJson, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dataJson))

}
