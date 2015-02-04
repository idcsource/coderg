/*
 * RouterPath
 * 路径路由器
 */

package coderg

import (
	"regexp"
	"reflect"
	"net/http"
	"github.com/msbranco/goconfig"
	"fmt"
)

// 动态路由信息
type RouterDynamic struct{
	Regexp *regexp.Regexp
	Equal string
	ControllorValue reflect.Value
	IfConfig bool
	Config map[string]*goconfig.ConfigFile
	RealNode string
}


// 总路由信息
type RouterPath struct {
	IndexRoute reflect.Value
	IndexRouteIsSet bool
	Static map[string]*regexp.Regexp
	Dynamic []RouterDynamic
	NotFind reflect.Value
	NotFindIsSet bool
	Service *ServiceStruct
}


// 新建路由
func NewRouterPath( service *ServiceStruct ) (*RouterPath){
	rp := &RouterPath{ IndexRouteIsSet : false, Static : make(map[string]*regexp.Regexp), Dynamic : make([]RouterDynamic,0,10), NotFindIsSet : false, Service : service};
	return rp;
}

// 新建路由
func (r *RouterPath) Init( service *ServiceStruct ){
	r = &RouterPath{ IndexRouteIsSet : false, Static : make(map[string]*regexp.Regexp), Dynamic : make([]RouterDynamic,0,10), NotFindIsSet : false, Service : service};
}

//添加静态路由
func (r *RouterPath) AddStatic(url, path string) (err error){
	url = PathMustBegin(url);
	url = "^" + url + "/(.*)";
	tu, err := regexp.Compile(url);
	if err != nil {
		return;
	}
	r.Static[path] = tu;
	return;
}

//添加动态路由
func (r *RouterPath) AddDynamic(url string, c ControllorInterface ){
	url = PathMustBegin(url);
	reUrl := "^" + url + "/(.*)";
	tu, _ := regexp.Compile(reUrl);
	eq := url;
	cv := reflect.ValueOf(c);
	rd := RouterDynamic{Regexp: tu, Equal : eq , ControllorValue: cv, IfConfig : false};
	r.Dynamic = append(r.Dynamic,rd);
	return;
}

//添加IndexRoute
func (r *RouterPath) AddIndexRoute(c ControllorInterface ){
	cv := reflect.ValueOf(c);
	r.IndexRoute = cv;
	r.IndexRouteIsSet = true;
	return;
}

//添加404路由
func (r *RouterPath) AddNotFind(c ControllorInterface ){
	cv := reflect.ValueOf(c);
	r.NotFind = cv;
	r.NotFindIsSet = true;
	return;
}

//路由
func (r *RouterPath) Router(httpw http.ResponseWriter, httpr *http.Request, rt Runtime, s *ServiceStruct){
	//先进行静态路由的判断
	if( len(r.Static) != 0 ){
		for k, v := range r.Static {
			if v.MatchString(rt.NowRoutePath) {
				nameA := v.FindStringSubmatch(rt.NowRoutePath);
				if len(nameA) > 1 {
					name := nameA[1];
					file := r.Service.GlobalRelativePath + DirMustEnd(k) + name;
					http.ServeFile(httpw, httpr, file);
				}
				return;
			}
		}
	}
	//再开始动态路由判断
	if ( len(r.Dynamic) != 0 ){
		for _, v := range r.Dynamic {
			if (v.Regexp.MatchString(rt.NowRoutePath) || rt.NowRoutePath == v.Equal || rt.NowRoutePath == v.Equal + "/") {
				nowPathA := v.Regexp.FindStringSubmatch(rt.NowRoutePath);
				nowP := "";
				if len(nowPathA) > 1 {
					nowP = nowPathA[1];
				}
				conftree := make(map[string]*goconfig.ConfigFile);
				if v.IfConfig == true {
					conftree = v.Config
				}
				realnode := "";
				if len(v.RealNode) != 0 {
					realnode = v.RealNode;
				}
				nr := Runtime{ AllRoutePath: httpr.URL.Path, NowRoutePath : nowP , ConfigTree: conftree, RealNode: realnode};
				in := make([]reflect.Value, 4);
				in[0] = reflect.ValueOf(httpw);
				in[1] = reflect.ValueOf(httpr);
				in[2] = reflect.ValueOf(s);
				in[3] = reflect.ValueOf(nr);
				v.ControllorValue.MethodByName("Init").Call(in);
				v.ControllorValue.MethodByName("Exec").Call(nil);
				return;
			}
		}
	}
	//看看IndexRoute
	if (r.IndexRouteIsSet){
		if ( httpr.URL.Path == "/" || httpr.URL.Path == "" ){
			nr := Runtime{ AllRoutePath: httpr.URL.Path, NowRoutePath : httpr.URL.Path };
			in := make([]reflect.Value, 4);
			in[0] = reflect.ValueOf(httpw);
			in[1] = reflect.ValueOf(httpr);
			in[2] = reflect.ValueOf(s);
			in[3] = reflect.ValueOf(nr);
			r.IndexRoute.MethodByName("Init").Call(in);
			r.IndexRoute.MethodByName("Exec").Call(nil);
			return;
		}
	}
	//还没有的话就交给404
	if (r.NotFindIsSet) {
		nr := Runtime{ AllRoutePath: httpr.URL.Path, NowRoutePath : "" };
		in := make([]reflect.Value, 4);
		in[0] = reflect.ValueOf(httpw);
		in[1] = reflect.ValueOf(httpr);
		in[2] = reflect.ValueOf(s);
		in[3] = reflect.ValueOf(nr);
		r.NotFind.MethodByName("Init").Call(in);
		r.NotFind.MethodByName("Exec").Call(nil);
		return;
	}else{
		httpw.WriteHeader(404);
		fmt.Fprint(httpw,"404 Page Not Found");
		return;
	}
	
}

//HTTP的路由
func (r *RouterPath) ServeHTTP(httpw http.ResponseWriter, httpr *http.Request) {
	rt := Runtime{ AllRoutePath: httpr.URL.Path, NowRoutePath : httpr.URL.Path };
	r.Router(httpw, httpr,rt, r.Service);
	return;
}
