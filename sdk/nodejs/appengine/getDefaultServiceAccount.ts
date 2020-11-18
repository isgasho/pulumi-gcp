// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to retrieve the default App Engine service account for the specified project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const default = gcp.appengine.getDefaultServiceAccount({});
 * export const defaultAccount = _default.then(_default => _default.email);
 * ```
 */
export function getDefaultServiceAccount(args?: GetDefaultServiceAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetDefaultServiceAccountResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:appengine/getDefaultServiceAccount:getDefaultServiceAccount", {
        "project": args.project,
    }, opts);
}

/**
 * A collection of arguments for invoking getDefaultServiceAccount.
 */
export interface GetDefaultServiceAccountArgs {
    /**
     * The project ID. If it is not provided, the provider project is used.
     */
    readonly project?: string;
}

/**
 * A collection of values returned by getDefaultServiceAccount.
 */
export interface GetDefaultServiceAccountResult {
    /**
     * The display name for the service account.
     */
    readonly displayName: string;
    /**
     * Email address of the default service account used by App Engine in this project.
     */
    readonly email: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The fully-qualified name of the service account.
     */
    readonly name: string;
    readonly project: string;
    /**
     * The unique id of the service account.
     */
    readonly uniqueId: string;
}
