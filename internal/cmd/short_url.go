package cmd

import (
	"compressURL/internal/model/entity"
	"compressURL/internal/service"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"io"
	"net/http"
)

// IPInfo 结构体用于存储IP地址的详细信息
type IPInfo struct {
	Query         string  `json:"query"`         // 查询的IP地址
	Status        string  `json:"status"`        // 查询状态，成功或失败
	Continent     string  `json:"continent"`     // 大洲名称
	ContinentCode string  `json:"continentCode"` // 大洲代码
	Country       string  `json:"country"`       // 国家名称
	CountryCode   string  `json:"countryCode"`   // 国家代码
	Region        string  `json:"region"`        // 地区或州的短代码（FIPS或ISO）
	RegionName    string  `json:"regionName"`    // 地区或州名称
	City          string  `json:"city"`          // 城市名称
	District      string  `json:"district"`      // 位置的区
	Lat           float64 `json:"lat"`           // 纬度
	Lon           float64 `json:"lon"`           // 经度
}

func getIpInfo(ctx context.Context, clientIp string) (IPInfo, error) {
	url := fmt.Sprintf("http://ip-api.com/json/%s?fields=status,message,continent,continentCode,country,countryCode,region,regionName,city,district,lat,lon,query&lang=zh-CN", clientIp)
	res, err := g.Client().Get(ctx, url)

	if err != nil {
		return IPInfo{}, fmt.Errorf("failed to make request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return IPInfo{}, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return IPInfo{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var ipInfo IPInfo
	if err := json.Unmarshal(body, &ipInfo); err != nil {
		return IPInfo{}, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	if ipInfo.Status != "success" {
		return IPInfo{}, fmt.Errorf("IP info retrieval failed: %s", ipInfo.Status)
	}

	return ipInfo, nil
}

func saveVisitsInfo(ctx context.Context, r *ghttp.Request, shortUrlInfo entity.ShortUrl) {
	clientIp := r.GetClientIp()

	ipInfo, err := getIpInfo(ctx, clientIp)

	if err != nil {
		g.Log().Error(ctx, "获取 ip 信息失败:", err)
	}

	err = service.ShortUrlVisits().Create(ctx, entity.ShortUrlVisits{
		Id:              0,
		UserId:          shortUrlInfo.UserId,
		ShortUrl:        shortUrlInfo.ShortUrl,
		RawUrl:          shortUrlInfo.RawUrl,
		Ip:              clientIp,
		UserAgent:       r.Header.Get("User-Agent"),
		SecChUa:         r.Header.Get("Sec-Ch-Ua"),
		SecChUaMobile:   r.Header.Get("Sec-Ch-Ua-Mobile"),
		SecChUaPlatform: r.Header.Get("Sec-Ch-Ua-Platform"),
		SecFetchUser:    r.Header.Get("Sec-Fetch-User"),
		Continent:       ipInfo.Continent,
		ContinentCode:   ipInfo.ContinentCode,
		Country:         ipInfo.Country,
		CountryCode:     ipInfo.CountryCode,
		Region:          ipInfo.Region,
		RegionName:      ipInfo.RegionName,
		City:            ipInfo.City,
		District:        ipInfo.District,
		Lat:             ipInfo.Lat,
		Lon:             ipInfo.Lon,
	})

	if err != nil {
		g.Log().Error(ctx, "创建短链访问信息失败:", err)
	}
}

// 注册短链服务
func registerShortURLService(s *ghttp.Server, ctx context.Context) {
	s.BindHandler("/:id", func(r *ghttp.Request) {
		url := gstr.SubStr(r.Request.RequestURI, 1, len(r.Request.RequestURI))

		if url == "favicon.ico" {
			return
		}

		req, err := service.ShortUrl().GenOne(ctx, url)

		g.Log().Debug(ctx, url, err)

		if err != nil {
			r.Response.Write("未获取到对应的地址，请检查链接是否正确！")
			return
		}

		go saveVisitsInfo(ctx, r, req)

		http.Redirect(r.Response.ResponseWriter, r.Request, req.RawUrl, http.StatusFound)
	})
}
