package uploader

import (
	"errors"

	"source.golabs.io/core/shipper/pkg/appcontext"
)

type Uploader interface {
	Install() error
	Uninstall() error
	Upload(bundle string, file string) error
}

type uploader struct {
	ctx *appcontext.AppContext
}

func NewUploader(ctx *appcontext.AppContext) Uploader {
	return &uploader{
		ctx: ctx,
	}
}

func (u *uploader) Install() error {
	return nil
}

func (u *uploader) Uninstall() error {
	return nil
}

func (u *uploader) Upload(bundle string, file string) error {
	log := u.ctx.GetLogger()
	conf := u.ctx.GetConfig()

	if conf.IsMissing() {
		log.Fatalln("It seems you have an empty config. Please run *shipper install* first")
		return errors.New("Need to install shipper first")
	}

	if bundle == "" {
		log.Fatalln("Please enter the bundleID that you're uploading for")
		return errors.New("BundleID missing")
	}

	if file == "" {
		log.Fatalln("Please enter the path of the file that you would like to upload")
		return errors.New("File path is missing")
	}

	return nil
}

// 	toUpload, err := os.Open(file)
// 	if err != nil {
// 		return err
// 	}
// 	defer toUpload.Close()

// 	serverURL := fmt.Sprintf("%s?key=%s&bundle=%s&file=%s", config.UploadServer(), key, bundle, file)
// 	logger.Infof(serverURL)

// 	response, err := http.Post(serverURL, "binary/octet-stream", toUpload)
// 	if err != nil {
// 		return err
// 	}
// 	defer response.Body.Close()

// 	message, _ := ioutil.ReadAll(response.Body)
// 	logger.Infoln(string(message))

// 	return nil
// }
