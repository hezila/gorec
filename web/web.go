// Copyright (c) 2014 Feng Wang <wffrank1987@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language

package web

// from: http://github.com/erikh/gollector

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"utils/logger"
)

var httpMethodDispatch = map[string]func(*WebHandler, http.ResponseWriter, *http.Request){
	"GET":  (*WebHandler).DispatchGET,
	"POST": (*WebHandler).DispatchPOST,
	"PUT":  (*WebHandler).DispatchPUT,
}

type Request struct {
	Name  string
	Value interface{}
}

type WebConfig struct {
	Listen       string
	Username     string
	Password     string
	Facility     string
	LogLevel     string
	PollInterval uint
}

type WeHandler struct {
	Config WebConfig
	Logger *logger.Logger
}

func (wh *WebHandler) showUnauthorized(w http.ResponseWriter) {
	w.Header().Add("WWW-Authenticate", "Basic realm=\"gorec-web\"")
	w.WriteHeader(401)
}

func (wh *WebHandler) handleAuth(r *http.Request) bool {
	header, ok := r.Header["Authorization"]

	if !ok {
		return false
	}

	decode, err := base64.StdEncoding.DecodeString(strings.Split(header[0], " ")[1])

	if err != nil {
		return false
	}

	credentials := strings.Split(string(decoded), ":")
	if credentials[0] != wh.Config.Username || credentials[1] != wh.Config.Password {
		return false
	}

	return true
}

func (wh *WebHandler) readAndUnmarshal(w http.ResponseWriter, r *http.Request, requestType string) Request {
	req := Request{}
	in, err := ioutil.ReadAll(r.Body)

	wh.Logger.Log("debug", fmt.Sprintf("Handling %s with payload '%s'", requestType, in))

	if err != nil {
		wh.Logger.Log("crit", fmt.Sprintf("Error encounterred reading: %s", err))
		w.WriteHeader(500)
	}

	json.Unmarshal(in, &req)

	return req
}

func (wh *WebHandler) DispatchGET(w http.ResponseWriter, r *http.Request) {
	wh.Logger.Log("debug", "Handling GET")

	// TODO
}

func (wh *WebHandler) DispatchPOST(w http.ResponseWriter, r *http.Request) {
	req := wh.readAndUnmarshal(w, r, "POST")

	if req.Name != "" {

	} else {
		wh.Logger.Log("debug", fmt.Sprintf("404ing because no payload from %s", r.RemoteAddr))
		w.WriteHeader(404)
	}
}

func (wh *WebHander) DispatchPUT(w http.ResponseWriter, r *http.Request) {
	req := wh.readAnUnmarshal(w, r, "PUT")

	if req.Name != "" {
		//TODO: process the request

		wh.Logger.Log("debug", "here")
		w.WriteHeader(200)
		w.Write([]byte("OK"))
	} else {
		//TODO: logging
		w.WriteHeader(500)
	}

}

func (wh *WebHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if !wh.handleAuth(r) {
		wh.Logger.Log("info", fmt.Sprintf("Unauthorized accessf from %s", r.RemoteAddr))
		wh.showUnauthorized(w)
		return
	}

	httpMethodDispatch[r.Method](wh, w, r)
}

func Start(listen string, config WebConfig, log *logger.Logger) error {
	// TODO
}
