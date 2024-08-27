package formatters

import (
	"bitbucket/cliformat"
	"bitbucket/constants"
	"fmt"
	"time"
)

type PullRequestsResponse struct {
	Page    int           `json:"page"`
	Values  []PullRequest `json:"values"`
	Size    int           `json:"size"`
	Pagelen int           `json:"pagelen"`
}
type CreatePullRequestsResponse struct {
	Links struct {
		Html struct {
			Href string `json:"href"`
		} `json:"html"`
	} `json:"links"`
}

type PullRequestAuthor struct {
	DisplayName string `json:"display_name"`
	Nickname    string `json:"nickname"`
}

type PullRequest struct {
	ID string `json:"id"`

	Title       string            `json:"title"`
	Description string            `json:"description"`
	State       string            `json:"state"`
	Author      PullRequestAuthor `json:"author"`
	Destination struct {
		Branch struct {
			Name string `json:"name"`
		} `json:"branch"`
	} `json:"destination"`
	Source struct {
		Branch struct {
			Name string `json:"name"`
		} `json:"branch"`
	} `json:"source"`
	Links struct {
		Html struct {
			Href string `json:"href"`
		} `json:"html"`
	} `json:"links"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
}

type StateUpdate struct {
	Update struct {
		State       string `json:"state"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Reviewers   []struct {
			DisplayName string `json:"display_name"`
			Links       struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
			} `json:"links"`
			Type      string `json:"type"`
			UUID      string `json:"uuid"`
			AccountID string `json:"account_id"`
			Nickname  string `json:"nickname"`
		} `json:"reviewers"`
		Changes struct {
		} `json:"changes"`
		Reason string `json:"reason"`
		Author struct {
			DisplayName string `json:"display_name"`
			Links       struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
			} `json:"links"`
			Type      string `json:"type"`
			UUID      string `json:"uuid"`
			AccountID string `json:"account_id"`
			Nickname  string `json:"nickname"`
		} `json:"author"`
		Date        time.Time `json:"date"`
		Destination struct {
			Branch struct {
				Name string `json:"name"`
			} `json:"branch"`
			Commit struct {
				Hash  string `json:"hash"`
				Links struct {
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
					HTML struct {
						Href string `json:"href"`
					} `json:"html"`
				} `json:"links"`
				Type string `json:"type"`
			} `json:"commit"`
			Repository struct {
				Type     string `json:"type"`
				FullName string `json:"full_name"`
				Links    struct {
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
					HTML struct {
						Href string `json:"href"`
					} `json:"html"`
					Avatar struct {
						Href string `json:"href"`
					} `json:"avatar"`
				} `json:"links"`
				Name string `json:"name"`
				UUID string `json:"uuid"`
			} `json:"repository"`
		} `json:"destination"`
		Source struct {
			Branch struct {
				Name string `json:"name"`
			} `json:"branch"`
			Commit struct {
				Hash  string `json:"hash"`
				Links struct {
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
					HTML struct {
						Href string `json:"href"`
					} `json:"html"`
				} `json:"links"`
				Type string `json:"type"`
			} `json:"commit"`
			Repository struct {
				Type     string `json:"type"`
				FullName string `json:"full_name"`
				Links    struct {
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
					HTML struct {
						Href string `json:"href"`
					} `json:"html"`
					Avatar struct {
						Href string `json:"href"`
					} `json:"avatar"`
				} `json:"links"`
				Name string `json:"name"`
				UUID string `json:"uuid"`
			} `json:"repository"`
		} `json:"source"`
	} `json:"update,omitempty"`
	Comment struct {
		ID        int       `json:"id"`
		CreatedOn time.Time `json:"created_on"`
		UpdatedOn time.Time `json:"updated_on"`
		Content   struct {
			Type   string `json:"type"`
			Raw    string `json:"raw"`
			Markup string `json:"markup"`
			HTML   string `json:"html"`
		} `json:"content"`
		User struct {
			DisplayName string `json:"display_name"`
			Links       struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
			} `json:"links"`
			Type      string `json:"type"`
			UUID      string `json:"uuid"`
			AccountID string `json:"account_id"`
			Nickname  string `json:"nickname"`
		} `json:"user"`
		Deleted bool `json:"deleted"`
		Parent  struct {
			ID    int `json:"id"`
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
			} `json:"links"`
		} `json:"parent"`
		Inline struct {
		} `json:"inline"`
		Pending bool   `json:"pending"`
		Type    string `json:"type"`
		Links   struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			HTML struct {
				Href string `json:"href"`
			} `json:"html"`
		} `json:"links"`
		Pullrequest struct {
			Type  string `json:"type"`
			ID    int    `json:"id"`
			Title string `json:"title"`
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
			} `json:"links"`
		} `json:"pullrequest"`
	} `json:"comment,omitempty"`
	ChangesRequested struct {
		Date time.Time `json:"date"`
		User struct {
			DisplayName string `json:"display_name"`
			Links       struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
			} `json:"links"`
			Type      string `json:"type"`
			UUID      string `json:"uuid"`
			AccountID string `json:"account_id"`
			Nickname  string `json:"nickname"`
		} `json:"user"`
		Pullrequest struct {
			Type  string `json:"type"`
			ID    int    `json:"id"`
			Title string `json:"title"`
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
			} `json:"links"`
		} `json:"pullrequest"`
	} `json:"changes_requested,omitempty"`
	Approval struct {
		Date time.Time `json:"date"`
		User struct {
			DisplayName string `json:"display_name"`
			Links       struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
			} `json:"links"`
			Type      string `json:"type"`
			UUID      string `json:"uuid"`
			AccountID string `json:"account_id"`
			Nickname  string `json:"nickname"`
		} `json:"user"`
		Pullrequest struct {
			Type  string `json:"type"`
			ID    int    `json:"id"`
			Title string `json:"title"`
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
			} `json:"links"`
		} `json:"pullrequest"`
	} `json:"approval,omitempty"`
}
type PullRequestActivities struct {
	Values []StateUpdate `json:"values"`
}

