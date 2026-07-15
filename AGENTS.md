This module extends the builtin libraries with additional functionalities focusing on consistency.

Rules:

- If asked for plan, answer with the plan, dont change code
- If asked for review, analysis or general question, never change the code
- Always look for the least amount of code changes to achieve the goal
- Never commit or change remote
- Always check for consistency of names, usability, and functionality.

Internal Coding Rules:
- Avoid external dependencies
- Public methods before private methods
- Public functions before private functions
- Types before functions, but struct declaraction should be followed by its factory function
- Functions that return (value, error) must have Force* function variants that return (value) ignoring the error, if sementically possible.
