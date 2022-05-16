## Virtualization basics

### history of virtualization
- ...
### what is virtualization
- virtualization provide abstraction on top of the actual resources we want to virtualize. The level at which this abstraction is applied changes the way that different virtualization techniques look.
- at a higher level, there aer two major virtualization techniques based on the level of abstraction: virtual machine based, container based.

### VM based virtualization
- the vm based approach virtualize the complete OS. the abstraction it presents to the VM are virtual devices like virtual disks, virtual CPUs, and Virtual NICs. In other words, we can state that this is virtualizing the complete ISA (instruction set architecture); as a example, the x86 ISA.
- VMs are very similar to other processes in the host OS. VMs execute in a hardware-isolated virtual address space and at a lower privilege level than the host OS. The primary difference between a process and a VM is the ABI(application binary interface) exposed by the host to VM.

### container based virtualization
- this form of virtualization doesn't abstract the hardware but uses techniques within the Linux kernel to isolate access paths for different resources. it carves out a logical boundary within the same operating system.

#### hypervisors
- a special piece of software is used to virtualize the OS, called the hypervisor. The hypervisor itself has two parts:
1. Virtual Machine Monitor(VMM): used for trapping and emulating the privileged instruction set (which can only the kernel of the operating system can perform)
2. Device model: Used for virtualizing the I/O devices.

##### Virtual Machine Monitor (VMM)
- since the hardware is not available directly on virtual machine (although in some cases it can be), the VMM traps privileged instructions that access the hardware like (disk, network card) and executes these instructions on behalf of the virtual machine
- the VMM has to satisfy three properties:
1. isolation: should isolate guest(VMs) from each other.
2. equivalency: should behave the same, with or without virtualization. This means we run the majority of the instructions on the physical hardware without any translation, and so on...
3. performance: should perform as good as it does without any virtualization

- some of the common functionalities of the VMM:
1. does not allow the VM to access privileged state; that is, things like manipulating the state of certain host registers should not be allowed from the VM.
2. handles exception and interrupts. If a network call was issued from within a virtual machine, it will be trapped in the VMM and emulated. On receipt of a response over the physical network/NIC, the CPU will generate an interrupt and deliver it to the actual virtual machine it's addressed to
3. handles CPU virtualization by running the majority of the instructions natively( within the virtual CPU of the VM) and only trapping for certain privileged instruction. This means the performance is almost as good as code running directly on the hardware.
4. handles memory mapped I/O by mapping the calls to the virtual device-mapped memory in the guest to the actual physical device-mapped memory.

##### Device Model
- the device model of the hypervisor handles the I/o virtualization again by trapping and emulating and then delivering interrupts back to the specific VM.

#### Memory Virtualization

#### CPU virtualization

#### IO Virtualization