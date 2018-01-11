package main

import (
	"fmt"
	"regexp"
	"encoding/json"
)

func main(){
	a := "{\"time\":1515660798096004704,\"lev\":\"Dbg\",\"msg\":\"Consumed message Value dynamic==={\"room_id\":2167,\"role\":\"teacher\",\"person_id\":9,\"timestamp\":1515660883,\"video_upload_speed\":\"76kB/s\",\"audio_upload_speed\":\"5kB/s\",\"download_speed\":\"\",\"current_status\":\"\",\"ip\":\"\"}\",\"path\":\"rabbit_consumer.go:55\",\"func\":\"main.(*ReceiverA).OnReceive\"}";
	re, _ := regexp.Compile("^msg.*$")
	submatchall := re.FindAll([]byte(a), -1)
	b:= map[string]interface{}{}
	err:= json.Unmarshal([]byte(a),&b)
	fmt.Println(err,b,submatchall)
}
