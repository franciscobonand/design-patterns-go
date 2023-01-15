# Prototype pattern

- A partially or fully initialized object that you copy (clone) and make use of
    - Requires deep copy support
- An existing (partially or fully constructed) design is a Prototype
- Cloning can be made convenient (e.g., via a Factory)
- To implement a prototype:
    - Partially construct an object and store it somewhere
    - Deep copy the prototype
    - Customize the resulting instance
- A prototype factory provides a convenient API for using prototypes
