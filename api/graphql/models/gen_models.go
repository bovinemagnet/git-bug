// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"github.com/MichaelMure/git-bug/entities/bug"
	"github.com/MichaelMure/git-bug/entity/dag"
	"github.com/MichaelMure/git-bug/repository"
)

// An object that has an author.
type Authored interface {
	IsAuthored()
}

type AddCommentAndCloseBugInput struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The name of the repository. If not set, the default repository is used.
	RepoRef *string `json:"repoRef"`
	// The bug ID's prefix.
	Prefix string `json:"prefix"`
	// The message to be added to the bug.
	Message string `json:"message"`
	// The collection of file's hash required for the first message.
	Files []repository.Hash `json:"files"`
}

type AddCommentAndCloseBugPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The affected bug.
	Bug BugWrapper `json:"bug"`
	// The resulting AddComment operation.
	CommentOperation *bug.AddCommentOperation `json:"commentOperation"`
	// The resulting SetStatusOperation.
	StatusOperation *bug.SetStatusOperation `json:"statusOperation"`
}

type AddCommentAndReopenBugInput struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The name of the repository. If not set, the default repository is used.
	RepoRef *string `json:"repoRef"`
	// The bug ID's prefix.
	Prefix string `json:"prefix"`
	// The message to be added to the bug.
	Message string `json:"message"`
	// The collection of file's hash required for the first message.
	Files []repository.Hash `json:"files"`
}

type AddCommentAndReopenBugPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The affected bug.
	Bug BugWrapper `json:"bug"`
	// The resulting AddComment operation.
	CommentOperation *bug.AddCommentOperation `json:"commentOperation"`
	// The resulting SetStatusOperation.
	StatusOperation *bug.SetStatusOperation `json:"statusOperation"`
}

type AddCommentInput struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The name of the repository. If not set, the default repository is used.
	RepoRef *string `json:"repoRef"`
	// The bug ID's prefix.
	Prefix string `json:"prefix"`
	// The message to be added to the bug.
	Message string `json:"message"`
	// The collection of file's hash required for the first message.
	Files []repository.Hash `json:"files"`
}

type AddCommentPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The affected bug.
	Bug BugWrapper `json:"bug"`
	// The resulting operation.
	Operation *bug.AddCommentOperation `json:"operation"`
}

