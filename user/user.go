package user

import (
	"errors"
	// "fmt"
)

type User struct {
	Name  string
	Age   int
	Email string
}

func (u User) ValidateUser() (string, []error) {
	c1, c2, c3, c4 := false, false, false, false
	var err []error
	for i, j := range u.Email {
		if string(j) == "@" {
			c1 = true
		}
		if j == '.' && i == len(u.Email)-4 {
			if u.Email[i+1] == 'c' && u.Email[i+2] == 'o' && u.Email[i+3] == 'm' && int(i+4) == len(u.Email) {
				c2 = true
			}
		}
	}
	if !c1 || !c2 {
		err = append(err, errors.New("Email: wrong email address\n"))
	}
	if len(u.Email) < 1 {
		err = append(err, errors.New("Email: couldn't be empty\n"))
	}
	if len(u.Name) >= 6 {
		c3 = true
	} else {
		err = append(err, errors.New("Name: no less than 6 character\n"))
	}
	if u.Age > 0 && u.Age <= 120 {
		c4 = true
	} else {
		err = append(err, errors.New("Age: age between 1 and 120\n"))
	}
	if c1 && c2 && c3 && c4 {
		return "Successfully registred", nil
	}
	return "Errors: ", err
}
