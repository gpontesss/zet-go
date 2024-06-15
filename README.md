# zet-go

This is a remake of my shell script [zet-cmd] project, where I tried to create
a proof-of-concept Zettelkasten application. It included some basic features,
but it proved early to not be scalable; it already showed performance issues in
a small repository. Composing shell programs to accomplish complex parsing and
filtering has a lot of costs, different than having the same buffers loaded in
the application level and performing searching in it. That's the main reason I'm
rewritting it in golang. I also have goals for the feature which include
building visualizations of trees, which justifies the usage of a language with
finer control.

## What's planned

For now, these are the basic features present in [zet-cmd] that I plan to
duplicate here:

- [ ] Initializing a new repository;
- [ ] Syncing the local repository with the remote;
- [ ] Creating a new zettel;
- [x] Listing zets in a repository. (Already showed a performance improvement of
    1000%, from 150ms to 15ms);
- [ ] Editting an existing zettel;
- [ ] Querying zets by metadata. (Tags, for now.);
- [ ] Performing fuzzy-searching in the repository.

## Running it

For now, there's nothing fancy. To run it:

```sh
go run main.go
```

## Running tests

To test the application:

```sh
go test ./lib/**
```

[zet-cmd]: https://github.com/gpontesss/zet-cmd
