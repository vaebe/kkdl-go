package login

import (
	"compressURL/internal/dao"
	"compressURL/internal/model"
	"compressURL/internal/model/entity"
	"compressURL/internal/service"
	"compressURL/utility"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"golang.org/x/net/context"
)

type sLogin struct {
}

func init() {
	service.RegisterLogin(New())
}

func New() *sLogin {
	return &sLogin{}
}

// GetUserInfo 获取用户信息
func (s *sLogin) GetUserInfo(ctx context.Context, in model.LoginInput) (entity.User, error) {
	userInfo := entity.User{}

	if in.AccountType == "01" {
		err := dao.User.Ctx(ctx).Where(dao.User.Columns().Email, in.Email).Scan(&userInfo)
		if err != nil {
			glog.Error(ctx, "GetUserInfo", err)
			return userInfo, gerror.New("用户不存在！")
		}
	} else {
		err := dao.User.Ctx(ctx).Where(dao.User.Columns().WxId, in.WxId).Scan(&userInfo)
		if err != nil {
			glog.Error(ctx, "GetUserInfo", err)
			return userInfo, gerror.New("用户不存在！")
		}
	}

	if utility.EncryptPassword(in.Password, userInfo.Salt) != userInfo.Password {
		glog.Error(ctx, "GetUserInfo", "账号或者密码不正确!")
		return userInfo, gerror.New("账号或者密码不正确!")
	}

	return userInfo, nil
}

// UserLogin 用户登录
func (s *sLogin) UserLogin(ctx context.Context, in model.LoginInput) (userInfo entity.User, token string, tokenExpire string, err error) {
	userInfo, err = service.Login().GetUserInfo(ctx, in)
	if err != nil {
		return entity.User{}, "", "", err
	}

	token, expire := service.Auth().AuthInstance().LoginHandler(ctx)
	tokenExpire = gtime.NewFromTime(expire).Format("Y-m-d H:i:s")

	return userInfo, token, tokenExpire, nil
}

// SignOutLogin 退出登录
func (s *sLogin) SignOutLogin(ctx context.Context) error {
	service.Auth().AuthInstance().LogoutHandler(ctx)
	// 用户退出登录删除 redis token
	return nil
}
