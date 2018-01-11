package main

import (
	"encoding/json"
	"fmt"
	//"strings"
	"regexp"
	"strings"
)

func main(){
	a := "{\"time\":1515660798096004704,\"lev\":\"Dbg\",\"msg\":\"Consumed message Value dynamic==={\"room_id\":2167,\"role\":\"teacher\",\"person_id\":9,\"timestamp\":1515660883,\"video_upload_speed\":\"76kB/s\",\"audio_upload_speed\":\"5kB/s\",\"download_speed\":\"\",\"current_status\":\"\",\"ip\":\"\"}\n\",\"path\":\"rabbit_consumer.go:55\",\"func\":\"main.(*ReceiverA).OnReceive\"}\n";
	//a := "msgdksjfldsfjoiejfl\n";
	bb:= "abcde"
	//must a =
	a = strings.Replace(a,"\n","",-1)
	a = strings.Replace(a,"\\","",-1)
	fmt.Println(fmt.Sprintf("%s",a))
	//re, _ := regexp.Compile("msg\":\".*,\"path")
	re, _ := regexp.Compile("\":.*?\"")

	reb, _ := regexp.Compile("^a.*$")
	//a = re.ReplaceAllString(a, "")
	submatchall := re.FindAll([]byte(a), -1)
	submatchallb := reb.FindAll([]byte(bb), -1)
	fmt.Println(string(submatchall[0]))
	b:= map[string]interface{}{}
	err:= json.Unmarshal([]byte(a),&b)
	fmt.Println(err,b,submatchall,submatchallb)
}
