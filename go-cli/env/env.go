package env

type Env struct {
	m map[string]string
}

func NewEnv() *Env {
	return &Env{make(map[string]string)}
}

func (e *Env) Init() {
	e.m["env1"] = "env1 init"
	e.m["env2"] = "env2 init"
}

func (e *Env) GetMap() map[string]string {
	return e.m
}

func (e *Env) Add(key string, value string) {
	e.m[key] = value
}
