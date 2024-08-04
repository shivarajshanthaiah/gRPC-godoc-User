package services

import (
	doc "github.com/shivaraj-shanthaiah/godoc-user/pkg/doctors/handler"
	userpb "github.com/shivaraj-shanthaiah/godoc-user/pkg/pb"
)

func (u *UserService) FetchDoctorByIDService(p *userpb.DoctorID) (*userpb.DoctorModel, error) {
	result, err := doc.FetchDoctorByIDHandler(u.doc, p)
	if err != nil {
		return nil, err
	}

	return &userpb.DoctorModel{
		DoctorName: result.DoctorName,
		Age:        result.Age,
		Speciality: result.Speciality,
		Gender:     result.Gender,
		Email:      result.Email,
	}, nil
}

func (u *UserService) FetchDoctorByNameService(p *userpb.DoctorName) (*userpb.DoctorModel, error) {
	result, err := doc.FetchDoctorByNameHandler(u.doc, p)
	if err != nil {
		return nil, err
	}

	return &userpb.DoctorModel{
		DoctorName: result.DoctorName,
		Age:        result.Age,
		Speciality: result.Speciality,
		Gender:     result.Gender,
		Email:      result.Email,
	}, nil
}

func (u *UserService) FetchAllDoctorsService(p *userpb.UNoParam) (*userpb.DoctorList, error) {
	result, err := doc.FetchAllDoctorsHandler(u.doc, p)
	if err != nil {
		return nil, err
	}
	var doctors []*userpb.DoctorModel
	for _, i := range result.Doctors {
		doctors = append(doctors, &userpb.DoctorModel{
			DoctorName: i.DoctorName,
			Age:        i.Age,
			Speciality: i.Speciality,
			Gender:     i.Speciality,
			Email:      i.Email,
		})
	}

	return &userpb.DoctorList{
		Doctors: doctors,
	}, nil
}
