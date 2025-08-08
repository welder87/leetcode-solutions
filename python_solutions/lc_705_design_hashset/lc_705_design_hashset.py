from common_datastructures import ListNode


class MyHashSet:
    def __init__(self):
        self._size = 4
        self._capacity = 0
        self._chains = [None] * self._size

    def add(self, key: int) -> None:
        idx = self._cals_idx(key)
        if self._chains[idx] is None:
            self._chains[idx] = ListNode(val=key)
            self._capacity += 1
            return
        prev, cur = None, self._chains[idx]
        while cur is not None:
            if cur.val == key:
                return
            prev = cur
            cur = cur.next
        prev.next = ListNode(val=key)
        self._capacity += 1

    def remove(self, key: int) -> None:
        idx = self._cals_idx(key)
        if (chain := self._chains[idx]) is None:
            return
        while chain is not None:
            if chain.val == key:
                nxt = chain.next
                chain.next = None
                chain = nxt
                self._chains[idx] = chain
                self._capacity -= 1
                return
            chain = chain.next

    def contains(self, key: int) -> bool:
        idx = self._cals_idx(key)
        if (chain := self._chains[idx]) is None:
            return False
        while chain is not None:
            if chain.val == key:
                return True
            chain = chain.next
        return False

    def _cals_idx(self, key: int) -> int:
        return hash(key) % self._size
