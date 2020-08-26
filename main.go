package main

import (
	"fmt"

	"github.com/HarvestStars/uniswap-go/core"
	"github.com/HarvestStars/uniswap-go/setup"
)

var action string

func main() {
	var personalShareA float64 = 0
	var personalShareB float64 = 0
	setup.InitTradingPair()
	var feeTotal float64 = 0
	for {
		actionSelect()
		switch action {
		case "swap":
			index, value := swapSelect()
			fmt.Print("\n")
			fmt.Print("----------------------- 交易详情 -----------------------\n")
			receieved := core.Swap(index, value, core.LavaEthPair, &feeTotal)
			if index == 0 {
				fmt.Printf("消耗ETH:      %f,           购买Lava     %f \n", value, receieved)

			} else {
				fmt.Printf("消耗Lava:     %f,           购买ETH      %f \n", value, receieved)
			}

		case "deposit":
			index, value := depositSelect()
			fmt.Print("\n")
			fmt.Print("----------------------- 交易详情 -----------------------\n")
			if index == 0 {
				depositB := core.LavaEthPair.BPool / core.LavaEthPair.APool * value
				MintA, MintB := core.Deposit(value, depositB, core.LavaEthPair)
				personalShareA += MintA
				personalShareB += MintB
				fmt.Printf("抵押(eth/lava)    %f, %f \n", value, depositB)
				fmt.Printf("用户本次share     %f, %f \n", MintA, MintB)
				fmt.Printf("用户本次占比      %f, %f \n", MintA/core.LavaEthPair.AShare, MintB/core.LavaEthPair.BShare)
				fmt.Printf("用户累计占比      %f, %f \n", personalShareA/core.LavaEthPair.AShare, personalShareB/core.LavaEthPair.BShare)
			} else {
				depositA := core.LavaEthPair.APool / core.LavaEthPair.BPool * value
				MintA, MintB := core.Deposit(depositA, value, core.LavaEthPair)
				personalShareA += MintA
				personalShareB += MintB
				fmt.Printf("抵押(eth/lava)    %f, %f \n", depositA, value)
				fmt.Printf("用户本次share     %f, %f \n", MintA, MintB)
				fmt.Printf("用户本次占比      %f, %f \n", MintA/core.LavaEthPair.AShare, MintB/core.LavaEthPair.BShare)
				fmt.Printf("用户累计占比      %f, %f \n", personalShareA/core.LavaEthPair.AShare, personalShareB/core.LavaEthPair.BShare)
			}

		case "withdraw":
			fmt.Print("\n")
			fmt.Print("----------------------- 交易详情 -----------------------\n")
			removedA, removedB := core.WithDraw(personalShareA, personalShareB, core.LavaEthPair)
			fmt.Printf("赎回ETH         %f, Lava         %f\n", removedA, removedB)

		}
		fmt.Printf("交易对常量 %f \n", core.LavaEthPair.Invariant)
		fmt.Printf("eth_pool储量为      %f,           lava_pool储量为     %f \n", core.LavaEthPair.APool, core.LavaEthPair.BPool)
		fmt.Printf("eth_pool全局share为 %f,           lava的最新share为   %f \n", core.LavaEthPair.AShare, core.LavaEthPair.BShare)
		fmt.Printf("eth的最新价格为:    %f lava/eth,  lava的最新价格为:   %f eth/lava \n", core.LavaEthPair.APrice, core.LavaEthPair.BPrice)
	}
}
