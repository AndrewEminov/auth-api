package note_v1

import (
	"context"

	converter "github.com/AndrewEminov/auth/internal/converter/note"
	desc "github.com/AndrewEminov/auth/pkg/note_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.noteService.Create(ctx, converter.ToInfo(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
