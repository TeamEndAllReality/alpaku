package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/TeamEndAllReality/cav2"
)

//Side - side of client
type Side int

const (
	//CLIENT side
	CLIENT Side = iota
	//SERVER side
	SERVER
	//BOTH sides
	BOTH
)

//RemoteConfig Represents a remote Repo config
type RemoteConfig struct {
	URI   string
	Sided Side
}

//RemoteDownload remote download repo
type RemoteDownload struct {
	GameVersion string
	Mods        []int
}

//GetConfig gets the config json data
func GetConfig(fileName string) *RemoteConfig {
	data, _ := ioutil.ReadFile(fileName)
	conf := &RemoteConfig{}
	json.Unmarshal(data, conf)
	return conf
}

//GetDownloads get remote repo
func GetDownloads(rec *RemoteConfig) *RemoteDownload {
	resp, _ := cav2.DoHTTPRequest(rec.URI)
	dl := &RemoteDownload{}
	json.Unmarshal(cav2.ResponseToBytes(resp), dl)
	return dl
}
