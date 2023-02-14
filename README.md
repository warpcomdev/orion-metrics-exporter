# orion-metrics-exporter

Exportador de métricas de Orion para prometheus

Exporta en formato prometheus las métricas publicadas por Orion en los endpoints:

- /admin/metrics
- /statistics

## Ejecución

El ejecutable admite tres parámetros:

- `port`: puerto por el que escucha el servicio.
- `metrics`: URL de métricas de orion (por defecto, `http://localhost:1026/admin/metrics`)
- `stats`: URL de estadísticas de orion (por defecto, `http://localhost:1026/statistics`)

## Métricas

Este es un ejemplo de las métricas que se obtienen:

```
orion_stats_counters_json_requests 46715.000000
orion_stats_counters_no_payload_requests 2724.000000
orion_stats_counters_invalid_requests 0.000000
orion_stats_counters_notifications_sent 58314.000000
orion_stats_timing_accumulated {at="JsonV1Parse"} 0.059632
orion_stats_timing_accumulated {at="JsonV2Parse"} 2.941718
orion_stats_timing_accumulated {at="MongoBackend"} 574.018799
orion_stats_timing_accumulated {at="MongoReadWait"} 117.062744
orion_stats_timing_accumulated {at="MongoWriteWait"} 128.056015
orion_stats_timing_accumulated {at="MongoCommandWait"} 0.008586
orion_stats_timing_accumulated {at="Render"} 5.528586
orion_stats_timing_accumulated_total 856.544556
orion_stats_timing_last {at="JsonV2Parse"} 0.000235
orion_stats_timing_last {at="MongoBackend"} 0.009387
orion_stats_timing_last {at="MongoReadWait"} 0.001339
orion_stats_timing_last {at="MongoWriteWait"} 0.003677
orion_stats_timing_last {at="Render"} 0.000205
orion_stats_timing_last_total 0.010196
orion_stats_semwait_request 0.000000
orion_stats_semwait_db_connection_pool 1.249825
orion_stats_semwait_transaction 0.047438
orion_stats_semwait_sub_cache 377.329590
orion_stats_semwait_connection_context 0.000000
orion_stats_semwait_timestat 0.013715
orion_stats_semwait_metrics 0.143647
orion_stats_notif_queue_in 59857.000000
orion_stats_notif_queue_out 59544.000000
orion_stats_notif_queue_reject 0.000000
orion_stats_notif_queue_sent_ok 58314.000000
orion_stats_notif_queue_sent_error 1204.000000
orion_stats_notif_queue_time_in_queue 0.175028
orion_stats_notif_queue_size 0.000000
orion_stats_notif_queue_in 59857.000000
orion_stats_counters_requests {url="/v2", method="GET"} 789.000000
orion_stats_counters_requests {url="/v2", method="POST"} 0.000000
orion_stats_counters_requests {url="/v2", method="PUT"} 0.000000
orion_stats_counters_requests {url="/v2", method="PATCH"} 0.000000
orion_stats_counters_requests {url="/v2", method="DELETE"} 0.000000
orion_stats_counters_requests {url="/v2", method="HEAD"} 0.000000
orion_stats_counters_requests {url="/v2/entities", method="GET"} 504.000000
orion_stats_counters_requests {url="/v2/entities", method="POST"} 1868.000000
orion_stats_counters_requests {url="/v2/entities", method="PUT"} 0.000000
orion_stats_counters_requests {url="/v2/entities", method="PATCH"} 0.000000
orion_stats_counters_requests {url="/v2/entities", method="DELETE"} 0.000000
orion_stats_counters_requests {url="/v2/entities", method="HEAD"} 0.000000
orion_stats_counters_requests {url="/v2/subscriptions", method="GET"} 2.000000
orion_stats_counters_requests {url="/v2/subscriptions", method="POST"} 0.000000
orion_stats_counters_requests {url="/v2/subscriptions", method="PUT"} 0.000000
orion_stats_counters_requests {url="/v2/subscriptions", method="PATCH"} 0.000000
orion_stats_counters_requests {url="/v2/subscriptions", method="DELETE"} 0.000000
orion_stats_counters_requests {url="/v2/subscriptions", method="HEAD"} 0.000000
orion_stats_counters_requests {url="/statistics", method="GET"} 9.000000
orion_stats_counters_requests {url="/statistics", method="POST"} 0.000000
orion_stats_counters_requests {url="/statistics", method="PUT"} 0.000000
orion_stats_counters_requests {url="/statistics", method="PATCH"} 0.000000
orion_stats_counters_requests {url="/statistics", method="DELETE"} 0.000000
orion_stats_counters_requests {url="/statistics", method="HEAD"} 0.000000
orion_stats_counters_requests {url="/admin/metrics", method="GET"} 7.000000
orion_stats_counters_requests {url="/admin/metrics", method="POST"} 0.000000
orion_stats_counters_requests {url="/admin/metrics", method="PUT"} 0.000000
orion_stats_counters_requests {url="/admin/metrics", method="PATCH"} 0.000000
orion_stats_counters_requests {url="/admin/metrics", method="DELETE"} 0.000000
orion_stats_counters_requests {url="/admin/metrics", method="HEAD"} 0.000000
orion_stats_counters_requests {url="/v2/entities/{id}[/attrs]", method="GET"} 1380.000000
orion_stats_counters_requests {url="/v2/entities/{id}[/attrs]", method="POST"} 27323.000000
orion_stats_counters_requests {url="/v2/entities/{id}[/attrs]", method="PUT"} 0.000000
orion_stats_counters_requests {url="/v2/entities/{id}[/attrs]", method="PATCH"} 13311.000000
orion_stats_counters_requests {url="/v2/entities/{id}[/attrs]", method="DELETE"} 7.000000
orion_stats_counters_requests {url="/v2/entities/{id}[/attrs]", method="HEAD"} 0.000000
orion_stats_counters_requests {url="/v2/types", method="GET"} 3.000000
orion_stats_counters_requests {url="/v2/types", method="POST"} 0.000000
orion_stats_counters_requests {url="/v2/types", method="PUT"} 0.000000
orion_stats_counters_requests {url="/v2/types", method="PATCH"} 0.000000
orion_stats_counters_requests {url="/v2/types", method="DELETE"} 0.000000
orion_stats_counters_requests {url="/v2/types", method="HEAD"} 0.000000
orion_stats_counters_requests {url="/v2/subscriptions/{id}", method="GET"} 3.000000
orion_stats_counters_requests {url="/v2/subscriptions/{id}", method="POST"} 0.000000
orion_stats_counters_requests {url="/v2/subscriptions/{id}", method="PUT"} 0.000000
orion_stats_counters_requests {url="/v2/subscriptions/{id}", method="PATCH"} 1.000000
orion_stats_counters_requests {url="/v2/subscriptions/{id}", method="DELETE"} 0.000000
orion_stats_counters_requests {url="/v2/subscriptions/{id}", method="HEAD"} 0.000000
orion_stats_counters_requests {url="/v2/op/update", method="GET"} 0.000000
orion_stats_counters_requests {url="/v2/op/update", method="POST"} 4071.000000
orion_stats_counters_requests {url="/v2/op/update", method="PUT"} 0.000000
orion_stats_counters_requests {url="/v2/op/update", method="PATCH"} 0.000000
orion_stats_counters_requests {url="/v2/op/update", method="DELETE"} 0.000000
orion_stats_counters_requests {url="/v2/op/update", method="HEAD"} 0.000000
orion_stats_counters_requests_legacy {url="/ngsi10/queryContext", method="GET"} 0.000000
orion_stats_counters_requests_legacy {url="/ngsi10/queryContext", method="POST"} 80.000000
orion_stats_counters_requests_legacy {url="/ngsi10/queryContext", method="PUT"} 0.000000
orion_stats_counters_requests_legacy {url="/ngsi10/queryContext", method="PATCH"} 0.000000
orion_stats_counters_requests_legacy {url="/ngsi10/queryContext", method="DELETE"} 0.000000
orion_stats_counters_requests_legacy {url="/ngsi10/queryContext", method="HEAD"} 0.000000
orion_stats_counters_requests_legacy {url="/v1/contextEntities/{id}", method="GET"} 0.000000
orion_stats_counters_requests_legacy {url="/v1/contextEntities/{id}", method="POST"} 0.000000
orion_stats_counters_requests_legacy {url="/v1/contextEntities/{id}", method="PUT"} 17.000000
orion_stats_counters_requests_legacy {url="/v1/contextEntities/{id}", method="PATCH"} 0.000000
orion_stats_counters_requests_legacy {url="/v1/contextEntities/{id}", method="DELETE"} 0.000000
orion_stats_counters_requests_legacy {url="/v1/contextEntities/{id}", method="HEAD"} 0.000000
orion_stats_counters_requests_legacy {url="/v1/contextEntities/{id}/attributes/{name}", method="GET"} 2.000000
orion_stats_counters_requests_legacy {url="/v1/contextEntities/{id}/attributes/{name}", method="POST"} 0.000000
orion_stats_counters_requests_legacy {url="/v1/contextEntities/{id}/attributes/{name}", method="PUT"} 0.000000
orion_stats_counters_requests_legacy {url="/v1/contextEntities/{id}/attributes/{name}", method="PATCH"} 0.000000
orion_stats_counters_requests_legacy {url="/v1/contextEntities/{id}/attributes/{name}", method="DELETE"} 0.000000
orion_stats_counters_requests_legacy {url="/v1/contextEntities/{id}/attributes/{name}", method="HEAD"} 0.000000
orion_stats_counters_requests_legacy {url="/v1/updateContext", method="GET"} 0.000000
orion_stats_counters_requests_legacy {url="/v1/updateContext", method="POST"} 32.000000
orion_stats_counters_requests_legacy {url="/v1/updateContext", method="PUT"} 0.000000
orion_stats_counters_requests_legacy {url="/v1/updateContext", method="PATCH"} 0.000000
orion_stats_counters_requests_legacy {url="/v1/updateContext", method="DELETE"} 0.000000
orion_stats_counters_requests_legacy {url="/v1/updateContext", method="HEAD"} 0.000000
orion_stats_counters_requests_legacy {url="/ngsi10/updateContext", method="GET"} 0.000000
orion_stats_counters_requests_legacy {url="/ngsi10/updateContext", method="POST"} 11.000000
orion_stats_counters_requests_legacy {url="/ngsi10/updateContext", method="PUT"} 0.000000
orion_stats_counters_requests_legacy {url="/ngsi10/updateContext", method="PATCH"} 0.000000
orion_stats_counters_requests_legacy {url="/ngsi10/updateContext", method="DELETE"} 0.000000
orion_stats_counters_requests_legacy {url="/ngsi10/updateContext", method="HEAD"} 0.000000
orion_outgoing_transaction_request_size {service="alcobendas", subservice="contadores"} 3256.000000
orion_outgoing_transactions {service="alcobendas", subservice="contadores"} 3.000000
orion_incoming_transaction_request_size {service="alcobendas", subservice="contadores"} 869.000000
orion_service_time {service="alcobendas", subservice="contadores"} 0.015255
orion_incoming_transactions {service="alcobendas", subservice="contadores"} 4.000000
orion_incoming_transaction_request_size {service="alcobendas", subservice="lamoraleja/contadores"} 200.000000
orion_service_time {service="alcobendas", subservice="lamoraleja/contadores"} 0.004769
orion_incoming_transactions {service="alcobendas", subservice="lamoraleja/contadores"} 1.000000
orion_incoming_transaction_request_size {service="alcobendas", subservice="patrimonio"} 18235.000000
orion_service_time {service="alcobendas", subservice="patrimonio"} 0.007474
orion_incoming_transactions {service="alcobendas", subservice="patrimonio"} 57.000000
orion_outgoing_transaction_request_size {service="alcobendas", subservice="patrimonio"} 34338.000000
orion_outgoing_transactions {service="alcobendas", subservice="patrimonio"} 57.000000
orion_incoming_transaction_request_size {service="alcobendas", subservice="wifi"} 27011.000000
orion_service_time {service="alcobendas", subservice="wifi"} 0.008933
orion_incoming_transactions {service="alcobendas", subservice="wifi"} 45.000000
orion_outgoing_transaction_request_size {service="alcobendas", subservice="wifi"} 109732.000000
orion_outgoing_transactions {service="alcobendas", subservice="wifi"} 42.000000
orion_incoming_transaction_response_size {service="alcobendas", subservice="activos"} 17400.000000
orion_service_time {service="alcobendas", subservice="activos"} 0.003412
orion_incoming_transactions {service="alcobendas", subservice="activos"} 12.000000
orion_outgoing_transaction_request_size {service="alcobendas", subservice="lamoraleja/alumbrado"} 29701.000000
orion_outgoing_transaction_response_size {service="alcobendas", subservice="lamoraleja/alumbrado"} 288.000000
orion_outgoing_transactions {service="alcobendas", subservice="lamoraleja/alumbrado"} 17.000000
orion_incoming_transaction_request_size {service="alcobendas", subservice="lamoraleja/alumbrado"} 3961.000000
orion_service_time {service="alcobendas", subservice="lamoraleja/alumbrado"} 0.047706
orion_incoming_transactions {service="alcobendas", subservice="lamoraleja/alumbrado"} 9.000000
orion_incoming_transactions {service="alcobendas", subservice="lamoraleja/residuos"} 3.000000
orion_outgoing_transaction_request_size {service="alcobendas", subservice="lamoraleja/residuos"} 2640.000000
orion_outgoing_transaction_response_size {service="alcobendas", subservice="lamoraleja/residuos"} 72.000000
orion_outgoing_transactions {service="alcobendas", subservice="lamoraleja/residuos"} 4.000000
orion_incoming_transaction_request_size {service="alcobendas", subservice="lamoraleja/residuos"} 818.000000
orion_service_time {service="alcobendas", subservice="lamoraleja/residuos"} 0.005485
orion_outgoing_transaction_request_size {service="alcobendas", subservice="playas"} 106958.000000
orion_outgoing_transaction_response_size {service="alcobendas", subservice="playas"} 2232.000000
orion_outgoing_transactions {service="alcobendas", subservice="playas"} 133.000000
orion_incoming_transaction_request_size {service="alcobendas", subservice="playas"} 23288.000000
orion_service_time {service="alcobendas", subservice="playas"} 0.052177
orion_incoming_transactions {service="alcobendas", subservice="playas"} 84.000000
orion_service_time {service="alcobendas", subservice="riego"} 0.008248
orion_incoming_transactions {service="alcobendas", subservice="riego"} 23.000000
orion_outgoing_transaction_request_size {service="alcobendas", subservice="riego"} 15835.000000
orion_outgoing_transactions {service="alcobendas", subservice="riego"} 23.000000
orion_incoming_transaction_request_size {service="alcobendas", subservice="riego"} 8278.000000
orion_incoming_transaction_request_size {service="alcobendas", subservice="valdelasfuentes/residuos"} 1642.000000
orion_service_time {service="alcobendas", subservice="valdelasfuentes/residuos"} 0.005906
orion_incoming_transactions {service="alcobendas", subservice="valdelasfuentes/residuos"} 7.000000
orion_outgoing_transaction_request_size {service="alcobendas", subservice="valdelasfuentes/residuos"} 2555.000000
orion_outgoing_transactions {service="alcobendas", subservice="valdelasfuentes/residuos"} 3.000000
orion_incoming_transaction_request_size {service="alcobendas", subservice="valdelasfuentes/riego"} 4960.000000
orion_service_time {service="alcobendas", subservice="valdelasfuentes/riego"} 0.051172
orion_incoming_transactions {service="alcobendas", subservice="valdelasfuentes/riego"} 13.000000
orion_outgoing_transaction_request_size {service="alcobendas", subservice="valdelasfuentes/riego"} 10554.000000
orion_outgoing_transactions {service="alcobendas", subservice="valdelasfuentes/riego"} 13.000000
orion_incoming_transactions {service="alcobendas", subservice="lamoraleja/riego"} 12.000000
orion_outgoing_transaction_request_size {service="alcobendas", subservice="lamoraleja/riego"} 7677.000000
orion_outgoing_transactions {service="alcobendas", subservice="lamoraleja/riego"} 12.000000
orion_incoming_transaction_request_size {service="alcobendas", subservice="lamoraleja/riego"} 4585.000000
orion_service_time {service="alcobendas", subservice="lamoraleja/riego"} 0.006386
orion_service_time {service="alcobendas", subservice="valdelasfuentes/alumbrado"} 0.007509
orion_incoming_transactions {service="alcobendas", subservice="valdelasfuentes/alumbrado"} 24.000000
orion_outgoing_transaction_request_size {service="alcobendas", subservice="valdelasfuentes/alumbrado"} 81088.000000
orion_outgoing_transaction_response_size {service="alcobendas", subservice="valdelasfuentes/alumbrado"} 828.000000
orion_outgoing_transactions {service="alcobendas", subservice="valdelasfuentes/alumbrado"} 47.000000
orion_incoming_transaction_request_size {service="alcobendas", subservice="valdelasfuentes/alumbrado"} 9030.000000
orion_incoming_transaction_request_size {service="alcobendas", subservice="valdelasfuentes/contadores"} 692.000000
orion_service_time {service="alcobendas", subservice="valdelasfuentes/contadores"} 0.013359
orion_incoming_transactions {service="alcobendas", subservice="valdelasfuentes/contadores"} 3.000000
orion_outgoing_transaction_request_size {service="alcobendas", subservice="valdelasfuentes/contadores"} 2022.000000
orion_outgoing_transactions {service="alcobendas", subservice="valdelasfuentes/contadores"} 3.000000
orion_incoming_transaction_request_size {service="alcobendas", subservice="vehiculoelectrico"} 819.000000
orion_service_time {service="alcobendas", subservice="vehiculoelectrico"} 0.005688
orion_incoming_transactions {service="alcobendas", subservice="vehiculoelectrico"} 3.000000
orion_outgoing_transaction_request_size {service="alcobendas", subservice="vehiculoelectrico"} 3450.000000
orion_outgoing_transactions {service="alcobendas", subservice="vehiculoelectrico"} 3.000000
orion_service_time {service="alcobendas", subservice="agua"} 0.007785
orion_incoming_transactions {service="alcobendas", subservice="agua"} 5.000000
orion_outgoing_transaction_request_size {service="alcobendas", subservice="agua"} 7706.000000
orion_outgoing_transactions {service="alcobendas", subservice="agua"} 5.000000
orion_incoming_transaction_request_size {service="alcobendas", subservice="agua"} 2941.000000
orion_incoming_transaction_request_size {service="alcobendas", subservice="alumbrado"} 31949.000000
orion_service_time {service="alcobendas", subservice="alumbrado"} 0.011609
orion_incoming_transactions {service="alcobendas", subservice="alumbrado"} 75.000000
orion_outgoing_transaction_request_size {service="alcobendas", subservice="alumbrado"} 191666.000000
orion_outgoing_transaction_response_size {service="alcobendas", subservice="alumbrado"} 2520.000000
orion_outgoing_transactions {service="alcobendas", subservice="alumbrado"} 145.000000
orion_service_time {service="alcobendas", subservice="medioambiente"} 0.006666
orion_incoming_transactions {service="alcobendas", subservice="medioambiente"} 10.000000
orion_outgoing_transaction_request_size {service="alcobendas", subservice="medioambiente"} 29682.000000
orion_outgoing_transaction_response_size {service="alcobendas", subservice="medioambiente"} 36.000000
orion_outgoing_transactions {service="alcobendas", subservice="medioambiente"} 11.000000
orion_incoming_transaction_request_size {service="alcobendas", subservice="medioambiente"} 18993.000000
orion_service_time {service="alcobendas", subservice="residuos"} 0.287193
orion_incoming_transactions {service="alcobendas", subservice="residuos"} 55.000000
orion_outgoing_transaction_request_size {service="alcobendas", subservice="residuos"} 26521.000000
orion_outgoing_transaction_response_size {service="alcobendas", subservice="residuos"} 900.000000
orion_outgoing_transactions {service="alcobendas", subservice="residuos"} 50.000000
orion_incoming_transaction_request_size {service="alcobendas", subservice="residuos"} 9907.000000
```
