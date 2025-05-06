
# Table of Contents

1.  [Active Directory Basics](#orgb342461)
    1.  [Introduction to Active Directory](#org08a213b)
        1.  [**Definition**: Microsoft's directory service for Windows domain networks](#org86d69e4)
        2.  [**Purpose**: Centralized domain management service that authenticates and authorizes users and computers in a network](#org2c0943b)
    2.  [Key Components of Active Directory](#org18d62c2)
        1.  [**Domain Controllers**: Servers that store AD data and handle authentication](#orgfb1dbe6)
        2.  [**Domain**: Logical group of objects (users, computers, resources) sharing the same AD database](#orgb231bce)
        3.  [**Forest**: Collection of domains sharing a common schema and global catalog](#orgba172bc)
        4.  [**Organization Units** (OUs): Containers for organizing AD objects](#org0a7f2b9)
        5.  [**Group Policy Objects** (GPOs): Collections of settings that can be applied to users and computers](#orga0f1dc7)
    3.  [Common AD Administrative Tasks](#orga83ceb6)
        1.  [User account creation, modification, and deletion](#org8a29256)
        2.  [Group management and permission assigment](#org342404b)
        3.  [Password policy implementation](#org64f2c65)
        4.  [Computer acocunt management](#orgd80c6fb)
        5.  [Trust relationships configuration](#org5258d41)
    4.  [AD Management Tools](#org2c9399c)
        1.  [Active Directory Users and Computers (ADUC)](#org2980c23)
        2.  [Active Directory Administrative Center (ADAC)](#orgc526256)
        3.  [Active Directory Domains and Trusts](#org228da55)
        4.  [Active Directory Sites and Services](#org8ad8e90)
        5.  [LDAP utilities and PowerShell cmdlets](#org3466581)
2.  [BASH Scripting Basics](#org5f93d80)
    1.  [BASH is a type of Shell.](#org730b544)
        1.  [Other UNIX shells include: zsh (used on macOS), fish, dash, and more&#x2026;](#org7c62a29)
    2.  [Basic BASH Commands](#org9f4cdd4)
    3.  [shebang #!](#orgc4d00b9)
        1.  [#!/bin/bash](#orgd687bba)
    4.  [which](#orgc038fb2)
        1.  [**which bash** (/usr/bin/bash)](#orgb505fbb)
    5.  [Variables](#orgb6efb3b)
        1.  [$](#orgcbd501d)
        2.  [Naming conventions](#org7a36805)
    6.  [I/O](#org5827ef7)
        1.  [Can read user input using the **read** command](#org00dedcb)
    7.  [Conditional statements (if/else)](#orgaa2b25d)
    8.  [Looping](#org5b2b29c)
        1.  [**while** loop](#orge08a8ef)
        2.  [**for** loop](#org376b7e3)
    9.  [Case statements](#org9c30055)
        1.  [Used to compare a given value agianst a list of patterns and execute a block of code based on the first pattern that matches.](#org259fded)
    10. [Example of BASH Scripts I Use:](#org4ecd667)
        1.  [KERNEL.SH](#orged44b48)
        2.  [~/.zprofile](#org2e5cbea)
3.  [Group Policy Basics](#org84a6eb2)
    1.  [Introduction to Group Policy](#orgd6a9615)
        1.  [**Group Policy** is a feature of Windows that provides centralized management and configuraiton of operating systems, applicaitons, and user settings.](#orgce0598f)
    2.  [Main Goals](#orge921f2a)
        1.  [Ensure consistent security policies and user experience.](#orge26c958)
        2.  [Reduce administrative overhead through automation.](#org1d4edb1)
        3.  [Apply settings at scale, such as disabling USB ports or configuring homepages.](#org0dfba45)
    3.  [Group Policy Components](#orgad3d659)
        1.  [**Group Policy Objects (GPOs)**](#org92d6e04)
        2.  [**Group Policy Container (GPC)**](#org756eaf7)
        3.  [**Group Policy Template (GPT)**](#org9ad34af)
    4.  [Policy Settings Categories](#orgc9c5bca)
        1.  [**Computer Configuration**](#orgae2cfd1)
        2.  [**User Configuration**](#org4eb58e0)
    5.  [Common Group Policy Settings](#org343e3f8)
        1.  [Security Settings (password policies, account lockout, etc.)](#org6f43b78)
        2.  [Software Installation and Updates](#org24cef5b)
        3.  [Script Deployment (startup, shutdown, logon, logoff)](#org2f23eee)
        4.  [Folder Redirection and Offline Files](#org83dc6f5)
        5.  [Administrative Templates (registry-based policy settings)](#org67d90fd)
    6.  [Group Policy Management](#org1754595)
        1.  [Group Policy Managment Console (GPMC)](#org41d8804)
        2.  [GPO creation and linking](#org5fc359c)
        3.  [Security filtering and WMI filtering](#orgade0d28)
        4.  [Group Policy inheritance and precedence](#org93b221f)
        5.  [Group Policy troubleshooting tools](#orge0cdd88)
4.  [Intermediate Python for Cybersecurity Professionals](#org695011f)
    1.  [Python for Security Operations](#org14bd073)
        1.  [Log parsing and analysis](#org5f51132)
        2.  [Network scanning and reconnaissance](#orge7a99af)
        3.  [Automated security assessments](#org29b4e2c)
        4.  [Incident response automation](#org33f1b39)
    2.  [Essential Libraries for Security](#orgdc794e3)
        1.  [**Requests**: HTTP library for API interactions](#org9572c01)
        2.  [**Scapy**: Packet manipulation](#org2a4f5c8)
        3.  [**Paramiko**: SSH implemenation](#org69685fd)
        4.  [**Beautiful Soup**: Web scraping and parsing](#orga70ea4d)
        5.  [**Pandas**: Data analysis and manipulation](#org5ae9baf)
    3.  [Security Script Examples](#orgd024474)
        1.  [Simple port scanner](#orge54f3b1)
        2.  [**Usage**: port<sub>scan</sub>("192.168.1.1", [22, 80, 443])](#org744ca55)
    4.  [Automated Security Testing](#orgf95efe1)
        1.  [Writing scripts for vulnerability checks](#org8ae8e4e)
        2.  [Automating repetitive security tasks](#org0140ffb)
        3.  [Developing custom security tools](#org1d2c113)
5.  [Linux CLI Basics](#org2b3e54d)
    1.  [Getting Started with the Command Line](#orgf3437b9)
        1.  [Terminal emulator and shell options](#org5678592)
        2.  [Command syntax and structure](#org2141549)
        3.  [Getting help (**man**, **info**, **&#x2013;help**)](#orgc43e2fd)
    2.  [Essential Commands for System Administrators](#orgc604b43)
        1.  [**System Information**: **uname**, **hostname**, **lsb<sub>release</sub>**, **fastfetch** & **neofetch** (for aesthetics)](#org3fb4aa0)
        2.  [**Process Management**: **ps**, **top**, **htop**, **btop** (my personal favorite process manager), **kill**, **nice**](#orgf2c1cf7)
        3.  [**User Management**: **useradd**, **usermod**, **userdel**, **passwd**](#org52896bb)
        4.  [**Package Management**: **apt** (Ubuntu & Debian), **yum** (RHEL & Fedora), **dnf** (RHEL & Fedora), **pacman** (Arch Linux), **xbps** (Void Linux)](#org0843a26)
        5.  [**Service Management**: **systemctl** (systemd), **service**, **sv** (runit), **vsv** (runit), **ln -s *etc/sv/<service> /var/service*** (runit)](#org1982a70)
        6.  [**Networking**: **ip**, **netstat**, **ss**, **ping**, **traceroute**, **dig**, **nslookup**](#org58c13fc)
    3.  [Test Processing Commands](#org1241459)
        1.  [**grep**, **sed**, **awk**, **cut**, **sort**, **cat**, **uniq**, **wc**](#org466cc4a)
    4.  [Text Editors](#orgca2a549)
        1.  [**vi**, **vim**, **neovim** (my favorite terminal text editor)](#org9c2cd32)
        2.  [**GNU Emacs**, **Spacemacs** (my favorite GUI text editor), **Doom Emacs**](#orgd570c85)
        3.  [**GNU nano**, **micro**](#org2edd143)
    5.  [System Monitoring and Performance Analysis](#org45809af)
        1.  [**top**, **htop**, **btop**](#orgdddeae7)
6.  [Linux File System Basics](#org125a898)
    1.  [Linux File System Types](#org4db552c)
        1.  [**ext2** (Second Extended Filesystem)](#orgbfe65f4)
        2.  [**ext3** (Third Extended Filesystem)](#org4ad2a9f)
        3.  [**ext4** (Fourth Extended Filesystem)](#orgebbc11a)
        4.  [**XFS**](#orgdfe3e35)
        5.  [**Btrfs** (B-tree File System)](#org657549d)
        6.  [**F2FS** (Flash-Friendly File System)](#orged9cf36)
        7.  [**VFAT\*/\*FAT32\*/\*exFAT**](#orgb87c6dd)
        8.  [**NTFS** (New Technology File System)](#org90fbf7d)
        9.  [**ISO 9660\*/\*UDF**](#org5209148)
    2.  [File System Hierarchy Standard (FHS)](#org9272a54)
        1.  [Key directories:](#orgb12439d)
    3.  [File System Management](#orgddd87c9)
        1.  [**Mounting file systems**: mount, umount](#orgb1561e5)
        2.  [**File system tools**: df, du, fsck](#org7885bf5)
        3.  [**Disk partitioning**: fdisk, cfdisk (my favorite terminal based partitioning tool), parted, gparted (my favorited GUI partitioning tool)](#org1b7f3e7)
        4.  [**Logical Volume Managment (LVM)** basics](#org2b56fda)
    4.  [File Types in Linux](#orgf4f5ad1)
        1.  [**Regular files (-)**](#orgb4ba406)
        2.  [**Directories (d)**](#orga0681fa)
        3.  [**Symbolic links (l)**](#org1322f2e)
        4.  [**Device files (b, c)**](#org52240ec)
        5.  [**Named pipes (p) and sockets (s)**](#orgdc4ec00)
    5.  [File Permissions and Ownership](#org7d67689)
        1.  [**Permission Basics**](#org0c5784e)
        2.  [**Permission Types**](#org3894467)
        3.  [**Permission notation**](#org45708c7)
        4.  [**chmod**, **chown**, and **chgrp** commands](#org9f2e368)
        5.  [**Special permissions**: SUID, SGID, and sticky bit](#orgac5b421)
7.  [PowerShell for Security Professionals](#org25614d7)
    1.  [PowerShell Security Features](#org742ef98)
        1.  [Execution policies](#org166e407)
        2.  [Constrained language mode](#org7472756)
        3.  [Script Block Logging](#orgbade34b)
        4.  [System-wide transcription](#orgbff34f6)
        5.  [Anti-malware Scan Interface (AMSI)](#orgd181bdb)
    2.  [Security Cmdlets and Modules](#org5288798)
        1.  [**Get-Acl** and **Set-Acl**: Managing access control lists](#org75125c2)
        2.  [**ConvertTo-SecureString** and **ConvertFrom-SecureString**: Working with secure strings](#org8d9c8df)
        3.  [**Get-Credential**: Securely requesting credentials](#org958cb35)
        4.  [**Set-ExecutionPolicy**: Managing PowerShell execution policies](#orgd997f1f)
        5.  [Security-related modules: **SecurityFever**, **PSAD**, **PSWinReporting**](#orgd8817e9)
    3.  [Security Monitoring and Auditing](#org59011ae)
        1.  [Querying event logs for security events](#org7b024e7)
        2.  [Monitoring file and registry changes](#org609ae04)
        3.  [Auditing system configuration](#org40d09c9)
        4.  [Detecting unauthorized PowerShell usage](#orgebabdd1)
    4.  [Common Security Tasks in PowerShell](#org0bcf4bb)
8.  [PowerShell Basics](#org44c94d1)
    1.  [Introduction to PowerShell](#orgb226eaf)
        1.  [PowerShell vs. Command Prompt](#orgba6529f)
        2.  [PowerShell versions and features](#org57c5aa1)
        3.  [PowerShell ISE and PowerShell 7](#orgb405fec)
    2.  [PowerShell Syntax and Structure](#org2686616)
        1.  [**Cmdlets**: verb-noun naming convention](#orgb4066ae)
        2.  [Parameters and parameter sets](#org83ca210)
        3.  [Pipelines and object handling](#orgb491922)
        4.  [Variables and data types](#org0c687f7)
        5.  [Comparison operators](#orga57755e)
        6.  [Flow control (if, switch, loops)](#orgddbfc12)
    3.  [Essential PowerShell Cmdlets](#org5560263)
        1.  [**Get-Command**: Discover available commands](#org74a804b)
        2.  [**Get-Help**: Access documentation and examples](#org6770e86)
        3.  [**Get-Process**: List running processes](#orgc50125d)
        4.  [**Get-Service**: List system services](#org6fb474b)
        5.  [**Get-Item** and **Get-ChildItem**: Work with files and directories](#org8fa9368)
        6.  [**Import-Module**: Load additional functionality](#org75c608b)
        7.  [**Invoke-Command**: Run commands on remote computers](#org646c8e7)
    4.  [PowerShell Scripting](#orga2d0495)
        1.  [Script creation and execution](#org0e47c08)
        2.  [Functions and modules](#org95b426d)
        3.  [Error handling and debugging](#orgcce519a)
        4.  [Script signing and security](#orgaf50d43)
        5.  [PowerShell profiles](#org2231056)
9.  [Python Basics](#org2c8a2f1)
    1.  [Python Environment Setup](#org675ac72)
        1.  [Installing Python](#org566b360)
        2.  [Package management with **pip**](#orgc6df576)
        3.  [Virtual environments (venv, conda)](#org1b18588)
        4.  [Code editors and IDEs (VS Code, PyCharm, NeoVim, Emacs)](#org9c4eefd)
    2.  [Python Syntax Fundamentals](#org000709c)
        1.  [Variables and data types](#orgc4b33f9)
        2.  [Operators and expressions](#orgc13cf76)
        3.  [Control flow: if statements, loops](#orgf4d1984)
        4.  [Functions and modules](#org71ea77d)
        5.  [Error handling with try/except](#orgf19b901)
    3.  [Data Structures](#orgf4a2e08)
        1.  [Lists, tuples, and sets](#orgadbf585)
        2.  [Dictionaries](#org1a8fb2d)
        3.  [List comprehensions](#orgae71f5b)
        4.  [String manipulation](#orgd46ea3f)
    4.  [File Operations](#orgc087b6b)
        1.  [Reading and writing text files](#org596f075)
        2.  [Working with CSV and JSON](#org5ae0d2f)
        3.  [File path manipulation](#org1411b7f)
    5.  [Python for System Administration](#org5124318)
10. [Windows File System Basics](#org4cc097f)
    1.  [NTFS File System](#org30800d8)
        1.  [Features and benefits of NTFS](#org2481e7f)
        2.  [NTFS permissions model](#org80d36d1)
        3.  [Access Control Lists (ACLs)](#org54de680)
        4.  [NTDS special permissions](#orgd1ca2d8)
        5.  [Alternate Data Streams](#orgb8c133d)
    2.  [NTFS Folder Structure](#org6e9270c)
        1.  [Windows system folders](#orga8ab58b)
        2.  [Program Files structure](#orgbe18b50)
        3.  [User profiles and directories](#org5168259)
        4.  [System32 and SysWOW64](#orgffd3128)
        5.  [Windows Registry as a hierarchical database](#org3b024f1)
    3.  [File and Folder Management](#orgcd56d3f)
        1.  [File Explorer vs. Command Line tools](#org87ef9b6)
        2.  [PowerShell file management cmdlets](#org8213290)
        3.  [Advanced file operations](#org791e276)
        4.  [File system auditing](#orge190e15)
    4.  [Security Features](#org209b831)
        1.  [BitLocker drive encryption](#org53cb4fd)
        2.  [EFS (Encrypting File System)](#org91adcb4)
        3.  [File ownership and inheritance](#orge239088)
        4.  [File system quotas](#org1a5c4a1)
        5.  [Shadow copies and previous versions](#org7af04de)
    5.  [Common File System Management Tasks](#orgee3ad6e)
        1.  [Sharing folders and setting permissions](#org955894f)
        2.  [Using administrative shares](#orgfac33b3)
        3.  [Managing disk quotas](#orgd881c53)
        4.  [Configuring file screening](#org2e2cd8e)
        5.  [Implementing NTFS compression](#org68d60ce)
11. [Documentation Basics](#org32f03f9)
    1.  [Importance of Documentation for System Administrators](#orgfefb487)
        1.  [**Knowledge retention**: Preserving institutional knowledge](#org98dace6)
        2.  [**Troubleshooting efficiency**: Recording solutions for future reference](#orgcb847fc)
        3.  [**Onboarding**: Helping new team members get up to speed](#org0a75e73)
        4.  [**Disaster recovery**: Documenting cirtical systems and procedures](#org1eaac42)
        5.  [**Compliance**: Meeting regulatory and audit requirements](#orgd95d8e4)
    2.  [Effective Documentation Practices](#orgdc4380e)
        1.  [Keep documentation current and regularly updated](#orgeaf969d)
        2.  [Use conistent formatting and templates](#org6c383a8)
        3.  [Include dates, versions, and author information](#org0bc0644)
        4.  [Document assumptions and dependencies](#orgdb52bdf)
        5.  [Balance detail with readability](#org3bb670a)
        6.  [Link related documentation where appropriate](#org835c8e8)
        7.  [Use diagrams and screenshots for visual clarity](#org9cfec27)
        8.  [Include examples for complex procedures](#orgf5e3f44)
    3.  [Documentation Tools and Formats](#org78e1f78)
        1.  [**Plain Text Editors**:](#org396a440)
        2.  [**Markup Languages**:](#org8c57dae)
        3.  [**Wiki Systems**:](#orgd9ccdbf)
        4.  [**Specialized Documentation Tools**:](#org7e94cc3)
    4.  [Documentation Workflows for SysAdmins](#org66a36da)
        1.  [**Incident documentation**: Recording incidencts, troubleshooting steps, and resolutions](#org8716465)
        2.  [**Change management**: Documenting system changes, approvals, and rollback procedures](#org96239ad)
        3.  [**Standard Operating Procedures (SOPs)**: Creating step-by-step instructions for routine tasks](#orge421050)
        4.  [**Architecture documentation**: Recording system design and dependencies](#orgdfb4a9b)
        5.  [**Runbooks**: Creating operation guides for specific scenarios](#orgcdb9c9b)



<a id="orgb342461"></a>

# Active Directory Basics


<a id="org08a213b"></a>

## Introduction to Active Directory


<a id="org86d69e4"></a>

### **Definition**: Microsoft's directory service for Windows domain networks


<a id="org2c0943b"></a>

### **Purpose**: Centralized domain management service that authenticates and authorizes users and computers in a network


<a id="org18d62c2"></a>

## Key Components of Active Directory


<a id="orgfb1dbe6"></a>

### **Domain Controllers**: Servers that store AD data and handle authentication


<a id="orgb231bce"></a>

### **Domain**: Logical group of objects (users, computers, resources) sharing the same AD database


<a id="orgba172bc"></a>

### **Forest**: Collection of domains sharing a common schema and global catalog


<a id="org0a7f2b9"></a>

### **Organization Units** (OUs): Containers for organizing AD objects


<a id="orga0f1dc7"></a>

### **Group Policy Objects** (GPOs): Collections of settings that can be applied to users and computers


<a id="orga83ceb6"></a>

## Common AD Administrative Tasks


<a id="org8a29256"></a>

### User account creation, modification, and deletion


<a id="org342404b"></a>

### Group management and permission assigment


<a id="org64f2c65"></a>

### Password policy implementation


<a id="orgd80c6fb"></a>

### Computer acocunt management


<a id="org5258d41"></a>

### Trust relationships configuration


<a id="org2c9399c"></a>

## AD Management Tools


<a id="org2980c23"></a>

### Active Directory Users and Computers (ADUC)


<a id="orgc526256"></a>

### Active Directory Administrative Center (ADAC)


<a id="org228da55"></a>

### Active Directory Domains and Trusts


<a id="org8ad8e90"></a>

### Active Directory Sites and Services


<a id="org3466581"></a>

### LDAP utilities and PowerShell cmdlets


<a id="org5f93d80"></a>

# BASH Scripting Basics


<a id="org730b544"></a>

## BASH is a type of Shell.


<a id="org7c62a29"></a>

### Other UNIX shells include: zsh (used on macOS), fish, dash, and more&#x2026;


<a id="org9f4cdd4"></a>

## Basic BASH Commands

-   **cd**: Change directory
-   **ls**: List contents of the current directory
-   **mkdir**: Create a new directory
-   **touch**: Create a new file
-   **rm**: Remove a file or directory
-   **cp**: Copy a file or directory
-   **mv**: Move or rename a file or directory
-   **echo**: Print text to the terminal
-   **cat**: Concatenate a print the contents of a file
-   **grep**: Search for a pattern in a file
-   **chmod**: Change the permissions of a file or directory
-   **sudo**: Run a command with administrative privileges.
-   **df**: Display the amount of disk space available.
-   **history**: Show a list of previously executed commands.
-   **ps**: Display informaito about running processes.
-   **date**: Displays the currnet date
-   **pwd**: Displays the present working directory


<a id="orgc4d00b9"></a>

## shebang #!

-   Bash scripts start with a shebang


<a id="orgd687bba"></a>

### #!/bin/bash


<a id="orgc038fb2"></a>

## which

-   Can find your shell path:


<a id="orgb505fbb"></a>

### **which bash** (/usr/bin/bash)


<a id="orgb6efb3b"></a>

## Variables


<a id="orgcbd501d"></a>

### $

1.  $ is required to access an existing variable's value.


<a id="org7a36805"></a>

### Naming conventions

1.  Should start with a letter or an underscore (\_)

2.  Can contian letters, numbers, and underscores (\_)

3.  Are case-sensitive

4.  Should not contian spaces or special characters

5.  Use descriptive names that reflect the purpose of the variable

6.  Avoid using reserved keywords such as **if**, **then**, **else**, **fi**


<a id="org5827ef7"></a>

## I/O


<a id="org00dedcb"></a>

### Can read user input using the **read** command

    #!/bin/bash
    echo "What's your name?"
    
    read entered_name
    
    echo -e "\nWelcome to bash tutorial" $entered_name


<a id="orgaa2b25d"></a>

## Conditional statements (if/else)

    if [[ condition ]];
    then
        statement
    elif [[ condition ]]; then
        statement
    else
        do this by default
    fi


<a id="org5b2b29c"></a>

## Looping


<a id="orge08a8ef"></a>

### **while** loop

    #!/bin/bash
    i=1
    while [[ $i le 10 ]] ; do
        echo "$i"
        (( i += 1 ))
    done


<a id="org376b7e3"></a>

### **for** loop

    #!/bin/bash
    for i in {1..5}
    do
        echo $i
    done


<a id="org9c30055"></a>

## Case statements


<a id="org259fded"></a>

### Used to compare a given value agianst a list of patterns and execute a block of code based on the first pattern that matches.

    case expression in
        pattern1)
            # code to execute if expression matches pattern1
            ;;
        pattern2)
            # code to execute if expression matches pattern2
            ;;
        pattern3)
            # code to execute if expression matches pattern3
            ;;
        *)
            # code to execute if none of the above patterns match expression
            ;;
    esac


<a id="org4ecd667"></a>

## Example of BASH Scripts I Use:


<a id="orged44b48"></a>

### KERNEL.SH

Waybar module to show Linux kernel version and if it's up to date or out of date **(kernel.sh)**:

      #!/bin/bash
    
    CACHE_DIR="${XDG_CACHE_HOME:-$HOME/.cache}/waybar"
    CACHE_FILE="$CACHE_DIR/kernel_status"
    CACHE_TTL=$((6 * 60 * 60)) # 6 hours in seconds
    
    mkdir -p "$CACHE_DIR"
    
    KERNEL=$(uname -r)
    
    # Check if cache exists and is fresh and then check for available kernel update
    if [ -f "$CACHE_FILE" ] && [ $(( $(date +%s) - $(stat -c %Y "$CACHE_FILE") )) -lt $CACHE_TTL ]; then
        STATUS=$(cat "$CACHE_FILE")
    else
        if xbps-install -un | grep -q '^linux[[:space:]]'; then
            STATUS="Out of date"
        else
            STATUS="Up to date"
        fi
        echo "$STATUS" > "$CACHE_FILE"
    fi
    
    # Output JSON for Waybar
    printf '{"text":"%s", "tooltip":"%s"}\n' "$KERNEL" "$STATUS"


<a id="org2e5cbea"></a>

### ~/.zprofile

This is how I start my Wayland Compositor (Sway) on login **(~/.zprofile)**:

      if [ -z "$WAYLAND_DISPLAY" ] && [ "$(tty)" = "/dev/tty1" ]; then
        echo "sway Starting: $( date )"
        if test -z "${XDG_RUNTIME_DIR}"; then
            export XDG_RUNTIME_DIR=/tmp/${UID}-runtime-dir
            if ! test -d "${XDG_RUNTIME_DIR}"; then
                mkdir "${XDG_RUNTIME_DIR}"
                chmod 0700 "${XDG_RUNTIME_DIR}"
            fi
        fi
        dbus-run-session sway
    fi


<a id="org84a6eb2"></a>

# Group Policy Basics


<a id="orgd6a9615"></a>

## Introduction to Group Policy


<a id="orgce0598f"></a>

### **Group Policy** is a feature of Windows that provides centralized management and configuraiton of operating systems, applicaitons, and user settings.

1.  It works via **Active Directory (AD)**, allowing admins to enforce rules and restrictions across all computers and users in a domain.


<a id="orge921f2a"></a>

## Main Goals


<a id="orge26c958"></a>

### Ensure consistent security policies and user experience.


<a id="org1d4edb1"></a>

### Reduce administrative overhead through automation.


<a id="org0dfba45"></a>

### Apply settings at scale, such as disabling USB ports or configuring homepages.

1.  *Example use case*: Preventing non-admins from accessing Control Panel.


<a id="orgad3d659"></a>

## Group Policy Components


<a id="org92d6e04"></a>

### **Group Policy Objects (GPOs)**

1.  **GPOs** are containers that hold policy settings.

2.  Can be linked to **Sites**, **Domains**, or **Organizational Units (OUs)**.

3.  Each GPO contains both **Computer** and **User** configurations.

4.  GPOs are versioned and replicated via Active Directory.


<a id="org756eaf7"></a>

### **Group Policy Container (GPC)**

1.  The **GPC** is stored in Active Directory.

2.  Contains metadata about the GPO (GUID, version info, links).

3.  Acts as the "pointer" ot the file system-based GPT.


<a id="org9ad34af"></a>

### **Group Policy Template (GPT)**

1.  Stored in **SYSVOL** on domain controllers (**\\\domain\SYSVOL**).

2.  Contains the actual settings (scripts, ADM/ADMX templates, etc.)

3.  Synchronized between domain controllers using **DFS Replication**.


<a id="orgc9c5bca"></a>

## Policy Settings Categories


<a id="orgae2cfd1"></a>

### **Computer Configuration**

1.  Applies to **computer accounts**, regardless of who logs in.

2.  Runs during system startup.

3.  Settings include:

    1.  Windows Updates
    
    2.  Security configurations
    
    3.  Scripts (startup/shudown)
    
    4.  Services and device restrictions


<a id="org4eb58e0"></a>

### **User Configuration**

1.  Applies to **user accounts**, regardless of the machine.

2.  Runs during user login.

3.  Settings include:

    1.  Desktop environments
    
    2.  Folder redirection
    
    3.  Logon/logoff scripts
    
    4.  Start menu and taskbar configurations


<a id="org343e3f8"></a>

## Common Group Policy Settings


<a id="org6f43b78"></a>

### Security Settings (password policies, account lockout, etc.)


<a id="org24cef5b"></a>

### Software Installation and Updates


<a id="org2f23eee"></a>

### Script Deployment (startup, shutdown, logon, logoff)


<a id="org83dc6f5"></a>

### Folder Redirection and Offline Files


<a id="org67d90fd"></a>

### Administrative Templates (registry-based policy settings)


<a id="org1754595"></a>

## Group Policy Management


<a id="org41d8804"></a>

### Group Policy Managment Console (GPMC)


<a id="org5fc359c"></a>

### GPO creation and linking


<a id="orgade0d28"></a>

### Security filtering and WMI filtering


<a id="org93b221f"></a>

### Group Policy inheritance and precedence


<a id="orge0cdd88"></a>

### Group Policy troubleshooting tools


<a id="org695011f"></a>

# Intermediate Python for Cybersecurity Professionals


<a id="org14bd073"></a>

## Python for Security Operations


<a id="org5f51132"></a>

### Log parsing and analysis


<a id="orge7a99af"></a>

### Network scanning and reconnaissance


<a id="org29b4e2c"></a>

### Automated security assessments


<a id="org33f1b39"></a>

### Incident response automation


<a id="orgdc794e3"></a>

## Essential Libraries for Security


<a id="org9572c01"></a>

### **Requests**: HTTP library for API interactions


<a id="org2a4f5c8"></a>

### **Scapy**: Packet manipulation


<a id="org69685fd"></a>

### **Paramiko**: SSH implemenation


<a id="orga70ea4d"></a>

### **Beautiful Soup**: Web scraping and parsing


<a id="org5ae9baf"></a>

### **Pandas**: Data analysis and manipulation


<a id="orgd024474"></a>

## Security Script Examples


<a id="orge54f3b1"></a>

### Simple port scanner


<a id="org744ca55"></a>

### **Usage**: port<sub>scan</sub>("192.168.1.1", [22, 80, 443])

    import socket
    
    def port_scan(target, ports):
        print(f"Scanning {target} for open ports...")
        for port in ports:
            s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
            s.settimeout(1)
            result = s.connect_ex((target, port))
            if result == 0:
                print(f"Port {port}: Open")
            s.close()


<a id="orgf95efe1"></a>

## Automated Security Testing


<a id="org8ae8e4e"></a>

### Writing scripts for vulnerability checks


<a id="org0140ffb"></a>

### Automating repetitive security tasks


<a id="org1d2c113"></a>

### Developing custom security tools


<a id="org2b3e54d"></a>

# Linux CLI Basics


<a id="orgf3437b9"></a>

## Getting Started with the Command Line


<a id="org5678592"></a>

### Terminal emulator and shell options

1.  **Terminal Emulators**: Alacritty, Konsole, Windows Terminal, Foot, Kitty, WezTerm, VTerm, etc.

2.  **Shells**: sh, bash, PowerShell, zsh, fish, dash, cmd


<a id="org2141549"></a>

### Command syntax and structure


<a id="orgc43e2fd"></a>

### Getting help (**man**, **info**, **&#x2013;help**)


<a id="orgc604b43"></a>

## Essential Commands for System Administrators


<a id="org3fb4aa0"></a>

### **System Information**: **uname**, **hostname**, **lsb<sub>release</sub>**, **fastfetch** & **neofetch** (for aesthetics)


<a id="orgf2c1cf7"></a>

### **Process Management**: **ps**, **top**, **htop**, **btop** (my personal favorite process manager), **kill**, **nice**


<a id="org52896bb"></a>

### **User Management**: **useradd**, **usermod**, **userdel**, **passwd**


<a id="org0843a26"></a>

### **Package Management**: **apt** (Ubuntu & Debian), **yum** (RHEL & Fedora), **dnf** (RHEL & Fedora), **pacman** (Arch Linux), **xbps** (Void Linux)


<a id="org1982a70"></a>

### **Service Management**: **systemctl** (systemd), **service**, **sv** (runit), **vsv** (runit), **ln -s *etc/sv/<service> /var/service*** (runit)


<a id="org58c13fc"></a>

### **Networking**: **ip**, **netstat**, **ss**, **ping**, **traceroute**, **dig**, **nslookup**


<a id="org1241459"></a>

## Test Processing Commands


<a id="org466cc4a"></a>

### **grep**, **sed**, **awk**, **cut**, **sort**, **cat**, **uniq**, **wc**

1.  Text stream manipulation and filtering

2.  Regular expression (regex) basics


<a id="orgca2a549"></a>

## Text Editors


<a id="org9c2cd32"></a>

### **vi**, **vim**, **neovim** (my favorite terminal text editor)


<a id="orgd570c85"></a>

### **GNU Emacs**, **Spacemacs** (my favorite GUI text editor), **Doom Emacs**


<a id="org2edd143"></a>

### **GNU nano**, **micro**


<a id="org45809af"></a>

## System Monitoring and Performance Analysis


<a id="orgdddeae7"></a>

### **top**, **htop**, **btop**

1.  **Disk Usage**: df, du, ncdu

2.  **Memory usage**: free, vmstat

3.  **I/O Statistics**: iostat, iotop

4.  **Network Statistics**: netstat, iftop, iptraf


<a id="org125a898"></a>

# Linux File System Basics


<a id="org4db552c"></a>

## Linux File System Types


<a id="orgbfe65f4"></a>

### **ext2** (Second Extended Filesystem)

1.  Older file system, no journaling

2.  Suitable for USB drives or flash storage

3.  Simple and low overhead


<a id="org4ad2a9f"></a>

### **ext3** (Third Extended Filesystem)

1.  ext2 + **journaling** (tracks changes before committing)

2.  More reliable than ext2


<a id="orgebbc11a"></a>

### **ext4** (Fourth Extended Filesystem)

1.  Improved version of ext3

2.  **Default** in many Linux distros

3.  Supports larger files and volumes

4.  Backwards compatible with ext3 and ext2


<a id="orgdfe3e35"></a>

### **XFS**

1.  High-performance journaling file system

2.  Great for large files and parallel I/O

3.  Common in servers (e.g., CentOS, RHEL)


<a id="org657549d"></a>

### **Btrfs** (B-tree File System)

1.  Also known as *"better FS"* (the right way to call it !!)

2.  Advanced features: **snapshots**, **compression**, **checksums**

3.  Aimed to replace ext4 in the future

4.  Still maturing

    1.  I use btrfs on my PC and laptop, but usually go with ext4 on servers with older hardware


<a id="orged9cf36"></a>

### **F2FS** (Flash-Friendly File System)

1.  Optimized for NAND flash (e.g., SSDs, eMMC)

2.  Developed by Samsung


<a id="orgb87c6dd"></a>

### **VFAT\*/\*FAT32\*/\*exFAT**

1.  Compatible with Windows/macOS

2.  Good for USB drives

3.  No journaling or permission support


<a id="org90fbf7d"></a>

### **NTFS** (New Technology File System)

1.  Windows file system, read/write support on Linux via **ntfs-3g**

2.  Used for shared external drives

    1.  I use NTFS on my drive labeled "extra" which stores files for use between my Linux and Windows boots


<a id="org5209148"></a>

### **ISO 9660\*/\*UDF**

1.  Used for optical discs (CD/DVD)

2.  ISO 9660 is older; UDF supports larger files and newer systems


<a id="org9272a54"></a>

## File System Hierarchy Standard (FHS)

1.  *This is true for **most** Linux distributions*

    1.  A notable exception is **NixOS**


<a id="orgb12439d"></a>

### Key directories:

1.  **/**: Root directory

2.  **/etc**: System configuration files (default program configuration files go here)

3.  **/var**: Variable data (logs, spool files, temporary files)

4.  **/tmp**: Temporary files

5.  **/usr**: User programs, libraries, documentation

6.  **/home**: User home directories (*home/<user>*)

7.  **/boot**: Boot loader files

8.  **/mnt**: Mount point for external filesystems

9.  **/dev**: Device files

10. **/proc** and **/sys**: Virtual file systems for system information


<a id="orgddd87c9"></a>

## File System Management


<a id="orgb1561e5"></a>

### **Mounting file systems**: mount, umount


<a id="org7885bf5"></a>

### **File system tools**: df, du, fsck


<a id="org1b7f3e7"></a>

### **Disk partitioning**: fdisk, cfdisk (my favorite terminal based partitioning tool), parted, gparted (my favorited GUI partitioning tool)


<a id="org2b56fda"></a>

### **Logical Volume Managment (LVM)** basics


<a id="orgf4f5ad1"></a>

## File Types in Linux


<a id="orgb4ba406"></a>

### **Regular files (-)**

1.  Text files, binary files, images, etc.

2.  Viewed with commands like **cat**, **less**, **more**

3.  Identified with **ls l** showing a **-** as the first character


<a id="orga0681fa"></a>

### **Directories (d)**

1.  Containers for other files and directories

2.  Created with **mkdir** command

3.  Navigated with **cd** command

4.  Identified with **ls -l** showing a **d** as the first character


<a id="org1322f2e"></a>

### **Symbolic links (l)**

1.  Pointers to other files or directories

2.  Created with \*ln -s <target<sub>file</sub>> <link<sub>name</sub>>

3.  Can span across file systems

4.  Identified with **ls -l** showing an **l** as the first character

5.  Show the path they to to in **ls -l** output


<a id="org52240ec"></a>

### **Device files (b, c)**

1.  **Block devices (b)**: Accessed in chunks, like hard drives

    1.  Example: **/dev/sda** (first SCSI disk)

2.  **Character devices (c)**: Accessed one character at a time

    1.  Example: **/dev/tty1** (first terminal)

3.  Created with **mknod** command (usually handled by udev)

4.  Located primarility in **/dev** directory


<a id="orgdc4ec00"></a>

### **Named pipes (p) and sockets (s)**

1.  **Named pipes**: Allow processes to communicate

    1.  Created with **mkfifo** command
    
    2.  Used for inter-process communication

2.  **Unix domain sockets**: Used for local inter-process communication

    1.  Created by srever processes
    
    2.  Example: X11 socket (***tmp*.X11-unix/x0**)


<a id="org7d67689"></a>

## File Permissions and Ownership


<a id="org0c5784e"></a>

### **Permission Basics**

1.  **User (u)**: The owner of the file

2.  **Group (g)**: The group associated with the file

3.  **Others (o)**: Everyone else


<a id="org3894467"></a>

### **Permission Types**

1.  **Read (r, 4)**:

    1.  For files: View file contents
    
    2.  For directories: List directory contents

2.  **Write (w, 2)**:

    1.  For files: Modify file contents
    
    2.  For directories: Create, delete, or rename files within the directory

3.  **Execute (x, 1)**:

    1.  For files: Run the file as a program or script
    
    2.  For directories: Access files within the directory (traverse)


<a id="org45708c7"></a>

### **Permission notation**

1.  **Symbolic notation** (rwx)

    1.  Example: **rw-r&#x2013;r--** (User: read+write, Group: read, Others: read)
    
    2.  Position matters: **rwxrwxrwx** - user, group, others

2.  **Numeric/Octal notation** (0-7)

    1.  Add values: r=4, w=2, x=1
    
    2.  Example: **644** = rw-r&#x2013;r&#x2013; (User: 6=4+2, Group: 4, Others: 4)
    
    3.  Common permissions:
    
        1.  **755** (rwxr-xr-x): Executable files and directories
        
        2.  **644** (rw-r&#x2013;r&#x2013;): Regular files
        
        3.  **600** (rw--&#x2013;&#x2014;): Private files


<a id="org9f2e368"></a>

### **chmod**, **chown**, and **chgrp** commands

1.  **chmod**: Change permissions

2.  **chown**: Change owner

3.  **chgrp**: Change group


<a id="orgac5b421"></a>

### **Special permissions**: SUID, SGID, and sticky bit


<a id="org25614d7"></a>

# PowerShell for Security Professionals

1.  Contrary to what you might believe, PowerShell is *not* exclusive to Windows

2.  PowerShell can be installed on Linux distributions and FreeBSD


<a id="org742ef98"></a>

## PowerShell Security Features


<a id="org166e407"></a>

### Execution policies


<a id="org7472756"></a>

### Constrained language mode


<a id="orgbade34b"></a>

### Script Block Logging


<a id="orgbff34f6"></a>

### System-wide transcription


<a id="orgd181bdb"></a>

### Anti-malware Scan Interface (AMSI)

1.  PowerShell commands are **not** case-sensitive, so "Get-Acl" and "get-acl" are both valid.

2.  "Get-Acl" is the preferred method and the standard.


<a id="org5288798"></a>

## Security Cmdlets and Modules


<a id="org75125c2"></a>

### **Get-Acl** and **Set-Acl**: Managing access control lists


<a id="org8d9c8df"></a>

### **ConvertTo-SecureString** and **ConvertFrom-SecureString**: Working with secure strings


<a id="org958cb35"></a>

### **Get-Credential**: Securely requesting credentials


<a id="orgd997f1f"></a>

### **Set-ExecutionPolicy**: Managing PowerShell execution policies


<a id="orgd8817e9"></a>

### Security-related modules: **SecurityFever**, **PSAD**, **PSWinReporting**


<a id="org59011ae"></a>

## Security Monitoring and Auditing


<a id="org7b024e7"></a>

### Querying event logs for security events


<a id="org609ae04"></a>

### Monitoring file and registry changes


<a id="org40d09c9"></a>

### Auditing system configuration


<a id="orgebabdd1"></a>

### Detecting unauthorized PowerShell usage


<a id="org0bcf4bb"></a>

## Common Security Tasks in PowerShell

    # Get security event logs
    Get-EventLog -LogName Security -Newest 50
    
    # Check for failed login attemps
    Get-EventLog -LogName Security -EntryType FailureAudit | Where-Object {$_.EventID -eq 4625}
    
    # List running processes and their paths
    Get-Process | Select-Object Name, Path | Sort-Object Name
    
    # Get service details including startup accounts
    Get-WmiObject -Class Win32_Service | Select-Object Name, StartName, PathName


<a id="org44c94d1"></a>

# PowerShell Basics


<a id="orgb226eaf"></a>

## Introduction to PowerShell


<a id="orgba6529f"></a>

### PowerShell vs. Command Prompt

1.  **PowerShell** is object-oriented, using .NET objects, whereas **CMD** is text-based.

2.  Poershell supports **scripting**, **automation**, **remote management**, and integrates with **Windows APIs** and **WMI/COM**.

3.  CMD is limited to basic command execution; PowerShell is more powerfull and flexible.


<a id="org57c5aa1"></a>

### PowerShell versions and features

1.  **Windows PowerShell** (v1-5.1): Built on .NET Framework; Windows-only.

2.  **PowerShell 7+** (PowerShell Core): Cross-platform, built on .NET Core/.NET 5+.

3.  Newer version offer parallel execution (**ForEach-Object -Parallel**), better error handling, and more modules.


<a id="orgb405fec"></a>

### PowerShell ISE and PowerShell 7

1.  **ISE** (Integrated Scripting Environment): GUI editor for scripts with debugging tools (Windows only, deprecated in favor of VS Code).

2.  **PowerShell 7**: Preferred modern version, cross-platform, maintained actively with new features and bug fixes.


<a id="org2686616"></a>

## PowerShell Syntax and Structure


<a id="orgb4066ae"></a>

### **Cmdlets**: verb-noun naming convention

1.  Cmdlets follow a strict **Verb-Noun** pattern (e.g., **Get-Process**, **Set-Date**)

2.  Use **Get-Verb** to see allowed verbs.

3.  Promotes readabilit and consistency.


<a id="org83ca210"></a>

### Parameters and parameter sets

1.  Cmdlets accept **parameters** to control behavior (e.g., **-Name**, **-Path**)

2.  **Parameter sets** allow different combinations of parameters for the same cmdlet (mutually exclusive sets).


<a id="orgb491922"></a>

### Pipelines and object handling

1.  Output from one c,dlet can be passed to another using the **|** (pipe) operator.

2.  Unlike CMD, PowerShell pipes **objects**, not just text.

    1.  *Example*: **Get-Process | Where-Object {$<sub>.CPU</sub> -gt 10}**


<a id="org0c687f7"></a>

### Variables and data types

1.  Variables are declaed with **$** (e.g., **$name = "Alice"**).

2.  Supports various types: **strings**, **integers**, **arrays**, **hashtables**, **custom objects**.

3.  Type enforcement: **[int]$age = 25**


<a id="orga57755e"></a>

### Comparison operators

1.  Use **-eq** (equals), **-ne** (not equal), **-lt** (less than), **-gt** (greater than), etc.

2.  String comparison: **-like**, **-match**, **-contains**

    1.  *Example*: **$name -like "A**"\* (wildcard match)


<a id="orgddbfc12"></a>

### Flow control (if, switch, loops)

1.  **if\*/\*elseif\*/\*else** for conditionals.

2.  **switch** for multi-branch selection.

3.  *Loops*: **for**, **foreach**, **while**, **do-until**

    1.  *Example*:
    
            foreach ($item in $items) {
                Write-Output $item
            }


<a id="org5560263"></a>

## Essential PowerShell Cmdlets


<a id="org74a804b"></a>

### **Get-Command**: Discover available commands

1.  *Example*: **Get-Command -Name **Service****


<a id="org6770e86"></a>

### **Get-Help**: Access documentation and examples

1.  Displays help info for any cmdlet.

2.  Use **-Examples** or **-Online** for more guidance.

    1.  *Example*: **Get-Help Get-Process -Examples**


<a id="orgc50125d"></a>

### **Get-Process**: List running processes

1.  Useful for monitoring and filtering.

    1.  *Example*: **Get-Process | Sort-Object CPU -Descending**


<a id="org6fb474b"></a>

### **Get-Service**: List system services

1.  Control services with **Start-Service**, **Stop-Service**.


<a id="org8fa9368"></a>

### **Get-Item** and **Get-ChildItem**: Work with files and directories

1.  **Get-Item**: Gets a single file or folder.

2.  **Get-ChildItem** (alias **ls**): Lists contents of directories, can recurse with **-Recurse**


<a id="org75c608b"></a>

### **Import-Module**: Load additional functionality

1.  Loads external modules (e.g., **ActiveDirectory**, **Az**, **PSReadLine**)

2.  Required to access advanced cmdlets not available by default.


<a id="org646c8e7"></a>

### **Invoke-Command**: Run commands on remote computers

1.  Executes PowerShell scripts or cmdlets remotely (*requires PS Remoting enabled*).

    1.  *Example*: **Invoke-Command -ComputerName BSCS-PBX-DEBIAN -ScriptBlock { Get-Service }**


<a id="orga2d0495"></a>

## PowerShell Scripting


<a id="org0e47c08"></a>

### Script creation and execution

1.  Save scripts with **.ps1** extension.

2.  Run scripts with **.\\<script>.ps1**; may need to adjust execution policy:

    1.  **Set-ExecutionPolicy RemoteSigned**


<a id="org95b426d"></a>

### Functions and modules

1.  Define reusable blocks with **funciton** keyword.

2.  Group related functions into a module (**.psm1**) for reuse and sharing.


<a id="orgcce519a"></a>

### Error handling and debugging

1.  Use **try/catch/finally** blocks to handle exceptions.

2.  **$Error** holds recent errors.

3.  Use **Write-Debug**, **Write-Verbose** for logging and diagnostics.


<a id="orgaf50d43"></a>

### Script signing and security

1.  Scripts may need to be signed with a code-signing certificate.

2.  Execution policies:

    1.  **Restricted**: No scripts run.
    
    2.  **AllSigned**: Only signed scripts run.
    
    3.  **RemoteSigned**: Local scripts run, downloaded ones must be signed.


<a id="org2231056"></a>

### PowerShell profiles

1.  Customize your environment using profile scripts (**$PROFILE**)

2.  Add aliases, functions, variables, or modules for startup


<a id="org2c8a2f1"></a>

# Python Basics


<a id="org675ac72"></a>

## Python Environment Setup


<a id="org566b360"></a>

### Installing Python

1.  Download from Python's website or use your system's package manager (*xbps* for me)

2.  Ensure Python is added to your system's **PATH** for CLI use.

3.  Check installation with: **python &#x2013;version** or **python3 &#x2013;version**


<a id="orgc6df576"></a>

### Package management with **pip**

1.  **pip** is Python's default package manager.

2.  Install external libraries: **pip install <library>**

3.  Use **pip list**, **pip show**, **pip uninstall** to manage packages.

4.  Can also install from *requirements.txt*: **pip install -r requirements.txt**.


<a id="org1b18588"></a>

### Virtual environments (venv, conda)

1.  Avoid conflicts by isolating dependencies per project.

2.  **venv** (built-in for most Python versions):

        python -m venv venv
        source venv/bin/activate # Unix
        venv\Scripts\activate    # Windows

3.  **conda** (Anaconda/Miniconda): full-featured environment + package manager.


<a id="org9c4eefd"></a>

### Code editors and IDEs (VS Code, PyCharm, NeoVim, Emacs)

1.  **VS Code**: Lightweight, extensible, rich Python extension; by Microsoft.

2.  **PyCharm**: Full-featured IDE with powerful refactoring/debugging tools; by JetBrains.

3.  **NeoVim/Emacs**: Highly customizable editors, popular in dev-heavy workflows.


<a id="org000709c"></a>

## Python Syntax Fundamentals


<a id="orgc4b33f9"></a>

### Variables and data types

1.  **Dynamically typed**: types inferred at runtime.

2.  Common types: **int**, **float**, **str**, **bool**, **list**, **dict**, **NoneType**.


<a id="orgc13cf76"></a>

### Operators and expressions

1.  **Arithmetic**: +, -, **, *, /*, %, \***

2.  **Comparison**: `=, !`, >, <, >=, <=

3.  **Logical**: and, or, not

4.  **Membership**: in, not in


<a id="orgf4d1984"></a>

### Control flow: if statements, loops

1.  Conditional logic:

        if x > 10:
            print("Large")
        elif x == 10:
            print("Equal")
        else:
            print("Small")

2.  Loops:

    1.  **for** loop: iterate over a sequence
    
    2.  **while** loop: repreat while condition is true
    
    3.  **break**, **continue**, **else** clause for loops


<a id="org71ea77d"></a>

### Functions and modules

1.  Define reusable code with **def**:

        def greet(name):
            return f"Hello, {name}"

2.  Import standard or cusotm modules:

        import math
        from custom_module import greet


<a id="orgf19b901"></a>

### Error handling with try/except

1.  Catch nd handle runtime exceptions:

        try:
            result = 10 / x
        except ZeroDivisionError:
            print("Cannot divide by zero.")
        finally:
            print("Cleanup code here.")


<a id="orgf4a2e08"></a>

## Data Structures


<a id="orgadbf585"></a>

### Lists, tuples, and sets

1.  **List**: Ordered, mutable: **my<sub>list</sub> = [1, 2, 3]**

2.  **Tuple**: Ordered, immutable: **my<sub>tuple</sub> = (1, 2)**

3.  **Set**: Unordered, unique items: **my<sub>set</sub> = {1, 2, 3}**


<a id="org1a8fb2d"></a>

### Dictionaries

1.  Key-value pairs:

        user = {"name": "Alive", "age": 30}
        print(user["name"])


<a id="orgae71f5b"></a>

### List comprehensions

1.  Concise syntax for creating lists:

        squares = [x**2 for x in range(10)]


<a id="orgd46ea3f"></a>

### String manipulation

1.  Strings are immutable.

2.  Common methods: **.lower()**, **.upper()**, **.strip()**, **.split()**, **.replace()**, **.find()**

3.  f-strings for formatting:

        name = "Alice"
        print(f"Hello, {name}")


<a id="orgc087b6b"></a>

## File Operations


<a id="org596f075"></a>

### Reading and writing text files

1.  Use **open()** with context manager:

        with open("file.txt", "r") as f:
            content = f.read()
        with open("output.txt", "w") as f:
            f.write("Hello")


<a id="org5ae0d2f"></a>

### Working with CSV and JSON

1.  **CSV**:

        import csv
        with open("data.csv") as f:
            reader = csv.reader(f)
            for row in reader:
                print(row)

2.  **JSON**:

        import json
        data = {"name": "Alex"}
        json_str = json.dumps(data)
        new_data = json.loads(json_str)


<a id="org1411b7f"></a>

### File path manipulation

1.  Use **os** and **pathlib**:

        import os
        from pathlib import Path
        
        current = Path.cwd()
        new_file = current / "folder" / "file.txt"
        print(new_file.exists())


<a id="org5124318"></a>

## Python for System Administration

      # Simple system monitoring script
      import psutil
      import platform
      from datetime import datetime
    
      def get_system_info():
        print("="*40, "System Information", "="*40)
        uname = platform.uname()
        print(f"System: {uname.system}")
        print(f"Node Name: {uname.node}")
        print(f"Release: {uname.release}")
        print(f"Version: {uname.version}")
        print(f"Machine: {uname.machine}")
        print(f"Processor: {uname.processor}")
    
        # Boot Time
        print("="*40, "Boot Time", "="*40)
        boot_time_timestamp = psutil.boot_time()
        bt = datetime.fromtimestamp(boot_time_timestamp)
        print(f"Boot Time: {bt.year}/{bt.month}/{bt.day} {bt.hour}:{bt.minute}:{bt.second}")
    
        # CPU Info
        print("="*40, "CPU Info", "="*40)
        print(f"Physical cores: {psutil.cpu_count(logical=False)}")
        print(f"Total cores: {psutil.cpu_count(logical=True)}")
    
        # Memory Information
        print("="*40, "Memory Information", "="*40)
        svmem = psutil.virtual_memory()
        print(f"Total: {get_size(svmem.total)}")
        print(f"Available: {get_size(svmem.available)}")
        print(f"Used: {get_size(svmem.used)}")
        print(f"Percentage: {svmem.percent}%")
    
    def get_size(bytes, suffix="B"):
        factor = 1024
        for unit in ["", "K", "M", "G", "T", "P"]:
            if bytes < factor:
                return f"{bytes:.2f}{unit}{suffix}"
            bytes /= factor
    
    if __name__ == "__main__":
        get_system_info()


<a id="org4cc097f"></a>

# Windows File System Basics

1.  **NTFS (New Technology File System)** is the default file system for modern Windows.


<a id="org30800d8"></a>

## NTFS File System


<a id="org2481e7f"></a>

### Features and benefits of NTFS

1.  **Journaling** for improved reliability.

2.  **Large file support** (up to 16 TB on modern systems).

3.  **Security** through file and folder permissions.

4.  **Compression**, **encryption**, and **disk quotas** support.

5.  **Transaction logging** to help recover from crashes.


<a id="org80d36d1"></a>

### NTFS permissions model

1.  Use **Discretionary Access Control** to define who can access files/folders.

2.  Two main permission types:

    1.  **Basic permissions**: Read, Write, Modify, Full Control
    
    2.  **Advanced permissions**: Fine-grained (e.g., Delete, Change Permissions, Take Ownership).

3.  Permissions can be **Allow** or **Deny**.


<a id="org54de680"></a>

### Access Control Lists (ACLs)

1.  **ACLs** store the security descriptors for files/folders.

2.  Each object has:

    1.  **DACL (Discretionary ACL)**: Lists users/groups and their permissions.
    
    2.  **SACL (System ACL)**: Controls auditing for object access.

3.  Managed via GUI (Security tab) or CLI (**icacls**, PowerShell).


<a id="orgd1ca2d8"></a>

### NTDS special permissions

1.  **NTDS** refers to the **NT Directory Services** (Active Directory database).

2.  Certain permissions are only applicable in AD environments:

    1.  **Replicating Directory Changes**, **Full Control of AD objects**, etc.

3.  Often managed via tools like **ADSI Edit** or Group Policy


<a id="orgb8c133d"></a>

### Alternate Data Streams

1.  NTFS allows storing **multiple data streams** within a single file.

2.  Commonly used for metadata (e.g, file download origin).

3.  Can be used maliciously to hide data: **echo Hidden > normal.txt:hidden.txt**

4.  Use **dir /r** or PowerShell to view ADS.


<a id="org6e9270c"></a>

## NTFS Folder Structure


<a id="orga8ab58b"></a>

### Windows system folders

1.  Core folders include:

    1.  **C:\Windows**: OS files.
    
    2.  **C:\Program Files**: Installed apps.
    
    3.  **C:\Users**: User profiles and data.


<a id="orgbe18b50"></a>

### Program Files structure

1.  Two main directories:

    1.  **C:\Program Files**: For 64-bit applications.
    
    2.  **C:\Program Files (x86)**: For 32-bit applications on 64-bit Windows.

2.  Helps prevent DLL conflicts and ensures compatibility


<a id="org5168259"></a>

### User profiles and directories

1.  Stored in **C:\Users\\<username>**.

2.  Each profile includes:

    1.  **Desktop**, **Documents**, **Downloads**, **AppData** (hidden)

3.  **AppData** contains local and roaming settings.


<a id="orgffd3128"></a>

### System32 and SysWOW64

1.  **System32**: Contains core 64-bit system files (despite the misleading name).

2.  **SysWOW64**: Contains 32-bit binaries on 64-bit systems (Windows-on-Windows 64).

3.  Necessary for compatibility with legacy applications.


<a id="org3b024f1"></a>

### Windows Registry as a hierarchical database

1.  Central configuration store for OS, users, apps, and hardware.

2.  Structure like folders and keys:

    1.  Root keys: **HKEY<sub>LOCAL</sub><sub>MACHINE</sub>**, **HKEY<sub>CURRENT</sub><sub>USER</sub>**, etc.

3.  Accessed with **regedit.exe** or via PowerShell (**Get-ItemProperty**, etc.)


<a id="orgcd56d3f"></a>

## File and Folder Management


<a id="org87ef9b6"></a>

### File Explorer vs. Command Line tools

1.  **File Explorer**: GUI tool for managing files (drag/drop, rename, properties).

2.  **Command Line tools**:

    1.  **cmd**: dir, copy, xcopy, robocopy
    
    2.  **PowerShell**: modern and scriptable file management.


<a id="org8213290"></a>

### PowerShell file management cmdlets

1.  Key cmdlets:

    1.  **Get-ChildItem** (**ls**)
    
    2.  **Copy-Item**, **Move-Item**
    
    3.  **Remove-Item**, **Rename-Item**
    
    4.  **New-Item** to create files/folders.
    
    5.  Example: **Copy-Item -Path "C:\File.txt" -Destination "D:\Backup"**


<a id="org791e276"></a>

### Advanced file operations

1.  **Roboopy** for efficient file copy/mirroring with retry logic.

2.  **Symbolic links** (**mklink**): create virtual pointers to files or folders.

3.  **Scheduled Tasks** or scripts to automate file cleanup or backup.


<a id="orge190e15"></a>

### File system auditing

1.  Configure via **Local Security Policy** or **Group Policy**

2.  Use SACLs to define audit rules for objects.

3.  Events logged in the **Security Event Log** (Event Viewer).


<a id="org209b831"></a>

## Security Features


<a id="org53cb4fd"></a>

### BitLocker drive encryption

1.  Protects data even if drive is removed.

2.  Uses **TPM (Trusted Platform Module)** or password/key for unlocking.


<a id="org91adcb4"></a>

### EFS (Encrypting File System)

1.  Encrypts individual files/folders at the file system level.

2.  Tied to user account; only the same user or authorized recovery agent can decrypt.

3.  Enabled via file properties -> Advanced -> Encrypt.


<a id="orge239088"></a>

### File ownership and inheritance

1.  **Ownership** determines who can modify permissions.

2.  By default, the creator is the owner.

3.  **Inheritance**: Child objects inherit permissions from parent folder unless explicitly blocked.


<a id="org1a5c4a1"></a>

### File system quotas

1.  Limit user disk space on NTFS volumes.

2.  Set per user via **fsutil** or GUI (Volume Properties -> Quota).

3.  Helps prevent users from filling up disk space.


<a id="org7af04de"></a>

### Shadow copies and previous versions

1.  Snapshot feature for recovering changed or deleted files.

2.  Accessed via file/folder properties -> Previous Versions tab.

3.  Powered by **Volume Shadow Copy Service (VSS)**.


<a id="orgee3ad6e"></a>

## Common File System Management Tasks


<a id="org955894f"></a>

### Sharing folders and setting permissions


<a id="orgfac33b3"></a>

### Using administrative shares


<a id="orgd881c53"></a>

### Managing disk quotas


<a id="org2e2cd8e"></a>

### Configuring file screening


<a id="org68d60ce"></a>

### Implementing NTFS compression


<a id="org32f03f9"></a>

# Documentation Basics


<a id="orgfefb487"></a>

## Importance of Documentation for System Administrators


<a id="org98dace6"></a>

### **Knowledge retention**: Preserving institutional knowledge


<a id="orgcb847fc"></a>

### **Troubleshooting efficiency**: Recording solutions for future reference


<a id="org0a75e73"></a>

### **Onboarding**: Helping new team members get up to speed


<a id="org1eaac42"></a>

### **Disaster recovery**: Documenting cirtical systems and procedures


<a id="orgd95d8e4"></a>

### **Compliance**: Meeting regulatory and audit requirements


<a id="orgdc4380e"></a>

## Effective Documentation Practices


<a id="orgeaf969d"></a>

### Keep documentation current and regularly updated


<a id="org6c383a8"></a>

### Use conistent formatting and templates


<a id="org0bc0644"></a>

### Include dates, versions, and author information


<a id="orgdb52bdf"></a>

### Document assumptions and dependencies


<a id="org3bb670a"></a>

### Balance detail with readability


<a id="org835c8e8"></a>

### Link related documentation where appropriate


<a id="org9cfec27"></a>

### Use diagrams and screenshots for visual clarity


<a id="orgf5e3f44"></a>

### Include examples for complex procedures


<a id="org78e1f78"></a>

## Documentation Tools and Formats


<a id="org396a440"></a>

### **Plain Text Editors**:

1.  Advantanges: Universal comaptibility, lightweight, version-control friendly

2.  Common editors: Vim, Emacs, Nano, Notepad++


<a id="org8c57dae"></a>

### **Markup Languages**:

1.  Markdown:

    1.  Lightweight markup language with plain-text formatting syntax
    
    2.  GitHub-flavored Markdown for code repositories
    
    3.  Tools: Typhora, VS Code with Markdown extensions, Obsidian, Emacs

2.  Org Mode:

    1.  (*I used Org Mode within Spacemacs to write this course!*)
    
    1.  Powerful markup language and notetaking/organizational tool within Emacs
    
    2.  Features: TODO lists, project planning, literate programming
    
    3.  Excellent for technical documentation with executable code blocks

3.  LaTeX:

    1.  Professional typesetting system for complex documents
    
    2.  Excellent for formal documentation with mathematical notation
    
        1.  A lot of scientific papers are written using LaTeX (but some are written using Markdown or Org Mode)
    
    3.  Tools: TeXStudio, Overleaf (online), Notepad++, Emacs, etc.


<a id="orgd9ccdbf"></a>

### **Wiki Systems**:

1.  **MediaWiki**: Powers Wikipedia, highly customizable

2.  **DokuWiki**: File-based wiki that doesn't require a database

3.  **Confluence**: Commercial solution with extensive integration options

    1.  Benefits: Collaboration editing, version history, search capabilities


<a id="org7e94cc3"></a>

### **Specialized Documentation Tools**:

1.  **IT Documentation**: ITGlue, Document360

2.  **Network diagrams**: draw.io, Lucidchart, Visio

3.  **API Documentation**: Swagger, ReadTheDocs

4.  **Knowledge bases**: BookStack, MkDocs


<a id="org66a36da"></a>

## Documentation Workflows for SysAdmins


<a id="org8716465"></a>

### **Incident documentation**: Recording incidencts, troubleshooting steps, and resolutions


<a id="org96239ad"></a>

### **Change management**: Documenting system changes, approvals, and rollback procedures


<a id="orge421050"></a>

### **Standard Operating Procedures (SOPs)**: Creating step-by-step instructions for routine tasks


<a id="orgdfb4a9b"></a>

### **Architecture documentation**: Recording system design and dependencies


<a id="orgcdb9c9b"></a>

### **Runbooks**: Creating operation guides for specific scenarios

