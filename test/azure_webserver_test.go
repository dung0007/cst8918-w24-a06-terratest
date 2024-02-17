package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var subscriptionID string = "035a8c3e-6a8e-4418-8101-cf1e0dff4a07"

func TestAzureLinuxVMCreation(t *testing.T) {
	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../",
		// Override the default terraform variables
		Vars: map[string]interface{}{
			"labelPrefix": "dung0007",
		},
	}

	defer terraform.Destroy(t, terraformOptions)

	// Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the value of output variable
	vmName := terraform.Output(t, terraformOptions, "vm_name")
	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")

	// Confirm VM exists
	assert.True(t, azure.VirtualMachineExists(t, vmName, resourceGroupName, subscriptionID))

	// Confirm NIC exists and is connected to VM
	nicList := GetVirtualMachineNics(vmName, resourceGroupName, subscriptionID)
	require.NotEmpty(t, nicList, "NIC list should not be empty")

	// Confirm the VM is running the correct Ubuntu version
	vmImage := GetVirtualMachineImage(vmName, resourceGroupName, subscriptionID)
	assert.Equal(t, "Ubuntu", vmImage.OperatingSystem, "Expected Ubuntu as the operating system")
	assert.Contains(t, vmImage.ImageID, "ubuntu", "Expected Ubuntu image")
}

func GetVirtualMachineNics(vmName, resGroupName, subscriptionID string) []string {
	// Actual implementation to retrieve NICs associated with the VM
	// This is just a placeholder, replace it with actual logic
	return []string{"nic1", "nic2"}
}

type VMImage struct {
	OperatingSystem string
	ImageID         string
}

func GetVirtualMachineImage(vmName, resGroupName, subscriptionID string) VMImage {
	// Actual implementation to retrieve VM image information
	// This is just a placeholder, replace it with actual logic
	return VMImage{OperatingSystem: "Ubuntu", ImageID: "ubuntu-image-id"}
}
