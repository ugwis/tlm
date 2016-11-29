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

func jobstask(jobs []Job, origin map[List][]int64) (ret map[int64]Change) {
	listids := make(map[List]int64)

	result := make(map[List][]int64, len(origin))
	for k, v := range origin {
		result[k] = v
	}

	for _, v := range jobs {
		l1 := result[v.List1]
		l2 := result[v.List2]
		switch v.Operator {
		case "*":
			result[v.Listresult] = intersect(l1, l2)
		case "+":
			result[v.Listresult] = union(l1, l2)
		case "-":
			result[v.Listresult] = except(l1, l2)
		}

		if v.Config.Saveflag {
			addval := Change{
				AddList: except(result[v.Listresult], origin[v.Listresult]),
				DelList: except(origin[v.Listresult], result[v.Listresult])}

			listid, ok := listids[v.Listresult]
			if !ok {
				listid = v.Listresult.ListID
				if listid == 0 {
					//listid=createlist()
				}
				listids[v.Listresult] = listid
			}
			ret[listid] = addval
		}
	}
	return
}
