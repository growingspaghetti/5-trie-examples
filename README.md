# 5-trie-examples

I'm developing a dictionary app and wanted to add prefix-search functionality to it. [my dictionary app](https://github.com/growingspaghetti/websters-1913-console-dictionary)

This repository has got 5 elementary implementations of trie. Their differences would contribute to your understanding of how the data structure evolved.

[1. playground](https://play.golang.org/p/hgfKw5y5Avv)
```bash
go run object/*
```

[2. playground](https://play.golang.org/p/K0VHxdfvpxs)
```bash
go run table/*
```

[3. playground](https://play.golang.org/p/8w3W5FEcv_i)
```bash
go run array/*
```

[4. playground](https://play.golang.org/p/wDEgeo0GBU7)
```bash
go run triple/*
```

[5. playground](https://play.golang.org/p/qry5tFlxBan)
```bash
go run double/*
```

# background

For example, you've got dictionary data saved in SQLite:

| category | word | meaning |
|:--:|:--|:--|
|1|bbc|British Broadcasting Service|
|2|cbc|Canadian Broadcasting Service|
|2|cbc|Cipher Block Chaining|
|3|cc|carbon copy|

If a user pressed `c`,  I'd like to run either
```sql
SELECT * FROM dict WHERE word LIKE 'c%'
SELECT * FROM dict WHERE category IN (2, 3)
```

Because WordNet and Webster's dictionary have 100 thousands of entries, I'd like to avoid building a trie tree in runtime.

```rust
struct Node {
    children: [Option<Box<Node>>; 27],
    id: u32,
}
```

However, since a tree is a recursive structure, its depth is arbitrary so that the content needs to be escaped to the heap space for the compiler to determine the size of an entire structure. This trie structure is dynamic and cannot be built at the compilation time.

# object-based trie

# sparse-table trie

## serializing table into a single array

# shortening the array

## triple-array trie

## double-array trie

