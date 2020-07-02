package main

import (
	"github.com/TeamEndAllReality/cav2"
	"github.com/TeamEndAllReality/mcpaku/downloader"
)

const (
	conf = "config.json"
)

func main() {
	cd := GetConfig(conf)
	data := GetDownloads(cd)
	addns, _ := cav2.GetAddons(data.Mods)
	for _, addn := range addns {
		downloader.ProcCurseAddon(addn, data.GameVersion)
	}
}
