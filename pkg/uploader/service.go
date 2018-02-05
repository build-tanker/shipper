package uploader

import (
	"fmt"
	"os/user"

	"github.com/pkg/errors"

	"source.golabs.io/core/shipper/pkg/appcontext"
	"source.golabs.io/core/shipper/pkg/filesystem"
)

// Service - service to install, uninstall and upload files from shipper
type Service interface {
	Install(server string) error
	Uninstall() error
	Upload(bundle string, file string) error
}

type service struct {
	ctx    *appcontext.AppContext
	client Client
	fs     filesystem.FileSystem
}

// NewService - create a new service to install, uninstall and upload files from shipper
func NewService(ctx *appcontext.AppContext) Service {
	client := NewClient(ctx)
	fs := filesystem.NewFileSystem()
	return &service{
		ctx:    ctx,
		client: client,
		fs:     fs,
	}
}

func (s *service) Install(server string) error {
	log := s.ctx.GetLogger()
	conf := s.ctx.GetConfig()

	if !conf.IsMissing() {
		log.Errorln("Install failed. It seems you have a non-empty config\n$ shipper uninstall")
		return errors.New("Non empty config already present")
	}

	if server == "" {
		log.Errorln("Install failed. Please enter the server that you would like to register with\n$ shipper install --server http://public.betas.in")
		return errors.New("Server flag missing")
	}

	// Get accessKey from client
	accessKey, err := s.client.GetAccessKey(server)
	if err != nil {
		return errors.Wrap(err, "Could not get Access Key")
	}

	// Save config file with accessKey and Server
	err = s.writeConfigFile(server, accessKey)
	if err != nil {
		return errors.Wrap(err, "Could not write config file")
	}

	return nil
}

func (s *service) Uninstall() error {
	log := s.ctx.GetLogger()
	conf := s.ctx.GetConfig()

	if conf.IsMissing() {
		log.Errorln("Uninstall failed. It seems you don't have a valid config file")
		return errors.New("No config file found")
	}

	// Remove accessKey from client
	err := s.client.DeleteAccessKey(conf.Server, conf.AccessKey)
	if err != nil {
		return errors.Wrap(err, "Could not delete access key")
	}
	// Delete config file
	err = s.deleteConfigFile()
	if err != nil {
		return errors.Wrap(err, "Could not delete config file")
	}

	return nil
}

func (s *service) Upload(bundle string, file string) error {
	log := s.ctx.GetLogger()
	conf := s.ctx.GetConfig()

	if conf.IsMissing() {
		log.Errorln("It seems you have an empty config. Please run *shipper install* first")
		return errors.New("Need to install shipper first")
	}

	if bundle == "" {
		log.Errorln("Please enter the bundleID that you're uploading for")
		return errors.New("BundleID missing")
	}

	if file == "" {
		log.Errorln("Please enter the path of the file that you would like to upload")
		return errors.New("File path is missing")
	}

	// Get upload URL from client
	url, err := s.client.GetUploadURL()
	if err != nil {
		return errors.Wrap(err, "Could not get upload URL")
	}

	// Start file upload from filesystem
	err = s.client.UploadFile(url, file)
	if err != nil {
		return errors.Wrap(err, "Could not upload file")
	}
	// On completion tell client file upload is done with url

	return nil
}

func (s *service) writeConfigFile(server, accessKey string) error {
	usr, err := user.Current()
	if err != nil {
		return errors.Wrap(err, "Could not find current user in writeConfigFile")
	}

	configFilePath := usr.HomeDir + "/.shipper.toml"
	configData := `[application]
server = "%s"
accessKey = "%s"`
	data := []byte(fmt.Sprintf(configData, server, accessKey))
	return s.fs.WriteCompleteFileToDisk(configFilePath, data, 0644)
}

func (s *service) deleteConfigFile() error {
	usr, err := user.Current()
	if err != nil {
		return errors.Wrap(err, "Could not find current user in deleteConfigFile")
	}

	configFilePath := usr.HomeDir + "/.shipper.toml"
	return s.fs.DeleteFileFromDisk(configFilePath)
}
