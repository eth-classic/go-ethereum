package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"os"

	"gopkg.in/urfave/cli.v1"


	"github.com/ether-core/go-ethereum/node"
	"github.com/ether-core/go-ethereum/rpc"
)

var (
	apiCommand = cli.Command{
		Action: execAPI,
		Name:   "api",
		Usage:  "Run any API command",
		Description: `
	The api command allows you to communicate (via IPC) with a running geth instance
	and run any RPC API method.

	Each parameter should be passed as JSON representation:
	- no quotations for numbers or booleans,
	- strings must be correctly quoted, like '"some value"' (quotes must be
	  included in string passed to application),
	- complex objects could be passed as JSON string.

	Examples:

		$ geth api eth getBlockByNumber 123 true
		$ geth eth getBlockByNumber '"latest"' true
		$ geth --chain morden api eth sendTransaction '{"from": "0x396599f365093186742c17aab158bf515e978bc7", "gas": "0x5208", "gasPrice": "0x02540be400", "to": "0xa02cee0fc1d3fb4dde86b79fe93e4140671fd949"}'

	Output will be written to stderr in JSON format.
		`,
	}
)

func execAPI(ctx *cli.Context) error {
	client, err := getClient(ctx)
	if err != nil {
		return err
	}

	if err := validateArguments(ctx, client); err != nil {
		return err
	}

	result, err := callRPC(ctx, client)
	if err != nil {
		return err
	}
	return prettyPrint(result)
}

func getClient(ctx *cli.Context) (rpc.Client, error) {
	chainDir := MustMakeChainDataDir(ctx)
	var uri = "ipc:" + node.DefaultIPCEndpoint(chainDir)
	return rpc.NewClient(uri)
}

func validateArguments(ctx *cli.Context, client rpc.Client) error {
	if len(ctx.Args()) < 2 {
		return fmt.Errorf("api command requires at least 2 arguments (module and method), %d provided",
			len(ctx.Args()))
	}
	modules, err := client.SupportedModules()
	if err != nil {
		return err
	}

	module := ctx.Args()[0]
	if _, ok := modules[module]; !ok {
		return fmt.Errorf("unknown API module: %s", module)
	}

	return nil
}

func callRPC(ctx *cli.Context, client rpc.Client) (interface{}, error) {
	var (
		module = ctx.Args()[0]
		method = ctx.Args()[1]
		args   = ctx.Args()[2:]
	)
	req := rpc.JSONRequest{
		Id:      json.RawMessage(strconv.Itoa(rand.Int())),
		Method:  module + "_" + method,
		Version: "2.0",
		Payload: json.RawMessage("[" + strings.Join(args, ",") + "]"),
	}

	if err := client.Send(req); err != nil {
		return nil, err
	}

	var res rpc.JSONResponse
	if err := client.Recv(&res); err != nil {
		return nil, err
	}
	if res.Error != nil {
		return nil, fmt.Errorf("error in %s_%s: %s (code: %d)",
			module, method, res.Error.Message, res.Error.Code)
	}
	if res.Result != nil {
		return res.Result, nil
	}

	return nil, errors.New("no API response")
}

func prettyPrint(result interface{}) error {
	jsonBytes, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return err
	}
	os.Stderr.Write(jsonBytes)
	return nil
}
