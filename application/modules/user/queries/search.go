package queries

import (
	"encoding/json"
	"fmt"
	"hulk/go-webservice/domain/entities"
	"hulk/go-webservice/infrastructure/services"
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

func SearchUser(keyword string) ([]entities.User, error) {
	queryFormat := `{
        "search_type": "match",
        "query":{
            "term": "%s",
        },
        "from": 0,
        "max_results": 20,
        "_source": []
    }`
	query := fmt.Sprintf(queryFormat, keyword)
	var users []entities.User
	dataStr, err := services.FTSInstance.Search("users", query)
	if err != nil {
		return users, err
	}

	var parseResponse SearchRes[entities.User]
	err = json.Unmarshal([]byte(dataStr), &parseResponse)
	if err != nil {
		return users, err
	}

	for i := range parseResponse.Hits.Hits {
		user := parseResponse.Hits.Hits[i].Source
		users = append(users, user)
	}

	return users, nil
}
