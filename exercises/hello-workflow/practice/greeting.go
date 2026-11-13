package hello

import (
	"go.temporal.io/sdk/workflow"
)

func GreetSomeone(ctx workflow.Context, name string) (string, error) {
	return "Bonjour " + name + " !", nil
}
