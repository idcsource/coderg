# 运行时数据——Runtime #

和之前所说的ServiceStruct不同，Runtime是在执行你的控制器的时候生成的数据，这些数据可能是在Router中生成，也可能是在控制器的Init()函数中生成，但最终都将装入Runtime中供你的控制器使用。

**AllRoutePath**： 这是你当前页面的URL除去域名后的完整路径，比如你的页面地址是http://domain.name:1234/node1/node2/page1，那么这个值则是“/node1/node2/page1”。

**NowRoutePath**： 经过路由之后，AllRoutePath被路由利用了一部分来到达你的页面，而NowRoutePath这是URL还剩余的部分，比如http://domain.name/node/page1/参数1/参数2，那么这个值就是“/参数1/参数2”。

**RealNode**： 当前节点树的名称，如node1/node2，因为在CoderG里一个控制器或单元可以在节点的任何位置，所以让它们知道现在身处哪里就很重要了，如果你没有用节点来管理这个控制器，那么这个值为空，也就是""。

**ConfigTree**： 这是个很有趣的东西，它是map[string]*goconfig.ConfigFile类型，其中的string为节点树名，它记录了从当前节点树到父节点树的一整串配置文件。例如你现在的节点树是/abc/node1/sp，那么在ConfigTree里就是这么记录的：[/abc/node1/sp]config、[/abc/node1]config、[/abc]config。如果某个节点没有配置文件，那么*goconfig.ConfigFile将会是一个nil。

**MyConfig**： 此值是在Controllor的Init()中生成的，其实就是把ConfigTree中的第一个取出来，如果当前的节点没有配置文件，则会一只向上遍历寻找，如果一直没有找到有配置文件的父节点，那么这个值将把ServiceStruct里的WebConfig拿过来用。反正就是保证不为空。

**UrlRequest**： 同样是在Controllor的Init()中生成，其实就是把NowRoutePath拿过来用“/”分割成字符串切片，唯一智能点的地方就是字符串切片中的空字符串将被剔除。如果你认为你的参数应该是int或者float之类的，那你就自己转数据类型呗，因为这里一切都是字符串。
