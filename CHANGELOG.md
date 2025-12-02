# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2024-12-02

### Added
- Initial release of AmoCRM Go Client
- OAuth 2.0 authentication with automatic token refresh
- Permanent token authentication (recommended for server integrations)
- File-based token storage
- Rate limiting (7 requests per second by default)
- Context support for all API calls
- Comprehensive error handling

#### Entities Support
- Contacts (create, read, update, batch operations)
- Companies (create, read, update, batch operations)
- Leads (create, read, update, batch operations, entity linking)
- Tasks (create, read, update, complete)
- Notes (create, read)
- Webhooks (subscribe, unsubscribe, list)
- Catalogs (list)
- Account information

#### Features
- Type-safe API interactions
- Automatic token refresh on expiration
- Debug logging support
- Custom HTTP client support
- Configurable rate limiting
- Custom token storage interface

#### Examples
- Basic usage with permanent token
- OAuth 2.0 authentication flow
- Batch operations
- Entity linking

#### Documentation
- Comprehensive README with usage examples
- Development guide
- MIT License

### Notes
- Based on AmoCRM API v4
- Inspired by [amocrm-api-php-v4](https://github.com/dedomorozoff/amocrm-api-php-v4)
- Requires Go 1.21 or higher

[1.0.0]: https://github.com/yourusername/amocrm-go/releases/tag/v1.0.0
