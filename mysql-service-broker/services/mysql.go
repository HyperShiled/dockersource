package services

import (
	"io"
	"database/sql"
	"crypto/rand"
	"crypto/md5"
	_ "github.com/go-sql-driver/mysql"
	"encoding/base64"
	"encoding/hex"
	"github.com/astaxie/beego/logs"
)



// create mysql database and database user.
// grant user to database
// return databaseName, userName, userPassword, error.
func createDatabaseAndUser(conn, planValue string) (*string, *string, *string, error) {
	logs.Info("The plan value is ", planValue)
	//初始化mysql的链接串
	db, err := sql.Open("mysql", conn)

	if err != nil {
		return nil, nil, nil, err
	}
	//测试是否能联通
	err = db.Ping()

	if err != nil {
		return nil, nil, nil, err
	}

	defer db.Close()

	//不能以instancdID为数据库名字，需要创建一个不带-的数据库名
	dbname := "d" + getguid()[0:15]
	_, err = db.Query("CREATE DATABASE " + dbname + " DEFAULT CHARACTER SET utf8 DEFAULT COLLATE utf8_general_ci")

	if err != nil {
		return nil, nil, nil, err
	}

	newusername := getguid()[0:15]
	newpassword := getguid()[0:15]

	_, err = db.Query("GRANT ALL ON " + dbname + ".* TO '" + newusername + "'@'%' IDENTIFIED BY '" + newpassword + "'")

	if err != nil {
		return nil, nil, nil, err
	}

	return &dbname, &newusername, &newpassword, err
}

// delete oracle database and database user.
// unGrant user to database
// return error.
func deleteDatabaseAndUser(conn, database, username string) error {
	//初始化mysql的链接串
	db, err := sql.Open("mysql", conn)
	if err != nil {
		return err
	}
	//测试是否能联通
	err = db.Ping()

	if err != nil {
		return err
	}

	defer db.Close()

	//删除数据库
	_, err = db.Query("DROP DATABASE " + database)

	if err != nil {
		return err
	}

	//删除用户
	_, err = db.Query("DROP USER " + username)

	if err != nil {
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