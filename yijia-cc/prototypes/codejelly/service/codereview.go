package service

import (
	"context"
	"errors"
	"github.com/teamyapp/experimental/yijia-cc/prototypes/codejelly/entity"
)

type PullRequestFilter int

const (
	prNeedAttention PullRequestFilter = iota
	prOpenedByMe
	prReviewedByMe
	prMergedByMe
)

type CodeReview struct {
	//repo vcs.Repository
}

func (c CodeReview) ListPullRequests(ctx context.Context, filter PullRequestFilter) []entity.PullRequest{
	// requesterId int
	panic("implement me")
}

func (c CodeReview) FindPullRequest(ctx context.Context, pullRequestId int) (entity.PullRequest, error){
	//requesterId int
	panic ("implement me")
}

func (c CodeReview) OpenPullRequest(ctx context.Context, pullRequest entity.PullRequest) (pullRequestId int, err error) {
	// requesterId int
	panic("implement me")
}

func (c CodeReview) ListPullRequestActions(ctx context.Context, pullRequestId int) []entity.PullRequestAction{
	//requesterId int
	panic("implement me")
}

func (c CodeReview) TakePullRequestAction(ctx context.Context, pullRequestId int, action entity.PullRequestAction) error {
	//requesterId int
	panic("implement me")
}

func (c CodeReview) ListMergeOptions(ctx context.Context, pullRequestId int) []entity.MergeOption {
	//requesterId int
	panic("implement me")
}

func (c CodeReview) ListReviewers(ctx context.Context, pullRequestId int) []entity.Reviewer {
	//requesterId int
	panic("implement me")
}

func (c CodeReview) TakeReviewerAction(ctx context.Context, action entity.EditReviewerAction, reviewerId string, pullRequestId int) {
	panic("implement me")
}

func (c CodeReview) ListThreads(ctx context.Context, pullRequestId int) []entity.Thread {
	panic("implement me")
}

func (c CodeReview) ListThreadActions(ctx context.Context, commentId int) []entity.ThreadAction {
	panic("implement me")
}

func (c CodeReview) ListCommentActions(ctx context.Context, commentId int) []entity.CommentAction {
	panic("implement me")
}

func (c CodeReview) TakeCommentAction(ctx context.Context, commentId int) []entity.CommentAction{
	panic("implement me")
}

func pickFilePath(fromFilePath string, toFilePath string) (string, error) {
	if len(fromFilePath) == 0 && len(toFilePath) == 0 {
		return "", errors.New("invalid fileDiff")
	}

	if len(fromFilePath) == 0 {
		return toFilePath, nil
	}

	if len(toFilePath) == 0 {
		return toFilePath, nil
	}

	return toFilePath, nil
}


