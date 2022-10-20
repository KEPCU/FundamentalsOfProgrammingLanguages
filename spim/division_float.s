.data
    msg1: .asciiz "\nIngrese un numero: "
    msg2: .asciiz "\nIngrese otro numero: "    
    msg6: .asciiz "\nLa division es: "
.text

    # imprimimos el primer mensaje
    main:
    li $v0,4
    la $a0,msg1
    syscall

    # recibimos el ingreso del usuario en $f1
    li $v0,6
    syscall
    mov.s $f1, $f0

    # imprimimos el segundo mensaje    
    li $v0,4
    la $a0,msg2
    syscall

    # recibimos el ingreso del usuario en $f2
    li $v0,6
    syscall
    mov.s $f2, $f0

    # imprimimos     
    li $v0,4
    la $a0,msg6
    syscall

    div.s $f12, $f1, $f2

    # imprimir float
    li $v0,2
    syscall
    
   
    li $v0,10
    syscall