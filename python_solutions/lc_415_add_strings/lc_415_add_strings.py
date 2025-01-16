class Solution:
    def addStrings(self, num1: str, num2: str) -> str:
        i = len(num1) - 1
        j = len(num2) - 1
        zero = ord("0")
        result = []
        ost = 0
        while i >= 0 or j >= 0:
            n1 = ord(num1[i]) if i >= 0 else zero
            n2 = ord(num2[j]) if j >= 0 else zero
            val = n1 - zero + n2 - zero + ost
            if val > 9:
                ost = 1
                result.append(str(val - 10))
            else:
                ost = 0
                result.append(str(val))
            i -= 1
            j -= 1
        if ost > 0:
            result.append("1")
        return "".join(reversed(result))
