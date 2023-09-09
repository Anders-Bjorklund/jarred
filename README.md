# jarred

A simple tool to simplify uploading every jar, war and ear from your local Maven repository in ~/.m2 to a remote repository.
It can upload a subset based on regex inclusion ( all files / paths looking like... ) or regex exclusion ( everything but files / paths looking like... )
Written in Go and can be compiled to a binary for your favourite os. Needs to have Maven installed on your path. 
