package handler

import (
	"context"
	"github.com/bingbingmo/sample-server/internal/example"
	"github.com/bingbingmo/sample-server/proto"
)

// SampleServerHandler sample server handler
type SampleServerHandler struct {
	sample.SampleServerHandler
}

// NewSampleServerHandler sample server handler 생성
func NewSampleServerHandler() (*SampleServerHandler, error) {
	return &SampleServerHandler{}, nil
}

func (h *SampleServerHandler) CreateVM(ctx context.Context, in *sample.CreateVMRequest, out *sample.CreateVMResponse) error {
	return example.CreateVM(ctx, in, out)
}

func (h *SampleServerHandler) GetVMList(ctx context.Context, in *sample.Empty, out *sample.GetVMListResponse) error {
	return example.GetVMList(ctx, in, out)
}

func (h *SampleServerHandler) GetVM(ctx context.Context, in *sample.GetVMRequest, out *sample.GetVMResponse) error {
	return example.GetVM(ctx, in, out)
}

func (h *SampleServerHandler) GetVMStatus(ctx context.Context, in *sample.GetVMStatusRequest, out *sample.GetVMStatusResponse) error {
	return example.GetVMStatus(ctx, in, out)
}

func (h *SampleServerHandler) DeleteVM(ctx context.Context, in *sample.DeleteVMRequest, out *sample.Empty) error {
	return example.DeleteVM(ctx, in, out)
}

func (h *SampleServerHandler) CheckNameAvailability(ctx context.Context, in *sample.CheckNameAvailabilityRequest, out *sample.Empty) error {
	return example.CheckNameAvailability(ctx, in, out)
}
