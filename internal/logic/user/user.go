package user

import (
	v1 "compressURL/api/user/v1"
	"compressURL/internal/dao"
	"compressURL/internal/model/entity"
	"compressURL/internal/service"
	"compressURL/utility"
	"errors"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/gogf/gf/v2/util/guid"
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

func (s *sUser) Create(ctx context.Context, in entity.User) (string, error) {
	// 昵称不存在生成默认昵称
	if in.NickName == "" {
		in.NickName = "kk" + grand.S(4)
	}

	// 生成随机盐值
	in.Salt = grand.S(10)

	userId := guid.S()

	// Email、WxId 只会存在一个且不可为空
	data := g.Map{
		"Id":          userId,
		"Email":       nil,
		"WxId":        nil,
		"Password":    utility.EncryptPassword(in.Password, in.Salt),
		"NickName":    in.NickName,
		"AccountType": nil,
		"Role":        in.Role,
		"Salt":        in.Salt,
	}

	if in.WxId != "" {
		data["WxId"] = in.WxId
		data["AccountType"] = "02"
	}

	if in.Email != "" {
		data["Email"] = in.Email
		data["AccountType"] = "01"
	}

	_, err := dao.User.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return "", err
	}
	return userId, nil
}

// Update 更新用户信息
func (s *sUser) Update(ctx context.Context, in entity.User) error {
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

// Remove 删除用户
func (s *sUser) Remove(ctx context.Context, id string) error {
	res, err := dao.User.Ctx(ctx).Where(dao.User.Columns().Id, id).Delete()

	if num, _ := res.RowsAffected(); num == 0 {
		return gerror.New("需要删除的数据不存在！")
	}

	return err
}

func (s *sUser) GetUserInfo(ctx context.Context, id string) (*v1.GetUserInfoRes, error) {
	userInfo := v1.GetUserInfoRes{}

	err := dao.User.Ctx(ctx).Where(dao.User.Columns().Id, id).Scan(&userInfo)

	if err != nil {
		return nil, errors.New("未查询到用户数据！")
	}
	return &userInfo, nil
}

// Detail 获取用户详情
func (s *sUser) Detail(ctx context.Context, id string) (entity.User, error) {
	userInfo := entity.User{}

	err := dao.User.Ctx(ctx).Where(dao.User.Columns().Id, id).Scan(&userInfo)

	if err != nil {
		return userInfo, errors.New("未查询到用户数据！")
	}
	return userInfo, nil
}

// GetUserList 获取用户列表
func (s *sUser) GetUserList(ctx context.Context, in v1.GetUserListReq) ([]entity.User, int, error) {
	var userList []entity.User

	db := dao.User.Ctx(ctx).OmitEmptyWhere().
		Where(dao.User.Columns().WxId, in.WxId).
		Where(dao.User.Columns().Email, in.Email).
		Where(dao.User.Columns().NickName, in.NickName)

	total, _ := db.Count()

	err := db.Page(in.PageNo, in.PageSize).Scan(&userList)

	if err != nil {
		return nil, 0, err
	}
	return userList, total, nil
}
