package tritium

import(
	tp "athena/proto"
)

type Engine interface {
	Run(transform *tp.Transform, input string, vars map[string]string) (data string, exports [][]string, logs []string)
}