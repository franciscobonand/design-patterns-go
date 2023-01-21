# Bridge pattern

- Prevents 'Cartesian product' complexity explosion
    - Example:
        - Common type ThreadScheduler
        - Can be preemptive or cooperative
        - Can run on Windows or Unix
        - This ends up with a 2x2 scenario (WindowsPTS, UnixPTS, WindowsCTS, UnixCTS)
- Bridge patterns avoids the entity explosion
- Bridge is a mechanism that decouples an interface from an implementation
    - Both exist as hierarchies
- Can be thought as a stronger form of encapsulation
