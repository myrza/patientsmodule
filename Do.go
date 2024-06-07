/*
v1.1.0 версия с сортировкой по возрасту

*/

package patientsmodule

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sort"
)

type paitent struct {
	Name  string `json:"name"`
	Age   int    `"json:"age"`
	Email string `"json:"email"`
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

	f, err = ioutil.TempFile("./", tgt)
	if err != nil {
		return err
	}

	err = json.NewEncoder(f).Encode(res)
	if err != nil {
		return err
	}

	err = f.Close()
	if err != nil {
		return err
	}

	f.Close()

	return nil
}
