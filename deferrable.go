package deferrable
import(
  "fmt"
)
func LinearDefer(mainFunc func(),deferFuncs []func()) {
    for i := len(deferFuncs) - 1; i >= 0; i-- {
        defer deferFuncs[i]()
    }
    mainFunc()
}

func Defer(deferFuncs ...func())[]func(){
  fmt.Print("cheese")
  return deferFuncs
}

