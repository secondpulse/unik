package osv

import (
	"github.com/emc-advanced-dev/unik/pkg/types"
	"github.com/emc-advanced-dev/pkg/errors"
)

const OSV_VIRTUALBOX_MEMORY = 512

type OsvVirtualboxCompiler struct {}

func (osvCompiler *OsvVirtualboxCompiler) CompileRawImage(params types.CompileImageParams) (_ *types.RawImage, err error) {
	resultFile, err := compileRawImage(params, false)
	if err != nil {
		return nil, errors.New("failed to compile raw osv image", err)
	}
	return &types.RawImage{
		LocalImagePath: resultFile,
		StageSpec: types.StageSpec{
			ImageFormat: types.ImageFormat_QCOW2,
		},
		RunSpec: types.RunSpec{
			DeviceMappings: []types.DeviceMapping{
				types.DeviceMapping{MountPoint: "/", DeviceName: "/dev/sda1"},
			},
			StorageDriver: types.StorageDriver_SATA,
			DefaultInstanceMemory: OSV_VIRTUALBOX_MEMORY,
		},
	}, nil
}