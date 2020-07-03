package main

import (
	"github.com/TeamEndAllReality/alpaku/data"
	"github.com/TeamEndAllReality/alpaku/downloader"
	"github.com/TeamEndAllReality/alpaku/global"
	"github.com/TeamEndAllReality/cav2"
)

const (
	conf = "config.json"
)

func main() {
	cd := data.GetConfig(conf)
	data := data.GetDownloads(cd)
	addns, _ := cav2.GetAddons(data.Mods.GetRelevent(cd.Sided))
	for _, addn := range addns {
		global.WG.Add(1)
		go downloader.ProcCurseAddon(addn, cd.GameVersion /*, cd.Modloader*/)
	}
	global.WG.Wait()
}
