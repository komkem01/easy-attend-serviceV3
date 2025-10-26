package student

import (
	"context"
	"log/slog"
	"time"

	entitiesdto "github.com/easy-attend-serviceV3/app/modules/entities/dto"
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
	StudentCode string    `json:"student_code"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Phone       string    `json:"phone"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (s *Service) ListService(ctx context.Context, request *ListServiceRequest) ([]*ListServiceResponse, *base.ResponsePaginate, error) {
	span, log := utils.LogSpanFromContext(ctx)
	span.AddEvent(`student.svc.list.start`)

	data, err := s.db.GetListStudent(ctx, &entitiesdto.StudentListResponse{})
	if err != nil {
		log.With(slog.Any(`body`, request)).Error(err)
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
			StudentCode: v.StudentCode,
			FirstName:   v.FirstName,
			LastName:    v.LastName,
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

	span.AddEvent(`student.svc.list.success`)
	return response, page, nil
}
