package example

import (
	"context"
	sample "sample/proto"
)

// CreateVM CreateVM
func CreateVM(ctx context.Context, in *sample.CreateVMRequest, out *sample.CreateVMResponse) error {
	return nil
}

// GetVMList GetVMList
func GetVMList(ctx context.Context, in *sample.Empty, out *sample.GetVMListResponse) error {
	return nil
}

// GetVM GetVM
func GetVM(ctx context.Context, in *sample.GetVMRequest, out *sample.GetVMResponse) error {
	return nil
}

// GetVMStatus GetVMStatus
func GetVMStatus(ctx context.Context, in *sample.GetVMStatusRequest, out *sample.GetVMStatusResponse) error {
	return nil
}

// DeleteVM DeleteVM
func DeleteVM(ctx context.Context, in *sample.DeleteVMRequest, out *sample.Empty) error {
	return nil
}

// CheckNameAvailability CheckNameAvailability
func CheckNameAvailability(ctx context.Context, in *sample.CheckNameAvailabilityRequest, out *sample.Empty) error {
	return nil
}
