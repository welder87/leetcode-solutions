class Solution:
    def validMountainArray(self, arr: list[int]) -> bool:
        length = len(arr)
        if length < 3:
            return False
        last_index = length - 1
        for i in range(1, length):
            prev = i - 1
            if (i == last_index and arr[prev] < arr[i]) or arr[prev] == arr[i]:
                return False
            if arr[prev] > arr[i]:
                if prev == 0:
                    return False
                break
        for i in range(i, length):
            prev = i - 1
            if arr[prev] == arr[i]:
                return False
            if arr[prev] < arr[i]:
                return False
        return True
