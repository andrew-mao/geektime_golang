package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type Dao struct {
	Host     string `toml:"host"`
	Port     int64  `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Database string `toml:"database"`
	SslMode  string `toml:"sslmode"`
	TimeZone string `toml:"timezone"`
	Encode   string `toml:"encode"`
	db       *sql.DB
}

//连接数据库
func (d *Dao) Connect() (err error) {
	sqlCon := ` fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", m.Username, m.Password, m.Host, m.Port, m.Database)`
	d.db, err = sql.Open("mysql", sqlCon)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("connect to mysql !!! \n %v", sqlCon))
	}
	return nil
}
func (m *Dao) Close() (err error) {
	err = m.db.Close()
	if err != nil {
		return errors.Wrap(err, "数据库关闭连接失败")
	}
	return nil
}
