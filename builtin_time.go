package flower

import (
	"time"

	"github.com/AldieNightStar/golisper"
)

func builtinTime(s *Scope) {
	d := newBuitinDict()
	d.m["sleep"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 1 {
			return nil, errNotEnoughArgs(s.LastLine, "time sleep", 1, 0)
		}
		f, err := EvalCast[float64]("sleep", s, args[0], 0)
		if err != nil {
			return nil, err
		}
		time.Sleep(time.Millisecond * time.Duration(f))
		return nil, nil
	})

	s.Memory["time"] = d
}
