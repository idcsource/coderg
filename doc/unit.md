# 单元是个好东西 #

“单元”是CoderG中核心理念的重要组成部分。有了这个概念，你就可以很方便得对功能进行集合，也可以相对方便得做到功能的复用。

“单元”实际上就是实现某一功能的一系列Controllor（控制器）的组合。

举个例子，你的Web站点需要一个新闻类的模块，它需要能管理、发布和展现新闻，而且是带有二级或三级分类的。这个新闻模块肯定需要许多页面和文件，如果你的Web站点功能比较多，那么把类似新闻模块这样的功能在代码层面上划分成独立的“单元”，显然在管理与维护代码方面会有很多好处。

然后，上面的例子再继续下去。如果你的Web站点依然需要一个管理下载文件的模块，比如叫什么“招标文件下载”。也许你发现这个模块的除了多了一个文件上传和下载的功能外，其他功能都与前面的新闻类模块一样。那么你只需要写一个提供下载的“单元”就可以了。无论是“招标文件下载”还是“企业新闻”都可以直接用这一个家伙就可以了。

当然，我知道你们大部分人现在就是在用类似的方式，我只不过把这个方式叫做“单元”而已。当然，实际上不是这么简单。

在codergweb里面，“单元”在名叫unit的路径下，很直接。在unit路径下有一个askanswer文件夹，它提供了“有问不一定必答”的功能，这里面所有的文件都遵循CoderG中关于Controllor的要求，这个在之后的文档里再详述，这里还是首先看一下unit路径。

在unit路径下，除了askanswer这个文件夹外，还有名为AskAnswerDoor.go的文件。为了能使用CoderG提供的所有功能，为“单元”准备这样一个“门口（Door）”文件是必需的，这个文件告诉“节点”（在之后的文档里详述）这个“单元”内的Controllor是如何路由的。

“门口”需要满足UnitDoorInterface这个接口，此接口在coderg中的UnitRouter.go里定义。下面贴出某一个版本的AskAnswerDoor.go文件（可能和你拿到的代码不一样）并加以注释讲解。

	package unit  //这个文件输入unit包（codergweb/unit）

	import (
		"coderg"    //载入CoderG核心包
		"codergweb/unit/askanswer"    //载入askanswer单元
	)
	
	type AskAnswerDoor struct {
			
	}
	func (s *AskAnswerDoor) Router()(route *coderg.UnitRouter){
		route = coderg.NewUnitRouter();  //新建一个单元路由
		//往下四行均为添加具体的单元路由信息
		route.Add("admin-ad",&askanswer.AdminAnswerDelControllor{});
		route.Add("admin",&askanswer.AdminControllor{});
		route.Add("add",&askanswer.AddControllor{});
		route.Add("/",&askanswer.MainControllor{});
		return;
	}

通过这个文件，“节点”将知道askanswer内各个控制器是如何进行路由的，并最终在运行的时候生成类似如下结构的路由表(其中的AskAnswer是节点的名称)：

	/AskAnswer/admin-add	-->  &askanswer.AdminAnswerDelControllor{}
	/AskAnswer/admin		-->  &askanswer.AdminControllor{}
	/AskAnswer/add			-->  &askanswer.AddControllor{}
	/AskAnswer/				-->  &askanswer.MainControllor{}

“单元”这个东西大改就是这么回事了，“节点”和“控制器”以及其他未解之谜都将在后面的文档中涉及。请继续。
