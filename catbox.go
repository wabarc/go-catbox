// Copyright 2021 Wayback Archiver. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package catbox // import "github.com/wabarc/go-catbox"

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
	"github.com/wabarc/helper"
)

const (
	ENDPOINT = "https://catbox.moe/user/api.php"
)

type Catbox struct {
	Client   *http.Client
	Userhash string
}

func New(client *http.Client) *Catbox {
	if client == nil {
		client = &http.Client{
			Timeout: 30 * time.Second,
		}
	}

	return &Catbox{
		Client: client,
	}
}

func (cat *Catbox) Upload(path string) (string, error) {
	switch {
	case helper.IsURL(path):
		return cat.urlUpload(path)
	case helper.Exists(path):
		return cat.fileUpload(path)
	default:
		return "", errors.New(`path invalid`)
	}
}

func (cat *Catbox) fileUpload(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if size := helper.FileSize(path); size > 209715200 {
		return "", fmt.Errorf("File too large, size: %d MB", size/1024/1024)
	}

	r, w := io.Pipe()
	m := multipart.NewWriter(w)

	go func() {
		defer w.Close()
		defer m.Close()

		m.WriteField("reqtype", "fileupload")
		m.WriteField("userhash", cat.Userhash)
		part, err := m.CreateFormFile("fileToUpload", filepath.Base(file.Name()))
		if err != nil {
			return
		}

		if _, err = io.Copy(part, file); err != nil {
			return
		}
	}()

	req, _ := http.NewRequest(http.MethodPost, ENDPOINT, r)
	req.Header.Add("Content-Type", m.FormDataContentType())

	resp, err := cat.Client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (cat *Catbox) urlUpload(url string) (string, error) {
	b := new(bytes.Buffer)
	w := multipart.NewWriter(b)
	w.WriteField("reqtype", "urlupload")
	w.WriteField("userhash", cat.Userhash)
	w.WriteField("url", url)

	req, _ := http.NewRequest(http.MethodPost, ENDPOINT, b)
	req.Header.Add("Content-Type", w.FormDataContentType())

	resp, err := cat.Client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (cat *Catbox) Delete(files ...string) error {
	// TODO
	return nil
}
