// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

let __config = new pulumi.Config("gcp");

export let accessContextManagerCustomEndpoint: string | undefined = __config.get("accessContextManagerCustomEndpoint");
export let accessToken: string | undefined = __config.get("accessToken");
export let appEngineCustomEndpoint: string | undefined = __config.get("appEngineCustomEndpoint");
export let artifactRegistryCustomEndpoint: string | undefined = __config.get("artifactRegistryCustomEndpoint");
export let batching: { enableBatching?: boolean, sendAfter?: string } | undefined = __config.getObject<{ enableBatching?: boolean, sendAfter?: string }>("batching");
export let bigQueryCustomEndpoint: string | undefined = __config.get("bigQueryCustomEndpoint");
export let bigqueryConnectionCustomEndpoint: string | undefined = __config.get("bigqueryConnectionCustomEndpoint");
export let bigqueryDataTransferCustomEndpoint: string | undefined = __config.get("bigqueryDataTransferCustomEndpoint");
export let bigqueryReservationCustomEndpoint: string | undefined = __config.get("bigqueryReservationCustomEndpoint");
export let bigtableCustomEndpoint: string | undefined = __config.get("bigtableCustomEndpoint");
export let billingCustomEndpoint: string | undefined = __config.get("billingCustomEndpoint");
export let binaryAuthorizationCustomEndpoint: string | undefined = __config.get("binaryAuthorizationCustomEndpoint");
export let cloudBillingCustomEndpoint: string | undefined = __config.get("cloudBillingCustomEndpoint");
export let cloudBuildCustomEndpoint: string | undefined = __config.get("cloudBuildCustomEndpoint");
export let cloudFunctionsCustomEndpoint: string | undefined = __config.get("cloudFunctionsCustomEndpoint");
export let cloudIotCustomEndpoint: string | undefined = __config.get("cloudIotCustomEndpoint");
export let cloudRunCustomEndpoint: string | undefined = __config.get("cloudRunCustomEndpoint");
export let cloudSchedulerCustomEndpoint: string | undefined = __config.get("cloudSchedulerCustomEndpoint");
export let cloudTasksCustomEndpoint: string | undefined = __config.get("cloudTasksCustomEndpoint");
export let composerCustomEndpoint: string | undefined = __config.get("composerCustomEndpoint");
export let computeBetaCustomEndpoint: string | undefined = __config.get("computeBetaCustomEndpoint");
export let computeCustomEndpoint: string | undefined = __config.get("computeCustomEndpoint");
export let containerAnalysisCustomEndpoint: string | undefined = __config.get("containerAnalysisCustomEndpoint");
export let containerBetaCustomEndpoint: string | undefined = __config.get("containerBetaCustomEndpoint");
export let containerCustomEndpoint: string | undefined = __config.get("containerCustomEndpoint");
export let credentials: string | undefined = __config.get("credentials") || utilities.getEnv("GOOGLE_CREDENTIALS", "GOOGLE_CLOUD_KEYFILE_JSON", "GCLOUD_KEYFILE_JSON");
export let dataFusionCustomEndpoint: string | undefined = __config.get("dataFusionCustomEndpoint");
export let dataflowCustomEndpoint: string | undefined = __config.get("dataflowCustomEndpoint");
export let dataprocBetaCustomEndpoint: string | undefined = __config.get("dataprocBetaCustomEndpoint");
export let dataprocCustomEndpoint: string | undefined = __config.get("dataprocCustomEndpoint");
export let datastoreCustomEndpoint: string | undefined = __config.get("datastoreCustomEndpoint");
export let deploymentManagerCustomEndpoint: string | undefined = __config.get("deploymentManagerCustomEndpoint");
export let dialogflowCustomEndpoint: string | undefined = __config.get("dialogflowCustomEndpoint");
export let dnsBetaCustomEndpoint: string | undefined = __config.get("dnsBetaCustomEndpoint");
export let dnsCustomEndpoint: string | undefined = __config.get("dnsCustomEndpoint");
export let filestoreCustomEndpoint: string | undefined = __config.get("filestoreCustomEndpoint");
export let firebaseCustomEndpoint: string | undefined = __config.get("firebaseCustomEndpoint");
export let firestoreCustomEndpoint: string | undefined = __config.get("firestoreCustomEndpoint");
export let gameServicesCustomEndpoint: string | undefined = __config.get("gameServicesCustomEndpoint");
export let healthcareCustomEndpoint: string | undefined = __config.get("healthcareCustomEndpoint");
export let iamCredentialsCustomEndpoint: string | undefined = __config.get("iamCredentialsCustomEndpoint");
export let iamCustomEndpoint: string | undefined = __config.get("iamCustomEndpoint");
export let iapCustomEndpoint: string | undefined = __config.get("iapCustomEndpoint");
export let identityPlatformCustomEndpoint: string | undefined = __config.get("identityPlatformCustomEndpoint");
export let kmsCustomEndpoint: string | undefined = __config.get("kmsCustomEndpoint");
export let loggingCustomEndpoint: string | undefined = __config.get("loggingCustomEndpoint");
export let mlEngineCustomEndpoint: string | undefined = __config.get("mlEngineCustomEndpoint");
export let monitoringCustomEndpoint: string | undefined = __config.get("monitoringCustomEndpoint");
export let osLoginCustomEndpoint: string | undefined = __config.get("osLoginCustomEndpoint");
export let project: string | undefined = __config.get("project") || utilities.getEnv("GOOGLE_PROJECT", "GOOGLE_CLOUD_PROJECT", "GCLOUD_PROJECT", "CLOUDSDK_CORE_PROJECT");
export let pubsubCustomEndpoint: string | undefined = __config.get("pubsubCustomEndpoint");
export let redisCustomEndpoint: string | undefined = __config.get("redisCustomEndpoint");
export let region: string | undefined = __config.get("region") || utilities.getEnv("GOOGLE_REGION", "GCLOUD_REGION", "CLOUDSDK_COMPUTE_REGION");
export let requestTimeout: string | undefined = __config.get("requestTimeout");
export let resourceManagerCustomEndpoint: string | undefined = __config.get("resourceManagerCustomEndpoint");
export let resourceManagerV2beta1CustomEndpoint: string | undefined = __config.get("resourceManagerV2beta1CustomEndpoint");
export let runtimeConfigCustomEndpoint: string | undefined = __config.get("runtimeConfigCustomEndpoint");
export let runtimeconfigCustomEndpoint: string | undefined = __config.get("runtimeconfigCustomEndpoint");
export let scopes: string[] | undefined = __config.getObject<string[]>("scopes");
export let secretManagerCustomEndpoint: string | undefined = __config.get("secretManagerCustomEndpoint");
export let securityCenterCustomEndpoint: string | undefined = __config.get("securityCenterCustomEndpoint");
export let securityScannerCustomEndpoint: string | undefined = __config.get("securityScannerCustomEndpoint");
export let serviceDirectoryCustomEndpoint: string | undefined = __config.get("serviceDirectoryCustomEndpoint");
export let serviceManagementCustomEndpoint: string | undefined = __config.get("serviceManagementCustomEndpoint");
export let serviceNetworkingCustomEndpoint: string | undefined = __config.get("serviceNetworkingCustomEndpoint");
export let serviceUsageCustomEndpoint: string | undefined = __config.get("serviceUsageCustomEndpoint");
export let sourceRepoCustomEndpoint: string | undefined = __config.get("sourceRepoCustomEndpoint");
export let spannerCustomEndpoint: string | undefined = __config.get("spannerCustomEndpoint");
export let sqlCustomEndpoint: string | undefined = __config.get("sqlCustomEndpoint");
export let storageCustomEndpoint: string | undefined = __config.get("storageCustomEndpoint");
export let storageTransferCustomEndpoint: string | undefined = __config.get("storageTransferCustomEndpoint");
export let tpuCustomEndpoint: string | undefined = __config.get("tpuCustomEndpoint");
export let userProjectOverride: boolean | undefined = __config.getObject<boolean>("userProjectOverride");
export let vpcAccessCustomEndpoint: string | undefined = __config.get("vpcAccessCustomEndpoint");
export let zone: string | undefined = __config.get("zone") || utilities.getEnv("GOOGLE_ZONE", "GCLOUD_ZONE", "CLOUDSDK_COMPUTE_ZONE");
