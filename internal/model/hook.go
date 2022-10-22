package model

type MergeRequestHook struct {
	Title     string   `json:"title"`
	State     string   `json:"state"`
	Action    string   `json:"action"`
	Username  string   `json:"username"`
	Url       string   `json:"url"`
	Project   string   `json:"project"`
	Assignees []string `json:"assignees"`
}
