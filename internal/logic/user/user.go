package user

import (
	v1 "compressURL/api/user/v1"
	"compressURL/internal/dao"
	"compressURL/internal/model"
	"compressURL/internal/model/entity"
	"compressURL/internal/service"
	"compressURL/utility"
	"errors"
	"fmt"
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

// Create todo 创建完用户后应该返回用户信息
func (s *sUser) Create(ctx context.Context, in entity.User) (string, error) {
	// 昵称不存在生成默认昵称
	if in.NickName == "" {
		in.NickName = "kk" + grand.S(4)
	}

	// 头像不存在生成随机头像
	if in.Avatar == "" {
		in.Avatar = fmt.Sprintf("https://api.dicebear.com/7.x/bottts-neutral/svg?seed=%s&size=64", in.NickName)
	}

	data := g.Map{
		"Id":          in.Id,
		"Email":       nil,
		"Password":    nil,
		"NickName":    in.NickName,
		"AccountType": in.AccountType,
		"Role":        in.Role,
		"Salt":        nil,
		"Avatar":      in.Avatar,
	}

	if data["AccountType"] == "01" {
		// 生成随机盐值
		salt := grand.S(10)

		data["Email"] = in.Email
		data["Id"] = guid.S()
		data["Salt"] = salt
		data["Password"] = utility.EncryptPassword(in.Password, salt)
	}

	_, err := dao.User.Ctx(ctx).Data(data).InsertAndGetId()
	if err != nil {
		return "", err
	}

	return data["Id"].(string), nil
}

// Update 更新用户信息
func (s *sUser) Update(ctx context.Context, in entity.User) error {
	// 获取用户信息
	userInfo, err := service.User().Detail(ctx, model.UserQueryInput{Id: in.Id})

	if err != nil {
		return err
	}

	if in.Password != "" {
		// 对修改后的密码加密
		in.Password = utility.EncryptPassword(in.Password, userInfo.Salt)
	}

	// 账号未更改不做更新
	if in.Email == userInfo.Email {
		in.Email = ""
	}

	// 查询该邮箱是否已经被绑定
	count, err := dao.User.Ctx(ctx).Where(dao.User.Columns().Email, in.Email).Count()
	if err != nil {
		return err
	}

	// 存在一个用户且邮箱不为空
	if count == 1 && in.Email != "" {
		return gerror.New("邮箱已存在！")
	}

	// 更新数据
	_, err = dao.User.Ctx(ctx).Data(in).
		OmitEmpty().
		FieldsEx(dao.User.Columns().Id).
		Where(dao.User.Columns().Id, in.Id).
		Update()
	return err
}

// Delete 删除用户
func (s *sUser) Delete(ctx context.Context, id string) error {
	res, err := dao.User.Ctx(ctx).Where(dao.User.Columns().Id, id).Delete()

	if num, _ := res.RowsAffected(); num == 0 {
		return gerror.New("需要删除的数据不存在！")
	}

	return err
}

// Detail 获取用户详情
func (s *sUser) Detail(ctx context.Context, in model.UserQueryInput) (*entity.User, error) {
	userInfo := entity.User{}

	db := dao.User.Ctx(ctx).
		OmitEmptyWhere().
		Where(dao.User.Columns().Id, in.Id).
		Where(dao.User.Columns().Email, in.Email)

	total, err := db.Where(db).Count()

	if total == 0 {
		g.Log().Debug(ctx, "user GetOne 未查询到用户,邮箱:", in.Email, "id:", in.Id)
		return nil, nil
	}

	err = db.Scan(&userInfo)

	if err != nil {
		g.Log().Error(ctx, "User Detail error:", err)
		return nil, err
	}

	return &userInfo, nil
}

// GetOne 根据 id 获取用户信息,隐藏关键信息
func (s *sUser) GetOne(ctx context.Context, in model.UserQueryInput) (*v1.GetOneRes, error) {
	userInfo := v1.GetOneRes{}

	db := dao.User.Ctx(ctx).
		OmitEmptyWhere().
		Where(dao.User.Columns().Id, in.Id).
		Where(dao.User.Columns().Email, in.Email)

	total, err := db.Where(db).Count()

	if total == 0 {
		g.Log().Debug(ctx, "user GetOne 未查询到用户,邮箱:", in.Email, "id:", in.Id)
		return nil, nil
	}

	err = db.Scan(&userInfo)

	if err != nil {
		g.Log().Error(ctx, "User GetOne error:", err)
		return nil, err
	}

	return &userInfo, nil
}

// GetUserInfoByWxId 根据 wxIdOpenId 获取用户信息 todo WxId 后期会合并到 id 中到时删除
func (s *sUser) GetUserInfoByWxId(ctx context.Context, wxId string) (*v1.GetOneRes, error) {
	userInfo := v1.GetOneRes{}

	err := dao.User.Ctx(ctx).Where(dao.User.Columns().WxId, wxId).Scan(&userInfo)

	if err != nil {
		return nil, errors.New("未查询到用户数据！")
	}
	return &userInfo, nil
}

// GetUserList 获取用户列表
func (s *sUser) GetUserList(ctx context.Context, in v1.GetListReq) ([]entity.User, int, error) {
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
