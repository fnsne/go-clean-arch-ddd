# go-clean-arch-ddd

This is an simple template for clean architecture in golang with Domain Driven Design.

`interface` put all the interface for call out and call in.
> the position to put the interfaces this module will use.
> The Clean Architecture's port

`infra` is put all the implementation for call out and call in.
> the position to put the implementation for this module will use.
> The Clean Architecture's adapter

This template can use for monolithic or microservice architecture.
> For monolithic, you can put all module in to monomain module with one main.go.
> For microservice, you can launch main.go within each module.


## How to use
1. Clone this repository
2. every time you want to create new module, just copy `template` folder and rename it to your module name.

## account is a module example