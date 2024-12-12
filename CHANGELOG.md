# Changelog

## [0.5.4](https://github.com/onfleet/gonfleet/compare/v0.5.3...v0.5.4) - 2024-12-12
* Add
    * `RoutePlan` API

## [0.5.3](https://github.com/onfleet/gonfleet/compare/v0.5.2...v0.5.3) - 2024-12-10
* Change
    * use `cenkalti/backoff` package for rate limiting / exp backoff
* Add
    * task `Dependencies` to `TaskListQueryParams`

## [0.5.2](https://github.com/onfleet/gonfleet/compare/v0.5.1...v0.5.2) - 2024-12-10
* Add
    * task `CustomFields` support
* Fix
    * task `TriangleColor`

## [0.5.1](https://github.com/onfleet/gonfleet/compare/v0.5.0...v0.5.1) - 2024-01-04
* Add
    * providers/manifest `Generate`

## [0.5.0](https://github.com/onfleet/gonfleet/compare/v0.4.0...v0.5.0) - 2023-10-18
* Add
    * `Tasks.BatchCreateAsync`
    * `Tasks.GetBatchJobStatus`

## [0.4.0](https://github.com/onfleet/gonfleet/compare/v0.3.1...v0.4.0) - 2023-06-01
* Change
    * `Worker.ListTasks` parameter changed from `WorkerTasksListQueryParams` to `*WorkerTasksListQueryParams`
    * `Teams.ListTasks` parameter changed from `TeamTasksListQueryParams` to `*TeamTasksListQueryParams`

## [0.3.1](https://github.com/onfleet/gonfleet/compare/v0.3.0...v0.3.1) - 2023-04-06
* Change
    * `TaskListQueryParams` field `LastID` to `LastId`

## [0.3.0](https://github.com/onfleet/gonfleet/compare/v0.2.3...v0.3.0) - 2023-04-06
* Add
    * task `List`

## [0.2.3](https://github.com/onfleet/gonfleet/compare/v0.2.2...v0.2.3) - 2023-03-31
* Fix
    * resync changelog with current version

## [0.2.2](https://github.com/onfleet/gonfleet/compare/v0.2.1...v0.2.2) - 2023-03-31
* Change
    * README task creation example to use RecipientCreateParams
* Fix
    * `DestinationCreateParams` to included `Metadata`

## [0.2.1](https://github.com/onfleet/gonfleet/compare/v0.2.0...v0.2.1) - 2023-03-30
* Remove
    * util package. no longer needed / relevant

## [0.2.0](https://github.com/onfleet/gonfleet/compare/v0.1.6...v0.2.0) - 2023-03-30
* Add
    * README example of task creation
* Fix
    * task creation comment suggests `DestinationCreateParams` over `Destination`

## [0.1.6](https://github.com/onfleet/gonfleet/compare/v0.1.5...v0.1.6) - 2023-03-30
* Add
    * list admins, workers, destinations, tasks, and recipients with metdata query. Method per service named `ListWithMetadataQuery`
    * task `Create`, `BatchCreate`, `Update`, `ForceComplete`, `Clone`, `Delete`, `AutoAssignMulti`
* Change
    * netw package renamed netwrk

## [0.1.5](https://github.com/onfleet/gonfleet/compare/v0.1.4...v0.1.5) - 2023-03-26
* Change
    * service handler comments only reference official onfleet api docs via url

## [0.1.4](https://github.com/onfleet/gonfleet/compare/v0.1.3...v0.1.4) - 2023-03-26
* Add
    * team `ListTasks`, `GetWorkerEta`, `AutoDispatch`, `Get`

## [0.1.3](https://github.com/onfleet/gonfleet/compare/v0.1.2...v0.1.3) - 2023-03-26
* Add
    * worker `SetSchedule`, `ListWorkersByLocation`, `ListTasks`, `Update`, `Delete`

## [0.1.2](https://github.com/onfleet/gonfleet/compare/v0.1.1...v0.1.2) - 2023-03-26
* Add
    * init README.md
    * worker `Create`, `GetWithQuery`, `ListWithQuery`

## [0.1.1](https://github.com/onfleet/gonfleet/compare/v0.1.0...v0.1.1) - 2023-03-26
* Change
    * pkg `Name` changed to onfleet/gonfleet 

## 0.1.0 - 2023-03-26
* Change
    * CHANGELOG.md structure / formatting
