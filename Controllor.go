// 此为CoderG的对控制器的定义的文件
// 使用 GNU GPL v3 许可证授权

package coderg

import(
	"net/http"
	"strings"
)

// ControllorInterface 此为控制器接口的定义
type ControllorInterface interface {
	 Init(w http.ResponseWriter, r *http.Request, service *ServiceStruct, rt Runtime)
	 Exec()
}

type Controllor struct {
	Runtime Runtime
	Service *ServiceStruct
	W http.ResponseWriter
	R *http.Request
}

// Init 函数负责对控制器进行初始化
// 根据Runtime中的NowRoutePath，使用strings.Split函数根据“/”分割并作为Runtime中的UrlRequest（已清除可能的为空的字符串）
func (c *Controllor) Init (w http.ResponseWriter, r *http.Request, service *ServiceStruct, rt Runtime) {
	c.W = w;
	c.R = r;
	c.Runtime = rt;
	c.Service = service;
	urlRequest := strings.Split(c.Runtime.NowRoutePath, "/");
	for _, v := range urlRequest{
		if ( len(v) != 0 ) {
			c.Runtime.UrlRequest = append(c.Runtime.UrlRequest, v);
		}
	}
	for _, conf := range c.Runtime.ConfigTree {
		if conf != nil {
			c.Runtime.MyConfig = conf;
			break;
		}
	}
	if c.Runtime.MyConfig == nil {
		c.Runtime.MyConfig = c.Service.WebConfig;
	}
}

func (c *Controllor) Exec() {
	
}
