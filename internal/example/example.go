package example

import (
	"context"
	"fmt"
	sample "github.com/bingbingmo/sample-server/proto"
	"github.com/google/uuid"
	"github.com/micro/go-micro/v2/logger"
	"math/rand"
)

// CreateVM CreateVM
func CreateVM(ctx context.Context, req *sample.CreateVMRequest) (string, error) {
	logger.Info("Create VM")

	return uuid.New().String(), nil
}

// GetVMList GetVMList
func GetVMList(ctx context.Context) ([]*sample.Array, error) {
	logger.Info("Get VM list")
	var vms []*sample.Array
	for i := 0; i < 5; i++ {
		vms = append(vms, &sample.Array{Name: fmt.Sprintf("vm-%d", i), Id: uuid.New().String()})
	}

	return vms, nil
}

// GetVM GetVM
func GetVM(ctx context.Context, req *sample.GetVMRequest) (*sample.Array, error) {
	logger.Infof("Get VM(%s)", req.GetUuid())

	if req.GetUuid() == "" {
		return nil, nil
	}

	return &sample.Array{Id: req.GetUuid(), Name: "vm"}, nil
}

// GetVMStatus GetVMStatus
func GetVMStatus(ctx context.Context, req *sample.GetVMStatusRequest) (int32, error) {
	logger.Info("Get VM status")

	return rand.Int31n(100), nil
}

// DeleteVM DeleteVM
func DeleteVM(ctx context.Context, req *sample.DeleteVMRequest) error {
	logger.Info("Delete VM")
	return nil
}

// CheckNameAvailability CheckNameAvailability
func CheckNameAvailability(ctx context.Context, req *sample.CheckNameAvailabilityRequest) error {
	logger.Info("Check VM name is availability")
	return nil
}
