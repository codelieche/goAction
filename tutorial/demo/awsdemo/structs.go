package awsdemo

type ProjectConfig struct {
	// 图片上传的目标桶
	AwsRegion string     // aws的区域
	TargetBucket  string // 目标桶
	TargetRootDir string // 目标桶根目录

	// aws密匙
	AwsSecretKeyId     string // Aws Secret Key Id
	AwsSecretAccessKey string // Aws Secret Access Key
}

