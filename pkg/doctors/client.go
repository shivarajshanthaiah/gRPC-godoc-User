package doctors

import (
	"log"

	docpb "github.com/shivaraj-shanthaiah/godoc-user/pkg/doctors/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial() (docpb.DoctorServicesClient, error) {
	grpc, err := grpc.Dial(":8084", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error dialing to grpc client :8084")
		return nil, err
	}
	log.Printf("successfully connected to client :8084")
	return docpb.NewDoctorServicesClient(grpc), nil
}
