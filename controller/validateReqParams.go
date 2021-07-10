package controller

import (
	"fmt"
	"strings"
)

type ReqParamsCreateUser struct {
	Name string `json:"name"`
}

func (reqC ReqParamsCreateUser) Validate() error {
	err := Name(reqC.Name).Validate()
	return err
}

type ReqParamsUpdateUser struct {
	Name string `json:"name"`
}

func (reqU ReqParamsUpdateUser) Validate() error {
	err := Name(reqU.Name).Validate()
	return err
}

type Name string

func (n Name) Validate() error {
	slice := strings.Split(string(n), "")
	len := len(slice)
	if len <= 0 {
		err := fmt.Errorf("%s", "Filled in with 0 characters.")
		return err
	}

	for i := 0; i < len; i++ {
		if slice[i] == " " || slice[i] == "　" {
			err := fmt.Errorf("%s", "Included space string")
			return err
		}
	}

	return nil
}

type ReqParamsDrawGacha struct {
	Times int `json:"times"`
}

func (reqDG ReqParamsDrawGacha) Validate() error {
	err := Times(reqDG.Times).Validate()
	return err
}

type Times int

func (t Times) Validate() error {

	if int(t) <= 0 {
		err := fmt.Errorf("%s", "The value of times must be at least 1.")
		return err
	}
	return nil
}
