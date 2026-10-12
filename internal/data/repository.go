package data

import (
	graphql "github.com/cli/shurcooL-graphql"
)

type BranchProtectionRules struct {
	Nodes []struct {
		RequiredApprovingReviewCount int
		RequiresApprovingReviews     graphql.Boolean
		RequiresCodeOwnerReviews     graphql.Boolean
		RequiresStatusChecks         graphql.Boolean
		RequiredStatusCheckContexts  []graphql.String
	}
}

type GitHubLogin string

type Owner struct {
	Login string
}

type Repository struct {
	Name                  string
	Owner                 Owner
	NameWithOwner         string
	IsArchived            bool
	ViewerPermission      string
	BranchProtectionRules BranchProtectionRules `graphql:"branchProtectionRules(first: 1)"`
}

// ViewerCanPush reports whether the viewer has write access to the repository,
// which GitHub requires to approve workflow runs from fork pull requests.
func (r Repository) ViewerCanPush() bool {
	switch r.ViewerPermission {
	case "ADMIN", "MAINTAIN", "WRITE":
		return true
	}
	return false
}
