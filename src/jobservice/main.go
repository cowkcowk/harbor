package main

import (
	"context"
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/goharbor/harbor/src/common"
	"github.com/goharbor/harbor/src/jobservice/common/utils"

	"github.com/goharbor/harbor/src/lib/cache"
	cfgLib "github.com/goharbor/harbor/src/lib/config"
)

func main() {
	cfgLib.DefaultCfgManager = common.RestCfgManager
	if err := cfgLib.DefaultMgr().Load(context.Background()); err != nil {
		panic(fmt.Sprintf("failed to load configuration, error: %v", err))
	}

	configPath := flag.String("c", "", "Specify the yaml config file path")
	flag.Parse()

	if configPath == nil || utils.IsEmptyStr(*configPath) {
		flag.Parse()
		panic("no config file is specified")
	}

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
