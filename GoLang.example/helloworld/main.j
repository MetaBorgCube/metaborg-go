
.source 
.class public main
.super java/lang/Object

  .method public static main ([Ljava/lang/String;)V
    .limit locals 99	
    .limit stack 99	
    aload_0	
    invokevirtual java/lang/main/test ()V	
    return
  .end method

  .method public test ()V
    .limit locals 99	
    .limit stack 99	
    getstatic java/lang/System/out Ljava/io/PrintStream;	
    ldc hello world	
    invokevirtual java/io/PrintStream/println (Ljava/lang/String;)V	
    getstatic java/lang/System/out Ljava/io/PrintStream;	
    return
  .end method