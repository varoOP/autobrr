package release_checker

import (
	"context"
	"testing"
)

func TestGitHubReleaseChecker_CheckLatest(t *testing.T) {
	type fields struct {
		Repo string
	}
	type args struct {
		ctx     context.Context
		version string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:   "test_1",
			fields: fields{Repo: "autobrr/autobrr"},
			args: args{
				ctx:     context.TODO(),
				version: "v0.26.0",
			},
			want:    false,
			wantErr: false,
		},
		{
			name:   "test_2",
			fields: fields{Repo: "autobrr/autobrr"},
			args: args{
				ctx:     context.TODO(),
				version: "v0.25.0",
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GitHubReleaseChecker{
				Repo: tt.fields.Repo,
			}
			got, err := g.CanUpdate(tt.args.ctx, tt.args.version)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckLatest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckLatest() got = %v, want %v", got, tt.want)
			}
		})
	}
}
