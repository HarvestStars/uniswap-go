package core

// function [Token_B_out, invariant_new, reserve_A_New, reserve_B_New, price_A_new] = swap(Token_A_in, invariant, reserve_A, reserve_B)
//     reserve_A_New = reserve_A + Token_A_in;
//     fee = Token_A_in .* 0.0025;
//     reserve_B_New = invariant ./ (reserve_A_New - fee);
//     Token_B_out = reserve_B - reserve_B_New;
//     invariant_new = reserve_A_New .*  reserve_B_New;
//     price_A_new = Token_B_out./Token_A_in;
// end

// Swap return the token receieved which is brought using inputIndex
// exp. Eth/Lava
// inputIndex = 0: return Eth input
// inputIndex = 1: return Lava input
func Swap(inputIndex int, expenseValue float64, pair *TradingPair, feeTotal *float64) float64 {
	invariant := pair.Invariant
	fee := expenseValue * 0.003
	*feeTotal = *feeTotal + fee
	var TokenOut float64
	if inputIndex == 0 {
		reserveANew := pair.APool + expenseValue
		reserveBNew := invariant / (reserveANew - fee)
		TokenOut = pair.BPool - reserveBNew
		pair.APool = reserveANew
		pair.BPool = reserveBNew
		pair.Invariant = reserveANew * reserveBNew
		pair.APrice = TokenOut / expenseValue
		pair.BPrice = expenseValue / TokenOut
	} else {
		reserveBNew := pair.BPool + expenseValue
		reserveANew := invariant / (reserveBNew - fee)
		TokenOut = pair.APool - reserveANew
		pair.APool = reserveANew
		pair.BPool = reserveBNew
		pair.Invariant = reserveANew * reserveBNew
		pair.BPrice = TokenOut / expenseValue
		pair.APrice = expenseValue / TokenOut
	}
	return TokenOut
}
