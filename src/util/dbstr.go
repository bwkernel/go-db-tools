package util

import (
	"strings"
)

type DbConnectInfo struct {
	Server, User, Pwd, Database string
}

func (self *DbConnectInfo) Convert(connectStr string) {
	if len(connectStr) == 0 {
		panic("The input connect str is not formatted correctly. ")
	}

	connectArr := strings.Split(connectStr, ",")
	self.Server = connectArr[0]
	self.User = connectArr[1]
	self.Pwd = connectArr[2]
	self.Database = connectArr[3]

}
