package service_interfaces

import pb "github.com/shivaraj-shanthaiah/godoc-user/pkg/pb"

type UserServiceInterface interface {
	SignupService(userpb *pb.SignupRequest) (*pb.CommonResponse, error)
	LoginService(userpb *pb.LoginRequest) (*pb.CommonResponse, error)
	ProfileService(userpb *pb.UserID) (*pb.Profile, error)
	UpdateProfileService(userpb *pb.ProfileUpdate) (*pb.Profile, error)

	FindAllUsersService(p *pb.UNoParam) (*pb.Users, error)
	EditUserService(p *pb.SignupRequest) (*pb.CommonResponse, error)
	FindUserByEmailService(p *pb.Email) (*pb.Profile, error)
	FindUserByIDService(p *pb.UserID) (*pb.Profile, error)
	DeleteUserService(p *pb.UserID) (*pb.CommonResponse, error)

	FetchDoctorByIDService(p *pb.DoctorID) (*pb.DoctorModel, error)
	FetchDoctorByNameService(p *pb.DoctorName) (*pb.DoctorModel, error)
	FetchAllDoctorsService(p *pb.UNoParam) (*pb.DoctorList, error)
}
