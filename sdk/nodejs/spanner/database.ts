// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * A Cloud Spanner Database which is hosted on a Spanner instance.
 *
 * To get more information about Database, see:
 *
 * * [API documentation](https://cloud.google.com/spanner/docs/reference/rest/v1/projects.instances.databases)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/spanner/)
 *
 * > **Warning:** It is strongly recommended to set `lifecycle { preventDestroy = true }` on databases in order to prevent accidental data loss.
 *
 * ## Example Usage
 * ### Spanner Database Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const main = new gcp.spanner.Instance("main", {
 *     config: "regional-europe-west1",
 *     displayName: "main-instance",
 * });
 * const database = new gcp.spanner.Database("database", {
 *     instance: main.name,
 *     ddls: [
 *         "CREATE TABLE t1 (t1 INT64 NOT NULL,) PRIMARY KEY(t1)",
 *         "CREATE TABLE t2 (t2 INT64 NOT NULL,) PRIMARY KEY(t2)",
 *     ],
 *     deletionProtection: false,
 * });
 * ```
 *
 * ## Import
 *
 * Database can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:spanner/database:Database default projects/{{project}}/instances/{{instance}}/databases/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:spanner/database:Database default instances/{{instance}}/databases/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:spanner/database:Database default {{project}}/{{instance}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:spanner/database:Database default {{instance}}/{{name}}
 * ```
 */
export class Database extends pulumi.CustomResource {
    /**
     * Get an existing Database resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatabaseState, opts?: pulumi.CustomResourceOptions): Database {
        return new Database(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:spanner/database:Database';

    /**
     * Returns true if the given object is an instance of Database.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Database {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Database.__pulumiType;
    }

    /**
     * An optional list of DDL statements to run inside the newly created
     * database. Statements can create tables, indexes, etc. These statements
     * execute atomically with the creation of the database: if there is an
     * error in any statement, the database is not created.
     */
    public readonly ddls!: pulumi.Output<string[] | undefined>;
    public readonly deletionProtection!: pulumi.Output<boolean | undefined>;
    /**
     * The instance to create the database on.
     */
    public readonly instance!: pulumi.Output<string>;
    /**
     * A unique identifier for the database, which cannot be changed after
     * the instance is created. Values are of the form [a-z][-a-z0-9]*[a-z0-9].
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * An explanation of the status of the database.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;

    /**
     * Create a Database resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabaseArgs | DatabaseState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DatabaseState | undefined;
            inputs["ddls"] = state ? state.ddls : undefined;
            inputs["deletionProtection"] = state ? state.deletionProtection : undefined;
            inputs["instance"] = state ? state.instance : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["state"] = state ? state.state : undefined;
        } else {
            const args = argsOrState as DatabaseArgs | undefined;
            if (!args || args.instance === undefined) {
                throw new Error("Missing required property 'instance'");
            }
            inputs["ddls"] = args ? args.ddls : undefined;
            inputs["deletionProtection"] = args ? args.deletionProtection : undefined;
            inputs["instance"] = args ? args.instance : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["state"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Database.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Database resources.
 */
export interface DatabaseState {
    /**
     * An optional list of DDL statements to run inside the newly created
     * database. Statements can create tables, indexes, etc. These statements
     * execute atomically with the creation of the database: if there is an
     * error in any statement, the database is not created.
     */
    readonly ddls?: pulumi.Input<pulumi.Input<string>[]>;
    readonly deletionProtection?: pulumi.Input<boolean>;
    /**
     * The instance to create the database on.
     */
    readonly instance?: pulumi.Input<string>;
    /**
     * A unique identifier for the database, which cannot be changed after
     * the instance is created. Values are of the form [a-z][-a-z0-9]*[a-z0-9].
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * An explanation of the status of the database.
     */
    readonly state?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Database resource.
 */
export interface DatabaseArgs {
    /**
     * An optional list of DDL statements to run inside the newly created
     * database. Statements can create tables, indexes, etc. These statements
     * execute atomically with the creation of the database: if there is an
     * error in any statement, the database is not created.
     */
    readonly ddls?: pulumi.Input<pulumi.Input<string>[]>;
    readonly deletionProtection?: pulumi.Input<boolean>;
    /**
     * The instance to create the database on.
     */
    readonly instance: pulumi.Input<string>;
    /**
     * A unique identifier for the database, which cannot be changed after
     * the instance is created. Values are of the form [a-z][-a-z0-9]*[a-z0-9].
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}
