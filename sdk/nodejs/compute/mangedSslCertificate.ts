// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * @deprecated gcp.compute.MangedSslCertificate has been deprecated in favor of gcp.compute.ManagedSslCertificate
 */
export class MangedSslCertificate extends pulumi.CustomResource {
    /**
     * Get an existing MangedSslCertificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MangedSslCertificateState, opts?: pulumi.CustomResourceOptions): MangedSslCertificate {
        pulumi.log.warn("MangedSslCertificate is deprecated: gcp.compute.MangedSslCertificate has been deprecated in favor of gcp.compute.ManagedSslCertificate")
        return new MangedSslCertificate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/mangedSslCertificate:MangedSslCertificate';

    /**
     * Returns true if the given object is an instance of MangedSslCertificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MangedSslCertificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MangedSslCertificate.__pulumiType;
    }

    /**
     * The unique identifier for the resource.
     */
    public readonly certificateId!: pulumi.Output<number>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Expire time of the certificate.
     */
    public /*out*/ readonly expireTime!: pulumi.Output<string>;
    /**
     * Properties relevant to a managed certificate. These will be used if the certificate is managed (as indicated by a value
     * of 'MANAGED' in 'type').
     */
    public readonly managed!: pulumi.Output<outputs.compute.MangedSslCertificateManaged | undefined>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
     * comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
     * '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
     * must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash. These are in the same
     * namespace as the managed SSL certificates.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Domains associated with the certificate via Subject Alternative Name.
     */
    public /*out*/ readonly subjectAlternativeNames!: pulumi.Output<string[]>;
    /**
     * Enum field whose value is always 'MANAGED' - used to signal to the API which type this is. Default value: "MANAGED"
     * Possible values: ["MANAGED"]
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a MangedSslCertificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated gcp.compute.MangedSslCertificate has been deprecated in favor of gcp.compute.ManagedSslCertificate */
    constructor(name: string, args?: MangedSslCertificateArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated gcp.compute.MangedSslCertificate has been deprecated in favor of gcp.compute.ManagedSslCertificate */
    constructor(name: string, argsOrState?: MangedSslCertificateArgs | MangedSslCertificateState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("MangedSslCertificate is deprecated: gcp.compute.MangedSslCertificate has been deprecated in favor of gcp.compute.ManagedSslCertificate")
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as MangedSslCertificateState | undefined;
            inputs["certificateId"] = state ? state.certificateId : undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["expireTime"] = state ? state.expireTime : undefined;
            inputs["managed"] = state ? state.managed : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["subjectAlternativeNames"] = state ? state.subjectAlternativeNames : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as MangedSslCertificateArgs | undefined;
            inputs["certificateId"] = args ? args.certificateId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["managed"] = args ? args.managed : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["expireTime"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["subjectAlternativeNames"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(MangedSslCertificate.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MangedSslCertificate resources.
 */
export interface MangedSslCertificateState {
    /**
     * The unique identifier for the resource.
     */
    readonly certificateId?: pulumi.Input<number>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp?: pulumi.Input<string>;
    /**
     * An optional description of this resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Expire time of the certificate.
     */
    readonly expireTime?: pulumi.Input<string>;
    /**
     * Properties relevant to a managed certificate. These will be used if the certificate is managed (as indicated by a value
     * of 'MANAGED' in 'type').
     */
    readonly managed?: pulumi.Input<inputs.compute.MangedSslCertificateManaged>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
     * comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
     * '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
     * must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash. These are in the same
     * namespace as the managed SSL certificates.
     */
    readonly name?: pulumi.Input<string>;
    readonly project?: pulumi.Input<string>;
    readonly selfLink?: pulumi.Input<string>;
    /**
     * Domains associated with the certificate via Subject Alternative Name.
     */
    readonly subjectAlternativeNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Enum field whose value is always 'MANAGED' - used to signal to the API which type this is. Default value: "MANAGED"
     * Possible values: ["MANAGED"]
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MangedSslCertificate resource.
 */
export interface MangedSslCertificateArgs {
    /**
     * The unique identifier for the resource.
     */
    readonly certificateId?: pulumi.Input<number>;
    /**
     * An optional description of this resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Properties relevant to a managed certificate. These will be used if the certificate is managed (as indicated by a value
     * of 'MANAGED' in 'type').
     */
    readonly managed?: pulumi.Input<inputs.compute.MangedSslCertificateManaged>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
     * comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
     * '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
     * must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash. These are in the same
     * namespace as the managed SSL certificates.
     */
    readonly name?: pulumi.Input<string>;
    readonly project?: pulumi.Input<string>;
    /**
     * Enum field whose value is always 'MANAGED' - used to signal to the API which type this is. Default value: "MANAGED"
     * Possible values: ["MANAGED"]
     */
    readonly type?: pulumi.Input<string>;
}
