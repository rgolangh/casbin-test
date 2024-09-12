package main

import (
	"fmt"
	"github.com/casbin/casbin"
)

func main() {
    e := casbin.NewEnforcer("model.conf", "policy.csv")
    fmt.Println(e.Enforce("user:default/guest", "orchestrator.workflows.foo", "read"))
    fmt.Println(e.Enforce("user:default/guest", "orchestrator.foo", "read"))
}
