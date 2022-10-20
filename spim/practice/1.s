.data
    out_string: .asciiz "\nThe Sum is: "
    out_string2: .asciiz "\nThe Subtraction is: "
    out_string3: .asciiz "\nThe Multiplication is: "
    out_string4: .asciiz "\nThe Division is: "
    out_string5: .asciiz "\nThe Average is: "
.text
main:
    li $t3, 2
    li $v0, 5
    syscall
    move $t0, $v0

    li $v0, 5
    syscall
    move $t1, $v0
   
    # SUM
    li $v0, 4
    la $a0, out_string
    syscall

    add $a0, $t0, $t1

    li $v0, 1
    syscall

    # SUB
    li $v0, 4
    la $a0, out_string2
    syscall

    sub $a0, $t0, $t1

    li $v0, 1
    syscall

    # MUL
    li $v0, 4
    la $a0, out_string3
    syscall

    mul $a0, $t0, $t1

    li $v0, 1
    syscall

    # DIV
    li $v0, 4
    la $a0, out_string4
    syscall

    

    div $a0, $t0, $t1

    li $v0, 1
    syscall

    # Average
    li $v0, 4
    la $a0, out_string5
    syscall

    add $t2, $t0, $t1
    div $a0, $t2, $t3

    li $v0, 1
    syscall

    jr $ra