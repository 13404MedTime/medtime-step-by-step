package function

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	botToken        = "6519849383:AAHw5BnPuFvtER6MNW6cNgcrVG6bMvElgac"
	chatID          = "1546926238"
	baseUrl         = "https://api.admin.u-code.io"
	logFunctionName = "ucode-template"
	IsHTTP          = true // if this is true banchmark test works.
)

/*
Answer below questions before starting the function.

When the function invoked?
 - table_slug -> AFTER | BEFORE | HTTP -> CREATE | UPDATE | MULTIPLE_UPDATE | DELETE | APPEND_MANY2MANY | DELETE_MANY2MANY
What does it do?
- Explain the purpose of the function.(O'zbekcha yozilsa ham bo'ladi.)
*/

// Request structures
type (
	// Handle request body
	NewRequestBody struct {
		RequestData HttpRequest `json:"request_data"`
		Auth        AuthData    `json:"auth"`
		Data        Data        `json:"data"`
	}

	HttpRequest struct {
		Method  string      `json:"method"`
		Path    string      `json:"path"`
		Headers http.Header `json:"headers"`
		Params  url.Values  `json:"params"`
		Body    []byte      `json:"body"`
	}

	AuthData struct {
		Type string                 `json:"type"`
		Data map[string]interface{} `json:"data"`
	}

	// Function request body >>>>> GET_LIST, GET_LIST_SLIM, CREATE, UPDATE
	Request struct {
		Data map[string]interface{} `json:"data"`
	}

	// most common request structure -> UPDATE, MULTIPLE_UPDATE, CREATE, DELETE
	Data struct {
		AppId      string `json:"app_id"`
		Method     string `json:"method"`
		ObjectData struct {
			UserID   string  `json:"user_id"`
			Steps    int     `json:"steps"`
			Km       float64 `json:"km"`
			MoveTime struct {
				Hour   float64 `json:"hour"`
				Minute int     `json:"minute"`
			} `json:"move_time"`
			Date string `json:"date"`
		} `json:"object_data"`
		ObjectIds []string `json:"object_ids"`
		TableSlug string   `json:"table_slug"`
		UserId    string   `json:"user_id"`
	}

	FunctionRequest struct {
		BaseUrl     string  `json:"base_url"`
		TableSlug   string  `json:"table_slug"`
		AppId       string  `json:"app_id"`
		Request     Request `json:"request"`
		DisableFaas bool    `json:"disable_faas"`
	}
)

type StepsToday struct {
	Status      string `json:"status"`
	Description string `json:"description"`
	Data        struct {
		TableSlug string `json:"table_slug"`
		Data      struct {
			Count    int `json:"count"`
			Response []struct {
				CleintsID string  `json:"cleints_id,omitempty"`
				Date      string  `json:"date"`
				Distance  int     `json:"distance"`
				GUID      string  `json:"guid"`
				Hour      int     `json:"hour"`
				Minutes   int     `json:"minutes"`
				StepCount int     `json:"step_count"`
				Time      float64 `json:"time"`
			} `json:"response"`
		} `json:"data"`
	} `json:"data"`
	CustomMessage string `json:"custom_message"`
}