// The connection type for Bug.
type BugConnection struct {
	// A list of edges.
	Edges []*BugEdge   `json:"edges"`
	Nodes []BugWrapper `json:"nodes"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int `json:"totalCount"`
}

// An edge in a connection.
type BugEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node BugWrapper `json:"node"`
}

type ChangeLabelInput struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The name of the repository. If not set, the default repository is used.
	RepoRef *string `json:"repoRef"`
	// The bug ID's prefix.
	Prefix string `json:"prefix"`
	// The list of label to add.
	Added []string `json:"added"`
	// The list of label to remove.
	Removed []string `json:"Removed"`
}

type ChangeLabelPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The affected bug.
	Bug BugWrapper `json:"bug"`
	// The resulting operation.
	Operation *bug.LabelChangeOperation `json:"operation"`
	// The effect each source label had.
	Results []*bug.LabelChangeResult `json:"results"`
}

type CloseBugInput struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The name of the repository. If not set, the default repository is used.
	RepoRef *string `json:"repoRef"`
	// The bug ID's prefix.
	Prefix string `json:"prefix"`
}

type CloseBugPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The affected bug.
	Bug BugWrapper `json:"bug"`
	// The resulting operation.
	Operation *bug.SetStatusOperation `json:"operation"`
}

type CommentConnection struct {
	Edges      []*CommentEdge `json:"edges"`
	Nodes      []*bug.Comment `json:"nodes"`
	PageInfo   *PageInfo      `json:"pageInfo"`
	TotalCount int            `json:"totalCount"`
}

type CommentEdge struct {
	Cursor string       `json:"cursor"`
	Node   *bug.Comment `json:"node"`
}

type EditCommentInput struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The name of the repository. If not set, the default repository is used.
	RepoRef *string `json:"repoRef"`
	// The bug ID's prefix.
	Prefix string `json:"prefix"`
	// The ID of the comment to be changed.
	Target string `json:"target"`
	// The new message to be set.
	Message string `json:"message"`
	// The collection of file's hash required for the first message.
	Files []repository.Hash `json:"files"`
}

type EditCommentPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The affected bug.
	Bug BugWrapper `json:"bug"`
	// The resulting operation.
	Operation *bug.EditCommentOperation `json:"operation"`
}

type IdentityConnection struct {
	Edges      []*IdentityEdge   `json:"edges"`
	Nodes      []IdentityWrapper `json:"nodes"`
	PageInfo   *PageInfo         `json:"pageInfo"`
	TotalCount int               `json:"totalCount"`
}

type IdentityEdge struct {
	Cursor string          `json:"cursor"`
	Node   IdentityWrapper `json:"node"`
}

type LabelConnection struct {
	Edges      []*LabelEdge `json:"edges"`
	Nodes      []bug.Label  `json:"nodes"`
	PageInfo   *PageInfo    `json:"pageInfo"`
	TotalCount int          `json:"totalCount"`
}

type LabelEdge struct {
	Cursor string    `json:"cursor"`
	Node   bug.Label `json:"node"`
}

type NewBugInput struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The name of the repository. If not set, the default repository is used.
	RepoRef *string `json:"repoRef"`
	// The title of the new bug.
	Title string `json:"title"`
	// The first message of the new bug.
	Message string `json:"message"`
	// The collection of file's hash required for the first message.
	Files []repository.Hash `json:"files"`
}

type NewBugPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The created bug.
	Bug BugWrapper `json:"bug"`
	// The resulting operation.
	Operation *bug.CreateOperation `json:"operation"`
}

type OpenBugInput struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The name of the repository. If not set, the default repository is used.
	RepoRef *string `json:"repoRef"`
	// The bug ID's prefix.
	Prefix string `json:"prefix"`
}

type OpenBugPayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The affected bug.
	Bug BugWrapper `json:"bug"`
	// The resulting operation.
	Operation *bug.SetStatusOperation `json:"operation"`
}

// The connection type for an Operation
type OperationConnection struct {
	Edges      []*OperationEdge `json:"edges"`
	Nodes      []dag.Operation  `json:"nodes"`
	PageInfo   *PageInfo        `json:"pageInfo"`
	TotalCount int              `json:"totalCount"`
}

// Represent an Operation
type OperationEdge struct {
	Cursor string        `json:"cursor"`
	Node   dag.Operation `json:"node"`
}

// Information about pagination in a connection.
type PageInfo struct {
	// When paginating forwards, are there more items?
	HasNextPage bool `json:"hasNextPage"`
	// When paginating backwards, are there more items?
	HasPreviousPage bool `json:"hasPreviousPage"`
	// When paginating backwards, the cursor to continue.
	StartCursor string `json:"startCursor"`
	// When paginating forwards, the cursor to continue.
	EndCursor string `json:"endCursor"`
}

type SetTitleInput struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The name of the repository. If not set, the default repository is used.
	RepoRef *string `json:"repoRef"`
	// The bug ID's prefix.
	Prefix string `json:"prefix"`
	// The new title.
	Title string `json:"title"`
}

type SetTitlePayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationID *string `json:"clientMutationId"`
	// The affected bug.
	Bug BugWrapper `json:"bug"`
	// The resulting operation
	Operation *bug.SetTitleOperation `json:"operation"`
}

// The connection type for TimelineItem
type TimelineItemConnection struct {
	Edges      []*TimelineItemEdge `json:"edges"`
	Nodes      []bug.TimelineItem  `json:"nodes"`
	PageInfo   *PageInfo           `json:"pageInfo"`
	TotalCount int                 `json:"totalCount"`
}

// Represent a TimelineItem
type TimelineItemEdge struct {
	Cursor string           `json:"cursor"`
	Node   bug.TimelineItem `json:"node"`
}
