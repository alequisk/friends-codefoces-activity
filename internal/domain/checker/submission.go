package checker

type Problem struct {
	Name string `json:"name"`
}

type Submission struct {
	ID      int64   `json:"id"`
	Problem Problem `json:"problem"`
}
