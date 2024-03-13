package models

import "encoding/json"

type KafkaErrorMessage struct {
	Origin string `json:"origin"`
	Url    string `json:"url"`
	Error  string `json:"error"`
}

type KafkaMessage struct {
	Origin string `json:"origin"`
}

type EvaluationRequest struct {
	InstitutionID string `json:"institution_id"`
	URL           string `json:"url"`
}

type EvaluationResponse struct {
	InstitutionID    string          `json:"institution_id"`
	Origin           string          `json:"origin"`
	StartTime        int64           `json:"start_time"`
	EndTime          int64           `json:"end_time"`
	EvaluationResult json.RawMessage `json:"evaluation_result"`
}

func CreateKafkaEvaluationResponseMessage(institutionID string, origin string,
	startTime int64, endTime int64, evaluationResult interface{}) (string, error) {
	resultJson, err := json.Marshal(evaluationResult)
	if err != nil {
		return "", err
	}

	evalResponse := EvaluationResponse{
		StartTime:        startTime,
		EndTime:          endTime,
		Origin:           origin,
		InstitutionID:    institutionID,
		EvaluationResult: json.RawMessage(resultJson),
	}

	jsonData, err := json.Marshal(evalResponse)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
