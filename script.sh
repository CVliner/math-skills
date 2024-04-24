#!/bin/bash
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color

for ((i=1;i<100;i++)) # цикл будет длится 100 раз потом завершится или завершится если найдет неверный ответ
do
math1=$(./math-skills) # запускаем math skills
math2=$(go run . data.txt) # тут запускаете свой .go файл

echo -e "---Math-Skills---"
echo -e "$math1\n"

echo -e "-----main.go-----"
echo -e "$math2\n"

if [[ $math1 == $math2 ]]
  then
    printf "${GREEN}------EQUAL------\n\n${NC}"
    else 
    printf "${RED}----NOT-EQUAL----\n\n${NC}"
    break
fi
done
