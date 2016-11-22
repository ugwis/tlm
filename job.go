package main

func intersect(l1, l2 []int64) (ret []int64) {
	i, j := 0, 0
	for i < len(l1) && j < len(l2) {
		if l1[i] == l2[j] {
			ret = append(ret, l1[i])
			i++
			j++
		} else if l1[i] < l2[j] {
			i++
		} else {
			j++
		}
	}
	return
}

func union(l1, l2 []int64) (ret []int64) {
	i, j := 0, 0
	for i < len(l1) && j < len(l2) {
		if l1[i] == l2[j] {
			ret = append(ret, l1[i])
			i++
			j++
		} else if l1[i] < l2[j] {
			ret = append(ret, l1[i])
			i++
		} else {
			ret = append(ret, l2[i])
			j++
		}
	}
	ret = append(ret, l1[i:]...)
	ret = append(ret, l2[j:]...)
	return
}

func except(l1, l2 []int64) (ret []int64) {
	i, j := 0, 0
	for i < len(l1) && j < len(l2) {
		if l1[i] == l2[j] {
			i++
			j++
		} else if l1[i] < l2[j] {
			ret = append(ret, l1[i])
			i++
		} else {
			j++
		}
	}
	ret = append(ret, l1[i:]...)
	return
}

func jobstask(jobs []Job, ret map[List][]int64) map[List][]int64 {
	for _, v := range jobs {
		l1 := ret[v.List1]
		l2 := ret[v.List2]
		switch v.Operator {
		case "*":
			ret[v.Listresult] = intersect(l1, l2)
		case "+":
			ret[v.Listresult] = union(l1, l2)
		case "-":
			ret[v.Listresult] = except(l1, l2)
		}
	}
	return ret
}
