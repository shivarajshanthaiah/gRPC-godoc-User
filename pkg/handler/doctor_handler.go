package handler

import (
	"context"
	"log"

	userpb "github.com/shivaraj-shanthaiah/godoc-user/pkg/pb"
)

func (u *UserHandler) UserFetchDoctorByID(ctx context.Context, p *userpb.DoctorID) (*userpb.DoctorModel, error) {
	result, err := u.svc.FetchDoctorByIDService(p)
	if err != nil {
		log.Println("error while fetching dotor by id")
		return nil, err
	}
	return result, nil
}

func (u *UserHandler) UserFetchDoctorByName(ctx context.Context, p *userpb.DoctorName) (*userpb.DoctorModel, error) {
	result, err := u.svc.FetchDoctorByNameService(p)
	if err != nil {
		log.Println("error while fetching doctors by name")
		return nil, err
	}
	return result, nil
}

func (u *UserHandler) UserFetchAllDoctors(ctx context.Context, p *userpb.UNoParam) (*userpb.DoctorList, error) {
	result, err := u.svc.FetchAllDoctorsService(p)
	if err != nil {
		log.Println("error while fetching all doctors")
		return nil, err
	}
	return result, nil
}
