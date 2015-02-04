// 此为CoderG的运行时结构体文件
// 使用 GNU GPL v3 许可证授权

package coderg

import(
	"github.com/msbranco/goconfig"
)

// Runtime 运行时数据
// 此运行时数据由CoderG中的Router或Controllor生成
type Runtime struct {
	AllRoutePath		string		//整个的RoutePath，也就是除域名外的完整路径
	NowRoutePath		string		//AllRoutePath经过层级路由之后剩余的部分，通常是提供的Request参数
	RealNode			string		//当前节点的树名，如/node1/node2，如果没有使用节点则此处为空
	ConfigTree 			map[string]*goconfig.ConfigFile //节点的配置文件树，从当前节点开始到最上层父节点，以节点名称为键名
	MyConfig			*goconfig.ConfigFile  //当前节点的配置文件，从ConfigTree中获取，如当前节点没有配置文件，则去寻找父节点，直到载入站点的配置文件
	UrlRequest			[]string	//Url请求的整理，具体行文见Controllor
}
