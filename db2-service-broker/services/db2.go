package services

import (
	_ "bitbucket.org/phiggins/db2cli"
	"crypto/md5"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"strings"
	"github.com/astaxie/beego/logs"
)

// DATABASE=sample; HOSTNAME=192.168.21.87; PORT=50000; PROTOCOL=TCPIP; UID=db2admin; PWD=12345678; FILE_ROOT=C:\\Tablespaces\\
func getDB2ConnectionInfo(connUri string) (*string, *string, *string, *string, error) {
	parts := strings.Split(connUri, "; ")
	if len(parts) != 7 {
		return nil, nil, nil, nil, errors.New("connUri is illegal.")
	}
	if !strings.Contains(connUri, "DATABASE=") ||
		!strings.Contains(connUri, "HOSTNAME=") ||
		!strings.Contains(connUri, "PORT=") ||
		!strings.Contains(connUri, "PROTOCOL=") ||
		!strings.Contains(connUri, "UID=") ||
		!strings.Contains(connUri, "PWD=") ||
		!strings.Contains(connUri, "FILE_ROOT=") {
		return nil, nil, nil, nil, errors.New("connUri is illegal.")
	}
	var name string = ""
	var password string = ""
	var fileRoot string = ""
	for _, part := range parts {
		content := strings.Split(part, "=")
		if strings.Contains(part, "UID=") {
			name = content[1]
		}
		if strings.Contains(part, "PWD=") {
			password = content[1]
		}
		if strings.Contains(part, "FILE_ROOT=") {
			fileRoot = content[1]
		}
	}
	index := strings.LastIndex(connUri, "; ")
	var url string = connUri[0:index]

	return &name, &password, &url, &fileRoot, nil
}

// create db2 database and database user.
// return databaseName, userName, userPassword, error.
func createDatabaseAndUser(conn, planValue string) (*string, *string, *string, *string, error) {
	logs.Info("The plan value is ", planValue)
	name, password, url, fileRoot, err := getDB2ConnectionInfo(conn)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	//初始化db2的链接串
	db, err := sql.Open("db2-cli", *url)

	if err != nil {
		return nil, nil, nil, nil, err
	}
	//测试是否能联通
	err = db.Ping()

	if err != nil {
		return nil, nil, nil, nil, err
	}

	defer db.Close()

	size := "-1"
	pageSize := "32k"
	bufferPoolName := fmt.Sprintf("bp%s%s", getguid()[0:10], pageSize)
	tablespaceName := fmt.Sprintf("tp%s", getguid()[0:10])
	fileLocation := fmt.Sprintf("%s%s", *fileRoot, tablespaceName) // "C:\\Tablespaces\\tp1"
	//fileDefaultSize := "5g"
	// create bufferpool tp1bp32k all nodes size -1 pagesize 32k
	// create regular tablespace  tp1 pagesize 32k managed by database using(file 'C:\Tablespaces\tp1' 5g) bufferpool tp1bp32k
	bufferPoolStatement := fmt.Sprintf("create bufferpool %s all nodes size %s pagesize %s;", bufferPoolName, size, pageSize)
	tablespaceStatement := fmt.Sprintf("create regular tablespace %s pagesize %s managed by database using(file '%s' %s) bufferpool %s;",
		tablespaceName, pageSize, fileLocation, planValue, bufferPoolName)
	logs.Info("create bufferPoolStatement: ", bufferPoolStatement)
	logs.Info("create tablespaceStatement: ", tablespaceStatement)
	// 创建 buffer pool
	db.Query(bufferPoolStatement)
	// 创建 tablespace
	db.Query(tablespaceStatement)
	return &bufferPoolName, &tablespaceName, name, password, err
}

// delete oracle database and database user.
// unGrant user to database
// return error.
func deleteDatabaseAndUser(conn, bufferpool, tablespace, username string) error {
	_, _, url, _, err := getDB2ConnectionInfo(conn)
	if err != nil {
		return err
	}
	//初始化mysql的链接串
	db, err := sql.Open("db2-cli", *url)
	if err != nil {
		return err
	}
	//测试是否能联通
	err = db.Ping()

	if err != nil {
		return err
	}

	defer db.Close()

	tablespaceStatement := fmt.Sprintf("drop tablespace %s", tablespace)
	bufferPoolStatement := fmt.Sprintf("drop bufferpool %s", bufferpool)

	logs.Info("drop tablespaceStatement : ", tablespaceStatement)
	logs.Info("drop bufferPoolStatement : ", bufferPoolStatement)

	//删除 tablespace
	db.Query(tablespaceStatement)

	//删除 bufferPool
	db.Query(bufferPoolStatement)
	logs.Info("username: ", username)
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
