import pytest


@pytest.fixture
def trie():
    raise NotImplementedError


@pytest.fixture
def empty_trie(trie):
    return trie


@pytest.fixture
def basic_trie(trie):
    words = ["apple", "app", "banana", "bat", "cat"]
    for word in words:
        trie.insert(word)
    return trie


@pytest.fixture
def overlapping_trie(trie):
    words = ["a", "ab", "abc", "abcd", "abcde"]
    for word in words:
        trie.insert(word)
    return trie


@pytest.fixture
def large_trie(trie):
    for i in range(100):
        trie.insert(f"word{i}")
    return trie
