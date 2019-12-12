---
title: MSGP
category: golang
date: 2019-12-09 14:26:00
---

"MessagePack is an efficient binary serialization format. It lets you exchange data among multiple languages like JSON. But it's faster and smaller. Small integers are encoded into a single byte, and typical short strings require only one extra byte in addition to the strings themselves."

![Speed Test](/images/msgp-speed.png)

For more detail specification, you can view [MSGP update proposal V5](https://gist.github.com/frsyuki/543255).

## Why Smaller
cause JSON is a __text based protocal__, it's easy to read by human while extra format characters are necessary. There's a example.

If we define a struct __User__ in golang,
```
type User struct {
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	Age       int    `json:"age"`
	Bio       string `json:"bio"`
}

```

the JSON marshal result will be
```
{
    "name":"DL",
    "gender":"male",
    "age":1024,
    "bio":"let's dance!"
}
```
there're __23__ bytes for brace, comma and colon. While, MSGP uses a Type System to reduce these extra overhead, for example:

```
fixmap stores a map whose length is upto 15 elements
+--------+~~~~~~~~~~~~~~~~~+
|1000XXXX|   N*2 objects   |
+--------+~~~~~~~~~~~~~~~~~+
where
* XXXX is a 4-bit unsigned integer which represents N
```
What's more, there's a compression algorithm of MSGP, which means less bytes will be used. For example:

```
false:
+--------+
|  0xc2  |
+--------+

true:
+--------+
|  0xc3  |
+--------+

```

## Why Faster
We usually use cJSON library to marshal/unmarshal JSON. cJSON uses __linked list__ to store a __tree__. We need to scan each char to locate it's node one by one. With useing MSGP, we do not need to match char cause there's a __type__ and __length__ of each kind of MSGP data type.

