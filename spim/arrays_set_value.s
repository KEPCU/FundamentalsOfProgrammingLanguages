.data
    text1: .asciiz "\n number: "
    text2: .asciiz "\n"
    array_1: .word 1:10
.text
main:
    

    li $t0, 0
    la $t1, array_1
    loop1: # set values

        # print text1
        li $v0, 4
        la $a0, text1
        syscall

        # get age
        li $v0,5
        syscall
        # print int
        move $a0, $v0
        li $v0, 1
        syscall
        move $t3,$a0

        bge $t0, 10, end_loop1
        sw $t3, 0($t1) # set value
        add $t1, $t1, 4 # increment pointer
        add $t0, $t0, 1 # i++
        j loop1
    end_loop1:

    # print text1
    li $v0, 4
    la $a0, text2
    syscall

    li $t0, 0
    la $t1, array_1
    loop2:
        # show values
        bge $t0, 10, end_loop2
        lw $t2, 0($t1)
        add $t1, $t1, 4
        li $v0, 1
        # print value
        move $a0, $t2
        syscall
        li $a0, 32
        # print space
        li $v0, 11
        syscall
        add $t0, $t0, 1
        # i++
        j loop2
    end_loop2:

    jr $ra

