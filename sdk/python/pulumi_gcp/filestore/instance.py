# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Instance(pulumi.CustomResource):
    create_time: pulumi.Output[str]
    description: pulumi.Output[str]
    etag: pulumi.Output[str]
    file_shares: pulumi.Output[dict]
    labels: pulumi.Output[dict]
    name: pulumi.Output[str]
    networks: pulumi.Output[list]
    project: pulumi.Output[str]
    tier: pulumi.Output[str]
    zone: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, description=None, file_shares=None, labels=None, name=None, networks=None, project=None, tier=None, zone=None, __props__=None, __name__=None, __opts__=None):
        """
        A Google Cloud Filestore instance.
        
        
        To get more information about Instance, see:
        
        * [API documentation](https://cloud.google.com/filestore/docs/reference/rest/v1beta1/projects.locations.instances/create)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/filestore/docs/creating-instances)
            * [Use with Kubernetes](https://cloud.google.com/filestore/docs/accessing-fileshares)
            * [Copying Data In/Out](https://cloud.google.com/filestore/docs/copying-data)
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/filestore_instance.html.markdown.
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

            __props__['description'] = description
            if file_shares is None:
                raise TypeError("Missing required property 'file_shares'")
            __props__['file_shares'] = file_shares
            __props__['labels'] = labels
            __props__['name'] = name
            if networks is None:
                raise TypeError("Missing required property 'networks'")
            __props__['networks'] = networks
            __props__['project'] = project
            if tier is None:
                raise TypeError("Missing required property 'tier'")
            __props__['tier'] = tier
            if zone is None:
                raise TypeError("Missing required property 'zone'")
            __props__['zone'] = zone
            __props__['create_time'] = None
            __props__['etag'] = None
        super(Instance, __self__).__init__(
            'gcp:filestore/instance:Instance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, create_time=None, description=None, etag=None, file_shares=None, labels=None, name=None, networks=None, project=None, tier=None, zone=None):
        """
        Get an existing Instance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/filestore_instance.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["create_time"] = create_time
        __props__["description"] = description
        __props__["etag"] = etag
        __props__["file_shares"] = file_shares
        __props__["labels"] = labels
        __props__["name"] = name
        __props__["networks"] = networks
        __props__["project"] = project
        __props__["tier"] = tier
        __props__["zone"] = zone
        return Instance(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

