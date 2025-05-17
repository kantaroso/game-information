package log

import (
	"os"
	"time"

	"go.uber.org/zap"
)

// 内部で使用するzapロガーインスタンス
var _logger *zap.Logger

// パッケージ読み込み時にロガーを初期化します。
// 環境変数 IS_PROD の値に応じてプロダクション設定か開発設定を切り替えます。
func init() {
	var err error
	// 修正：os.Getenvは第2引数を取らない
	isProd := os.Getenv("IS_PROD") == "true"

	if isProd {
		// 本番環境用:
		// - JSON形式
		// - 標準出力
		// - INFOレベル以上を出力
		config := zap.NewProductionConfig()
		level = zap.InfoLevel // Infoレベル以上を出力

		_logger, err = config.Build()

	} else {
		// 開発環境用（デフォルト）:
		// - Human-readableなコンソール形式
		// - 標準エラー出力
		// - DEBUGレベル以上を出力
		config := zap.NewDevelopmentConfig()
		level = zap.DebugLevel // Debugレベル以上を出力

		_logger, err = config.Build()
	}

	if err != nil {
		// ロガーの初期化に失敗した場合はパニック
		panic("failed to initialize zap logger: " + err.Error())
	}
}

// Debug はデバッグレベルのログを出力します。
// val が nil でない場合、"data" キーの値として出力します。
func Debug(msg string, val ...interface{}) {
	if len(val) > 0 && val[0] != nil {
		_logger.WithOptions(zap.AddCallerSkip(1)).Debug(msg, zap.Any("data", val[0]))
	} else {
		_logger.WithOptions(zap.AddCallerSkip(1)).Debug(msg)
	}
}

// Info は情報レベルのログを出力します。
// val が nil でない場合、"data" キーの値として出力します。
func Info(msg string, val ...interface{}) {
	if len(val) > 0 && val[0] != nil {
		_logger.WithOptions(zap.AddCallerSkip(1)).Info(msg, zap.Any("data", val[0]))
	} else {
		_logger.WithOptions(zap.AddCallerSkip(1)).Info(msg)
	}
}

// Error はエラーレベルのログを出力します。
// val が nil でない場合、"data" キーの値として出力します。
func Error(msg string, val ...interface{}) {
	if len(val) > 0 && val[0] != nil {
		_logger.WithOptions(zap.AddCallerSkip(1)).Error(msg, zap.Any("data", val[0]))
	} else {
		_logger.WithOptions(zap.AddCallerSkip(1)).Error(msg)
	}
}

// 追加の機能: 特定のフィールド名でログを残したい場合のヘルパー関数

// DebugWith はデバッグレベルのログを、指定したキーでvalを出力します。
func DebugWith(msg string, key string, val interface{}) {
	_logger.WithOptions(zap.AddCallerSkip(1)).Debug(msg, zap.Any(key, val))
}

// InfoWith は情報レベルのログを、指定したキーでvalを出力します。
func InfoWith(msg string, key string, val interface{}) {
	_logger.WithOptions(zap.AddCallerSkip(1)).Info(msg, zap.Any(key, val))
}

// ErrorWith はエラーレベルのログを、指定したキーでvalを出力します。
func ErrorWith(msg string, key string, val interface{}) {
	_logger.WithOptions(zap.AddCallerSkip(1)).Error(msg, zap.Any(key, val))
}

// 追加のヘルパー関数：エラーオブジェクト専用のログ出力

// ErrorErr はエラーオブジェクトをerrorフィールドとして出力するヘルパー関数です。
func ErrorErr(msg string, err error) {
	if err != nil {
		_logger.WithOptions(zap.AddCallerSkip(1)).Error(msg, zap.Error(err))
	} else {
		_logger.WithOptions(zap.AddCallerSkip(1)).Error(msg)
	}
}

// InfoErr は情報レベルでエラーオブジェクトをerrorフィールドとして出力するヘルパー関数です。
func InfoErr(msg string, err error) {
	if err != nil {
		_logger.WithOptions(zap.AddCallerSkip(1)).Info(msg, zap.Error(err))
	} else {
		_logger.WithOptions(zap.AddCallerSkip(1)).Info(msg)
	}
}

// DebugErr はデバッグレベルでエラーオブジェクトをerrorフィールドとして出力するヘルパー関数です。
func DebugErr(msg string, err error) {
	if err != nil {
		_logger.WithOptions(zap.AddCallerSkip(1)).Debug(msg, zap.Error(err))
	} else {
		_logger.WithOptions(zap.AddCallerSkip(1)).Debug(msg)
	}
}

// 従来のフィールド指定ログ出力も残しておく

// DebugFields はデバッグレベルのログを複数のフィールドと共に出力します。
func DebugFields(msg string, fields ...zap.Field) {
	_logger.WithOptions(zap.AddCallerSkip(1)).Debug(msg, fields...)
}

// InfoFields は情報レベルのログを複数のフィールドと共に出力します。
func InfoFields(msg string, fields ...zap.Field) {
	_logger.WithOptions(zap.AddCallerSkip(1)).Info(msg, fields...)
}

// ErrorFields はエラーレベルのログを複数のフィールドと共に出力します。
func ErrorFields(msg string, fields ...zap.Field) {
	_logger.WithOptions(zap.AddCallerSkip(1)).Error(msg, fields...)
}

// Field関数を提供して、zap.Fieldのエイリアスを作成
func String(key string, val string) zap.Field {
	return zap.String(key, val)
}

func Int(key string, val int) zap.Field {
	return zap.Int(key, val)
}

func Int64(key string, val int64) zap.Field {
	return zap.Int64(key, val)
}

func Float64(key string, val float64) zap.Field {
	return zap.Float64(key, val)
}

func Bool(key string, val bool) zap.Field {
	return zap.Bool(key, val)
}

func Duration(key string, val time.Duration) zap.Field {
	return zap.Duration(key, val)
}

func Time(key string, val time.Time) zap.Field {
	return zap.Time(key, val)
}

// Any は任意の型をフィールドとして出力します
func Any(key string, val interface{}) zap.Field {
	return zap.Any(key, val)
}