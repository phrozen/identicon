# identicon
A colection of Identicon algorithms written in Go for fun and experimentation.

## Install

`go get -u github.com/phrozen/identicon`

## Examples

### GitHub

Close port of the GitHub Identicon algorithm (12x12)

```go
img := identicon.GitHub("phrozen") // paletted image 1bpp
```

![phrozen.png](https://github.com/user-attachments/assets/f8af4238-0862-423c-bfcb-c2f63822c0c5)

### Github Alternate

A dark alternate that uses more of the hash information for a more distinct result (16x16)

```go
img := identicon.GitHubAlternate("phrozen") // paletted image 1bpp
```

![phrozen.png](https://github.com/user-attachments/assets/6d8244ef-77fb-4504-81c4-02366f3d6ce6)
