package User

import iuser "gomock/learn/IUser"

type User struct {
	Iuser iuser.IuserInterface
}

func (IU *User) AddCell() {
	IU.Iuser.AddCell(1, "sample string", true, 1)
}

func (IU *User) UpdateCell() {
	IU.Iuser.UpdateCell(1, "sample string1")
}
