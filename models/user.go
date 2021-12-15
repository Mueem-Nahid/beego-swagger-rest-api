package models

import (
	// "errors"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	UserList map[string]*User
)

func init() {
	UserList = make(map[string]*User)
	u := User{"Mueem", "Nahid", "mueem51@gmail.com", "01408454475", "1234", "29-04-1998"}
	UserList["user_11111"] = &u
}

type User struct {
	// Id        string
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Password  string
	DoB       string
}

func AddUser(u User) string {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "123456"
		dbname   = "user_db"
	)

	// u.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	// UserList[u.Id] = &u

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	sql := `INSERT INTO "user_info"("FirstName", "LastName", "Email", "Phone", "Password", "DoB") VALUES ($1, $2, $3, $4, $5, $6)`
	_, e := db.Exec(sql, u.FirstName, u.LastName, u.Email, u.Phone, u.Password, u.DoB)
	CheckError(e)

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("---------->> Connected <<----------")

	return u.FirstName
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

/*
func GetUser(uid string) (u *User, err error) {
	if u, ok := UserList[uid]; ok {
		return u, nil
	}
	return nil, errors.New("User not exists")
} */

func GetAllUsers() map[string]*User {
	return UserList
}

/*
func UpdateUser(uid string, uu *User) (a *User, err error) {
	if u, ok := UserList[uid]; ok {
		if uu.Username != "" {
			u.Username = uu.Username
		}
		if uu.Password != "" {
			u.Password = uu.Password
		}
		if uu.Profile.Age != 0 {
			u.Profile.Age = uu.Profile.Age
		}
		if uu.Profile.Address != "" {
			u.Profile.Address = uu.Profile.Address
		}
		if uu.Profile.Gender != "" {
			u.Profile.Gender = uu.Profile.Gender
		}
		if uu.Profile.Email != "" {
			u.Profile.Email = uu.Profile.Email
		}
		return u, nil
	}
	return nil, errors.New("User Not Exist")
}

func Login(username, password string) bool {
	for _, u := range UserList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func DeleteUser(uid string) {
	delete(UserList, uid)
} */
