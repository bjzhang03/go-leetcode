package easy_0021

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode
	if l1 == nil && l2 == nil {
		result = nil
	} else if l1 == nil && l2 != nil {
		result = l2
	} else if l1 != nil && l2 == nil {
		result = l1
	} else {
		if l1.Val < l2.Val {
			result = l1
			l1 = l1.Next
		} else {
			result = l2
			l2 = l2.Next
		}
		result.Next = nil
		tmp := result
		/*比较数据的大小*/
		for l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				tmp.Next = l1
				l1 = l1.Next
				tmp = tmp.Next
				tmp.Next = nil
			} else {
				tmp.Next = l2
				l2 = l2.Next
				tmp = tmp.Next
				tmp.Next = nil
			}
		}
		/*添加上最后面的数据*/
		if l1 != nil {
			tmp.Next = l1
		}
		if l2 != nil {
			tmp.Next = l2
		}
	}

	return result

}
