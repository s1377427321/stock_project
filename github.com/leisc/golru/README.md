# golru

### golang LRU with sync  
### Data list back insert, front pop

# Example
```
glru, err := golru.New(128)
	if err != nil {
		fmt.Println("create lru failed")
	}

	for index := 0; index < 125; index++ {
		node := golru.NewLRUNode(index, index)

		glru.AddNode(node)
	}

	fmt.Println("size = ", glru.Size())

	for index := 200; index < 210; index++ {
		node := golru.NewLRUNode(index, index)

		glru.AddNode(node)
	}
	fmt.Println("size = ", glru.Size())
```
