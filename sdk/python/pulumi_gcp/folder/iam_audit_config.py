# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class IamAuditConfig(pulumi.CustomResource):
    audit_log_configs: pulumi.Output[list]
    """
    The configuration for logging of each type of permission. This can be specified multiple times.

      * `exemptedMembers` (`list`)
      * `logType` (`str`)
    """
    etag: pulumi.Output[str]
    """
    The etag of iam policy
    """
    folder: pulumi.Output[str]
    service: pulumi.Output[str]
    """
    Service which will be enabled for audit logging. The special value allServices covers all services.
    """
    def __init__(__self__, resource_name, opts=None, audit_log_configs=None, folder=None, service=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a IamAuditConfig resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] audit_log_configs: The configuration for logging of each type of permission. This can be specified multiple times.
        :param pulumi.Input[str] service: Service which will be enabled for audit logging. The special value allServices covers all services.

        The **audit_log_configs** object supports the following:

          * `exemptedMembers` (`pulumi.Input[list]`)
          * `logType` (`pulumi.Input[str]`)
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

            if audit_log_configs is None:
                raise TypeError("Missing required property 'audit_log_configs'")
            __props__['audit_log_configs'] = audit_log_configs
            if folder is None:
                raise TypeError("Missing required property 'folder'")
            __props__['folder'] = folder
            if service is None:
                raise TypeError("Missing required property 'service'")
            __props__['service'] = service
            __props__['etag'] = None
        super(IamAuditConfig, __self__).__init__(
            'gcp:folder/iamAuditConfig:IamAuditConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, audit_log_configs=None, etag=None, folder=None, service=None):
        """
        Get an existing IamAuditConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] audit_log_configs: The configuration for logging of each type of permission. This can be specified multiple times.
        :param pulumi.Input[str] etag: The etag of iam policy
        :param pulumi.Input[str] service: Service which will be enabled for audit logging. The special value allServices covers all services.

        The **audit_log_configs** object supports the following:

          * `exemptedMembers` (`pulumi.Input[list]`)
          * `logType` (`pulumi.Input[str]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["audit_log_configs"] = audit_log_configs
        __props__["etag"] = etag
        __props__["folder"] = folder
        __props__["service"] = service
        return IamAuditConfig(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

