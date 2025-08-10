# AviUtlExEdit2PluginManager

AviUtlExEdit2PluginManager is a plugin management system for AviUtl and ExEdit2, popular video editing software primarily used in Japan. This repository is currently in early development stages.

**Always reference these instructions first and fallback to search or bash commands only when you encounter unexpected information that does not match the info here.**

## Current Repository State

This repository is currently minimal, containing only:
- README.md (basic project title)
- LICENSE (MIT License)
- .github/copilot-instructions.md (this file)

**IMPORTANT**: There is no buildable code yet. Do not attempt to build, compile, or run application code as none exists currently.

## Working Effectively

### Initial Repository Setup
- Clone the repository: `git clone https://github.com/seless-yuu/AviUtlExEdit2PluginManager.git`
- Navigate to repository: `cd AviUtlExEdit2PluginManager`
- Check current status: `git --no-pager status`
- View repository structure: `ls -la`

### Repository Navigation
- **Repository root**: `/` - Contains README.md, LICENSE, and .github directory
- **Documentation**: `.github/` - Contains this instruction file
- **Current file structure**:
  ```
  .
  ├── README.md
  ├── LICENSE
  └── .github/
      └── copilot-instructions.md
  ```

### Git Operations (Validated Commands)
- Check repository status: `git --no-pager status`
- View commit history: `git --no-pager log --oneline -10`
- Check branches: `git --no-pager branch -a`
- View changes: `git --no-pager diff`
- Find files by type: `find . -name "*.md" -o -name "*.json" -o -name "*.yml" -o -name "*.yaml"`

## Development Environment

### Platform Considerations
- **Target Platform**: Windows (AviUtl is Windows-only software)
- **Development Environment**: Windows recommended for AviUtl plugin development
- **Language**: Likely C/C++ for AviUtl plugins (based on AviUtl SDK requirements)

### Future Development Setup (When Code is Added)
When this repository contains actual code, expect to need:
- Visual Studio or compatible C/C++ compiler for Windows
- AviUtl SDK for plugin development
- Batch files for build automation (referenced in Issue #1)

## Known Issues and Limitations

### Current Limitations
- **No build system**: Repository contains no buildable code
- **No tests**: No test framework or tests exist
- **No CI/CD**: No GitHub Actions workflows configured
- **No dependencies**: No package.json, requirements.txt, or similar

### Validation Commands That Work
- Repository navigation: `pwd`, `ls -la`
- Git operations: `git --no-pager status`, `git --no-pager log`
- File searching: `find . -type f -name "pattern"`

### Commands That Will Fail
- **Build commands**: No Makefile, CMakeLists.txt, or build scripts exist
- **Test commands**: No test framework configured
- **Package managers**: No npm, pip, or other package management files

## Future Development Guidelines

### When Adding Code
- Create appropriate build scripts (batch files for Windows development)
- Add README.md content describing the project purpose and usage
- Set up proper directory structure for AviUtl plugin development
- Add build validation and testing procedures

### Expected Build Times (Future)
When build system is implemented:
- **Initial setup**: Unknown - NEVER CANCEL setup operations
- **Compilation**: Unknown - NEVER CANCEL build operations
- **Testing**: Unknown - NEVER CANCEL test operations
- **Set timeouts to 60+ minutes** for any build operations when they are added

## Validation Scenarios (Current)

### Repository Structure Validation
```bash
# Verify repository contents
ls -la
# Expected: README.md, LICENSE, .github directory

# Check .github directory
ls -la .github/
# Expected: copilot-instructions.md

# Verify git repository
git --no-pager status
# Expected: Clean working tree or staged changes
```

### Navigation Validation
```bash
# Test basic navigation
pwd
# Expected: Full path to repository root

# Test file finding
find . -name "*.md"
# Expected: ./README.md and ./.github/copilot-instructions.md
```

## Issue Context

### Open Issues
- **Issue #1**: Create utility scripts for development builds (batch files)
- **Issue #3**: Set up Copilot instructions (this file addresses this issue)

### Project Context
- **AviUtl**: Popular video editing software in Japan
- **Plugin System**: AviUtl supports plugins for extended functionality
- **Target Audience**: Japanese developers and AviUtl users
- **Development Language**: Comments in Japanese acceptable

## Common Tasks Reference

### Repository Root Contents
```bash
ls -la
total 24
drwxr-xr-x 4 runner docker 4096 Aug 10 10:01 .
drwxr-xr-x 3 runner docker 4096 Aug 10 09:58 ..
drwxr-xr-x 7 runner docker 4096 Aug 10 10:00 .git
drwxr-xr-x 2 runner docker 4096 Aug 10 10:01 .github
-rw-r--r-- 1 runner docker 1067 Aug 10 09:58 LICENSE
-rw-r--r-- 1 runner docker   28 Aug 10 09:58 README.md
```

### README.md Content
```bash
cat README.md
# AviUtlExEdit2PluginManager
```

### Git Repository Status
```bash
git --no-pager log --oneline -5
9765e43 (HEAD -> copilot/fix-3, origin/copilot/fix-3) Initial plan
7f578c3 (grafted) Initial commit
```

## Critical Reminders

- **NEVER** attempt to build non-existent code
- **ALWAYS** check repository state before assuming build capabilities exist
- **USE** these validated commands for repository exploration
- **WAIT** for actual code to be added before attempting compilation
- **REFERENCE** these instructions first before trying alternative approaches
- **DOCUMENT** any new build processes when they are eventually added