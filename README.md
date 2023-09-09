# jarred

A small tool to **simplify uploading jars from your local Maven repository in ~/.m2 to a remote repository** of your choice.  
  
* It can upload everything in your repository or a subset, based on regex inclusion ( all files / paths looking like... ) or regex exclusion ( everything but files / paths looking like... )  
* If you would like to see the nuber of uploads before you begin, you can **--dry-run**
* Written in Go and can be compiled to a binary for your favourite os. Needs to have Maven installed on your path. 
