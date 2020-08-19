# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class FolderFeed(pulumi.CustomResource):
    asset_names: pulumi.Output[list]
    """
    A list of the full names of the assets to receive updates. You must specify either or both of
    assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
    exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
    See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.
    """
    asset_types: pulumi.Output[list]
    """
    A list of types of the assets to receive updates. You must specify either or both of assetNames
    and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
    the feed. For example: "compute.googleapis.com/Disk"
    See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
    supported asset types.
    """
    billing_project: pulumi.Output[str]
    """
    The project whose identity will be used when sending messages to the
    destination pubsub topic. It also specifies the project for API
    enablement check, quota, and billing.
    """
    content_type: pulumi.Output[str]
    """
    Asset content type. If not specified, no content but the asset name and type will be returned.
    Possible values are `CONTENT_TYPE_UNSPECIFIED`, `RESOURCE`, `IAM_POLICY`, `ORG_POLICY`, and `ACCESS_POLICY`.
    """
    feed_id: pulumi.Output[str]
    """
    This is the client-assigned asset feed identifier and it needs to be unique under a specific parent.
    """
    feed_output_config: pulumi.Output[dict]
    """
    Output configuration for asset feed destination.
    Structure is documented below.

      * `pubsubDestination` (`dict`) - Destination on Cloud Pubsub.
        Structure is documented below.
        * `topic` (`str`) - Destination on Cloud Pubsub topic.
    """
    folder: pulumi.Output[str]
    """
    The folder this feed should be created in.
    """
    folder_id: pulumi.Output[str]
    """
    The ID of the folder where this feed has been created. Both [FOLDER_NUMBER] and folders/[FOLDER_NUMBER] are accepted.
    """
    name: pulumi.Output[str]
    """
    The format will be folders/{folder_number}/feeds/{client-assigned_feed_identifier}.
    """
    def __init__(__self__, resource_name, opts=None, asset_names=None, asset_types=None, billing_project=None, content_type=None, feed_id=None, feed_output_config=None, folder=None, __props__=None, __name__=None, __opts__=None):
        """
        Describes a Cloud Asset Inventory feed used to to listen to asset updates.


        To get more information about FolderFeed, see:

        * [API documentation](https://cloud.google.com/asset-inventory/docs/reference/rest/)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/asset-inventory/docs)


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] asset_names: A list of the full names of the assets to receive updates. You must specify either or both of
               assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
               exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
               See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.
        :param pulumi.Input[list] asset_types: A list of types of the assets to receive updates. You must specify either or both of assetNames
               and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
               the feed. For example: "compute.googleapis.com/Disk"
               See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
               supported asset types.
        :param pulumi.Input[str] billing_project: The project whose identity will be used when sending messages to the
               destination pubsub topic. It also specifies the project for API
               enablement check, quota, and billing.
        :param pulumi.Input[str] content_type: Asset content type. If not specified, no content but the asset name and type will be returned.
               Possible values are `CONTENT_TYPE_UNSPECIFIED`, `RESOURCE`, `IAM_POLICY`, `ORG_POLICY`, and `ACCESS_POLICY`.
        :param pulumi.Input[str] feed_id: This is the client-assigned asset feed identifier and it needs to be unique under a specific parent.
        :param pulumi.Input[dict] feed_output_config: Output configuration for asset feed destination.
               Structure is documented below.
        :param pulumi.Input[str] folder: The folder this feed should be created in.

        The **feed_output_config** object supports the following:

          * `pubsubDestination` (`pulumi.Input[dict]`) - Destination on Cloud Pubsub.
            Structure is documented below.
            * `topic` (`pulumi.Input[str]`) - Destination on Cloud Pubsub topic.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['asset_names'] = asset_names
            __props__['asset_types'] = asset_types
            if billing_project is None:
                raise TypeError("Missing required property 'billing_project'")
            __props__['billing_project'] = billing_project
            __props__['content_type'] = content_type
            if feed_id is None:
                raise TypeError("Missing required property 'feed_id'")
            __props__['feed_id'] = feed_id
            if feed_output_config is None:
                raise TypeError("Missing required property 'feed_output_config'")
            __props__['feed_output_config'] = feed_output_config
            if folder is None:
                raise TypeError("Missing required property 'folder'")
            __props__['folder'] = folder
            __props__['folder_id'] = None
            __props__['name'] = None
        super(FolderFeed, __self__).__init__(
            'gcp:cloudasset/folderFeed:FolderFeed',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, asset_names=None, asset_types=None, billing_project=None, content_type=None, feed_id=None, feed_output_config=None, folder=None, folder_id=None, name=None):
        """
        Get an existing FolderFeed resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] asset_names: A list of the full names of the assets to receive updates. You must specify either or both of
               assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
               exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
               See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.
        :param pulumi.Input[list] asset_types: A list of types of the assets to receive updates. You must specify either or both of assetNames
               and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
               the feed. For example: "compute.googleapis.com/Disk"
               See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
               supported asset types.
        :param pulumi.Input[str] billing_project: The project whose identity will be used when sending messages to the
               destination pubsub topic. It also specifies the project for API
               enablement check, quota, and billing.
        :param pulumi.Input[str] content_type: Asset content type. If not specified, no content but the asset name and type will be returned.
               Possible values are `CONTENT_TYPE_UNSPECIFIED`, `RESOURCE`, `IAM_POLICY`, `ORG_POLICY`, and `ACCESS_POLICY`.
        :param pulumi.Input[str] feed_id: This is the client-assigned asset feed identifier and it needs to be unique under a specific parent.
        :param pulumi.Input[dict] feed_output_config: Output configuration for asset feed destination.
               Structure is documented below.
        :param pulumi.Input[str] folder: The folder this feed should be created in.
        :param pulumi.Input[str] folder_id: The ID of the folder where this feed has been created. Both [FOLDER_NUMBER] and folders/[FOLDER_NUMBER] are accepted.
        :param pulumi.Input[str] name: The format will be folders/{folder_number}/feeds/{client-assigned_feed_identifier}.

        The **feed_output_config** object supports the following:

          * `pubsubDestination` (`pulumi.Input[dict]`) - Destination on Cloud Pubsub.
            Structure is documented below.
            * `topic` (`pulumi.Input[str]`) - Destination on Cloud Pubsub topic.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["asset_names"] = asset_names
        __props__["asset_types"] = asset_types
        __props__["billing_project"] = billing_project
        __props__["content_type"] = content_type
        __props__["feed_id"] = feed_id
        __props__["feed_output_config"] = feed_output_config
        __props__["folder"] = folder
        __props__["folder_id"] = folder_id
        __props__["name"] = name
        return FolderFeed(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

