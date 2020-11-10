package utils

const (
	ResourceName = "ucloud.cn/gpu-mem"
	CountName    = "ucloud.cn/gpu-count"

	EnvNVGPU              = "NVIDIA_VISIBLE_DEVICES"
	EnvResourceIndex      = "UCLOUD_CN_GPU_MEM_IDX"
	EnvResourceByPod      = "UCLOUD_CN_GPU_MEM_POD"
	EnvResourceByDev      = "UCLOUD_CN_GPU_MEM_DEV"
	EnvAssignedFlag       = "UCLOUD_CN_GPU_MEM_ASSIGNED"
	EnvResourceAssumeTime = "UCLOUD_CN_GPU_MEM_ASSUME_TIME"
)
