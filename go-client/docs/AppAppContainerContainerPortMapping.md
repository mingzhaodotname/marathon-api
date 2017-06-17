# AppAppContainerContainerPortMapping

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerPort** | [**AppNumberAnyPort**](app.number.AnyPort.md) | Refers to the port the application listens to inside of the container. It is optional and defaults to 0. For each containerPort with a value of 0 Marathon assigns the containerPort the same value as the assigned hostPort. This is especially useful for apps that advertise the port they are listening on to the outside world for P2P communication. Without containerPort: 0 they would erroneously advertise their private container port which is usually not the same as the externally visible host port.  | [default to null]
**HostPort** | [**AppNumberAnyPort**](app.number.AnyPort.md) | Retains the traditional meaning in Marathon, which is a random port from the range included in the Mesos resource offer. The resulting host ports for each task are exposed via the task details in the REST API and the Marathon web UI. hostPort is optional. In BRIDGE mode it defaults to 0 if left unspecified. In USER mode an unspecified hostPort does not allocate a port from a Mesos offer.  | [optional] [default to null]
**Labels** | [**AppLabelKvLabels**](app.label.KVLabels.md) | This can be used to add metadata to be interpreted by external applications such as firewalls. | [optional] [default to null]
**Name** | [**StringsLegacyName**](strings.LegacyName.md) | Name of the service hosted on this port. If provided, it must be unique over all port mappings. | [optional] [default to null]
**Protocol** | [**StringsNetworkProtocol**](strings.NetworkProtocol.md) |  | [optional] [default to null]
**ServicePort** | [**AppNumberAnyPort**](app.number.AnyPort.md) | Is a helper port intended for doing service discovery using a well-known port per service. The assigned servicePort value is not used/interpreted by Marathon itself but supposed to used by load balancer infrastructure. See Service Discovery Load Balancing doc page. The servicePort parameter is optional and defaults to 0. Like hostPort, If the value is 0, a random port will be assigned. If a servicePort value is assigned by Marathon then Marathon guarantees that its value is unique across the cluster. The values for random service ports are in the range [local_port_min, local_port_max] where local_port_min and local_port_max are command line options with default values of 10000 and 20000, respectively.  | [optional] [default to null]
**NetworkNames** | [**[]StringsName**](strings.Name.md) | List of the container networks associated with this mapping. If absent, then this mapping is associated with all defined container networks (for this application). A single item list is mandatory when &#x60;hostPort&#x60; is specified and multiple container networks are defined.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


