package user

func (one UserIDs) Intersect(another UserIDs) UserIDs {
	var ret UserIDs
	i, j := 0, 0
	for i < len(one) && j < len(another) {
		if one[i] == another[j] {
			ret = append(ret, one[i])
			i++
			j++
		} else if one[i] < another[j] {
			i++
		} else {
			j++
		}
	}
	return ret
}

func (one UserIDs) Union(another UserIDs) UserIDs {
	var ret UserIDs
	i, j := 0, 0
	for i < len(one) && j < len(another) {
		if one[i] == another[j] {
			ret = append(ret, one[i])
			i++
			j++
		} else if one[i] < another[j] {
			ret = append(ret, one[i])
			i++
		} else {
			ret = append(ret, another[j])
			j++
		}
	}
	ret = append(ret, one[i:]...)
	ret = append(ret, another[j:]...)

	return ret
}

func (one UserIDs) Except(another UserIDs) UserIDs {
	var ret UserIDs
	i, j := 0, 0
	for i < len(one) && j < len(another) {
		if one[i] == another[j] {
			i++
			j++
		} else if one[i] < another[j] {
			ret = append(ret, one[i])
			i++
		} else {
			j++
		}
	}
	ret = append(ret, one[i:]...)
	return ret
}
