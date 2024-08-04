package services

import (
	"github.com/shivaraj-shanthaiah/godoc-user/pkg/models"
	pb "github.com/shivaraj-shanthaiah/godoc-user/pkg/pb"
)

func (u *UserService) ProfileService(userpb *pb.UserID) (*pb.Profile, error) {
	user, err := u.repo.FindUserByID(uint(userpb.Id))
	if err != nil {
		return nil, err
	}

	response := &pb.Profile{
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Username:  user.UserName,
		Dob:       user.DoB,
		Gender:    user.Gender,
		Email:     user.Email,
		Phone:     user.Phone,
	}
	return response, nil
}

func (u *UserService) UpdateProfileService(userpb *pb.ProfileUpdate) (*pb.Profile, error) {
	user, err := u.repo.FindUserByID(uint(userpb.Userid))
	if err != nil {
		return nil, err
	}
	user = &models.User{
		UserID:    uint(user.UserID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
		DoB:       user.DoB,
		Gender:    user.Gender,
		Email:     user.Email,
		Phone:     user.Phone,
		Password:  user.Password,
	}

	err = u.repo.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return &pb.Profile{
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Username:  user.UserName,
		Dob:       user.DoB,
		Gender:    user.Gender,
		Email:     user.Email,
		Phone:     user.Phone,
	}, nil
}
