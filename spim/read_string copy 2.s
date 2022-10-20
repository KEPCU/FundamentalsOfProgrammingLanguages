.data
    out_string_1: .asciiz "\nEnter a string:\n"
    out_string_2: .asciiz "\nYou wrote:\n"
    buffer: .space 20
.text
main: # Start of code section
    # print string
    li $v0, 4
    la $a0, out_string_1
    syscall

    # read string
    li $a1, 10
    li $v0, 8
    la $a0, buffer
    syscall

    # print string
    li $v0, 4
    la $a0, out_string_2
    syscall

    # print string
    li $v0, 4
    la $a0, buffer
    syscall
   
    jr $ra