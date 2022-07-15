package collect

import (
	route "github.com/glide-im/api/internal/api/router"
	"github.com/mozillazg/request"
	"net/http"
)

func GetIpAddr(ip string) string {
	c := new(http.Client)
	req := request.NewRequest(c)
	resp, err := req.Get("https://ip.zxinc.org/api.php?type=json&ip=" + ip)
	if err != nil {
		return ""
	}
	j, err := resp.Json()
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	j, err = resp.Json()
	if err != nil {
		return ""
	}

	location, err := j.Get("data").Get("country").String()

	if err != nil {
		return ""
	}
	return location
}

func GetIp(ctx *route.Context) string {
	ip := ctx.Context.ClientIP()
	return ip
}

func GetBrowser(ctx *route.Context) string {
	ua := ctx.Context.GetHeader("User-Agent")
	return ua
}

func GetUserUa(ctx *route.Context) CollectData {
	ip := GetIp(ctx)
	collectData := CollectData{
		Ip:      ip,
		Region:  GetIpAddr(GetIp(ctx)),
		Browser: GetBrowser(ctx),
	}
	return collectData
}
