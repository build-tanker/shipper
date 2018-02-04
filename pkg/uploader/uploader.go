package uploader

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"source.golabs.io/core/shipper/pkg/config"
	"source.golabs.io/core/shipper/pkg/logger"
)

// Upload file for a specific buncle with an access key
func Upload(key string, bundle string, file string) error {
	logger.Infof("Key:%s Bundle:%s File:%s", key, bundle, file)

	toUpload, err := os.Open(file)
	if err != nil {
		return err
	}
	defer toUpload.Close()

	serverURL := fmt.Sprintf("%s?key=%s&bundle=%s&file=%s", config.UploadServer(), key, bundle, file)
	logger.Infof(serverURL)

	response, err := http.Post(serverURL, "binary/octet-stream", toUpload)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	message, _ := ioutil.ReadAll(response.Body)
	logger.Infoln(string(message))

	return nil
}
