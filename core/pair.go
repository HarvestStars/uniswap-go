package core

type TradingPair struct {
	Invariant float64
	APool     float64 `ini:"InitialEth"`
	BPool     float64 `ini:"InitialLava"`
	AShare    float64
	BShare    float64
	APrice    float64
	BPrice    float64
}

var LavaEthPair = &TradingPair{}
