package point

import (
	"bytes"
	"errors"
    "net/http"
	"io/ioutil"
    "encoding/json"
)

type pointerRequestError struct {
	Error		string      `json:"error"`
}

func pointRequest(url string, in interface{}, out interface{}) error {
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

	body, err := ioutil.ReadAll(response.Body)
	if err!=nil { return err }

	if response.StatusCode==400 {
		errorContainer := pointerRequestError{}
		err = json.Unmarshal(body, &errorContainer)
		if err != nil {
			return err
		}
		return errors.New(errorContainer.Error)
	}
	
	if out == nil {return nil}

	err = json.Unmarshal(body, &out)
	if err != nil {
		return err
	}
	return err

}