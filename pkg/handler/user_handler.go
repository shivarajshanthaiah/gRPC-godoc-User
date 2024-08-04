package handler

import (
	"context"
	"fmt"

	pb "github.com/shivaraj-shanthaiah/godoc-user/pkg/pb"
	service_interfaces "github.com/shivaraj-shanthaiah/godoc-user/pkg/services/interfaces"
)

type UserHandler struct {
	pb.UserServiceServer
	svc service_interfaces.UserServiceInterface
}

func NewUserHandler(svc service_interfaces.UserServiceInterface) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (u *UserHandler) UserSignup(ctx context.Context, p *pb.SignupRequest) (*pb.CommonResponse, error) {
	result, err := u.svc.SignupService(p)
	if err != nil {
		return &pb.CommonResponse{
			Status:  "Failed",
			Error:   "User signup Error",
			Message: "",
		}, err
	}

	msg := fmt.Sprintf("User Created %v", result)
	return &pb.CommonResponse{
		Status:  "Success",
		Error:   "",
		Message: msg,
	}, err
}

func (u *UserHandler) UserLogin(ctx context.Context, p *pb.LoginRequest) (*pb.CommonResponse, error) {
	result, err := u.svc.LoginService(p)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (u *UserHandler) UserProfile(ctx context.Context, p *pb.UserID) (*pb.Profile, error) {
	result, err := u.svc.ProfileService(p)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (u *UserHandler) UserProfileUpdate(ctx context.Context, p *pb.ProfileUpdate) (*pb.Profile, error) {
	result, err := u.svc.UpdateProfileService(p)
	if err != nil {
		return result, err
	}
	return result, nil
}
