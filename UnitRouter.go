// 此为CoderG与注册单元内路由功能有关的文件，
// 使用 GNU GPL v3 许可证授权

package coderg

import (
	"reflect"
)

type UnitRouterDynamic struct{
	Url string
	ControllorValue reflect.Value
}

type UnitRouter struct {
	Info []UnitRouterDynamic;
}

type UnitDoorInterface interface {
	 Router()(router *UnitRouter)
}

func NewUnitRouter()(router *UnitRouter){
	info := make([]UnitRouterDynamic,0,10);
	router = &UnitRouter{Info: info};
	return;
}

func (r *UnitRouter) Add(url string, c ControllorInterface ) (){
	cv := reflect.ValueOf(c);
	url = PathMustBegin(url);
	rd := UnitRouterDynamic{ Url : url , ControllorValue: cv};
	r.Info = append(r.Info,rd);
	return;
}
