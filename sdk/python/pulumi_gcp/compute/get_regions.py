# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetRegionsResult:
    """
    A collection of values returned by getRegions.
    """
    def __init__(__self__, names=None, project=None, status=None, id=None):
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        __self__.names = names
        """
        A list of regions available in the given project
        """
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        __self__.project = project
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        __self__.status = status
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetRegionsResult(GetRegionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRegionsResult(
            names=self.names,
            project=self.project,
            status=self.status,
            id=self.id)

def get_regions(project=None,status=None,opts=None):
    """
    Provides access to available Google Compute regions for a given project.
    See more about [regions and regions](https://cloud.google.com/compute/docs/regions-zones/) in the upstream docs.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/compute_regions.html.markdown.
    """
    __args__ = dict()

    __args__['project'] = project
    __args__['status'] = status
    if opts is None:
        opts = pulumi.ResourceOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:compute/getRegions:getRegions', __args__, opts=opts).value

    return AwaitableGetRegionsResult(
        names=__ret__.get('names'),
        project=__ret__.get('project'),
        status=__ret__.get('status'),
        id=__ret__.get('id'))
