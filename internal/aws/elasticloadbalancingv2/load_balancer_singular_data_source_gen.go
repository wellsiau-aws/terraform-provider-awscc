// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_elasticloadbalancingv2_load_balancer", loadBalancerDataSource)
}

// loadBalancerDataSource returns the Terraform awscc_elasticloadbalancingv2_load_balancer data source.
// This Terraform data source corresponds to the CloudFormation AWS::ElasticLoadBalancingV2::LoadBalancer resource.
func loadBalancerDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CanonicalHostedZoneID
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"canonical_hosted_zone_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DNSName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"dns_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether to evaluate inbound security group rules for traffic sent to a Network Load Balancer through privatelink.",
		//	  "type": "string"
		//	}
		"enforce_security_group_inbound_rules_on_private_link_traffic": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether to evaluate inbound security group rules for traffic sent to a Network Load Balancer through privatelink.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IpAddressType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Note: Internal load balancers must use the ``ipv4`` IP address type.\n [Application Load Balancers] The IP address type. The possible values are ``ipv4`` (for only IPv4 addresses), ``dualstack`` (for IPv4 and IPv6 addresses), and ``dualstack-without-public-ipv4`` (for IPv6 only public addresses, with private IPv4 and IPv6 addresses).\n Note: Application Load Balancer authentication only supports IPv4 addresses when connecting to an Identity Provider (IdP) or Amazon Cognito endpoint. Without a public IPv4 address the load balancer cannot complete the authentication process, resulting in HTTP 500 errors.\n [Network Load Balancers] The IP address type. The possible values are ``ipv4`` (for only IPv4 addresses) and ``dualstack`` (for IPv4 and IPv6 addresses). You can’t specify ``dualstack`` for a load balancer with a UDP or TCP_UDP listener.\n [Gateway Load Balancers] The IP address type. The possible values are ``ipv4`` (for only IPv4 addresses) and ``dualstack`` (for IPv4 and IPv6 addresses).",
		//	  "type": "string"
		//	}
		"ip_address_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Note: Internal load balancers must use the ``ipv4`` IP address type.\n [Application Load Balancers] The IP address type. The possible values are ``ipv4`` (for only IPv4 addresses), ``dualstack`` (for IPv4 and IPv6 addresses), and ``dualstack-without-public-ipv4`` (for IPv6 only public addresses, with private IPv4 and IPv6 addresses).\n Note: Application Load Balancer authentication only supports IPv4 addresses when connecting to an Identity Provider (IdP) or Amazon Cognito endpoint. Without a public IPv4 address the load balancer cannot complete the authentication process, resulting in HTTP 500 errors.\n [Network Load Balancers] The IP address type. The possible values are ``ipv4`` (for only IPv4 addresses) and ``dualstack`` (for IPv4 and IPv6 addresses). You can’t specify ``dualstack`` for a load balancer with a UDP or TCP_UDP listener.\n [Gateway Load Balancers] The IP address type. The possible values are ``ipv4`` (for only IPv4 addresses) and ``dualstack`` (for IPv4 and IPv6 addresses).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LoadBalancerArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"load_balancer_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LoadBalancerAttributes
		// CloudFormation resource type schema:
		//
		//	{
		//	  "arrayType": "AttributeList",
		//	  "description": "The load balancer attributes.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Specifies an attribute for an Application Load Balancer, a Network Load Balancer, or a Gateway Load Balancer.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The name of the attribute.\n The following attributes are supported by all load balancers:\n  +   ``deletion_protection.enabled`` - Indicates whether deletion protection is enabled. The value is ``true`` or ``false``. The default is ``false``.\n  +   ``load_balancing.cross_zone.enabled`` - Indicates whether cross-zone load balancing is enabled. The possible values are ``true`` and ``false``. The default for Network Load Balancers and Gateway Load Balancers is ``false``. The default for Application Load Balancers is ``true``, and cannot be changed.\n  \n The following attributes are supported by both Application Load Balancers and Network Load Balancers:\n  +   ``access_logs.s3.enabled`` - Indicates whether access logs are enabled. The value is ``true`` or ``false``. The default is ``false``.\n  +   ``access_logs.s3.bucket`` - The name of the S3 bucket for the access logs. This attribute is required if access logs are enabled. The bucket must exist in the same region as the load balancer and have a bucket policy that grants Elastic Load Balancing permissions to write to the bucket.\n  +   ``access_logs.s3.prefix`` - The prefix for the location in the S3 bucket for the access logs.\n  +   ``ipv6.deny_all_igw_traffic`` - Blocks internet gateway (IGW) access to the load balancer. It is set to ``false`` for internet-facing load balancers and ``true`` for internal load balancers, preventing unintended access to your internal load balancer through an internet gateway.\n  \n The following attributes are supported by only Application Load Balancers:\n  +   ``idle_timeout.timeout_seconds`` - The idle timeout value, in seconds. The valid range is 1-4000 seconds. The default is 60 seconds.\n  +   ``client_keep_alive.seconds`` - The client keep alive value, in seconds. The valid range is 60-604800 seconds. The default is 3600 seconds.\n  +   ``connection_logs.s3.enabled`` - Indicates whether connection logs are enabled. The value is ``true`` or ``false``. The default is ``false``.\n  +   ``connection_logs.s3.bucket`` - The name of the S3 bucket for the connection logs. This attribute is required if connection logs are enabled. The bucket must exist in the same region as the load balancer and have a bucket policy that grants Elastic Load Balancing permissions to write to the bucket.\n  +   ``connection_logs.s3.prefix`` - The prefix for the location in the S3 bucket for the connection logs.\n  +   ``routing.http.desync_mitigation_mode`` - Determines how the load balancer handles requests that might pose a security risk to your application. The possible values are ``monitor``, ``defensive``, and ``strictest``. The default is ``defensive``.\n  +   ``routing.http.drop_invalid_header_fields.enabled`` - Indicates whether HTTP headers with invalid header fields are removed by the load balancer (``true``) or routed to targets (``false``). The default is ``false``.\n  +   ``routing.http.preserve_host_header.enabled`` - Indicates whether the Application Load Balancer should preserve the ``Host`` header in the HTTP request and send it to the target without any change. The possible values are ``true`` and ``false``. The default is ``false``.\n  +   ``routing.http.x_amzn_tls_version_and_cipher_suite.enabled`` - Indicates whether the two headers (``x-amzn-tls-version`` and ``x-amzn-tls-cipher-suite``), which contain information about the negotiated TLS version and cipher suite, are added to the client request before sending it to the target. The ``x-amzn-tls-version`` header has information about the TLS protocol version negotiated with the client, and the ``x-amzn-tls-cipher-suite`` header has information about the cipher suite negotiated with the client. Both headers are in OpenSSL format. The possible values for the attribute are ``true`` and ``false``. The default is ``false``.\n  +   ``routing.http.xff_client_port.enabled`` - Indicates whether the ``X-Forwarded-For`` header should preserve the source port that the client used to connect to the load balancer. The possible values are ``true`` and ``false``. The default is ``false``.\n  +   ``routing.http.xff_header_processing.mode`` - Enables you to modify, preserve, or remove the ``X-Forwarded-For`` header in the HTTP request before the Application Load Balancer sends the request to the target. The possible values are ``append``, ``preserve``, and ``remove``. The default is ``append``.\n  +  If the value is ``append``, the Application Load Balancer adds the client IP address (of the last hop) to the ``X-Forwarded-For`` header in the HTTP request before it sends it to targets.\n  +  If the value is ``preserve`` the Application Load Balancer preserves the ``X-Forwarded-For`` header in the HTTP request, and sends it to targets without any change.\n  +  If the value is ``remove``, the Application Load Balancer removes the ``X-Forwarded-For`` header in the HTTP request before it sends it to targets.\n  \n  +   ``routing.http2.enabled`` - Indicates whether HTTP/2 is enabled. The possible values are ``true`` and ``false``. The default is ``true``. Elastic Load Balancing requires that message header names contain only alphanumeric characters and hyphens.\n  +   ``waf.fail_open.enabled`` - Indicates whether to allow a WAF-enabled load balancer to route requests to targets if it is unable to forward the request to AWS WAF. The possible values are ``true`` and ``false``. The default is ``false``.\n  \n The following attributes are supported by only Network Load Balancers:\n  +   ``dns_record.client_routing_policy`` - Indicates how traffic is distributed among the load balancer Availability Zones. The possible values are ``availability_zone_affinity`` with 100 percent zonal affinity, ``partial_availability_zone_affinity`` with 85 percent zonal affinity, and ``any_availability_zone`` with 0 percent zonal affinity.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value of the attribute.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"load_balancer_attributes": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The name of the attribute.\n The following attributes are supported by all load balancers:\n  +   ``deletion_protection.enabled`` - Indicates whether deletion protection is enabled. The value is ``true`` or ``false``. The default is ``false``.\n  +   ``load_balancing.cross_zone.enabled`` - Indicates whether cross-zone load balancing is enabled. The possible values are ``true`` and ``false``. The default for Network Load Balancers and Gateway Load Balancers is ``false``. The default for Application Load Balancers is ``true``, and cannot be changed.\n  \n The following attributes are supported by both Application Load Balancers and Network Load Balancers:\n  +   ``access_logs.s3.enabled`` - Indicates whether access logs are enabled. The value is ``true`` or ``false``. The default is ``false``.\n  +   ``access_logs.s3.bucket`` - The name of the S3 bucket for the access logs. This attribute is required if access logs are enabled. The bucket must exist in the same region as the load balancer and have a bucket policy that grants Elastic Load Balancing permissions to write to the bucket.\n  +   ``access_logs.s3.prefix`` - The prefix for the location in the S3 bucket for the access logs.\n  +   ``ipv6.deny_all_igw_traffic`` - Blocks internet gateway (IGW) access to the load balancer. It is set to ``false`` for internet-facing load balancers and ``true`` for internal load balancers, preventing unintended access to your internal load balancer through an internet gateway.\n  \n The following attributes are supported by only Application Load Balancers:\n  +   ``idle_timeout.timeout_seconds`` - The idle timeout value, in seconds. The valid range is 1-4000 seconds. The default is 60 seconds.\n  +   ``client_keep_alive.seconds`` - The client keep alive value, in seconds. The valid range is 60-604800 seconds. The default is 3600 seconds.\n  +   ``connection_logs.s3.enabled`` - Indicates whether connection logs are enabled. The value is ``true`` or ``false``. The default is ``false``.\n  +   ``connection_logs.s3.bucket`` - The name of the S3 bucket for the connection logs. This attribute is required if connection logs are enabled. The bucket must exist in the same region as the load balancer and have a bucket policy that grants Elastic Load Balancing permissions to write to the bucket.\n  +   ``connection_logs.s3.prefix`` - The prefix for the location in the S3 bucket for the connection logs.\n  +   ``routing.http.desync_mitigation_mode`` - Determines how the load balancer handles requests that might pose a security risk to your application. The possible values are ``monitor``, ``defensive``, and ``strictest``. The default is ``defensive``.\n  +   ``routing.http.drop_invalid_header_fields.enabled`` - Indicates whether HTTP headers with invalid header fields are removed by the load balancer (``true``) or routed to targets (``false``). The default is ``false``.\n  +   ``routing.http.preserve_host_header.enabled`` - Indicates whether the Application Load Balancer should preserve the ``Host`` header in the HTTP request and send it to the target without any change. The possible values are ``true`` and ``false``. The default is ``false``.\n  +   ``routing.http.x_amzn_tls_version_and_cipher_suite.enabled`` - Indicates whether the two headers (``x-amzn-tls-version`` and ``x-amzn-tls-cipher-suite``), which contain information about the negotiated TLS version and cipher suite, are added to the client request before sending it to the target. The ``x-amzn-tls-version`` header has information about the TLS protocol version negotiated with the client, and the ``x-amzn-tls-cipher-suite`` header has information about the cipher suite negotiated with the client. Both headers are in OpenSSL format. The possible values for the attribute are ``true`` and ``false``. The default is ``false``.\n  +   ``routing.http.xff_client_port.enabled`` - Indicates whether the ``X-Forwarded-For`` header should preserve the source port that the client used to connect to the load balancer. The possible values are ``true`` and ``false``. The default is ``false``.\n  +   ``routing.http.xff_header_processing.mode`` - Enables you to modify, preserve, or remove the ``X-Forwarded-For`` header in the HTTP request before the Application Load Balancer sends the request to the target. The possible values are ``append``, ``preserve``, and ``remove``. The default is ``append``.\n  +  If the value is ``append``, the Application Load Balancer adds the client IP address (of the last hop) to the ``X-Forwarded-For`` header in the HTTP request before it sends it to targets.\n  +  If the value is ``preserve`` the Application Load Balancer preserves the ``X-Forwarded-For`` header in the HTTP request, and sends it to targets without any change.\n  +  If the value is ``remove``, the Application Load Balancer removes the ``X-Forwarded-For`` header in the HTTP request before it sends it to targets.\n  \n  +   ``routing.http2.enabled`` - Indicates whether HTTP/2 is enabled. The possible values are ``true`` and ``false``. The default is ``true``. Elastic Load Balancing requires that message header names contain only alphanumeric characters and hyphens.\n  +   ``waf.fail_open.enabled`` - Indicates whether to allow a WAF-enabled load balancer to route requests to targets if it is unable to forward the request to AWS WAF. The possible values are ``true`` and ``false``. The default is ``false``.\n  \n The following attributes are supported by only Network Load Balancers:\n  +   ``dns_record.client_routing_policy`` - Indicates how traffic is distributed among the load balancer Availability Zones. The possible values are ``availability_zone_affinity`` with 100 percent zonal affinity, ``partial_availability_zone_affinity`` with 85 percent zonal affinity, and ``any_availability_zone`` with 0 percent zonal affinity.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value of the attribute.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The load balancer attributes.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LoadBalancerFullName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"load_balancer_full_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LoadBalancerName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"load_balancer_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the load balancer. This name must be unique per region per account, can have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, must not begin or end with a hyphen, and must not begin with \"internal-\".\n If you don't specify a name, AWS CloudFormation generates a unique physical ID for the load balancer. If you specify a name, you cannot perform updates that require replacement of this resource, but you can perform other updates. To replace the resource, specify a new name.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the load balancer. This name must be unique per region per account, can have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, must not begin or end with a hyphen, and must not begin with \"internal-\".\n If you don't specify a name, AWS CloudFormation generates a unique physical ID for the load balancer. If you specify a name, you cannot perform updates that require replacement of this resource, but you can perform other updates. To replace the resource, specify a new name.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Scheme
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The nodes of an Internet-facing load balancer have public IP addresses. The DNS name of an Internet-facing load balancer is publicly resolvable to the public IP addresses of the nodes. Therefore, Internet-facing load balancers can route requests from clients over the internet.\n The nodes of an internal load balancer have only private IP addresses. The DNS name of an internal load balancer is publicly resolvable to the private IP addresses of the nodes. Therefore, internal load balancers can route requests only from clients with access to the VPC for the load balancer.\n The default is an Internet-facing load balancer.\n You cannot specify a scheme for a Gateway Load Balancer.",
		//	  "type": "string"
		//	}
		"scheme": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The nodes of an Internet-facing load balancer have public IP addresses. The DNS name of an Internet-facing load balancer is publicly resolvable to the public IP addresses of the nodes. Therefore, Internet-facing load balancers can route requests from clients over the internet.\n The nodes of an internal load balancer have only private IP addresses. The DNS name of an internal load balancer is publicly resolvable to the private IP addresses of the nodes. Therefore, internal load balancers can route requests only from clients with access to the VPC for the load balancer.\n The default is an Internet-facing load balancer.\n You cannot specify a scheme for a Gateway Load Balancer.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SecurityGroups
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "[Application Load Balancers and Network Load Balancers] The IDs of the security groups for the load balancer.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"security_groups": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "[Application Load Balancers and Network Load Balancers] The IDs of the security groups for the load balancer.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SubnetMappings
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IDs of the subnets. You can specify only one subnet per Availability Zone. You must specify either subnets or subnet mappings, but not both.\n [Application Load Balancers] You must specify subnets from at least two Availability Zones. You cannot specify Elastic IP addresses for your subnets.\n [Application Load Balancers on Outposts] You must specify one Outpost subnet.\n [Application Load Balancers on Local Zones] You can specify subnets from one or more Local Zones.\n [Network Load Balancers] You can specify subnets from one or more Availability Zones. You can specify one Elastic IP address per subnet if you need static IP addresses for your internet-facing load balancer. For internal load balancers, you can specify one private IP address per subnet from the IPv4 range of the subnet. For internet-facing load balancer, you can specify one IPv6 address per subnet.\n [Gateway Load Balancers] You can specify subnets from one or more Availability Zones. You cannot specify Elastic IP addresses for your subnets.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Specifies a subnet for a load balancer.",
		//	    "properties": {
		//	      "AllocationId": {
		//	        "description": "[Network Load Balancers] The allocation ID of the Elastic IP address for an internet-facing load balancer.",
		//	        "type": "string"
		//	      },
		//	      "IPv6Address": {
		//	        "description": "[Network Load Balancers] The IPv6 address.",
		//	        "type": "string"
		//	      },
		//	      "PrivateIPv4Address": {
		//	        "description": "[Network Load Balancers] The private IPv4 address for an internal load balancer.",
		//	        "type": "string"
		//	      },
		//	      "SubnetId": {
		//	        "description": "The ID of the subnet.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "SubnetId"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"subnet_mappings": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: AllocationId
					"allocation_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "[Network Load Balancers] The allocation ID of the Elastic IP address for an internet-facing load balancer.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: IPv6Address
					"i_pv_6_address": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "[Network Load Balancers] The IPv6 address.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: PrivateIPv4Address
					"private_i_pv_4_address": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "[Network Load Balancers] The private IPv4 address for an internal load balancer.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: SubnetId
					"subnet_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The ID of the subnet.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The IDs of the subnets. You can specify only one subnet per Availability Zone. You must specify either subnets or subnet mappings, but not both.\n [Application Load Balancers] You must specify subnets from at least two Availability Zones. You cannot specify Elastic IP addresses for your subnets.\n [Application Load Balancers on Outposts] You must specify one Outpost subnet.\n [Application Load Balancers on Local Zones] You can specify subnets from one or more Local Zones.\n [Network Load Balancers] You can specify subnets from one or more Availability Zones. You can specify one Elastic IP address per subnet if you need static IP addresses for your internet-facing load balancer. For internal load balancers, you can specify one private IP address per subnet from the IPv4 range of the subnet. For internet-facing load balancer, you can specify one IPv6 address per subnet.\n [Gateway Load Balancers] You can specify subnets from one or more Availability Zones. You cannot specify Elastic IP addresses for your subnets.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Subnets
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IDs of the subnets. You can specify only one subnet per Availability Zone. You must specify either subnets or subnet mappings, but not both. To specify an Elastic IP address, specify subnet mappings instead of subnets.\n [Application Load Balancers] You must specify subnets from at least two Availability Zones.\n [Application Load Balancers on Outposts] You must specify one Outpost subnet.\n [Application Load Balancers on Local Zones] You can specify subnets from one or more Local Zones.\n [Network Load Balancers] You can specify subnets from one or more Availability Zones.\n [Gateway Load Balancers] You can specify subnets from one or more Availability Zones.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"subnets": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The IDs of the subnets. You can specify only one subnet per Availability Zone. You must specify either subnets or subnet mappings, but not both. To specify an Elastic IP address, specify subnet mappings instead of subnets.\n [Application Load Balancers] You must specify subnets from at least two Availability Zones.\n [Application Load Balancers on Outposts] You must specify one Outpost subnet.\n [Application Load Balancers on Local Zones] You can specify subnets from one or more Local Zones.\n [Network Load Balancers] You can specify subnets from one or more Availability Zones.\n [Gateway Load Balancers] You can specify subnets from one or more Availability Zones.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tags to assign to the load balancer.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Information about a tag.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key of the tag.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value of the tag.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key of the tag.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value of the tag.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The tags to assign to the load balancer.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of load balancer. The default is ``application``.",
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of load balancer. The default is ``application``.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::ElasticLoadBalancingV2::LoadBalancer",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ElasticLoadBalancingV2::LoadBalancer").WithTerraformTypeName("awscc_elasticloadbalancingv2_load_balancer")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"allocation_id":            "AllocationId",
		"canonical_hosted_zone_id": "CanonicalHostedZoneID",
		"dns_name":                 "DNSName",
		"enforce_security_group_inbound_rules_on_private_link_traffic": "EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic",
		"i_pv_6_address":           "IPv6Address",
		"ip_address_type":          "IpAddressType",
		"key":                      "Key",
		"load_balancer_arn":        "LoadBalancerArn",
		"load_balancer_attributes": "LoadBalancerAttributes",
		"load_balancer_full_name":  "LoadBalancerFullName",
		"load_balancer_name":       "LoadBalancerName",
		"name":                     "Name",
		"private_i_pv_4_address":   "PrivateIPv4Address",
		"scheme":                   "Scheme",
		"security_groups":          "SecurityGroups",
		"subnet_id":                "SubnetId",
		"subnet_mappings":          "SubnetMappings",
		"subnets":                  "Subnets",
		"tags":                     "Tags",
		"type":                     "Type",
		"value":                    "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
