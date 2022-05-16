# Container
- linux containers are made of three linux kernel primitives: linux namespaces, cgroups, layered file systems.

## namespaces - restricting the visibility 

- namespaces allow the kernel to provide isolation by restricting the visibility of the kernel resource like mount point, network subsystems among process scoped to different namespaces.

- a namespaces is a logical isolation within the Linux kernel. A namespaces controls  visibility within the kernel

- since there could be many user space applications running in parallel on a single linux kernel, we need a way to provide isolation between these use space-based applications.


### namespace types:
1. UTS: this namespace allows a process to see a separate hostname other than the actual global namespace one
2. PID: the processes within the PID namespace have a different process tree.
3. Mount: It controls which mount points a process should see.
4. Network: a network namespace gives a container a separate set of network subsystems. this means that the process within the network namespace will see different network interfaces, routes, and iptables.
5. IPC: this namespace scopes IPC constructs such as POSIX message queues. Between two process within the same namespace, IPC is enabled, but it will be restricted if two processes in two different namespaces try to communicate over IPC.
6. Cgroup: this namespaces restrict the visibility of the cgroup file system to the cgroup the process belong to. Without this restriction, a process could peek at the global cgroups via the /proc/self/cgroup hierarchy. This namespace effectively virtualize the cgroup itself.
7. time: change the date and time inside a container and adjusts the clocks for a container restored from a checkpoint.

## cgroups - resource control
- cgroup (abbreviate from control groups) is a Linux kernel feature that limits, account for, and isolates the resource usage (CPU, memory, disk I/O, network, etc...) of a collection of processes.
- there are different types of cgroups, based on which resources we want to control.
- CPU group