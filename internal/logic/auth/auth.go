package auth

import (
	"compressURL/internal/model"
	"compressURL/internal/model/entity"
	"compressURL/internal/service"
	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"golang.org/x/net/context"
	"time"
)

type sAuth struct {
}

func init() {
	initAuth()
	service.RegisterAuth(New())
}

func New() *sAuth {
	return &sAuth{}
}

var authInstance *jwt.GfJWTMiddleware

func initAuth() {
	authInstance = jwt.New(&jwt.GfJWTMiddleware{
		Realm:           "short_url",
		Key:             []byte("short_url_byd"),
		Timeout:         time.Hour * 5,
		MaxRefresh:      time.Hour * 5,
		IdentityKey:     "LoginUserId",
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		Authenticator:   Authenticator,
		Unauthorized:    Unauthorized,
		PayloadFunc:     PayloadFunc,
		IdentityHandler: IdentityHandler,
	})
}

func (s sAuth) AuthInstance() *jwt.GfJWTMiddleware {
	return authInstance
}

// GetLoginUserInfo 获取当前登录用户的信息
func (s sAuth) GetLoginUserInfo(ctx context.Context) (model.JWTPayloadInfo, error) {

	jwtInfo, token, err := authInstance.GetClaimsFromJWT(ctx)

	glog.Info(ctx, jwtInfo, token, err)

	// 为空直接返回
	if jwtInfo == nil {
		return model.JWTPayloadInfo{}, err
	}

	info := model.JWTPayloadInfo{
		Id:          jwtInfo["LoginUserId"].(string),
		AccountType: jwtInfo["LoginUserAccountType"].(string),
		Role:        jwtInfo["LoginUserRole"].(string),
		Token:       token,
	}

	return info, err
}

// PayloadFunc is a callback function that will be called during login.
// Using this function it is possible to add additional payload data to the webtoken.
// The data is then made available during requests via c.Get("JWT_PAYLOAD").
// Note that the payload is not encrypted.
// The attributes mentioned on jwt.io can't be used as keys for the map.
// Optional, by default no additional data will be set.
func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}

	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

// IdentityHandler get the identity from JWT and set the identity for every request
// Using this function, by r.GetParam("id") get identity
func IdentityHandler(ctx context.Context) interface{} {
	claims := jwt.ExtractClaims(ctx)
	return claims[authInstance.IdentityKey]
}

// Unauthorized is used to define customized Unauthorized callback function.
func Unauthorized(ctx context.Context, code int, message string) {
	r := g.RequestFromCtx(ctx)
	r.Response.WriteJson(g.Map{
		"code":    code,
		"message": message,
	})
	r.ExitAll()
}

// Authenticator is used to validate login parameters.
// It must return user data as user identifier, it will be stored in Claim Array.
// if your identityKey is 'id', your user data must have 'id'
// Check error (e) to determine the appropriate error message.
func Authenticator(ctx context.Context) (interface{}, error) {
	var loginInfo entity.User

	err := g.RequestFromCtx(ctx).GetCtxVar("loginInfo").Scan(&loginInfo)

	if err != nil {
		return nil, err
	}

	data := g.Map{
		"LoginUserId":          loginInfo.Id,
		"LoginUserAccountType": loginInfo.AccountType,
		"LoginUserRole":        loginInfo.Role,
	}

	return data, nil
}
