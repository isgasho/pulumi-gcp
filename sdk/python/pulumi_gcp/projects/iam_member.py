# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class IAMMember(pulumi.CustomResource):
    condition: pulumi.Output[dict]
    """
    An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
    Structure is documented below.

      * `description` (`str`)
      * `expression` (`str`)
      * `title` (`str`)
    """
    etag: pulumi.Output[str]
    """
    (Computed) The etag of the project's IAM policy.
    """
    member: pulumi.Output[str]
    project: pulumi.Output[str]
    """
    The project ID. If not specified for `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig`, uses the ID of the project configured with the provider.
    Required for `projects.IAMPolicy` - you must explicitly set the project, and it
    will not be inferred from the provider.
    """
    role: pulumi.Output[str]
    """
    The role that should be applied. Only one
    `projects.IAMBinding` can be used per role. Note that custom roles must be of the format
    `[projects|organizations]/{parent-name}/roles/{role-name}`.
    """
    def __init__(__self__, resource_name, opts=None, condition=None, member=None, project=None, role=None, __props__=None, __name__=None, __opts__=None):
        """
        Four different resources help you manage your IAM policy for a project. Each of these resources serves a different use case:

        * `projects.IAMPolicy`: Authoritative. Sets the IAM policy for the project and replaces any existing policy already attached.
        * `projects.IAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the project are preserved.
        * `projects.IAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the project are preserved.
        * `projects.IAMAuditConfig`: Authoritative for a given service. Updates the IAM policy to enable audit logging for the given service.


        > **Note:** `projects.IAMPolicy` **cannot** be used in conjunction with `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig` or they will fight over what your policy should be.

        > **Note:** `projects.IAMBinding` resources **can be** used in conjunction with `projects.IAMMember` resources **only if** they do not grant privilege to the same role.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/google_project_iam.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] condition: An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
               Structure is documented below.
        :param pulumi.Input[str] project: The project ID. If not specified for `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig`, uses the ID of the project configured with the provider.
               Required for `projects.IAMPolicy` - you must explicitly set the project, and it
               will not be inferred from the provider.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `projects.IAMBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.

        The **condition** object supports the following:

          * `description` (`pulumi.Input[str]`)
          * `expression` (`pulumi.Input[str]`)
          * `title` (`pulumi.Input[str]`)
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

            __props__['condition'] = condition
            if member is None:
                raise TypeError("Missing required property 'member'")
            __props__['member'] = member
            __props__['project'] = project
            if role is None:
                raise TypeError("Missing required property 'role'")
            __props__['role'] = role
            __props__['etag'] = None
        super(IAMMember, __self__).__init__(
            'gcp:projects/iAMMember:IAMMember',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, condition=None, etag=None, member=None, project=None, role=None):
        """
        Get an existing IAMMember resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] condition: An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
               Structure is documented below.
        :param pulumi.Input[str] etag: (Computed) The etag of the project's IAM policy.
        :param pulumi.Input[str] project: The project ID. If not specified for `projects.IAMBinding`, `projects.IAMMember`, or `projects.IAMAuditConfig`, uses the ID of the project configured with the provider.
               Required for `projects.IAMPolicy` - you must explicitly set the project, and it
               will not be inferred from the provider.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `projects.IAMBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.

        The **condition** object supports the following:

          * `description` (`pulumi.Input[str]`)
          * `expression` (`pulumi.Input[str]`)
          * `title` (`pulumi.Input[str]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["condition"] = condition
        __props__["etag"] = etag
        __props__["member"] = member
        __props__["project"] = project
        __props__["role"] = role
        return IAMMember(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

