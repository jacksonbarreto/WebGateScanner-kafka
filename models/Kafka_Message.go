package models

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
	InstitutionID    string                 `json:"institution_id"`
	Origin           string                 `json:"origin"`
	StartTime        int64                  `json:"start_time"`
	EndTime          int64                  `json:"end_time"`
	EvaluationResult map[string]interface{} `json:"evaluation_result"`
}
