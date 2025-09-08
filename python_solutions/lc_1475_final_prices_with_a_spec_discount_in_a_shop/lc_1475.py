class Solution:
    def finalPrices(self, prices: list[int]) -> list[int]:
        # Time complexity: O(n * n). Space complexity: O(1).
        for i in range(0, len(prices) - 1):
            for j in range(i + 1, len(prices)):
                if prices[j] <= prices[i]:
                    prices[i] = prices[i] - prices[j]
                    break
        return prices
