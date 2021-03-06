# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['SharedVPCHostProject']


class SharedVPCHostProject(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Enables the Google Compute Engine
        [Shared VPC](https://cloud.google.com/compute/docs/shared-vpc)
        feature for a project, assigning it as a Shared VPC host project.

        For more information, see,
        [the Project API documentation](https://cloud.google.com/compute/docs/reference/latest/projects),
        where the Shared VPC feature is referred to by its former name "XPN".

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        # A host project provides network resources to associated service projects.
        host = gcp.compute.SharedVPCHostProject("host", project="host-project-id")
        # A service project gains access to network resources provided by its
        # associated host project.
        service1 = gcp.compute.SharedVPCServiceProject("service1",
            host_project=host.project,
            service_project="service-project-id-1")
        service2 = gcp.compute.SharedVPCServiceProject("service2",
            host_project=host.project,
            service_project="service-project-id-2")
        ```

        ## Import

        Google Compute Engine Shared VPC host project feature can be imported using the `project`, e.g.

        ```sh
         $ pulumi import gcp:compute/sharedVPCHostProject:SharedVPCHostProject host host-project-id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project that will serve as a Shared VPC host project
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

            if project is None:
                raise TypeError("Missing required property 'project'")
            __props__['project'] = project
        super(SharedVPCHostProject, __self__).__init__(
            'gcp:compute/sharedVPCHostProject:SharedVPCHostProject',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            project: Optional[pulumi.Input[str]] = None) -> 'SharedVPCHostProject':
        """
        Get an existing SharedVPCHostProject resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project that will serve as a Shared VPC host project
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["project"] = project
        return SharedVPCHostProject(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project that will serve as a Shared VPC host project
        """
        return pulumi.get(self, "project")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

