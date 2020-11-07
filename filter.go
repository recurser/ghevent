package ghevent

import (
	"github.com/google/go-github/v32/github"
)

// FilterByAction decides whether the given event should be handled or dropped.
//
// Implements the "comma, ok" idiom, where the first return value is
// the event, ,
// and the second return value indicates not-ok.
func FilterByAction(event interface{}) (interface{}, bool) {
	switch e := event.(type) {

	// MembershipEvent is triggered when a user is added or removed from a team. See:
	//
	// - https://developer.github.com/webhooks/event-payloads/#membership
	// - https://github.com/google/go-github/blob/8da2410a408643f0a1781c0f748b9c3b7039402b/github/event_types.go#L405
	case *github.MembershipEvent:
		if e.Action != nil && (false ||
			*e.Action == "added" ||
			*e.Action == "removed") {
			return event, true
		}

	// MetaEvent is triggered when our webhook is deleted. See:
	//
	// - https://developer.github.com/webhooks/event-payloads/#meta
	// - https://github.com/google/go-github/blob/8da2410a408643f0a1781c0f748b9c3b7039402b/github/event_types.go#L651
	//
	// Action is always "deleted".
	case *github.MetaEvent:
		if e.Action != nil && *e.Action == "deleted" {
			return event, true
		}

	// OrganizationEvent is triggered when an organization is deleted and renamed, and when a user is added,
	// removed, or invited to an organization. See:
	//
	// - https://developer.github.com/webhooks/event-payloads/#organization
	// - https://github.com/google/go-github/blob/8da2410a408643f0a1781c0f748b9c3b7039402b/github/event_types.go#L460
	case *github.OrganizationEvent:
		if e.Action != nil && (false ||
			*e.Action == "deleted" ||
			*e.Action == "renamed" ||
			*e.Action == "member_added" ||
			*e.Action == "member_invited" ||
			*e.Action == "member_removed") {
			return event, true
		}

	// PullRequestEvent is triggered when multiple things happen to a pull request. See:
	//
	// - https://developer.github.com/webhooks/event-payloads/#pull_request
	// - https://github.com/google/go-github/blob/8da2410a408643f0a1781c0f748b9c3b7039402b/github/event_types.go#L610
	case *github.PullRequestEvent:
		if e.Action != nil && (false ||
			*e.Action == "closed" ||
			*e.Action == "opened" ||
			*e.Action == "reopened" ||
			*e.Action == "review_request_removed" ||
			*e.Action == "review_requested" ||
			*e.Action == "synchronize") {
			return event, true
		}

	// PullRequestReviewEvent is triggered when a review is submitted on a pull request. See:
	//
	// - https://developer.github.com/webhooks/event-payloads/#pull_request_review
	// - https://github.com/google/go-github/blob/8da2410a408643f0a1781c0f748b9c3b7039402b/github/event_types.go#L651
	//
	// Action is always "submitted".
	case *github.PullRequestReviewEvent:
		if e.Action != nil && *e.Action == "submitted" {
			return event, true
		}

	// RepositoryEvent is triggered when a repository is created, deleted etc. See:
	//
	// - https://developer.github.com/webhooks/event-payloads/#repository
	// - https://github.com/google/go-github/blob/8da2410a408643f0a1781c0f748b9c3b7039402b/github/event_types.go#L812
	case *github.RepositoryEvent:
		if e.Action != nil && (false ||
			*e.Action == "created" ||
			*e.Action == "deleted" ||
			*e.Action == "archived" ||
			*e.Action == "unarchived" ||
			*e.Action == "renamed" ||
			*e.Action == "transferred") {
			return event, true
		}

	// TeamEvent is triggered when an organization's team is created, modified or deleted. See:
	//
	// - https://developer.github.com/webhooks/event-payloads/#team
	// - https://github.com/google/go-github/blob/8da2410a408643f0a1781c0f748b9c3b7039402b/github/event_types.go#L911
	case *github.TeamEvent:
		if e.Action != nil && (false ||
			*e.Action == "created" ||
			*e.Action == "deleted" ||
			*e.Action == "edited") {
			return event, true
		}

	default:
		return nil, false
	}

	return nil, true
}
