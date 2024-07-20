package login

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"io"
	"time"

	"compressURL/api/login/v1"
)

type GithubUserInfo struct {
	AvatarURL         string    `json:"avatar_url"`
	Bio               string    `json:"bio"`
	Company           string    `json:"company"`
	CreatedAt         time.Time `json:"created_at"`
	Email             *string   `json:"email"`
	EventsURL         string    `json:"events_url"`
	Followers         int       `json:"followers"`
	FollowersURL      string    `json:"followers_url"`
	Following         int       `json:"following"`
	FollowingURL      string    `json:"following_url"`
	GistsURL          string    `json:"gists_url"`
	GravatarID        string    `json:"gravatar_id"`
	Hireable          *bool     `json:"hireable"`
	HTMLURL           string    `json:"html_url"`
	ID                int       `json:"id"`
	Location          *string   `json:"location"`
	Login             string    `json:"login"`
	Name              string    `json:"name"`
	NodeID            string    `json:"node_id"`
	OrganizationsURL  string    `json:"organizations_url"`
	PublicGists       int       `json:"public_gists"`
	PublicRepos       int       `json:"public_repos"`
	ReceivedEventsURL string    `json:"received_events_url"`
	ReposURL          string    `json:"repos_url"`
	SiteAdmin         bool      `json:"site_admin"`
	StarredURL        string    `json:"starred_url"`
	SubscriptionsURL  string    `json:"subscriptions_url"`
	TwitterUsername   *string   `json:"twitter_username"`
	Type              string    `json:"type"`
	UpdatedAt         time.Time `json:"updated_at"`
	URL               string    `json:"url"`
}

func (c *ControllerV1) GithubLogin(ctx context.Context, req *v1.GithubLoginReq) (res *v1.GithubLoginRes, err error) {
	githubClientID, _ := g.Cfg().Get(ctx, "githubConfig.client_id")
	githubClientSecret, _ := g.Cfg().Get(ctx, "githubConfig.client_secret")

	githubOAuthConfig := &oauth2.Config{
		ClientID:     githubClientID.String(),     // 从 GitHub 获取
		ClientSecret: githubClientSecret.String(), // 从 GitHub 获取
		RedirectURL:  "http://localhost:5173/login",
		Scopes:       []string{"user:email"},
		Endpoint:     github.Endpoint,
	}

	token, err := githubOAuthConfig.Exchange(context.Background(), req.Code)
	if err != nil {
		g.Log().Error(ctx, "Failed to get token:", err)
		return nil, err
	}

	client := githubOAuthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		g.Log().Error(ctx, "Failed to get user:", err)
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	var user map[string]GithubUserInfo

	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		g.Log().Error(ctx, "Failed to parse user info:", err)
		return nil, err
	}

	g.Log().Debug(ctx, "User info:", user)

	return nil, nil
}
