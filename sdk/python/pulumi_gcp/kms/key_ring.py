# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class KeyRing(pulumi.CustomResource):
    location: pulumi.Output[str]
    name: pulumi.Output[str]
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    self_link: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, location=None, name=None, project=None, __name__=None, __opts__=None):
        """
        A `KeyRing` is a toplevel logical grouping of `CryptoKeys`.
        
        
        > **Note:** KeyRings cannot be deleted from Google Cloud Platform.
        Destroying a Terraform-managed KeyRing will remove it from state but
        *will not delete the resource on the server.*
        
        
        To get more information about KeyRing, see:
        
        * [API documentation](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings)
        * How-to Guides
            * [Creating a key ring](https://cloud.google.com/kms/docs/creating-keys#create_a_key_ring)
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if location is None:
            raise TypeError("Missing required property 'location'")
        __props__['location'] = location

        __props__['name'] = name

        __props__['project'] = project

        __props__['self_link'] = None

        super(KeyRing, __self__).__init__(
            'gcp:kms/keyRing:KeyRing',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

