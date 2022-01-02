package dynamicprogramming

// 给你 k 种面值的硬币，面值分别为 c1, c2 ... ck，
// 每种硬币的数量无限，再给一个总金额 amount，
// 问你最少需要几枚硬币凑出这个金额，如果不可能凑出，算法返回 -1 。

// coins 中是可选硬币面值，amount 是目标金额
func coinChange(coins []int, amount int) int {
	mem := make(map[int]int)
	return dp(coins, amount, mem)
}

func dp(coins []int, amount int, mem map[int]int) int {
	res := 2<<31 - 1
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	for _, c := range coins {
		// 计算子问题的结果
		sub := dp(coins, amount-c, mem)
		// 子问题无解
		if sub == -1 {
			continue
		}
		// 在子问题中选择最优解，然后加一
		res = min(res, sub+1)
	}
	if res == 2<<31-1 {
		return -1
	}
	if res == 2<<31-1 {
		mem[amount] = -1
	} else {
		mem[amount] = res
	}
	return mem[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
