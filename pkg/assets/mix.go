package assets

import (
	"cms/pkg/logger"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Manifest() (manifests map[string]string) {
	curPath, _ := os.Getwd()
	bytes, err := ioutil.ReadFile(curPath + "/public/mix-manifest.json")
	if err != nil {
		fmt.Println(err)
		logger.Error("mix-manifest.json文件读取失败")
	}
	manifests = make(map[string]string)
	err = json.Unmarshal(bytes, &manifests)
	if err != nil {
		logger.Error("manifest.json解析失败")
	}
	for k, v := range manifests {
		if strings.HasSuffix(k, ".js") && k != "manifest.js" {
			nk := strings.Split(k, "-")
			nk[0] = strings.TrimLeft(nk[0], "/js/")
			manifests[UcFirst(nk[0])] = v
			delete(manifests, k)
		}
	}
	return manifests
}
