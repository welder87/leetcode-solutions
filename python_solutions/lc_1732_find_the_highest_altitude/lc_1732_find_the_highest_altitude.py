class Solution:
    def largestAltitude(self, gain: list[int]) -> int:
        length = len(gain)
        prefix_sum = [0] * (length + 1)
        for i in range(length):
            prefix_sum[i + 1] = prefix_sum[i] + gain[i]
        return max(prefix_sum)

    def largestAltitudeV1(self, gain: list[int]) -> int:
        current, mx = 0, 0
        for i in range(len(gain)):
            current += gain[i]
            if current > mx:
                mx = current
        return mx
