// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * An SslCertificate resource, used for HTTPS load balancing.  This resource
 * represents a certificate for which the certificate secrets are created and
 * managed by Google.
 * 
 * For a resource where you provide the key, see the
 * SSL Certificate resource.
 * 
 * To get more information about ManagedSslCertificate, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/sslCertificates)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/load-balancing/docs/ssl-certificates)
 * 
 * > **Warning:** This resource should be used with extreme caution!  Provisioning an SSL
 * certificate is complex.  Ensure that you understand the lifecycle of a
 * certificate before attempting complex tasks like cert rotation automatically.
 * This resource will "return" as soon as the certificate object is created,
 * but post-creation the certificate object will go through a "provisioning"
 * process.  The provisioning process can complete only when the domain name
 * for which the certificate is created points to a target pool which, itself,
 * points at the certificate.  Depending on your DNS provider, this may take
 * some time, and migrating from self-managed certificates to Google-managed
 * certificates may entail some downtime while the certificate provisions.
 * 
 * In conclusion: Be extremely cautious.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_managed_ssl_certificate.html.markdown.
 */
export class ManagedSslCertificate extends pulumi.CustomResource {
    /**
     * Get an existing ManagedSslCertificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ManagedSslCertificateState, opts?: pulumi.CustomResourceOptions): ManagedSslCertificate {
        return new ManagedSslCertificate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/managedSslCertificate:ManagedSslCertificate';

    /**
     * Returns true if the given object is an instance of ManagedSslCertificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ManagedSslCertificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ManagedSslCertificate.__pulumiType;
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
     * Properties relevant to a managed certificate. These will be used if the certificate is managed (as indicated by a
     * value of 'MANAGED' in 'type').
     */
    public readonly managed!: pulumi.Output<outputs.compute.ManagedSslCertificateManaged | undefined>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long,
     * and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
     * '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash. These are
     * in the same namespace as the managed SSL certificates.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Domains associated with the certificate via Subject Alternative Name.
     */
    public /*out*/ readonly subjectAlternativeNames!: pulumi.Output<string[]>;
    /**
     * Enum field whose value is always 'MANAGED' - used to signal to the API which type this is.
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a ManagedSslCertificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ManagedSslCertificateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ManagedSslCertificateArgs | ManagedSslCertificateState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ManagedSslCertificateState | undefined;
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
            const args = argsOrState as ManagedSslCertificateArgs | undefined;
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
        const aliasOpts = { aliases: [{ type: "gcp:compute/mangedSslCertificate:MangedSslCertificate" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ManagedSslCertificate.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ManagedSslCertificate resources.
 */
export interface ManagedSslCertificateState {
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
     * Properties relevant to a managed certificate. These will be used if the certificate is managed (as indicated by a
     * value of 'MANAGED' in 'type').
     */
    readonly managed?: pulumi.Input<inputs.compute.ManagedSslCertificateManaged>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long,
     * and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
     * '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash. These are
     * in the same namespace as the managed SSL certificates.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * Domains associated with the certificate via Subject Alternative Name.
     */
    readonly subjectAlternativeNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Enum field whose value is always 'MANAGED' - used to signal to the API which type this is.
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ManagedSslCertificate resource.
 */
export interface ManagedSslCertificateArgs {
    /**
     * The unique identifier for the resource.
     */
    readonly certificateId?: pulumi.Input<number>;
    /**
     * An optional description of this resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Properties relevant to a managed certificate. These will be used if the certificate is managed (as indicated by a
     * value of 'MANAGED' in 'type').
     */
    readonly managed?: pulumi.Input<inputs.compute.ManagedSslCertificateManaged>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long,
     * and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
     * '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash. These are
     * in the same namespace as the managed SSL certificates.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Enum field whose value is always 'MANAGED' - used to signal to the API which type this is.
     */
    readonly type?: pulumi.Input<string>;
}
