package types

type FrigateEvent struct {
	Id           string   `json:"id"`
	Label        string   `json:"label"`
	Camera       string   `json:"camera"`
	Score        float32  `json:"score"`
	TopScore     float32  `json:"top_score"`
	StartTime    float64  `json:"start_time"`
	EndTime      float64  `json:"end_time"`
	IsStationary bool     `json:"stationary"`
	HasClip      bool     `json:"has_clip"`
	Thumbnail    string   `json:"thumbnail"`
	EnteredZones []string `json:"entered_zones"`
	CurrentZones []string `json:"current_zones"`
}
