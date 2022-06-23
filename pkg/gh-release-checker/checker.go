package release_checker

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/jbowes/semver"
)

// Release is a GitHub release
type Release struct {
	Draft      bool   `json:"draft"`
	Prerelease bool   `json:"prerelease"`
	TagName    string `json:"tag_name"`
}

func (r *Release) IsPreOrDraft() bool {
	if r.Draft || r.Prerelease {
		return true
	}
	return false
}

type GitHubReleaseChecker struct {
	// user/repo-name or org/repo-name
	Repo string
}

func (g *GitHubReleaseChecker) CanUpdate(ctx context.Context, version string) (bool, error) {
	releases, err := g.get(ctx)
	if err != nil {
		return false, err
	}

	for _, release := range releases {
		if release.IsPreOrDraft() {
			continue
		}

		currentTag := stripVPrefix(version)
		currentVersion, err := semver.Parse(currentTag)
		if err != nil {
			return false, err
		}

		latestTag := stripVPrefix(release.TagName)
		latestVersion, err := semver.Parse(latestTag)
		if err != nil {
			return false, err
		}

		v := latestVersion.Compare(currentVersion)
		if v == 1 {
			return true, nil
		}
	}

	return false, nil
}

func (g *GitHubReleaseChecker) get(ctx context.Context) ([]Release, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/releases", g.Repo)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")
	c := http.DefaultClient

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error getting releases for %v: %s", g.Repo, resp.Status)
	}

	var releases []Release
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&releases); err != nil {
		return nil, err
	}

	return releases, nil
}

func stripVPrefix(tag string) string {
	return strings.TrimPrefix(tag, "v")
}
