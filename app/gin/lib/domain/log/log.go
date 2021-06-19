package log

import (
	"fmt"
	"local.packages/game-information/debug"
)

// "local.packages/game-information/lib/domain/log"
const prefixInfo = "[info] "
const prefixDebug = "[debug] "
const prefixError = "[error] "

// Info 通常ログ
func Info(message string) {
	output(prefixInfo + message)
}

// Debug 開発環境用ログ
func Debug(message string) {
	output(prefixDebug + message)
}

// DebugStruct 開発環境用ログ（構造体を出力）
func DebugStruct(v interface{}) {
	debug.LogStruct(v)
}

// Error エラーログ
func Error(message string) {
	output(prefixError + message)
}

func output(message string) {
	fmt.Println(message)
}
