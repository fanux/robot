package main

import (
	"fmt"
	"github.com/fanux/robot/issue"
	"os"
	"testing"
)

func Test_payEvent(t *testing.T) {
	body := []byte(`{
  "action": "opened",
  "issue": {
    "url": "https://api.github.com/repos/fanux/sshcmd/issues/5",
    "repository_url": "https://api.github.com/repos/fanux/sshcmd",
    "labels_url": "https://api.github.com/repos/fanux/sshcmd/issues/5/labels{/name}",
    "comments_url": "https://api.github.com/repos/fanux/sshcmd/issues/5/comments",
    "events_url": "https://api.github.com/repos/fanux/sshcmd/issues/5/events",
    "html_url": "https://github.com/fanux/sshcmd/issues/5",
    "id": 581023143,
    "node_id": "MDU6SXNzdWU1ODEwMjMxNDM=",
    "number": 5,
    "title": "test pay comment",
    "user": {
      "login": "fanux",
      "id": 8912557,
      "node_id": "MDQ6VXNlcjg5MTI1NTc=",
      "avatar_url": "https://avatars2.githubusercontent.com/u/8912557?v=4",
      "gravatar_id": "",
      "url": "https://api.github.com/users/fanux",
      "html_url": "https://github.com/fanux",
      "followers_url": "https://api.github.com/users/fanux/followers",
      "following_url": "https://api.github.com/users/fanux/following{/other_user}",
      "gists_url": "https://api.github.com/users/fanux/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/fanux/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/fanux/subscriptions",
      "organizations_url": "https://api.github.com/users/fanux/orgs",
      "repos_url": "https://api.github.com/users/fanux/repos",
      "events_url": "https://api.github.com/users/fanux/events{/privacy}",
      "received_events_url": "https://api.github.com/users/fanux/received_events",
      "type": "User",
      "site_admin": false
    },
    "labels": [

    ],
    "state": "open",
    "locked": false,
    "assignee": null,
    "assignees": [

    ],
    "milestone": null,
    "comments": 0,
    "created_at": "2020-03-14T04:05:48Z",
    "updated_at": "2020-03-14T04:05:48Z",
    "closed_at": null,
    "author_association": "OWNER",
    "body": "/pay 55"
  },
  "repository": {
    "id": 214574711,
    "node_id": "MDEwOlJlcG9zaXRvcnkyMTQ1NzQ3MTE=",
    "name": "sshcmd",
    "full_name": "fanux/sshcmd",
    "private": false,
    "owner": {
      "login": "fanux",
      "id": 8912557,
      "node_id": "MDQ6VXNlcjg5MTI1NTc=",
      "avatar_url": "https://avatars2.githubusercontent.com/u/8912557?v=4",
      "gravatar_id": "",
      "url": "https://api.github.com/users/fanux",
      "html_url": "https://github.com/fanux",
      "followers_url": "https://api.github.com/users/fanux/followers",
      "following_url": "https://api.github.com/users/fanux/following{/other_user}",
      "gists_url": "https://api.github.com/users/fanux/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/fanux/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/fanux/subscriptions",
      "organizations_url": "https://api.github.com/users/fanux/orgs",
      "repos_url": "https://api.github.com/users/fanux/repos",
      "events_url": "https://api.github.com/users/fanux/events{/privacy}",
      "received_events_url": "https://api.github.com/users/fanux/received_events",
      "type": "User",
      "site_admin": false
    },
    "html_url": "https://github.com/fanux/sshcmd",
    "description": "instead sshpass",
    "fork": true,
    "url": "https://api.github.com/repos/fanux/sshcmd",
    "forks_url": "https://api.github.com/repos/fanux/sshcmd/forks",
    "keys_url": "https://api.github.com/repos/fanux/sshcmd/keys{/key_id}",
    "collaborators_url": "https://api.github.com/repos/fanux/sshcmd/collaborators{/collaborator}",
    "teams_url": "https://api.github.com/repos/fanux/sshcmd/teams",
    "hooks_url": "https://api.github.com/repos/fanux/sshcmd/hooks",
    "issue_events_url": "https://api.github.com/repos/fanux/sshcmd/issues/events{/number}",
    "events_url": "https://api.github.com/repos/fanux/sshcmd/events",
    "assignees_url": "https://api.github.com/repos/fanux/sshcmd/assignees{/user}",
    "branches_url": "https://api.github.com/repos/fanux/sshcmd/branches{/branch}",
    "tags_url": "https://api.github.com/repos/fanux/sshcmd/tags",
    "blobs_url": "https://api.github.com/repos/fanux/sshcmd/git/blobs{/sha}",
    "git_tags_url": "https://api.github.com/repos/fanux/sshcmd/git/tags{/sha}",
    "git_refs_url": "https://api.github.com/repos/fanux/sshcmd/git/refs{/sha}",
    "trees_url": "https://api.github.com/repos/fanux/sshcmd/git/trees{/sha}",
    "statuses_url": "https://api.github.com/repos/fanux/sshcmd/statuses/{sha}",
    "languages_url": "https://api.github.com/repos/fanux/sshcmd/languages",
    "stargazers_url": "https://api.github.com/repos/fanux/sshcmd/stargazers",
    "contributors_url": "https://api.github.com/repos/fanux/sshcmd/contributors",
    "subscribers_url": "https://api.github.com/repos/fanux/sshcmd/subscribers",
    "subscription_url": "https://api.github.com/repos/fanux/sshcmd/subscription",
    "commits_url": "https://api.github.com/repos/fanux/sshcmd/commits{/sha}",
    "git_commits_url": "https://api.github.com/repos/fanux/sshcmd/git/commits{/sha}",
    "comments_url": "https://api.github.com/repos/fanux/sshcmd/comments{/number}",
    "issue_comment_url": "https://api.github.com/repos/fanux/sshcmd/issues/comments{/number}",
    "contents_url": "https://api.github.com/repos/fanux/sshcmd/contents/{+path}",
    "compare_url": "https://api.github.com/repos/fanux/sshcmd/compare/{base}...{head}",
    "merges_url": "https://api.github.com/repos/fanux/sshcmd/merges",
    "archive_url": "https://api.github.com/repos/fanux/sshcmd/{archive_format}{/ref}",
    "downloads_url": "https://api.github.com/repos/fanux/sshcmd/downloads",
    "issues_url": "https://api.github.com/repos/fanux/sshcmd/issues{/number}",
    "pulls_url": "https://api.github.com/repos/fanux/sshcmd/pulls{/number}",
    "milestones_url": "https://api.github.com/repos/fanux/sshcmd/milestones{/number}",
    "notifications_url": "https://api.github.com/repos/fanux/sshcmd/notifications{?since,all,participating}",
    "labels_url": "https://api.github.com/repos/fanux/sshcmd/labels{/name}",
    "releases_url": "https://api.github.com/repos/fanux/sshcmd/releases{/id}",
    "deployments_url": "https://api.github.com/repos/fanux/sshcmd/deployments",
    "created_at": "2019-10-12T04:01:13Z",
    "updated_at": "2020-02-15T13:40:23Z",
    "pushed_at": "2020-02-15T13:40:21Z",
    "git_url": "git://github.com/fanux/sshcmd.git",
    "ssh_url": "git@github.com:fanux/sshcmd.git",
    "clone_url": "https://github.com/fanux/sshcmd.git",
    "svn_url": "https://github.com/fanux/sshcmd",
    "homepage": "",
    "size": 15,
    "stargazers_count": 1,
    "watchers_count": 1,
    "language": "Go",
    "has_issues": true,
    "has_projects": true,
    "has_downloads": true,
    "has_wiki": true,
    "has_pages": false,
    "forks_count": 0,
    "mirror_url": null,
    "archived": false,
    "disabled": false,
    "open_issues_count": 4,
    "license": {
      "key": "apache-2.0",
      "name": "Apache License 2.0",
      "spdx_id": "Apache-2.0",
      "url": "https://api.github.com/licenses/apache-2.0",
      "node_id": "MDc6TGljZW5zZTI="
    },
    "forks": 0,
    "open_issues": 4,
    "watchers": 1,
    "default_branch": "master"
  },
  "sender": {
    "login": "fanux",
    "id": 8912557,
    "node_id": "MDQ6VXNlcjg5MTI1NTc=",
    "avatar_url": "https://avatars2.githubusercontent.com/u/8912557?v=4",
    "gravatar_id": "",
    "url": "https://api.github.com/users/fanux",
    "html_url": "https://github.com/fanux",
    "followers_url": "https://api.github.com/users/fanux/followers",
    "following_url": "https://api.github.com/users/fanux/following{/other_user}",
    "gists_url": "https://api.github.com/users/fanux/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/fanux/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/fanux/subscriptions",
    "organizations_url": "https://api.github.com/users/fanux/orgs",
    "repos_url": "https://api.github.com/users/fanux/repos",
    "events_url": "https://api.github.com/users/fanux/events{/privacy}",
    "received_events_url": "https://api.github.com/users/fanux/received_events",
    "type": "User",
    "site_admin": false
  }
}`)

	type args struct {
		body []byte
	}
	tests := []struct {
		name string
		args args
		want *issue.IssueCommentEvent
	}{
		{
			"test /pay 55",
			args{body},
			nil,
		},
	}
	for _, tt := range tests {
		fmt.Println(os.Getenv("GITHUB_USER"), os.Getenv("GITHUB_PASSWD"))
		// 注意主进程先退出可能会评论失败，在函数里sleep一下解决，长时间运行的程序无此问题
		e := payEvent(tt.args.body)
		if *e.Repo.FullName != "fanux/sshcmd" {
			t.Errorf("%s",*e.Repo.FullName)
		}
	}
}