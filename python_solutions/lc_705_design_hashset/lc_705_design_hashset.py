from common_datastructures import ListNode


class MyHashSet:
    def __init__(self):
        self._size = 4
        self._power = 2
        self._capacity = 0
        self._chains = [None] * self._size

    def add(self, key: int) -> None:
        self._check_resize()
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

    def _check_resize(self):
        if 3 * self._capacity > 2 * self._size:
            self._resize()

    def _resize(self):
        self._capacity = 0
        self._power += 1
        self._size = 2**self._power
        old_chains = self._chains
        self._chains = [None] * self._size
        for idx in range(len(old_chains)):
            node = old_chains[idx]
            old_chains[idx] = None
            while node is not None:
                self.add(node.val)
                node = node.next
        old_chains = None
