package downloader

import (
	"strconv"

	"github.com/TeamEndAllReality/cav2"
)

func init() {
	proced = make(map[int]bool)
}

var (
	proced map[int]bool
)

//DownloadCurseFile Downloads a file off of curseforge
func DownloadCurseFile(addn *cav2.File, name string) {
	h, _ := cav2.GetFileHash("mods/" + name + ".jar")
	if int64(h) != addn.PackageFingerprint {
		println("Downloading: " + name)
		inf, _ := cav2.DoHTTPRequest(addn.DownloadURL)
		cav2.WriteBytesToFile(cav2.ResponseToBytes(inf), "mods/"+name+".jar")
		println("Downloaded: " + name)
	} else {
		println("Hash Matched (Not Downloading): " + name)
	}
}

//ProcCurseAddon Queues a file for downloading with dependencies from curse
func ProcCurseAddon(addn *cav2.Addon, gv string) {
	if _, ok := proced[addn.ID]; !ok {
		proced[addn.ID] = true
		for _, adl := range addn.GameVersionLatestFiles {
			if adl.GameVersion == gv {
				url, _ := cav2.GetAddonFile(addn.ID, adl.ProjectFileID)
				for _, depend := range url.Dependencies {
					if depend.Type == 1 {
						dep, _ := cav2.GetAddon(strconv.Itoa(depend.AddonID))
						ProcCurseAddon(dep, gv)
					}
				}
				DownloadCurseFile(url, addn.Name)
				break
			}
		}
	}
}
