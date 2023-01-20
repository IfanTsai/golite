package frontend

import (
	"fmt"
	"strings"
)

type typeStatement int

const (
	statementInsert typeStatement = iota
	statementSelect
)

type Statement struct {
	statementType typeStatement
}

func PrepareStatement(line string) *Statement {
	tokens := strings.Split(line, " ")
	switch tokens[0] {
	case "insert":
		return &Statement{statementType: statementInsert}
	case "select":
		return &Statement{statementType: statementSelect}
	}

	return nil
}

func ExecuteStatement(statement *Statement) {
	switch statement.statementType {
	case statementInsert:
		fmt.Println("This is where we would do an insert.")
	case statementSelect:
		fmt.Println("This is where we would do a select.")
	}
}
