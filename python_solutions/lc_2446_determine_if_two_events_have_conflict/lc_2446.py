class Solution:
    def haveConflict(self, event1: list[str], event2: list[str]) -> bool:
        # Time complexity: O(1). Space complexity: O(1).
        ev1 = [self._convert(event1[0]), self._convert(event1[1])]
        ev2 = [self._convert(event2[0]), self._convert(event2[1])]
        return (
            ev2[0] <= ev1[0] <= ev2[1]
            or ev2[0] <= ev1[1] <= ev2[1]
            or ev1[0] <= ev2[0] <= ev1[1]
            or ev1[0] <= ev2[1] <= ev1[1]
        )

    def _convert(self, val: str) -> int:
        splitted = val.split(":")
        if splitted[0][0] == "0":
            h = int(splitted[0][1])
        else:
            h = int(splitted[0])
        if splitted[1][0] == "0":
            m = int(splitted[1][1])
        else:
            m = int(splitted[1])
        return h * 60 + m
