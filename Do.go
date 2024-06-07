/*
v2.1.1 версия с xml без сортировки

*/

package patientsmodule

import (
	//"encoding/json"
	"encoding/json"
	"sort"

	"encoding/xml"
	"os"
)

type paitent struct {
	Name  string `json:"name"`
	Age   int    `"json:"age"`
	Email string `"json:"email"`
}

type xml_patient struct {
	Name  string `xml:"Name"`
	Age   int    `xml:"Age"`
	Email string `xml:"Email"`
}
type patients struct {
	List []xml_patient `xml:"Patient"`
}

func Do(src string, tgt string) error {

	f, err := os.Open(src)
	if err != nil {

		return err
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	res := make([]paitent, 0, 3)

	for dec.More() {
		var p paitent
		err := dec.Decode(&p)
		if err != nil {
			return err
		}
		res = append(res, p)

	}

	// отсортируем по годам
	sort.Slice(res[:], func(i, j int) bool {
		return res[i].Age < res[j].Age
	})

	f, err = os.Create(tgt)
	if err != nil {
		return err
	}

	f.WriteString(xml.Header)

	x1 := xml_patient{
		Name:  res[0].Name,
		Age:   res[0].Age,
		Email: res[0].Email,
	}
	x2 := xml_patient{
		Name:  res[1].Name,
		Age:   res[1].Age,
		Email: res[1].Email,
	}
	x3 := xml_patient{
		Name:  res[2].Name,
		Age:   res[2].Age,
		Email: res[2].Email,
	}
	d := patients{}
	d.List = append(d.List, x1)
	d.List = append(d.List, x2)
	d.List = append(d.List, x3)

	enc := xml.NewEncoder(f)
	enc.Indent("", "    ")
	err = enc.Encode(d)
	if err != nil {
		return err
	}
	f.Close()
	f, err = os.Open(tgt)
	if err != nil {
		return err
	}
	defer f.Close()

	err = xml.NewDecoder(f).Decode(&d)
	if err != nil {
		return err
	}

	return nil
}
