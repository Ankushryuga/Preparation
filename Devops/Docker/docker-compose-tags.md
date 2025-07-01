#  1.Top-Level Tags:

| Tag        | Description                                                                                     |
| ---------- | ----------------------------------------------------------------------------------------------- |
| `version`  | Specifies the Compose file format version. (e.g., `"3.9"`, `"3.8"`). *Optional in Compose v2+*. |
| `services` | Defines all the containers (services) that will be run.                                         |
| `volumes`  | Declares named volumes that can be reused across services.                                      |
| `networks` | Defines custom networks for containers to communicate.                                          |
| `configs`  | Defines configuration files that can be mounted inside containers. (Swarm only)                 |
| `secrets`  | Defines sensitive data (e.g., passwords) to be used with services. (Swarm only)                 |


# 2.Service Sections Tags:
# Each service can include many options:

| Tag                    | Description                                                      |
| ---------------------- | ---------------------------------------------------------------- |
| `image`                | Image to use (e.g., `nginx:latest`).                             |
| `build`                | Build context and Dockerfile to build the image locally.         |
| `command`              | Override default command for the container.                      |
| `container_name`       | Assign a custom name to the container.                           |
| `depends_on`           | Specify service startup order (not a wait mechanism).            |
| `ports`                | Map container ports to host ports (e.g., `"8080:80"`).           |
| `volumes`              | Mount host paths or named volumes.                               |
| `environment`          | Set environment variables (e.g., `VAR=value`).                   |
| `env_file`             | Load environment variables from a `.env`-like file.              |
| `networks`             | Assign the service to one or more networks.                      |
| `restart`              | Restart policy (`no`, `always`, `on-failure`, `unless-stopped`). |
| `stdin_open`           | Keep STDIN open (useful for interactive containers).             |
| `tty`                  | Allocate a TTY (useful for terminal-based apps).                 |
| `entrypoint`           | Override the default entrypoint.                                 |
| `working_dir`          | Set the working directory inside the container.                  |
| `healthcheck`          | Define how Docker checks container health.                       |
| `labels`               | Add metadata labels.                                             |
| `logging`              | Configure logging drivers and options.                           |
| `extra_hosts`          | Add host-to-IP mappings.                                         |
| `hostname`             | Set the container’s hostname.                                    |
| `dns` / `dns_search`   | Customize DNS settings for the container.                        |
| `privileged`           | Give extended privileges to the container.                       |
| `cap_add` / `cap_drop` | Add or drop Linux capabilities.                                  |
| `secrets`              | Mount secrets into the container (Swarm only).                   |
| `configs`              | Mount configuration files (Swarm only).                          |
| `user`                 | Set user (UID or username) to run as inside container.           |
| `deploy`               | Swarm-only deployment configs (replicas, resources, etc.).       |
| `init`                 | Use an init process inside the container.                        |
| `isolation`            | Container isolation technology (Windows-only).                   |
| `platform`             | Target platform (e.g., `linux/amd64`, `linux/arm64`).            |
| `shm_size`             | Set shared memory size (e.g., for PostgreSQL).                   |


# 3. Build Suboptions:
# When using build, you can specify:
| Tag          | Description                                         |
| ------------ | --------------------------------------------------- |
| `context`    | Path to the build context (folder with Dockerfile). |
| `dockerfile` | Custom Dockerfile name or path.                     |
| `args`       | Build-time variables.                               |
| `cache_from` | Images to use for cache.                            |
| `labels`     | Metadata labels to add to the image.                |
| `target`     | Specify build stage in multi-stage builds.          |
| `secrets`    | Secret values used during build.                    |


# 4.Volume Section:
| Tag           | Description                          |
| ------------- | ------------------------------------ |
| `driver`      | Specify volume driver.               |
| `driver_opts` | Options passed to the volume driver. |
| `external`    | Use an existing volume.              |
| `labels`      | Metadata labels for the volume.      |
| `name`        | Custom name for the volume.          |


# 5.Networks Section:
| Tag           | Description                                        |
| ------------- | -------------------------------------------------- |
| `driver`      | Network driver to use (e.g., `bridge`, `overlay`). |
| `driver_opts` | Options passed to the driver.                      |
| `ipam`        | IP address management configuration.               |
| `external`    | Use an existing network.                           |
| `labels`      | Add metadata to networks.                          |
| `name`        | Custom network name.                               |


# 6.Deploy:
| Tag              | Description                                 |
| ---------------- | ------------------------------------------- |
| `mode`           | `replicated` or `global`.                   |
| `replicas`       | Number of replicas (only for `replicated`). |
| `resources`      | Set resource limits and reservations.       |
| `restart_policy` | Configure restart behavior on failure.      |
| `placement`      | Constraints for where containers run.       |
| `update_config`  | Rolling update strategy.                    |
| `labels`         | Deployment-level metadata.                  |
