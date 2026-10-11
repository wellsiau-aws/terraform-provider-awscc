resource "awscc_servicediscovery_http_namespace" "example" {
  name        = "example-http-namespace"
  description = "Example AWS Cloud Map HTTP namespace for API-based service discovery"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

output "http_namespace_id" {
  description = "The ID of the Cloud Map HTTP namespace"
  value       = awscc_servicediscovery_http_namespace.example.http_namespace_id
}

output "http_namespace_arn" {
  description = "The ARN of the Cloud Map HTTP namespace"
  value       = awscc_servicediscovery_http_namespace.example.arn
}
