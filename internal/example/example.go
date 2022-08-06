package example

import (
	"context"
	sample "github.com/bingbingmo/sample-server/proto"
	"github.com/micro/go-micro/v2/logger"
)

// CreateVM CreateVM
func CreateVM(ctx context.Context, in *sample.CreateVMRequest, out *sample.CreateVMResponse) error {
	logger.Info("Create VM")
	return nil
}

// GetVMList GetVMList
func GetVMList(ctx context.Context, in *sample.Empty, out *sample.GetVMListResponse) error {
	logger.Info("Get VM list")
	return nil
}

// GetVM GetVM
func GetVM(ctx context.Context, in *sample.GetVMRequest, out *sample.GetVMResponse) error {
	logger.Info("Get VM")
	return nil
}

// GetVMStatus GetVMStatus
func GetVMStatus(ctx context.Context, in *sample.GetVMStatusRequest, out *sample.GetVMStatusResponse) error {
	logger.Info("Get VM status")
	return nil
}

// DeleteVM DeleteVM
func DeleteVM(ctx context.Context, in *sample.DeleteVMRequest, out *sample.Empty) error {
	logger.Info("Delete VM")
	return nil
}

// CheckNameAvailability CheckNameAvailability
func CheckNameAvailability(ctx context.Context, in *sample.CheckNameAvailabilityRequest, out *sample.Empty) error {
	logger.Info("Check VM name is availability")
	return nil
}
