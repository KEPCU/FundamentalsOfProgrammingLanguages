.data
    out_string: .asciiz "\hello world!\n"
.text
    main:
    li $v0, 4 # system call code for printing = 4
    la $a0, out_string # load address to $a0
    syscall
    jr $ra # return to caller

    