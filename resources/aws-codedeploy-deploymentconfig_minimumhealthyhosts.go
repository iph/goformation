package resources

// AWS::CodeDeploy::DeploymentConfig.MinimumHealthyHosts AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentconfig-minimumhealthyhosts.html
type AWSCodeDeployDeploymentConfigMinimumHealthyHosts struct {

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentconfig-minimumhealthyhosts.html#cfn-codedeploy-deploymentconfig-minimumhealthyhosts-type
	Type string `json:"Type"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentconfig-minimumhealthyhosts.html#cfn-codedeploy-deploymentconfig-minimumhealthyhosts-value
	Value int64 `json:"Value"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeDeployDeploymentConfigMinimumHealthyHosts) AWSCloudFormationType() string {
	return "AWS::CodeDeploy::DeploymentConfig.MinimumHealthyHosts"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodeDeployDeploymentConfigMinimumHealthyHosts) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}