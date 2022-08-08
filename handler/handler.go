package handler

import (
	"context"
	"github.com/bingbingmo/sample-server/internal/example"
	"github.com/bingbingmo/sample-server/proto"
	"github.com/micro/go-micro/v2/errors"
)

// SampleServerHandler sample server handler
type SampleServerHandler struct {
	sample.SampleServerHandler
}

// NewSampleServerHandler sample server handler 생성
func NewSampleServerHandler() (*SampleServerHandler, error) {
	return &SampleServerHandler{}, nil
}

func (h *SampleServerHandler) CreateVM(ctx context.Context, req *sample.CreateVMRequest, rsp *sample.CreateVMResponse) error {
	var err error

	rsp.Id, err = example.CreateVM(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (h *SampleServerHandler) GetVMList(ctx context.Context, _ *sample.Empty, rsp *sample.GetVMListResponse) error {
	var err error

	rsp.Arrays, err = example.GetVMList(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (h *SampleServerHandler) GetVM(ctx context.Context, req *sample.GetVMRequest, rsp *sample.GetVMResponse) error {
	var err error

	rsp.Array, err = example.GetVM(ctx, req)
	if err != nil {
		return err
	}

	if rsp.Array == nil {
		return errors.NotFound("sample", "detail")
	}

	return nil
}

func (h *SampleServerHandler) GetVMStatus(ctx context.Context, req *sample.GetVMStatusRequest, rsp *sample.GetVMStatusResponse) error {
	var err error

	rsp.CpuUtilization, err = example.GetVMStatus(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (h *SampleServerHandler) DeleteVM(ctx context.Context, req *sample.DeleteVMRequest, _ *sample.Empty) error {
	err := example.DeleteVM(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (h *SampleServerHandler) CheckNameAvailability(ctx context.Context, req *sample.CheckNameAvailabilityRequest, _ *sample.Empty) error {
	err := example.CheckNameAvailability(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
