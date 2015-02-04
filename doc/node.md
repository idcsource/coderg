# 节点是神马玩意 #

“节点”（Node）是一个非常非常好玩的玩意儿，有了它，那复杂的“单元”以及“门口”之类的东西才有用处。而且“节点”所生成的路由表和CoderG提供的普通的路由表（之后的文档详述）功能是可以同时存在的，所以你可以混着用，只要保证所设定的节点标识与路由路径不重复就可以了。

“节点”以及派生出来的“节点树”（NodeTree，不需要解释这个概念了吧）实际上不过是一种组织Web站点路径结构的工具，是对普通路由的增强。形象点就是你造一棵圣诞树，首先往树干上安装树枝，然后往树枝上挂各种礼物和装饰品，而这些树枝上还可以继续分叉，然后挂更多的礼物和装饰品。

“节点”功能的主要代码在coderg的NodeTree.go文件中，但生成路由表的时候则会用到coderg中Router.go文件里的部分结构体(struct)。

通常来说，“节点”和“节点树”配置和生成都是放在站点的main.go中，下面是codergweb中某一个版本main()函数里涉及到“节点”的代码（可能和你拿到的代码不一样）：

	service.NodeTree = coderg.NewNodeTree(&controllor.IndexControllor{});
		// ↑ 新建一个节点树，并且指定Web站点首页所对应的Controllor
	service.NodeTree.AddNode("WhyIp","",&controllor.WhyIpControllor{});
		// ↑ 添加一个节点，节点标识为“WhyIp”，节点配置文件为空，最后是对应的Controllor
	service.NodeTree.AddNode("Docs","nodetree/docs.cfg",&controllor.DocsControllor{});
		// ↑ 添加一个叫“Docs”的节点，与上面唯一的不同是定义了这个节点的配置文件
	service.NodeTree.AddUnit("AskAnswer","nodetree/askanswer.cfg",&unit.AskAnswerDoor{});
		// ↑ 添加“AskAnswer”节点，但这个节点是一个“单元”，于是最后是对应的“入口”文件
	service.NodeTree.Router(service);
		// ↑ 生成路由表

然后它将生成一个类似如下结构的路由表（含askanswer单元内路由）：

	/WhyIp/					-->  &controllor.WhyIpControllor{}
	/Docs/					-->  &controllor.DocsControllor{}
	/AskAnswer/admin-add	-->  &askanswer.AdminAnswerDelControllor{}
	/AskAnswer/admin		-->  &askanswer.AdminControllor{}
	/AskAnswer/add			-->  &askanswer.AddControllor{}
	/AskAnswer/				-->  &askanswer.MainControllor{}
	/	  					-->  &controllor.IndexControllor{}

这里来说一下节点配置文件。每个“节点”都可以拥有自己的配置文件（配置文件格式见代码，这里不废话），这是一个很有用的东西。比如在“单元”那篇文档中我们所举“招标文件下载”和“企业新闻”共用一个“单元”的例子，我们完全可以在“企业新闻”这个节点的配置文件上加上关闭下载功能或其他的设置项，而你在写这个“单元”的具体代码时只要读取配置内容（如何读取配置在“控制器”中详述）就可以做到将自身轻易的从“下载单元”变身成为“新闻单元”（在“控制器”里你可以获得如何让“单元”知道自己是被挂在哪个“节点”的方法）。

因为“节点”这个家伙造成大量的“单元”重用，除了如上面那段所说在“单元”编写的时候需要可以根据配置文件判断自己在哪个节点以及自我变身外，如果你的“单元”涉及到数据库的读写，那么你可能在设计表的时候多加上一个字段来标注“节点”。当然我知道你们其实一直都在这么做。
