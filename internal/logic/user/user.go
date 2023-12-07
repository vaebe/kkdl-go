package user

import (
	v1 "compressURL/api/user/v1"
	"compressURL/internal/dao"
	"compressURL/internal/model"
	"compressURL/internal/service"
	"compressURL/utility"
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
	in.Salt = grand.S(10)
	in.Password = utility.EncryptPassword(in.Password, in.Salt)

	id, err := dao.User.Ctx(ctx).Data(in).InsertAndGetId()
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
	_, err := dao.User.Ctx(ctx).Data(in).
		FieldsEx(dao.User.Columns().Id).
		Where(dao.User.Columns().Id, in.Id).
		Update()
	return err
}

func (s *sUser) GetUserInfo(ctx context.Context, id int64) (info *v1.GetUserInfoRes, error error) {
	userInfo := v1.GetUserInfoRes{}

	err := dao.User.Ctx(ctx).Where(dao.User.Columns().Id, id).Scan(&userInfo)

	g.Log().Info(ctx, &userInfo)

	if err != nil {
		return &userInfo, err
	}
	return &userInfo, nil
}
