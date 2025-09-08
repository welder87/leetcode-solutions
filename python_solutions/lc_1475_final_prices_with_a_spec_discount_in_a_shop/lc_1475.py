class Solution:
    def finalPrices(self, prices: list[int]) -> list[int]:
        # Time complexity: O(n * n). Space complexity: O(1).
        for i in range(0, len(prices) - 1):
            for j in range(i + 1, len(prices)):
                if prices[j] <= prices[i]:
                    prices[i] = prices[i] - prices[j]
                    break
        return prices

    def finalPricesV1(self, prices: list[int]) -> list[int]:
        # Time complexity: O(n + n). Space complexity: O(n + n).
        ans = [*prices]
        stack = []
        for idx, price in enumerate(prices):
            while stack and prices[stack[-1]] >= price:
                prev_idx = stack.pop()
                ans[prev_idx] = prices[prev_idx] - price
            stack.append(idx)
        return ans
