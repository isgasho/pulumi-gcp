// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Get information about a Google Cloud Run. For more information see
 * the [official documentation](https://cloud.google.com/run/docs/)
 * and [API](https://cloud.google.com/run/docs/apis).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const run_service = pulumi.output(gcp.cloudrun.getService({
 *     location: "us-central1",
 *     name: "my-service",
 * }, { async: true }));
 * ```
 */
export function getService(args: GetServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:cloudrun/getService:getService", {
        "location": args.location,
        "name": args.name,
        "project": args.project,
    }, opts);
}

/**
 * A collection of arguments for invoking getService.
 */
export interface GetServiceArgs {
    /**
     * The location of the cloud run instance. eg us-central1
     */
    readonly location: string;
    /**
     * The name of the Cloud Run Service.
     */
    readonly name: string;
    /**
     * The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: string;
}

/**
 * A collection of values returned by getService.
 */
export interface GetServiceResult {
    readonly autogenerateRevisionName: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly location: string;
    readonly metadatas: outputs.cloudrun.GetServiceMetadata[];
    readonly name: string;
    readonly project?: string;
    readonly statuses: outputs.cloudrun.GetServiceStatus[];
    readonly templates: outputs.cloudrun.GetServiceTemplate[];
    readonly traffics: outputs.cloudrun.GetServiceTraffic[];
}
