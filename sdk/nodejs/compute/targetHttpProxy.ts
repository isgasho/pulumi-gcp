// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Represents a TargetHttpProxy resource, which is used by one or more global
 * forwarding rule to route incoming HTTP requests to a URL map.
 * 
 * 
 * To get more information about TargetHttpProxy, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/v1/targetHttpProxies)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/compute/docs/load-balancing/http/target-proxies)
 * 
 * ## Example Usage - Target Http Proxy Basic
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const defaultHttpHealthCheck = new gcp.compute.HttpHealthCheck("default", {
 *     checkIntervalSec: 1,
 *     requestPath: "/",
 *     timeoutSec: 1,
 * });
 * const defaultBackendService = new gcp.compute.BackendService("default", {
 *     healthChecks: defaultHttpHealthCheck.selfLink,
 *     portName: "http",
 *     protocol: "HTTP",
 *     timeoutSec: 10,
 * });
 * const defaultURLMap = new gcp.compute.URLMap("default", {
 *     defaultService: defaultBackendService.selfLink,
 *     hostRules: [{
 *         hosts: ["mysite.com"],
 *         pathMatcher: "allpaths",
 *     }],
 *     pathMatchers: [{
 *         defaultService: defaultBackendService.selfLink,
 *         name: "allpaths",
 *         pathRules: [{
 *             paths: ["/*"],
 *             service: defaultBackendService.selfLink,
 *         }],
 *     }],
 * });
 * const defaultTargetHttpProxy = new gcp.compute.TargetHttpProxy("default", {
 *     urlMap: defaultURLMap.selfLink,
 * });
 * ```
 */
export class TargetHttpProxy extends pulumi.CustomResource {
    /**
     * Get an existing TargetHttpProxy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TargetHttpProxyState, opts?: pulumi.CustomResourceOptions): TargetHttpProxy {
        return new TargetHttpProxy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/targetHttpProxy:TargetHttpProxy';

    /**
     * Returns true if the given object is an instance of TargetHttpProxy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TargetHttpProxy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TargetHttpProxy.__pulumiType;
    }

    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    public /*out*/ readonly proxyId!: pulumi.Output<number>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    public readonly urlMap!: pulumi.Output<string>;

    /**
     * Create a TargetHttpProxy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TargetHttpProxyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TargetHttpProxyArgs | TargetHttpProxyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TargetHttpProxyState | undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["proxyId"] = state ? state.proxyId : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["urlMap"] = state ? state.urlMap : undefined;
        } else {
            const args = argsOrState as TargetHttpProxyArgs | undefined;
            if (!args || args.urlMap === undefined) {
                throw new Error("Missing required property 'urlMap'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["urlMap"] = args ? args.urlMap : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["proxyId"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        super(TargetHttpProxy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TargetHttpProxy resources.
 */
export interface TargetHttpProxyState {
    readonly creationTimestamp?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly proxyId?: pulumi.Input<number>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    readonly urlMap?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TargetHttpProxy resource.
 */
export interface TargetHttpProxyArgs {
    readonly description?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly urlMap: pulumi.Input<string>;
}
