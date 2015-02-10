// 此为CoderG的Service文件
// 使用 GNU GPL v3 许可证授权

package coderg

import(
	"path/filepath"
	"os"
	"net/http"
	"fmt"
	"sync"
	"database/sql"
	
	"github.com/msbranco/goconfig"
)

// ServiceGo 为唯一的全局变量，为了协调各服务
var ServiceGo sync.WaitGroup

type ServiceStruct struct {
	GlobalRelativePath string  //本地路径
	Grp string //GlobalRelativePath
	WebConfig *goconfig.ConfigFile  //配置文件
	MainRouter *RouterPath  //总路由
	//TemplateConfig TemplateConfig  //模板配置缓存
	NodeTree *NodeTree //节点树
	Template TemplateCache  //模板准备缓存，这个没有自动的，只有手动的准备和注册
	Database *sql.DB  //数据库连接
}

// ServicePrepare 此函数为“服务”的准备，需要一个配置文件，并返回ServiceStruct
func ServicePrepare(configfile string) (Service *ServiceStruct){
	Service = &ServiceStruct{};
	Service.GlobalRelativePath = DirMustEnd(filepath.Dir(os.Args[0]));
	Service.Grp = Service.GlobalRelativePath;
	webconfig, err := goconfig.ReadConfigFile(Service.GlobalRelativePath + configfile);
	if(err != nil){
		errs :=  CodergError(1, err.Error());
		fmt.Fprintln(os.Stderr, errs);
		os.Exit(1);
	}
	Service.WebConfig = webconfig;
	Service.MainRouter = NewRouterPath(Service);
	
	Service.Template = make(TemplateCache);
	/*
	Service.TemplateConfig = make(TemplateConfig);
	if tc, _ := Service.WebConfig.GetString("templateconfig","cache"); tc == "yes" {
		tpath, _ := Service.WebConfig.GetString("templateconfig","path");
		Service.TemplateConfig.AllCache(tpath);
	}
	*/
	ifdb, err := webconfig.GetBool("server","database");
	if err != nil {
		errs :=  CodergError(2, "");
		fmt.Fprintln(os.Stderr, errs);
		os.Exit(1);
	}
	if ifdb == true {
		Service.Database = DatabasePrepare(Service.WebConfig);
	}
	return;
}

func (s *ServiceStruct) serviceGo(){
	var ifHttps bool;
	ifHttps, e1 := s.WebConfig.GetBool("server", "https");
	if e1 != nil {
		ifHttps = false;
	}
	var thecert, thekey string;
	if ifHttps == true {
		var e2, e3 error;
		thecert, e2 = s.WebConfig.GetString("server","sslcert");
		thekey, e3 = s.WebConfig.GetString("server","sslkey");
		if e2 != nil || e3 != nil {
			fmt.Fprintln(os.Stderr,"找不到SSL所需要的文件");
			os.Exit(1);
		}
	}
	var thePort string;
	thePort, e4 := s.WebConfig.GetString("server","port");
	if e4 != nil {
		if ifHttps == false {
			thePort = "80"
		}else{
			thePort = "443";
		}
	}
	thePort = ":" + thePort;
	
	var err error;
	if ifHttps == true {
		err = http.ListenAndServeTLS(thePort, s.GlobalRelativePath + thecert, s.GlobalRelativePath + thekey, s.MainRouter);
	}else{
		err = http.ListenAndServe(thePort, s.MainRouter);
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "服务器动出错：", err);
		os.Exit(1)
	}
}

func (s *ServiceStruct) Start(){
	ServiceGo.Add(1);
	go s.serviceGo();
}

func (s *ServiceStruct) Wait(){
	ServiceGo.Wait();
}
