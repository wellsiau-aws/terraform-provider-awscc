// Code generated by generators/resource/main.go; DO NOT EDIT.

package amplify

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/generic"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("aws_amplify_domain", domain)
}

// domain returns the Terraform aws_amplify_domain resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Amplify::Domain resource type.
func domain(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"app_id": {
			// Property: AppId
			// CloudFormation resource type schema:
			/*
			   {
			     "maxLength": 20,
			     "minLength": 1,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Required: true,
			// AppId is a force-new attribute.
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			/*
			   {
			     "maxLength": 1000,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Computed: true,
		},
		"auto_sub_domain_creation_patterns": {
			// Property: AutoSubDomainCreationPatterns
			// CloudFormation resource type schema:
			/*
			   {
			     "items": {
			       "maxLength": 2048,
			       "minLength": 1,
			       "pattern": "",
			       "type": "string"
			     },
			     "type": "array",
			     "uniqueItems": false
			   }
			*/
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
		},
		"auto_sub_domain_iam_role": {
			// Property: AutoSubDomainIAMRole
			// CloudFormation resource type schema:
			/*
			   {
			     "maxLength": 1000,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Optional: true,
		},
		"certificate_record": {
			// Property: CertificateRecord
			// CloudFormation resource type schema:
			/*
			   {
			     "maxLength": 1000,
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Computed: true,
		},
		"domain_name": {
			// Property: DomainName
			// CloudFormation resource type schema:
			/*
			   {
			     "maxLength": 255,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Required: true,
			// DomainName is a force-new attribute.
		},
		"domain_status": {
			// Property: DomainStatus
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Computed: true,
		},
		"enable_auto_sub_domain": {
			// Property: EnableAutoSubDomain
			// CloudFormation resource type schema:
			/*
			   {
			     "type": "boolean"
			   }
			*/
			Type:     types.BoolType,
			Optional: true,
		},
		"status_reason": {
			// Property: StatusReason
			// CloudFormation resource type schema:
			/*
			   {
			     "maxLength": 1000,
			     "type": "string"
			   }
			*/
			Type:     types.StringType,
			Computed: true,
		},
		"sub_domain_settings": {
			// Property: SubDomainSettings
			// CloudFormation resource type schema:
			/*
			   {
			     "items": {
			       "additionalProperties": false,
			       "properties": {
			         "BranchName": {
			           "maxLength": 255,
			           "minLength": 1,
			           "pattern": "",
			           "type": "string"
			         },
			         "Prefix": {
			           "maxLength": 255,
			           "pattern": "",
			           "type": "string"
			         }
			       },
			       "$ref": "#/definitions/SubDomainSetting",
			       "required": [
			         "Prefix",
			         "BranchName"
			       ],
			       "type": "object"
			     },
			     "maxItems": 255,
			     "type": "array",
			     "uniqueItems": false
			   }
			*/
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"branch_name": {
						// Property: BranchName
						// CloudFormation resource type schema:
						/*
						   {
						     "maxLength": 255,
						     "minLength": 1,
						     "pattern": "",
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Required: true,
					},
					"prefix": {
						// Property: Prefix
						// CloudFormation resource type schema:
						/*
						   {
						     "maxLength": 255,
						     "pattern": "",
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Required: true,
					},
				},
				schema.ListNestedAttributesOptions{
					MaxItems: 255,
				},
			),
			Required: true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "The AWS::Amplify::Domain resource allows you to connect a custom domain to your app.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Amplify::Domain").WithTerraformTypeName("aws_amplify_domain").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_amplify_domain", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
