package req

import (
	"bytes"
    "net/http"
	"io/ioutil"
    "encoding/json"
)

func GetJSON(url string, i interface{}) error {
	
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type","application/json")

	response, err := (&http.Client{}).Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body, i)
	if err != nil {
		return err
	}
	return nil
}


func PostJSON(url string, in interface{}, out interface{}) error {
	jsonData, err := json.Marshal(in)
	if err != nil {
		return err
	}
	
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err!=nil {
		return err
	}

	req.Header.Add("Content-Type","application/json")
	response, err := (&http.Client{}).Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if out == nil {return nil}
	body, err := ioutil.ReadAll(response.Body)
	if err!=nil {return err}

	err = json.Unmarshal(body, &out)
	if err != nil {
		return err
	}
	return err

}
