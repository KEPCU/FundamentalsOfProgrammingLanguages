.data
    out_string: .asciiz "\nThe Sum is: "
    out_string2: .asciiz "\nThe Average is: "
    out_string3: .asciiz "\nThe smallest is: "
    out_string4: .asciiz "\nThe largest is: "
    out_string5: .asciiz "\nEnter a number: "
    out_string6: .asciiz "\nEnter the number of numbers: "
.text
main:
    li $v0, 4
    la $a0, out_string6
    syscall

    li $v0,5
    syscall
    move $t0,$v0

    li $t1, 0 
    li $t3, 0 
    li $t4, 0 

    # 1
    li $v0, 4
    la $a0, out_string5
    syscall

    li $v0,5
    syscall

    add $t2,$t2,$v0

    add $t1, $t1, 1

    add $t3, $t3, $v0
    add $t4, $t4, $v0



    Loop:
        
        beq $t0,$t1, Exit

        li $v0, 4
        la $a0, out_string5
        syscall

        li $v0,5
        syscall

        add $t2,$t2,$v0

        add $t1, $t1, 1

        bge $t3, $v0, LABEL_IF 
        LABEL_ELSE:
            b END_LABEL_IF
        LABEL_IF:
            move $t3,$v0
        END_LABEL_IF:
            li $v0, 4

    j Loop
    Exit:
    li $v0,4
    la $a0,out_string
    syscall

    li $v0,1
    move $a0,$t2
    syscall

    # Average
    li $v0, 4
    la $a0, out_string2
    syscall

    mtc1 $t2, $f1
    cvt.s.w $f1, $f1
    mtc1 $t0, $f2
    cvt.s.w $f2, $f2
    div.s $f12, $f1, $f2

    li $v0, 2
    syscall

    # 
    li $v0, 4
    la $a0, out_string3
    syscall
    move $a0,$t3
    li $v0, 1
    syscall

    # 
    li $v0, 4
    la $a0, out_string4
    syscall

    move $a0,$t4

    li $v0, 1
    syscall

    jr $ra