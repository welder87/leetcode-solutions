from typing import Optional


class Node:
    __slots__ = ("children", "is_end")

    def __init__(self):
        self.children: list[Optional[Node]] = [None] * 26
        self.is_end: bool = False


class Trie:
    # Space complexity: O(alphabet_size * key_size * key_count).

    __slots__ = ("root", "NIL")

    def __init__(self):
        self.root = Node()
        self.NIL = ord("a")

    def insert(self, word: str) -> None:
        # Time complexity: O(n).
        node = self.root
        for char in word:
            idx = ord(char) - self.NIL
            cur = node.children[idx]
            if cur is None:
                cur = Node()
                node.children[idx] = cur
            node = cur
        node.is_end = True

    def search(self, word: str) -> bool:
        # Time complexity: O(n). Space complexity: O(1).
        node = self.root
        for char in word:
            cur = node.children[ord(char) - self.NIL]
            if cur is None:
                return False
            node = cur
        return node.is_end

    def startsWith(self, prefix: str) -> bool:
        # Time complexity: O(n). Space complexity: O(1).
        node = self.root
        for char in prefix:
            cur = node.children[ord(char) - self.NIL]
            if cur is None:
                return False
            node = cur
        return True


NIL = ord("a")


class TrieV1:
    # Space complexity: O(alphabet_size * key_size * key_count).

    __slots__ = ("children", "is_end")

    def __init__(self):
        self.children: list[Optional[TrieV1]] = [None] * 26
        self.is_end: bool = False

    def insert(self, word: str) -> None:
        # Time complexity: O(n).
        node = self
        for char in word:
            idx = ord(char) - NIL
            cur = node.children[idx]
            if cur is None:
                cur = TrieV1()
                node.children[idx] = cur
            node = cur
        node.is_end = True

    def search(self, word: str) -> bool:
        # Time complexity: O(n). Space complexity: O(1).
        node = self
        for char in word:
            cur = node.children[ord(char) - NIL]
            if cur is None:
                return False
            node = cur
        return node.is_end

    def startsWith(self, prefix: str) -> bool:
        # Time complexity: O(n). Space complexity: O(1).
        node = self
        for char in prefix:
            cur = node.children[ord(char) - NIL]
            if cur is None:
                return False
            node = cur
        return True
