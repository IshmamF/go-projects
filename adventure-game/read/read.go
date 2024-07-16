package read

import (
	"encoding/json"
	"os"
)

type Option struct {
	Text string `json:"text"`
	Arc string `json:"arc"`
}

type Details struct {
	Title string `json:"title"`
	Stories []string `json:"story"`
	Options []Option `json:"options"`
}

type Story map[string]Details

func ReadJSON(path string) Story {
	fd := FileReader(path)
	var chapters Story
	AddData(chapters, fd)
	return chapters
}

func FileReader (path string) *os.File {
	fd, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return fd
}

func AddData(chapters Story, fd *os.File) {
	jsonDecoder := json.NewDecoder(fd)
	if err := jsonDecoder.Decode(&chapters); err != nil {
		panic(err)
	}
}