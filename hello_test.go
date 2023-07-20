package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		/*
			%q为  占位符、 点进去Errorf 这个函数、就可以看到，如果结果是错的，就会把got和want这两个参数，放到%q的位置拼接起来
			作用是：方便查看为什么测试结果是错的。
		*/
		t.Errorf("got '%q' want '%q'", got, want)
	}
}
