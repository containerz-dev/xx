// Copyright 2020 The containerz Authors.
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/davecgh/go-spew/spew"

	"github.com/containerz-dev/xx/internal/cmd/xcode-release/pkg/xcoderelease"
)

var rootURI = &url.URL{
	Scheme: "https",
	Host:   "xcodereleases.com",
	Path:   "data.json",
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rootURI.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	dec := json.NewDecoder(bytes.NewReader(data))
	var xrs xcoderelease.Xcodereleases
	if err := dec.Decode(&xrs); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("xrs: %s\n", spew.Sdump(xrs))
}
