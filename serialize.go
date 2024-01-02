package main

import (
	"encoding/json"
	"fmt"
)

const rawResp = `
{
   "header": {
	   "code": 0,
	   "message": ""
   },
   "data": [{
	   "type": "user",
	   "id": 100,
	   "attributes": {
		   "email": "bob@yandex.ru",
		   "article_ids": [10, 11, 12]
	   }
   }]
}
`

type Header struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Attributes struct {
	Email      string `json:"email"`
	ArticleIds []int  `json:"article_ids"`
}

type User struct {
	Type       string     `json:"type"`
	Id         int        `json:"id"`
	Attributes Attributes `json:"attributes"`
}

type Response struct {
	Header Header `json:"header"`
	Data   []User `json:"data,omitempty"`
}

func ReadResponse(rawResp string) (Response, error) {
	resp := Response{}
	if err := json.Unmarshal([]byte(rawResp), &resp); err != nil {
		return Response{}, fmt.Errorf("JSON unmarshal: %w", err)
	}
	return resp, nil
}

func main() {
	user := User{
		Type: "user",
		Id:   100,
		Attributes: Attributes{
			Email:      "bob@yandex.ru",
			ArticleIds: []int{10, 11, 12},
		},
	}
	response := Response{
		Header: Header{
			Code:    0,
			Message: "",
		},
		Data: []User{user},
	}
	jsData, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("response: %v", string(jsData))

	resp, err := ReadResponse(rawResp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v+\n", resp)
}
