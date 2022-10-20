.data
    text1: .asciiz "\n ingrese un elemento del vector "
    text2: .asciiz " : "
    text3: .asciiz "\n el producto escalar es: "
    text4: .asciiz "\n"
    array_1: .word 1:3
    array_2: .word 1:3
.text
main:

    li $t0, 0
    la $t1, array_1

    # set values 1
    Loop: 
        bge $t0, 3, End_Loop

        # print text1
        li $v0, 4
        la $a0, text1
        syscall

        # print int
        la $a0, 1
        li $v0, 1
        syscall

        # print text2
        li $v0, 4
        la $a0, text2
        syscall

        # get int
        li $v0,5
        syscall

        # print int
        move $a0, $v0
        li $v0, 1
        syscall
        move $t3,$a0

        sw $t3, 0($t1) # set value
        add $t1, $t1, 4 # increment pointer
        add $t0, $t0, 1 # i++
        j Loop
    End_Loop:

    # print text4
    li $v0, 4
    la $a0, text4
    syscall

    li $t0, 0
    la $t1, array_2

    # set values 2
    Loop2: 
        bge $t0, 3, End_Loop2

        # print text1
        li $v0, 4
        la $a0, text1
        syscall

        # print int
        la $a0, 2
        li $v0, 1
        syscall

        # print text2
        li $v0, 4
        la $a0, text2
        syscall

        # get int
        li $v0,5
        syscall

        # print int
        move $a0, $v0
        li $v0, 1
        syscall
        move $t3,$a0

        sw $t3, 0($t1) # set value
        add $t1, $t1, 4 # increment pointer
        add $t0, $t0, 1 # i++
        j Loop2
    End_Loop2:

    # print text3
    li $v0, 4
    la $a0, text3
    syscall

    li $t0, 0
    la $t1, array_1
    la $t3, array_2
    la $t5, 1
    
    la $a0, 0
    Loop3:

        bge $t0, 10, End_Loop3

        # get values array 1
        lw $t2, 0($t1)
        add $t1, $t1, 4


        # get values array 2
        lw $t4, 0($t3)
        add $t3, $t3, 4

        # get product
        mul $t5, $t2, $t4
        add $a0, $a0, $t5

        # i++
        add $t0, $t0, 1
       
        j Loop3
    End_Loop3:

    li $v0, 1
    syscall

    jr $ra

