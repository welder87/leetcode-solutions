package problem2446

/*
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
*/
func haveConflict(event1 []string, event2 []string) bool {
	return !(event1[0] > event2[1] || event1[1] < event2[0])
}
