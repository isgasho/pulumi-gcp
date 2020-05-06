// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class Table extends pulumi.CustomResource {
    /**
     * Get an existing Table resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TableState, opts?: pulumi.CustomResourceOptions): Table {
        return new Table(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:bigtable/table:Table';

    /**
     * Returns true if the given object is an instance of Table.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Table {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Table.__pulumiType;
    }

    /**
     * A group of columns within a table which share a common configuration. This can be specified multiple times. Structure is documented below.
     */
    public readonly columnFamilies!: pulumi.Output<outputs.bigtable.TableColumnFamily[] | undefined>;
    /**
     * The name of the Bigtable instance.
     */
    public readonly instanceName!: pulumi.Output<string>;
    /**
     * The name of the table.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    public readonly splitKeys!: pulumi.Output<string[] | undefined>;

    /**
     * Create a Table resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TableArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TableArgs | TableState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TableState | undefined;
            inputs["columnFamilies"] = state ? state.columnFamilies : undefined;
            inputs["instanceName"] = state ? state.instanceName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["splitKeys"] = state ? state.splitKeys : undefined;
        } else {
            const args = argsOrState as TableArgs | undefined;
            if (!args || args.instanceName === undefined) {
                throw new Error("Missing required property 'instanceName'");
            }
            inputs["columnFamilies"] = args ? args.columnFamilies : undefined;
            inputs["instanceName"] = args ? args.instanceName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["splitKeys"] = args ? args.splitKeys : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Table.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Table resources.
 */
export interface TableState {
    /**
     * A group of columns within a table which share a common configuration. This can be specified multiple times. Structure is documented below.
     */
    readonly columnFamilies?: pulumi.Input<pulumi.Input<inputs.bigtable.TableColumnFamily>[]>;
    /**
     * The name of the Bigtable instance.
     */
    readonly instanceName?: pulumi.Input<string>;
    /**
     * The name of the table.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly splitKeys?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Table resource.
 */
export interface TableArgs {
    /**
     * A group of columns within a table which share a common configuration. This can be specified multiple times. Structure is documented below.
     */
    readonly columnFamilies?: pulumi.Input<pulumi.Input<inputs.bigtable.TableColumnFamily>[]>;
    /**
     * The name of the Bigtable instance.
     */
    readonly instanceName: pulumi.Input<string>;
    /**
     * The name of the table.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly splitKeys?: pulumi.Input<pulumi.Input<string>[]>;
}
