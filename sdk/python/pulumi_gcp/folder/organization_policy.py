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

__all__ = ['OrganizationPolicy']


class OrganizationPolicy(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 boolean_policy: Optional[pulumi.Input[pulumi.InputType['OrganizationPolicyBooleanPolicyArgs']]] = None,
                 constraint: Optional[pulumi.Input[str]] = None,
                 folder: Optional[pulumi.Input[str]] = None,
                 list_policy: Optional[pulumi.Input[pulumi.InputType['OrganizationPolicyListPolicyArgs']]] = None,
                 restore_policy: Optional[pulumi.Input[pulumi.InputType['OrganizationPolicyRestorePolicyArgs']]] = None,
                 version: Optional[pulumi.Input[int]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Allows management of Organization policies for a Google Folder. For more information see
        [the official
        documentation](https://cloud.google.com/resource-manager/docs/organization-policy/overview) and
        [API](https://cloud.google.com/resource-manager/reference/rest/v1/folders/setOrgPolicy).

        ## Example Usage

        To set policy with a [boolean constraint](https://cloud.google.com/resource-manager/docs/organization-policy/quickstart-boolean-constraints):

        ```python
        import pulumi
        import pulumi_gcp as gcp

        serial_port_policy = gcp.folder.OrganizationPolicy("serialPortPolicy",
            boolean_policy=gcp.folder.OrganizationPolicyBooleanPolicyArgs(
                enforced=True,
            ),
            constraint="compute.disableSerialPortAccess",
            folder="folders/123456789")
        ```

        To set a policy with a [list constraint](https://cloud.google.com/resource-manager/docs/organization-policy/quickstart-list-constraints):

        ```python
        import pulumi
        import pulumi_gcp as gcp

        services_policy = gcp.folder.OrganizationPolicy("servicesPolicy",
            constraint="serviceuser.services",
            folder="folders/123456789",
            list_policy=gcp.folder.OrganizationPolicyListPolicyArgs(
                allow=gcp.folder.OrganizationPolicyListPolicyAllowArgs(
                    all=True,
                ),
            ))
        ```

        Or to deny some services, use the following instead:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        services_policy = gcp.folder.OrganizationPolicy("servicesPolicy",
            constraint="serviceuser.services",
            folder="folders/123456789",
            list_policy=gcp.folder.OrganizationPolicyListPolicyArgs(
                deny=gcp.folder.OrganizationPolicyListPolicyDenyArgs(
                    values=["cloudresourcemanager.googleapis.com"],
                ),
                suggested_value="compute.googleapis.com",
            ))
        ```

        To restore the default folder organization policy, use the following instead:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        services_policy = gcp.folder.OrganizationPolicy("servicesPolicy",
            constraint="serviceuser.services",
            folder="folders/123456789",
            restore_policy=gcp.folder.OrganizationPolicyRestorePolicyArgs(
                default=True,
            ))
        ```

        ## Import

        Folder organization policies can be imported using any of the follow formats

        ```sh
         $ pulumi import gcp:folder/organizationPolicy:OrganizationPolicy policy folders/folder-1234/constraints/serviceuser.services
        ```

        ```sh
         $ pulumi import gcp:folder/organizationPolicy:OrganizationPolicy policy folder-1234/serviceuser.services
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['OrganizationPolicyBooleanPolicyArgs']] boolean_policy: A boolean policy is a constraint that is either enforced or not. Structure is documented below.
        :param pulumi.Input[str] constraint: The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
        :param pulumi.Input[str] folder: The resource name of the folder to set the policy for. Its format is folders/{folder_id}.
        :param pulumi.Input[pulumi.InputType['OrganizationPolicyListPolicyArgs']] list_policy: A policy that can define specific values that are allowed or denied for the given constraint. It
               can also be used to allow or deny all values. Structure is documented below.
        :param pulumi.Input[pulumi.InputType['OrganizationPolicyRestorePolicyArgs']] restore_policy: A restore policy is a constraint to restore the default policy. Structure is documented below.
        :param pulumi.Input[int] version: Version of the Policy. Default version is 0.
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

            __props__['boolean_policy'] = boolean_policy
            if constraint is None:
                raise TypeError("Missing required property 'constraint'")
            __props__['constraint'] = constraint
            if folder is None:
                raise TypeError("Missing required property 'folder'")
            __props__['folder'] = folder
            __props__['list_policy'] = list_policy
            __props__['restore_policy'] = restore_policy
            __props__['version'] = version
            __props__['etag'] = None
            __props__['update_time'] = None
        super(OrganizationPolicy, __self__).__init__(
            'gcp:folder/organizationPolicy:OrganizationPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            boolean_policy: Optional[pulumi.Input[pulumi.InputType['OrganizationPolicyBooleanPolicyArgs']]] = None,
            constraint: Optional[pulumi.Input[str]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            folder: Optional[pulumi.Input[str]] = None,
            list_policy: Optional[pulumi.Input[pulumi.InputType['OrganizationPolicyListPolicyArgs']]] = None,
            restore_policy: Optional[pulumi.Input[pulumi.InputType['OrganizationPolicyRestorePolicyArgs']]] = None,
            update_time: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[int]] = None) -> 'OrganizationPolicy':
        """
        Get an existing OrganizationPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['OrganizationPolicyBooleanPolicyArgs']] boolean_policy: A boolean policy is a constraint that is either enforced or not. Structure is documented below.
        :param pulumi.Input[str] constraint: The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
        :param pulumi.Input[str] etag: (Computed) The etag of the organization policy. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.
        :param pulumi.Input[str] folder: The resource name of the folder to set the policy for. Its format is folders/{folder_id}.
        :param pulumi.Input[pulumi.InputType['OrganizationPolicyListPolicyArgs']] list_policy: A policy that can define specific values that are allowed or denied for the given constraint. It
               can also be used to allow or deny all values. Structure is documented below.
        :param pulumi.Input[pulumi.InputType['OrganizationPolicyRestorePolicyArgs']] restore_policy: A restore policy is a constraint to restore the default policy. Structure is documented below.
        :param pulumi.Input[str] update_time: (Computed) The timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds, representing when the variable was last updated. Example: "2016-10-09T12:33:37.578138407Z".
        :param pulumi.Input[int] version: Version of the Policy. Default version is 0.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["boolean_policy"] = boolean_policy
        __props__["constraint"] = constraint
        __props__["etag"] = etag
        __props__["folder"] = folder
        __props__["list_policy"] = list_policy
        __props__["restore_policy"] = restore_policy
        __props__["update_time"] = update_time
        __props__["version"] = version
        return OrganizationPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="booleanPolicy")
    def boolean_policy(self) -> pulumi.Output[Optional['outputs.OrganizationPolicyBooleanPolicy']]:
        """
        A boolean policy is a constraint that is either enforced or not. Structure is documented below.
        """
        return pulumi.get(self, "boolean_policy")

    @property
    @pulumi.getter
    def constraint(self) -> pulumi.Output[str]:
        """
        The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
        """
        return pulumi.get(self, "constraint")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        (Computed) The etag of the organization policy. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def folder(self) -> pulumi.Output[str]:
        """
        The resource name of the folder to set the policy for. Its format is folders/{folder_id}.
        """
        return pulumi.get(self, "folder")

    @property
    @pulumi.getter(name="listPolicy")
    def list_policy(self) -> pulumi.Output[Optional['outputs.OrganizationPolicyListPolicy']]:
        """
        A policy that can define specific values that are allowed or denied for the given constraint. It
        can also be used to allow or deny all values. Structure is documented below.
        """
        return pulumi.get(self, "list_policy")

    @property
    @pulumi.getter(name="restorePolicy")
    def restore_policy(self) -> pulumi.Output[Optional['outputs.OrganizationPolicyRestorePolicy']]:
        """
        A restore policy is a constraint to restore the default policy. Structure is documented below.
        """
        return pulumi.get(self, "restore_policy")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        (Computed) The timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds, representing when the variable was last updated. Example: "2016-10-09T12:33:37.578138407Z".
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[int]:
        """
        Version of the Policy. Default version is 0.
        """
        return pulumi.get(self, "version")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

