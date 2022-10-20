.data
text_1: .asciiz "\nEl primero es el mayor!\n"
text_2: .asciiz "\nEl segundo es el mayor!\n"
.text
main:
li $t1, 7
li $t2, 6
bge $t1, $t2, LABEL_IF #si el primero es mayor
LABEL_ELSE:
    la $a0, text_2
    b END_LABEL_IF
LABEL_IF:
    la $a0, text_1
END_LABEL_IF:
    li $v0, 4
    syscall
jr $ra