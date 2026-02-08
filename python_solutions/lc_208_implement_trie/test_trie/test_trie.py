import pytest

from lc_208_implement_trie.lc_208 import Trie
from lc_208_implement_trie.test_trie.helpers import Cases


def test_empty_trie(empty_trie):
    assert not empty_trie.root.is_end
    assert len(empty_trie.root.children) == 26


class TestCases(Cases):
    pass


@pytest.fixture
def trie():
    return Trie()
