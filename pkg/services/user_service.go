package services

import (
	"fmt"

	docpb "github.com/shivaraj-shanthaiah/godoc-user/pkg/doctors/pb"
	"github.com/shivaraj-shanthaiah/godoc-user/pkg/models"
	pb "github.com/shivaraj-shanthaiah/godoc-user/pkg/pb"
	repo_interfaces "github.com/shivaraj-shanthaiah/godoc-user/pkg/repo/interfaces"
	service_interfaces "github.com/shivaraj-shanthaiah/godoc-user/pkg/services/interfaces"
)

type UserService struct {
	repo repo_interfaces.UserRepoInterface
	doc  docpb.DoctorServicesClient
}

func (u *UserService) SignupService(userpb *pb.SignupRequest) (*pb.CommonResponse, error) {
	hashedPassword, err := u.HashPassword(userpb.Password)
	if err != nil {
		return &pb.CommonResponse{
			Status:  "Failed",
			Error:   err.Error(),
			Message: "Password hashing error",
		}, err
	}
	user := &models.User{
		FirstName: userpb.Firstname,
		LastName:  userpb.Lastname,
		UserName:  userpb.Firstname + " " + userpb.Lastname,
		DoB:       userpb.Dob,
		Gender:    userpb.Gender,
		Email:     userpb.Email,
		Phone:     userpb.Phone,
		Role:      userpb.Role,
		Password:  hashedPassword,
	}
	err = u.repo.CreateUser(user)
	if err != nil {
		return &pb.CommonResponse{
			Status:  "Failed",
			Error:   err.Error(),
			Message: "User Signup error",
		}, err
	}
	return &pb.CommonResponse{
		Status:  "Succes",
		Error:   "",
		Message: "User Singnedup successfully",
	}, err

}

func (u *UserService) LoginService(userpb *pb.LoginRequest) (*pb.CommonResponse, error) {

	user, err := u.repo.FindUserByEmail(userpb.Email)

	if err != nil {
		return &pb.CommonResponse{
			Status:  "Failed",
			Error:   err.Error(),
			Message: "Error Finding user by email",
		}, err
	}
	if err := u.CheckPassword(userpb.Password, user.Password); err != nil {

		return &pb.CommonResponse{
			Status:  "Failed",
			Error:   err.Error(),
			Message: "Invalid Password",
		}, err
	}

	// if user.Role != "user" {
	// 	return &pb.CommonResponse{
	// 		Status:  "Failed",
	// 		Error:   "",
	// 		Message: "Invalid Role",
	// 	}, err
	// }

	token, err := u.GenerateToken(userpb.Email)
	if err != nil {
		fmt.Println("This is log in", err)
		return &pb.CommonResponse{
			Status:  "Failed",
			Error:   err.Error(),
			Message: "Error while Generating Token",
		}, err
	}

	return &pb.CommonResponse{
		Status:  "Succes",
		Error:   "",
		Message: token,
	}, nil
}

func (u *UserService) FindAllUsersService(p *pb.UNoParam) (*pb.Users, error) {
	result, err := u.repo.FindAllUsers()
	if err != nil {
		return nil, err
	}
	var users []*pb.Profile
	for _, i := range result {
		users = append(users, &pb.Profile{
			Firstname: i.FirstName,
			Lastname:  i.LastName,
			Username:  i.LastName,
			Email:     i.Email,
			Dob:       i.DoB,
			Gender:    i.Gender,
			Phone:     i.Phone,
		})
	}
	response := &pb.Users{
		Userlist: users,
	}
	return response, nil
}

func (u *UserService) EditUserService(p *pb.SignupRequest) (*pb.CommonResponse, error) {
	user, err := u.repo.FindUserByEmail(p.Email)
	if err != nil {
		return nil, err
	}
	user = &models.User{
		UserID:    uint(user.UserID),
		FirstName: p.Firstname,
		LastName:  p.Lastname,
		UserName:  user.UserName,
		DoB:       p.Dob,
		Gender:    p.Gender,
		Email:     user.Email,
		Phone:     p.Phone,
		Password:  user.Password,
	}

	err = u.repo.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return &pb.CommonResponse{
		Status:  "Success",
		Error:   "",
		Message: "User details have been successfully edited",
	}, nil
}

func (u *UserService) FindUserByEmailService(p *pb.Email) (*pb.Profile, error) {
	user, err := u.repo.FindUserByEmail(p.Email)
	if err != nil {
		return nil, err
	}
	return &pb.Profile{
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Username:  user.UserName,
		Gender:    user.Gender,
		Dob:       user.DoB,
		Email:     user.Email,
		Phone:     user.Phone,
	}, nil
}

func (u *UserService) FindUserByIDService(p *pb.UserID) (*pb.Profile, error) {
	user, err := u.repo.FindUserByID(uint(p.Id))
	if err != nil {
		return nil, err
	}
	return &pb.Profile{
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Username:  user.UserName,
		Gender:    user.Gender,
		Dob:       user.DoB,
		Email:     user.Email,
		Phone:     user.Phone,
	}, nil
}

func (u *UserService) DeleteUserService(p *pb.UserID) (*pb.CommonResponse, error) {
	err := u.repo.DeleteUser(uint(p.Id))
	if err != nil {
		return nil, err
	}
	return &pb.CommonResponse{
		Status:  "Success",
		Error:   "",
		Message: "User deleted succesfully",
	}, nil
}

func NewUserService(repo repo_interfaces.UserRepoInterface, doc docpb.DoctorServicesClient) service_interfaces.UserServiceInterface {
	return &UserService{
		repo: repo,
		doc:  doc,
	}
}
