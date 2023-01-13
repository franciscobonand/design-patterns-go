# Factory pattern

- Ways of controlling how an object is constructed
    - Used when object creation logic is too convoluted
- Factories are components responsible for wholesale object creation (non-piecewise, unlike Builder)
    - Using a separate function - Factory function (constructor)
    - Using a separated struct - Factory (can be passed as a parameter) 
- A Factory function is a helper function for making struct instances
- A factory is any entity that can take care of object creation
    - Can be a function or a dedicated struct
