package jobs

type JobPost struct {
	Title    string  `json:"title"`
	Category string  `json:"category"`
	JobDesc  JobDesc `json:"jobDesc"`
}
