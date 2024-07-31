package login

import (
	"compressURL/api/login/v1"
	"compressURL/internal/model"
	"compressURL/internal/model/entity"
	"compressURL/internal/service"
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"net/http"
	"strconv"
	"time"
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

func getGithubUserInfo(ctx context.Context, code string) (*GithubUserInfo, error) {
	cfg := g.Cfg()
	githubOAuthConfig := &oauth2.Config{
		ClientID:     cfg.MustGet(ctx, "githubConfig.client_id").String(),
		ClientSecret: cfg.MustGet(ctx, "githubConfig.client_secret").String(),
		RedirectURL:  cfg.MustGet(ctx, "githubConfig.redirect_uri").String(),
		Scopes:       []string{"user"},
		Endpoint:     github.Endpoint,
	}

	token, err := githubOAuthConfig.Exchange(ctx, code)
	if err != nil {
		return nil, gerror.New("获取 github token失败！")
	}

	client := githubOAuthConfig.Client(ctx, token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		return nil, gerror.New("获取 github 用户信息失败！")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, gerror.New("无法获取 github 用户信息！")
	}

	var userInfo GithubUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		g.Log().Error(ctx, "Failed to parse user info:", err)
		return nil, gerror.New("格式化 github 用户信息失败！")
	}

	return &userInfo, nil
}

// GithubLogin github登录
func (c *ControllerV1) GithubLogin(ctx context.Context, req *v1.GithubLoginReq) (*v1.GithubLoginRes, error) {
	githubUserInfo, err := getGithubUserInfo(ctx, req.Code)
	if err != nil {
		return nil, err
	}

	// 获取用户信息
	userInfo, err := service.User().Detail(ctx, model.UserQueryInput{Id: strconv.Itoa(githubUserInfo.ID)})

	// 用户不存在则创建用户
	if userInfo == nil && err == nil {
		userInfo = &entity.User{
			Id:          strconv.Itoa(githubUserInfo.ID),
			NickName:    githubUserInfo.Login,
			Role:        "01",
			Avatar:      githubUserInfo.AvatarURL,
			AccountType: "03",
		}

		if _, err = service.User().Create(ctx, *userInfo); err != nil {
			return nil, err
		}
	}

	if err != nil {
		return nil, err
	}

	info := getLoginRes(ctx, *userInfo)
	return (*v1.GithubLoginRes)(info), nil
}
