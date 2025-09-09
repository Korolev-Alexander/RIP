package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"smartdevices/data" // ! Имя должно совпадать с module в go.mod
	"strconv"
	"strings"
)

var (
	tmplDevices       = template.Must(template.ParseFiles("templates/layout.html", "templates/devices.html"))
	tmplDeviceDetail  = template.Must(template.ParseFiles("templates/layout.html", "templates/device_detail.html"))
	tmplRequestDetail = template.Must(template.ParseFiles("templates/layout.html", "templates/request_detail.html"))
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/devices", devicesHandler)
	http.HandleFunc("/devices/", deviceDetailHandler)
	http.HandleFunc("/requests/", requestDetailHandler)
	http.ListenAndServe(":8080", nil)
}

func devicesHandler(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	devices := data.Devices
	var filtered []data.Device
	if search != "" {
		for _, d := range devices {
			if strings.Contains(strings.ToLower(d.Name), strings.ToLower(search)) ||
				strings.Contains(strings.ToLower(d.Model), strings.ToLower(search)) {
				filtered = append(filtered, d)
			}
		}
	} else {
		filtered = devices
	}
	tmplDevices.ExecuteTemplate(w, "layout.html", map[string]interface{}{
		"Devices": filtered,
		"Search":  search,
	})
}

func deviceDetailHandler(w http.ResponseWriter, r *http.Request) {
	idStr := filepath.Base(r.URL.Path)
	id, _ := strconv.Atoi(idStr)
	var device *data.Device
	for _, d := range data.Devices {
		if d.ID == id {
			device = &d
			break
		}
	}
	if device == nil {
		http.NotFound(w, r)
		return
	}
	tmplDeviceDetail.ExecuteTemplate(w, "layout.html", map[string]interface{}{
		"Device": device,
	})
}

func requestDetailHandler(w http.ResponseWriter, r *http.Request) {
	idStr := filepath.Base(r.URL.Path)
	id, _ := strconv.Atoi(idStr)
	var device *data.Device
	for _, d := range data.Devices {
		if d.ID == id {
			device = &d
			break
		}
	}
	if device == nil {
		http.NotFound(w, r)
		return
	}
	tmplRequestDetail.ExecuteTemplate(w, "layout.html", map[string]interface{}{
		"Device": device,
	})
}
