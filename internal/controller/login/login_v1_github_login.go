package login

import (
	"compressURL/internal/model/entity"
	"compressURL/internal/service"
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"io"
	"strconv"
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

func getGithubUserInfo(ctx context.Context, req *v1.GithubLoginReq) (res *GithubUserInfo, err error) {
	githubClientID, _ := g.Cfg().Get(ctx, "githubConfig.client_id")
	githubClientSecret, _ := g.Cfg().Get(ctx, "githubConfig.client_secret")
	githubRedirectURL, _ := g.Cfg().Get(ctx, "githubConfig.redirect_uri")

	githubOAuthConfig := &oauth2.Config{
		ClientID:     githubClientID.String(),     // 从 GitHub 获取
		ClientSecret: githubClientSecret.String(), // 从 GitHub 获取
		RedirectURL:  githubRedirectURL.String(),
		Scopes:       []string{"user"}, // 客户端 https://github.com/login/oauth/authorize 不设置 scope 参数这里配置无效
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

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		g.Log().Error(ctx, "Failed to parse user info:", err)
		return nil, err
	}

	return res, nil
}

// GithubLogin github登录
func (c *ControllerV1) GithubLogin(ctx context.Context, req *v1.GithubLoginReq) (res *v1.GithubLoginRes, err error) {
	githubUserInfo, err := getGithubUserInfo(ctx, req)
	g.Log().Debug(ctx, githubUserInfo)

	if err != nil {
		return nil, err
	}

	// 获取用户信息
	userInfo, err := service.User().Detail(ctx, strconv.Itoa(githubUserInfo.ID))

	// 用户不存在则创建用户
	if err != nil {
		userInfo = entity.User{
			Id:          strconv.Itoa(githubUserInfo.ID),
			NickName:    githubUserInfo.Login,
			Role:        "01",
			Avatar:      githubUserInfo.AvatarURL,
			AccountType: "03",
		}

		_, err = service.User().Create(ctx, userInfo)

		if err != nil {
			return nil, err
		}
	}

	// 设置登录用户信息
	g.RequestFromCtx(ctx).SetCtxVar("loginInfo", userInfo)
	token, expire := service.Auth().AuthInstance().LoginHandler(ctx)
	tokenExpire := gtime.NewFromTime(expire).Format("Y-m-d H:i:s")

	return &v1.GithubLoginRes{
		Token:       token,
		TokenExpire: tokenExpire,
		UserInfo: entity.User{
			Id:          userInfo.Id,
			Email:       userInfo.Email,
			NickName:    userInfo.NickName,
			AccountType: userInfo.AccountType,
			Role:        userInfo.Role,
			Avatar:      userInfo.Avatar,
		},
	}, nil
}
