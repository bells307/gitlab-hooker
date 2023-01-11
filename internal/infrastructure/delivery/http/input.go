package http

import (
	"github.com/bells307/gitlab-hooker/internal/domain/merge_request"
	"github.com/bells307/gitlab-hooker/internal/domain/pipeline"
	"time"
)

type MergeRequestHookInput struct {
	ObjectKind string `json:"object_kind"`
	EventType  string `json:"event_type"`
	User       struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Username  string `json:"username"`
		AvatarURL string `json:"avatar_url"`
		Email     string `json:"email"`
	} `json:"user"`
	Project struct {
		ID                int    `json:"id"`
		Name              string `json:"name"`
		Description       string `json:"description"`
		WebURL            string `json:"web_url"`
		AvatarURL         string `json:"avatar_url"`
		GitSSHURL         string `json:"git_ssh_url"`
		GitHTTPURL        string `json:"git_http_url"`
		Namespace         string `json:"namespace"`
		VisibilityLevel   int    `json:"visibility_level"`
		PathWithNamespace string `json:"path_with_namespace"`
		DefaultBranch     string `json:"default_branch"`
		Homepage          string `json:"homepage"`
		URL               string `json:"url"`
		SSHURL            string `json:"ssh_url"`
		HTTPURL           string `json:"http_url"`
	} `json:"project"`
	Repository struct {
		Name        string `json:"name"`
		URL         string `json:"url"`
		Description string `json:"description"`
		Homepage    string `json:"homepage"`
	} `json:"repository"`
	ObjectAttributes struct {
		ID                          int    `json:"id"`
		Iid                         int    `json:"iid"`
		TargetBranch                string `json:"target_branch"`
		SourceBranch                string `json:"source_branch"`
		SourceProjectID             int    `json:"source_project_id"`
		AuthorID                    int    `json:"author_id"`
		AssigneeIds                 []int  `json:"assignee_ids"`
		AssigneeID                  int    `json:"assignee_id"`
		ReviewerIds                 []int  `json:"reviewer_ids"`
		Title                       string `json:"title"`
		CreatedAt                   string `json:"created_at"`
		UpdatedAt                   string `json:"updated_at"`
		MilestoneID                 int    `json:"milestone_id"`
		State                       string `json:"state"`
		BlockingDiscussionsResolved bool   `json:"blocking_discussions_resolved"`
		WorkInProgress              bool   `json:"work_in_progress"`
		FirstContribution           bool   `json:"first_contribution"`
		MergeStatus                 string `json:"merge_status"`
		TargetProjectID             int    `json:"target_project_id"`
		Description                 string `json:"description"`
		URL                         string `json:"url"`
		Source                      struct {
			Name              string `json:"name"`
			Description       string `json:"description"`
			WebURL            string `json:"web_url"`
			AvatarURL         string `json:"avatar_url"`
			GitSSHURL         string `json:"git_ssh_url"`
			GitHTTPURL        string `json:"git_http_url"`
			Namespace         string `json:"namespace"`
			VisibilityLevel   int    `json:"visibility_level"`
			PathWithNamespace string `json:"path_with_namespace"`
			DefaultBranch     string `json:"default_branch"`
			Homepage          string `json:"homepage"`
			URL               string `json:"url"`
			SSHURL            string `json:"ssh_url"`
			HTTPURL           string `json:"http_url"`
		} `json:"source"`
		Target struct {
			Name              string `json:"name"`
			Description       string `json:"description"`
			WebURL            string `json:"web_url"`
			AvatarURL         string `json:"avatar_url"`
			GitSSHURL         string `json:"git_ssh_url"`
			GitHTTPURL        string `json:"git_http_url"`
			Namespace         string `json:"namespace"`
			VisibilityLevel   int    `json:"visibility_level"`
			PathWithNamespace string `json:"path_with_namespace"`
			DefaultBranch     string `json:"default_branch"`
			Homepage          string `json:"homepage"`
			URL               string `json:"url"`
			SSHURL            string `json:"ssh_url"`
			HTTPURL           string `json:"http_url"`
		} `json:"target"`
		LastCommit struct {
			ID        string    `json:"id"`
			Message   string    `json:"message"`
			Timestamp time.Time `json:"timestamp"`
			URL       string    `json:"url"`
			Author    struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			} `json:"author"`
		} `json:"last_commit"`
		Labels []struct {
			ID          int    `json:"id"`
			Title       string `json:"title"`
			Color       string `json:"color"`
			ProjectID   int    `json:"project_id"`
			CreatedAt   string `json:"created_at"`
			UpdatedAt   string `json:"updated_at"`
			Template    bool   `json:"template"`
			Description string `json:"description"`
			Type        string `json:"type"`
			GroupID     int    `json:"group_id"`
		} `json:"labels"`
		Action string `json:"action"`
	} `json:"object_attributes"`
	Labels []struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Color       string `json:"color"`
		ProjectID   int    `json:"project_id"`
		CreatedAt   string `json:"created_at"`
		UpdatedAt   string `json:"updated_at"`
		Template    bool   `json:"template"`
		Description string `json:"description"`
		Type        string `json:"type"`
		GroupID     int    `json:"group_id"`
	} `json:"labels"`
	Changes struct {
		UpdatedByID struct {
			Previous int `json:"previous"`
			Current  int `json:"current"`
		} `json:"updated_by_id"`
		UpdatedAt struct {
			Previous string `json:"previous"`
			Current  string `json:"current"`
		} `json:"updated_at"`
		Labels struct {
			Previous []struct {
				ID          int    `json:"id"`
				Title       string `json:"title"`
				Color       string `json:"color"`
				ProjectID   int    `json:"project_id"`
				CreatedAt   string `json:"created_at"`
				UpdatedAt   string `json:"updated_at"`
				Template    bool   `json:"template"`
				Description string `json:"description"`
				Type        string `json:"type"`
				GroupID     int    `json:"group_id"`
			} `json:"previous"`
			Current []struct {
				ID          int    `json:"id"`
				Title       string `json:"title"`
				Color       string `json:"color"`
				ProjectID   int    `json:"project_id"`
				CreatedAt   string `json:"created_at"`
				UpdatedAt   string `json:"updated_at"`
				Template    bool   `json:"template"`
				Description string `json:"description"`
				Type        string `json:"type"`
				GroupID     int    `json:"group_id"`
			} `json:"current"`
		} `json:"labels"`
	} `json:"changes"`
	Assignees []struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Username  string `json:"username"`
		AvatarURL string `json:"avatar_url"`
	} `json:"assignees"`
	Reviewers []struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Username  string `json:"username"`
		AvatarURL string `json:"avatar_url"`
	} `json:"reviewers"`
}

