package service

import (
	"context"
	"log"
	"os"

	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/storage/inmem"
	"github.com/open-policy-agent/opa/util"
)

type Request struct {
	Role   string `json:"role"`
	Method string `json:"method"`
	Path   string `json:"path"`
}

func Auth(ctx context.Context, input Request) (bool, error) {
	b, err := os.ReadFile(os.Getenv("AUTH_DATA"))
	if err != nil {
		log.Printf("%v", err)
		return false, err
	}

	var data map[string]interface{}
	err = util.UnmarshalJSON(b, &data)
	if err != nil {
		log.Printf("%v", err)
		return false, err
	}
	store := inmem.NewFromObject(data)
	b2, err := os.ReadFile(os.Getenv("POLICY_AUTH"))
	if err != nil {
		log.Printf("%v", err)
		return false, err
	}

	r := rego.New(
		rego.Query("data.auth.allow"),
		rego.Store(store),
		rego.Module("auth.rego", string(b2)),
		rego.Input(input))
	result, err := r.Eval(ctx)
	return result.Allowed(), err
}
