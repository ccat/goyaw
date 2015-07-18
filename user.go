package goyaw

import (
	"encoding/hex"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/naoina/genmai"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type User struct {
	Id             int64  `db:"pk"`     // column:"tbl_id"
	UserName       string `db:"unique"` // size:"255"
	HashedPassword string
	Active         bool
	CreatedAt      *time.Time
}

type UserMgmt struct {
	UserDB *genmai.DB
}

type UserDBconfig struct {
	Type   string
	Config string
}

func NewUserDB(userDBconfig *UserDBconfig) *UserMgmt {
	var dbIns *UserMgmt = new(UserMgmt)
	var err error

	if userDBconfig.Type == "sqlite3" {
		dbIns.UserDB, err = genmai.New(&genmai.SQLite3Dialect{}, userDBconfig.Config)
	} else if userDBconfig.Type == "mysql" {
		dbIns.UserDB, err = genmai.New(&genmai.MySQLDialect{}, userDBconfig.Config)
	} else if userDBconfig.Type == "postgresql" {
		dbIns.UserDB, err = genmai.New(&genmai.PostgresDialect{}, userDBconfig.Config)
	}
	if err != nil {
		panic(err)
	}

	err = dbIns.UserDB.CreateTableIfNotExists(&User{})
	if err != nil {
		panic(err)
	}
	return dbIns
}

func (self *UserMgmt) CreateUser(username string, password string) error {
	converted, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return errors.New("Invalid Password")
	}
	hashedPass := hex.EncodeToString(converted[:])
	t := time.Now()
	records := []User{
		{UserName: username, HashedPassword: hashedPass, Active: true, CreatedAt: &t},
	}

	_, err = self.UserDB.Insert(records)

	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE") {
			return errors.New("Username already used")
		} else {
			return errors.New("Invalid Password")
		}

	}
	return nil
}

func (self *UserMgmt) Auth(username string, password string) error {

	var results []User
	err := self.UserDB.Select(&results, self.UserDB.Where("user_name", "=", username))
	//err := self.UserDB.Select(&results)
	if err != nil {
		panic(err)
	}
	if len(results) == 0 {
		return errors.New("Invalid username")
	}
	if results[0].Active == false {
		return errors.New("User is not active")

	}

	hashedPass, err := hex.DecodeString(results[0].HashedPassword)
	if err != nil {
		panic(err)
	}

	err = bcrypt.CompareHashAndPassword(hashedPass, []byte(password))

	if err != nil {
		return errors.New("Invalid Password")
	}

	return nil
}

func nilFunc() {
	var tempMysql *mysql.NullTime = nil
	var tempPq *pq.Error = nil
	sqlite3.Version()
	tempMysql = tempMysql
	tempPq = tempPq
}