func (in MergeRequestHookInput) ToDomain() merge_request.MergeRequest {
	return merge_request.MergeRequest{
		Title:    in.ObjectAttributes.Title,
		Username: in.User.Name,
		Project:  in.Project.Name,
		URL:      in.ObjectAttributes.URL,
		State:    merge_request.MergeRequestState(in.ObjectAttributes.State),
		Action:   merge_request.MergeRequestAction(in.ObjectAttributes.Action),
	}
}

type PipelineHookInput struct {
	ObjectKind       string `json:"object_kind"`
	ObjectAttributes struct {
		ID         int      `json:"id"`
		Ref        string   `json:"ref"`
		Tag        bool     `json:"tag"`
		Sha        string   `json:"sha"`
		BeforeSha  string   `json:"before_sha"`
		Source     string   `json:"source"`
		Status     string   `json:"status"`
		Stages     []string `json:"stages"`
		CreatedAt  string   `json:"created_at"`
		FinishedAt string   `json:"finished_at"`
		Duration   int      `json:"duration"`
		Variables  []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"variables"`
	} `json:"object_attributes"`
	MergeRequest struct {
		ID              int    `json:"id"`
		Iid             int    `json:"iid"`
		Title           string `json:"title"`
		SourceBranch    string `json:"source_branch"`
		SourceProjectID int    `json:"source_project_id"`
		TargetBranch    string `json:"target_branch"`
		TargetProjectID int    `json:"target_project_id"`
		State           string `json:"state"`
		MergeStatus     string `json:"merge_status"`
		URL             string `json:"url"`
	} `json:"merge_request"`
	User struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Username  string `json:"username"`
		AvatarURL string `json:"avatar_url"`
		Email     string `json:"email"`
	} `json:"user"`
	Project struct {
		ID                int    `json:"id"`
		Name              string `json:"name"`
		Description       string `json:"description"`
		WebURL            string `json:"web_url"`
		AvatarURL         string `json:"avatar_url"`
		GitSSHURL         string `json:"git_ssh_url"`
		GitHTTPURL        string `json:"git_http_url"`
		Namespace         string `json:"namespace"`
		VisibilityLevel   int    `json:"visibility_level"`
		PathWithNamespace string `json:"path_with_namespace"`
		DefaultBranch     string `json:"default_branch"`
	} `json:"project"`
	Commit struct {
		ID        string    `json:"id"`
		Message   string    `json:"message"`
		Timestamp time.Time `json:"timestamp"`
		URL       string    `json:"url"`
		Author    struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"author"`
	} `json:"commit"`
	SourcePipeline struct {
		Project struct {
			ID                int    `json:"id"`
			WebURL            string `json:"web_url"`
			PathWithNamespace string `json:"path_with_namespace"`
		} `json:"project"`
		PipelineID int `json:"pipeline_id"`
		JobID      int `json:"job_id"`
	} `json:"source_pipeline"`
	Builds []struct {
		ID             int         `json:"id"`
		Stage          string      `json:"stage"`
		Name           string      `json:"name"`
		Status         string      `json:"status"`
		CreatedAt      string      `json:"created_at"`
		StartedAt      interface{} `json:"started_at"`
		FinishedAt     interface{} `json:"finished_at"`
		Duration       interface{} `json:"duration"`
		QueuedDuration interface{} `json:"queued_duration"`
		FailureReason  interface{} `json:"failure_reason"`
		When           string      `json:"when"`
		Manual         bool        `json:"manual"`
		AllowFailure   bool        `json:"allow_failure"`
		User           struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			Username  string `json:"username"`
			AvatarURL string `json:"avatar_url"`
			Email     string `json:"email"`
		} `json:"user"`
		Runner        interface{} `json:"runner"`
		ArtifactsFile struct {
			Filename interface{} `json:"filename"`
			Size     interface{} `json:"size"`
		} `json:"artifacts_file"`
		Environment struct {
			Name           string `json:"name"`
			Action         string `json:"action"`
			DeploymentTier string `json:"deployment_tier"`
		} `json:"environment"`
	} `json:"builds"`
}

func (in PipelineHookInput) ToDomain() pipeline.Pipeline {
	return pipeline.Pipeline{
		Project: in.Project.Name,
		Branch:  in.ObjectAttributes.Ref,
		Status:  pipeline.PipelineStatus(in.ObjectAttributes.Status),
	}
}
