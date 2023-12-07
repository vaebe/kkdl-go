package user

import (
	v1 "compressURL/api/user/v1"
	"compressURL/internal/dao"
	"compressURL/internal/model"
	"compressURL/internal/model/entity"
	"compressURL/internal/service"
	"compressURL/utility"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
	"golang.org/x/net/context"
)

type sUser struct {
}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

func (s *sUser) Create(ctx context.Context, in model.UserCreateInput) (userid int64, error error) {
	// 昵称不存在生成默认昵称
	if in.Nickname == "" {
		in.Nickname = "UU" + grand.S(4)
	}

	// 生成随机盐值
	in.Salt = grand.S(10)

	// Email、WxId 只会存在一个切不可为空
	data := g.Map{
		"Email":       nil,
		"WxId":        nil,
		"Password":    utility.EncryptPassword(in.Password, in.Salt),
		"Nickname":    in.Nickname,
		"AccountType": in.AccountType,
		"Role":        in.Role,
		"Salt":        in.Salt,
	}

	if in.WxId != "" {
		data["WxId"] = in.WxId
	}

	if in.Email != "" {
		data["Email"] = in.Email
	}

	id, err := dao.User.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *sUser) Remove(ctx context.Context, id int64) error {
	_, err := dao.User.Ctx(ctx).Where(dao.User.Columns().Id, id).Delete()
	return err
}

func (s *sUser) Update(ctx context.Context, in model.UserUpdateInput) error {
	// 获取用户信息
	userInfo, err := service.User().Detail(ctx, in.Id)

	if err != nil {
		return err
	}

	// 对修改后的密码加密
	in.Password = utility.EncryptPassword(in.Password, userInfo.Salt)

	// 更新数据
	_, err = dao.User.Ctx(ctx).Data(in).
		FieldsEx(dao.User.Columns().Id).
		Where(dao.User.Columns().Id, in.Id).
		Update()
	return err
}

func (s *sUser) GetUserInfo(ctx context.Context, id int64) (info *v1.GetUserInfoRes, error error) {
	userInfo := v1.GetUserInfoRes{}

	err := dao.User.Ctx(ctx).Where(dao.User.Columns().Id, id).Scan(&userInfo)

	if err != nil {
		return nil, errors.New("未查询到用户数据！")
	}
	return &userInfo, nil
}

// Detail 获取用户详情
func (s *sUser) Detail(ctx context.Context, id int64) (info entity.User, error error) {
	userInfo := entity.User{}

	err := dao.User.Ctx(ctx).Where(dao.User.Columns().Id, id).Scan(&userInfo)

	if err != nil {
		return userInfo, errors.New("未查询到用户数据！")
	}
	return userInfo, nil
}
