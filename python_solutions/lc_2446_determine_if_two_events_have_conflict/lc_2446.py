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

    def haveConflictV1(self, event1: list[str], event2: list[str]) -> bool:
        """
        Determine if two events have a time conflict.

        The solution checks for non-conflict conditions:
        - if one event ends before the other starts
            (either event1[1] < event2[0] or event1[0] > event2[1]),
            there's no conflict.
        - Otherwise, the events must overlap.

        Since string comparison works correctly with the HH:MM format,
        we can directly compare the time strings without converting them to numbers.

        Time complexity: O(1). Space complexity: O(1).
        Solution: https://algo.monster/liteproblems/2446.
        Method: De Morgan's law.

        Args:
            event1: List containing [start_time, end_time] as strings in "HH:MM" format
            event2: List containing [start_time, end_time] as strings in "HH:MM" format

        Returns:
            bool: True if events overlap, False otherwise
        """
        # Two events DO NOT conflict if:
        # 1. Event1 starts after event2 ends (event1[0] > event2[1])
        # 2. Event1 ends before event2 starts (event1[1] < event2[0])
        #
        # Therefore, events HAVE a conflict when neither of these conditions is true
        # Using De Morgan's law: not (A or B) = (not A) and (not B)
        # Which means: events overlap when event1 doesn't start after event2 ends
        # AND event1 doesn't end before event2 starts

        return not (event1[0] > event2[1] or event1[1] < event2[0])

    def haveConflictV2(self, event1: list[str], event2: list[str]) -> bool:
        # Time complexity: O(1). Space complexity: O(1).
        # Solution: https://rutube.ru/video/bf8a2ceb596586bdf3f9c10d56e4e927/.
        return max(event1[0], event2[0]) <= min(event1[1], event2[1])
