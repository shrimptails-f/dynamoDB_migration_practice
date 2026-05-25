package env

// EnvInterface はEnvの抽象です
type EnvInterface interface {
	GetEnv(key string) (string, error)
}
