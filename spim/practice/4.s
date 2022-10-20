.data
    out_string: .asciiz "\n no es multiplo "
    out_string1: .asciiz "\nel numero "
    out_string2: .asciiz "\n es multiplo "
    out_string4: .asciiz "\ningrese un numero:  "

.text
main:
    li $v0, 4
    la $a0, out_string4
    syscall

    li $v0,5
    syscall
    move $t0,$v0

    li $v0, 1
    syscall

    li $t1, 1 

    Loop:
        
        beq 21,$t1, Exit

        li $v0, 4 
        la $a0, out_string1 
        syscall

        li $v0, 1
        syscall

        div $t1, $t0


        # mostrar residuo
        mfhi $a0 # de aqui sacamos el residuo
        li $v0, 1
        syscall

        

        bge $a0, 0, LABEL_IF 
        LABEL_ELSE:
            la $a0, out_string
            b END_LABEL_IF
        LABEL_IF:
            la $a0, out_string2
        END_LABEL_IF:
            li $v0, 4
            syscall


        add $t1,$t1,1

    j Loop
    Exit:

    jr $ra