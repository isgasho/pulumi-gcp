# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Application']


class Application(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_domain: Optional[pulumi.Input[str]] = None,
                 database_type: Optional[pulumi.Input[str]] = None,
                 feature_settings: Optional[pulumi.Input[pulumi.InputType['ApplicationFeatureSettingsArgs']]] = None,
                 iap: Optional[pulumi.Input[pulumi.InputType['ApplicationIapArgs']]] = None,
                 location_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 serving_status: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Allows creation and management of an App Engine application.

        > App Engine applications cannot be deleted once they're created; you have to delete the
           entire project to delete the application. This provider will report the application has been
           successfully deleted; this is a limitation of the provider, and will go away in the future.
           This provider is not able to delete App Engine applications.

        > **Warning:** All arguments including `iap.oauth2_client_secret` will be stored in the raw
        state as plain-text. [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        my_project = gcp.organizations.Project("myProject",
            project_id="your-project-id",
            org_id="1234567")
        app = gcp.appengine.Application("app",
            project=my_project.project_id,
            location_id="us-central")
        ```

        ## Import

        Applications can be imported using the ID of the project the application belongs to, e.g.

        ```sh
         $ pulumi import gcp:appengine/application:Application app your-project-id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_domain: The domain to authenticate users with when using App Engine's User API.
        :param pulumi.Input[str] database_type: The type of the Cloud Firestore or Cloud Datastore database associated with this application.
               Can be `CLOUD_FIRESTORE` or `CLOUD_DATASTORE_COMPATIBILITY` for new
               instances.  To support old instances, the value `CLOUD_DATASTORE` is accepted
               by the provider, but will be rejected by the API.
        :param pulumi.Input[pulumi.InputType['ApplicationFeatureSettingsArgs']] feature_settings: A block of optional settings to configure specific App Engine features:
        :param pulumi.Input[pulumi.InputType['ApplicationIapArgs']] iap: Settings for enabling Cloud Identity Aware Proxy
        :param pulumi.Input[str] location_id: The [location](https://cloud.google.com/appengine/docs/locations)
               to serve the app from.
        :param pulumi.Input[str] project: The project ID to create the application under.
               ~>**NOTE:** GCP only accepts project ID, not project number. If you are using number,
               you may get a "Permission denied" error.
        :param pulumi.Input[str] serving_status: The serving status of the app.
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

            __props__['auth_domain'] = auth_domain
            __props__['database_type'] = database_type
            __props__['feature_settings'] = feature_settings
            __props__['iap'] = iap
            if location_id is None:
                raise TypeError("Missing required property 'location_id'")
            __props__['location_id'] = location_id
            __props__['project'] = project
            __props__['serving_status'] = serving_status
            __props__['app_id'] = None
            __props__['code_bucket'] = None
            __props__['default_bucket'] = None
            __props__['default_hostname'] = None
            __props__['gcr_domain'] = None
            __props__['name'] = None
            __props__['url_dispatch_rules'] = None
        super(Application, __self__).__init__(
            'gcp:appengine/application:Application',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            app_id: Optional[pulumi.Input[str]] = None,
            auth_domain: Optional[pulumi.Input[str]] = None,
            code_bucket: Optional[pulumi.Input[str]] = None,
            database_type: Optional[pulumi.Input[str]] = None,
            default_bucket: Optional[pulumi.Input[str]] = None,
            default_hostname: Optional[pulumi.Input[str]] = None,
            feature_settings: Optional[pulumi.Input[pulumi.InputType['ApplicationFeatureSettingsArgs']]] = None,
            gcr_domain: Optional[pulumi.Input[str]] = None,
            iap: Optional[pulumi.Input[pulumi.InputType['ApplicationIapArgs']]] = None,
            location_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            serving_status: Optional[pulumi.Input[str]] = None,
            url_dispatch_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationUrlDispatchRuleArgs']]]]] = None) -> 'Application':
        """
        Get an existing Application resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: Identifier of the app, usually `{PROJECT_ID}`
        :param pulumi.Input[str] auth_domain: The domain to authenticate users with when using App Engine's User API.
        :param pulumi.Input[str] code_bucket: The GCS bucket code is being stored in for this app.
        :param pulumi.Input[str] database_type: The type of the Cloud Firestore or Cloud Datastore database associated with this application.
               Can be `CLOUD_FIRESTORE` or `CLOUD_DATASTORE_COMPATIBILITY` for new
               instances.  To support old instances, the value `CLOUD_DATASTORE` is accepted
               by the provider, but will be rejected by the API.
        :param pulumi.Input[str] default_bucket: The GCS bucket content is being stored in for this app.
        :param pulumi.Input[str] default_hostname: The default hostname for this app.
        :param pulumi.Input[pulumi.InputType['ApplicationFeatureSettingsArgs']] feature_settings: A block of optional settings to configure specific App Engine features:
        :param pulumi.Input[str] gcr_domain: The GCR domain used for storing managed Docker images for this app.
        :param pulumi.Input[pulumi.InputType['ApplicationIapArgs']] iap: Settings for enabling Cloud Identity Aware Proxy
        :param pulumi.Input[str] location_id: The [location](https://cloud.google.com/appengine/docs/locations)
               to serve the app from.
        :param pulumi.Input[str] name: Unique name of the app, usually `apps/{PROJECT_ID}`
        :param pulumi.Input[str] project: The project ID to create the application under.
               ~>**NOTE:** GCP only accepts project ID, not project number. If you are using number,
               you may get a "Permission denied" error.
        :param pulumi.Input[str] serving_status: The serving status of the app.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationUrlDispatchRuleArgs']]]] url_dispatch_rules: A list of dispatch rule blocks. Each block has a `domain`, `path`, and `service` field.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["app_id"] = app_id
        __props__["auth_domain"] = auth_domain
        __props__["code_bucket"] = code_bucket
        __props__["database_type"] = database_type
        __props__["default_bucket"] = default_bucket
        __props__["default_hostname"] = default_hostname
        __props__["feature_settings"] = feature_settings
        __props__["gcr_domain"] = gcr_domain
        __props__["iap"] = iap
        __props__["location_id"] = location_id
        __props__["name"] = name
        __props__["project"] = project
        __props__["serving_status"] = serving_status
        __props__["url_dispatch_rules"] = url_dispatch_rules
        return Application(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> pulumi.Output[str]:
        """
        Identifier of the app, usually `{PROJECT_ID}`
        """
        return pulumi.get(self, "app_id")

    @property
    @pulumi.getter(name="authDomain")
    def auth_domain(self) -> pulumi.Output[str]:
        """
        The domain to authenticate users with when using App Engine's User API.
        """
        return pulumi.get(self, "auth_domain")

    @property
    @pulumi.getter(name="codeBucket")
    def code_bucket(self) -> pulumi.Output[str]:
        """
        The GCS bucket code is being stored in for this app.
        """
        return pulumi.get(self, "code_bucket")

    @property
    @pulumi.getter(name="databaseType")
    def database_type(self) -> pulumi.Output[str]:
        """
        The type of the Cloud Firestore or Cloud Datastore database associated with this application.
        Can be `CLOUD_FIRESTORE` or `CLOUD_DATASTORE_COMPATIBILITY` for new
        instances.  To support old instances, the value `CLOUD_DATASTORE` is accepted
        by the provider, but will be rejected by the API.
        """
        return pulumi.get(self, "database_type")

    @property
    @pulumi.getter(name="defaultBucket")
    def default_bucket(self) -> pulumi.Output[str]:
        """
        The GCS bucket content is being stored in for this app.
        """
        return pulumi.get(self, "default_bucket")

    @property
    @pulumi.getter(name="defaultHostname")
    def default_hostname(self) -> pulumi.Output[str]:
        """
        The default hostname for this app.
        """
        return pulumi.get(self, "default_hostname")

    @property
    @pulumi.getter(name="featureSettings")
    def feature_settings(self) -> pulumi.Output['outputs.ApplicationFeatureSettings']:
        """
        A block of optional settings to configure specific App Engine features:
        """
        return pulumi.get(self, "feature_settings")

    @property
    @pulumi.getter(name="gcrDomain")
    def gcr_domain(self) -> pulumi.Output[str]:
        """
        The GCR domain used for storing managed Docker images for this app.
        """
        return pulumi.get(self, "gcr_domain")

    @property
    @pulumi.getter
    def iap(self) -> pulumi.Output['outputs.ApplicationIap']:
        """
        Settings for enabling Cloud Identity Aware Proxy
        """
        return pulumi.get(self, "iap")

    @property
    @pulumi.getter(name="locationId")
    def location_id(self) -> pulumi.Output[str]:
        """
        The [location](https://cloud.google.com/appengine/docs/locations)
        to serve the app from.
        """
        return pulumi.get(self, "location_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Unique name of the app, usually `apps/{PROJECT_ID}`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The project ID to create the application under.
        ~>**NOTE:** GCP only accepts project ID, not project number. If you are using number,
        you may get a "Permission denied" error.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="servingStatus")
    def serving_status(self) -> pulumi.Output[str]:
        """
        The serving status of the app.
        """
        return pulumi.get(self, "serving_status")

    @property
    @pulumi.getter(name="urlDispatchRules")
    def url_dispatch_rules(self) -> pulumi.Output[Sequence['outputs.ApplicationUrlDispatchRule']]:
        """
        A list of dispatch rule blocks. Each block has a `domain`, `path`, and `service` field.
        """
        return pulumi.get(self, "url_dispatch_rules")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

