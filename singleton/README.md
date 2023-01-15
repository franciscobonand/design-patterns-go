# Singleton pattern

- For some components it only makes sense to have one in the system
    - Database repository
    - Object factory (if it is stateless)
- Used in cases such as:
    - Construction call is expensive
    - Want to prevent anyone creating additional copies
    - Need to take car of lazy instantiation
- **Singleton is a component which is instantiated only once**
- For lazy instantiation and thread-safety use `sync.Once`
- Adhere to DIP: depend on interfaces, not concrete types

## Problems with Singleton

- Often breaks the dependency inversion principle
