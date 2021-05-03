package data

import (
	"encoding/json"
	"io/ioutil"

	"github.com/TeamEndAllReality/cav2"
)

//Side client or server
type Side string

const (
	//CLIENT SIDE
	CLIENT Side = "CLIENT"
	//BOTH SIDES
	BOTH = "BOTH"
	//SERVER SIDE
	SERVER = "SERVER"
)

//RemoteConfig Represents a remote Repo config
type RemoteConfig struct {
	GameVersion string
	Modloader   string
	URL         string
	Sided       Side
}

//Mod mod item
type Mod struct {
	ID    int
	Sided Side
}

//ModLists mod sides modlist
type ModLists []Mod

//GetRelevent gets relevent for this side
func (c ModLists) GetRelevent(side Side) []int {
	mods := []int{}
	for _, Mod := range c {
		if Mod.Sided == side || Mod.Sided == BOTH {
			mods = append(mods, Mod.ID)
		}
	}
	return mods
}

//RemoteDownload remote download repo
type RemoteDownload struct {
	Mods ModLists
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
	resp, _ := cav2.DoHTTPRequest(rec.URL)
	dl := &RemoteDownload{}
	json.Unmarshal(cav2.ResponseToBytes(resp), dl)
	return dl
}
