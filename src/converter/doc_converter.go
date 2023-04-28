package converter

import (
	"code.sajari.com/docconv"
)

func ConvertDoc(path string) (*docconv.Response, error) {
	res, err := docconv.ConvertPath(path)
	if err != nil {
		return nil, err
	}
	// fmt.Println(res.Body)
	return res, nil
}
