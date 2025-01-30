class Solution:
    def removeElement(self, nums: list[int], val: int) -> int:
        i, j, cnt = 0, len(nums) - 1, 0
        while i <= j:
            if nums[j] == val:
                nums[j], j, cnt = None, j - 1, cnt + 1
            elif nums[i] == val:
                nums[i], nums[j], i, j, cnt = nums[j], None, i + 1, j - 1, cnt + 1
            else:
                i += 1
        return len(nums) - cnt
