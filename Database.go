// 此为CoderG中与数据库连接初始化有关的文件
// 使用 GNU GPL v3 许可证授权
// 依赖第三方库：
// 		github.com/msbranco/goconfig
//		github.com/lib/pq

package coderg

import(
	"database/sql"
	"fmt"
	"os"
	
	"github.com/msbranco/goconfig"
	_ "github.com/lib/pq"
)


// DatabasePrepare 数据库准备函数
// 此函数根据传入配置文件中关于数据库连接的信息进行配置，使用Golang自己的数据库接口，返回*sql.DB类型的数据库连接
// 此函数目前只处理针对PostgreSQL数据库的连接
// 具体查看配置文件终的配置参数为：
// [database]
// 		server		# 数据库服务器的地址
// 		port		# 数据库服务器的端口号
// 		user		# 数据库访问用户名
//		passwd		# 数据库访问用户的密码
//		dbname		# 数据库名
func DatabasePrepare(conf *goconfig.ConfigFile) (db *sql.DB) {
	db_type , e := conf.GetString("database","type");
	if( e != nil ){
		errs :=  CodergError(3, "");
		fmt.Fprintln(os.Stderr, errs);
		os.Exit(1);
	}
	switch db_type {
		case "postgres":
			db = connPostgres(conf);
		default:
			errs := CodergError(5,db_type);
			fmt.Fprintln(os.Stderr, errs);
			os.Exit(1);
	}
	return db;
}

func connPostgres(conf *goconfig.ConfigFile) (*sql.DB) {
	db_server, e1 := conf.GetString("database","server");
	db_port, e2 := conf.GetString("database","port");
	db_user, e3 := conf.GetString("database","user");
	db_passwd, e4 := conf.GetString("database","passwd");
	db_dbname, e5 := conf.GetString("database","dbname");
	if( e1 != nil || e2 != nil || e3 != nil || e4 != nil || e5 != nil ){
		errs :=  CodergError(3, "");
		fmt.Fprintln(os.Stderr, errs);
		os.Exit(1);
	}
	connection_string := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=disable", db_dbname, db_user, db_passwd, db_server, db_port);
	db, err := sql.Open("postgres", connection_string);
	if err != nil {
		errs :=  CodergError(4, err.Error());
		fmt.Fprintln(os.Stderr, errs);
		os.Exit(1);
	}
	return db
}
