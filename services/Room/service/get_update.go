package service

import (
	"net/http"
	"net/url"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"YouHome_server/services/Room/entities"
)

func getRooms(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	m, _ := url.ParseQuery(r.URL.RawQuery)
	userId := m["userId"][0]
	data := entities.GetRooms(userId)
	fmt.Fprint(w, data)
	return
}
func nameHandler(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	if r.Method == "GET"{
		getName(w,r);
	}else{
		updateName(w,r)
	}
	return
}
func getName(w http.ResponseWriter,r*http.Request){
	r.ParseForm()
	m, _ := url.ParseQuery(r.URL.RawQuery)
	sceneId := m["sceneId"][0]
	data := entities.GetRoomName(sceneId)
	fmt.Fprint(w, data)
	return
}
func updateName(w http.ResponseWriter,r*http.Request){
	r.ParseForm()
	var user map[string]interface{}
	data, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(data, &user)
	fmt.Println(string(data))

	sceneId := user["sceneId"].(string)
	sceneName := user["sceneName"].(string)
	result := entities.UpdateRoomName(sceneName,sceneId)
	fmt.Fprint(w,result)
}

