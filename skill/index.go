package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type HTTPRequest struct {
	Body string `json:"body"`
}

func (req *HTTPRequest) Event() (event Event) {
	err := json.Unmarshal([]byte(req.Body), &event)
	if err != nil {
		log.Fatalln(err)
	}
	return
}

type Event struct {
	Version string      `json:"version"`
	Session interface{} `json:"session"`
	Request Request     `json:"request"`
}

type Request struct {
	Command           string `json:"command"`
	OriginalUtterance string `json:"original_utterance"`
}

type HTTPResponse struct {
	Body       Response `json:"body"`
	StatusCode int      `json:"statusCode"`
}

type Response struct {
	Version  string      `json:"version"`
	Session  interface{} `json:"session"`
	Response Answer      `json:"response"`
}

type Answer struct {
	Text       string `json:"text"`
	EndSession string `json:"end_session"`
}

func Handler(request HTTPRequest) (httpResponse *HTTPResponse, reqErr error) {
	event := request.Event()
	httpResponse = &HTTPResponse{
		Body: Response{
			Version: event.Version,
			Session: event.Session,
			Response: Answer{
				EndSession: "false",
			},
		},
		StatusCode: 200,
	}
	answer := &httpResponse.Body.Response.Text

	text := strings.ToLower(event.Request.Command)
	for k, day := range daysAfterToday {
		if strings.Contains(text, day) {
			dayAnswer, err := getDayAfterToday(k)
			if err != nil {
				*answer = err.Error()
				return
			}
			*answer = fmt.Sprintf(
				"%s у Вас: %s",
				cases.Title(language.Russian).String(day),
				dayAnswer,
			)
			return
		}
	}
	for k, day := range weekdays {
		if strings.Contains(text, day) {
			dayAnswer, err := getByWeekday(k)
			if err != nil {
				*answer = err.Error()
				return
			}
			*answer = fmt.Sprintf(
				"%s %s у Вас: %s",
				func() string {
					if []rune(day)[0] == 'в' {
						return "Во"
					}
					return "В"
				}(),
				cases.Title(language.Russian).String(day),
				dayAnswer,
			)
			return
		}
	}
	return
}
