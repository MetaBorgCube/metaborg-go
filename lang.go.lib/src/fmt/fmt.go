package fmt;

import "errors";

func Print(...interface{}) (int, error);
func Printf(string, ...interface{}) (int, error);
func Println(...interface{}) (int, error);
