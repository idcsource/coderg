# CoderG的服务、准备、启动 #

在CoderG的源码包中，有一个叫Service.go的文件，这里面是关于CoderG“服务(Service)”的相关代码。这里的“服务”只是对CoderG内部一块功能的称谓，与服务器上的服务无关。

“服务”提供了一个（也是CoderG中唯一）的全局变量——ServiceGo，是一个sync.WaitGroup类型。因为CoderG可以同时启动多个服务（不同端口号，go那神奇的go），所以需要有这么一个变量进行协调。简单说无论如何你需要在所有服务启动代码结束之后加上一句ServiceGo.Wait()，当然这个代码也已经封装进“服务”里了。

ServicePrepare函数是对服务的准备动作，需要提供一个配置文件，之后这个函数将返回ServiceStruct结构体类型的对象。这个ServiceStruct对象将包含这个“服务”的几乎所有信息，节点树的信息、路由的信息、数据库连接等的信息均在这里面。因为内容比较多，所以会单独准备一篇文档。

一个要注意的地方。任何被CoderG封装过的代码，类似给ServicePrepare看的配置文件，文件地址都需要直接使用与可执行文件的相对路径（副作用就是你不能把其他文件放到可执行文件的外面去），而类似对模板文件的准备这种没有被CoderG封装的功能，你可以使用“服务”中的GlobalRelativePath变量加到文件的相对路径前，保证文件可以被正确访问。

当你把节点树、路由等设置完毕之后，就可以使用Start()方法启动服务器。而最后，你只需要使用Wait()方法作为main函数的结尾就可以了。

下面给出一个站点main.go文件，此代码是从codergweb项目演示站点的代码修改而来。

	package main

	import(
		"text/template"
		"coderg"
		"codergweb/controllor"
		"codergweb/unit"
	)

	func main(){
	
		service :=coderg.ServicePrepare("codergweb.cfg");
			// ↑ ServicePrepare函数，并提供了配置文件

		service.MainRouter.AddStatic("/img","static/img");
		service.MainRouter.AddStatic("/css","static/css");
		service.MainRouter.AddStatic("/js","static/js");
			// ↑ 添加静态文件路由
	
		/*
		 * 节点树结构
		 * /
		 *  - WhyIp
		 *  - AskAnswer
		 *      - add
		 *  - Docs
		 */
		service.NodeTree = coderg.NewNodeTree(&controllor.IndexControllor{});
		service.NodeTree.AddNode("WhyIp","",&controllor.WhyIpControllor{});
		service.NodeTree.AddNode("Docs","nodetree/docs.cfg",&controllor.DocsControllor{});
		service.NodeTree.AddUnit("AskAnswer","nodetree/askanswer.cfg",&unit.AskAnswerDoor{});
	
		service.NodeTree.Router(service);
			// ↑ 将节点树的路由信息写入路由器
	
		local := service.GlobalRelativePath;
		service.Template["index"], _ = template.ParseFiles(local + "template/index.tmpl",local + "template/mainout.tmpl");
		service.Template["docs"], _ = template.ParseFiles(local + "template/docs.tmpl",local + "template/mainout.tmpl");
		service.Template["whyip"], _ = template.ParseFiles(local + "template/whyip.tmpl",local + "template/mainout.tmpl");
		service.Template["askanswer"], _ = template.ParseFiles(local + "template/askanswer.tmpl",local + "template/mainout.tmpl");
			// ↑ 对模板文件进行缓存
	
		service.Start();
			// ↑ 启动服务

		service2 := coderg.ServicePrepare("otherweb.cfg")
		......
		service2.Start();
			// ↑ 添加另一个服务，使用不同的配置文件，监听不同的端口

		service.Wait();
			// ↑ 也就是执行全局变量的ServiceGo.Wait()
	}

怎么这篇文档显得那么正经？不行，这样不行！
