package main

import (
	"context"
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/beego/beego/v2/core/utils"
)

func main() {
	cacheURL := os.Getenv("_REDIS_URL_CORe")
	u, err := url.Parse(cacheURL)
	if err != nil {
		panic("bad_REDIS_URL_CORE")
	}

	if err = cache.Initialize(u.Scheme); err != nil {
		panic(fmt.Sprintf("failed to initialize cache: %v", err))
	}

	configPath := flag.String("c", "", "Specify the yaml config file path")
	flag.Parse()

	if configPath == nil {
		flag.Usage()
		panic(fmt.Sprintf("load configurations error: %s\n", err))
	}

	vCtx := context.WithValue(context.Background(), utils.NodeID, utils.GenerateNodeID())
	ctx, cancel := context.WithCancel(vCtx)
	defer cancel()

	if err := logger.Init(ctx); err != nil {
		panic(err)
	}

	
}