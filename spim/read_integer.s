.data
.text
main: # Start of code section
    li $v0,5 # code 5 == read integer
    syscall # Invoke the operating system
    move $a0, $v0
    li $v0,1 # code 1 == print integer
    syscall # Invoke the operating system
    jr $ra