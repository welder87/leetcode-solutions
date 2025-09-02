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

    def replaceElementsV2(self, arr: list[int]) -> list[int]:
        # Time complexity: O(n). Space complexity: O(1).
        # Solution: https://algo.monster/liteproblems/1299
        # Initialize the maximum value found to the right of the current element
        max_to_the_right = -1
        # Reverse iterate through the array
        for i in range(len(arr) - 1, -1, -1):
            # Store the current element before updating
            current_element = arr[i]
            # Replace the current element with the maximum value found so far
            arr[i] = max_to_the_right
            # Update the maximum value by comparing it with the previously stored current element
            max_to_the_right = max(max_to_the_right, current_element)
        # Return the modified array
        return arr
