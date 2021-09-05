package api

/*
Frontend
- PR View
	- View unified view / view split view with comments
	- View unified view / view split view given fileDiffHeader


CodeReviewService: depends on a repo
- List Pull requests
	- List PRs that need attention
	- List "Opened By Me"
	- List "Reviewed By Me"
	- List "Merged"
	- Find PR by ID ( should include file diff headers & stats)
- PR Actions
	- Open PR
	- Approve PR
	- Request Changes PR
	- Abandon PR
	- Merge PR
		- List Merge Options (rebase, squash, etc.)
- Reviewer
	- Add reviewers to PR
	- Edit reviewers to PR
	- List reviewers
- Comment
	- Leave comments
		- Highlight on certain words
		- Stack comments into threads
		- Group thread for same selected code
	- Edit comments
	- Delete comments
	- Mark thread resolved
	- Mark thread unresolved
	- List threads

TeamService
- List all repos owned by a team
- List all teams for a user
- List members for a team
- Add user to a team
- Create a team
- Delete a team
- Create a repo

UserService
- List all repos owned by the user
- List all users
- Create a repo

RepoService
- List all branches for a repo
- List change history for a certain branch
- optional
	- code search

DiffService: depends on a repo
- List file diffs given FromBranchName, ToBranchName
- Show statistics of diff

// Not yet
// web API -
//
// code review service - stitching view and code review service together

*/