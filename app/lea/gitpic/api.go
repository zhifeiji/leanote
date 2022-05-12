package gitpic

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/zhifeiji/leanote/app/lea"
)

type GitHubAPI struct {
	owner string
	repo  string
	token string
}

func NewGitHubAPI(owner, repo, token string) *GitHubAPI {
	return &GitHubAPI{
		owner: owner,
		repo:  repo,
		token: token,
	}
}

func (api *GitHubAPI) UploadFile(path string, content []byte) (err error) {
	lea.Logf("begin UploadFile")
	ref, err := api.getRef()
	if err != nil {
		return
	}
	lea.Logf("step 1, getRef:%+v", ref)
	commit, err := api.getCommit(ref.Object.Sha)
	if err != nil {
		return
	}
	lea.Logf("step 2, getCommit:%+v", commit)
	blob, err := api.newBlob(base64.StdEncoding.EncodeToString(content), "base64")
	if err != nil {
		return
	}
	lea.Logf("step 3, newBlob:%+v", blob)
	tree, err := api.newTree(commit.Sha, path, blob.Sha)
	if err != nil {
		return
	}
	lea.Logf("step 4, newTree:%+v", tree)
	newCommit, err := api.newCommit(path, commit.Sha, tree.Sha)
	if err != nil {
		return
	}
	lea.Logf("step 5, newCommit:%+v", newCommit)
	newRef, err := api.updateRef(newCommit.Sha)
	if err != nil {
		return
	}
	lea.Logf("step 6, updateRef:%+v", newRef)
	return nil
}

//Ref ...
//{
//  "ref": "refs/heads/main",
//  "node_id": "REF_kwDOHTb-mK9yZWZzL2hlYWRzL21haW4",
//  "url": "https://api.github.com/repos/pictmp13/tm23456/git/refs/heads/main",
//  "object": {
//    "sha": "ab32d3a4ccb6084ad8dcdb3de0ade1ee5a121b4a",
//    "type": "commit",
//    "url": "https://api.github.com/repos/pictmp13/tm23456/git/commits/ab32d3a4ccb6084ad8dcdb3de0ade1ee5a121b4a"
//  }
//}
type Ref struct {
	Ref    string `json:"ref"`
	NodeID string `json:"node_id"`
	URL    string `json:"url"`
	Object Object `json:"object"`
}

