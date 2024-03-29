#!/bin/bash
try() {
  expected="$1"
  input="$2"

  # ./9cc -> ./9gg と変更
  ./9gg "$input" > tmp.asm
  as -o tmp.o tmp.asm
  gcc -o tmp tmp.o
  ./tmp
  actual="$?"

  if [ "$actual" = "$expected" ]; then
    echo "$input => $actual"
  else
    echo "$input => $expected expected, but got $actual"
    exit 1
  fi
}

try 0 0
try 42 42
try 21 "5+20-4"

echo OK