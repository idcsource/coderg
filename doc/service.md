# 非常必需但实在混乱的ServiceStruct #

此篇文档将讲述目前为止ServiceStruct里所有的内容，目测将会非常无聊。

**GlobalRelativePath**： 此为本地相对路径，继续以codergweb为例，这个值取决于你在哪里执行codergweb这个程序。如果你在codergweb所在路径执行，这个东西是没有用的，但如果你在别的路径这个值就很有用了，你肯定不打算天天面对找不到文件的。这个值实际上就是filepath.Dir(os.Args[0])获取到的，并且以“/”进行了结尾。被CoderG封装的功能，比如载入配置文件等，CoderG会自己把这个路径加上，所以你不要多此一举。但如果你使用Go的模板引擎这种CoderG根本没有去管的东西，那就是你自己的决定了。

**Grp**： 这是在写此文档的时候刚刚加上的，实际上就是感觉GlobalRelativePath太长了，简写一下呗！

**WebConfig**： 当你在执行ServicePrepare时提供了配置文件的地址，然后配置文件的信息将被解析并放入WebConfig。关于CoderG的配置文件，均依赖了第三方库"github.com/msbranco/goconfig"，这个其他文档上会说明。

**MainRouter**： 这是整站的路由，当你在main()函数里添加路由信息的时候就用到了它，具体看讨论“路由”的文档吧。

**NodeTree**： 很明显，这就是“节点树”。其实“节点树”的信息可以不需要写入Service里，但如果你的站点程序想获取整个节点情况的话，比如生成站点地图之类的，这个还是需要的。

**Template**： 只是一个很简单的map[string]*template.Template类型，用来让你提前准备模板文件的，多占点内存，少点运行时的IO操作而已。

**Database**： 数据库句柄而已，使用Go语言自己的数据库接口。不过因为本人在codergweb上使用的是PostgreSQL，所以CoderG中的数据库连接也只是可以连接PostgreSQL，并使用了第三方"github.com/lib/pq"数据库驱动。这里以后会加上MySQL的支持。

OK，没有了。就这么无聊。
