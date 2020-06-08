// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A NotificationChannel is a medium through which an alert is delivered
 * when a policy violation is detected. Examples of channels include email, SMS,
 * and third-party messaging applications. Fields containing sensitive information
 * like authentication tokens or contact info are only partially populated on retrieval.
 *
 *
 * To get more information about NotificationChannel, see:
 *
 * * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannels)
 * * How-to Guides
 *     * [Notification Options](https://cloud.google.com/monitoring/support/notification-options)
 *     * [Monitoring API Documentation](https://cloud.google.com/monitoring/api/v3/)
 *
 *
 * ## Example Usage - Notification Channel Basic
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const basic = gcp.monitoring.getNotificationChannel({
 *     displayName: "Test Notification Channel",
 * });
 * const alertPolicy = new gcp.monitoring.AlertPolicy("alertPolicy", {
 *     displayName: "My Alert Policy",
 *     notificationChannels: [basic.then(basic => basic.name)],
 *     combiner: "OR",
 *     conditions: [{
 *         displayName: "test condition",
 *         condition_threshold: {
 *             filter: "metric.type=\"compute.googleapis.com/instance/disk/write_bytes_count\" AND resource.type=\"gce_instance\"",
 *             duration: "60s",
 *             comparison: "COMPARISON_GT",
 *             aggregations: [{
 *                 alignmentPeriod: "60s",
 *                 perSeriesAligner: "ALIGN_RATE",
 *             }],
 *         },
 *     }],
 * });
 * ```
 */
export function getNotificationChannel(args?: GetNotificationChannelArgs, opts?: pulumi.InvokeOptions): Promise<GetNotificationChannelResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:monitoring/getNotificationChannel:getNotificationChannel", {
        "displayName": args.displayName,
        "labels": args.labels,
        "project": args.project,
        "type": args.type,
        "userLabels": args.userLabels,
    }, opts);
}

/**
 * A collection of arguments for invoking getNotificationChannel.
 */
export interface GetNotificationChannelArgs {
    /**
     * The display name for this notification channel.
     */
    readonly displayName?: string;
    /**
     * Labels (corresponding to the
     * NotificationChannelDescriptor schema) to filter the notification channels by.
     */
    readonly labels?: {[key: string]: string};
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: string;
    /**
     * The type of the notification channel.
     */
    readonly type?: string;
    /**
     * User-provided key-value labels to filter by.
     */
    readonly userLabels?: {[key: string]: string};
}

/**
 * A collection of values returned by getNotificationChannel.
 */
export interface GetNotificationChannelResult {
    readonly description: string;
    readonly displayName?: string;
    readonly enabled: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly labels?: {[key: string]: string};
    readonly name: string;
    readonly project?: string;
    readonly sensitiveLabels: outputs.monitoring.GetNotificationChannelSensitiveLabel[];
    readonly type?: string;
    readonly userLabels?: {[key: string]: string};
    readonly verificationStatus: string;
}
