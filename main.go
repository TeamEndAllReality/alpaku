package main

import (
	"github.com/TeamEndAllReality/alpaku/downloader"
	"github.com/TeamEndAllReality/alpaku/global"
	"github.com/TeamEndAllReality/cav2"
)

const (
	conf = "config.json"
)

func main() {
	cd := GetConfig(conf)
	data := GetDownloads(cd)

	addns, _ := cav2.GetAddons(data.Mods.GetRelevent(cd.Sided))
	for _, addn := range addns {
		global.WG.Add(1)
		go downloader.ProcCurseAddon(addn, cd.GameVersion)
	}
	global.WG.Wait()
}
