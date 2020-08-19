// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A Service-Level Objective (SLO) describes the level of desired good
 * service. It consists of a service-level indicator (SLI), a performance
 * goal, and a period over which the objective is to be evaluated against
 * that goal. The SLO can use SLIs defined in a number of different manners.
 * Typical SLOs might include "99% of requests in each rolling week have
 * latency below 200 milliseconds" or "99.5% of requests in each calendar
 * month return successfully."
 *
 *
 * To get more information about Slo, see:
 *
 * * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/services.serviceLevelObjectives)
 * * How-to Guides
 *     * [Service Monitoring](https://cloud.google.com/monitoring/service-monitoring)
 *     * [Monitoring API Documentation](https://cloud.google.com/monitoring/api/v3/)
 *
 * ## Example Usage
 *
 * ### Monitoring Slo Appengine
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const default = gcp.monitoring.getAppEngineService({
 *     moduleId: "default",
 * });
 * const appengSlo = new gcp.monitoring.Slo("appengSlo", {
 *     service: _default.then(_default => _default.serviceId),
 *     sloId: "ae-slo",
 *     displayName: "Test SLO for App Engine",
 *     goal: 0.9,
 *     calendarPeriod: "DAY",
 *     basic_sli: {
 *         latency: {
 *             threshold: "1s",
 *         },
 *     },
 * });
 * ```
 */
export class Slo extends pulumi.CustomResource {
    /**
     * Get an existing Slo resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SloState, opts?: pulumi.CustomResourceOptions): Slo {
        return new Slo(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:monitoring/slo:Slo';

    /**
     * Returns true if the given object is an instance of Slo.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Slo {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Slo.__pulumiType;
    }

    /**
     * Basic Service-Level Indicator (SLI) on a well-known service type.
     * Performance will be computed on the basis of pre-defined metrics.
     * SLIs are used to measure and calculate the quality of the Service's
     * performance with respect to a single aspect of service quality.
     * Exactly one of the following must be set:
     * `basicSli`, `requestBasedSli`, `windowsBasedSli`
     * Structure is documented below.
     */
    public readonly basicSli!: pulumi.Output<outputs.monitoring.SloBasicSli | undefined>;
    /**
     * A calendar period, semantically "since the start of the current
     * <calendarPeriod>".
     * Possible values are `DAY`, `WEEK`, `FORTNIGHT`, and `MONTH`.
     */
    public readonly calendarPeriod!: pulumi.Output<string | undefined>;
    /**
     * Name used for UI elements listing this SLO.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * The fraction of service that must be good in order for this objective
     * to be met. 0 < goal <= 0.999
     */
    public readonly goal!: pulumi.Output<number>;
    /**
     * The full resource name for this service. The syntax is:
     * projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]/serviceLevelObjectives/[SLO_NAME]
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * A request-based SLI defines a SLI for which atomic units of
     * service are counted directly.
     * A SLI describes a good service.
     * It is used to measure and calculate the quality of the Service's
     * performance with respect to a single aspect of service quality.
     * Exactly one of the following must be set:
     * `basicSli`, `requestBasedSli`, `windowsBasedSli`
     * Structure is documented below.
     */
    public readonly requestBasedSli!: pulumi.Output<outputs.monitoring.SloRequestBasedSli | undefined>;
    /**
     * A rolling time period, semantically "in the past X days".
     * Must be between 1 to 30 days, inclusive.
     */
    public readonly rollingPeriodDays!: pulumi.Output<number | undefined>;
    /**
     * ID of the service to which this SLO belongs.
     */
    public readonly service!: pulumi.Output<string>;
    /**
     * The id to use for this ServiceLevelObjective. If omitted, an id will be generated instead.
     */
    public readonly sloId!: pulumi.Output<string>;
    /**
     * A windows-based SLI defines the criteria for time windows.
     * goodService is defined based off the count of these time windows
     * for which the provided service was of good quality.
     * A SLI describes a good service. It is used to measure and calculate
     * the quality of the Service's performance with respect to a single
     * aspect of service quality.
     * Exactly one of the following must be set:
     * `basicSli`, `requestBasedSli`, `windowsBasedSli`
     * Structure is documented below.
     */
    public readonly windowsBasedSli!: pulumi.Output<outputs.monitoring.SloWindowsBasedSli | undefined>;

    /**
     * Create a Slo resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SloArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SloArgs | SloState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SloState | undefined;
            inputs["basicSli"] = state ? state.basicSli : undefined;
            inputs["calendarPeriod"] = state ? state.calendarPeriod : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["goal"] = state ? state.goal : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["requestBasedSli"] = state ? state.requestBasedSli : undefined;
            inputs["rollingPeriodDays"] = state ? state.rollingPeriodDays : undefined;
            inputs["service"] = state ? state.service : undefined;
            inputs["sloId"] = state ? state.sloId : undefined;
            inputs["windowsBasedSli"] = state ? state.windowsBasedSli : undefined;
        } else {
            const args = argsOrState as SloArgs | undefined;
            if (!args || args.goal === undefined) {
                throw new Error("Missing required property 'goal'");
            }
            if (!args || args.service === undefined) {
                throw new Error("Missing required property 'service'");
            }
            inputs["basicSli"] = args ? args.basicSli : undefined;
            inputs["calendarPeriod"] = args ? args.calendarPeriod : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["goal"] = args ? args.goal : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["requestBasedSli"] = args ? args.requestBasedSli : undefined;
            inputs["rollingPeriodDays"] = args ? args.rollingPeriodDays : undefined;
            inputs["service"] = args ? args.service : undefined;
            inputs["sloId"] = args ? args.sloId : undefined;
            inputs["windowsBasedSli"] = args ? args.windowsBasedSli : undefined;
            inputs["name"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Slo.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Slo resources.
 */
export interface SloState {
    /**
     * Basic Service-Level Indicator (SLI) on a well-known service type.
     * Performance will be computed on the basis of pre-defined metrics.
     * SLIs are used to measure and calculate the quality of the Service's
     * performance with respect to a single aspect of service quality.
     * Exactly one of the following must be set:
     * `basicSli`, `requestBasedSli`, `windowsBasedSli`
     * Structure is documented below.
     */
    readonly basicSli?: pulumi.Input<inputs.monitoring.SloBasicSli>;
    /**
     * A calendar period, semantically "since the start of the current
     * <calendarPeriod>".
     * Possible values are `DAY`, `WEEK`, `FORTNIGHT`, and `MONTH`.
     */
    readonly calendarPeriod?: pulumi.Input<string>;
    /**
     * Name used for UI elements listing this SLO.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The fraction of service that must be good in order for this objective
     * to be met. 0 < goal <= 0.999
     */
    readonly goal?: pulumi.Input<number>;
    /**
     * The full resource name for this service. The syntax is:
     * projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]/serviceLevelObjectives/[SLO_NAME]
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * A request-based SLI defines a SLI for which atomic units of
     * service are counted directly.
     * A SLI describes a good service.
     * It is used to measure and calculate the quality of the Service's
     * performance with respect to a single aspect of service quality.
     * Exactly one of the following must be set:
     * `basicSli`, `requestBasedSli`, `windowsBasedSli`
     * Structure is documented below.
     */
    readonly requestBasedSli?: pulumi.Input<inputs.monitoring.SloRequestBasedSli>;
    /**
     * A rolling time period, semantically "in the past X days".
     * Must be between 1 to 30 days, inclusive.
     */
    readonly rollingPeriodDays?: pulumi.Input<number>;
    /**
     * ID of the service to which this SLO belongs.
     */
    readonly service?: pulumi.Input<string>;
    /**
     * The id to use for this ServiceLevelObjective. If omitted, an id will be generated instead.
     */
    readonly sloId?: pulumi.Input<string>;
    /**
     * A windows-based SLI defines the criteria for time windows.
     * goodService is defined based off the count of these time windows
     * for which the provided service was of good quality.
     * A SLI describes a good service. It is used to measure and calculate
     * the quality of the Service's performance with respect to a single
     * aspect of service quality.
     * Exactly one of the following must be set:
     * `basicSli`, `requestBasedSli`, `windowsBasedSli`
     * Structure is documented below.
     */
    readonly windowsBasedSli?: pulumi.Input<inputs.monitoring.SloWindowsBasedSli>;
}

/**
 * The set of arguments for constructing a Slo resource.
 */
export interface SloArgs {
    /**
     * Basic Service-Level Indicator (SLI) on a well-known service type.
     * Performance will be computed on the basis of pre-defined metrics.
     * SLIs are used to measure and calculate the quality of the Service's
     * performance with respect to a single aspect of service quality.
     * Exactly one of the following must be set:
     * `basicSli`, `requestBasedSli`, `windowsBasedSli`
     * Structure is documented below.
     */
    readonly basicSli?: pulumi.Input<inputs.monitoring.SloBasicSli>;
    /**
     * A calendar period, semantically "since the start of the current
     * <calendarPeriod>".
     * Possible values are `DAY`, `WEEK`, `FORTNIGHT`, and `MONTH`.
     */
    readonly calendarPeriod?: pulumi.Input<string>;
    /**
     * Name used for UI elements listing this SLO.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * The fraction of service that must be good in order for this objective
     * to be met. 0 < goal <= 0.999
     */
    readonly goal: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * A request-based SLI defines a SLI for which atomic units of
     * service are counted directly.
     * A SLI describes a good service.
     * It is used to measure and calculate the quality of the Service's
     * performance with respect to a single aspect of service quality.
     * Exactly one of the following must be set:
     * `basicSli`, `requestBasedSli`, `windowsBasedSli`
     * Structure is documented below.
     */
    readonly requestBasedSli?: pulumi.Input<inputs.monitoring.SloRequestBasedSli>;
    /**
     * A rolling time period, semantically "in the past X days".
     * Must be between 1 to 30 days, inclusive.
     */
    readonly rollingPeriodDays?: pulumi.Input<number>;
    /**
     * ID of the service to which this SLO belongs.
     */
    readonly service: pulumi.Input<string>;
    /**
     * The id to use for this ServiceLevelObjective. If omitted, an id will be generated instead.
     */
    readonly sloId?: pulumi.Input<string>;
    /**
     * A windows-based SLI defines the criteria for time windows.
     * goodService is defined based off the count of these time windows
     * for which the provided service was of good quality.
     * A SLI describes a good service. It is used to measure and calculate
     * the quality of the Service's performance with respect to a single
     * aspect of service quality.
     * Exactly one of the following must be set:
     * `basicSli`, `requestBasedSli`, `windowsBasedSli`
     * Structure is documented below.
     */
    readonly windowsBasedSli?: pulumi.Input<inputs.monitoring.SloWindowsBasedSli>;
}
