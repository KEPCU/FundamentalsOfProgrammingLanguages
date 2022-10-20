.data
    out_string: .asciiz "\nThe Sum is:\n"
.text
main:
    li $t1, 0 #index
    li $t2, 5 #iterations

    Loop:
        beq $t2,$t1, Exit
        li $v0, 4
        li $v0,5
        syscall

        add $t3,$t3,$v0

        add $t1, $t1, 1
        # i++

    j Loop
    Exit:
    li $v0,4
    la $a0,out_string
    syscall

    li $v0,1
    move $a0,$t3
    syscall

    jr $ra