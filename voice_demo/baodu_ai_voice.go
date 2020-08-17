package main

import (
	"fmt"
	"github.com/chenqinghe/baidu-ai-go-sdk/voice"
	"log"
	"os"
)

var apiKey, apiSecret string
var client *voice.VoiceClient

func init() {
	apiKey = "xQr8wlnW9nKwQIrF6t9fwoQR"
	apiSecret = "SYOXr7KskiGKnn0xO5T0SPkQ1xq7qVmA"
	client = voice.NewVoiceClient(apiKey, apiSecret)
}

func txtToVoice(text string, fileName string) {
	bts, err := client.TextToSpeech("你好")
	if err != nil {
		log.Fatal(err)
	}
	if err := writeBytesToFile(bts, fileName); err != nil { //writeBytesToFile需要自己实现
		log.Fatal(err)
	}
}

func writeBytesToFile(voiceByteArray []byte, fileName string) error {
	return nil
}

func SpeechToText() {
	if err := client.Auth(); err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile("16k.pcm", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	rs, err := client.SpeechToText(
		f,
		voice.Format("pcm"),
		voice.Channel(1),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rs)
}
func main() {
	SpeechToText()
}
