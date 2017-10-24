package services

import (
	"io"
	"crypto/rand"
	"crypto/md5"
	"gopkg.in/mgo.v2"
	"encoding/base64"
	"encoding/hex"
	"strings"
	"github.com/pkg/errors"
	"github.com/astaxie/beego/logs"
)

func getMongoConnectionInfo(connUri string) (*string, *string, *string, error) {
	// myUserAdmin:myUserAdmin@192.168.21.87:27017
	parts := strings.Split(connUri, "@")
	if len(parts) != 2 {
		return nil, nil, nil, errors.New("connUri is illegal.")
	}
	user := strings.Split(parts[0], ":")
	if len(user) != 2 {
		return nil, nil, nil, errors.New("connUri is illegal.")
	}
	name := user[0]
	password := user[1]
	url := connUri
	return &name, &password, &url, nil
}

// create postgresql database and database user.
// grant user to database
// return databaseName, userName, userPassword, error.
func createDatabaseAndUser(url string, username string, password string, instanceId string, planValue string) (*string, *string, *string, error) {
	logs.Info("The plan value is ", planValue)

	//初始化mongodb的链接串
	session, err := mgo.Dial(url) //连接数据库
	if err != nil {
		return nil, nil, nil, err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	mongodb := session.DB("admin") //数据库名称
	err = mongodb.Login(username, password)
	if err != nil {
		return nil, nil, nil, err
	}

	//创建一个名为instanceId的数据库，并随机的创建用户名和密码，这个用户名是该数据库的管理员
	newdb := session.DB(instanceId)
	newusername := getguid()
	newpassword := getguid()

	//这个服务很快，所以通过同步模式直接返回了
	err = newdb.UpsertUser(&mgo.User{
		Username: newusername,
		Password: newpassword,
		Roles: []mgo.Role{
			mgo.Role(mgo.RoleDBAdmin),
		},
	})

	if err != nil {
		return nil, nil, nil, err
	}

	return &instanceId, &newusername, &newpassword, err
}

// delete mongodb database and database user.
// unGrant user to database
// return error.
func deleteDatabaseAndUser(url, adminUser, adminPassword, databaseName, userName string) error {
	session, err := mgo.Dial(url) //连接数据库
	if err != nil {
		logs.Error("dial database: ", err)
		return err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	mongodb := session.DB("admin") //数据库名称
	err = mongodb.Login(adminUser, adminPassword)
	if err != nil {
		logs.Error("login database: ", err)
		return err
	}

	//选择服务创建的数据库
	userdb := session.DB(databaseName)
	//这个服务很快，所以通过同步模式直接返回了
	err = userdb.DropDatabase()
	if err != nil {
		logs.Error("drop database : ", err)
		return err
	}
	err = userdb.RemoveUser(userName)
	if err != nil {
		logs.Error("remove user : ", err)
		return err
	}
	return nil
}

func getguid() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return getmd5string(base64.URLEncoding.EncodeToString(b))
}

func getmd5string(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}