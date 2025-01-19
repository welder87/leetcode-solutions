class Solution:
    def sortArrayByParityII(self, nums: list[int]) -> list[int]:
        odd = []
        even = []
        for num in nums:
            if num % 2 == 0:
                even.append(num)
            else:
                odd.append(num)
        result = [None] * len(nums)
        i = 0
        j = 0
        for k in range(len(result)):
            if k % 2 == 0:
                result[k] = even[i]
                i += 1
            else:
                result[k] = odd[j]
                j += 1
        return result
