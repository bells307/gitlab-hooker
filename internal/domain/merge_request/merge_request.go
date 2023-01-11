package merge_request

type MergeRequest struct {
	Title    string
	Username string
	Project  string
	URL      string
	State    MergeRequestState
	Action   MergeRequestAction
}

type MergeRequestState string

const (
	StateOpened MergeRequestState = "opened"
	StateClosed                   = "closed"
	StateLocked                   = "locked"
	StateMerged                   = "merged"
)

type MergeRequestAction string

const (
	ActionOpen       MergeRequestAction = "open"
	ActionClose                         = "close"
	ActionReopen                        = "reopen"
	ActionUpdate                        = "update"
	ActionApproved                      = "approved"
	ActionUnapproved                    = "unapproved"
	ActionApproval                      = "approval"
	ActionUnapproval                    = "unapproval"
	ActionMerge                         = "merge"
)
