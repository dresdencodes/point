package point

import (
	"point/req"
	"point/config"
)

func Pointer(endpoint string, data interface{}, rec interface{}) (interface{}, error) {
	err := pointRequest(pointsURL() + endpoint, data, rec)
	return rec, err
}

func PointerJSONFile(endpoint string, rec interface{}) (interface{}, error) {
	err := req.GetJSON(pointsURL() + endpoint, rec)
	return rec, err
}