# 延续无聊的控制器 #

控制器这种东西，本来就是无聊的，无论你是叫它Controller还是Controllor，都不能改变这一点。当然我就叫它Controllor了。关于控制器的定义在Controllor.go文件中。

所有控制器，包括你要加入Router中的或者通过UnitDoor加入到NodeTree中的，都需要满足ControllorInterface接口，具体的接口定义为：

	type ControllorInterface interface {
		 Init(w http.ResponseWriter, r *http.Request, service *ServiceStruct, rt Runtime)
		 Exec()
	}

CoderG实现了一个自己的Controllor结构体和类，主要目的是实现接口中的Init()方法，对控制器进行初始化。

Controllor结构体如下：

	type Controllor struct {
		Runtime Runtime			// 运行时数据，见runtime文档
		Service *ServiceStruct	// Service数据，看service文档
		W http.ResponseWriter	// 这个就不用讲了吧
		R *http.Request			// 这个也不用讲了吧
	}

很简单，很无聊，到此为止。
