
.source 
.class public main
.super java/lang/Object

  .method public <init> ()V
    aload_0	
    invokenonvirtual java/lang/Object/<init> ()V	
    return
  .end method

  .method public static main ([Ljava/lang/String;)V
    .limit locals 99	
    .limit stack 99	
    ldc hello world	
    astore 0	
    getstatic java/lang/System/out Ljava/io/PrintStream;	
    aload 0	
    invokevirtual java/io/PrintStream/println (Ljava/lang/String;)V	
    new main	
    dup	
    invokenonvirtual main/<init> ()V	
    invokevirtual main/test ()V	
    return
  .end method

  .method public test ()V
    .limit locals 99	
    .limit stack 99	
    getstatic java/lang/System/out Ljava/io/PrintStream;	
    ldc Hello 2	
    invokevirtual java/io/PrintStream/println (Ljava/lang/String;)V	
    return
  .end method