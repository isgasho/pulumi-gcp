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
    'GetAppEngineServiceResult',
    'AwaitableGetAppEngineServiceResult',
    'get_app_engine_service',
]

@pulumi.output_type
class GetAppEngineServiceResult:
    """
    A collection of values returned by getAppEngineService.
    """
    def __init__(__self__, display_name=None, id=None, module_id=None, name=None, project=None, service_id=None, telemetries=None):
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if module_id and not isinstance(module_id, str):
            raise TypeError("Expected argument 'module_id' to be a str")
        pulumi.set(__self__, "module_id", module_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if service_id and not isinstance(service_id, str):
            raise TypeError("Expected argument 'service_id' to be a str")
        pulumi.set(__self__, "service_id", service_id)
        if telemetries and not isinstance(telemetries, list):
            raise TypeError("Expected argument 'telemetries' to be a list")
        pulumi.set(__self__, "telemetries", telemetries)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="moduleId")
    def module_id(self) -> str:
        return pulumi.get(self, "module_id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> Optional[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> str:
        return pulumi.get(self, "service_id")

    @property
    @pulumi.getter
    def telemetries(self) -> Sequence['outputs.GetAppEngineServiceTelemetryResult']:
        return pulumi.get(self, "telemetries")


class AwaitableGetAppEngineServiceResult(GetAppEngineServiceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAppEngineServiceResult(
            display_name=self.display_name,
            id=self.id,
            module_id=self.module_id,
            name=self.name,
            project=self.project,
            service_id=self.service_id,
            telemetries=self.telemetries)


def get_app_engine_service(module_id: Optional[str] = None,
                           project: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAppEngineServiceResult:
    """
    A Monitoring Service is the root resource under which operational aspects of a
    generic service are accessible. A service is some discrete, autonomous, and
    network-accessible unit, designed to solve an individual concern

    An App Engine monitoring service is automatically created by GCP to monitor
    App Engine services.

    To get more information about Service, see:

    * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/services)
    * How-to Guides
        * [Service Monitoring](https://cloud.google.com/monitoring/service-monitoring)
        * [Monitoring API Documentation](https://cloud.google.com/monitoring/api/v3/)

    ## Example Usage
    ### Monitoring App Engine Service

    ```python
    import pulumi
    import pulumi_gcp as gcp

    bucket = gcp.storage.Bucket("bucket")
    object = gcp.storage.BucketObject("object",
        bucket=bucket.name,
        source=pulumi.FileAsset("./test-fixtures/appengine/hello-world.zip"))
    myapp = gcp.appengine.StandardAppVersion("myapp",
        version_id="v1",
        service="myapp",
        runtime="nodejs10",
        entrypoint=gcp.appengine.StandardAppVersionEntrypointArgs(
            shell="node ./app.js",
        ),
        deployment=gcp.appengine.StandardAppVersionDeploymentArgs(
            zip=gcp.appengine.StandardAppVersionDeploymentZipArgs(
                source_url=pulumi.Output.all(bucket.name, object.name).apply(lambda bucketName, objectName: f"https://storage.googleapis.com/{bucket_name}/{object_name}"),
            ),
        ),
        env_variables={
            "port": "8080",
        },
        delete_service_on_destroy=False)
    srv = myapp.service.apply(lambda service: gcp.monitoring.get_app_engine_service(module_id=service))
    ```


    :param str module_id: The ID of the App Engine module underlying this
           service. Corresponds to the moduleId resource label in the [gae_app](https://cloud.google.com/monitoring/api/resources#tag_gae_app) monitored resource, or the service/module name.
    :param str project: The ID of the project in which the resource belongs.
           If it is not provided, the provider project is used.
    """
    __args__ = dict()
    __args__['moduleId'] = module_id
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:monitoring/getAppEngineService:getAppEngineService', __args__, opts=opts, typ=GetAppEngineServiceResult).value

    return AwaitableGetAppEngineServiceResult(
        display_name=__ret__.display_name,
        id=__ret__.id,
        module_id=__ret__.module_id,
        name=__ret__.name,
        project=__ret__.project,
        service_id=__ret__.service_id,
        telemetries=__ret__.telemetries)
