// linux/mksysnum.pl -Wall -Werror -static -I/tmp/include -m32 /tmp/include/asm/unistd.h
// Code generated by the command above; see README.md. DO NOT EDIT.

// +build 386,linux

package unix

const (
	SYS_RESTART_SYSCALL        = 0
	SYS_EXIT                   = 1
	SYS_FORK                   = 2
	SYS_READ                   = 3
	SYS_WRITE                  = 4
	SYS_OPEN                   = 5
	SYS_CLOSE                  = 6
	SYS_WAITPID                = 7
	SYS_CREAT                  = 8
	SYS_LINK                   = 9
	SYS_UNLINK                 = 10
	SYS_EXECVE                 = 11
	SYS_CHDIR                  = 12
	SYS_TIME                   = 13
	SYS_MKNOD                  = 14
	SYS_CHMOD                  = 15
	SYS_LCHOWN                 = 16
	SYS_BREAK                  = 17
	SYS_OLDSTAT                = 18
	SYS_LSEEK                  = 19
	SYS_GETPID                 = 20
	SYS_MOUNT                  = 21
	SYS_UMOUNT                 = 22
	SYS_SETUID                 = 23
	SYS_GETUID                 = 24
	SYS_STIME                  = 25
	SYS_PTRACE                 = 26
	SYS_ALARM                  = 27
	SYS_OLDFSTAT               = 28
	SYS_PAUSE                  = 29
	SYS_UTIME                  = 30
	SYS_STTY                   = 31
	SYS_GTTY                   = 32
	SYS_ACCESS                 = 33
	SYS_NICE                   = 34
	SYS_FTIME                  = 35
	SYS_SYNC                   = 36
	SYS_KILL                   = 37
	SYS_RENAME                 = 38
	SYS_MKDIR                  = 39
	SYS_RMDIR                  = 40
	SYS_DUP                    = 41
	SYS_PIPE                   = 42
	SYS_TIMES                  = 43
	SYS_PROF                   = 44
	SYS_BRK                    = 45
	SYS_SETGID                 = 46
	SYS_GETGID                 = 47
	SYS_SIGNAL                 = 48
	SYS_GETEUID                = 49
	SYS_GETEGID                = 50
	SYS_ACCT                   = 51
	SYS_UMOUNT2                = 52
	SYS_LOCK                   = 53
	SYS_IOCTL                  = 54
	SYS_FCNTL                  = 55
	SYS_MPX                    = 56
	SYS_SETPGID                = 57
	SYS_ULIMIT                 = 58
	SYS_OLDOLDUNAME            = 59
	SYS_UMASK                  = 60
	SYS_CHROOT                 = 61
	SYS_USTAT                  = 62
	SYS_DUP2                   = 63
	SYS_GETPPID                = 64
	SYS_GETPGRP                = 65
	SYS_SETSID                 = 66
	SYS_SIGACTION              = 67
	SYS_SGETMASK               = 68
	SYS_SSETMASK               = 69
	SYS_SETREUID               = 70
	SYS_SETREGID               = 71
	SYS_SIGSUSPEND             = 72
	SYS_SIGPENDING             = 73
	SYS_SETHOSTNAME            = 74
	SYS_SETRLIMIT              = 75
	SYS_GETRLIMIT              = 76
	SYS_GETRUSAGE              = 77
	SYS_GETTIMEOFDAY           = 78
	SYS_SETTIMEOFDAY           = 79
	SYS_GETGROUPS              = 80
	SYS_SETGROUPS              = 81
	SYS_SELECT                 = 82
	SYS_SYMLINK                = 83
	SYS_OLDLSTAT               = 84
	SYS_READLINK               = 85
	SYS_USELIB                 = 86
	SYS_SWAPON                 = 87
	SYS_REBOOT                 = 88
	SYS_READDIR                = 89
	SYS_MMAP                   = 90
	SYS_MUNMAP                 = 91
	SYS_TRUNCATE               = 92
	SYS_FTRUNCATE              = 93
	SYS_FCHMOD                 = 94
	SYS_FCHOWN                 = 95
	SYS_GETPRIORITY            = 96
	SYS_SETPRIORITY            = 97
	SYS_PROFIL                 = 98
	SYS_STATFS                 = 99
	SYS_FSTATFS                = 100
	SYS_IOPERM                 = 101
	SYS_SOCKETCALL             = 102
	SYS_SYSLOG                 = 103
	SYS_SETITIMER              = 104
	SYS_GETITIMER              = 105
	SYS_STAT                   = 106
	SYS_LSTAT                  = 107
	SYS_FSTAT                  = 108
	SYS_OLDUNAME               = 109
	SYS_IOPL                   = 110
	SYS_VHANGUP                = 111
	SYS_IDLE                   = 112
	SYS_VM86OLD                = 113
	SYS_WAIT4                  = 114
	SYS_SWAPOFF                = 115
	SYS_SYSINFO                = 116
	SYS_IPC                    = 117
	SYS_FSYNC                  = 118
	SYS_SIGRETURN              = 119
	SYS_CLONE                  = 120
	SYS_SETDOMAINNAME          = 121
	SYS_UNAME                  = 122
	SYS_MODIFY_LDT             = 123
	SYS_ADJTIMEX               = 124
	SYS_MPROTECT               = 125
	SYS_SIGPROCMASK            = 126
	SYS_CREATE_MODULE          = 127
	SYS_INIT_MODULE            = 128
	SYS_DELETE_MODULE          = 129
	SYS_GET_KERNEL_SYMS        = 130
	SYS_QUOTACTL               = 131
	SYS_GETPGID                = 132
	SYS_FCHDIR                 = 133
	SYS_BDFLUSH                = 134
	SYS_SYSFS                  = 135
	SYS_PERSONALITY            = 136
	SYS_AFS_SYSCALL            = 137
	SYS_SETFSUID               = 138
	SYS_SETFSGID               = 139
	SYS__LLSEEK                = 140
	SYS_GETDENTS               = 141
	SYS__NEWSELECT             = 142
	SYS_FLOCK                  = 143
	SYS_MSYNC                  = 144
	SYS_READV                  = 145
	SYS_WRITEV                 = 146
	SYS_GETSID                 = 147
	SYS_FDATASYNC              = 148
	SYS__SYSCTL                = 149
	SYS_MLOCK                  = 150
	SYS_MUNLOCK                = 151
	SYS_MLOCKALL               = 152
	SYS_MUNLOCKALL             = 153
	SYS_SCHED_SETPARAM         = 154
	SYS_SCHED_GETPARAM         = 155
	SYS_SCHED_SETSCHEDULER     = 156
	SYS_SCHED_GETSCHEDULER     = 157
	SYS_SCHED_YIELD            = 158
	SYS_SCHED_GET_PRIORITY_MAX = 159
	SYS_SCHED_GET_PRIORITY_MIN = 160
	SYS_SCHED_RR_GET_INTERVAL  = 161
	SYS_NANOSLEEP              = 162
	SYS_MREMAP                 = 163
	SYS_SETRESUID              = 164
	SYS_GETRESUID              = 165
	SYS_VM86                   = 166
	SYS_QUERY_MODULE           = 167
	SYS_POLL                   = 168
	SYS_NFSSERVCTL             = 169
	SYS_SETRESGID              = 170
	SYS_GETRESGID              = 171
	SYS_PRCTL                  = 172
	SYS_RT_SIGRETURN           = 173
	SYS_RT_SIGACTION           = 174
	SYS_RT_SIGPROCMASK         = 175
	SYS_RT_SIGPENDING          = 176
	SYS_RT_SIGTIMEDWAIT        = 177
	SYS_RT_SIGQUEUEINFO        = 178
	SYS_RT_SIGSUSPEND          = 179
	SYS_PREAD64                = 180
	SYS_PWRITE64               = 181
	SYS_CHOWN                  = 182
	SYS_GETCWD                 = 183
	SYS_CAPGET                 = 184
	SYS_CAPSET                 = 185
	SYS_SIGALTSTACK            = 186
	SYS_SENDFILE               = 187
	SYS_GETPMSG                = 188
	SYS_PUTPMSG                = 189
	SYS_VFORK                  = 190
	SYS_UGETRLIMIT             = 191
	SYS_MMAP2                  = 192
	SYS_TRUNCATE64             = 193
	SYS_FTRUNCATE64            = 194
	SYS_STAT64                 = 195
	SYS_LSTAT64                = 196
	SYS_FSTAT64                = 197
	SYS_LCHOWN32               = 198
	SYS_GETUID32               = 199
	SYS_GETGID32               = 200
	SYS_GETEUID32              = 201
	SYS_GETEGID32              = 202
	SYS_SETREUID32             = 203
	SYS_SETREGID32             = 204
	SYS_GETGROUPS32            = 205
	SYS_SETGROUPS32            = 206
	SYS_FCHOWN32               = 207
	SYS_SETRESUID32            = 208
	SYS_GETRESUID32            = 209
	SYS_SETRESGID32            = 210
	SYS_GETRESGID32            = 211
	SYS_CHOWN32                = 212
	SYS_SETUID32               = 213
	SYS_SETGID32               = 214
	SYS_SETFSUID32             = 215
	SYS_SETFSGID32             = 216
	SYS_PIVOT_ROOT             = 217
	SYS_MINCORE                = 218
	SYS_MADVISE                = 219
	SYS_GETDENTS64             = 220
	SYS_FCNTL64                = 221
	SYS_GETTID                 = 224
	SYS_READAHEAD              = 225
	SYS_SETXATTR               = 226
	SYS_LSETXATTR              = 227
	SYS_FSETXATTR              = 228
	SYS_GETXATTR               = 229
	SYS_LGETXATTR              = 230
	SYS_FGETXATTR              = 231
	SYS_LISTXATTR              = 232
	SYS_LLISTXATTR             = 233
	SYS_FLISTXATTR             = 234
	SYS_REMOVEXATTR            = 235
	SYS_LREMOVEXATTR           = 236
	SYS_FREMOVEXATTR           = 237
	SYS_TKILL                  = 238
	SYS_SENDFILE64             = 239
	SYS_FUTEX                  = 240
	SYS_SCHED_SETAFFINITY      = 241
	SYS_SCHED_GETAFFINITY      = 242
	SYS_SET_THREAD_AREA        = 243
	SYS_GET_THREAD_AREA        = 244
	SYS_IO_SETUP               = 245
	SYS_IO_DESTROY             = 246
	SYS_IO_GETEVENTS           = 247
	SYS_IO_SUBMIT              = 248
	SYS_IO_CANCEL              = 249
	SYS_FADVISE64              = 250
	SYS_EXIT_GROUP             = 252
	SYS_LOOKUP_DCOOKIE         = 253
	SYS_EPOLL_CREATE           = 254
	SYS_EPOLL_CTL              = 255
	SYS_EPOLL_WAIT             = 256
	SYS_REMAP_FILE_PAGES       = 257
	SYS_SET_TID_ADDRESS        = 258
	SYS_TIMER_CREATE           = 259
	SYS_TIMER_SETTIME          = 260
	SYS_TIMER_GETTIME          = 261
	SYS_TIMER_GETOVERRUN       = 262
	SYS_TIMER_DELETE           = 263
	SYS_CLOCK_SETTIME          = 264
	SYS_CLOCK_GETTIME          = 265
	SYS_CLOCK_GETRES           = 266
	SYS_CLOCK_NANOSLEEP        = 267
	SYS_STATFS64               = 268
	SYS_FSTATFS64              = 269
	SYS_TGKILL                 = 270
	SYS_UTIMES                 = 271
	SYS_FADVISE64_64           = 272
	SYS_VSERVER                = 273
	SYS_MBIND                  = 274
	SYS_GET_MEMPOLICY          = 275
	SYS_SET_MEMPOLICY          = 276
	SYS_MQ_OPEN                = 277
	SYS_MQ_UNLINK              = 278
	SYS_MQ_TIMEDSEND           = 279
	SYS_MQ_TIMEDRECEIVE        = 280
	SYS_MQ_NOTIFY              = 281
	SYS_MQ_GETSETATTR          = 282
	SYS_KEXEC_LOAD             = 283
	SYS_WAITID                 = 284
	SYS_ADD_KEY                = 286
	SYS_REQUEST_KEY            = 287
	SYS_KEYCTL                 = 288
	SYS_IOPRIO_SET             = 289
	SYS_IOPRIO_GET             = 290
	SYS_INOTIFY_INIT           = 291
	SYS_INOTIFY_ADD_WATCH      = 292
	SYS_INOTIFY_RM_WATCH       = 293
	SYS_MIGRATE_PAGES          = 294
	SYS_OPENAT                 = 295
	SYS_MKDIRAT                = 296
	SYS_MKNODAT                = 297
	SYS_FCHOWNAT               = 298
	SYS_FUTIMESAT              = 299
	SYS_FSTATAT64              = 300
	SYS_UNLINKAT               = 301
	SYS_RENAMEAT               = 302
	SYS_LINKAT                 = 303
	SYS_SYMLINKAT              = 304
	SYS_READLINKAT             = 305
	SYS_FCHMODAT               = 306
	SYS_FACCESSAT              = 307
	SYS_PSELECT6               = 308
	SYS_PPOLL                  = 309
	SYS_UNSHARE                = 310
	SYS_SET_ROBUST_LIST        = 311
	SYS_GET_ROBUST_LIST        = 312
	SYS_SPLICE                 = 313
	SYS_SYNC_FILE_RANGE        = 314
	SYS_TEE                    = 315
	SYS_VMSPLICE               = 316
	SYS_MOVE_PAGES             = 317
	SYS_GETCPU                 = 318
	SYS_EPOLL_PWAIT            = 319
	SYS_UTIMENSAT              = 320
	SYS_SIGNALFD               = 321
	SYS_TIMERFD_CREATE         = 322
	SYS_EVENTFD                = 323
	SYS_FALLOCATE              = 324
	SYS_TIMERFD_SETTIME        = 325
	SYS_TIMERFD_GETTIME        = 326
	SYS_SIGNALFD4              = 327
	SYS_EVENTFD2               = 328
	SYS_EPOLL_CREATE1          = 329
	SYS_DUP3                   = 330
	SYS_PIPE2                  = 331
	SYS_INOTIFY_INIT1          = 332
	SYS_PREADV                 = 333
	SYS_PWRITEV                = 334
	SYS_RT_TGSIGQUEUEINFO      = 335
	SYS_PERF_EVENT_OPEN        = 336
	SYS_RECVMMSG               = 337
	SYS_FANOTIFY_INIT          = 338
	SYS_FANOTIFY_MARK          = 339
	SYS_PRLIMIT64              = 340
	SYS_NAME_TO_HANDLE_AT      = 341
	SYS_OPEN_BY_HANDLE_AT      = 342
	SYS_CLOCK_ADJTIME          = 343
	SYS_SYNCFS                 = 344
	SYS_SENDMMSG               = 345
	SYS_SETNS                  = 346
	SYS_PROCESS_VM_READV       = 347
	SYS_PROCESS_VM_WRITEV      = 348
	SYS_KCMP                   = 349
	SYS_FINIT_MODULE           = 350
	SYS_SCHED_SETATTR          = 351
	SYS_SCHED_GETATTR          = 352
	SYS_RENAMEAT2              = 353
	SYS_SECCOMP                = 354
	SYS_GETRANDOM              = 355
	SYS_MEMFD_CREATE           = 356
	SYS_BPF                    = 357
	SYS_EXECVEAT               = 358
	SYS_SOCKET                 = 359
	SYS_SOCKETPAIR             = 360
	SYS_BIND                   = 361
	SYS_CONNECT                = 362
	SYS_LISTEN                 = 363
	SYS_ACCEPT4                = 364
	SYS_GETSOCKOPT             = 365
	SYS_SETSOCKOPT             = 366
	SYS_GETSOCKNAME            = 367
	SYS_GETPEERNAME            = 368
	SYS_SENDTO                 = 369
	SYS_SENDMSG                = 370
	SYS_RECVFROM               = 371
	SYS_RECVMSG                = 372
	SYS_SHUTDOWN               = 373
	SYS_USERFAULTFD            = 374
	SYS_MEMBARRIER             = 375
	SYS_MLOCK2                 = 376
	SYS_COPY_FILE_RANGE        = 377
	SYS_PREADV2                = 378
	SYS_PWRITEV2               = 379
	SYS_PKEY_MPROTECT          = 380
	SYS_PKEY_ALLOC             = 381
	SYS_PKEY_FREE              = 382
<<<<<<< HEAD
	SYS_STATX                  = 383
	SYS_ARCH_PRCTL             = 384
=======
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a
)
