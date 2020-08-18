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
	apiKey = "kVcnfD9iW2XVZSMaLMrtLYIz"
	apiSecret = "O9o1O213UgG5LFn0bDGNtoRN3VWl2du6"
	client = voice.NewVoiceClient(apiKey, apiSecret)
}

func txtToVoice(text string, fileName string) {
	bts, err := client.TextToSpeech("你好")
	if err != nil {
		log.Fatal(err)
	}
	writeBytesToFile(bts, fileName)
}

func writeBytesToFile(voiceByteArray []byte, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Println(err)
	}
	file.Write(voiceByteArray)
}

func SpeechToText() {
	if err := client.Auth(); err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile("F:\\go_demo\\backend\\gin_demo\\voice_demo\\16k.wav", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	rs, err := client.SpeechToText(
		f,
		voice.Format("wav"),
		voice.Channel(1),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rs)
}
func main() {
	//SpeechToText()
	txtToVoice("dd", "result.wav")
}
