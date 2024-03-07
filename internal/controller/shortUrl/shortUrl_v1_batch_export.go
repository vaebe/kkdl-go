package shortUrl

import (
	"compressURL/internal/model"
	"compressURL/internal/service"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/xuri/excelize/v2"

	"compressURL/api/shortUrl/v1"
)

func (c *ControllerV1) BatchExport(ctx context.Context, req *v1.BatchExportReq) (res *v1.BatchExportRes, err error) {
	loginUserInfo, err := service.Auth().GetLoginUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	userId := loginUserInfo.Id

	// 非管理员用户导出需要携带用户 id
	if loginUserInfo.Role == "00" {
		userId = ""
		if loginUserInfo.Id == "" {
			return nil, gerror.Newf("用户id不能为空: %s", err)
		}
	}

	inParams := v1.GetListReq{
		PageParams: model.PageParams{
			PageSize: 9999,
			PageNo:   1,
		},
		Title:  req.Title,
		RawUrl: req.RawUrl,
	}

	// 获取数据
	dataList, _, err := service.ShortUrl().GetList(ctx, inParams, userId)
	if err != nil {
		glog.Errorf(ctx, "获取导出数据错误: %s", err)
	}

	// 导出 excel
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			glog.Errorf(ctx, "关闭文件错误: %s", err)
		}
	}()

	// 创建一个工作表
	sheetIndex, err := f.NewSheet("Sheet1")
	if err != nil {
		return nil, gerror.Newf("创建工作表错误: %s", err)
	}

	if err = f.SetColWidth("Sheet1", "A", "H", 20); err != nil {
		return nil, gerror.Newf("设置列宽错误: %s", err)
	}

	// 设置工作簿的默认工作表
	f.SetActiveSheet(sheetIndex)

	// 插入表头
	titleList := []string{"短链名称", "短链", "跳转链接", "创建时间", "过期时间", "短链分组id"}
	for i, v := range titleList {
		cellName, _ := excelize.CoordinatesToCellName(i+1, 1)
		if err := f.SetCellValue("Sheet1", cellName, v); err != nil {
			return nil, gerror.Newf("设置表头错误: %s", err)
		}
	}

	// 插入内容
	for i, v := range dataList {
		row := i + 2 // 从第二行开始插入数据
		cells := []interface{}{v.Title, v.ShortUrl, v.RawUrl, v.CreatedAt, v.ExpirationTime, v.GroupId}
		if err := f.SetSheetRow("Sheet1", fmt.Sprintf("A%d", row), &cells); err != nil {
			return nil, gerror.Newf("设置表数据错误: %s", err)
		}
	}

	fileName := fmt.Sprintf("短链列表%s.xlsx", gtime.Datetime())
	saveFilePath := "resource/template/shortUrl/" + fileName
	if err := f.SaveAs(saveFilePath); err != nil {
		return nil, gerror.Newf("保存导出文件错误: %s", err)
	}

	g.RequestFromCtx(ctx).Response.ServeFileDownload(saveFilePath, fileName)

	// 完成后删除文件
	if err = gfile.Remove(saveFilePath); err != nil {
		glog.Error(ctx, "删除文件失败", err)
	}

	return nil, nil
}