func FormatPullrequest(pullrequest PullRequest) string {

	state := pullrequest.State
	if state == constants.PULLREQUEST_MERGED {

		state = cliformat.Success(state)
	} else if state == constants.PULLREQUEST_DECLIEND {
		state = cliformat.Error(state)
	} else {
		state = cliformat.Info(state)
	}

	location := time.Now().Location()
	createdOn := pullrequest.CreatedOn.In(location).Format("2006-01-02 15:04")
	updatedOn := pullrequest.UpdatedOn.In(location).Format("2006-01-02 15:04")
	createdBy := pullrequest.Author.DisplayName
	branches := pullrequest.Source.Branch.Name + " -> " + pullrequest.Destination.Branch.Name
	url := pullrequest.Links.Html.Href
	return fmt.Sprintf("%s %s %s %s \n%s\n%s\n", cliformat.RightPad(createdBy, 20, " "), createdOn, updatedOn, state, branches, url)
}

func FormatPullrequestResponse(pullrequests PullRequestsResponse) {

	for i := 0; i < len(pullrequests.Values); i++ {
		fmt.Println(FormatPullrequest(pullrequests.Values[i]))
	}
}

func formatPullRequestState(prStateUpdate StateUpdate) string {
	state := ""
	description := ""
	date := time.Now()
	if !prStateUpdate.Approval.Date.IsZero() {
		state = cliformat.Success("Approved")
		description = prStateUpdate.Approval.User.DisplayName
		date = prStateUpdate.Approval.Date
	}
	if !prStateUpdate.ChangesRequested.Date.IsZero() {
		state = cliformat.Info("Changes requested")
		description = prStateUpdate.ChangesRequested.User.DisplayName
		date = prStateUpdate.ChangesRequested.Date

	}
	if !prStateUpdate.Comment.CreatedOn.IsZero() {
		state = cliformat.BlueState("Comment")
		description = prStateUpdate.Comment.Content.Raw
		date = prStateUpdate.Comment.CreatedOn
	}
	location := time.Now().Location()
	formattedDate := date.In(location).Format("2006-01-02 15:04")
	return fmt.Sprintf("- %s %s %s", state, description, formattedDate)

}
func FormatPullrequestActivitites(prActivities PullRequestActivities) {
	approved := 0
	changedRequested := 0
	comments := 0
	approvedItems := ""
	changedRequestedItems := ""
	commentsItems := ""

	for i := 0; i < len(prActivities.Values); i++ {
		if !prActivities.Values[i].Approval.Date.IsZero() {
			approved += 1
			approvedItems += formatPullRequestState(prActivities.Values[i]) + "\n"
		}
		if !prActivities.Values[i].ChangesRequested.Date.IsZero() {
			changedRequested += 1

			changedRequestedItems += formatPullRequestState(prActivities.Values[i]) + "\n"
		}
		if !prActivities.Values[i].Comment.CreatedOn.IsZero() {
			comments += 1
			commentsItems += formatPullRequestState(prActivities.Values[i]) + "\n"

		}

	}
	results := cliformat.Success(fmt.Sprintf("%d Approved", approved)) + cliformat.Info(fmt.Sprintf(" %d Changes requested", changedRequested)) + cliformat.BlueState(fmt.Sprintf(" %d Comments", comments))

	fmt.Printf("\n%s\n%s\n%s\n%s\n",
		commentsItems,
		changedRequestedItems,
		approvedItems,
		results,
	)
}
