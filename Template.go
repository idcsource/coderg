// 此为CoderG与模板（以及模板配置文件）缓存有关的文件
// 使用 GNU GPL v3 许可证授权
// 用处不大，可以不用

package coderg

import (
	"os"
	"regexp"
	"fmt"
	"text/template"
	"path/filepath"
	
	"github.com/msbranco/goconfig"
)


type TemplateConfig  map[string]*goconfig.ConfigFile

// 此数据类型在Service中初始化，提供手动在运行前缓存模板的功能
type TemplateCache map[string]*template.Template


// AllCache 这个函数提供了缓存thePath路径下（遍历子路经）所有.cfg文件的功能
// 数据库的保存格式为map[string]*goconfig.ConfigFile
// 其中string为配置文件名不包含.cfg的部分，如果在子路经下则包含子路经，例如：abc/def
func (tc TemplateConfig) AllCache(thePath string) {
	spath := DirMustEnd(filepath.Dir(os.Args[0]));
	thePath = DirMustEnd(thePath);
	thePath = spath + thePath;
	tc.ReadAndAdd(thePath,"");
}

func (tc TemplateConfig) ReadAndAdd(dir, path string){
	opendir, err := os.Open(dir);
	if(err != nil){
		fmt.Fprintln(os.Stderr, "无法读取模板配置文件：", err);
		os.Exit(1);
	}
	defer opendir.Close();
	allfile , _ := opendir.Readdir(0);
	checkfile, _ := regexp.Compile("(.+).cfg$");
	for _, onefile := range allfile {
		if onefile.IsDir() == true {
			newdir := dir + DirMustEnd(onefile.Name());
			newpath := DirMustEnd(onefile.Name()) + path;
			tc.ReadAndAdd(newdir, newpath);
		}else if checkfile.MatchString(onefile.Name()) {
			fileMarks := checkfile.FindStringSubmatch(onefile.Name());
			fileMark := path + fileMarks[1];
			newpath := dir + onefile.Name();
			tc[fileMark], _ = goconfig.ReadConfigFile(newpath);
		}
	}
}
