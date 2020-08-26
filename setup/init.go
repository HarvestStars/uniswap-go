package setup

import (
	"log"

	"github.com/HarvestStars/uniswap-go/core"
	"github.com/go-ini/ini"
)

// InitTradingPair 初始化交易对
func InitTradingPair() {
	cfg, err := ini.Load("./conf/config.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	mapTo(cfg, "uniswap", core.LavaEthPair)
	core.LavaEthPair.Invariant = core.CalcInvariant(core.LavaEthPair.APool, core.LavaEthPair.BPool)
	core.LavaEthPair.APrice = core.LavaEthPair.BPool / core.LavaEthPair.APool
	core.LavaEthPair.BPrice = core.LavaEthPair.APool / core.LavaEthPair.BPool
	core.LavaEthPair.AShare = core.CalcInitialShare(core.LavaEthPair.APool, core.LavaEthPair.BPool)
	core.LavaEthPair.BShare = core.CalcInitialShare(core.LavaEthPair.APool, core.LavaEthPair.BPool)
}

func mapTo(cfg *ini.File, section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}
