package api

import (
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"github.com/iceopen/go-mta/config"
	"github.com/iceopen/go-mta/utils"
)

// 应用每天的pv\uv\vv\iv数据。
func CtrCoreData(start_date string, end_date string, idx string) string {
	m := make(map[string]string, 4)
	m["app_id"] = config.APP_ID
	m["end_date"] = end_date
	m["idx"] = idx
	m["start_date"] = start_date

	sign := utils.SigMake(m)
	param := ""
	for k, v := range m {
		param = param + k + "=" + v + "&"
	}
	param = param + "sign=" + sign
	url := "http://mta.qq.com/h5/api/ctr_core_data?" + param
	logs.Info("ctr_core_data : " + url)
	req := httplib.Get(url)
	str, err := req.String()
	if err != nil {
		logs.Error(503, err)
	}
	logs.Info("ctr_core_data : " + url + " result : " + str)
	return str
}
