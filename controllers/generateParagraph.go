package controllers

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/vigneshkk18/go-apis/utils"
)

func GetRandomParagraph(difficulty uint) (string, error) {
	// Construct request url with difficulty level passed in.
	reqUrl := fmt.Sprintf(utils.PublicApis["generateParagraph"], difficulty/5, difficulty)
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
