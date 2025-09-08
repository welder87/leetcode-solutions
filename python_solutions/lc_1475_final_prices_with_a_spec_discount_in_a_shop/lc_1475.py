from collections import deque


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

    def finalPricesV2(self, prices: list[int]) -> list[int]:
        # Time complexity: O(n + n). Space complexity:  O(n + n).
        # Solution: https://leetcode.com/problems/final-prices-with-a-special-discount-in-a-shop/editorial/
        # Create a copy of prices array to store discounted prices
        result = prices.copy()
        stack = deque()
        for i in range(len(prices)):
            # Process items that can be discounted by current price
            while stack and prices[stack[-1]] >= prices[i]:
                # Apply discount to previous item using current price
                result[stack.pop()] -= prices[i]
            # Add current index to stack
            stack.append(i)
        return result
