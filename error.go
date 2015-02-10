package coderg

import(
	"fmt"
)

var codergErrorType = map[int]string{
	0: "一般性错误：",
	1: "无法读取站点配置文件：",
	2: "站点配置文件不完整。",
	3: "站点配置文件中的数据库连接配置不完整。",
	4: "连接数据库出错：",
	5: "未知的数据库类型：",
	6: "无法读取节点配置文件，节点：",
};


type CodergErrorStruct struct {
	err string
}

func CodergError (etype int, content string) (CodergErrorStruct) {
	return CodergErrorStruct{
		err: codergErrorType[etype] + content,
	}
}


func (e CodergErrorStruct) Error() string {
	return fmt.Sprint(e.err);
}
