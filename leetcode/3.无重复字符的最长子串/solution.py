# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        return self.addTwoNumbers_(l1, l2, 0)

    def addTwoNumbers_(self, l1: ListNode, l2: ListNode, last_promotion_val) -> ListNode:
        if l1 is None and l2 is None and last_promotion_val == 0:
            return None

        l1_va = 0
        l2_va = 0

        if l1 is not None:
            l1_va = l1.val

        if l2 is not None:
            l2_va = l2.val

        sum_val = l1_va + l2_va + last_promotion_val

        new_node = ListNode()
        new_node.val = sum_val % 10
        promotion_val = int(sum_val / 10)

        l1_next = None
        l2_next = None
        if l1 is not None and l1.next is not None:
            l1_next = l1.next
        if l2 is not None and l2.next is not None:
            l2_next = l2.next

        new_node.next = self.addTwoNumbers_(l1_next, l2_next, promotion_val)
        return new_node


def gen_list(arr):
    return gen_list_(arr, 0, len(arr) - 1)


def gen_list_(arr, i, last_index):
    node = ListNode()
    node.val = arr[i]
    if i < last_index:
        node.next = gen_list_(arr, i + 1, last_index)
    return node


def print_list(l: ListNode):
    if l is not None:
        print(l.val, end="")
        if l.next is not None:
            print("->", end="")
            print_list(l.next)


if __name__ == "__main__":
    l1 = gen_list([1, 2, 3, 4, 5, 6])
    l2 = gen_list([1, 2, 3, 4, 0, 0, 1])
    print_list(Solution().addTwoNumbers(l1, l2))
