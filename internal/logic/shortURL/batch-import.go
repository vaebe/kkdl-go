package shortURL

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/xuri/excelize/v2"
	"golang.org/x/net/context"
)

func (s *sShortURL) BatchImport(ctx context.Context, file *ghttp.UploadFile) error {
	filesReader, err := file.Open()

	if err != nil {
		return err
	}

	f, err := excelize.OpenReader(filesReader)
	if err != nil {

		return err
	}

	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return err
	}

	// 首行数据不存在返回
	if len(rows) == 0 {
		return err
	}

	firstRows := rows[0]
	glog.Info(ctx, firstRows)

	// todo 增加合法性校验、数据插入逻辑
	//for _, row := range rows {
	//	for _, colCell := range row {
	//		fmt.Print(colCell, "\t")
	//	}
	//	fmt.Println()
	//}

	return nil
}
