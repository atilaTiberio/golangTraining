package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type dataSets struct {
	Pagination struct {
		PageSize int `json:"pageSize"`
		Page     int `json:"page"`
		Total    int `json:"total"`
	} `json:"pagination"`
	Results []struct {
		ID        string   `json:"_id"`
		Variables []string `json:"variables"`
		Count     int      `json:"count"`
		URL       string   `json:"url"`
		Endpoint  string   `json:"endpoint"`
	} `json:"results"`
}

func main() {

	response, err := http.Get("https://api.datos.gob.mx/v1/api-catalog")

	if err != nil {
		fmt.Println("Something went wrong ", err)
	}

	body, error := ioutil.ReadAll(response.Body)

	if error != nil {
		fmt.Println("Something went wrong ", error)
	}

	var resultado dataSets
	error = json.Unmarshal(body, &resultado)

	if error != nil {
		fmt.Println(error)
	}

	fmt.Printf("pageSize: %v, page: %v, total %v\n", resultado.Pagination.PageSize, resultado.Pagination.Page, resultado.Pagination.Total)

	for _, _val := range resultado.Results{
		fmt.Printf("Count: %v URL: %v, variables: %v\n",_val.Count,_val.URL,_val.Variables)
	}

}
