# Write the README.md content to a downloadable file at /mnt/data/README.md

content = """# Client Registration Guide — Custom OAuth2 / OIDC Server

This document lists **all practical ways** to register OAuth2/OIDC clients for your **Custom Authorization Server** (Spring Authorization Server–based), with copy-pasteable examples.

> Replace host/IPs as needed. Examples assume:
>
> ```
> ISSUER     = http://10.201.240.154:8080
> DISCOVERY  = http://10.201.240.154:8080/.well-known/openid-configuration
> DCR        = http://10.201.240.154:8080/connect/register
> ```
>
> If your server requires a bootstrap token for DCR:
> ```
> export AUTH_DCR_INITIAL_TOKEN='<your-initial-token>'
> ```

---

## Contents

- [Prerequisites](#prerequisites)
- [What to Register (quick reference)](#what-to-register-quick-reference)
- [A) Dynamic Client Registration (DCR)](#a-dynamic-client-registration-dcr)
  - [SPA (Public + PKCE)](#spa-public--pkce)
  - [Backend (Confidential, Client Credentials)](#backend-confidential-client-credentials)
  - [Native App (Loopback)](#native-app-loopback)
  - [Multiple Redirect URIs](#multiple-redirect-uris)
  - [Update / Delete (if implemented)](#update--delete-if-implemented)
- [B) Custom Admin API (`/admin/clients`)](#b-custom-admin-api-adminclients)
- [C) Programmatic Seeding (Java @Startup)](#c-programmatic-seeding-java-startup)
- [D) SQL Seeding (Flyway/Liquibase)](#d-sql-seeding-flywayliquibase)
- [E) Postman / GUI](#e-postman--gui)
- [F) Shell Script / Makefile](#f-shell-script--makefile)
- [Troubleshooting](#troubleshooting)
- [Security Notes](#security-notes)
- [Client Metadata Field Reference](#client-metadata-field-reference)

---

## Prerequisites

- Auth server is running and reachable at `$ISSUER`.
- Verify discovery:
  ```bash
  curl -s $ISSUER/.well-known/openid-configuration | jq .
