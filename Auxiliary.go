// 此为CoderG的辅助函数文件
// 使用 GNU GPL v3 许可证授权

package coderg

import (
	"regexp"
	"crypto/sha1"
	"io"
	"fmt"
)

// DirMustEnd 路径必须以斜线结尾
func DirMustEnd(dir string) (string) {
	matched , _ := regexp.MatchString("/$", dir)
	if matched == false {
		dir = dir + "/"
	}
	return dir
}

// PathMustBegin 路径必须以斜线开始
func PathMustBegin(path string) (string){
	matched , _ := regexp.MatchString("^/", path)
	if matched == false {
		path = "/" + path
	}
	return path
}

func GetSha1(data string) string {
    t := sha1.New();
    io.WriteString(t,data);
    return fmt.Sprintf("%x",t.Sum(nil));
}