type Object struct {
	Sha  string `json:"sha"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

//getRef
//获取 Ref
//curl \
//-H "Accept: application/vnd.github.v3+json" \
//  -H "Authorization: token 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 " \
//https://api.github.com/repos/pictmp13/tm23456/git/refs/heads/main
//
//{
//  "ref": "refs/heads/main",
//  "node_id": "REF_kwDOHTb-mK9yZWZzL2hlYWRzL21haW4",
//  "url": "https://api.github.com/repos/pictmp13/tm23456/git/refs/heads/main",
//  "object": {
//    "sha": "ab32d3a4ccb6084ad8dcdb3de0ade1ee5a121b4a",
//    "type": "commit",
//    "url": "https://api.github.com/repos/pictmp13/tm23456/git/commits/ab32d3a4ccb6084ad8dcdb3de0ade1ee5a121b4a"
//  }
//}
func (api *GitHubAPI) getRef() (*Ref, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/git/refs/heads/main", api.owner, api.repo)
	ref := &Ref{}
	err := api.call(http.MethodGet, url, nil, ref)
	return ref, err
}

//Commit
//{
//  "sha": "ab32d3a4ccb6084ad8dcdb3de0ade1ee5a121b4a",
//  "node_id": "C_kwDOHTb-mNoAKGFiMzJkM2E0Y2NiNjA4NGFkOGRjZGIzZGUwYWRlMWVlNWExMjFiNGE",
//  "url": "https://api.github.com/repos/pictmp13/tm23456/git/commits/ab32d3a4ccb6084ad8dcdb3de0ade1ee5a121b4a",
//  "html_url": "https://github.com/pictmp13/tm23456/commit/ab32d3a4ccb6084ad8dcdb3de0ade1ee5a121b4a",
//  "author": {
//    "name": "pictmp13",
//    "email": "62287395+pictmp13@users.noreply.github.com",
//    "date": "2022-05-09T05:13:06Z"
//  },
//  "committer": {
//    "name": "GitHub",
//    "email": "noreply@github.com",
//    "date": "2022-05-09T05:13:06Z"
//  },
//  "tree": {
//    "sha": "0d8d8967eba9ebe04308aaa1003b5360e35f6930",
//    "url": "https://api.github.com/repos/pictmp13/tm23456/git/trees/0d8d8967eba9ebe04308aaa1003b5360e35f6930"
//  },
//  "message": "Initial commit",
//  "parents": [
//
//  ],
//  "verification": {
//    "verified": true,
//    "reason": "valid",
//  }
//}
type Commit struct {
	Sha     string   `json:"sha"`
	Tree    Tree     `json:"tree"`
	parents []string `json:"parents"`
}

//getCommit
//2. 获取 Commit
//
//curl \
//  -H "Accept: application/vnd.github.v3+json" \
//  -H "Authorization: token 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 " \
//  https://api.github.com/repos/pictmp13/tm23456/git/commits/ab32d3a4ccb6084ad8dcdb3de0ade1ee5a121b4a
//
//{
//  "sha": "ab32d3a4ccb6084ad8dcdb3de0ade1ee5a121b4a",
//  "node_id": "C_kwDOHTb-mNoAKGFiMzJkM2E0Y2NiNjA4NGFkOGRjZGIzZGUwYWRlMWVlNWExMjFiNGE",
//  "url": "https://api.github.com/repos/pictmp13/tm23456/git/commits/ab32d3a4ccb6084ad8dcdb3de0ade1ee5a121b4a",
//  "html_url": "https://github.com/pictmp13/tm23456/commit/ab32d3a4ccb6084ad8dcdb3de0ade1ee5a121b4a",
//  "author": {
//    "name": "pictmp13",
//    "email": "62287395+pictmp13@users.noreply.github.com",
//    "date": "2022-05-09T05:13:06Z"
//  },
//  "committer": {
//    "name": "GitHub",
//    "email": "noreply@github.com",
//    "date": "2022-05-09T05:13:06Z"
//  },
//  "tree": {
//    "sha": "0d8d8967eba9ebe04308aaa1003b5360e35f6930",
//    "url": "https://api.github.com/repos/pictmp13/tm23456/git/trees/0d8d8967eba9ebe04308aaa1003b5360e35f6930"
//  },
//  "message": "Initial commit",
//  "parents": [
//
//  ],
//  "verification": {
//    "verified": true,
//    "reason": "valid",
//  }
//}
func (api *GitHubAPI) getCommit(commitSha string) (*Commit, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/git/commits/%s", api.owner, api.repo, commitSha)
	commit := &Commit{}
	err := api.call(http.MethodGet, url, nil, commit)
	return commit, err
}

type BlobContent struct {
	Content  string `json:"content"`
	Encoding string `json:"encoding"`
}

type Blob struct {
	Sha string `json:"sha"`
}

//newBlob
//3. 生成 Blob
//
//curl \
//  -X POST \
//  -H "Accept: application/vnd.github.v3+json" \
//  -H "Authorization: token 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 " \
//  https://api.github.com/repos/pictmp13/tm23456/git/blobs \
//  -d '{"content":"Content of the blob","encoding":"utf-8"}'
//
//{
//  "sha": "929246f65aab4d636cb229c790f966afc332c124",
//  "url": "https://api.github.com/repos/pictmp13/tm23456/git/blobs/929246f65aab4d636cb229c790f966afc332c124"
//}
func (api *GitHubAPI) newBlob(content string, encoding string) (*Blob, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/git/blobs", api.owner, api.repo)
	blob := &Blob{}
	err := api.call(http.MethodPost, url, &BlobContent{Content: content, Encoding: encoding}, blob)
	return blob, err
}

const (
	defaultMode = "100644"
	defaultType = "blob"
)

type TreeItem struct {
	Path string `json:"path"`
	Mode string `json:"mode"`
	Type string `json:"type"`
	Sha  string `json:"sha"`
}

type TreeContent struct {
	BaseTree string     `json:"base_tree"`
	Tree     []TreeItem `json:"tree"`
}

type Tree struct {
	Sha string `json:"sha"`
}

//newTree
//4. 生成 tree
//
//curl \
//  -X POST \
//  -H "Accept: application/vnd.github.v3+json" \
//  -H "Authorization: token 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 " \
//  https://api.github.com/repos/pictmp13/tm23456/git/trees \
//  -d '{"base_tree":"ab32d3a4ccb6084ad8dcdb3de0ade1ee5a121b4a","tree":[{"path":"tmp/123.txt","mode":"100644","type":"blob","sha":"929246f65aab4d636cb229c790f966afc332c124"}]}'
//
//{
//  "sha": "f961f7cad5169d131923f4b75af65df9c64b531a",
//  "url": "https://api.github.com/repos/pictmp13/tm23456/git/trees/f961f7cad5169d131923f4b75af65df9c64b531a",
//  "tree": [
//    {
//      "path": "README.md",
//      "mode": "100644",
//      "type": "blob",
//      "sha": "5c8cf6841d6b94fbb94f861e46c9e559fd1853bb",
//      "size": 9,
//      "url": "https://api.github.com/repos/pictmp13/tm23456/git/blobs/5c8cf6841d6b94fbb94f861e46c9e559fd1853bb"
//    },
//    {
//      "path": "tmp",
//      "mode": "040000",
//      "type": "tree",
//      "sha": "cd6dc376e92f394d53113a22c251a3e70e6a9115",
//      "url": "https://api.github.com/repos/pictmp13/tm23456/git/trees/cd6dc376e92f394d53113a22c251a3e70e6a9115"
//    }
//  ],
//  "truncated": false
//}
func (api *GitHubAPI) newTree(baseTree, path, sha string) (*Tree, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/git/trees", api.owner, api.repo)
	tree := &Tree{}
	err := api.call(http.MethodPost, url, &TreeContent{BaseTree: baseTree, Tree: []TreeItem{TreeItem{Path: path,
		Mode: defaultMode, Type: defaultType, Sha: sha}}}, tree)
	return tree, err
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Date  string `json:"date"`
}

type CommitContent struct {
	Message string   `json:"message"`
	Author  Author   `json:"author"`
	Parents []string `json:"parents"`
	Tree    string   `json:"tree"`
}

const (
	defaultAuthor = "pictmp13"
	defaultEmail  = "pictmp13@github.com"
)

//newCommit
//5. 生成 Commit
//
//curl \
//  -X POST \
//  -H "Accept: application/vnd.github.v3+json" \
//  -H "Authorization: token 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 " \
//  https://api.github.com/repos/pictmp13/tm23456/git/commits \
//  -d '{"message":"my commit message1111111","author":{"name":"pictmp13_user","email":"pictmp13@gmail.com", "date":"2022-01-09T16:13:30+12:00"},"parents":["ab32d3a4ccb6084ad8dcdb3de0ade1ee5a121b4a"],"tree":"f961f7cad5169d131923f4b75af65df9c64b531a"}'
//
//{
//  "sha": "9cca4a3ecd50d476428f2cf173abc234d834fd51",
//  "node_id": "C_kwDOHTb-mNoAKDljY2E0YTNlY2Q1MGQ0NzY0MjhmMmNmMTczYWJjMjM0ZDgzNGZkNTE",
//  "url": "https://api.github.com/repos/pictmp13/tm23456/git/commits/9cca4a3ecd50d476428f2cf173abc234d834fd51",
//  "html_url": "https://github.com/pictmp13/tm23456/commit/9cca4a3ecd50d476428f2cf173abc234d834fd51",
//  "author": {
//    "name": "pictmp13_user",
//    "email": "pictmp13@gmail.com",
//    "date": "2022-01-09T04:13:30Z"
//  },
//  "committer": {
//    "name": "pictmp13_user",
//    "email": "pictmp13@gmail.com",
//    "date": "2022-01-09T04:13:30Z"
//  },
//  "tree": {
//    "sha": "f961f7cad5169d131923f4b75af65df9c64b531a",
//    "url": "https://api.github.com/repos/pictmp13/tm23456/git/trees/f961f7cad5169d131923f4b75af65df9c64b531a"
//  },
//  "message": "my commit message1111111",
//  "parents": [
//    {
//      "sha": "ab32d3a4ccb6084ad8dcdb3de0ade1ee5a121b4a",
//      "url": "https://api.github.com/repos/pictmp13/tm23456/git/commits/ab32d3a4ccb6084ad8dcdb3de0ade1ee5a121b4a",
//      "html_url": "https://github.com/pictmp13/tm23456/commit/ab32d3a4ccb6084ad8dcdb3de0ade1ee5a121b4a"
//    }
//  ],
//  "verification": {
//    "verified": false,
//    "reason": "unsigned",
//    "signature": null,
//    "payload": null
//  }
//}
func (api *GitHubAPI) newCommit(message, parent, tree string) (*Commit, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/git/commits", api.owner, api.repo)
	commit := &Commit{}
	err := api.call(http.MethodPost, url, &CommitContent{
		Message: message,
		Author:  Author{Name: defaultAuthor, Email: defaultEmail, Date: time.Now().Format(time.RFC3339)},
		Parents: []string{parent},
		Tree:    tree,
	}, commit)
	return commit, err
}

type RefContent struct {
	Sha   string `json:"sha"`
	Force bool   `json:"force"`
}

//updateRef
//6. 更新 Ref
//
//curl \
//  -X PATCH \
//  -H "Accept: application/vnd.github.v3+json" \
//  -H "Authorization: token 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 " \
//  https://api.github.com/repos/pictmp13/tm23456/git/refs/heads/main \
//  -d '{"sha":"9cca4a3ecd50d476428f2cf173abc234d834fd51","force":true}'
//
//
//{
//  "ref": "refs/heads/main",
//  "node_id": "REF_kwDOHTb-mK9yZWZzL2hlYWRzL21haW4",
//  "url": "https://api.github.com/repos/pictmp13/tm23456/git/refs/heads/main",
//  "object": {
//    "sha": "9cca4a3ecd50d476428f2cf173abc234d834fd51",
//    "type": "commit",
//    "url": "https://api.github.com/repos/pictmp13/tm23456/git/commits/9cca4a3ecd50d476428f2cf173abc234d834fd51"
//  }
//}
func (api *GitHubAPI) updateRef(sha string) (*Ref, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/git/refs/heads/main", api.owner, api.repo)
	ref := &Ref{}
	err := api.call(http.MethodPatch, url, &RefContent{Sha: sha, Force: true}, ref)
	return ref, err
}

func (api *GitHubAPI) call(method string, url string, body interface{}, res interface{}) error {
	reqBody := &strings.Reader{}
	if body != nil {
		byteBody, err := json.Marshal(body)
		if err != nil {
			return err
		}

		reqBody = strings.NewReader(string(byteBody))
	}
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", api.token))

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return fmt.Errorf("status is %d, body: %s", resp.StatusCode, string(respByte))
	}

	if len(respByte) > 0 {
		return json.Unmarshal(respByte, res)
	}
	return nil
}
