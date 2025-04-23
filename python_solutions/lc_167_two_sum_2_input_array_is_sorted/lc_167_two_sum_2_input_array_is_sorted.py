class Solution:
    def twoSum(self, numbers: list[int], target: int) -> list[int]:
        # Time complexity: O(n log n). Space complexity: O(1).
        # modified itmo
        for i in range(len(numbers)):
            num = target - numbers[i]
            left, right = -1, len(numbers)
            while right > left + 1:
                mid = left + (right - left) // 2
                if numbers[mid] == num and mid != i:
                    return [i + 1, mid + 1]
                if numbers[mid] <= num:
                    left = mid
                else:
                    right = mid
        return []

    def twoSumV1(self, numbers: list[int], target: int) -> list[int]:
        # Time complexity: O(n log n). Space complexity: O(1).
        # modified leetcode
        for i in range(len(numbers)):
            num = target - numbers[i]
            left, right = 0, len(numbers) - 1
            while left < right:
                mid = left + (right - left) // 2
                if numbers[mid] == num and mid != i:
                    return [min(mid, i) + 1, max(mid, i) + 1]
                if numbers[mid] <= num:
                    left = mid + 1
                else:
                    right = mid
        return []

    def twoSumV2(self, numbers: list[int], target: int) -> list[int]:
        # Time complexity: O(n). Space complexity: O(1).
        i, j = 0, len(numbers) - 1
        while i < j:
            sm = numbers[i] + numbers[j]
            if sm == target:
                return [i + 1, j + 1]
            if sm < target:
                i += 1
            else:
                j -= 1
        return []
