# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GameServerConfig(pulumi.CustomResource):
    config_id: pulumi.Output[str]
    deployment_id: pulumi.Output[str]
    description: pulumi.Output[str]
    fleet_configs: pulumi.Output[list]
    labels: pulumi.Output[dict]
    location: pulumi.Output[str]
    name: pulumi.Output[str]
    project: pulumi.Output[str]
    scaling_configs: pulumi.Output[list]
    def __init__(__self__, resource_name, opts=None, config_id=None, deployment_id=None, description=None, fleet_configs=None, labels=None, location=None, project=None, scaling_configs=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a GameServerConfig resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        
        The **fleet_configs** object supports the following:
        
          * `fleetSpec` (`pulumi.Input[str]`)
          * `name` (`pulumi.Input[str]`)
        
        The **scaling_configs** object supports the following:
        
          * `fleetAutoscalerSpec` (`pulumi.Input[str]`)
          * `name` (`pulumi.Input[str]`)
          * `schedules` (`pulumi.Input[list]`)
        
            * `cronJobDuration` (`pulumi.Input[str]`)
            * `cronSpec` (`pulumi.Input[str]`)
            * `endTime` (`pulumi.Input[str]`)
            * `startTime` (`pulumi.Input[str]`)
        
          * `selectors` (`pulumi.Input[list]`)
        
            * `labels` (`pulumi.Input[dict]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/game_services_game_server_config.html.markdown.
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

            if config_id is None:
                raise TypeError("Missing required property 'config_id'")
            __props__['config_id'] = config_id
            if deployment_id is None:
                raise TypeError("Missing required property 'deployment_id'")
            __props__['deployment_id'] = deployment_id
            __props__['description'] = description
            if fleet_configs is None:
                raise TypeError("Missing required property 'fleet_configs'")
            __props__['fleet_configs'] = fleet_configs
            __props__['labels'] = labels
            __props__['location'] = location
            __props__['project'] = project
            __props__['scaling_configs'] = scaling_configs
            __props__['name'] = None
        super(GameServerConfig, __self__).__init__(
            'gcp:gameservices/gameServerConfig:GameServerConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, config_id=None, deployment_id=None, description=None, fleet_configs=None, labels=None, location=None, name=None, project=None, scaling_configs=None):
        """
        Get an existing GameServerConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        
        The **fleet_configs** object supports the following:
        
          * `fleetSpec` (`pulumi.Input[str]`)
          * `name` (`pulumi.Input[str]`)
        
        The **scaling_configs** object supports the following:
        
          * `fleetAutoscalerSpec` (`pulumi.Input[str]`)
          * `name` (`pulumi.Input[str]`)
          * `schedules` (`pulumi.Input[list]`)
        
            * `cronJobDuration` (`pulumi.Input[str]`)
            * `cronSpec` (`pulumi.Input[str]`)
            * `endTime` (`pulumi.Input[str]`)
            * `startTime` (`pulumi.Input[str]`)
        
          * `selectors` (`pulumi.Input[list]`)
        
            * `labels` (`pulumi.Input[dict]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/game_services_game_server_config.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["config_id"] = config_id
        __props__["deployment_id"] = deployment_id
        __props__["description"] = description
        __props__["fleet_configs"] = fleet_configs
        __props__["labels"] = labels
        __props__["location"] = location
        __props__["name"] = name
        __props__["project"] = project
        __props__["scaling_configs"] = scaling_configs
        return GameServerConfig(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

