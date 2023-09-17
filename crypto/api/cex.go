package api

import ("github.com/amsem/crypto/data"
	"fmt"
	"strings"
	"net/http"
	"io"
	"encoding/json"
	)
const apiUrl string = "https://cex.io/api/ticker/%s/EUR"
func GetRate(c string) (*data.Rate, error){
	if len(c) !=3 {
		return nil, fmt.Errorf("3 char required; %v received", len(c))
	}
	upperC := strings.ToUpper(c)
	response, err := http.Get(fmt.Sprintf(apiUrl, upperC))
	if err != nil{
		return nil, err
	}
	var res ResCEX
	if response.StatusCode == http.StatusOK{
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil{
				return nil, err
			}
		//j := string(bodyBytes)
		err = json.Unmarshal(bodyBytes, &res)
		if err != nil{
			return nil, err
		}
		//fmt.Println(j)
	}else {
		return nil, fmt.Errorf("status Code received: %v\n", response.StatusCode)
	}
	r := data.Rate{Currency: c, Price: res.Bid}
	return &r, nil
}
