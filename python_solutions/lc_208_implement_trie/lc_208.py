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


class TrieV2:
    """
    Solution: https://algo.monster/liteproblems/208

    A Trie (prefix tree) data structure implementation for efficient string operations.
    Each node contains an array of 26 children (for lowercase English letters a-z)
    and a boolean flag to mark word endings.
    """

    def __init__(self):
        """Initialize the Trie node with empty children array and end-of-word flag."""
        self.children = [None] * 26  # Array to store 26 lowercase letters
        self.is_end_of_word = (
            False  # Flag to mark if this node represents end of a word
        )

    def insert(self, word: str) -> None:
        """
        Insert a word into the Trie.
        Time Complexity: O(n) where n is the length of the word
        Space Complexity: O(n) in worst case when creating new nodes

        Args:
            word: The word to be inserted into the Trie
        """
        current_node = self

        for char in word:
            # Calculate the index for current character (0-25 for a-z)
            char_index = ord(char) - ord("a")

            # Create new node if path doesn't exist
            if current_node.children[char_index] is None:
                current_node.children[char_index] = TrieV2()

            # Move to the child node
            current_node = current_node.children[char_index]

        # Mark the last node as end of word
        current_node.is_end_of_word = True

    def search(self, word: str) -> bool:
        """
        Search for a complete word in the Trie.

        Args:
            word: The word to search for

        Returns:
            True if the word exists in the Trie, False otherwise
        """
        final_node = self._search_prefix(word)
        return final_node is not None and final_node.is_end_of_word

    def startsWith(self, prefix: str) -> bool:
        """
        Check if any word in the Trie starts with the given prefix.

        Args:
            prefix: The prefix to search for

        Returns:
            True if any word starts with the prefix, False otherwise
        """
        prefix_node = self._search_prefix(prefix)
        return prefix_node is not None

    def _search_prefix(self, prefix: str) -> "TrieV2":
        """
        Helper method to traverse the Trie following the given prefix.

        Args:
            prefix: The prefix string to traverse

        Returns:
            The Trie node at the end of the prefix path, or None if path doesn't exist
        """
        current_node = self

        for char in prefix:
            # Calculate the index for current character
            char_index = ord(char) - ord("a")

            # Return None if path doesn't exist
            if current_node.children[char_index] is None:
                return None

            # Move to the child node
            current_node = current_node.children[char_index]

        return current_node
