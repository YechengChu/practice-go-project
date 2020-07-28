// models.user.go

package main

import (
	"errors"
	"regexp"
	"strings"
	"unicode"
)

// Check if the username and password combination is valid
func isUserValid(username, password string) bool {
	hasAcc, passwordInDB := query(username)
	if hasAcc {
		if passwordInDB == password {
			return true
		}
	}
	return false
}

// Register a new user with the given username and password
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

  insert(username, password)

	return username, nil
}

// Check if the supplied username is available
func isUsernameAvailable(username string) bool {
  hasAcc,_ := query(username)
	if hasAcc {
		return false
	}
	return true
}
