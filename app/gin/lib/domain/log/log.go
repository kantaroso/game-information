package log

import (
	"os"

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

		_logger, err = config.Build()

	} else {
		// 開発環境用（デフォルト）:
		// - Human-readableなコンソール形式
		// - 標準エラー出力
		// - DEBUGレベル以上を出力
		config := zap.NewDevelopmentConfig()

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