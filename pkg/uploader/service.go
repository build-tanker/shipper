package uploader

import (
	"errors"

	"source.golabs.io/core/shipper/pkg/appcontext"
)

type Service interface {
	Install() error
	Uninstall() error
	Upload(bundle string, file string) error
}

type service struct {
	ctx *appcontext.AppContext
}

func NewUploader(ctx *appcontext.AppContext) Service {
	return &service{
		ctx: ctx,
	}
}

func (s *service) Install() error {
	log := s.ctx.GetLogger()
	conf := s.ctx.GetConfig()

	if conf.IsMissing() == false {
		log.Fatalln("Install failed. It seems you have a non-empty config at $HOME/.shipper.toml")
		return errors.New("Non empty config already present")
	}

	return nil
}

func (s *service) Uninstall() error {
	log := s.ctx.GetLogger()
	conf := s.ctx.GetConfig()

	if conf.IsMissing() {
		log.Fatalln("Uninstall failed. It seems you don't have a valid config file")
		return errors.New("No config file found")
	}

	return nil
}

func (s *service) Upload(bundle string, file string) error {
	log := s.ctx.GetLogger()
	conf := s.ctx.GetConfig()

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
