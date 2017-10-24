package services

import (
	"io"
	"database/sql"
	"crypto/rand"
	"crypto/md5"
	_ "github.com/lib/pq"
	"encoding/base64"
	"encoding/hex"
	"github.com/astaxie/beego/logs"
)

// create postgresql database and database user.
// grant user to database
// return databaseName, userName, userPassword, error.
func createDatabaseAndUser(conn, planValue string) (*string, *string, *string, error) {
	logs.Info("The plan value is ", planValue)

	//初始化postgres的链接串
	db, err := sql.Open("postgres", conn)

	if err != nil {
		return nil, nil, nil, err
	}
	//测试是否能联通
	err = db.Ping()

	if err != nil {
		return nil, nil, nil, err
	}

	defer db.Close()

	//不能以instancdID为数据库名字，需要创建一个不带-的数据库名 pg似乎必须用字母开头的变量
	dbname := "d" + getguid()[0:15]
	newusername := "u" + getguid()[0:15]
	newpassword := "p" + getguid()[0:15]
	_, err = db.Query("CREATE USER " + newusername + " WITH PASSWORD '" + newpassword + "'")

	if err != nil {
		return nil, nil, nil, err
	}
	//_, err = db.Query("CREATE DATABASE " + dbname + " WITH OWNER =" + newusername + " ENCODING = 'UTF8'")
	_, err = db.Query("CREATE DATABASE " + dbname + " ENCODING = 'UTF8'")

	if err != nil {
		return nil, nil, nil, err
	}

	_, err = db.Query("GRANT ALL PRIVILEGES ON DATABASE " + dbname + " TO " + newusername)

	if err != nil {
		return nil, nil, nil, err
	}
	return &dbname, &newusername, &newpassword, err
}

// delete postgresql database and database user.
// unGrant user to database
// return error.
func deleteDatabaseAndUser(conn, database, username string) error {
	//初始化postgres的链接串
	db, err := sql.Open("postgres", conn)

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