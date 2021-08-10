// Code generated by generators/resource/main.go; DO NOT EDIT.

package codegurureviewer

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
	registry.AddResourceTypeFactory("aws_codegurureviewer_repository_association", repositoryAssociation)
}

// repositoryAssociation returns the Terraform aws_codegurureviewer_repository_association resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::CodeGuruReviewer::RepositoryAssociation resource type.
func repositoryAssociation(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"association_arn": {
			// Property: AssociationArn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The Amazon Resource Name (ARN) of the repository association.",
			     "maxLength": 256,
			     "minLength": 0,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The Amazon Resource Name (ARN) of the repository association.",
			Type:        types.StringType,
			Computed:    true,
		},
		"bucket_name": {
			// Property: BucketName
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The name of the S3 bucket associated with an associated S3 repository. It must start with `codeguru-reviewer-`.",
			     "maxLength": 63,
			     "minLength": 3,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The name of the S3 bucket associated with an associated S3 repository. It must start with `codeguru-reviewer-`.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// BucketName is a force-new attribute.
		},
		"connection_arn": {
			// Property: ConnectionArn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The Amazon Resource Name (ARN) of an AWS CodeStar Connections connection.",
			     "maxLength": 256,
			     "minLength": 0,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The Amazon Resource Name (ARN) of an AWS CodeStar Connections connection.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// ConnectionArn is a force-new attribute.
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Name of the repository to be associated.",
			     "maxLength": 100,
			     "minLength": 1,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "Name of the repository to be associated.",
			Type:        types.StringType,
			Required:    true,
			// Name is a force-new attribute.
		},
		"owner": {
			// Property: Owner
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The owner of the repository. For a Bitbucket repository, this is the username for the account that owns the repository.",
			     "maxLength": 100,
			     "minLength": 1,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The owner of the repository. For a Bitbucket repository, this is the username for the account that owns the repository.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// Owner is a force-new attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The tags associated with a repository association.",
			     "items": {
			       "additionalProperties": false,
			       "description": "A key-value pair to associate with a resource.",
			       "properties": {
			         "Key": {
			           "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. The allowed characters across services are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.",
			           "maxLength": 128,
			           "minLength": 1,
			           "type": "string"
			         },
			         "Value": {
			           "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. The allowed characters across services are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.",
			           "maxLength": 256,
			           "minLength": 0,
			           "type": "string"
			         }
			       },
			       "$ref": "#/definitions/Tag",
			       "required": [
			         "Value",
			         "Key"
			       ],
			       "type": "object"
			     },
			     "maxItems": 50,
			     "type": "array",
			     "uniqueItems": false
			   }
			*/
			Description: "The tags associated with a repository association.",
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. The allowed characters across services are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.",
						     "maxLength": 128,
						     "minLength": 1,
						     "type": "string"
						   }
						*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. The allowed characters across services are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.",
						Type:        types.StringType,
						Required:    true,
					},
					"value": {
						// Property: Value
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. The allowed characters across services are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.",
						     "maxLength": 256,
						     "minLength": 0,
						     "type": "string"
						   }
						*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. The allowed characters across services are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.",
						Type:        types.StringType,
						Required:    true,
					},
				},
				schema.ListNestedAttributesOptions{
					MaxItems: 50,
				},
			),
			Optional: true,
			Computed: true,
			// Tags is a force-new attribute.
		},
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The type of repository to be associated.",
			     "enum": [
			       "CodeCommit",
			       "Bitbucket",
			       "GitHubEnterpriseServer",
			       "S3Bucket"
			     ],
			     "type": "string"
			   }
			*/
			Description: "The type of repository to be associated.",
			Type:        types.StringType,
			Required:    true,
			// Type is a force-new attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "This resource schema represents the RepositoryAssociation resource in the Amazon CodeGuru Reviewer service.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CodeGuruReviewer::RepositoryAssociation").WithTerraformTypeName("aws_codegurureviewer_repository_association").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_codegurureviewer_repository_association", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
