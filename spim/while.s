.data
    out_string: .asciiz "\nFLPS"
.text
main:
    li $t1, 0 #index
    li $t2, 10 #iterations

    Loop:
        beq $t2,$t1, Exit
        li $v0, 4
        la $a0, out_string
        syscall
        add $t1, $t1, 1
        #
        i++

        j Loop
    Exit:
    jr $ra