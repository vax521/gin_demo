package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Servers struct {
	XMLName xml.Name  `xml:"servers"`
	Version string    `xml:"version,attr"`
	Svs     []server1 `xml:"server"`
}
type server1 struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIp"`
}

func main() {
	//struct转xml
	v := &Servers{Version: "1"}
	v.Svs = append(v.Svs, server1{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	v.Svs = append(v.Svs, server1{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})

	output, err := xml.MarshalIndent(v, " ", "")
	if err != nil {
		fmt.Printf("error:%v\n", err)
	}
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)

}
