.data
text_1: .asciiz "\nIngrese el primer lado: "
text_2: .asciiz "\nIngrese el segundo lado: "
text_3: .asciiz "\nIngrese el tercer lado: "
text_4: .asciiz "\nEl triangulo es valido!\n"
text_5: .asciiz "\nEl triangulo no es valido!\n"
.text
main:
    li $v0, 4
    la $a0, text_1
    syscall

    li $v0,5
    syscall
    move $a0,$v0
    li $v0,1
    syscall
    move $t1,$a0

    li $v0, 4
    la $a0, text_2
    syscall

    li $v0,5
    syscall
    move $a0,$v0
    li $v0,1
    syscall
    move $t2,$a0

    li $v0, 4
    la $a0, text_3
    syscall

    li $v0,5
    syscall
    move $a0,$v0
    li $v0,1
    syscall
    move $t3,$a0

    add $t0,$t1,$t2
    bge $t0, $t3, LABEL_IF 
    LABEL_ELSE:
        la $a0, text_5
        li $v0, 4
        syscall
        jr $ra
        b END_LABEL_IF

    LABEL_IF:
        la $a0, text_4
    END_LABEL_IF:
        li $v0, 4
        syscall


    add $t0,$t1,$t3
    bge $t0, $t2, LABEL_IF2 
    LABEL_ELSE:
        la $a0, text_5
        li $v0, 4
        syscall
        jr $ra
        b END_LABEL_IF2

    LABEL_IF:
        la $a0, text_4
    END_LABEL_IF:
        li $v0, 4
        syscall


    add $t0,$t2,$t3
    bge $t0, $t1, LABEL_IF3 
    LABEL_ELSE:
        la $a0, text_5
        b END_LABEL_IF3
    LABEL_IF:
        la $a0, text_4
    END_LABEL_IF:
        li $v0, 4
        syscall

    jr $ra