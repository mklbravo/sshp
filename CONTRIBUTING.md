# Contributing to SSHP

We appreciate your interest in improving SSHP! Here’s how you can help:

## How to Contribute
- **Issues:** If you spot bugs, need help, or want to suggest features, please open a GitHub issue with as much detail as possible.
- **Pull Requests:**
  1. Fork this repository and create a new branch for your feature or fix.
  2. Make clear, incremental commits with good descriptions.
  3. Ensure your code passes existing tests and adheres to project conventions (see below).
  4. Open a pull request, referencing any related issues or discussions.
  5. Describe your change clearly in the PR description.

## Project Structure
- `domain/`: Core entities, value objects, and repository interfaces.
- `application/`: Use cases/business logic, service and port definitions.
- `infrastructure/`: Concrete adapter implementations (e.g., JSON, SSH, env, persistence).
- `tui/`: Terminal UI components, models, and update logic.
- `internal/config/`: Configuration loading for different environments.

## Code Style & Guidelines
- Write clear, idiomatic Go code.
- Use dependency injection and strive for maintainable, testable modules.
- Favor modular, clean code and avoid large, monolithic changes.
- Add/update comments and documentation where helpful.

## Before You Submit
- Run tests locally (if applicable).
- Check for correct behavior in a real terminal (the TUI requires a TTY).
- Ensure your feature/fix integrates cleanly with the layered architecture.

## License
All contributions are covered by the project’s license; please see LICENSE file for details.

Thank you for helping make SSHP better!