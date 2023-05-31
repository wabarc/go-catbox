// Copyright 2021 Wayback Archiver. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package catbox // import "github.com/wabarc/go-catbox"

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFileUpload(t *testing.T) {
	content := make([]byte, 5000)
	tmpfile, err := ioutil.TempFile("", "go-catbox-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(content); err != nil {
		t.Fatal(err)
	}

	if _, err := New(nil).Upload(tmpfile.Name()); err != nil {
		t.Fatal(err)
	}
}

func TestRawUpload(t *testing.T) {
	content := make([]byte, 5000)
	tmpfile, err := ioutil.TempFile("", "go-catbox-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	b, err := ioutil.ReadFile(tmpfile.Name())

	if err != nil {
		t.Fatal(err)
	}

	if _, err := tmpfile.Write(content); err != nil {
		t.Fatal(err)
	}

	if _, err := New(nil).Upload(b, "test"); err != nil {
		t.Fatal(err)
	}
}

func TestURLUpload(t *testing.T) {
	url := "https://www.gstatic.com/webp/gallery/1.webp"
	if _, err := New(nil).Upload(url); err != nil {
		t.Fatal(err)
	}
}
