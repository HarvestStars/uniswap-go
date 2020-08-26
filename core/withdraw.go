package core

// function [removed_A, removed_B, reverse_A_new, reverse_B_new, total_share_A_new, total_share_B_new, invariant] = withdraw(share_A, share_B, reverse_A, reverse_B, total_share_A, total_share_B)
//     removed_A = share_A ./ total_share_A .* reverse_A;
//     removed_B = share_B ./ total_share_B .* reverse_B;
//     reverse_A_new = reverse_A - removed_A;
//     reverse_B_new = reverse_B - removed_B;
//     total_share_A_new = total_share_A - share_A;
//     total_share_B_new = total_share_B - share_B;
//     invariant = reverse_A_new .* reverse_B_new;
// end

// WithDraw returns AB removed from pool
func WithDraw(shareA, shareB float64, pair *TradingPair) (float64, float64) {
	removedA := shareA * pair.APool / pair.AShare
	removedB := shareB * pair.BPool / pair.BShare
	pair.APool = pair.APool - removedA
	pair.BPool = pair.BPool - removedB
	pair.AShare = pair.AShare - shareA
	pair.BShare = pair.BShare - shareB
	pair.Invariant = pair.APool * pair.BPool
	return removedA, removedB
}
