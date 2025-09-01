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

    def replaceElementsV1(self, arr: list[int]) -> list[int]:
        # Time complexity: O(n). Space complexity: O(1).
        mx = -1
        for i in range(-1, -(len(arr) + 1), -1):
            if arr[i] > mx:
                arr[i], mx = mx, arr[i]
            else:
                arr[i] = mx
        return arr
