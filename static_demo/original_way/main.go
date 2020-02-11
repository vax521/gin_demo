package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

/**
  自己造轮子方式
*/
var extensionToContentType = map[string]string{
	".html": "text/html; charset=utf-8",
	".css":  "text/css; charset=utf-8",
	".js":   "application/javascript",
	".xml":  "text/xml; charset=utf-8",
	".jpg":  "image/jpeg",
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	path := "./static_demo" + r.URL.Path
	fmt.Println(path)
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	d, err := file.Stat()
	if err != nil {
		log.Fatal(err)
		return
	}
	if d.IsDir() {
		DirList(w, r, file)
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		return
	}
	ext := filepath.Ext(path)
	if contentType := extensionToContentType[ext]; contentType != "" {
		w.Header().Set("Content-Type", contentType)
	}

	w.Header().Set("Content-Length", strconv.FormatInt(d.Size(), 10))
	w.Write(data)
}

func DirList(w http.ResponseWriter, r *http.Request, f http.File) {
	dirs, err := f.Readdir(-1)
	if err != nil {
		return
	}
	sort.Slice(dirs, func(i, j int) bool { return dirs[i].Name() < dirs[j].Name() })

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<pre>\n")
	for _, d := range dirs {
		name := d.Name()
		if d.IsDir() {
			name += "/"
		}
		url := url.URL{Path: name}
		fmt.Fprintf(w, "<a href=\"%s\">%s</a>\n", url.String(), name)
	}
	fmt.Fprintf(w, "</pre>\n")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/static/", fileHandler)
	server := &http.Server{
		Addr:    ":8090",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
