.data # Data declaration section

.text # Assembly language instructions
    main: # Start of code section
    li $a0, 5
    li $v0, 1
    syscall
jr $ra