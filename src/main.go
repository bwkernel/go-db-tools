package main

import (
	"fmt"
	"os"

	. "tcminplay/db-tools/src/actions"
)

func main() {
	command := os.Args[1]

	actionMap := regAction()
	if len(actionMap) == 0 {
		fmt.Println("No action.")
		return
	}

	action, ok := actionMap[command]
	if ok == false {
		fmt.Println("Get action failed!")
		return
	}

	err := action.Handle(os.Args)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Action execution completed!")
	}
}

func regAction() map[string]Action {
	actionMap := make(map[string]Action)
	actionMap["execsqlfile"] = &ExecuteSqlFileAction{}

	return actionMap
}
