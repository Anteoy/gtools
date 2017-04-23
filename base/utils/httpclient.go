package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	//"encoding/json"
	"io/ioutil"
	"net/http"
)

func HttpGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("httpget err=%v\n", err)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("httpget err=%v\n", err)
		return nil, err
	}
	return body, nil
}

//body_type := "application/json;charset=utf-8"
//body_type := "application/x-www-form-urlencoded"
func HttpPost(url string, bodyType string, req *bytes.Buffer) ([]byte, error) {
	resp, err := http.Post(url, bodyType, req)
	if err != nil {
		fmt.Printf("httppost err=%v\n", err)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("httppost err=%v\n", err)
		return nil, err
	}
	return body, nil
}

type UserLogin struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
type IdT struct {
	Id int64 `json:"id"`
}

func main() {
	/*url := "http://localhost:14500/Oper/login"*/
	body_type := "application/json;charset=utf-8"
	url := "http://localhost:9797/consumer/get_studentapp_by_id"
	// body_type := "application/x-www-form-urlencoded"
	// body, _ := HttpPost(url, body_type, bytes.NewBuffer([]byte(`{"id":1`)))
	// fmt.Printf("post body=%s\n", body)
	// fmt.Println("post body=%s\n", string(body))

	userLogin := IdT{Id: 1}
	fmt.Printf("post body=%v\n", userLogin)
	req, _ := json.Marshal(userLogin)
	body, _ := HttpPost(url, body_type, bytes.NewBuffer(req))
	//body, _ := HttpPost(url, body_type, bytes.NewBuffer([]byte(`{"username":"admin","password":"123456"}`)))
	//slog.Debugf("post body=%s\n", body)
	fmt.Printf("post body=%s\n", body)
	fmt.Println("post body=%s\n", string(body))
}
