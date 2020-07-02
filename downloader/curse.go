package downloader

import (
	"strconv"
	"sync"

	"github.com/TeamEndAllReality/alpaku/global"
	"github.com/TeamEndAllReality/alpaku/utils"
	"github.com/TeamEndAllReality/cav2"
)

func init() {
	proced = &sync.Map{}
}

var (
	proced *sync.Map
)

//DownloadCurseFile Downloads a file off of curseforge
func DownloadCurseFile(addn *cav2.File, name string) {
	sane := utils.Sanitize(name)
	defer global.WG.Done()
	h, _ := cav2.GetFileHash("mods/" + sane + ".jar")
	if int64(h) != addn.PackageFingerprint {
		println("Downloading: " + sane)
		inf, _ := cav2.DoHTTPRequest(addn.DownloadURL)
		cav2.WriteBytesToFile(cav2.ResponseToBytes(inf), "mods/"+sane+".jar")
		println("Downloaded: " + sane)
	} else {
		println("Hash Matched (Not Downloading): " + sane)
	}
}

//ProcCurseAddon Queues a file for downloading with dependencies from curse
func ProcCurseAddon(addn *cav2.Addon, gv string) {
	defer global.WG.Done()
	if _, ok := proced.LoadOrStore(addn.ID, true); !ok {
		for _, adl := range addn.GameVersionLatestFiles {
			if adl.GameVersion == gv {
				url, _ := cav2.GetAddonFile(addn.ID, adl.ProjectFileID)
				global.WG.Add(1)
				go func() {
					defer global.WG.Done()
					for _, depend := range url.Dependencies {
						if depend.Type == 3 {
							dep, _ := cav2.GetAddon(strconv.Itoa(depend.AddonID))
							global.WG.Add(1)
							go ProcCurseAddon(dep, gv)
						}
					}
				}()
				global.WG.Add(1)
				go DownloadCurseFile(url, addn.Name)
				break
			}
		}
	}
}
