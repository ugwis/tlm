package user

type UserIDs []UserID

func (u *UserIDs) New(i []int64) {
	for _, v := range i {
		var one UserID
		one.New(v)
		*u = append(*u, one)
	}
}

func (u UserIDs) Len() int {
	return len(u)
}

func (u UserIDs) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func (u UserIDs) Less(i, j int) bool {
	return u[i] < u[j]
}

type UserID int64

func (u *UserID) New(i int64) {
	*u = UserID(i)
}
