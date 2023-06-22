package grpcapi

import (
	"context"
	"log"

	"github.com/ddomd/maily/internal/mdb"
	pb "github.com/ddomd/maily/proto"
)


func (s *Server) GetEmail(ctx context.Context, req *pb.GetEmailRequest) (*pb.EmailResponse, error) {
	email, err := s.DB.GetEmail(req.Id)
	if err != nil {
		return nil, err
	}

	res := dbToPbEmail(email)

	log.Printf("GRPC: Retrieved email entry(id:%d), client Addr: %s\n", email.ID, getIpFromCtx(ctx))
	return &pb.EmailResponse{Email: res}, nil 
}

func (s *Server) GetAll(ctx context.Context, req *pb.GetAllRequest) (*pb.EmailBatchResponse, error) {
	emails, err := s.DB.GetAllEmails()
	if err != nil {
		return nil, err
	}

	res := dbToPbEmails(emails)
	
	log.Printf("GRPC: Retrieved all email entries, client Addr:%s\n", getIpFromCtx(ctx))
	return &pb.EmailBatchResponse{Emails: res}, nil
}

func (s *Server) GetAllSubscribed(ctx context.Context, req *pb.GetAllRequest) (*pb.EmailBatchResponse, error) {
	emails, err := s.DB.GetAllSubscribed()
	if err != nil {
		return nil, err
	}

	res := dbToPbEmails(emails)
	
	log.Printf("GRPC: Retrieved all subscribed email entries, client Addr:%s\n", getIpFromCtx(ctx))
	return &pb.EmailBatchResponse{Emails: res}, nil
}

func (s *Server) GetBatch(ctx context.Context, req *pb.GetBatchRequest) (*pb.EmailBatchResponse, error) {
	params := mdb.BatchParams{
		Offset: int(req.Offset), 
		Limit: int(req.Limit),
	}

	emails, err := s.DB.GetBatchEmails(params)
	if err != nil {
		return nil, err
	}

	res := dbToPbEmails(emails)
	
	log.Printf("GRPC: Retrieved email batch(id:%d-id:%d), client Addr:%s\n", emails[0].ID, emails[len(emails)-1].ID, getIpFromCtx(ctx))
	return &pb.EmailBatchResponse{Emails: res}, nil
}

//
func (s *Server) GetBatchSubscribed(ctx context.Context, req *pb.GetBatchRequest) (*pb.EmailBatchResponse, error) {
	params := mdb.BatchParams{
		Offset: int(req.Offset), 
		Limit: int(req.Limit),
	}

	emails, err := s.DB.GetBatchSubscribed(params)
	if err != nil {
		return nil, err
	}

	res := dbToPbEmails(emails)
	
	log.Printf("GRPC: Retrieved subscribed email batch(id:%d-id:%d), client Addr:%s\n", emails[0].ID, emails[len(emails)-1].ID, getIpFromCtx(ctx))
	return &pb.EmailBatchResponse{Emails: res}, nil
}

//
func (s *Server) CreateEmail(ctx context.Context, req *pb.CreateEmailRequest) (*pb.EmailResponse, error) {
	email, err := s.DB.CreateEmail(req.EmailAddress)
	if err != nil {
		return nil, err
	}

	res := dbToPbEmail(email)

	log.Printf("GRPC: Created new email entry(id:%d), client Addr: %s\n", email.ID, getIpFromCtx(ctx))
	return &pb.EmailResponse{Email: res}, nil
}

func (s *Server) UpdateEmail(ctx context.Context, req *pb.UpdateEmailRequest) (*pb.EmailResponse, error) {
	email, err := s.DB.UpdateEmail(req.Id, req.OptOut)
	if err != nil {
		return nil, err
	}

	res := dbToPbEmail(email)

	log.Printf("GRPC: Updated email entry(id:%d), client Addr: %s\n", email.ID, getIpFromCtx(ctx))
	return &pb.EmailResponse{Email: res}, nil
}

func (s *Server) DeleteEmail(ctx context.Context, req *pb.DeleteEmailRequest) (*pb.EmailResponse, error) {
	email, err := s.DB.DeleteEmail(req.Id)
	if err != nil {
		return nil, err
	}

	res := dbToPbEmail(email)

	log.Printf("GRPC: Removed email entry(id:%d), client Addr: %s\n", email.ID, getIpFromCtx(ctx))
	return &pb.EmailResponse{Email: res}, nil
}

func (s *Server) DeleteUnsubscribed(ctx context.Context, req *pb.DeleteUnsubscribedRequest) (*pb.EmailBatchResponse, error) {
	emails, err := s.DB.DeleteUnsubscribed()
	if err != nil {
		return nil, err
	}

	res := dbToPbEmails(emails)
	
	log.Printf("GRPC: Deleted all unsubscribed email entries, client Addr:%s\n", getIpFromCtx(ctx))
	return &pb.EmailBatchResponse{Emails: res}, nil
}

func (s *Server) DeleteUnsubscribedBefore(ctx context.Context, req *pb.DeleteUnsubscribedBeforeRequest) (*pb.EmailBatchResponse, error) {
	emails, err := s.DB.DeleteUnsubscribedBefore(req.Date)
	if err != nil {
		return nil, err
	}

	res := dbToPbEmails(emails)
	
	log.Printf("GRPC: Deleted all email entries unsubscribed before: %d, client Addr:%s\n", req.Date, getIpFromCtx(ctx))
	return &pb.EmailBatchResponse{Emails: res}, nil
}
