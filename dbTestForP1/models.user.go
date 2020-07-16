// models.user.go

package main

import (
	"errors"
	"regexp"
	"strings"
	"unicode"
)

// type user struct {
// 	Username string `json:"username"`
// 	Password string `json:"-"`
// }
//
// // For this demo, we're storing the user list in memory
// // We also have some users predefined.
// // In a real application, this list will most likely be fetched
// // from a database. Moreover, in production settings, you should
// // store passwords securely by salting and hashing them instead
// // of using them as we're doing in this demo
// var userList = []user{
// 	user{Username: "user1", Password: "pass1"},
// 	user{Username: "user2", Password: "pass2"},
// 	user{Username: "user3", Password: "pass3"},
// }

// Check if the username and password combination is valid
func isUserValid(username, password string) bool {
	// for _, u := range userList {
	// 	if u.Username == username && u.Password == password {
	// 		return true
	// 	}
	// }
	hasAcc, passwordInDB := query(username)
	if hasAcc {
		if passwordInDB == password {
			return true
		}
	}
	return false
}

// Register a new user with the given username and password
// NOTE: For this demo, we
func registerNewUser(username, password string) (string, error) {
	if strings.TrimSpace(password) == "" {
		return "", errors.New("The password can't be empty")
	} else if !isUsernameAvailable(username) {
		return "", errors.New("The username isn't available")
	}
	//pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	isEmail := reg.MatchString(username)
	// 匹配手机
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg2 := regexp.MustCompile(regular)
	isMobile := reg2.MatchString(username)
	// 验证用户名
	if (isEmail || isMobile) != true {
		return "", errors.New("Your user name should either be phone number or email address!")
	}
	// 验证密码长度
	if len(password) < 8 {
		return "", errors.New("Your password is too short!")
	}
	// 验证密码是否有大小写
	var hasUpperCase, hasLowercase bool
	for _, c := range password {
		switch {
		case unicode.IsUpper(c):
			hasUpperCase = true
			if hasLowercase {
				break
			}
		case unicode.IsLower(c):
			hasLowercase = true
			if hasUpperCase {
				break
			}
		}
	}
	if (hasLowercase && hasUpperCase) != true {
		return "", errors.New("Your password should contain both upper and lower cases!")
	}

	// u := user{Username: username, Password: password}
	//
	// userList = append(userList, u)
  insert(username, password)

	return username, nil
}

// Check if the supplied username is available
func isUsernameAvailable(username string) bool {
	// for _, u := range userList {
	// 	if u.Username == username {
	// 		return false
	// 	}
	// }
  hasAcc,_ := query(username)
	if hasAcc {
		return false
	}
	return true
}
