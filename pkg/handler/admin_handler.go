package handler

import (
	"context"

	pb "github.com/shivaraj-shanthaiah/godoc-user/pkg/pb"
)

func (u *UserHandler) FindAllUsers(ctx context.Context, p *pb.UNoParam) (*pb.Users, error) {
	result, err := u.svc.FindAllUsersService(p)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *UserHandler) EditUser(ctx context.Context, p *pb.SignupRequest) (*pb.CommonResponse, error) {
	result, err := u.svc.EditUserService(p)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *UserHandler) DeleteUser(ctx context.Context, p *pb.UserID) (*pb.CommonResponse, error) {
	result, err := u.svc.DeleteUserService(p)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *UserHandler) FindUserByID(ctx context.Context, p *pb.UserID) (*pb.Profile, error) {
	result, err := u.svc.FindUserByIDService(p)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *UserHandler) FindUserByEmail(ctx context.Context, p *pb.Email) (*pb.Profile, error) {
	result, err := u.svc.FindUserByEmailService(p)
	if err != nil {
		return nil, err
	}
	return result, nil
}
