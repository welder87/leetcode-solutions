class Solution:
    def replaceElements(self, arr: list[int]) -> list[int]:
        # Time complexity: O(n^2). Space complexity: O(1).
        # Method: brute force
        for i in range(len(arr) - 1):
            mx = i + 1
            for j in range(i + 2, len(arr)):
                if arr[mx] < arr[j]:
                    mx = j
            arr[i] = arr[mx]
        arr[-1] = -1
        return arr
