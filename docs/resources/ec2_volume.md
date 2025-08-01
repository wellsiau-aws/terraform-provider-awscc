---
page_title: "awscc_ec2_volume Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Specifies an Amazon Elastic Block Store (Amazon EBS) volume.
  When you use CFNlong to update an Amazon EBS volume that modifies Iops, Size, or VolumeType, there is a cooldown period before another operation can occur. This can cause your stack to report being in UPDATE_IN_PROGRESS or UPDATE_ROLLBACK_IN_PROGRESS for long periods of time.
  Amazon EBS does not support sizing down an Amazon EBS volume. CFNlong does not attempt to modify an Amazon EBS volume to a smaller size on rollback.
  Some common scenarios when you might encounter a cooldown period for Amazon EBS include:
  You successfully update an Amazon EBS volume and the update succeeds. When you attempt another update within the cooldown window, that update will be subject to a cooldown period.You successfully update an Amazon EBS volume and the update succeeds but another change in your update-stack call fails. The rollback will be subject to a cooldown period.
  For more information, see Requirements for EBS volume modifications https://docs.aws.amazon.com/ebs/latest/userguide/modify-volume-requirements.html.
  DeletionPolicy attribute
  To control how CFNlong handles the volume when the stack is deleted, set a deletion policy for your volume. You can choose to retain the volume, to delete the volume, or to create a snapshot of the volume. For more information, see DeletionPolicy attribute https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html.
  If you set a deletion policy that creates a snapshot, all tags on the volume are included in the snapshot.
---

# awscc_ec2_volume (Resource)

