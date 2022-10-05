package controllers

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/vigneshkk18/go-apis/utils"
)

func GetDifficultyOptions() ([]string, error) {
	return utils.MapKeys(utils.DifficultyMap), nil
}

func GetRandomParagraph(difficulty uint) (string, error) {
	// Construct request url with difficulty level passed in.
	reqUrl := fmt.Sprintf(utils.PublicApis["generateParagraph"], difficulty)
	resp, err := http.Get(reqUrl)

	if err != nil {
		return "", errors.New("unable to generate paragraph")
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", errors.New("unable to generate paragraph")
	}

	return string(body), nil
}
