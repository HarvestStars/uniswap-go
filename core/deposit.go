package core

// function [reverse_A_new, reverse_B_new, share_A_mint, share_B_mint, share_A_total_new, share_B_total_new, invariant] = deposit(value_A, value_B, reverse_A, reverse_B, total_share_A, total_share_B)
//     reverse_A_new = reverse_A + value_A;
//     reverse_B_new = reverse_B + value_B;
//     share_A_mint = mint_share(value_A, reverse_A, total_share_A);
//     share_B_mint = mint_share(value_B, reverse_B, total_share_B);
//     share_A_total_new = share_A_mint + total_share_A;
//     share_B_total_new = share_B_mint + total_share_B;
//     invariant = reverse_A_new .* reverse_B_new;
// end

// Deposit returns AB share minted.
func Deposit(depositA, depositB float64, pair *TradingPair) (float64, float64) {
	MintA := CalcShare(depositA, pair.APool, pair.AShare)
	MintB := CalcShare(depositB, pair.BPool, pair.BShare)
	pair.APool += depositA
	pair.BPool += depositB
	pair.AShare += MintA
	pair.BShare += MintB
	pair.Invariant = pair.APool * pair.BPool
	return MintA, MintB
}
