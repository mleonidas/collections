# collections

This is just playing around with golang 1.18beta1 
generics 

Don't use this for anything real!!! :)

### Example
```go
items := slice.From(1,2,3,4,5,5,5,6)

filtered := slice.Filter(*items, func(i int) bool {
    return i == 5
})

```
