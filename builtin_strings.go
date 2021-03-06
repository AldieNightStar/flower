package flower

import (
	"fmt"
	"strings"
	"sync"

	"github.com/AldieNightStar/golisper"
)

func builtinString(s *Scope) {
	str := newBuitinDict()
	str.m["iterate"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 1 {
			return nil, errNotEnoughArgs(s.LastLine, "str iterate", 1, 0)
		}
		str, err := EvalCast("str iterate", s, args[0], "")
		if err != nil {
			return nil, err
		}
		return &builtinStringIterator{str}, nil
	})
	str.m["join"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 2 {
			return nil, errNotEnoughArgs(s.LastLine, "str iterate", 2, len(args))
		}
		list, err := EvalCast[*builtinList]("str iterate", s, args[0], nil)
		if err != nil {
			return nil, err
		}
		joiner, err := EvalCast("str iterate", s, args[1], "")
		if err != nil {
			return nil, err
		}
		arr := make([]string, 0, len(list.list))
		for _, el := range list.list {
			arr = append(arr, fmt.Sprint(el))
		}
		return strings.Join(arr, joiner), nil
	})
	str.m["concat"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		sb := strings.Builder{}
		sb.Grow(32)
		evaled, err := s.EvalArrayValues(args)
		if err != nil {
			return nil, err
		}
		arr := make([]string, 0, len(evaled))
		for _, el := range evaled {
			arr = append(arr, fmt.Sprint(el))
		}
		return strings.Join(arr, ""), nil
	})
	str.m["str"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 1 {
			return nil, errNotEnoughArgs(s.LastLine, "str", 1, 0)
		}
		val, err := s.Eval(args[0])
		if err != nil {
			return nil, err
		}
		return fmt.Sprint(val), nil
	})
	str.m["sub"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 2 {
			return nil, errNotEnoughArgs(s.LastLine, "str sub", 2, len(args))
		}
		src, err := EvalCast("str sub", s, args[0], "")
		if err != nil {
			return nil, err
		}
		startF, err := EvalCast[float64]("str sub", s, args[1], 0)
		if err != nil {
			return nil, err
		}
		endF, err := EvalCast[float64]("str sub", s, args[2], 0)
		if err != nil {
			return nil, err
		}
		start := int(startF)
		end := int(endF)
		return src[start:end], nil
	})
	str.m["len"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 1 {
			return nil, errNotEnoughArgs(s.LastLine, "str len", 1, 0)
		}
		str, err := EvalCast("str len", s, args[0], "")
		if err != nil {
			return nil, err
		}
		return float64(len(str)), nil
	})
	str.m["split"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 2 {
			return nil, errNotEnoughArgs(s.LastLine, "str split", 2, len(args))
		}
		str, err := EvalCast("str split", s, args[0], "")
		if err != nil {
			return nil, err
		}
		sep, err := EvalCast("str split", s, args[1], "")
		if err != nil {
			return nil, err
		}
		var strs []string
		if len(args) > 2 {
			count, err := EvalCast[float64]("str split", s, args[2], 0)
			if err != nil {
				return nil, err
			}
			strs = strings.SplitN(str, sep, int(count))
		} else {
			strs = strings.Split(str, sep)
		}
		arr := make([]any, 0, len(strs))
		for _, s := range strs {
			arr = append(arr, s)
		}
		return &builtinList{arr, &sync.Mutex{}}, nil
	})
	str.m["find"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 2 {
			return nil, errNotEnoughArgs(s.LastLine, "str find", 2, len(args))
		}
		str, err := EvalCast("str find", s, args[0], "")
		if err != nil {
			return nil, err
		}
		sub, err := EvalCast("str find", s, args[1], "")
		if err != nil {
			return nil, err
		}
		return float64(strings.Index(str, sub)), nil
	})
	str.m["at"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 2 {
			return nil, errNotEnoughArgs(s.LastLine, "str at", 2, len(args))
		}
		str, err := EvalCast("str at", s, args[0], "")
		if err != nil {
			return nil, err
		}
		idf, err := EvalCast[float64]("str at", s, args[1], 0)
		if err != nil {
			return nil, err
		}
		id := int(idf)
		if id < 0 || id >= len(str) {
			return "", nil
		}
		return str[id : id+1], nil
	})
	str.m["rep"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 2 {
			return nil, errNotEnoughArgs(s.LastLine, "str rep", 2, len(args))
		}
		str, err := EvalCast("str rep", s, args[0], "")
		if err != nil {
			return nil, err
		}
		str1, err := EvalCast("str rep", s, args[1], "")
		if err != nil {
			return nil, err
		}
		str2, err := EvalCast("str rep", s, args[2], "")
		if err != nil {
			return nil, err
		}
		return strings.ReplaceAll(str, str1, str2), nil
	})
	str.m["mul"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 2 {
			return nil, errNotEnoughArgs(s.LastLine, "str mul", 2, len(args))
		}
		str, err := EvalCast("str mul", s, args[0], "")
		if err != nil {
			return nil, err
		}
		countF, err := EvalCast[float64]("str mul", s, args[1], 0)
		if err != nil {
			return nil, err
		}
		count := int(countF)
		return strings.Repeat(str, count), nil
	})
	s.Memory["print"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		elems, err := s.EvalArrayValues(args)
		if err != nil {
			return nil, err
		}
		fmt.Println(elems...)
		return nil, nil
	})

	s.Memory["str"] = str
}
