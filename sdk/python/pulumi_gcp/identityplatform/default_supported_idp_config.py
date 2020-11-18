# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['DefaultSupportedIdpConfig']


class DefaultSupportedIdpConfig(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 idp_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Configurations options for authenticating with a the standard set of Identity Toolkit-trusted IDPs.

        You must enable the
        [Google Identity Platform](https://console.cloud.google.com/marketplace/details/google-cloud-platform/customer-identity) in
        the marketplace prior to using this resource.

        ## Example Usage
        ### Identity Platform Default Supported Idp Config Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        idp_config = gcp.identityplatform.DefaultSupportedIdpConfig("idpConfig",
            client_id="client-id",
            client_secret="secret",
            enabled=True,
            idp_id="playgames.google.com")
        ```

        ## Import

        DefaultSupportedIdpConfig can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:identityplatform/defaultSupportedIdpConfig:DefaultSupportedIdpConfig default projects/{{project}}/defaultSupportedIdpConfigs/{{idp_id}}
        ```

        ```sh
         $ pulumi import gcp:identityplatform/defaultSupportedIdpConfig:DefaultSupportedIdpConfig default {{project}}/{{idp_id}}
        ```

        ```sh
         $ pulumi import gcp:identityplatform/defaultSupportedIdpConfig:DefaultSupportedIdpConfig default {{idp_id}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_id: OAuth client ID
        :param pulumi.Input[str] client_secret: OAuth client secret
        :param pulumi.Input[bool] enabled: If this IDP allows the user to sign in
        :param pulumi.Input[str] idp_id: ID of the IDP. Possible values include:
               * `apple.com`
               * `facebook.com`
               * `gc.apple.com`
               * `github.com`
               * `google.com`
               * `linkedin.com`
               * `microsoft.com`
               * `playgames.google.com`
               * `twitter.com`
               * `yahoo.com`
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
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

            if client_id is None:
                raise TypeError("Missing required property 'client_id'")
            __props__['client_id'] = client_id
            if client_secret is None:
                raise TypeError("Missing required property 'client_secret'")
            __props__['client_secret'] = client_secret
            __props__['enabled'] = enabled
            if idp_id is None:
                raise TypeError("Missing required property 'idp_id'")
            __props__['idp_id'] = idp_id
            __props__['project'] = project
            __props__['name'] = None
        super(DefaultSupportedIdpConfig, __self__).__init__(
            'gcp:identityplatform/defaultSupportedIdpConfig:DefaultSupportedIdpConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            client_id: Optional[pulumi.Input[str]] = None,
            client_secret: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            idp_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None) -> 'DefaultSupportedIdpConfig':
        """
        Get an existing DefaultSupportedIdpConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_id: OAuth client ID
        :param pulumi.Input[str] client_secret: OAuth client secret
        :param pulumi.Input[bool] enabled: If this IDP allows the user to sign in
        :param pulumi.Input[str] idp_id: ID of the IDP. Possible values include:
               * `apple.com`
               * `facebook.com`
               * `gc.apple.com`
               * `github.com`
               * `google.com`
               * `linkedin.com`
               * `microsoft.com`
               * `playgames.google.com`
               * `twitter.com`
               * `yahoo.com`
        :param pulumi.Input[str] name: The name of the DefaultSupportedIdpConfig resource
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["client_id"] = client_id
        __props__["client_secret"] = client_secret
        __props__["enabled"] = enabled
        __props__["idp_id"] = idp_id
        __props__["name"] = name
        __props__["project"] = project
        return DefaultSupportedIdpConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[str]:
        """
        OAuth client ID
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> pulumi.Output[str]:
        """
        OAuth client secret
        """
        return pulumi.get(self, "client_secret")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        If this IDP allows the user to sign in
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="idpId")
    def idp_id(self) -> pulumi.Output[str]:
        """
        ID of the IDP. Possible values include:
        * `apple.com`
        * `facebook.com`
        * `gc.apple.com`
        * `github.com`
        * `google.com`
        * `linkedin.com`
        * `microsoft.com`
        * `playgames.google.com`
        * `twitter.com`
        * `yahoo.com`
        """
        return pulumi.get(self, "idp_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the DefaultSupportedIdpConfig resource
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

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

