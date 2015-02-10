package coderg

import(
	"reflect"
	"path/filepath"
	//"fmt"
	"os"
	"regexp"
	//"strings"
	
	"github.com/msbranco/goconfig"
)

type NodeTree struct{
	Mark string  //用来做路由的，也就是未来显示在连接上的地址
	Config *goconfig.ConfigFile //配置文件
	IfChildren bool  //是否有下层
	IfDoor bool   //是否为入口
	ControllorValue reflect.Value  //控制器的反射值
	Children map[string]*NodeTree  //下层的信息，map的键为Mark
}

//新建节点树，并输入跟节点的控制器
func NewNodeTree(c ControllorInterface) (*NodeTree) {
	cv := reflect.ValueOf(c);
	nt := &NodeTree{ Mark : "", IfChildren : false, ControllorValue : cv, Children : make(map[string]*NodeTree) };
	return nt;
}

//添加Children，返回Children的子节点书
func (nt *NodeTree) AddNode (mark, conf string, c ControllorInterface) (*NodeTree, error){
	var wrong error;
	cv := reflect.ValueOf(c);
	nt.IfChildren = true;
	cnt := &NodeTree{ Mark : mark, IfChildren : false, IfDoor : false, ControllorValue : cv, Children : make(map[string]*NodeTree) };
	if len(conf) != 0 {
		globalPath := DirMustEnd(filepath.Dir(os.Args[0]));
		conf = globalPath + conf;
		config, err := goconfig.ReadConfigFile(conf);
		if err != nil {
			wrong = CodergError(6, mark);
		}
		cnt.Config = config;
	}
	nt.Children[mark] = cnt;
	return nt.Children[mark], wrong;
}

func (nt *NodeTree) AddUnit (mark, conf string, d UnitDoorInterface) (*NodeTree, error){
	var wrong error;
	cv := reflect.ValueOf(d);
	nt.IfChildren = true;
	cnt := &NodeTree{ Mark : mark, IfChildren : false, IfDoor : true, ControllorValue : cv, Children : make(map[string]*NodeTree) };
	if len(conf) != 0 {
		globalPath := DirMustEnd(filepath.Dir(os.Args[0]));
		conf = globalPath + conf;
		config, err := goconfig.ReadConfigFile(conf);
		if err != nil {
			wrong = CodergError(6, mark);
		}
		cnt.Config = config;
	}
	nt.Children[mark] = cnt;
	return nt.Children[mark], wrong;
}

//生成路由，直接加入Service里的MainRouter中
func (nt *NodeTree) Router (s *ServiceStruct){
	
	if nt.IfChildren == true {
		conf := make(map[string]*goconfig.ConfigFile);
		nt.rangeRoute(s.MainRouter, "", nt.Children, conf);
	}
	s.MainRouter.IndexRoute = nt.ControllorValue;
	s.MainRouter.IndexRouteIsSet = true;
	
	return;
}

func (nt *NodeTree) rangeRoute (r *RouterPath, path string, children map[string]*NodeTree , conf map[string]*goconfig.ConfigFile){
	for _, v := range children {
		nowpath := path + "/" + v.Mark ;
		nowconf := make(map[string]*goconfig.ConfigFile);
		nowconf[nowpath] = v.Config;
		for fpath, fconf := range conf {
			nowconf[fpath] = fconf;
		}
		if v.IfChildren == true {
			nt.rangeRoute(r, nowpath, v.Children, nowconf);
		}
		if v.IfDoor == true {
			dpa := v.ControllorValue.MethodByName("Router").Call(nil);
			routeA := dpa[0].Interface().(*UnitRouter);
			for _, route := range routeA.Info {
				therege := "";
				lipath := "";
				if route.Url == "/" {
					therege = "^" + nowpath + "/(.*)";
					lipath = nowpath;
				}else{
					lipath = nowpath + route.Url;
					therege = "^" + lipath + "/(.*)" ;
				}
				tu, _ := regexp.Compile(therege);
				rd := RouterDynamic{Regexp: tu, Equal : lipath , ControllorValue: route.ControllorValue, IfConfig : true, Config: nowconf, RealNode: nowpath};
				r.Dynamic = append(r.Dynamic,rd);
			}
		}else{	
			therege := "^" + nowpath + "/(.*)" ;
			eq := nowpath;
			tu, err := regexp.Compile(therege);
			if err != nil {
				continue;
			}
			rd := RouterDynamic{Regexp: tu, Equal : eq , ControllorValue: v.ControllorValue, IfConfig : true, Config: nowconf, RealNode:nowpath};
			r.Dynamic = append(r.Dynamic,rd);
		}
	}
}
