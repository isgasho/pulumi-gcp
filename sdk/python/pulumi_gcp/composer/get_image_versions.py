# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'GetImageVersionsResult',
    'AwaitableGetImageVersionsResult',
    'get_image_versions',
]

@pulumi.output_type
class GetImageVersionsResult:
    """
    A collection of values returned by getImageVersions.
    """
    def __init__(__self__, id=None, image_versions=None, project=None, region=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if image_versions and not isinstance(image_versions, list):
            raise TypeError("Expected argument 'image_versions' to be a list")
        pulumi.set(__self__, "image_versions", image_versions)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="imageVersions")
    def image_versions(self) -> Sequence['outputs.GetImageVersionsImageVersionResult']:
        """
        A list of composer image versions available in the given project and location. Each `image_version` contains:
        """
        return pulumi.get(self, "image_versions")

    @property
    @pulumi.getter
    def project(self) -> str:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")


class AwaitableGetImageVersionsResult(GetImageVersionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetImageVersionsResult(
            id=self.id,
            image_versions=self.image_versions,
            project=self.project,
            region=self.region)


def get_image_versions(project: Optional[str] = None,
                       region: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetImageVersionsResult:
    """
    Provides access to available Cloud Composer versions in a region for a given project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    all = gcp.composer.get_image_versions()
    test = gcp.composer.Environment("test",
        region="us-central1",
        config=gcp.composer.EnvironmentConfigArgs(
            software_config=gcp.composer.EnvironmentConfigSoftwareConfigArgs(
                image_version=all.image_versions[0].image_version_id,
            ),
        ))
    ```


    :param str project: The ID of the project to list versions in.
           If it is not provided, the provider project is used.
    :param str region: The location to list versions in.
           If it is not provider, the provider region is used.
    """
    __args__ = dict()
    __args__['project'] = project
    __args__['region'] = region
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:composer/getImageVersions:getImageVersions', __args__, opts=opts, typ=GetImageVersionsResult).value

    return AwaitableGetImageVersionsResult(
        id=__ret__.id,
        image_versions=__ret__.image_versions,
        project=__ret__.project,
        region=__ret__.region)
