package flower

import (
	"fmt"
	"sync"

	"github.com/AldieNightStar/golisper"
)

type builtinTypeGetter interface {
	Type() string
}

func getTypeOf(a any) string {
	if a == nil {
		return "nil"
	}
	if t, ok := a.(builtinTypeGetter); ok {
		return t.Type()
	}
	if lispVal, ok := a.(*golisper.Value); ok {
		if lispVal.Type == golisper.TYPE_ETC_STRING {
			return "token:etc_string"
		} else if lispVal.Type == golisper.TYPE_NUMBER {
			return "token:number"
		} else if lispVal.Type == golisper.TYPE_STRING {
			return "token:string"
		} else if lispVal.Type == golisper.TYPE_TAG {
			return "token:tag"
		}
	}
	if _, ok := a.(string); ok {
		return "string"
	}
	if _, ok := a.(float64); ok {
		return "number"
	}
	if _, ok := a.(bool); ok {
		return "bool"
	}
	if _, ok := a.(error); ok {
		return "error"
	}
	if _, ok := a.(builtinIterator); ok {
		return "iterator"
	}
	if _, ok := a.(builtinIteration); ok {
		return "iteration"
	}
	if _, ok := a.(*sync.Mutex); ok {
		return "mutex"
	}
	return fmt.Sprintf("%T", a)
}
