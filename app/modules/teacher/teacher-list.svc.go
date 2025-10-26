package teacher

import (
	"context"
	"time"

	"github.com/easy-attend-serviceV3/app/utils"
	"github.com/easy-attend-serviceV3/app/utils/base"
	"github.com/google/uuid"
)

type ListServiceRequest struct {
	base.RequestPaginate
}

type ListServiceResponse struct {
	ID          uuid.UUID `json:"id"`
	SchoolID    uuid.UUID `json:"school_id"`
	ClassroomID uuid.UUID `json:"classroom_id"`
	PrefixID    uuid.UUID `json:"prefix_id"`
	GenderID    uuid.UUID `json:"gender_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (s *Service) ListService(ctx context.Context, request *ListServiceRequest) ([]*ListServiceResponse, *base.ResponsePaginate, error) {
	span, log := utils.LogSpanFromContext(ctx)
	span.AddEvent(`teacher.svc.list.start`)

	data, err := s.db.GetListTeacher(ctx)
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}

	var response []*ListServiceResponse
	for _, v := range data {
		response = append(response, &ListServiceResponse{
			ID:          v.ID,
			SchoolID:    v.SchoolID,
			ClassroomID: v.ClassroomID,
			PrefixID:    v.PrefixID,
			GenderID:    v.GenderID,
			FirstName:   v.FirstName,
			LastName:    v.LastName,
			Email:       v.Email,
			Phone:       v.Phone,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		})
	}

	// Create pagination response
	page := &base.ResponsePaginate{
		Page:  int64(request.Page),
		Size:  int64(request.Size),
		Total: int64(len(response)),
	}

	span.AddEvent(`teacher.svc.list.success`)
	return response, page, nil
}
