package unsplash

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/probeldev/fastlauncher/model"
)

type unsplashParser struct{}

func GetUnsplashParser() *unsplashParser {
	u := unsplashParser{}

	return &u
}

func (u *unsplashParser) getDataJson(data string) (
	string,
	error,
) {

	data = strings.ReplaceAll(data, ">", ">\n")
	data = strings.ReplaceAll(data, "<", "\n<")
	dataArray := strings.Split(data, "\n")
	for _, data := range dataArray {
		if strings.Contains(data, `"download":`) {
			return data, nil
		}
	}

	return "", errors.New("json not found")

}

func (u *unsplashParser) getJson() ([]byte, error) {
	fn := "unsplashParser:getJson"

	resp, err := http.Get("https://unsplash.com/")
	if err != nil {
		log.Println(fn, err)
		return []byte{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(fn, err)
		return []byte{}, err
	}

	jsonData, err := u.getDataJson(string(body))
	if err != nil {
		log.Println(fn, err)
		return []byte{}, err
	}

	return []byte(jsonData), nil
}

func (u *unsplashParser) GetJsonData() []byte {
	countTry := 0
	for {
		data, err := u.getJson()
		if err != nil {
			if countTry < 5 {
				countTry++
				time.Sleep(30 * time.Second)
			} else {
				countTry = 0
				time.Sleep(10 * time.Minute)
			}
			continue
		}

		return data
	}
}

func (u *unsplashParser) GetImageUrls() (
	[]string,
	error,
) {
	jsonData := u.GetJsonData()
	var downloadURLs []string

	var data model.Root
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		return downloadURLs, err
	}

	photos := data.ReduxInitialState.Entities.Photos
	for _, p := range photos {
		downloadURLs = append(downloadURLs, p.Links.Download)
	}

	if len(downloadURLs) == 0 {
		return downloadURLs, errors.New("No download URLs found")
	}

	return downloadURLs, nil
}
