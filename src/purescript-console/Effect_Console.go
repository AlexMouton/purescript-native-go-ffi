package purescript_console

import (
	"fmt"
	"os"
	. "purescript"
)

func init() {
	exports := Foreign("Effect.Console")

	exports["log"] = func(s Any) Any {
		return func() Any {
			fmt.Println(s)
			return nil
		}
	}

	exports["warn"] = func(s Any) Any {
		return func() Any {
			fmt.Println(s)
			return nil
		}
	}

	exports["error"] = func(s Any) Any {
		return func() Any {
			fmt.Fprintln(os.Stderr, s)
			return nil
		}
	}

}
