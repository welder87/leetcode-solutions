from lc_705_design_hashset.lc_705_design_hashset import MyHashSet


def test_v0(subtests):
    st = MyHashSet()
    with subtests.test(i=0):
        assert st._capacity == 0
        assert st._size == 4
        assert st._chains == [None] * 4
    with subtests.test(i=1):
        assert st.add(1) is None
        assert len([item for item in st._chains if item is None]) == 3
        assert st._capacity == 1
        assert st._size == 4
        assert st._chains[1].val == 1
        assert st._chains[1].next is None
    with subtests.test(i=2):
        assert st.add(2) is None
        assert len([item for item in st._chains if item is None]) == 2
        assert st._capacity == 2
        assert st._size == 4
        assert st._chains[1].val == 1
        assert st._chains[1].next is None
        assert st._chains[2].val == 2
        assert st._chains[2].next is None
    with subtests.test(i=3):
        assert st.contains(1) is True
    with subtests.test(i=5):
        assert st.contains(5) is False
    with subtests.test(i=6):
        assert st.add(2) is None
        assert len([item for item in st._chains if item is None]) == 2
        assert st._capacity == 2
        assert st._size == 4
        assert st._chains[1].val == 1
        assert st._chains[1].next is None
        assert st._chains[2].val == 2
        assert st._chains[2].next is None
    with subtests.test(i=7):
        assert st.add(5) is None
        assert st._capacity == 3
        assert st._size == 4
        assert len([item for item in st._chains if item is None]) == 2
        assert st._chains[1].val == 1
        assert st._chains[1].next is not None
        assert st._chains[1].next.val == 5
        assert st._chains[1].next.next is None
        assert st._chains[2].val == 2
        assert st._chains[2].next is None
    with subtests.test(i=8):
        assert st.remove(1) is None
        assert st._capacity == 2
        assert st._size == 4
        assert len([item for item in st._chains if item is None]) == 2
        assert st._chains[1].val == 5
        assert st._chains[1].next is None
        assert st._chains[2].val == 2
        assert st._chains[2].next is None
    with subtests.test(i=9):
        assert st.remove(5) is None
        assert st._capacity == 1
        assert st._size == 4
        assert len([item for item in st._chains if item is None]) == 3
        assert st._chains[2].val == 2
        assert st._chains[2].next is None
