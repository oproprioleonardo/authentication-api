package session

type Response struct {
	ID          string `json:"id"`
	ProfileId   string `json:"profileId"`
	Active      bool   `json:"active"`
	LastServer  string `json:"lastServer"`
	Ip          string `json:"lastIp"`
	FinalizedAt int64  `json:"finalizedAt"`
	StartedAt   int64  `json:"startedAt"`
}