Specifies an Amazon Elastic Block Store (Amazon EBS) volume.
 When you use CFNlong to update an Amazon EBS volume that modifies ``Iops``, ``Size``, or ``VolumeType``, there is a cooldown period before another operation can occur. This can cause your stack to report being in ``UPDATE_IN_PROGRESS`` or ``UPDATE_ROLLBACK_IN_PROGRESS`` for long periods of time.
 Amazon EBS does not support sizing down an Amazon EBS volume. CFNlong does not attempt to modify an Amazon EBS volume to a smaller size on rollback.
 Some common scenarios when you might encounter a cooldown period for Amazon EBS include:
  +  You successfully update an Amazon EBS volume and the update succeeds. When you attempt another update within the cooldown window, that update will be subject to a cooldown period.
  +  You successfully update an Amazon EBS volume and the update succeeds but another change in your ``update-stack`` call fails. The rollback will be subject to a cooldown period.
  
 For more information, see [Requirements for EBS volume modifications](https://docs.aws.amazon.com/ebs/latest/userguide/modify-volume-requirements.html).
  *DeletionPolicy attribute* 
 To control how CFNlong handles the volume when the stack is deleted, set a deletion policy for your volume. You can choose to retain the volume, to delete the volume, or to create a snapshot of the volume. For more information, see [DeletionPolicy attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html).
  If you set a deletion policy that creates a snapshot, all tags on the volume are included in the snapshot.

## Example Usage

To create ec2 volume with tags
```terraform
resource "awscc_ec2_volume" "example" {
  availability_zone = "us-west-2a"
  size              = 40

  tags = [
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `availability_zone` (String) The ID of the Availability Zone in which to create the volume. For example, ``us-east-1a``.

### Optional

- `auto_enable_io` (Boolean) Indicates whether the volume is auto-enabled for I/O operations. By default, Amazon EBS disables I/O to the volume from attached EC2 instances when it determines that a volume's data is potentially inconsistent. If the consistency of the volume is not a concern, and you prefer that the volume be made available immediately if it's impaired, you can configure the volume to automatically enable I/O.
- `encrypted` (Boolean) Indicates whether the volume should be encrypted. The effect of setting the encryption state to ``true`` depends on the volume origin (new or from a snapshot), starting encryption state, ownership, and whether encryption by default is enabled. For more information, see [Encryption by default](https://docs.aws.amazon.com/ebs/latest/userguide/work-with-ebs-encr.html#encryption-by-default) in the *Amazon EBS User Guide*.
 Encrypted Amazon EBS volumes must be attached to instances that support Amazon EBS encryption. For more information, see [Supported instance types](https://docs.aws.amazon.com/ebs/latest/userguide/ebs-encryption-requirements.html#ebs-encryption_supported_instances).
- `iops` (Number) The number of I/O operations per second (IOPS). For ``gp3``, ``io1``, and ``io2`` volumes, this represents the number of IOPS that are provisioned for the volume. For ``gp2`` volumes, this represents the baseline performance of the volume and the rate at which the volume accumulates I/O credits for bursting.
 The following are the supported values for each volume type:
  +  ``gp3``: 3,000 - 16,000 IOPS
  +  ``io1``: 100 - 64,000 IOPS
  +  ``io2``: 100 - 256,000 IOPS
  
 For ``io2`` volumes, you can achieve up to 256,000 IOPS on [instances built on the Nitro System](https://docs.aws.amazon.com/ec2/latest/instancetypes/ec2-nitro-instances.html). On other instances, you can achieve performance up to 32,000 IOPS.
 This parameter is required for ``io1`` and ``io2`` volumes. The default for ``gp3`` volumes is 3,000 IOPS. This parameter is not supported for ``gp2``, ``st1``, ``sc1``, or ``standard`` volumes.
- `kms_key_id` (String) The identifier of the kms-key-long to use for Amazon EBS encryption. If ``KmsKeyId`` is specified, the encrypted state must be ``true``.
 If you omit this property and your account is enabled for encryption by default, or *Encrypted* is set to ``true``, then the volume is encrypted using the default key specified for your account. If your account does not have a default key, then the volume is encrypted using the aws-managed-key.
 Alternatively, if you want to specify a different key, you can specify one of the following:
  +  Key ID. For example, 1234abcd-12ab-34cd-56ef-1234567890ab.
  +  Key alias. Specify the alias for the key, prefixed with ``alias/``. For example, for a key with the alias ``my_cmk``, use ``alias/my_cmk``. Or to specify the aws-managed-key, use ``alias/aws/ebs``.
  +  Key ARN. For example, arn:aws:kms:us-east-1:012345678910:key/1234abcd-12ab-34cd-56ef-1234567890ab.
  +  Alias ARN. For example, arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.
- `multi_attach_enabled` (Boolean) Indicates whether Amazon EBS Multi-Attach is enabled.
 CFNlong does not currently support updating a single-attach volume to be multi-attach enabled, updating a multi-attach enabled volume to be single-attach, or updating the size or number of I/O operations per second (IOPS) of a multi-attach enabled volume.
- `outpost_arn` (String) The Amazon Resource Name (ARN) of the Outpost.
- `size` (Number) The size of the volume, in GiBs. You must specify either a snapshot ID or a volume size. If you specify a snapshot, the default is the snapshot size. You can specify a volume size that is equal to or larger than the snapshot size.
 The following are the supported volumes sizes for each volume type:
  +  ``gp2`` and ``gp3``: 1 - 16,384 GiB
  +  ``io1``: 4 - 16,384 GiB
  +  ``io2``: 4 - 65,536 GiB
  +  ``st1`` and ``sc1``: 125 - 16,384 GiB
  +  ``standard``: 1 - 1024 GiB
- `snapshot_id` (String) The snapshot from which to create the volume. You must specify either a snapshot ID or a volume size.
- `tags` (Attributes List) The tags to apply to the volume during creation. (see [below for nested schema](#nestedatt--tags))
- `throughput` (Number) The throughput to provision for a volume, with a maximum of 1,000 MiB/s.
 This parameter is valid only for ``gp3`` volumes. The default value is 125.
 Valid Range: Minimum value of 125. Maximum value of 1000.
- `volume_initialization_rate` (Number)
- `volume_type` (String) The volume type. This parameter can be one of the following values:
  +  General Purpose SSD: ``gp2`` | ``gp3``
  +  Provisioned IOPS SSD: ``io1`` | ``io2``
  +  Throughput Optimized HDD: ``st1``
  +  Cold HDD: ``sc1``
  +  Magnetic: ``standard``
  
 For more information, see [Amazon EBS volume types](https://docs.aws.amazon.com/ebs/latest/userguide/ebs-volume-types.html).
 Default: ``gp2``

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `volume_id` (String)

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The tag key.
- `value` (String) The tag value.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_ec2_volume.example
  id = "volume_id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_ec2_volume.example "volume_id"
```