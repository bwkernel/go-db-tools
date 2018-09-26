package actions

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"
	"tcminplay/db-tools/src/util"

	_ "github.com/denisenkom/go-mssqldb"
)

type ExecuteSqlFileAction struct {
}

func (self *ExecuteSqlFileAction) Handle(args []string) error {
	if len(args) <= 2 {
		fmt.Println("Input format error, like this: sqltools execsqlfile mssqlconnectstring sqlfile...")
		return nil
	}
	connectInfo := util.DbConnectInfo{}
	connectInfo.Convert(args[2])
	connectStr := convertConnect(connectInfo)
	fmt.Println("Current connect: " + connectStr)

	db, err := sql.Open("mssql", connectStr)
	if err != nil {
		return err
	}
	defer db.Close()

	files, err := matchFilePaths(args[3:])
	if err != nil {
		return err
	}
	for _, v := range files {
		content, err := ioutil.ReadFile(v)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		sql := string(content)
		for _, s := range splitGo(sql) {
			r, err := db.Exec(s)
			if err != nil {
				return err
			}
			affected, err := r.RowsAffected()
			fmt.Println("Row affected: ", affected, ", err: ", err)
		}
	}
	return nil
}

func matchFilePaths(originalFiles []string) ([]string, error) {
	if len(originalFiles) == 0 {
		return nil, util.Err_NULL
	}

	for _, v := range originalFiles {
		if util.FileExists(v) {
			continue
		}

		dic := util.GetCurrentDirectory()
		file := dic + v
		if !util.FileExists(file) {
			panic(util.Err_NULL)
		}
		v = file
	}
	return originalFiles, nil
}

func splitGo(sql string) []string {
	if strings.Contains(sql, "go\n") {
		return strings.Split(sql, "go\n")
	} else {
		return strings.Split(sql, "GO\n")
	}
}

func convertConnect(connectInfo util.DbConnectInfo) string {
	return "server=" + connectInfo.Server + ";user id=" + connectInfo.User +
		";password=" + connectInfo.Pwd + ";database=" + connectInfo.Database
}
