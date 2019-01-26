// +build !opencl

package eth

import (
	"errors"
	"fmt"

	"github.com/ether-core/go-ethereum/logger"
	"github.com/ether-core/go-ethereum/logger/glog"
)

const disabledInfo = "Set GO_OPENCL and re-build to enable."

func (s *Ethereum) StartMining(threads int, gpus string) error {
	eb, err := s.Etherbase()
	if err != nil {
		err = fmt.Errorf("Cannot start mining without etherbase address: %v", err)
		glog.V(logger.Error).Infoln(err)
		return err
	}

	if gpus != "" {
		return errors.New("GPU mining disabled. " + disabledInfo)
	}

	// CPU mining
	go s.miner.Start(eb, threads)
	return nil
}

func GPUBench(gpuid uint64) {
	fmt.Println("GPU mining disabled. " + disabledInfo)
}

func PrintOpenCLDevices() {
	fmt.Println("OpenCL disabled. " + disabledInfo)
}
