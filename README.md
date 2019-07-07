# 「Go言語でつくるインタプリタ」学習リポジトリ

「Go言語でつくるインタプリタ」を読み、書籍内で定義されたMonkey言語のインタプリタを作成する。

```
$ go run main.go
Hello master! This is the Monkey programing language!
Feel free to type in commands
>> let i = 0;
>> i;
0
>> if (i > 3) {
..   puts("OK");
.. } else {
..   puts(i + 3);
.. }
4
>> while (i < 5) {
..   puts("Hello");
..   let i = i + 1;
.. }
Hello
Hello
Hello
Hello
Hello
>>
```
