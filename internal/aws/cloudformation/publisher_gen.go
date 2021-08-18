// Code generated by generators/resource/main.go; DO NOT EDIT.

package cloudformation

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("awscc_cloudformation_publisher", publisherResourceType)
}

// publisherResourceType returns the Terraform awscc_cloudformation_publisher resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::CloudFormation::Publisher resource type.
func publisherResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"accept_terms_and_conditions": {
			// Property: AcceptTermsAndConditions
			// CloudFormation resource type schema:
			// {
			//   "description": "Whether you accept the terms and conditions for publishing extensions in the CloudFormation registry. You must accept the terms and conditions in order to publish public extensions to the CloudFormation registry. The terms and conditions can be found at https://cloudformation-registry-documents.s3.amazonaws.com/Terms_and_Conditions_for_AWS_CloudFormation_Registry_Publishers.pdf",
			//   "type": "boolean"
			// }
			Description: "Whether you accept the terms and conditions for publishing extensions in the CloudFormation registry. You must accept the terms and conditions in order to publish public extensions to the CloudFormation registry. The terms and conditions can be found at https://cloudformation-registry-documents.s3.amazonaws.com/Terms_and_Conditions_for_AWS_CloudFormation_Registry_Publishers.pdf",
			Type:        types.BoolType,
			Required:    true,
			// AcceptTermsAndConditions is a force-new attribute.
		},
		"connection_arn": {
			// Property: ConnectionArn
			// CloudFormation resource type schema:
			// {
			//   "description": "If you are using a Bitbucket or GitHub account for identity verification, the Amazon Resource Name (ARN) for your connection to that account.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "If you are using a Bitbucket or GitHub account for identity verification, the Amazon Resource Name (ARN) for your connection to that account.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// ConnectionArn is a force-new attribute.
		},
		"identity_provider": {
			// Property: IdentityProvider
			// CloudFormation resource type schema:
			// {
			//   "description": "The type of account used as the identity provider when registering this publisher with CloudFormation.",
			//   "enum": [
			//     "AWS_Marketplace",
			//     "GitHub",
			//     "Bitbucket"
			//   ],
			//   "type": "string"
			// }
			Description: "The type of account used as the identity provider when registering this publisher with CloudFormation.",
			Type:        types.StringType,
			Computed:    true,
		},
		"publisher_id": {
			// Property: PublisherId
			// CloudFormation resource type schema:
			// {
			//   "description": "The publisher id assigned by CloudFormation for publishing in this region.",
			//   "maxLength": 40,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The publisher id assigned by CloudFormation for publishing in this region.",
			Type:        types.StringType,
			Computed:    true,
		},
		"publisher_profile": {
			// Property: PublisherProfile
			// CloudFormation resource type schema:
			// {
			//   "description": "The URL to the publisher's profile with the identity provider.",
			//   "maxLength": 1024,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The URL to the publisher's profile with the identity provider.",
			Type:        types.StringType,
			Computed:    true,
		},
		"publisher_status": {
			// Property: PublisherStatus
			// CloudFormation resource type schema:
			// {
			//   "description": "Whether the publisher is verified.",
			//   "enum": [
			//     "VERIFIED",
			//     "UNVERIFIED"
			//   ],
			//   "type": "string"
			// }
			Description: "Whether the publisher is verified.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Register as a publisher in the CloudFormation Registry.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFormation::Publisher").WithTerraformTypeName("awscc_cloudformation_publisher").WithTerraformSchema(schema)

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_cloudformation_publisher", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
