package episodeNumber

import (
	"fmt"
	"regexp"
	"strings"
)

func ClearYare(data string) string {
	regexp, err := regexp.Compile(`(\d+年)`)
	if err != nil {
		return data
	}
	yareKey := regexp.FindString(data)
	data = strings.ReplaceAll(data, yareKey, "")
	return data
}

func ClearNumber(numbers [][]string) string {
	dataNumber := ""
	for _, item := range numbers {
		numberText := item[1]
		if len(numberText) > 0 {
			dataNumber = numberText
		}
	}
	return dataNumber
}

func EpisodeNumber(text string) (string, error) {
	regexp, err := regexp.Compile(`(\d{2,3})\D`)
	if err != nil {
		return "", err
	}
	text = strings.ReplaceAll(text, "1080", "")
	text = strings.ReplaceAll(text, "720", "")
	text = strings.ReplaceAll(text, "264", "")
	text = strings.ReplaceAll(text, "265", "")
	data := ClearYare(text)
	datas := regexp.FindAllStringSubmatch(data, -1)
	number := ClearNumber(datas)
	return number, nil
}

func SeasonEpisode(name, season, text string) (string, error) {
	eNumber, err := EpisodeNumber(text)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%sS%sE%s", name, season, eNumber), nil
}
