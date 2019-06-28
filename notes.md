Command to search and replace imports. this usage of sed is  macOS/BSD specific and is not very portable.

```
$ sed -i '' -e 's/github.com\/hashicorp\/terraform/github.com\/hashicorp\/terraform-plugin-sdk/' $(find . -path ./vendor -prune -o -name '*.go' -print)
```

Currently sdk is private repo, you will need to force git to translate https:// requests from go modules to git over ssh

```
[url "ssh://git@github.com/"]
	insteadOf = https://github.com/
```

sdk is currently missing:

 - `helper/acctest`
 - `helper/encryption`
 - `helper/mutexkv`

Other than that the aws provider compiled!