package main

import (
	"context"
	"github.com/go-chassis/go-chassis"
	"github.com/go-chassis/go-chassis/client/rest"
	"github.com/go-chassis/go-chassis/core"
	"github.com/go-chassis/go-chassis/core/common"
	"github.com/go-chassis/go-chassis/pkg/util/httputil"
	"github.com/go-mesh/openlogging"
)

//if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/{path}/{to}/client/
func main() {
	//Init framework
	if err := chassis.Init(); err != nil {
		openlogging.Error("Init failed." + err.Error())
		return
	}

	req, err := rest.NewRequest("GET", "http://TLSService/sayhello/world", nil)
	if err != nil {
		openlogging.Error("new request failed.")
		return
	}

	ctx := context.WithValue(context.TODO(), common.ContextHeaderKey{}, map[string]string{
		"user": "peter",
	})
	resp, err := core.NewRestInvoker().ContextDo(ctx, req)
	if err != nil {
		openlogging.Error("do request failed.")
		return
	}
	defer resp.Body.Close()
	openlogging.Info("REST Server sayhello[GET]: " + string(httputil.ReadBody(resp)))
}
