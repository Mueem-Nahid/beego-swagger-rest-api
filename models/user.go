package models

import (
	// "errors"
	"database/sql"
	"fmt"
	"log"

	"github.com/beego/beego/v2/core/validation"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
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
	FirstName string `valid:"Required"`
	LastName  string `valid:"Required"`
	Email     string `valid:"Email; MaxSize(100)"`
	Phone     string `valid:"Required; Phone; Length(11)"`
	Password  string `valid:"Required; MinSize(4)"`
	// DoB       string `valid:"Required;Match(/(0?[1-9]|1[012])/(0?[1-9]|[12][0-9]|3[01])/((19|20)\\d\\d)/)"`
	DoB string
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func AddUser(u User) string {

	var message string
	valid := validation.Validation{}
	b, err := valid.Valid(&u)

	hash, _ := HashPassword(u.Password)
	fmt.Println("hashing: ", hash)

	if err != nil {

	}
	if !b {
		// validation does not pass
		// blabla...
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			message = err.Message
		}
	} else {
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

		sql := `INSERT INTO "user_info"("first_name", "last_name", "email", "phone", "password", "dob") VALUES ($1, $2, $3, $4, $5, $6)`
		_, e := db.Exec(sql, u.FirstName, u.LastName, u.Email, u.Phone, hash, u.DoB)
		CheckError(e)

		// check db
		err = db.Ping()
		CheckError(err)

		fmt.Println("---------->> Connected <<----------")
		message = "New user " + u.Email + " created"
		fmt.Println("New user " + message + " created")
	}
	return message
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
