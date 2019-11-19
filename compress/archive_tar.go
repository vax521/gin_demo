package main

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//创建一个缓冲区
	buff := new(bytes.Buffer)
	//创建一个新的tar存档
	tarFile := tar.NewWriter(buff)
	var files = []struct {
		name, body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling license."},
	}
	for _, file := range files {
		hdr := &tar.Header{
			Name: file.name,
			Mode: 0600,
			Size: int64(len(file.body)),
		}
		if err := tarFile.WriteHeader(hdr); err != nil {
			log.Fatalln(err)
		}
		if _, err := tarFile.Write([]byte(file.body)); err != nil {
			log.Fatalln(err)
		}
	}
	// 确保在Close时检查错误。
	if err := tarFile.Close(); err != nil {
		log.Fatalln(err)
	}

	//打开tar文档
	r := bytes.NewReader(buff.Bytes())
	tr := tar.NewReader(r)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Contents of %s:\n", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatalln(err)
		}
		fmt.Println()
	}

}
