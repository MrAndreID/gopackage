package gopackage

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/h2non/filetype"
	"github.com/sirupsen/logrus"
)

type SeaweedFSData struct {
	URL string
}

func NewSeaweedFS(host string, port string, ssl bool) (*SeaweedFSData, error) {
	var (
		tag    string        = "GoPackage.SeaweedFS.NewSeaweedFS."
		client *resty.Client = resty.New()
		url    string
	)

	if ssl {
		url = "https://" + host + ":" + port
	} else {
		url = "http://" + host + ":" + port
	}

	resp, err := client.R().EnableTrace().Get(url)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"tag":   tag + "01",
			"error": err.Error(),
		}).Error("failed to connect seaweedfs")

		return nil, err
	}

	if resp.StatusCode() != 200 {
		logrus.WithFields(logrus.Fields{
			"tag": tag + "02",
		}).Error("failed to connect seaweedfs")

		return nil, errors.New("failed to connect - seaweedfs")
	}

	return &SeaweedFSData{
		URL: url,
	}, nil
}

func (sfs *SeaweedFSData) Upload(base64File string) (string, error) {
	var (
		tag       string        = "GoPackage.SeaweedFS.Upload."
		client    *resty.Client = resty.New()
		publicURL string
		err       error
	)

	resp, err := client.R().EnableTrace().Get(sfs.URL + "/dir/assign")

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"tag":   tag + "01",
			"error": err.Error(),
		}).Error("failed to upload - seaweedfs")

		return publicURL, err
	}

	if resp.StatusCode() != 200 {
		logrus.WithFields(logrus.Fields{
			"tag": tag + "02",
		}).Error("failed to upload - seaweedfs")

		return publicURL, errors.New("failed to upload - seaweedfs")
	}

	responseBody := struct {
		FID       string `json:"fid"`
		PublicURL string `json:"publicUrl"`
	}{}

	err = json.Unmarshal(resp.Body(), &responseBody)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"tag":   tag + "03",
			"error": err.Error(),
		}).Error("failed to upload - seaweedfs")

		return publicURL, errors.New("failed to upload - seaweedfs")
	}

	if !strings.Contains(responseBody.PublicURL, "http") {
		responseBody.PublicURL = strings.Split(sfs.URL, "/")[0] + "//" + responseBody.PublicURL
	}

	publicURL = responseBody.PublicURL + "/" + responseBody.FID

	bytesFile, err := base64.StdEncoding.DecodeString(base64File)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"tag":   tag + "04",
			"error": err.Error(),
		}).Error("failed to upload - seaweedfs")

		return publicURL, errors.New("failed to upload - seaweedfs")
	}

	kind, err := filetype.Match(bytesFile)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"tag":   tag + "05",
			"error": err.Error(),
		}).Error("failed to upload - seaweedfs")

		return publicURL, errors.New("failed to upload - seaweedfs")
	}

	fileName := responseBody.FID + "." + kind.Extension

	resp, err = client.R().SetFileReader("file", fileName, bytes.NewReader(bytesFile)).Post(publicURL)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"tag":   tag + "06",
			"error": err.Error(),
		}).Error("failed to upload - seaweedfs")

		return publicURL, err
	}

	if resp.StatusCode() != 201 {
		logrus.WithFields(logrus.Fields{
			"tag": tag + "07",
		}).Error("failed to upload - seaweedfs")

		return publicURL, errors.New("failed to upload - seaweedfs")
	}

	return publicURL, nil
}

func (sfs *SeaweedFSData) Download(publicURL string) (string, error) {
	var (
		tag        string        = "GoPackage.SeaweedFS.Download."
		client     *resty.Client = resty.New()
		base64File string
		err        error
	)

	resp, err := client.R().EnableTrace().Get(publicURL)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"tag":   tag + "01",
			"error": err.Error(),
		}).Error("failed to download - seaweedfs")

		return base64File, err
	}

	if resp.StatusCode() != 200 {
		logrus.WithFields(logrus.Fields{
			"tag": tag + "02",
		}).Error("failed to download - seaweedfs")

		return base64File, errors.New("failed to download - seaweedfs")
	}

	base64File = base64.StdEncoding.EncodeToString(resp.Body())

	return base64File, nil
}

func (sfs *SeaweedFSData) Delete(publicURL string) error {
	var (
		tag    string        = "GoPackage.SeaweedFS.Delete."
		client *resty.Client = resty.New()
		err    error
	)

	resp, err := client.R().EnableTrace().Delete(publicURL)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"tag":   tag + "01",
			"error": err.Error(),
		}).Error("failed to delete - seaweedfs")

		return err
	}

	if resp.StatusCode() != 202 {
		logrus.WithFields(logrus.Fields{
			"tag": tag + "02",
		}).Error("failed to delete - seaweedfs")

		return errors.New("failed to delete - seaweedfs")
	}

	return nil
}
