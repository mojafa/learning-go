init functions
escape analysis stack is self cleaning
prog exec
stmt stmt idiom
,ok idiom


statement idiom
if x:=f(x);x<0 {
    return x
}else if x >z {
    return z
} else {
    return y
}

ok idiom
if x,ok:=f(x);ok {
    return x
}else if x >z {
    return z
} else {
    return y
}

func offset(tz string)int{
    if seconds,ok:=timezones[tz];ok {
        return seconds
    }
    log.Println("unknown time zone:",tz)
    return 0

}



map returns 0 value
maps are not ordered


for loops
nested loops
for range loops
map[lookup]value return
s:=map[string]int{"a":1,"b":2}
for k,v:=range s {
    fmt.Println(k,v)
}


func median(values []float64)float64{
    sort.Float64s(values)
    i := len(values)/2
    //if len(values)$1 ==1 {
        if len(values)%2==1 {
            return values[i]
        }
        v := (values[i-1]+values[i])/2
        return v
    }
}


vs = []float64{1,2,3,4,5}
fmt.Println(median(vs))


slices behave like pointers, like copy
copy of a copy

everything in go is passed by value
slice -pointer, value, copy

const type is defined when we use them in place of usage



in go we dont like panic, we like errors

const foo = boil ? 0 : 100;



//concat
s := make([]string, len(s1)+len(s2))
copy (s,s1)
copy (s[len(s1):],s2)
reiturn s
return append(s1,s2)