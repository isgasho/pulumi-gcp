# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Disk']


class Disk(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disk_encryption_key: Optional[pulumi.Input[pulumi.InputType['DiskDiskEncryptionKeyArgs']]] = None,
                 image: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 physical_block_size_bytes: Optional[pulumi.Input[int]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 resource_policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 snapshot: Optional[pulumi.Input[str]] = None,
                 source_image_encryption_key: Optional[pulumi.Input[pulumi.InputType['DiskSourceImageEncryptionKeyArgs']]] = None,
                 source_snapshot_encryption_key: Optional[pulumi.Input[pulumi.InputType['DiskSourceSnapshotEncryptionKeyArgs']]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Persistent disks are durable storage devices that function similarly to
        the physical disks in a desktop or a server. Compute Engine manages the
        hardware behind these devices to ensure data redundancy and optimize
        performance for you. Persistent disks are available as either standard
        hard disk drives (HDD) or solid-state drives (SSD).

        Persistent disks are located independently from your virtual machine
        instances, so you can detach or move persistent disks to keep your data
        even after you delete your instances. Persistent disk performance scales
        automatically with size, so you can resize your existing persistent disks
        or add more persistent disks to an instance to meet your performance and
        storage space requirements.

        Add a persistent disk to your instance when you need reliable and
        affordable storage with consistent performance characteristics.

        To get more information about Disk, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/v1/disks)
        * How-to Guides
            * [Adding a persistent disk](https://cloud.google.com/compute/docs/disks/add-persistent-disk)

        > **Warning:** All arguments including `disk_encryption_key.raw_key` will be stored in the raw
        state as plain-text. [Read more about secrets in state](https://www.pulumi.com/docs/intro/concepts/programming-model/#secrets).

        ## Example Usage
        ### Disk Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.compute.Disk("default",
            image="debian-8-jessie-v20170523",
            labels={
                "environment": "dev",
            },
            physical_block_size_bytes=4096,
            type="pd-ssd",
            zone="us-central1-a")
        ```

        ## Import

        Disk can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/disk:Disk default projects/{{project}}/zones/{{zone}}/disks/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/disk:Disk default {{project}}/{{zone}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/disk:Disk default {{zone}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/disk:Disk default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when
               you create the resource.
        :param pulumi.Input[pulumi.InputType['DiskDiskEncryptionKeyArgs']] disk_encryption_key: Encrypts the disk using a customer-supplied encryption key.
               After you encrypt a disk with a customer-supplied key, you must
               provide the same key if you use the disk later (e.g. to create a disk
               snapshot or an image, or to attach the disk to a virtual machine).
               Customer-supplied encryption keys do not protect access to metadata of
               the disk.
               If you do not provide an encryption key when creating the disk, then
               the disk will be encrypted using an automatically generated key and
               you do not need to provide a key to use the disk later.
               Structure is documented below.
        :param pulumi.Input[str] image: The image from which to initialize this disk. This can be
               one of: the image's `self_link`, `projects/{project}/global/images/{image}`,
               `projects/{project}/global/images/family/{family}`, `global/images/{image}`,
               `global/images/family/{family}`, `family/{family}`, `{project}/{family}`,
               `{project}/{image}`, `{family}`, or `{image}`. If referred by family, the
               images names must include the family name. If they don't, use the
               [compute.Image data source](https://www.terraform.io/docs/providers/google/d/compute_image.html).
               For instance, the image `centos-6-v20180104` includes its family name `centos-6`.
               These images can be referred by family name here.
        :param pulumi.Input[str] interface: Specifies the disk interface to use for attaching this disk, which is either SCSI or NVME. The default is SCSI. Default
               value: "SCSI" Possible values: ["SCSI", "NVME"]
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels to apply to this disk.  A list of key->value pairs.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035. Specifically, the name must be 1-63 characters long and match
               the regular expression `a-z?` which means the
               first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
        :param pulumi.Input[int] physical_block_size_bytes: Physical block size of the persistent disk, in bytes. If not present
               in a request, a default value is used. Currently supported sizes
               are 4096 and 16384, other sizes may be added in the future.
               If an unsupported value is requested, the error message will list
               the supported values for the caller's project.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] resource_policies: Resource policies applied to this disk for automatic snapshot creations.
               ~>**NOTE** This value does not support updating the
               resource policy, as resource policies can not be updated more than
               one at a time. Use
               `compute.DiskResourcePolicyAttachment`
               to allow for updating the resource policy attached to the disk.
        :param pulumi.Input[int] size: Size of the persistent disk, specified in GB. You can specify this
               field when creating a persistent disk using the `image` or
               `snapshot` parameter, or specify it alone to create an empty
               persistent disk.
               If you specify this field along with `image` or `snapshot`,
               the value must not be less than the size of the image
               or the size of the snapshot.
        :param pulumi.Input[str] snapshot: The source snapshot used to create this disk. You can provide this as
               a partial or full URL to the resource. If the snapshot is in another
               project than this disk, you must supply a full URL. For example, the
               following are valid values:
               * `https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot`
               * `projects/project/global/snapshots/snapshot`
               * `global/snapshots/snapshot`
               * `snapshot`
        :param pulumi.Input[pulumi.InputType['DiskSourceImageEncryptionKeyArgs']] source_image_encryption_key: The customer-supplied encryption key of the source image. Required if
               the source image is protected by a customer-supplied encryption key.
               Structure is documented below.
        :param pulumi.Input[pulumi.InputType['DiskSourceSnapshotEncryptionKeyArgs']] source_snapshot_encryption_key: The customer-supplied encryption key of the source snapshot. Required
               if the source snapshot is protected by a customer-supplied encryption
               key.
               Structure is documented below.
        :param pulumi.Input[str] type: URL of the disk type resource describing which disk type to use to
               create the disk. Provide this when creating the disk.
        :param pulumi.Input[str] zone: A reference to the zone where the disk resides.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['description'] = description
            __props__['disk_encryption_key'] = disk_encryption_key
            __props__['image'] = image
            __props__['interface'] = interface
            __props__['labels'] = labels
            __props__['name'] = name
            __props__['physical_block_size_bytes'] = physical_block_size_bytes
            __props__['project'] = project
            __props__['resource_policies'] = resource_policies
            __props__['size'] = size
            __props__['snapshot'] = snapshot
            __props__['source_image_encryption_key'] = source_image_encryption_key
            __props__['source_snapshot_encryption_key'] = source_snapshot_encryption_key
            __props__['type'] = type
            __props__['zone'] = zone
            __props__['creation_timestamp'] = None
            __props__['label_fingerprint'] = None
            __props__['last_attach_timestamp'] = None
            __props__['last_detach_timestamp'] = None
            __props__['self_link'] = None
            __props__['source_image_id'] = None
            __props__['source_snapshot_id'] = None
            __props__['users'] = None
        super(Disk, __self__).__init__(
            'gcp:compute/disk:Disk',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            creation_timestamp: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            disk_encryption_key: Optional[pulumi.Input[pulumi.InputType['DiskDiskEncryptionKeyArgs']]] = None,
            image: Optional[pulumi.Input[str]] = None,
            interface: Optional[pulumi.Input[str]] = None,
            label_fingerprint: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            last_attach_timestamp: Optional[pulumi.Input[str]] = None,
            last_detach_timestamp: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            physical_block_size_bytes: Optional[pulumi.Input[int]] = None,
            project: Optional[pulumi.Input[str]] = None,
            resource_policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            self_link: Optional[pulumi.Input[str]] = None,
            size: Optional[pulumi.Input[int]] = None,
            snapshot: Optional[pulumi.Input[str]] = None,
            source_image_encryption_key: Optional[pulumi.Input[pulumi.InputType['DiskSourceImageEncryptionKeyArgs']]] = None,
            source_image_id: Optional[pulumi.Input[str]] = None,
            source_snapshot_encryption_key: Optional[pulumi.Input[pulumi.InputType['DiskSourceSnapshotEncryptionKeyArgs']]] = None,
            source_snapshot_id: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'Disk':
        """
        Get an existing Disk resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] creation_timestamp: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when
               you create the resource.
        :param pulumi.Input[pulumi.InputType['DiskDiskEncryptionKeyArgs']] disk_encryption_key: Encrypts the disk using a customer-supplied encryption key.
               After you encrypt a disk with a customer-supplied key, you must
               provide the same key if you use the disk later (e.g. to create a disk
               snapshot or an image, or to attach the disk to a virtual machine).
               Customer-supplied encryption keys do not protect access to metadata of
               the disk.
               If you do not provide an encryption key when creating the disk, then
               the disk will be encrypted using an automatically generated key and
               you do not need to provide a key to use the disk later.
               Structure is documented below.
        :param pulumi.Input[str] image: The image from which to initialize this disk. This can be
               one of: the image's `self_link`, `projects/{project}/global/images/{image}`,
               `projects/{project}/global/images/family/{family}`, `global/images/{image}`,
               `global/images/family/{family}`, `family/{family}`, `{project}/{family}`,
               `{project}/{image}`, `{family}`, or `{image}`. If referred by family, the
               images names must include the family name. If they don't, use the
               [compute.Image data source](https://www.terraform.io/docs/providers/google/d/compute_image.html).
               For instance, the image `centos-6-v20180104` includes its family name `centos-6`.
               These images can be referred by family name here.
        :param pulumi.Input[str] interface: Specifies the disk interface to use for attaching this disk, which is either SCSI or NVME. The default is SCSI. Default
               value: "SCSI" Possible values: ["SCSI", "NVME"]
        :param pulumi.Input[str] label_fingerprint: The fingerprint used for optimistic locking of this resource. Used internally during updates.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels to apply to this disk.  A list of key->value pairs.
        :param pulumi.Input[str] last_attach_timestamp: Last attach timestamp in RFC3339 text format.
        :param pulumi.Input[str] last_detach_timestamp: Last detach timestamp in RFC3339 text format.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035. Specifically, the name must be 1-63 characters long and match
               the regular expression `a-z?` which means the
               first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
        :param pulumi.Input[int] physical_block_size_bytes: Physical block size of the persistent disk, in bytes. If not present
               in a request, a default value is used. Currently supported sizes
               are 4096 and 16384, other sizes may be added in the future.
               If an unsupported value is requested, the error message will list
               the supported values for the caller's project.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] resource_policies: Resource policies applied to this disk for automatic snapshot creations.
               ~>**NOTE** This value does not support updating the
               resource policy, as resource policies can not be updated more than
               one at a time. Use
               `compute.DiskResourcePolicyAttachment`
               to allow for updating the resource policy attached to the disk.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[int] size: Size of the persistent disk, specified in GB. You can specify this
               field when creating a persistent disk using the `image` or
               `snapshot` parameter, or specify it alone to create an empty
               persistent disk.
               If you specify this field along with `image` or `snapshot`,
               the value must not be less than the size of the image
               or the size of the snapshot.
        :param pulumi.Input[str] snapshot: The source snapshot used to create this disk. You can provide this as
               a partial or full URL to the resource. If the snapshot is in another
               project than this disk, you must supply a full URL. For example, the
               following are valid values:
               * `https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot`
               * `projects/project/global/snapshots/snapshot`
               * `global/snapshots/snapshot`
               * `snapshot`
        :param pulumi.Input[pulumi.InputType['DiskSourceImageEncryptionKeyArgs']] source_image_encryption_key: The customer-supplied encryption key of the source image. Required if
               the source image is protected by a customer-supplied encryption key.
               Structure is documented below.
        :param pulumi.Input[str] source_image_id: The ID value of the image used to create this disk. This value identifies the exact image that was used to create this
               persistent disk. For example, if you created the persistent disk from an image that was later deleted and recreated
               under the same name, the source image ID would identify the exact version of the image that was used.
        :param pulumi.Input[pulumi.InputType['DiskSourceSnapshotEncryptionKeyArgs']] source_snapshot_encryption_key: The customer-supplied encryption key of the source snapshot. Required
               if the source snapshot is protected by a customer-supplied encryption
               key.
               Structure is documented below.
        :param pulumi.Input[str] source_snapshot_id: The unique ID of the snapshot used to create this disk. This value identifies the exact snapshot that was used to create
               this persistent disk. For example, if you created the persistent disk from a snapshot that was later deleted and
               recreated under the same name, the source snapshot ID would identify the exact version of the snapshot that was used.
        :param pulumi.Input[str] type: URL of the disk type resource describing which disk type to use to
               create the disk. Provide this when creating the disk.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] users: Links to the users of the disk (attached instances) in form: project/zones/zone/instances/instance
        :param pulumi.Input[str] zone: A reference to the zone where the disk resides.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["creation_timestamp"] = creation_timestamp
        __props__["description"] = description
        __props__["disk_encryption_key"] = disk_encryption_key
        __props__["image"] = image
        __props__["interface"] = interface
        __props__["label_fingerprint"] = label_fingerprint
        __props__["labels"] = labels
        __props__["last_attach_timestamp"] = last_attach_timestamp
        __props__["last_detach_timestamp"] = last_detach_timestamp
        __props__["name"] = name
        __props__["physical_block_size_bytes"] = physical_block_size_bytes
        __props__["project"] = project
        __props__["resource_policies"] = resource_policies
        __props__["self_link"] = self_link
        __props__["size"] = size
        __props__["snapshot"] = snapshot
        __props__["source_image_encryption_key"] = source_image_encryption_key
        __props__["source_image_id"] = source_image_id
        __props__["source_snapshot_encryption_key"] = source_snapshot_encryption_key
        __props__["source_snapshot_id"] = source_snapshot_id
        __props__["type"] = type
        __props__["users"] = users
        __props__["zone"] = zone
        return Disk(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> pulumi.Output[str]:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        An optional description of this resource. Provide this property when
        you create the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="diskEncryptionKey")
    def disk_encryption_key(self) -> pulumi.Output[Optional['outputs.DiskDiskEncryptionKey']]:
        """
        Encrypts the disk using a customer-supplied encryption key.
        After you encrypt a disk with a customer-supplied key, you must
        provide the same key if you use the disk later (e.g. to create a disk
        snapshot or an image, or to attach the disk to a virtual machine).
        Customer-supplied encryption keys do not protect access to metadata of
        the disk.
        If you do not provide an encryption key when creating the disk, then
        the disk will be encrypted using an automatically generated key and
        you do not need to provide a key to use the disk later.
        Structure is documented below.
        """
        return pulumi.get(self, "disk_encryption_key")

    @property
    @pulumi.getter
    def image(self) -> pulumi.Output[Optional[str]]:
        """
        The image from which to initialize this disk. This can be
        one of: the image's `self_link`, `projects/{project}/global/images/{image}`,
        `projects/{project}/global/images/family/{family}`, `global/images/{image}`,
        `global/images/family/{family}`, `family/{family}`, `{project}/{family}`,
        `{project}/{image}`, `{family}`, or `{image}`. If referred by family, the
        images names must include the family name. If they don't, use the
        [compute.Image data source](https://www.terraform.io/docs/providers/google/d/compute_image.html).
        For instance, the image `centos-6-v20180104` includes its family name `centos-6`.
        These images can be referred by family name here.
        """
        return pulumi.get(self, "image")

    @property
    @pulumi.getter
    def interface(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the disk interface to use for attaching this disk, which is either SCSI or NVME. The default is SCSI. Default
        value: "SCSI" Possible values: ["SCSI", "NVME"]
        """
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter(name="labelFingerprint")
    def label_fingerprint(self) -> pulumi.Output[str]:
        """
        The fingerprint used for optimistic locking of this resource. Used internally during updates.
        """
        return pulumi.get(self, "label_fingerprint")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Labels to apply to this disk.  A list of key->value pairs.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="lastAttachTimestamp")
    def last_attach_timestamp(self) -> pulumi.Output[str]:
        """
        Last attach timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "last_attach_timestamp")

    @property
    @pulumi.getter(name="lastDetachTimestamp")
    def last_detach_timestamp(self) -> pulumi.Output[str]:
        """
        Last detach timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "last_detach_timestamp")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the resource. Provided by the client when the resource is
        created. The name must be 1-63 characters long, and comply with
        RFC1035. Specifically, the name must be 1-63 characters long and match
        the regular expression `a-z?` which means the
        first character must be a lowercase letter, and all following
        characters must be a dash, lowercase letter, or digit, except the last
        character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="physicalBlockSizeBytes")
    def physical_block_size_bytes(self) -> pulumi.Output[int]:
        """
        Physical block size of the persistent disk, in bytes. If not present
        in a request, a default value is used. Currently supported sizes
        are 4096 and 16384, other sizes may be added in the future.
        If an unsupported value is requested, the error message will list
        the supported values for the caller's project.
        """
        return pulumi.get(self, "physical_block_size_bytes")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="resourcePolicies")
    def resource_policies(self) -> pulumi.Output[Sequence[str]]:
        """
        Resource policies applied to this disk for automatic snapshot creations.
        ~>**NOTE** This value does not support updating the
        resource policy, as resource policies can not be updated more than
        one at a time. Use
        `compute.DiskResourcePolicyAttachment`
        to allow for updating the resource policy attached to the disk.
        """
        return pulumi.get(self, "resource_policies")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        The URI of the created resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[int]:
        """
        Size of the persistent disk, specified in GB. You can specify this
        field when creating a persistent disk using the `image` or
        `snapshot` parameter, or specify it alone to create an empty
        persistent disk.
        If you specify this field along with `image` or `snapshot`,
        the value must not be less than the size of the image
        or the size of the snapshot.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def snapshot(self) -> pulumi.Output[Optional[str]]:
        """
        The source snapshot used to create this disk. You can provide this as
        a partial or full URL to the resource. If the snapshot is in another
        project than this disk, you must supply a full URL. For example, the
        following are valid values:
        * `https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot`
        * `projects/project/global/snapshots/snapshot`
        * `global/snapshots/snapshot`
        * `snapshot`
        """
        return pulumi.get(self, "snapshot")

    @property
    @pulumi.getter(name="sourceImageEncryptionKey")
    def source_image_encryption_key(self) -> pulumi.Output[Optional['outputs.DiskSourceImageEncryptionKey']]:
        """
        The customer-supplied encryption key of the source image. Required if
        the source image is protected by a customer-supplied encryption key.
        Structure is documented below.
        """
        return pulumi.get(self, "source_image_encryption_key")

    @property
    @pulumi.getter(name="sourceImageId")
    def source_image_id(self) -> pulumi.Output[str]:
        """
        The ID value of the image used to create this disk. This value identifies the exact image that was used to create this
        persistent disk. For example, if you created the persistent disk from an image that was later deleted and recreated
        under the same name, the source image ID would identify the exact version of the image that was used.
        """
        return pulumi.get(self, "source_image_id")

    @property
    @pulumi.getter(name="sourceSnapshotEncryptionKey")
    def source_snapshot_encryption_key(self) -> pulumi.Output[Optional['outputs.DiskSourceSnapshotEncryptionKey']]:
        """
        The customer-supplied encryption key of the source snapshot. Required
        if the source snapshot is protected by a customer-supplied encryption
        key.
        Structure is documented below.
        """
        return pulumi.get(self, "source_snapshot_encryption_key")

    @property
    @pulumi.getter(name="sourceSnapshotId")
    def source_snapshot_id(self) -> pulumi.Output[str]:
        """
        The unique ID of the snapshot used to create this disk. This value identifies the exact snapshot that was used to create
        this persistent disk. For example, if you created the persistent disk from a snapshot that was later deleted and
        recreated under the same name, the source snapshot ID would identify the exact version of the snapshot that was used.
        """
        return pulumi.get(self, "source_snapshot_id")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        URL of the disk type resource describing which disk type to use to
        create the disk. Provide this when creating the disk.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def users(self) -> pulumi.Output[Sequence[str]]:
        """
        Links to the users of the disk (attached instances) in form: project/zones/zone/instances/instance
        """
        return pulumi.get(self, "users")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        A reference to the zone where the disk resides.
        """
        return pulumi.get(self, "zone")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

