/*
Package patron framework

Patron is a framework for creating microservices.

Patron is french for template or pattern, but it means also boss which we found out later (no pun intended).

The entry point of the framework is the Service.
The Service uses WithComponents to handle the processing of sync and async requests.
The Service can set up as many components as it wants, even multiple HTTP components provided the port does not collide.
The Service starts by default an HTTP component which hosts the debug, alive, ready and metric endpoints.
Any other endpoints will be added to the default HTTP Component as Routes.
The Service set's up by default logging with zerolog, tracing and metrics with jaeger and prometheus.

Patron provides abstractions for the following functionality of the framework:

  - Service, which orchestrates everything
  - components and processors, which provide an abstraction of adding processing functionality to the Service
  - asynchronous message processing (RabbitMQ, Kafka)
  - synchronous processing (HTTP)
  - metrics and tracing
  - logging
  - configuration management

Patron provides same defaults for making the usage as simple as possible.
For more details please check out the GitHub repository https://github.com/beatlabs/patron.
*/
package patron
