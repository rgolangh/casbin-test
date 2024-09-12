package main

import (
	"fmt"
	"github.com/casbin/casbin"
)

func main() {
    fmt.Println("casbin test")

    e := casbin.NewEnforcer("model.conf", "policy.csv")
    fmt.Println(e.Enforce("user:default/guest", "orchestrator.workflows.foo", "read"))
    

}
