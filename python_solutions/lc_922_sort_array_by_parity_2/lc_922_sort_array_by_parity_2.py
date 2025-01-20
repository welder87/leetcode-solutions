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

    def sortArrayByParityIIV1(self, nums: list[int]) -> list[int]:
        last_index = len(nums) - 1
        i = 0
        j = last_index
        while i < j:
            is_odd_1 = nums[i] % 2 == 0
            is_odd_2 = nums[j] % 2 == 0
            if is_odd_1 and not is_odd_2:
                nums[i], nums[j] = nums[j], nums[i]
                i, j = i + 1, j - 1
            elif is_odd_1 and is_odd_2 or not is_odd_1 and is_odd_2:
                j -= 1
            else:
                i += 1
        for i in range(len(nums) // 2):
            if i % 2 == 0:
                j = last_index - i
                nums[i], nums[j] = nums[j], nums[i]
        return nums
