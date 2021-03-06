# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['WorkloadIdentityPool']


class WorkloadIdentityPool(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 workload_identity_pool_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ## Import

        WorkloadIdentityPool can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:iam/workloadIdentityPool:WorkloadIdentityPool default projects/{{project}}/locations/global/workloadIdentityPools/{{workload_identity_pool_id}}
        ```

        ```sh
         $ pulumi import gcp:iam/workloadIdentityPool:WorkloadIdentityPool default {{project}}/{{workload_identity_pool_id}}
        ```

        ```sh
         $ pulumi import gcp:iam/workloadIdentityPool:WorkloadIdentityPool default {{workload_identity_pool_id}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the pool. Cannot exceed 256 characters.
        :param pulumi.Input[bool] disabled: Whether the pool is disabled. You cannot use a disabled pool to exchange tokens, or use
               existing tokens to access resources. If the pool is re-enabled, existing tokens grant
               access again.
        :param pulumi.Input[str] display_name: A display name for the pool. Cannot exceed 32 characters.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] workload_identity_pool_id: The ID to use for the pool, which becomes the final component of the resource name. This
               value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
               `gcp-` is reserved for use by Google, and may not be specified.
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
            __props__['disabled'] = disabled
            __props__['display_name'] = display_name
            __props__['project'] = project
            if workload_identity_pool_id is None:
                raise TypeError("Missing required property 'workload_identity_pool_id'")
            __props__['workload_identity_pool_id'] = workload_identity_pool_id
            __props__['name'] = None
            __props__['state'] = None
        super(WorkloadIdentityPool, __self__).__init__(
            'gcp:iam/workloadIdentityPool:WorkloadIdentityPool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            disabled: Optional[pulumi.Input[bool]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None,
            workload_identity_pool_id: Optional[pulumi.Input[str]] = None) -> 'WorkloadIdentityPool':
        """
        Get an existing WorkloadIdentityPool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the pool. Cannot exceed 256 characters.
        :param pulumi.Input[bool] disabled: Whether the pool is disabled. You cannot use a disabled pool to exchange tokens, or use
               existing tokens to access resources. If the pool is re-enabled, existing tokens grant
               access again.
        :param pulumi.Input[str] display_name: A display name for the pool. Cannot exceed 32 characters.
        :param pulumi.Input[str] name: The resource name of the pool as
               'projects/{project_number}/locations/global/workloadIdentityPools/{workload_identity_pool_id}'.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] state: The state of the pool. * STATE_UNSPECIFIED: State unspecified. * ACTIVE: The pool is active, and may be used in Google
               Cloud policies. * DELETED: The pool is soft-deleted. Soft-deleted pools are permanently deleted after approximately 30
               days. You can restore a soft-deleted pool using UndeleteWorkloadIdentityPool. You cannot reuse the ID of a soft-deleted
               pool until it is permanently deleted. While a pool is deleted, you cannot use it to exchange tokens, or use existing
               tokens to access resources. If the pool is undeleted, existing tokens grant access again.
        :param pulumi.Input[str] workload_identity_pool_id: The ID to use for the pool, which becomes the final component of the resource name. This
               value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
               `gcp-` is reserved for use by Google, and may not be specified.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["disabled"] = disabled
        __props__["display_name"] = display_name
        __props__["name"] = name
        __props__["project"] = project
        __props__["state"] = state
        __props__["workload_identity_pool_id"] = workload_identity_pool_id
        return WorkloadIdentityPool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the pool. Cannot exceed 256 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the pool is disabled. You cannot use a disabled pool to exchange tokens, or use
        existing tokens to access resources. If the pool is re-enabled, existing tokens grant
        access again.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        """
        A display name for the pool. Cannot exceed 32 characters.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name of the pool as
        'projects/{project_number}/locations/global/workloadIdentityPools/{workload_identity_pool_id}'.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The state of the pool. * STATE_UNSPECIFIED: State unspecified. * ACTIVE: The pool is active, and may be used in Google
        Cloud policies. * DELETED: The pool is soft-deleted. Soft-deleted pools are permanently deleted after approximately 30
        days. You can restore a soft-deleted pool using UndeleteWorkloadIdentityPool. You cannot reuse the ID of a soft-deleted
        pool until it is permanently deleted. While a pool is deleted, you cannot use it to exchange tokens, or use existing
        tokens to access resources. If the pool is undeleted, existing tokens grant access again.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="workloadIdentityPoolId")
    def workload_identity_pool_id(self) -> pulumi.Output[str]:
        """
        The ID to use for the pool, which becomes the final component of the resource name. This
        value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
        `gcp-` is reserved for use by Google, and may not be specified.
        """
        return pulumi.get(self, "workload_identity_pool_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

