# Builder pattern

- Some objects are simple and can be created in a single constructor call
- Other objects require a lot of ceremony to create:
    - Factory function with 10 arguments
- For those cases, opt for a piecewise construction
- **Builder provides an API for constructing an object step-by-step**
    - For objects with distinct sets of information (e.g. address and job data of a person) multiple builders can be used (se `facets.go`)

