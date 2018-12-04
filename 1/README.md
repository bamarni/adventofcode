# day 1

## part 1

    (echo "0"; cat frequencies.txt) | (tr '\n' ' '; echo) | bc

## part 2

``` bash
declare -A sum_history
sum_history["0"]=1
sum=0

while true; do
    for freq in $(cat frequencies.txt); do
        let sum+=$freq
        if [[ -n ${sum_history["$sum"]} ]]; then
            echo $sum
            exit 0
        fi
        sum_history["$sum"]=1
    done
done
```