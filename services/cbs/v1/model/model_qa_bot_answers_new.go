package model

import (
	"encoding/json"

	"strings"
)

//
type QaBotAnswersNew struct {
	// 问答机器人回复。

	Answers *[]QaBotAnswer `json:"answers,omitempty"`
	// 问答机器人推荐问题

	RecommendAnswers *[]RecomendAnswer `json:"recommend_answers,omitempty"`
}

func (o QaBotAnswersNew) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QaBotAnswersNew struct{}"
	}

	return strings.Join([]string{"QaBotAnswersNew", string(data)}, " ")
}