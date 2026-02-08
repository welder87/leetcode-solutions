import pytest


class Cases:
    def test_insert_single_word(self, empty_trie):
        empty_trie.insert("hello")
        assert empty_trie.search("hello")

    def test_insert_multiple_words(self, empty_trie):
        words = ["apple", "banana", "cherry"]
        for word in words:
            empty_trie.insert(word)

        for word in words:
            assert empty_trie.search(word)

    def test_insert_duplicate_word(self, empty_trie):
        empty_trie.insert("test")
        empty_trie.insert("test")
        assert empty_trie.search("test")

    @pytest.mark.parametrize(
        "word",
        (
            pytest.param("a", id="Single character"),
            pytest.param("antidisestablishmentarianism", id="Long word"),
        ),
    )
    def test_insert_various_words(self, empty_trie, word):
        empty_trie.insert(word)
        assert empty_trie.search(word)

    @pytest.mark.parametrize(
        ("word", "ans"),
        (
            ("apple", True),
            ("app", True),
            ("banana", True),
            ("bat", True),
            ("cat", True),
            ("dog", False),  # Not inserted
            ("ap", False),  # Prefix but not complete word
        ),
    )
    def test_search_basic_trie(self, basic_trie, word, ans):
        assert basic_trie.search(word) is ans

    def test_search_empty_trie(self, empty_trie):
        assert not empty_trie.search("anything")

    def test_search_overlapping_words(self, overlapping_trie):
        assert overlapping_trie.search("a")
        assert overlapping_trie.search("ab")
        assert overlapping_trie.search("abc")
        assert overlapping_trie.search("abcd")
        assert overlapping_trie.search("abcde")
        assert not overlapping_trie.search("abcdef")

    @pytest.mark.parametrize(
        ("prefix", "ans"),
        (
            ("app", True),
            ("apple", True),
            ("a", True),
            ("b", True),
            ("ba", True),
            ("ban", True),
            ("c", True),
            ("cat", True),
            ("d", False),  # No words starting with d
            ("appl", True),  # Prefix of apple
            ("bana", True),  # Prefix of banana
            ("catt", False),  # Not a prefix
        ),
    )
    def test_starts_with_basic(self, basic_trie, prefix, ans):
        assert basic_trie.startsWith(prefix) is ans

    def test_starts_with_empty_trie(self, empty_trie):
        assert not empty_trie.startsWith("any")

    def test_starts_with_overlapping(self, overlapping_trie):
        """Test startsWith with overlapping words"""
        assert overlapping_trie.startsWith("")
        assert overlapping_trie.startsWith("a")
        assert overlapping_trie.startsWith("ab")
        assert overlapping_trie.startsWith("abc")
        assert overlapping_trie.startsWith("abcd")
        assert overlapping_trie.startsWith("abcde")
        assert not overlapping_trie.startsWith("abcdef")
        assert not overlapping_trie.startsWith("b")
