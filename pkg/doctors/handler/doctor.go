package handler

import (
	"context"
	"log"

	docpb "github.com/shivaraj-shanthaiah/godoc-user/pkg/doctors/pb"
	userpb "github.com/shivaraj-shanthaiah/godoc-user/pkg/pb"
)

func FetchDoctorByIDHandler(client docpb.DoctorServicesClient, p *userpb.DoctorID) (*docpb.DoctorModel, error) {
	ctx := context.Background()

	response, err := client.FetchDoctorByID(ctx, &docpb.DoctorID{Id: p.Id})
	if err != nil {
		log.Printf("error while fetching doctor by name")
		return nil, err
	}
	return response, nil
}

func FetchDoctorByNameHandler(client docpb.DoctorServicesClient, p *userpb.DoctorName) (*docpb.DoctorModel, error) {
	ctx := context.Background()

	response, err := client.FetchDoctorByName(ctx, &docpb.DoctorName{Name: p.Name})
	if err != nil {
		log.Printf("error while fetching doctor by id")
		return nil, err
	}
	return response, nil
}

func FetchAllDoctorsHandler(client docpb.DoctorServicesClient, p *userpb.UNoParam) (*docpb.DoctorList, error) {
	ctx := context.Background()

	response, err := client.FetchAllDoctors(ctx, &docpb.NoParam{})
	if err != nil {
		log.Printf("error while fetching doctors")
		return nil, err
	}
	return response, nil
}
