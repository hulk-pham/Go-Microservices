package services

import (
	"fmt"
	"hulk/go-webservice/infrastructure/config"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type SearchRes[T any] struct {
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		Hits []struct {
			Id     string `json:"_id"`
			Source T      `json:"_source"`
		}
	} `json:"hits"`
}

// FTS = Full Text Search
type FTSService struct {
	httpClient         *http.Client
	baseUrl            string
	credentialUserName string
	credentialPassword string
}

var FTSInstance FTSService

func InitFTSService() {
	config := config.AppConfig()
	FTSInstance.httpClient = &http.Client{}
	FTSInstance.baseUrl = config.ZincAddress
	FTSInstance.credentialUserName = config.ZincUsername
	FTSInstance.credentialPassword = config.ZincPassword
}

func (m *FTSService) Request(uri string, method string, data string) (string, error) {
	url := fmt.Sprintf("%s/%s", m.baseUrl, uri)
	var payload io.Reader

	if data != "" {
		payload = strings.NewReader(data)
	}

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(m.credentialUserName, m.credentialPassword)

	resp, err := m.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(body), nil
}

func (m *FTSService) Search(indexName string, query string) (string, error) {
	dataStr, err := m.Request(fmt.Sprintf("api/%s/_search", indexName), "POST", query)
	if err != nil {
		return "", err
	}
	return dataStr, nil
}

func (m *FTSService) CreateDoc(indexName string, dataEncoded string) (string, error) {
	dataStr, err := m.Request(fmt.Sprintf("api/%s/_doc", indexName), "POST", dataEncoded)
	if err != nil {
		return "", err
	}
	return dataStr, nil
}

func (m *FTSService) UpdateDoc(indexName string, id string, dataEncoded string) (string, error) {
	dataStr, err := m.Request(fmt.Sprintf("api/%s/_update/%s", indexName, id), "POST", dataEncoded)
	if err != nil {
		return "", err
	}
	return dataStr, nil
}

func (m *FTSService) DeleteDoc(indexName string, id string) (string, error) {
	dataStr, err := m.Request(fmt.Sprintf("api/%s/_update/%s", indexName, id), "DELETE", "")
	if err != nil {
		return "", err
	}
	return dataStr, nil
}
