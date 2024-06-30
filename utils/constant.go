package utils

// Define constants for the commit message
const (
	SYSTEM_PROMPT = `## Generate a High-Quality Git Commit Message

Please create a concise and informative commit message following best practices for Git commit messages.

**Context:**

The following is a diff of the changes made in this commit:`
	COMMIT_MSG_STRUCTURE = `

**Commit Message Structure:**

<type>[optional scope]: <short summary>
[optional body]
[optional footer(s)]

**Commit Types:**

* **feat:** A new feature is introduced.
* **fix:** A bug fix is implemented.
* **chore:** Non-functional changes (e.g., updating dependencies).
* **refactor:** Code refactoring without fixing bugs or adding features.
* **docs:** Documentation updates.
* **style:** Code style changes (e.g., formatting).
* **test:** Adding or updating tests.
* **perf:** Performance improvements.
* **ci:** Continuous integration changes.
* **build:** Build system or external dependencies changes.
* **revert:** Reverting a previous commit.

**Guidelines:**

1. **Subject Line:**
   - Use imperative mood (e.g., "Add", "Fix", "Update").
   - Keep it concise (maximum 50 characters or less).
   - Capitalize the first letter.
   - Do not end with a period.
   - **Include an optional scope after the type to specify the part of the codebase affected (e.g., "feat(auth):").** 

2. **Body:**
   - Use if more explanation is needed.
   - Wrap text at 72 characters.
   - Explain the "what" and "why" of the changes.

3. **Footer:**
   - Include author information as follows:
     Authored-by: Jiwon Choi <devjiwonchoi@gmail.com>

**Example:**

fix(auth): Resolve login failure issue
This commit fixes the login failure issue that was occurring due to incorrect password hashing. The hashing algorithm has been updated to bcrypt.

Authored-by: Jiwon Choi <devjiwonchoi@gmail.com>

**Generate a commit message following these guidelines.**`
)
