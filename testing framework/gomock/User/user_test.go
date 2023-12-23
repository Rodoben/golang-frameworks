package User_test

import (
	"gomock/learn/User"
	"gomock/learn/mocks"

	"testing"

	"github.com/golang/mock/gomock"
)

func Test_AddCell(t *testing.T) {
	mockctrl := gomock.NewController(t)
	defer mockctrl.Finish()

	mockUser := mocks.NewMockIuserInterface(mockctrl)
	testUser := &User.User{Iuser: mockUser}
	mockUser.EXPECT().AddCell(1, "sample string", true, 1).Return(nil).Times(1)
	testUser.AddCell()

}

func Test_UpdateCell(t *testing.T) {
	mockctrl := gomock.NewController(t)
	defer mockctrl.Finish()

	mockUser := mocks.NewMockIuserInterface(mockctrl)
	testUser := &User.User{Iuser: mockUser}
	mockUser.EXPECT().UpdateCell(1, "sample string1").Return(nil).Times(1)
	testUser.UpdateCell()

}
