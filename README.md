# golang
Golang Learnings and Tools
#### Interfaces
Rules of Interface: 
concrete type: {map, struct, int, string, englishBot}
interface type: bot

--> interfaces are not generic type in golang as we see in c++, Java etc..
generic types was introduced in go 1.18+
## 📚 Useful Resources  

| Topic | Link |
|-------|------|
| 📝 How to Use Generics in Golang | [Read Here](https://medium.com/@samix.ys/how-to-use-generics-in-a-structs-and-interfaces-in-golang-69bd8dcbeb2d) |
| 💬 Discussion on Generics in Golang | [Join the Conversation](https://forum.golangbridge.org/t/trying-to-understand-generics-with-use-of-structs-and-interfaces/30611) |


--> interface are implicit
No "implements" keyword – Unlike Java, C++, or C#, there is no need to declare implements SomeInterface.

--> interfaces are a contract to help us manage types

--> interfaces are tough step #1 is understanding how to read them

--> make(....) vs direct declaration

```
make(...) only works with slices, maps, and channels.
```

```
For arrays, structs, pointers, use other initialization methods like new() or struct literals.
```

```
🔍 Key Differences – make(...) vs. Direct Declaration
Feature	make([]byte, size)	var b []byte (without make)
Allocates Memory?	✅ Yes	❌ No (nil slice)
Predefined Length?	✅ Yes	❌ No (must append to use)
Capacity Control?	✅ Yes (make([]byte, len, cap))	❌ No
Nil Check?	❌ Always allocated	✅ Can be nil
```







