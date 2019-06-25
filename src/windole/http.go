package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		bs, err := ioutil.ReadFile("ui.html")
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(200)
		w.Write(bs)
	})
}

func init() {
	http.HandleFunc("/act/open_system_proxy", func(w http.ResponseWriter, r *http.Request) {
		err := exec.Command("reg", "add", "HKCU\\Software\\Microsoft\\Windows\\CurrentVersion\\Internet Settings", "/v", "ProxyEnable", "/t", "REG_DWORD", "/d", "1", "/f").Run()
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(200)
		w.Write([]byte(ok))
	})
}

func init() {
	http.HandleFunc("/act/close_system_proxy", func(w http.ResponseWriter, r *http.Request) {
		err := exec.Command("reg", "add", "HKCU\\Software\\Microsoft\\Windows\\CurrentVersion\\Internet Settings", "/v", "ProxyEnable", "/t", "REG_DWORD", "/d", "0", "/f").Run()
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(200)
		w.Write([]byte(ok))
	})
}

func init() {
	http.HandleFunc("/act/restart_indole", func(w http.ResponseWriter, r *http.Request) {
		indole := r.URL.Query().Get("indole")
		config := r.URL.Query().Get("config")
		log.Println("1")
		err := restart(indole, config)
		log.Println("2")
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(200)
		w.Write([]byte(ok))
	})
}

func init() {
	http.HandleFunc("/act/exit", func(w http.ResponseWriter, r *http.Request) {
		cmd.Process.Kill()
		os.Exit(0)
		w.WriteHeader(200)
		w.Write([]byte(ok))
	})
}

const ok = "OK"
