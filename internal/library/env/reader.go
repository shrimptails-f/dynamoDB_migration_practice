package env

import (
	"fmt"
	"os"
	"strings"
)

// Env は両方の機能を持つ具象構造体です
type Env struct{}

var _ EnvInterface = (*Env)(nil)

// New は Env のインスタンスを返します
func New() *Env {
	return &Env{}
}

// GetEnv は環境変数を取得します。空文字の場合はエラーを返します。
func (o *Env) GetEnv(key string) (string, error) {
	value := os.Getenv(key)
	if strings.TrimSpace(value) == "" {
		return "", fmt.Errorf("環境変数 %s が設定されていません", key)
	}

	return value, nil
}
