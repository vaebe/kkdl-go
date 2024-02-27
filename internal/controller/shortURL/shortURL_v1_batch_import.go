package shortURL

import (
	"compressURL/api/shortURL/v1"
	"compressURL/internal/model/entity"
	"compressURL/internal/service"
	"compressURL/utility"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/xuri/excelize/v2"
	"regexp"
)

func checkDataFormat(i int, in []string) string {
	errStr := ""

	title := in[0]
	if title == "" {
		errStr += fmt.Sprintf("第 %d 行短链名称不能为空 %s \t", i, title)
	}

	url := in[1]
	re := regexp.MustCompile(`^https?://`)
	if !re.MatchString(url) {
		errStr += fmt.Sprintf("第 %d 行跳转链接格式错误 %s \t", i, url)
	}

	exTime := in[2]
	if gtime.NewFromStrFormat(exTime, "Y-m-d") == nil && gtime.NewFromStrFormat(exTime, "Y-m-d H:i:s") == nil {
		errStr += fmt.Sprintf("第 %d 行日期格式错误 %s \t", i, exTime)
	}

	return errStr
}

func (c *ControllerV1) BatchImport(ctx context.Context, req *v1.BatchImportReq) (res *v1.BatchImportRes, err error) {
	fileReader, err := req.File.Open()

	if err != nil {
		return nil, gerror.New("打开文件失败！")
	}

	f, err := excelize.OpenReader(fileReader)
	if err != nil {
		return nil, gerror.New("读取文件失败！")
	}

	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return nil, gerror.New("获取所有单元格失败！")
	}

	// 首行数据不存在返回
	if len(rows) == 0 {
		return nil, gerror.New("数据不存在！")
	}

	firstRowExpected := []string{"短链名称", "跳转链接", "过期时间"}
	if !utility.SliceEqual(rows[0], firstRowExpected) {
		return nil, gerror.Newf("首列数据不正确！应为: %s", firstRowExpected)
	}

	// 获取当前登陆用户信息
	loginUserInfo, err := service.Auth().GetLoginUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	errStr := ""
	var data []entity.ShortUrl
	for i, row := range rows {
		if i != 0 {
			errStr += checkDataFormat(i, row) + "\r\n"

			data = append(data, entity.ShortUrl{
				RawUrl:         row[1],
				ExpirationTime: gtime.NewFromStr(row[2]),
				UserId:         loginUserInfo.Id,
				Title:          row[0],
			})
		}
	}

	if errStr != "" {
		return nil, gerror.New(errStr)
	}

	_, err = service.ShortURL().BatchImport(ctx, data)
	return nil, err
}
