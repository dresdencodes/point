package point

import (
	"heiro/pkg/dg/dgreq"
	"heiro/pkg/point/config"
)


func Pointer(endpoint string, data interface{}, rec interface{}) (interface{}, error) {
	err := pointRequest(config.PointURL + endpoint, data, rec)
	return rec, err
}

func PointerJSONFile(endpoint string, rec interface{}) (interface{}, error) {
	err := dgreq.GetJSON(config.PointURL + endpoint, rec)
	return rec, err
}