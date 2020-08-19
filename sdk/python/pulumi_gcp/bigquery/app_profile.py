# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class AppProfile(pulumi.CustomResource):
    app_profile_id: pulumi.Output[str]
    """
    The unique name of the app profile in the form `[_a-zA-Z0-9][-_.a-zA-Z0-9]*`.
    """
    description: pulumi.Output[str]
    """
    Long form description of the use case for this app profile.
    """
    ignore_warnings: pulumi.Output[bool]
    """
    If true, ignore safety checks when deleting/updating the app profile.
    """
    instance: pulumi.Output[str]
    """
    The name of the instance to create the app profile within.
    """
    multi_cluster_routing_use_any: pulumi.Output[bool]
    """
    If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest cluster that is available
    in the event of transient errors or delays. Clusters in a region are considered equidistant. Choosing this option sacrifices read-your-writes
    consistency to improve availability.
    """
    name: pulumi.Output[str]
    """
    The unique name of the requested app profile. Values are of the form
    'projects/<project>/instances/<instance>/appProfiles/<appProfileId>'.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    single_cluster_routing: pulumi.Output[dict]
    """
    Use a single-cluster routing policy.
    Structure is documented below.

      * `allowTransactionalWrites` (`bool`) - If true, CheckAndMutateRow and ReadModifyWriteRow requests are allowed by this app profile.
        It is unsafe to send these requests to the same table/row/column in multiple clusters.
      * `cluster_id` (`str`) - The cluster to which read/write requests should be routed.
    """
    def __init__(__self__, resource_name, opts=None, app_profile_id=None, description=None, ignore_warnings=None, instance=None, multi_cluster_routing_use_any=None, project=None, single_cluster_routing=None, __props__=None, __name__=None, __opts__=None):
        """
        App profile is a configuration object describing how Cloud Bigtable should treat traffic from a particular end user application.



        ## Example Usage

        ### Bigtable App Profile Multicluster

        ```python
        import pulumi
        import pulumi_gcp as gcp

        instance = gcp.bigtable.Instance("instance",
            cluster=[{
                "cluster_id": "bt-instance",
                "zone": "us-central1-b",
                "num_nodes": 3,
                "storageType": "HDD",
            }],
            deletion_protection="true")
        ap = gcp.bigquery.AppProfile("ap",
            instance=instance.name,
            app_profile_id="bt-profile",
            multi_cluster_routing_use_any=True,
            ignore_warnings=True)
        ```

        ### Bigtable App Profile Singlecluster

        ```python
        import pulumi
        import pulumi_gcp as gcp

        instance = gcp.bigtable.Instance("instance",
            cluster=[{
                "cluster_id": "bt-instance",
                "zone": "us-central1-b",
                "num_nodes": 3,
                "storageType": "HDD",
            }],
            deletion_protection="true")
        ap = gcp.bigquery.AppProfile("ap",
            instance=instance.name,
            app_profile_id="bt-profile",
            single_cluster_routing={
                "cluster_id": "bt-instance",
                "allowTransactionalWrites": True,
            },
            ignore_warnings=True)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_profile_id: The unique name of the app profile in the form `[_a-zA-Z0-9][-_.a-zA-Z0-9]*`.
        :param pulumi.Input[str] description: Long form description of the use case for this app profile.
        :param pulumi.Input[bool] ignore_warnings: If true, ignore safety checks when deleting/updating the app profile.
        :param pulumi.Input[str] instance: The name of the instance to create the app profile within.
        :param pulumi.Input[bool] multi_cluster_routing_use_any: If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest cluster that is available
               in the event of transient errors or delays. Clusters in a region are considered equidistant. Choosing this option sacrifices read-your-writes
               consistency to improve availability.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[dict] single_cluster_routing: Use a single-cluster routing policy.
               Structure is documented below.

        The **single_cluster_routing** object supports the following:

          * `allowTransactionalWrites` (`pulumi.Input[bool]`) - If true, CheckAndMutateRow and ReadModifyWriteRow requests are allowed by this app profile.
            It is unsafe to send these requests to the same table/row/column in multiple clusters.
          * `cluster_id` (`pulumi.Input[str]`) - The cluster to which read/write requests should be routed.
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

            if app_profile_id is None:
                raise TypeError("Missing required property 'app_profile_id'")
            __props__['app_profile_id'] = app_profile_id
            __props__['description'] = description
            __props__['ignore_warnings'] = ignore_warnings
            __props__['instance'] = instance
            __props__['multi_cluster_routing_use_any'] = multi_cluster_routing_use_any
            __props__['project'] = project
            __props__['single_cluster_routing'] = single_cluster_routing
            __props__['name'] = None
        super(AppProfile, __self__).__init__(
            'gcp:bigquery/appProfile:AppProfile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, app_profile_id=None, description=None, ignore_warnings=None, instance=None, multi_cluster_routing_use_any=None, name=None, project=None, single_cluster_routing=None):
        """
        Get an existing AppProfile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_profile_id: The unique name of the app profile in the form `[_a-zA-Z0-9][-_.a-zA-Z0-9]*`.
        :param pulumi.Input[str] description: Long form description of the use case for this app profile.
        :param pulumi.Input[bool] ignore_warnings: If true, ignore safety checks when deleting/updating the app profile.
        :param pulumi.Input[str] instance: The name of the instance to create the app profile within.
        :param pulumi.Input[bool] multi_cluster_routing_use_any: If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest cluster that is available
               in the event of transient errors or delays. Clusters in a region are considered equidistant. Choosing this option sacrifices read-your-writes
               consistency to improve availability.
        :param pulumi.Input[str] name: The unique name of the requested app profile. Values are of the form
               'projects/<project>/instances/<instance>/appProfiles/<appProfileId>'.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[dict] single_cluster_routing: Use a single-cluster routing policy.
               Structure is documented below.

        The **single_cluster_routing** object supports the following:

          * `allowTransactionalWrites` (`pulumi.Input[bool]`) - If true, CheckAndMutateRow and ReadModifyWriteRow requests are allowed by this app profile.
            It is unsafe to send these requests to the same table/row/column in multiple clusters.
          * `cluster_id` (`pulumi.Input[str]`) - The cluster to which read/write requests should be routed.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["app_profile_id"] = app_profile_id
        __props__["description"] = description
        __props__["ignore_warnings"] = ignore_warnings
        __props__["instance"] = instance
        __props__["multi_cluster_routing_use_any"] = multi_cluster_routing_use_any
        __props__["name"] = name
        __props__["project"] = project
        __props__["single_cluster_routing"] = single_cluster_routing
        return AppProfile(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

