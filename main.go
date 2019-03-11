package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func main() {

	var topics = []string{
		"5618c159add4471100150637", // 浴室沉思
		"557ed045e4b0a573eb66b751", // 无用但有趣的冷知识
		"5a82a88df0eddb00179c1df7", // 今日烂梗
		"572c4e31d9595811007a0b6b", // 弱智金句病友会
		"56d177a27cb3331100465f72", // 今日金句
		"5aa21c7ae54af10017dc93f8", // 一个想法不一定对
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	topicdIndex := r1.Intn(len(topics))
	url := "https://app.jike.ruguoapp.com/1.0/squarePosts/list"
	jsonStr := []byte(`{"topicId": "` + topics[topicdIndex] + `", "limit": 20}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		panic(err)
	}

	s2 := rand.NewSource(time.Now().UnixNano())
	r2 := rand.New(s2)

	contentData := dat["data"].([]interface{})
	contentIndex := r2.Intn(len(contentData))
	content := contentData[contentIndex].(map[string]interface{})
	fmt.Println(content["content"].(string))
	fmt.Println("--", (content["topic"].(map[string]interface{})["content"].(string)))
}
