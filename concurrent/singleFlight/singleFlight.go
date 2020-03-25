package main

import (
	"errors"
	"fmt"
	"golang.org/x/sync/singleflight"
	"log"
	"sync"
)

var errNews = errors.New("not exit")

func getDataFromCache(key string) (string, error) {
	return "", errNews
}

func getDataFromDB(key string) (string, error) {
	log.Printf("got %s from db", key)
	return "data", nil
}

func getData(key string) (string, error) {
	data, err := getDataFromCache(key)
	if err == errNews {
		data, err := getDataFromDB(key)
		if err != nil {
			log.Fatal(err)
			return "", err
		} else {
			fmt.Println(data)
		}
	} else if err != nil {
		return "", err
	}
	return data, nil
}

var g singleflight.Group

func getDataWithSingleFlight(key string) (string, error) {
	data, err := getDataFromCache(key)
	if err == errNews {
		v, err, _ := g.Do(key, func() (interface{}, error) {
			return getDataFromDB(key)
		})
		if err != nil {
			log.Println(err)
			return "", err
		}

		//TOOD: set cache
		data = v.(string)
	} else if err != nil {
		return "", nil
	}
	return data, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(10)

	//模拟10个并发
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			//data, err := getData("key")
			data, err := getDataWithSingleFlight("key")
			if err != nil {
				log.Print(err)
				return
			}
			log.Println(data)
		}()
	}
	wg.Wait()

}
