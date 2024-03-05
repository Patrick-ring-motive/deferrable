package deferrable

func LinearDefer(mainFunc func(),deferFuncs []func()) {
    for i := len(deferFuncs) - 1; i >= 0; i-- {
        defer deferFuncs[i]()
    }
    mainFunc()
}

func Defer(deferFuncs ...func())[]func(){
  return deferFuncs
}

