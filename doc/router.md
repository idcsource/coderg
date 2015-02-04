# 看起来偷工减料的URL路由 #

CoderG中除了有一个根据节点和单元来进行路由生成功能外，还拥有一个极为简单的基础路由功能。与节点和单元结合起来的路由功能相比，这个基础路由功能简直就是偷工减料。为什么会发生这种偷工减料的问题，原因请转到下一段：

这个基础路由功能是首先被作者实现出来的，经过多次的修修补补，它的确是能用的。当然作者的目标从最初开始就是“节点”，只不过作者很清楚一个简单的单纯的路由功能肯定是基础。但当作者去实现“节点”的自动路由生成的时候，对路由中这个看URL并进行路由的动作进行了极大的扩充，这个基础路由实际上早就不基础了。只不过作者并没有去同时增强添加路由规则这个功能，于是它看起来依然是原始的。当然它的作用依然很大，因为目前你只能通过这个功能添加对404“找不到页面”的处理。

解释完毕。关于基础URL路由的文件是coderg中的Router.go。

使用的时候，因为ServicePrepare已经准备好了路由的结构。所以你只需要调用类似service.MainRouter.AddStatic()之类的就可以使用了。关于Service的东西，依然是以后再说。下面看代码：

	service.MainRouter.AddStatic("/img","static/img");
	service.MainRouter.AddStatic("/css","static/css");
	service.MainRouter.AddStatic("/js","static/js");
	service.MainRouter.AddDynamic("/Docs",&controllor.DocsControllor{});
	service.MainRouter.AddIndexRoute(&controllor.IndexControllor{});
	service.MainRouter.AddNotFind(&controllor.PageNotFindControllor{});

简单一看就可以明白了，添加对静态文件的路由就是AddStatic，添加控制器的路由就用AddDynamic，添加对首页的路由就用AddIndexRoute，添加404路由就用AddNotFind。而其简陋之处猛然看去就能发现它不会和“节点”似的可以指定配置文件。在为了实现节点在路由的时候可以得到必须的Runtime配置参数，路由中的RouterDynamic结构体中的内容被大量增加，可是如果你用AddDynamic的时候这些东西都用不到，最终也不会在Runtime里获取到。关于Runtime的内容，看后面讲运行时的文档。

当然，这个基础URL路由功能是可用的，而且在某些情况下是好用的，所以它会被保留，但不会增强，不过欢迎使用。
