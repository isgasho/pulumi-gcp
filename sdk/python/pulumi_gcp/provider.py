# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables


class Provider(pulumi.ProviderResource):
    def __init__(__self__, resource_name, opts=None, access_context_manager_custom_endpoint=None, access_token=None, app_engine_custom_endpoint=None, artifact_registry_custom_endpoint=None, batching=None, big_query_custom_endpoint=None, bigquery_connection_custom_endpoint=None, bigquery_data_transfer_custom_endpoint=None, bigquery_reservation_custom_endpoint=None, bigtable_custom_endpoint=None, billing_custom_endpoint=None, binary_authorization_custom_endpoint=None, cloud_billing_custom_endpoint=None, cloud_build_custom_endpoint=None, cloud_functions_custom_endpoint=None, cloud_identity_custom_endpoint=None, cloud_iot_custom_endpoint=None, cloud_run_custom_endpoint=None, cloud_scheduler_custom_endpoint=None, cloud_tasks_custom_endpoint=None, composer_custom_endpoint=None, compute_beta_custom_endpoint=None, compute_custom_endpoint=None, container_analysis_custom_endpoint=None, container_beta_custom_endpoint=None, container_custom_endpoint=None, credentials=None, data_catalog_custom_endpoint=None, data_fusion_custom_endpoint=None, dataflow_custom_endpoint=None, dataproc_beta_custom_endpoint=None, dataproc_custom_endpoint=None, datastore_custom_endpoint=None, deployment_manager_custom_endpoint=None, dialogflow_custom_endpoint=None, dns_beta_custom_endpoint=None, dns_custom_endpoint=None, filestore_custom_endpoint=None, firebase_custom_endpoint=None, firestore_custom_endpoint=None, game_services_custom_endpoint=None, healthcare_custom_endpoint=None, iam_credentials_custom_endpoint=None, iam_custom_endpoint=None, iap_custom_endpoint=None, identity_platform_custom_endpoint=None, kms_custom_endpoint=None, logging_custom_endpoint=None, memcache_custom_endpoint=None, ml_engine_custom_endpoint=None, monitoring_custom_endpoint=None, network_management_custom_endpoint=None, notebooks_custom_endpoint=None, os_login_custom_endpoint=None, project=None, pubsub_custom_endpoint=None, redis_custom_endpoint=None, region=None, request_timeout=None, resource_manager_custom_endpoint=None, resource_manager_v2beta1_custom_endpoint=None, runtime_config_custom_endpoint=None, runtimeconfig_custom_endpoint=None, scopes=None, secret_manager_custom_endpoint=None, security_center_custom_endpoint=None, security_scanner_custom_endpoint=None, service_directory_custom_endpoint=None, service_management_custom_endpoint=None, service_networking_custom_endpoint=None, service_usage_custom_endpoint=None, source_repo_custom_endpoint=None, spanner_custom_endpoint=None, sql_custom_endpoint=None, storage_custom_endpoint=None, storage_transfer_custom_endpoint=None, tpu_custom_endpoint=None, user_project_override=None, vpc_access_custom_endpoint=None, zone=None, __props__=None, __name__=None, __opts__=None):
        """
        The provider type for the google-beta package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.

        The **batching** object supports the following:

          * `enable_batching` (`pulumi.Input[bool]`)
          * `send_after` (`pulumi.Input[str]`)
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

            __props__['access_context_manager_custom_endpoint'] = access_context_manager_custom_endpoint
            __props__['access_token'] = access_token
            __props__['app_engine_custom_endpoint'] = app_engine_custom_endpoint
            __props__['artifact_registry_custom_endpoint'] = artifact_registry_custom_endpoint
            __props__['batching'] = pulumi.Output.from_input(batching).apply(json.dumps) if batching is not None else None
            __props__['big_query_custom_endpoint'] = big_query_custom_endpoint
            __props__['bigquery_connection_custom_endpoint'] = bigquery_connection_custom_endpoint
            __props__['bigquery_data_transfer_custom_endpoint'] = bigquery_data_transfer_custom_endpoint
            __props__['bigquery_reservation_custom_endpoint'] = bigquery_reservation_custom_endpoint
            __props__['bigtable_custom_endpoint'] = bigtable_custom_endpoint
            __props__['billing_custom_endpoint'] = billing_custom_endpoint
            __props__['binary_authorization_custom_endpoint'] = binary_authorization_custom_endpoint
            __props__['cloud_billing_custom_endpoint'] = cloud_billing_custom_endpoint
            __props__['cloud_build_custom_endpoint'] = cloud_build_custom_endpoint
            __props__['cloud_functions_custom_endpoint'] = cloud_functions_custom_endpoint
            __props__['cloud_identity_custom_endpoint'] = cloud_identity_custom_endpoint
            __props__['cloud_iot_custom_endpoint'] = cloud_iot_custom_endpoint
            __props__['cloud_run_custom_endpoint'] = cloud_run_custom_endpoint
            __props__['cloud_scheduler_custom_endpoint'] = cloud_scheduler_custom_endpoint
            __props__['cloud_tasks_custom_endpoint'] = cloud_tasks_custom_endpoint
            __props__['composer_custom_endpoint'] = composer_custom_endpoint
            __props__['compute_beta_custom_endpoint'] = compute_beta_custom_endpoint
            __props__['compute_custom_endpoint'] = compute_custom_endpoint
            __props__['container_analysis_custom_endpoint'] = container_analysis_custom_endpoint
            __props__['container_beta_custom_endpoint'] = container_beta_custom_endpoint
            __props__['container_custom_endpoint'] = container_custom_endpoint
            if credentials is None:
                credentials = utilities.get_env('GOOGLE_CREDENTIALS', 'GOOGLE_CLOUD_KEYFILE_JSON', 'GCLOUD_KEYFILE_JSON')
            __props__['credentials'] = credentials
            __props__['data_catalog_custom_endpoint'] = data_catalog_custom_endpoint
            __props__['data_fusion_custom_endpoint'] = data_fusion_custom_endpoint
            __props__['dataflow_custom_endpoint'] = dataflow_custom_endpoint
            __props__['dataproc_beta_custom_endpoint'] = dataproc_beta_custom_endpoint
            __props__['dataproc_custom_endpoint'] = dataproc_custom_endpoint
            __props__['datastore_custom_endpoint'] = datastore_custom_endpoint
            __props__['deployment_manager_custom_endpoint'] = deployment_manager_custom_endpoint
            __props__['dialogflow_custom_endpoint'] = dialogflow_custom_endpoint
            __props__['dns_beta_custom_endpoint'] = dns_beta_custom_endpoint
            __props__['dns_custom_endpoint'] = dns_custom_endpoint
            __props__['filestore_custom_endpoint'] = filestore_custom_endpoint
            __props__['firebase_custom_endpoint'] = firebase_custom_endpoint
            __props__['firestore_custom_endpoint'] = firestore_custom_endpoint
            __props__['game_services_custom_endpoint'] = game_services_custom_endpoint
            __props__['healthcare_custom_endpoint'] = healthcare_custom_endpoint
            __props__['iam_credentials_custom_endpoint'] = iam_credentials_custom_endpoint
            __props__['iam_custom_endpoint'] = iam_custom_endpoint
            __props__['iap_custom_endpoint'] = iap_custom_endpoint
            __props__['identity_platform_custom_endpoint'] = identity_platform_custom_endpoint
            __props__['kms_custom_endpoint'] = kms_custom_endpoint
            __props__['logging_custom_endpoint'] = logging_custom_endpoint
            __props__['memcache_custom_endpoint'] = memcache_custom_endpoint
            __props__['ml_engine_custom_endpoint'] = ml_engine_custom_endpoint
            __props__['monitoring_custom_endpoint'] = monitoring_custom_endpoint
            __props__['network_management_custom_endpoint'] = network_management_custom_endpoint
            __props__['notebooks_custom_endpoint'] = notebooks_custom_endpoint
            __props__['os_login_custom_endpoint'] = os_login_custom_endpoint
            if project is None:
                project = utilities.get_env('GOOGLE_PROJECT', 'GOOGLE_CLOUD_PROJECT', 'GCLOUD_PROJECT', 'CLOUDSDK_CORE_PROJECT')
            __props__['project'] = project
            __props__['pubsub_custom_endpoint'] = pubsub_custom_endpoint
            __props__['redis_custom_endpoint'] = redis_custom_endpoint
            if region is None:
                region = utilities.get_env('GOOGLE_REGION', 'GCLOUD_REGION', 'CLOUDSDK_COMPUTE_REGION')
            __props__['region'] = region
            __props__['request_timeout'] = request_timeout
            __props__['resource_manager_custom_endpoint'] = resource_manager_custom_endpoint
            __props__['resource_manager_v2beta1_custom_endpoint'] = resource_manager_v2beta1_custom_endpoint
            __props__['runtime_config_custom_endpoint'] = runtime_config_custom_endpoint
            __props__['runtimeconfig_custom_endpoint'] = runtimeconfig_custom_endpoint
            __props__['scopes'] = pulumi.Output.from_input(scopes).apply(json.dumps) if scopes is not None else None
            __props__['secret_manager_custom_endpoint'] = secret_manager_custom_endpoint
            __props__['security_center_custom_endpoint'] = security_center_custom_endpoint
            __props__['security_scanner_custom_endpoint'] = security_scanner_custom_endpoint
            __props__['service_directory_custom_endpoint'] = service_directory_custom_endpoint
            __props__['service_management_custom_endpoint'] = service_management_custom_endpoint
            __props__['service_networking_custom_endpoint'] = service_networking_custom_endpoint
            __props__['service_usage_custom_endpoint'] = service_usage_custom_endpoint
            __props__['source_repo_custom_endpoint'] = source_repo_custom_endpoint
            __props__['spanner_custom_endpoint'] = spanner_custom_endpoint
            __props__['sql_custom_endpoint'] = sql_custom_endpoint
            __props__['storage_custom_endpoint'] = storage_custom_endpoint
            __props__['storage_transfer_custom_endpoint'] = storage_transfer_custom_endpoint
            __props__['tpu_custom_endpoint'] = tpu_custom_endpoint
            __props__['user_project_override'] = pulumi.Output.from_input(user_project_override).apply(json.dumps) if user_project_override is not None else None
            __props__['vpc_access_custom_endpoint'] = vpc_access_custom_endpoint
            if zone is None:
                zone = utilities.get_env('GOOGLE_ZONE', 'GCLOUD_ZONE', 'CLOUDSDK_COMPUTE_ZONE')
            __props__['zone'] = zone
        super(Provider, __self__).__init__(
            'gcp',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
