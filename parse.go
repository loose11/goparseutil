/*
Package goparseutil implements an abstract version of the parse android sdk

*/

package goparseutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

const objecturl string = "https://api.parse.com/1/classes"
const fileurl string = "https://api.parse.com/1/files"

var applicationId string
var restAPIKey string

type ParseObject struct {
	name   string
	values map[string]interface{}
}

type ParseFile struct {
	filename  string
	extension string
	content   []byte
}

type ParseFileResponse struct {
	Name string
	Url  string
}

func Initialize(appId, restKey string) {
	applicationId = appId
	restAPIKey = restKey
}

func NewParseObject(name string) *ParseObject {
	values := make(map[string]interface{})
	return &ParseObject{name, values}
}

func (p *ParseObject) Add(key string, value interface{}) {
	p.values[key] = value
}

func (p *ParseObject) AddFile(key string, value string) {
	p.values[key] = map[string]string{
		"name":   value,
		"__type": "File",
	}
}

func (p *ParseObject) Save() {
	json, _ := json.Marshal(&p.values)

	req, err := http.NewRequest("POST", objecturl+"/"+p.name, bytes.NewBuffer(json))
	req.Header.Set("X-Parse-Application-Id", applicationId)
	req.Header.Set("X-Parse-REST-API-Key", restAPIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func NewParseFile(filename string) *ParseFile {
	extension := filepath.Ext(filename)
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	return &ParseFile{filename, extension, b}
}

func (f *ParseFile) Save() string {
	fmt.Println(fileurl + "/" + f.filename)
	req, err := http.NewRequest("POST", fileurl+"/"+f.filename, bytes.NewBuffer(f.content))
	req.Header.Set("X-Parse-Application-Id", applicationId)
	req.Header.Set("X-Parse-REST-API-Key", restAPIKey)
	req.Header.Set("Content-Type", GetMimeType(f.extension))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("response Body:", string(body))

	var response ParseFileResponse

	err = json.Unmarshal(body, &response)

	if err != nil {
		panic(err)
	}

	return response.Name
}
