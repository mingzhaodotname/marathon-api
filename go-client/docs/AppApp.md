# AppApp

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**StringsPathId**](strings.PathId.md) |  | [default to null]
**AcceptedResourceRoles** | **[]string** | A list of resource roles. Marathon considers only resource offers with roles in this list for launching tasks of this app. If you do not specify this, Marathon considers all resource offers with roles that have been configured by the &#x60;--default_accepted_resource_roles&#x60; command line flag. If no &#x60;--default_accepted_resource_roles&#x60; was given on startup, Marathon considers all resource offers. To register Marathon for a role, you need to specify the &#x60;--mesos_role&#x60; command line flag on startup. If you want to assign all resources of a slave to a role, you can use the &#x60;--default_role&#x60; argument when starting up the slave. If you need a more fine-grained configuration, you can use the &#x60;--resources&#x60; argument to specify resource shares per role. See [the Mesos attribute and resources documentation](http://mesos.apache.org/documentation/latest/attributes-resources/) for details  | [optional] [default to null]
**Args** | **[]string** | An array of strings that represents an alternative mode of specifying the command to run. This was motivated by safe usage of containerizer features like a custom Docker ENTRYPOINT. This args field may be used in place of cmd even when using the default command executor. This change mirrors API and semantics changes in the Mesos CommandInfo protobuf message starting with version &#x60;0.20.0&#x60;. Either &#x60;cmd&#x60; or &#x60;args&#x60; must be supplied. It is invalid to supply both &#x60;cmd&#x60; and &#x60;args&#x60; in the same app.  | [optional] [default to null]
**BackoffFactor** | **float64** | Configures exponential backoff behavior when launching potentially sick apps. This prevents sandboxes associated with consecutively failing tasks from filling up the hard disk on Mesos slaves. The backoff period is multiplied by the factor for each consecutive failure until it reaches maxLaunchDelaySeconds. This applies also to tasks that are killed due to failing too many health checks.  | [optional] [default to 1.15]
**BackoffSeconds** | **int32** | Configures exponential backoff behavior when launching potentially sick apps. This prevents sandboxes associated with consecutively failing tasks from filling up the hard disk on Mesos slaves. The backoff period is multiplied by the factor for each consecutive failure until it reaches maxLaunchDelaySeconds. This applies also to tasks that are killed due to failing too many health checks.  | [optional] [default to 1]
**Cmd** | **string** | The command that is executed.  This value is wrapped by Mesos via &#x60;/bin/sh -c ${app.cmd}&#x60;. Either &#x60;cmd&#x60; or &#x60;args&#x60; must be supplied. It is invalid to supply both &#x60;cmd&#x60; and &#x60;args&#x60; in the same app.  | [optional] [default to null]
**Constraints** | [**[]AppConstraintsAppConstraint**](app.constraints.AppConstraint.md) |  | [optional] [default to null]
**Container** | [**AppAppContainerContainer**](app.appContainer.Container.md) |  | [optional] [default to null]
**Cpus** | **float64** | The number of CPU shares this application needs per instance. This number does not have to be integer, but can be a fraction.\&quot;,  | [optional] [default to 1.0]
**Dependencies** | [**[]StringsPathId**](strings.PathId.md) | A list of services upon which this application depends An order is derived from the dependencies for performing start/stop and upgrade of the application. For example, an application /a relies on the services /b which itself relies on /c. To start all 3 applications, first /c is started than /b than /a.  | [optional] [default to null]
**Disk** | **float64** | How much disk space is needed for this application. This number does not have to be an integer, but can be a fraction.  | [optional] [default to 0.0]
**Env** | [**AppEnvLegacyEnvVars**](app.env.LegacyEnvVars.md) |  | [optional] [default to null]
**Executor** | **string** | The executor to use to launch this application. Different executors are available. The simplest one (and the one configured by default if none is given) is &#x60;//cmd&#x60;, which takes the cmd and executes that on the shell level.  | [optional] [default to null]
**Fetch** | [**[]AppArtifactArtifact**](app.artifact.Artifact.md) | Provided URIs are passed to Mesos fetcher module and resolved in runtime. | [optional] [default to null]
**HealthChecks** | [**[]AppHealthAppHealthCheck**](app.health.AppHealthCheck.md) |  | [optional] [default to null]
**Instances** | **int32** | The number of instances of this application to start. Please note: this number can be changed any time as needed to scale the application.  | [optional] [default to 1]
**Labels** | [**AppLabelKvLabels**](app.label.KVLabels.md) |  | [optional] [default to null]
**MaxLaunchDelaySeconds** | **int32** | Configures exponential backoff behavior when launching potentially sick apps. This prevents sandboxes associated with consecutively failing tasks from filling up the hard disk on Mesos slaves. The backoff period is multiplied by the factor for each consecutive failure until it reaches maxLaunchDelaySeconds. This applies also to tasks that are killed due to failing too many health checks.  | [optional] [default to 3600]
**Mem** | **float64** | The amount of memory in MB that is needed for the application per instance.  | [optional] [default to 128.0]
**Gpus** | **int32** | The amount of GPU cores that is needed for the application per instance.  | [optional] [default to 0]
**IpAddress** | [**AppNetworkIpAddress**](app.network.IpAddress.md) |  | [optional] [default to null]
**Networks** | [**[]AppNetworkNetwork**](app.network.Network.md) |  | [optional] [default to null]
**Ports** | [**[]AppNumberAnyPort**](app.number.AnyPort.md) | An array of required port resources on the agent host. The number of items in the array determines how many dynamic ports are allocated for every task. For every port that is zero, a globally unique (cluster-wide) port is assigned and provided as part of the app definition to be used in load balancing definitions.  | [optional] [default to null]
**PortDefinitions** | [**[]AppNetworkPortDefinition**](app.network.PortDefinition.md) | An array of required port resources on the agent host. The number of items in the array determines how many dynamic ports are allocated for every task. For every port definition with port number zero, a globally unique (cluster-wide) service port is assigned and provided as part of the app definition to be used in load balancing definitions.  | [optional] [default to null]
**ReadinessChecks** | [**[]AppReadinessReadinessCheck**](app.readiness.ReadinessCheck.md) |  | [optional] [default to null]
**Residency** | [**AppAppResidency**](app.AppResidency.md) |  | [optional] [default to null]
**RequirePorts** | **bool** | Applies only for host networking. Normally, the host ports of your tasks are automatically assigned. This corresponds to the requirePorts value false which is the default. If you need more control and want to specify your host ports in advance, you can set requirePorts to true. This way the ports you have specified are used as host ports. That also means that Marathon can schedule the associated tasks only on hosts that have the specified ports available.  | [optional] [default to null]
**Secrets** | [**AppSecretsSecrets**](app.secrets.Secrets.md) |  | [optional] [default to null]
**TaskKillGracePeriodSeconds** | **int32** | Configures the number of seconds between escalating from SIGTERM to SIGKILL when signalling tasks to terminate. Using this grace period, tasks should perform orderly shut down immediately upon receiving SIGTERM.  | [optional] [default to null]
**UpgradeStrategy** | [**AppUpgradeStrategy**](app.UpgradeStrategy.md) |  | [optional] [default to null]
**Uris** | [**[]StringsUri**](strings.Uri.md) | URIs defined here are resolved, before the application gets started. If the application has external dependencies, they should be defined here.  | [optional] [default to null]
**User** | **string** | The user to use to run the tasks on the agent. | [optional] [default to null]
**Version** | [**time.Time**](time.Time.md) | The version of this definition | [optional] [default to null]
**VersionInfo** | [**AppVersionInfoVersionInfo**](app.versionInfo.VersionInfo.md) |  | [optional] [default to null]
**KillSelection** | [**AppKillSelectionKillSelection**](app.killSelection.KillSelection.md) |  | [optional] [default to null]
**UnreachableStrategy** | [**AppUnreachableStrategyUnreachableStrategy**](app.unreachableStrategy.UnreachableStrategy.md) |  | [optional] [default to null]
**Tty** | [**AppContainerCommonsTty**](app.containerCommons.TTY.md) | Describes the information about (pseudo) TTY that can be attached to the process of this container.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


