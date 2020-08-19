# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetRegionInstanceGroupResult:
    """
    A collection of values returned by getRegionInstanceGroup.
    """
    def __init__(__self__, id=None, instances=None, name=None, project=None, region=None, self_link=None, size=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if instances and not isinstance(instances, list):
            raise TypeError("Expected argument 'instances' to be a list")
        __self__.instances = instances
        """
        List of instances in the group, as a list of resources, each containing:
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        String port name
        """
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        __self__.project = project
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        __self__.region = region
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        __self__.self_link = self_link
        if size and not isinstance(size, float):
            raise TypeError("Expected argument 'size' to be a float")
        __self__.size = size
        """
        The number of instances in the group.
        """
class AwaitableGetRegionInstanceGroupResult(GetRegionInstanceGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRegionInstanceGroupResult(
            id=self.id,
            instances=self.instances,
            name=self.name,
            project=self.project,
            region=self.region,
            self_link=self.self_link,
            size=self.size)

def get_region_instance_group(name=None,project=None,region=None,self_link=None,opts=None):
    """
    Get a Compute Region Instance Group within GCE.
    For more information, see [the official documentation](https://cloud.google.com/compute/docs/instance-groups/distributing-instances-with-regional-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/latest/regionInstanceGroups).

    ```python
    import pulumi
    import pulumi_gcp as gcp

    group = gcp.compute.get_region_instance_group(name="instance-group-name")
    ```

    The most common use of this datasource will be to fetch information about the instances inside regional managed instance groups, for instance:



    :param str name: The name of the instance group.  One of `name` or `self_link` must be provided.
    :param str project: The ID of the project in which the resource belongs.
           If `self_link` is provided, this value is ignored.  If neither `self_link`
           nor `project` are provided, the provider project is used.
    :param str region: The region in which the resource belongs.  If `self_link`
           is provided, this value is ignored.  If neither `self_link` nor `region` are
           provided, the provider region is used.
    :param str self_link: The link to the instance group.  One of `name` or `self_link` must be provided.
    """
    __args__ = dict()


    __args__['name'] = name
    __args__['project'] = project
    __args__['region'] = region
    __args__['selfLink'] = self_link
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:compute/getRegionInstanceGroup:getRegionInstanceGroup', __args__, opts=opts).value

    return AwaitableGetRegionInstanceGroupResult(
        id=__ret__.get('id'),
        instances=__ret__.get('instances'),
        name=__ret__.get('name'),
        project=__ret__.get('project'),
        region=__ret__.get('region'),
        self_link=__ret__.get('selfLink'),
        size=__ret__.get('size'))
