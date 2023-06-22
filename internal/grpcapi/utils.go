package grpcapi

import (
	"context"
	"time"

	"github.com/ddomd/maily/internal/mdb"
	"google.golang.org/grpc/peer"
	pb "github.com/ddomd/maily/proto"
)

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