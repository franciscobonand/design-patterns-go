# Adapter pattern

- A construct which adapts an existing interface X to conform to the required interface Y
- To create an Adapter:
    - Determine the API you have and the API you need
    - Create a component which aggregates (has a pointer to) the adaptee
    - Intermediate representations can pile up (use caching and other optimizations)
