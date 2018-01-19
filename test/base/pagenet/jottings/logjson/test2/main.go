package main

import (
	"fmt"
	"regexp"
	"encoding/json"
	"strings"
)

func main(){
	message := "{\"time\":1515660798096004704,\"lev\":\"Dbg\",\"msg\":\"Consumed message Value dynamic==={\"room_id\":2167,\"role\":\"teacher\",\"person_id\":9,\"timestamp\":1515660883,\"video_upload_speed\":\"76kB/s\",\"audio_upload_speed\":\"5kB/s\",\"download_speed\":\"\",\"current_status\":\"\",\"ip\":\"\"}\n\",\"path\":\"rabbit_consumer.go:55\",\"func\":\"main.(*ReceiverA).OnReceive\"}\n";
	message = strings.Replace(message,"\n","",-1)
	message = strings.Replace(message,"\\","",-1)
	fmt.Println(fmt.Sprintf("%s",message))
	re, _ := regexp.Compile("msg\":\".*,\"path")
	match_all := re.FindAll([]byte(message), -1)
	msg_col := string(match_all[0])
	msg_col_change := msg_col[6:len(msg_col)-7]
	msg_col_change = strings.Replace(msg_col_change,"\"","\\\"",-1)
	msg_col_new := "msg\":\""+ msg_col_change + "\",\"path"
	message = strings.Replace(message,msg_col,msg_col_new,-1)
	b:= map[string]interface{}{}
	err:= json.Unmarshal([]byte(message),&b)
	fmt.Println(err,b, message)
}
