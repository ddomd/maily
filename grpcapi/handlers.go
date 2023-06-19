package grpcapi

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc/peer"
	"github.com/ddomd/maily/internal/mdb"
	pb "github.com/ddomd/maily/proto"
)


func (cfg *Server) GetEmail(ctx context.Context, req *pb.GetEmailRequest) (*pb.EmailResponse, error) {
	email, err := cfg.DB.GetEmail(req.EmailAddress)
	if err != nil {
		return nil, err
	}

	res := dbToPbEmail(email)

	log.Printf("GRPC: Retrieved email entry(id:%d), client Addr: %s\n", email.ID, getIpFromCtx(ctx))
	return &pb.EmailResponse{Email: res}, nil 
}

func (cfg *Server) GetBatchEmails(ctx context.Context, req *pb.GetEmailBatchRequest) (*pb.EmailBatchResponse, error) {
	params := mdb.BatchParams{
		Offset: int(req.Offset), 
		Limit: int(req.Limit),
	}

	emails, err := cfg.DB.GetBatchEmails(params)
	if err != nil {
		return nil, err
	}

	res := dbToPbEmails(emails)
	
	log.Printf("GRPC: Retrieved email batch(id:%d-id:%d), client Addr:%s\n", emails[0].ID, emails[len(emails)-1].ID, getIpFromCtx(ctx))
	return &pb.EmailBatchResponse{Emails: res}, nil
}

func (cfg *Server) CreateEmail(ctx context.Context, req *pb.CreateEmailRequest) (*pb.EmailResponse, error) {
	email, err := cfg.DB.CreateEmail(req.EmailAddress)
	if err != nil {
		return nil, err
	}

	res := dbToPbEmail(email)

	log.Printf("GRPC: Created new email entry(id:%d), client Addr: %s\n", email.ID, getIpFromCtx(ctx))
	return &pb.EmailResponse{Email: res}, nil
}

func (cfg *Server) UpdateEmail(ctx context.Context, req *pb.UpdateEmailRequest) (*pb.EmailResponse, error) {
	email, err := cfg.DB.UpdateEmail(req.EmailAddress, req.OptOut)
	if err != nil {
		return nil, err
	}

	res := dbToPbEmail(email)

	log.Printf("GRPC: Updated email entry(id:%d), client Addr: %s\n", email.ID, getIpFromCtx(ctx))
	return &pb.EmailResponse{Email: res}, nil
}

func (cfg *Server) DeleteEmail(ctx context.Context, req *pb.DeleteEmailRequest) (*pb.EmailResponse, error) {
	email, err := cfg.DB.DeleteEmail(req.EmailAddress)
	if err != nil {
		return nil, err
	}

	res := dbToPbEmail(email)

	log.Printf("GRPC: Removed email entry(id:%d), client Addr: %s\n", email.ID, getIpFromCtx(ctx))
	return &pb.EmailResponse{Email: res}, nil
}

func getIpFromCtx(ctx context.Context) string {
	p, _ := peer.FromContext(ctx)
	return p.Addr.String()
}

func pbToDbEmail(email *pb.Email) mdb.Email {
	return mdb.Email{
		ID: email.Id,
		Address: email.Email,
		ConfirmedAt: time.Unix(email.ConfirmedAt, 0),
		OptOut: email.OptOut,
	}
}


func dbToPbEmail(email mdb.Email) *pb.Email {
	return &pb.Email{
		Id: email.ID,
		Email: email.Address,
		ConfirmedAt: email.ConfirmedAt.Unix(),
		OptOut: email.OptOut,
	}
}

func dbToPbEmails(emails []mdb.Email) []*pb.Email {
	pbEmails := make([]*pb.Email, 0, len(emails))

	for i := range emails {
		pbEmails = append(pbEmails, dbToPbEmail(emails[i]))
	}

	return pbEmails
}