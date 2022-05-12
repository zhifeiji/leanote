package gitpic

import (
	"reflect"
	"testing"
)

const (
	testOwner = ""
	testRepo  = ""
	testToken = ""
)

func TestGitHubAPI_UploadFile(t *testing.T) {
	type fields struct {
		owner string
		repo  string
		token string
	}
	type args struct {
		path    string
		content []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &GitHubAPI{
				owner: tt.fields.owner,
				repo:  tt.fields.repo,
				token: tt.fields.token,
			}
			if err := api.UploadFile(tt.args.path, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("UploadFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGitHubAPI_call(t *testing.T) {
	type fields struct {
		owner string
		repo  string
		token string
	}
	type args struct {
		method string
		url    string
		body   interface{}
		res    interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &GitHubAPI{
				owner: tt.fields.owner,
				repo:  tt.fields.repo,
				token: tt.fields.token,
			}
			if err := api.call(tt.args.method, tt.args.url, tt.args.body, tt.args.res); (err != nil) != tt.wantErr {
				t.Errorf("call() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGitHubAPI_getCommit(t *testing.T) {
	type fields struct {
		owner string
		repo  string
		token string
	}
	type args struct {
		commitSha string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Commit
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "正常", fields: fields{owner: testOwner, repo: testRepo, token: testToken},
			args: args{commitSha: "9cca4a3ecd50d476428f2cf173abc234d834fd51"},
			want: &Commit{Sha: "9cca4a3ecd50d476428f2cf173abc234d834fd51", Tree: Tree{Sha: "f961f7cad5169d131923f4b75af65df9c64b531a"}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &GitHubAPI{
				owner: tt.fields.owner,
				repo:  tt.fields.repo,
				token: tt.fields.token,
			}
			got, err := api.getCommit(tt.args.commitSha)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCommit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCommit() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGitHubAPI_getRef(t *testing.T) {
	type fields struct {
		owner string
		repo  string
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Ref
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "正常", fields: fields{owner: testOwner, repo: testRepo, token: testToken},
			want: &Ref{Ref: "refs/heads/main", NodeID: "REF_kwDOHTb-mK9yZWZzL2hlYWRzL21haW4", URL: "https://api.github.com/repos/pictmp13/tm23456/git/refs/heads/main",
				Object: Object{Sha: "9cca4a3ecd50d476428f2cf173abc234d834fd51", Type: "commit", URL: "https://api.github.com/repos/pictmp13/tm23456/git/commits/9cca4a3ecd50d476428f2cf173abc234d834fd51"}},
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &GitHubAPI{
				owner: tt.fields.owner,
				repo:  tt.fields.repo,
				token: tt.fields.token,
			}
			got, err := api.getRef()
			if (err != nil) != tt.wantErr {
				t.Errorf("getRef() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRef() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGitHubAPI_newBlob(t *testing.T) {
	type fields struct {
		owner string
		repo  string
		token string
	}
	type args struct {
		content  string
		encoding string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Blob
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "正常", fields: fields{owner: testOwner, repo: testRepo, token: testToken}, args: args{content: "aGVsbG8=",
			encoding: "base64"}, want: &Blob{Sha: "b6fc4c620b67d95f953a5c1c1230aaab5db5a1b0"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &GitHubAPI{
				owner: tt.fields.owner,
				repo:  tt.fields.repo,
				token: tt.fields.token,
			}
			got, err := api.newBlob(tt.args.content, tt.args.encoding)
			if (err != nil) != tt.wantErr {
				t.Errorf("newBlob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newBlob() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGitHubAPI_newCommit(t *testing.T) {
	type fields struct {
		owner string
		repo  string
		token string
	}
	type args struct {
		message string
		parent  string
		tree    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Commit
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "正常", fields: fields{owner: testOwner, repo: testRepo, token: testToken},
			args: args{message: "test message", parent: "9cca4a3ecd50d476428f2cf173abc234d834fd51",
				tree: "d8905c369d62da138be0b2a94f7af1e4f86e4bf5"},
			want: &Commit{Sha: "f04f69778c86ccc0c0856a5b884a90a0f2b36ce9", Tree: Tree{Sha: "d8905c369d62da138be0b2a94f7af1e4f86e4bf5"}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &GitHubAPI{
				owner: tt.fields.owner,
				repo:  tt.fields.repo,
				token: tt.fields.token,
			}
			got, err := api.newCommit(tt.args.message, tt.args.parent, tt.args.tree)
			if (err != nil) != tt.wantErr {
				t.Errorf("newCommit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newCommit() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGitHubAPI_newTree(t *testing.T) {
	type fields struct {
		owner string
		repo  string
		token string
	}
	type args struct {
		baseTree string
		path     string
		sha      string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Tree
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "正常", fields: fields{owner: testOwner, repo: testRepo, token: testToken},
			args: args{baseTree: "9cca4a3ecd50d476428f2cf173abc234d834fd51", path: "tmp/a.md",
				sha: "b6fc4c620b67d95f953a5c1c1230aaab5db5a1b0"},
			want: &Tree{Sha: "d8905c369d62da138be0b2a94f7af1e4f86e4bf5"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &GitHubAPI{
				owner: tt.fields.owner,
				repo:  tt.fields.repo,
				token: tt.fields.token,
			}
			got, err := api.newTree(tt.args.baseTree, tt.args.path, tt.args.sha)
			if (err != nil) != tt.wantErr {
				t.Errorf("newTree() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTree() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGitHubAPI_updateRef(t *testing.T) {
	type fields struct {
		owner string
		repo  string
		token string
	}
	type args struct {
		sha string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Ref
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "正常", fields: fields{owner: testOwner, repo: testRepo, token: testToken},
			args: args{sha: "f04f69778c86ccc0c0856a5b884a90a0f2b36ce9"}, want: &Ref{Ref: "refs/heads/main",
				NodeID: "REF_kwDOHTb-mK9yZWZzL2hlYWRzL21haW4",
				URL:    "https://api.github.com/repos/pictmp13/tm23456/git/refs/heads/main",
				Object: Object{Sha: "f04f69778c86ccc0c0856a5b884a90a0f2b36ce9", Type: "commit",
					URL: "https://api.github." +
						"com/repos/pictmp13/tm23456/git/commits/f04f69778c86ccc0c0856a5b884a90a0f2b36ce9"}},
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &GitHubAPI{
				owner: tt.fields.owner,
				repo:  tt.fields.repo,
				token: tt.fields.token,
			}
			got, err := api.updateRef(tt.args.sha)
			if (err != nil) != tt.wantErr {
				t.Errorf("updateRef() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateRef() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewGitHubAPI(t *testing.T) {
	type args struct {
		owner string
		repo  string
		token string
	}
	tests := []struct {
		name string
		args args
		want *GitHubAPI
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGitHubAPI(tt.args.owner, tt.args.repo, tt.args.token); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGitHubAPI() = %v, want %v", got, tt.want)
			}
		})
	}
}
