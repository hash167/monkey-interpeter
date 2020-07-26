package parser

import (
	"fmt"
	"strings"
)

var traceLevel int = 0

const traceHolderIdentString = "\t"

func indentLevel() string {
	return strings.Repeat(traceHolderIdentString, traceLevel-1)
}

func tracePrint(fs string) {
	fmt.Printf("%s%s\n", indentLevel(), fs)
}

func incIndent() { traceLevel = traceLevel + 1 }
func decIndent() { traceLevel = traceLevel - 1 }

func trace(msg string) string {
	incIndent()
	tracePrint("BEGIN" + msg)
	return msg
}

func untrace(msg string) {
	tracePrint("END" + msg)
	decIndent()
}
