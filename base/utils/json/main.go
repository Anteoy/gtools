package main

import (
	"fmt"
	"encoding/json"
)

var json_str string = ""

func RecResultJsonToPlain() {
	var recResult string
	var dat map[string]interface{}
	json.Unmarshal([]byte(json_str), &dat)

	if v, ok := dat["ws"]; ok {
		ws := v.([]interface{})
		for i, wsItem := range ws {
			wsMap := wsItem.(map[string]interface{})
			if vCw, ok := wsMap["cw"]; ok {
				cw := vCw.([]interface{})
				for i, cwItem := range cw {
					cwItemMap := cwItem.(map[string]interface{})
					if w, ok := cwItemMap["w"]; ok {
						recResult = recResult + w.(string)
					}
				}
			}
		}
	}
	fmt.Println(recResult)
}